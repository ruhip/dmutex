// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dsync/dsync.proto

package dsync

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

// Request the distributed lock
type LockReq struct {
	Node                 string               `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Tstmp                *timestamp.Timestamp `protobuf:"bytes,2,opt,name=tstmp,proto3" json:"tstmp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LockReq) Reset()         { *m = LockReq{} }
func (m *LockReq) String() string { return proto.CompactTextString(m) }
func (*LockReq) ProtoMessage()    {}
func (*LockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_71eb3d5f878593d6, []int{0}
}

func (m *LockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockReq.Unmarshal(m, b)
}
func (m *LockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockReq.Marshal(b, m, deterministic)
}
func (m *LockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockReq.Merge(m, src)
}
func (m *LockReq) XXX_Size() int {
	return xxx_messageInfo_LockReq.Size(m)
}
func (m *LockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LockReq.DiscardUnknown(m)
}

var xxx_messageInfo_LockReq proto.InternalMessageInfo

func (m *LockReq) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *LockReq) GetTstmp() *timestamp.Timestamp {
	if m != nil {
		return m.Tstmp
	}
	return nil
}

// for Reply, Inquire, and Relinquish messages
type Node struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_71eb3d5f878593d6, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

// Response to an Inquire message.  May either yield or relinquish
type InquireReply struct {
	Yield                bool     `protobuf:"varint,1,opt,name=yield,proto3" json:"yield,omitempty"`
	Relinquish           bool     `protobuf:"varint,2,opt,name=relinquish,proto3" json:"relinquish,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InquireReply) Reset()         { *m = InquireReply{} }
func (m *InquireReply) String() string { return proto.CompactTextString(m) }
func (*InquireReply) ProtoMessage()    {}
func (*InquireReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71eb3d5f878593d6, []int{2}
}

func (m *InquireReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InquireReply.Unmarshal(m, b)
}
func (m *InquireReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InquireReply.Marshal(b, m, deterministic)
}
func (m *InquireReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InquireReply.Merge(m, src)
}
func (m *InquireReply) XXX_Size() int {
	return xxx_messageInfo_InquireReply.Size(m)
}
func (m *InquireReply) XXX_DiscardUnknown() {
	xxx_messageInfo_InquireReply.DiscardUnknown(m)
}

var xxx_messageInfo_InquireReply proto.InternalMessageInfo

func (m *InquireReply) GetYield() bool {
	if m != nil {
		return m.Yield
	}
	return false
}

func (m *InquireReply) GetRelinquish() bool {
	if m != nil {
		return m.Relinquish
	}
	return false
}

// Response to a Validate message.  Return True if lock is still being held
type ValidateReply struct {
	Holding              bool     `protobuf:"varint,1,opt,name=holding,proto3" json:"holding,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateReply) Reset()         { *m = ValidateReply{} }
func (m *ValidateReply) String() string { return proto.CompactTextString(m) }
func (*ValidateReply) ProtoMessage()    {}
func (*ValidateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_71eb3d5f878593d6, []int{3}
}

func (m *ValidateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateReply.Unmarshal(m, b)
}
func (m *ValidateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateReply.Marshal(b, m, deterministic)
}
func (m *ValidateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateReply.Merge(m, src)
}
func (m *ValidateReply) XXX_Size() int {
	return xxx_messageInfo_ValidateReply.Size(m)
}
func (m *ValidateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateReply.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateReply proto.InternalMessageInfo

func (m *ValidateReply) GetHolding() bool {
	if m != nil {
		return m.Holding
	}
	return false
}

func init() {
	proto.RegisterType((*LockReq)(nil), "dsync.LockReq")
	proto.RegisterType((*Node)(nil), "dsync.Node")
	proto.RegisterType((*InquireReply)(nil), "dsync.InquireReply")
	proto.RegisterType((*ValidateReply)(nil), "dsync.ValidateReply")
}

func init() { proto.RegisterFile("dsync/dsync.proto", fileDescriptor_71eb3d5f878593d6) }

var fileDescriptor_71eb3d5f878593d6 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0x87, 0x19, 0x82, 0x6f, 0x6a, 0x62, 0xdd, 0x61, 0xe1, 0xa0, 0x4b, 0xf5, 0x80, 0x07,
	0x8b, 0x99, 0xdf, 0x60, 0xd9, 0xc5, 0xc4, 0xe8, 0x52, 0x8d, 0xf7, 0x0d, 0x2a, 0x6b, 0x2c, 0x94,
	0xd1, 0x92, 0xc8, 0xe7, 0xf5, 0x8b, 0x18, 0xda, 0x62, 0x36, 0x63, 0xbc, 0x40, 0xff, 0x2f, 0xbf,
	0xfe, 0xfb, 0xfe, 0xef, 0xc1, 0x59, 0xa6, 0xda, 0x32, 0x4d, 0xcc, 0x97, 0x54, 0xb5, 0xd4, 0x12,
	0xf9, 0x46, 0x44, 0x97, 0xb9, 0x94, 0xb9, 0x60, 0x89, 0x29, 0xae, 0x9b, 0xf7, 0x44, 0xf3, 0x82,
	0x29, 0xbd, 0x2a, 0x2a, 0xcb, 0xe1, 0x67, 0x08, 0x1e, 0x65, 0xfa, 0x41, 0xd9, 0x16, 0x21, 0x18,
	0x96, 0x32, 0x63, 0x13, 0x6f, 0xea, 0xc5, 0x47, 0xd4, 0x9c, 0xd1, 0x1d, 0xf8, 0x5a, 0xe9, 0xa2,
	0x9a, 0x1c, 0x4c, 0xbd, 0x78, 0x34, 0x8b, 0x88, 0xf5, 0x23, 0xbd, 0x1f, 0x79, 0xed, 0xfd, 0xa8,
	0x05, 0x71, 0x04, 0xc3, 0xa7, 0xee, 0xe6, 0x1f, 0x6e, 0x78, 0x01, 0xc7, 0x0f, 0xe5, 0xb6, 0xe1,
	0x35, 0xa3, 0xac, 0x12, 0x2d, 0x1a, 0x83, 0xdf, 0x72, 0x26, 0x32, 0x03, 0x85, 0xd4, 0x0a, 0x74,
	0x01, 0x50, 0x33, 0xc1, 0x3b, 0x50, 0x6d, 0xcc, 0xc3, 0x21, 0xdd, 0xa9, 0xe0, 0x1b, 0x38, 0x79,
	0x5b, 0x09, 0x9e, 0xad, 0xb4, 0xb3, 0x99, 0x40, 0xb0, 0x91, 0x22, 0xe3, 0x65, 0xee, 0x8c, 0x7a,
	0x39, 0xfb, 0xf2, 0x20, 0x5c, 0x70, 0xa5, 0x5f, 0xda, 0x32, 0x45, 0x31, 0x04, 0x94, 0x6d, 0x1b,
	0xa6, 0x34, 0x3a, 0x25, 0x76, 0x56, 0x2e, 0x7a, 0x34, 0x72, 0xba, 0xeb, 0x1c, 0x0f, 0xd0, 0x15,
	0xf8, 0xd6, 0x79, 0xb7, 0xfe, 0x1b, 0x8a, 0x01, 0xe8, 0x4f, 0x53, 0xff, 0x92, 0xb7, 0x10, 0xb8,
	0xd8, 0xfb, 0xd8, 0xb9, 0x13, 0xbb, 0x33, 0xc1, 0x03, 0x94, 0x40, 0xd8, 0xe7, 0xdb, 0xe7, 0xc7,
	0x4e, 0xec, 0xa5, 0xc7, 0x83, 0xf9, 0x35, 0x8c, 0xb9, 0x24, 0x79, 0x5d, 0xa5, 0x24, 0x2b, 0x1a,
	0xcd, 0x3e, 0x2d, 0x37, 0x87, 0x45, 0xf7, 0x5b, 0x76, 0xab, 0x5a, 0x7a, 0xeb, 0x43, 0xb3, 0xb3,
	0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xf4, 0x73, 0x97, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DistSyncClient is the client API for DistSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DistSyncClient interface {
	// Request for Distributed Lock
	Request(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*Node, error)
	// Request for Distributed Lock
	Reply(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error)
	// Relinquish Distributed Lock
	Relinquish(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error)
	// Inquire on status of Distributed Lock from head of queue
	Inquire(ctx context.Context, in *Node, opts ...grpc.CallOption) (*InquireReply, error)
	// Validate health of node holding Distributed Lock
	Validate(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ValidateReply, error)
}

type distSyncClient struct {
	cc *grpc.ClientConn
}

func NewDistSyncClient(cc *grpc.ClientConn) DistSyncClient {
	return &distSyncClient{cc}
}

func (c *distSyncClient) Request(ctx context.Context, in *LockReq, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/dsync.DistSync/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distSyncClient) Reply(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/dsync.DistSync/Reply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distSyncClient) Relinquish(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/dsync.DistSync/Relinquish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distSyncClient) Inquire(ctx context.Context, in *Node, opts ...grpc.CallOption) (*InquireReply, error) {
	out := new(InquireReply)
	err := c.cc.Invoke(ctx, "/dsync.DistSync/Inquire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distSyncClient) Validate(ctx context.Context, in *Node, opts ...grpc.CallOption) (*ValidateReply, error) {
	out := new(ValidateReply)
	err := c.cc.Invoke(ctx, "/dsync.DistSync/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistSyncServer is the server API for DistSync service.
type DistSyncServer interface {
	// Request for Distributed Lock
	Request(context.Context, *LockReq) (*Node, error)
	// Request for Distributed Lock
	Reply(context.Context, *Node) (*Node, error)
	// Relinquish Distributed Lock
	Relinquish(context.Context, *Node) (*Node, error)
	// Inquire on status of Distributed Lock from head of queue
	Inquire(context.Context, *Node) (*InquireReply, error)
	// Validate health of node holding Distributed Lock
	Validate(context.Context, *Node) (*ValidateReply, error)
}

func RegisterDistSyncServer(s *grpc.Server, srv DistSyncServer) {
	s.RegisterService(&_DistSync_serviceDesc, srv)
}

func _DistSync_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSyncServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsync.DistSync/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSyncServer).Request(ctx, req.(*LockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistSync_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSyncServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsync.DistSync/Reply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSyncServer).Reply(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistSync_Relinquish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSyncServer).Relinquish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsync.DistSync/Relinquish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSyncServer).Relinquish(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistSync_Inquire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSyncServer).Inquire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsync.DistSync/Inquire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSyncServer).Inquire(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistSync_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSyncServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dsync.DistSync/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSyncServer).Validate(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dsync.DistSync",
	HandlerType: (*DistSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _DistSync_Request_Handler,
		},
		{
			MethodName: "Reply",
			Handler:    _DistSync_Reply_Handler,
		},
		{
			MethodName: "Relinquish",
			Handler:    _DistSync_Relinquish_Handler,
		},
		{
			MethodName: "Inquire",
			Handler:    _DistSync_Inquire_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _DistSync_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dsync/dsync.proto",
}
