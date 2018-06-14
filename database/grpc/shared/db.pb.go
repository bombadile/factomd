// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package databaseGrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DBKey is a key in the leveldb. Using a type is optional to specify what type the
// value is. E.g: 'entry', 'dblock', etc
type DBKey struct {
	KeyType              string   `protobuf:"bytes,1,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	Bucket               []byte   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBKey) Reset()         { *m = DBKey{} }
func (m *DBKey) String() string { return proto.CompactTextString(m) }
func (*DBKey) ProtoMessage()    {}
func (*DBKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_fa4dc8a9d7778159, []int{0}
}
func (m *DBKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBKey.Unmarshal(m, b)
}
func (m *DBKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBKey.Marshal(b, m, deterministic)
}
func (dst *DBKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBKey.Merge(dst, src)
}
func (m *DBKey) XXX_Size() int {
	return xxx_messageInfo_DBKey.Size(m)
}
func (m *DBKey) XXX_DiscardUnknown() {
	xxx_messageInfo_DBKey.DiscardUnknown(m)
}

var xxx_messageInfo_DBKey proto.InternalMessageInfo

func (m *DBKey) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

func (m *DBKey) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *DBKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// DBValue is a returned value from a database access
type DBValue struct {
	ValType              string   `protobuf:"bytes,1,opt,name=val_type,json=valType,proto3" json:"val_type,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ContainedIn          []byte   `protobuf:"bytes,3,opt,name=containedIn,proto3" json:"containedIn,omitempty"`
	Sequence             int32    `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Error                string   `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBValue) Reset()         { *m = DBValue{} }
func (m *DBValue) String() string { return proto.CompactTextString(m) }
func (*DBValue) ProtoMessage()    {}
func (*DBValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_fa4dc8a9d7778159, []int{1}
}
func (m *DBValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBValue.Unmarshal(m, b)
}
func (m *DBValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBValue.Marshal(b, m, deterministic)
}
func (dst *DBValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBValue.Merge(dst, src)
}
func (m *DBValue) XXX_Size() int {
	return xxx_messageInfo_DBValue.Size(m)
}
func (m *DBValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DBValue.DiscardUnknown(m)
}

var xxx_messageInfo_DBValue proto.InternalMessageInfo

func (m *DBValue) GetValType() string {
	if m != nil {
		return m.ValType
	}
	return ""
}

func (m *DBValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DBValue) GetContainedIn() []byte {
	if m != nil {
		return m.ContainedIn
	}
	return nil
}

func (m *DBValue) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *DBValue) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*DBKey)(nil), "databaseGrpc.DBKey")
	proto.RegisterType((*DBValue)(nil), "databaseGrpc.DBValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseGrpcClient is the client API for DatabaseGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseGrpcClient interface {
	// A simple RPC.
	//
	// Searches the DB for the key and returns the value found.
	Retrieve(ctx context.Context, in *DBKey, opts ...grpc.CallOption) (*DBValue, error)
	// A server-to-client streaming RPC.
	//
	// Retrieves the entries contained by the dbkey as a stream
	RetrieveAllEntries(ctx context.Context, in *DBKey, opts ...grpc.CallOption) (DatabaseGrpc_RetrieveAllEntriesClient, error)
}

type databaseGrpcClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseGrpcClient(cc *grpc.ClientConn) DatabaseGrpcClient {
	return &databaseGrpcClient{cc}
}

