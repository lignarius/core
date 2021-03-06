// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package node is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	Info
	StatusParams
	SyncParams
	Block
	PushReturn
*/
package node

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

type Info struct {
	Version  string `protobuf:"bytes,1,opt,name=Version" json:"Version,omitempty"`
	Valid    bool   `protobuf:"varint,2,opt,name=Valid" json:"Valid,omitempty"`
	Length   uint64 `protobuf:"varint,3,opt,name=Length" json:"Length,omitempty"`
	LastHash []byte `protobuf:"bytes,4,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Info) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Info) GetLastHash() []byte {
	if m != nil {
		return m.LastHash
	}
	return nil
}

type StatusParams struct {
	Host string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
}

func (m *StatusParams) Reset()                    { *m = StatusParams{} }
func (m *StatusParams) String() string            { return proto.CompactTextString(m) }
func (*StatusParams) ProtoMessage()               {}
func (*StatusParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type SyncParams struct {
	LastHash []byte `protobuf:"bytes,1,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
}

func (m *SyncParams) Reset()                    { *m = SyncParams{} }
func (m *SyncParams) String() string            { return proto.CompactTextString(m) }
func (*SyncParams) ProtoMessage()               {}
func (*SyncParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SyncParams) GetLastHash() []byte {
	if m != nil {
		return m.LastHash
	}
	return nil
}

