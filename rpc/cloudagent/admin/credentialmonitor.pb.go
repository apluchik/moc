// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credentialmonitor.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CertificateStatus int32

const (
	CertificateStatus_Single  CertificateStatus = 0
	CertificateStatus_Overlap CertificateStatus = 1
)

var CertificateStatus_name = map[int32]string{
	0: "Single",
	1: "Overlap",
}

var CertificateStatus_value = map[string]int32{
	"Single":  0,
	"Overlap": 1,
}

func (x CertificateStatus) String() string {
	return proto.EnumName(CertificateStatus_name, int32(x))
}

func (CertificateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{0}
}

type CredentialMonitorRequest struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorRequest) Reset()         { *m = CredentialMonitorRequest{} }
func (m *CredentialMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorRequest) ProtoMessage()    {}
func (*CredentialMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{0}
}

func (m *CredentialMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorRequest.Unmarshal(m, b)
}
func (m *CredentialMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorRequest.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorRequest.Merge(m, src)
}
func (m *CredentialMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorRequest.Size(m)
}
func (m *CredentialMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorRequest proto.InternalMessageInfo

func (m *CredentialMonitorRequest) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

type CredentialMonitorResponse struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	Error                string             `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorResponse) Reset()         { *m = CredentialMonitorResponse{} }
func (m *CredentialMonitorResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorResponse) ProtoMessage()    {}
func (*CredentialMonitorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{1}
}

func (m *CredentialMonitorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorResponse.Unmarshal(m, b)
}
func (m *CredentialMonitorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorResponse.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorResponse.Merge(m, src)
}
func (m *CredentialMonitorResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorResponse.Size(m)
}
func (m *CredentialMonitorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorResponse proto.InternalMessageInfo

func (m *CredentialMonitorResponse) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

func (m *CredentialMonitorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CredentialMonitor struct {
	Certificate          []byte            `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               CertificateStatus `protobuf:"varint,2,opt,name=status,proto3,enum=moc.cloudagent.admin.CertificateStatus" json:"status,omitempty"`
	NewCertificate       string            `protobuf:"bytes,3,opt,name=newCertificate,proto3" json:"newCertificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CredentialMonitor) Reset()         { *m = CredentialMonitor{} }
func (m *CredentialMonitor) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitor) ProtoMessage()    {}
func (*CredentialMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{2}
}

func (m *CredentialMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitor.Unmarshal(m, b)
}
func (m *CredentialMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitor.Marshal(b, m, deterministic)
}
func (m *CredentialMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitor.Merge(m, src)
}
func (m *CredentialMonitor) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitor.Size(m)
}
func (m *CredentialMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitor proto.InternalMessageInfo

func (m *CredentialMonitor) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *CredentialMonitor) GetStatus() CertificateStatus {
	if m != nil {
		return m.Status
	}
	return CertificateStatus_Single
}

func (m *CredentialMonitor) GetNewCertificate() string {
	if m != nil {
		return m.NewCertificate
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.cloudagent.admin.CertificateStatus", CertificateStatus_name, CertificateStatus_value)
	proto.RegisterType((*CredentialMonitorRequest)(nil), "moc.cloudagent.admin.CredentialMonitorRequest")
	proto.RegisterType((*CredentialMonitorResponse)(nil), "moc.cloudagent.admin.CredentialMonitorResponse")
	proto.RegisterType((*CredentialMonitor)(nil), "moc.cloudagent.admin.CredentialMonitor")
}

func init() { proto.RegisterFile("credentialmonitor.proto", fileDescriptor_0e30998b3c2b899e) }

var fileDescriptor_0e30998b3c2b899e = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0xff, 0xf2, 0xaf, 0x38, 0x95, 0xd2, 0x2e, 0x45, 0xa3, 0xa7, 0x92, 0x83, 0x16,
	0xd1, 0x5d, 0xa8, 0x0f, 0x20, 0x2a, 0xe2, 0x49, 0x84, 0x14, 0x2f, 0xde, 0xd2, 0xed, 0xb4, 0x5d,
	0x48, 0x76, 0xd2, 0xcd, 0x44, 0xef, 0x9e, 0x7c, 0x0a, 0x9f, 0x55, 0x4c, 0xc4, 0x16, 0x53, 0xa1,
	0x17, 0x8f, 0x3b, 0x33, 0xdf, 0xef, 0xfb, 0x76, 0x77, 0xe0, 0xc0, 0x78, 0x9c, 0xa2, 0x63, 0x1b,
	0x27, 0x29, 0x39, 0xcb, 0xe4, 0x55, 0xe6, 0x89, 0x49, 0xf6, 0x53, 0x32, 0xca, 0x24, 0x54, 0x4c,
	0xe3, 0x39, 0x3a, 0x56, 0xf1, 0x34, 0xb5, 0x2e, 0x5c, 0x42, 0x70, 0xf3, 0x2d, 0xb8, 0xaf, 0x04,
	0x11, 0x2e, 0x0b, 0xcc, 0x59, 0x3e, 0x42, 0xaf, 0xd6, 0x0b, 0xc4, 0x40, 0x0c, 0xdb, 0xa3, 0x13,
	0xb5, 0x89, 0xa6, 0xea, 0xa8, 0x3a, 0x21, 0x7c, 0x13, 0x70, 0xb8, 0xc1, 0x33, 0xcf, 0xc8, 0xe5,
	0xf8, 0x47, 0xa6, 0xb2, 0x0f, 0xff, 0x6f, 0xbd, 0x27, 0x1f, 0xfc, 0x1b, 0x88, 0xe1, 0x6e, 0x54,
	0x1d, 0xc2, 0x77, 0xb1, 0xc1, 0x4d, 0x0e, 0xa0, 0x6d, 0xd0, 0xb3, 0x9d, 0x59, 0x13, 0x33, 0x96,
	0xe6, 0x7b, 0xd1, 0x7a, 0x49, 0x5e, 0x42, 0x2b, 0xe7, 0x98, 0x8b, 0xbc, 0xc4, 0x75, 0x7e, 0x4d,
	0xb6, 0x92, 0x8c, 0xcb, 0xf1, 0xe8, 0x4b, 0x26, 0x8f, 0xa1, 0xe3, 0xf0, 0x65, 0xad, 0x1f, 0x34,
	0xcb, 0x5c, 0x3f, 0xaa, 0xa7, 0x67, 0xd0, 0xab, 0x41, 0x24, 0x40, 0x6b, 0x6c, 0xdd, 0x3c, 0xc1,
	0x6e, 0x43, 0xb6, 0x61, 0xe7, 0xe1, 0x19, 0x7d, 0x12, 0x67, 0x5d, 0x31, 0x7a, 0x15, 0xb0, 0x5f,
	0xbb, 0xce, 0xd5, 0x67, 0x20, 0xb9, 0x80, 0xe6, 0x1d, 0xb2, 0x54, 0xdb, 0x3e, 0x61, 0xb5, 0x02,
	0x47, 0x7a, 0xeb, 0xf9, 0xea, 0xfb, 0xc2, 0xc6, 0xb5, 0x7e, 0x3a, 0x9f, 0x5b, 0x5e, 0x14, 0x13,
	0x65, 0x28, 0xd5, 0xa9, 0x35, 0x9e, 0x72, 0x9a, 0xb1, 0x4e, 0xc9, 0x68, 0x9f, 0x19, 0xbd, 0x82,
	0xe9, 0x12, 0x36, 0x69, 0x95, 0xfb, 0x79, 0xf1, 0x11, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xf2, 0x49,
	0xaf, 0xba, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialMonitorAgentClient is the client API for CredentialMonitorAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialMonitorAgentClient interface {
	Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error)
}

type credentialMonitorAgentClient struct {
	cc *grpc.ClientConn
}

func NewCredentialMonitorAgentClient(cc *grpc.ClientConn) CredentialMonitorAgentClient {
	return &credentialMonitorAgentClient{cc}
}

func (c *credentialMonitorAgentClient) Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error) {
	out := new(CredentialMonitorResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.admin.CredentialMonitorAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialMonitorAgentServer is the server API for CredentialMonitorAgent service.
type CredentialMonitorAgentServer interface {
	Get(context.Context, *CredentialMonitorRequest) (*CredentialMonitorResponse, error)
}

// UnimplementedCredentialMonitorAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialMonitorAgentServer struct {
}

func (*UnimplementedCredentialMonitorAgentServer) Get(ctx context.Context, req *CredentialMonitorRequest) (*CredentialMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCredentialMonitorAgentServer(s *grpc.Server, srv CredentialMonitorAgentServer) {
	s.RegisterService(&_CredentialMonitorAgent_serviceDesc, srv)
}

func _CredentialMonitorAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialMonitorAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.admin.CredentialMonitorAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialMonitorAgentServer).Get(ctx, req.(*CredentialMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialMonitorAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.admin.CredentialMonitorAgent",
	HandlerType: (*CredentialMonitorAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CredentialMonitorAgent_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentialmonitor.proto",
}