func (c *databaseGrpcClient) Retrieve(ctx context.Context, in *DBKey, opts ...grpc.CallOption) (*DBValue, error) {
	out := new(DBValue)
	err := c.cc.Invoke(ctx, "/databaseGrpc.databaseGrpc/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseGrpcClient) RetrieveAllEntries(ctx context.Context, in *DBKey, opts ...grpc.CallOption) (DatabaseGrpc_RetrieveAllEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DatabaseGrpc_serviceDesc.Streams[0], "/databaseGrpc.databaseGrpc/RetrieveAllEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &databaseGrpcRetrieveAllEntriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DatabaseGrpc_RetrieveAllEntriesClient interface {
	Recv() (*DBValue, error)
	grpc.ClientStream
}

type databaseGrpcRetrieveAllEntriesClient struct {
	grpc.ClientStream
}

func (x *databaseGrpcRetrieveAllEntriesClient) Recv() (*DBValue, error) {
	m := new(DBValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DatabaseGrpcServer is the server API for DatabaseGrpc service.
type DatabaseGrpcServer interface {
	// A simple RPC.
	//
	// Searches the DB for the key and returns the value found.
	Retrieve(context.Context, *DBKey) (*DBValue, error)
	// A server-to-client streaming RPC.
	//
	// Retrieves the entries contained by the dbkey as a stream
	RetrieveAllEntries(*DBKey, DatabaseGrpc_RetrieveAllEntriesServer) error
}

func RegisterDatabaseGrpcServer(s *grpc.Server, srv DatabaseGrpcServer) {
	s.RegisterService(&_DatabaseGrpc_serviceDesc, srv)
}

func _DatabaseGrpc_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGrpcServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.databaseGrpc/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGrpcServer).Retrieve(ctx, req.(*DBKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseGrpc_RetrieveAllEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DatabaseGrpcServer).RetrieveAllEntries(m, &databaseGrpcRetrieveAllEntriesServer{stream})
}

type DatabaseGrpc_RetrieveAllEntriesServer interface {
	Send(*DBValue) error
	grpc.ServerStream
}

type databaseGrpcRetrieveAllEntriesServer struct {
	grpc.ServerStream
}

func (x *databaseGrpcRetrieveAllEntriesServer) Send(m *DBValue) error {
	return x.ServerStream.SendMsg(m)
}

var _DatabaseGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "databaseGrpc.databaseGrpc",
	HandlerType: (*DatabaseGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Retrieve",
			Handler:    _DatabaseGrpc_Retrieve_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveAllEntries",
			Handler:       _DatabaseGrpc_RetrieveAllEntries_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "db.proto",
}

func init() { proto.RegisterFile("db.proto", fileDescriptor_db_fa4dc8a9d7778159) }

var fileDescriptor_db_fa4dc8a9d7778159 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x31, 0x25, 0x6d, 0x38, 0x3a, 0xa0, 0xe3, 0x43, 0xa6, 0x53, 0x94, 0x29, 0x53, 0x84,
	0x60, 0x61, 0xa5, 0x0a, 0x42, 0x08, 0xa6, 0x08, 0xb1, 0x22, 0x27, 0xb9, 0xa1, 0x8a, 0xe5, 0x04,
	0xc7, 0xb1, 0xe4, 0x3f, 0xd1, 0xdf, 0x8c, 0xec, 0x06, 0xd4, 0xb2, 0xb1, 0xdd, 0xf3, 0xda, 0x7e,
	0xac, 0xd7, 0x86, 0xb8, 0xa9, 0xf2, 0x5e, 0x77, 0xa6, 0xc3, 0x65, 0x23, 0x8c, 0xa8, 0xc4, 0x40,
	0xcf, 0xba, 0xaf, 0xd3, 0x37, 0x88, 0x8a, 0xf5, 0x2b, 0x39, 0xbc, 0x81, 0xb8, 0x25, 0xf7, 0x69,
	0x5c, 0x4f, 0x9c, 0x25, 0x2c, 0x3b, 0x2d, 0x17, 0x2d, 0xb9, 0x77, 0xd7, 0x13, 0x5e, 0xc3, 0xbc,
	0x1a, 0xeb, 0x96, 0x0c, 0x3f, 0x4e, 0x58, 0xb6, 0x2c, 0x27, 0xc2, 0x73, 0x98, 0xb5, 0xe4, 0xf8,
	0x2c, 0x84, 0x7e, 0x4c, 0xb7, 0x0c, 0x16, 0xc5, 0xfa, 0x43, 0xc8, 0x91, 0xbc, 0xd0, 0x0a, 0x79,
	0x20, 0xb4, 0x42, 0x06, 0xe1, 0x25, 0x44, 0xd6, 0xef, 0x99, 0x7c, 0x3b, 0xc0, 0x04, 0xce, 0xea,
	0x4e, 0x19, 0xb1, 0x51, 0xd4, 0xbc, 0xa8, 0x49, 0xbb, 0x1f, 0xe1, 0x0a, 0xe2, 0x81, 0xbe, 0x46,
	0x52, 0x35, 0xf1, 0x93, 0x84, 0x65, 0x51, 0xf9, 0xcb, 0xde, 0x49, 0x5a, 0x77, 0x9a, 0x47, 0xe1,
	0xae, 0x1d, 0xdc, 0x6d, 0x19, 0x1c, 0xf4, 0xc5, 0x07, 0x88, 0x4b, 0x32, 0x7a, 0x43, 0x96, 0xf0,
	0x22, 0xdf, 0x5f, 0xca, 0xc3, 0x3b, 0xac, 0xae, 0xfe, 0x86, 0xa1, 0x4d, 0x7a, 0x84, 0x05, 0xe0,
	0xcf, 0xc9, 0x47, 0x29, 0x9f, 0x94, 0x1f, 0x87, 0xff, 0x39, 0x6e, 0x59, 0x35, 0x0f, 0x9f, 0x70,
	0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x08, 0x66, 0xa0, 0x1c, 0x90, 0x01, 0x00, 0x00,
}