type Block struct {
	Content   string `protobuf:"bytes,1,opt,name=Content" json:"Content,omitempty"`
	Nonce     uint32 `protobuf:"varint,2,opt,name=Nonce" json:"Nonce,omitempty"`
	Previous  []byte `protobuf:"bytes,3,opt,name=Previous,proto3" json:"Previous,omitempty"`
	Type      string `protobuf:"bytes,5,opt,name=Type" json:"Type,omitempty"`
	PubKey    string `protobuf:"bytes,6,opt,name=PubKey" json:"PubKey,omitempty"`
	Date      int64  `protobuf:"varint,7,opt,name=Date" json:"Date,omitempty"`
	Signature string `protobuf:"bytes,8,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Block) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Block) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Block) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *Block) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Block) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Block) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *Block) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type PushReturn struct {
}

func (m *PushReturn) Reset()                    { *m = PushReturn{} }
func (m *PushReturn) String() string            { return proto.CompactTextString(m) }
func (*PushReturn) ProtoMessage()               {}
func (*PushReturn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Info)(nil), "Info")
	proto.RegisterType((*StatusParams)(nil), "StatusParams")
	proto.RegisterType((*SyncParams)(nil), "SyncParams")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*PushReturn)(nil), "PushReturn")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DistributionService service

type DistributionServiceClient interface {
	GetInfo(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*Info, error)
	Synchronize(ctx context.Context, in *SyncParams, opts ...grpc.CallOption) (DistributionService_SynchronizeClient, error)
	AddBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*PushReturn, error)
}

type distributionServiceClient struct {
	cc *grpc.ClientConn
}

func NewDistributionServiceClient(cc *grpc.ClientConn) DistributionServiceClient {
	return &distributionServiceClient{cc}
}

func (c *distributionServiceClient) GetInfo(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc.Invoke(ctx, "/DistributionService/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributionServiceClient) Synchronize(ctx context.Context, in *SyncParams, opts ...grpc.CallOption) (DistributionService_SynchronizeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DistributionService_serviceDesc.Streams[0], c.cc, "/DistributionService/Synchronize", opts...)
	if err != nil {
		return nil, err
	}
	x := &distributionServiceSynchronizeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributionService_SynchronizeClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type distributionServiceSynchronizeClient struct {
	grpc.ClientStream
}

func (x *distributionServiceSynchronizeClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributionServiceClient) AddBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*PushReturn, error) {
	out := new(PushReturn)
	err := grpc.Invoke(ctx, "/DistributionService/AddBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistributionService service

type DistributionServiceServer interface {
	GetInfo(context.Context, *StatusParams) (*Info, error)
	Synchronize(*SyncParams, DistributionService_SynchronizeServer) error
	AddBlock(context.Context, *Block) (*PushReturn, error)
}

func RegisterDistributionServiceServer(s *grpc.Server, srv DistributionServiceServer) {
	s.RegisterService(&_DistributionService_serviceDesc, srv)
}

func _DistributionService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributionServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributionService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributionServiceServer).GetInfo(ctx, req.(*StatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributionService_Synchronize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributionServiceServer).Synchronize(m, &distributionServiceSynchronizeServer{stream})
}

type DistributionService_SynchronizeServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type distributionServiceSynchronizeServer struct {
	grpc.ServerStream
}

func (x *distributionServiceSynchronizeServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributionService_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributionServiceServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributionService/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributionServiceServer).AddBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DistributionService",
	HandlerType: (*DistributionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _DistributionService_GetInfo_Handler,
		},
		{
			MethodName: "AddBlock",
			Handler:    _DistributionService_AddBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synchronize",
			Handler:       _DistributionService_Synchronize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0xae, 0xd3, 0x30,
	0x14, 0x84, 0x6b, 0x6e, 0x92, 0xe6, 0x9e, 0xe6, 0x6e, 0x0c, 0x42, 0x56, 0xc4, 0x22, 0xf5, 0x02,
	0x65, 0x15, 0x21, 0x78, 0x02, 0xa0, 0x12, 0x45, 0x54, 0x28, 0x72, 0x50, 0xf7, 0x6e, 0x62, 0x1a,
	0x43, 0xb1, 0x2b, 0xdb, 0xa9, 0x54, 0xb6, 0x3c, 0x11, 0x6f, 0x88, 0xec, 0xa4, 0x3f, 0x62, 0x77,
	0xbe, 0xd1, 0xd1, 0x78, 0x3c, 0x36, 0x80, 0xd2, 0x9d, 0xa8, 0x8e, 0x46, 0x3b, 0x4d, 0x7f, 0x40,
	0xf4, 0x59, 0x7d, 0xd7, 0x98, 0xc0, 0x7c, 0x2b, 0x8c, 0x95, 0x5a, 0x11, 0x54, 0xa0, 0xf2, 0x91,
	0x5d, 0x10, 0xbf, 0x80, 0x78, 0xcb, 0x0f, 0xb2, 0x23, 0xcf, 0x0a, 0x54, 0xa6, 0x6c, 0x04, 0xfc,
	0x12, 0x92, 0x8d, 0x50, 0x7b, 0xd7, 0x93, 0x87, 0x02, 0x95, 0x11, 0x9b, 0x08, 0xe7, 0x90, 0x6e,
	0xb8, 0x75, 0x6b, 0x6e, 0x7b, 0x12, 0x15, 0xa8, 0xcc, 0xd8, 0x95, 0x29, 0x85, 0xac, 0x71, 0xdc,
	0x0d, 0xb6, 0xe6, 0x86, 0xff, 0xb2, 0x18, 0x43, 0xb4, 0xd6, 0xd6, 0x4d, 0x07, 0x86, 0x99, 0x96,
	0x00, 0xcd, 0x59, 0xb5, 0xd3, 0xc6, 0xbd, 0x1b, 0xfa, 0xcf, 0xed, 0x2f, 0x82, 0xf8, 0xc3, 0x41,
	0xb7, 0x3f, 0x7d, 0xf6, 0x8f, 0x5a, 0x39, 0xa1, 0x2e, 0x56, 0x17, 0xf4, 0xd9, 0xbf, 0x6a, 0xd5,
	0x8a, 0x90, 0xfd, 0x89, 0x8d, 0xe0, 0x5d, 0x6b, 0x23, 0x4e, 0x52, 0x0f, 0x36, 0xa4, 0xcf, 0xd8,
	0x95, 0x7d, 0xa6, 0x6f, 0xe7, 0xa3, 0x20, 0xf1, 0x98, 0xc9, 0xcf, 0xfe, 0xae, 0xf5, 0xb0, 0xfb,
	0x22, 0xce, 0x24, 0x09, 0xea, 0x44, 0x7e, 0x77, 0xc5, 0x9d, 0x20, 0xf3, 0x02, 0x95, 0x0f, 0x2c,
	0xcc, 0xf8, 0x15, 0x3c, 0x36, 0x72, 0xaf, 0xb8, 0x1b, 0x8c, 0x20, 0x69, 0x58, 0xbf, 0x09, 0x34,
	0x03, 0xa8, 0x07, 0xdb, 0x33, 0xe1, 0x06, 0xa3, 0xde, 0xfe, 0x41, 0xf0, 0x7c, 0x25, 0xad, 0x33,
	0x72, 0x37, 0x38, 0xa9, 0x55, 0x23, 0xcc, 0x49, 0xb6, 0x02, 0x2f, 0x61, 0xfe, 0x49, 0xb8, 0xf0,
	0x2c, 0x4f, 0xd5, 0x7d, 0x63, 0x79, 0x5c, 0x79, 0x95, 0xce, 0xf0, 0x6b, 0x58, 0xf8, 0x9a, 0x7a,
	0xa3, 0x95, 0xfc, 0x2d, 0xf0, 0xa2, 0xba, 0x95, 0x96, 0x27, 0x55, 0xa8, 0x85, 0xce, 0xde, 0x20,
	0xbc, 0x84, 0xf4, 0x7d, 0xd7, 0x8d, 0x35, 0x4d, 0x7a, 0xbe, 0xa8, 0x6e, 0x19, 0xe8, 0x6c, 0x97,
	0x84, 0x8f, 0xf0, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x56, 0x9a, 0x1b, 0x16, 0x02,
	0x00, 0x00,
}
