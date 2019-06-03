package state

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
)

// Toggle this var to swap Old/New Holding
var UseDependentHolding = true

// This hold a slice of messages dependent on a hash
type HoldingList struct {
	holding    map[[32]byte][]interfaces.IMsg
	s          *State // for debug logging
	dependents map[[32]byte]bool
}

func (l *HoldingList) Init(s *State) {
	l.holding = make(map[[32]byte][]interfaces.IMsg)
	l.s = s
	l.dependents = make(map[[32]byte]bool)
}

func (l *HoldingList) Messages() map[[32]byte][]interfaces.IMsg {
	return l.holding
}

func (l *HoldingList) GetSize() int {
	return len(l.dependents)
}

func (l *HoldingList) Exists(h [32]byte) bool {
	return l.dependents[h]
}

// Add a message to a dependent holding list
func (l *HoldingList) Add(h [32]byte, msg interfaces.IMsg) bool {

	if l.dependents[msg.GetMsgHash().Fixed()] {
		return false
	}

	if l.holding[h] == nil {
		l.holding[h] = []interfaces.IMsg{msg}
	} else {
		l.holding[h] = append(l.holding[h], msg)
	}

	l.dependents[msg.GetMsgHash().Fixed()] = true
	return true
}

// get and remove the list of dependent message for a hash
func (l *HoldingList) Get(h [32]byte) []interfaces.IMsg {
	rval := l.holding[h]
	delete(l.holding, h)

	for _, msg := range rval {
		l.s.LogMessage("newHolding", "DequeueFromDependantHolding()", msg)
		delete(l.dependents, msg.GetMsgHash().Fixed())
	}
	return rval
}

// clean stale messages from holding
func (l *HoldingList) Review() {
	for h := range l.holding {
		dh := l.holding[h]
		if nil == l {
			continue
		}
		for _, msg := range dh {
			if l.s.IsMsgStale(msg) < 0 {
				l.Get(h) // remove from holding
				l.s.LogMessage("newHolding", "RemoveFromDependantHolding()", msg)
				continue
			}
		}
	}
}

// Add a message to a dependent holding list
func (s *State) Add(h [32]byte, msg interfaces.IMsg) int {
	if msg == nil {
		panic("Empty Message")
	}

	if !UseDependentHolding {
		return 0
	}

	if s.Hold.Add(h, msg) {
		// return negative value so message is marked invalid and not processed in standard holding
		s.LogMessage("newHolding", fmt.Sprintf("AddToDependantHolding(%x)", h[:4]), msg)
		s.LogMessage("newHolding", fmt.Sprintf("DUP_AddToDependantHolding(%x)", h[:4]), msg)
	}

	return -2
}

// get and remove the list of dependent message for a hash
func (s *State) Get(h [32]byte) []interfaces.IMsg {
	return s.Hold.Get(h)
}

// Execute a list of messages from holding that are dependant on a hash
// the hash may be a EC address or a CainID or a height (ok heights are not really hashes but we cheat on that)
func (s *State) ExecuteFromHolding(h [32]byte) {
	// get the list of messages waiting on this hash
	l := s.Get(h)
	if l != nil {
		s.LogPrintf("newHolding", "ExecuteFromDependantHolding(%x)[%d]", len(l), h[:4])
		// add the messages to the msgQueue so they get executed as space is available
		for _, m := range l {
			s.LogPrintf("newHolding", "DeleteFromDependantHolding M-%x", m.GetMsgHash().Bytes()[:3])
		}

		go func() {
			for _, m := range l {
				s.LogMessage("msgQueue", "enqueue_from_dependant_holding", m)
				s.msgQueue <- m
			}
		}()
	} else {
		s.LogPrintf("newHolding", "ExecuteFromDependantHolding(%x) nothing waiting", h[:4])
	}
}

// put a height in the first 4 bytes of a hash so we can use it to look up dependent message in holding
func HeightToHash(height uint32) [32]byte {
	var h [32]byte
	h[0] = byte((height >> 24) & 0xFF)
	h[1] = byte((height >> 16) & 0xFF)
	h[2] = byte((height >> 8) & 0xFF)
	h[3] = byte((height >> 0) & 0xFF)
	return h
}