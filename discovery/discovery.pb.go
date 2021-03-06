// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery.proto

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	discovery.proto

It has these top-level messages:
	RegisterRequest
	DeRegisterRequest
	HeartBeatRequest
	LocationRequest
	StatusRequest
	DiscoveryResponse
	ServiceLocation
*/
package discovery

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

type ServiceStatus int32

const (
	ServiceStatus_OFFLINE              ServiceStatus = 0
	ServiceStatus_ONLINE               ServiceStatus = 1
	ServiceStatus_DEGRADED_PERFORMANCE ServiceStatus = 2
	ServiceStatus_PARTIAL_OUTAGE       ServiceStatus = 3
	ServiceStatus_UNDER_MAINTENANCE    ServiceStatus = 5
)

var ServiceStatus_name = map[int32]string{
	0: "OFFLINE",
	1: "ONLINE",
	2: "DEGRADED_PERFORMANCE",
	3: "PARTIAL_OUTAGE",
	5: "UNDER_MAINTENANCE",
}
var ServiceStatus_value = map[string]int32{
	"OFFLINE":              0,
	"ONLINE":               1,
	"DEGRADED_PERFORMANCE": 2,
	"PARTIAL_OUTAGE":       3,
	"UNDER_MAINTENANCE":    5,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}
func (ServiceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AppRelease int32

const (
	AppRelease_STABLE    AppRelease = 0
	AppRelease_BETA      AppRelease = 1
	AppRelease_ALPHA     AppRelease = 2
	AppRelease_PRE_ALPHA AppRelease = 3
)

var AppRelease_name = map[int32]string{
	0: "STABLE",
	1: "BETA",
	2: "ALPHA",
	3: "PRE_ALPHA",
}
var AppRelease_value = map[string]int32{
	"STABLE":    0,
	"BETA":      1,
	"ALPHA":     2,
	"PRE_ALPHA": 3,
}

func (x AppRelease) String() string {
	return proto.EnumName(AppRelease_name, int32(x))
}
func (AppRelease) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusTarget int32

const (
	StatusTarget_INSTANCE StatusTarget = 0
	StatusTarget_SERVICE  StatusTarget = 1
	StatusTarget_BOTH     StatusTarget = 2
)

var StatusTarget_name = map[int32]string{
	0: "INSTANCE",
	1: "SERVICE",
	2: "BOTH",
}
var StatusTarget_value = map[string]int32{
	"INSTANCE": 0,
	"SERVICE":  1,
	"BOTH":     2,
}

func (x StatusTarget) String() string {
	return proto.EnumName(StatusTarget_name, int32(x))
}
func (StatusTarget) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type RegisterRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	ServiceHost  string     `protobuf:"bytes,3,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort  int32      `protobuf:"varint,4,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Release      AppRelease `protobuf:"varint,5,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RegisterRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *RegisterRequest) GetServiceHost() string {
	if m != nil {
		return m.ServiceHost
	}
	return ""
}

func (m *RegisterRequest) GetServicePort() int32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *RegisterRequest) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

type DeRegisterRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Release      AppRelease `protobuf:"varint,3,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *DeRegisterRequest) Reset()                    { *m = DeRegisterRequest{} }
func (m *DeRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeRegisterRequest) ProtoMessage()               {}
func (*DeRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeRegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeRegisterRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *DeRegisterRequest) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

type HeartBeatRequest struct {
	AppId        string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string     `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Release      AppRelease `protobuf:"varint,3,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *HeartBeatRequest) Reset()                    { *m = HeartBeatRequest{} }
func (m *HeartBeatRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartBeatRequest) ProtoMessage()               {}
func (*HeartBeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HeartBeatRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *HeartBeatRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *HeartBeatRequest) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

type LocationRequest struct {
	AppId   string     `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Release AppRelease `protobuf:"varint,2,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *LocationRequest) Reset()                    { *m = LocationRequest{} }
func (m *LocationRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationRequest) ProtoMessage()               {}
func (*LocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LocationRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *LocationRequest) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

type StatusRequest struct {
	AppId        string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string        `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Message      string        `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Status       ServiceStatus `protobuf:"varint,4,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
	Target       StatusTarget  `protobuf:"varint,5,opt,name=target,enum=discovery.StatusTarget" json:"target,omitempty"`
	Release      AppRelease    `protobuf:"varint,6,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StatusRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *StatusRequest) GetInstanceUuid() string {
	if m != nil {
		return m.InstanceUuid
	}
	return ""
}

func (m *StatusRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StatusRequest) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_OFFLINE
}

func (m *StatusRequest) GetTarget() StatusTarget {
	if m != nil {
		return m.Target
	}
	return StatusTarget_INSTANCE
}

func (m *StatusRequest) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

type DiscoveryResponse struct {
	Recorded bool `protobuf:"varint,1,opt,name=recorded" json:"recorded,omitempty"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DiscoveryResponse) GetRecorded() bool {
	if m != nil {
		return m.Recorded
	}
	return false
}

type ServiceLocation struct {
	AppId   string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Address []string      `protobuf:"bytes,2,rep,name=address" json:"address,omitempty"`
	Status  ServiceStatus `protobuf:"varint,3,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
	Release AppRelease    `protobuf:"varint,4,opt,name=release,enum=discovery.AppRelease" json:"release,omitempty"`
}

func (m *ServiceLocation) Reset()                    { *m = ServiceLocation{} }
func (m *ServiceLocation) String() string            { return proto.CompactTextString(m) }
func (*ServiceLocation) ProtoMessage()               {}
func (*ServiceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServiceLocation) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ServiceLocation) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ServiceLocation) GetStatus() ServiceStatus {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_OFFLINE
}

func (m *ServiceLocation) GetRelease() AppRelease {
	if m != nil {
		return m.Release
	}
	return AppRelease_STABLE
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "discovery.RegisterRequest")
	proto.RegisterType((*DeRegisterRequest)(nil), "discovery.DeRegisterRequest")
	proto.RegisterType((*HeartBeatRequest)(nil), "discovery.HeartBeatRequest")
	proto.RegisterType((*LocationRequest)(nil), "discovery.LocationRequest")
	proto.RegisterType((*StatusRequest)(nil), "discovery.StatusRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "discovery.DiscoveryResponse")
	proto.RegisterType((*ServiceLocation)(nil), "discovery.ServiceLocation")
	proto.RegisterEnum("discovery.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterEnum("discovery.AppRelease", AppRelease_name, AppRelease_value)
	proto.RegisterEnum("discovery.StatusTarget", StatusTarget_name, StatusTarget_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Discovery service

type DiscoveryClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/DeRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/HeartBeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error) {
	out := new(ServiceLocation)
	err := grpc.Invoke(ctx, "/discovery.Discovery/GetLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Register(context.Context, *RegisterRequest) (*DiscoveryResponse, error)
	DeRegister(context.Context, *DeRegisterRequest) (*DiscoveryResponse, error)
	HeartBeat(context.Context, *HeartBeatRequest) (*DiscoveryResponse, error)
	Status(context.Context, *StatusRequest) (*DiscoveryResponse, error)
	GetLocation(context.Context, *LocationRequest) (*ServiceLocation, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/DeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).DeRegister(ctx, req.(*DeRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).GetLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _Discovery_DeRegister_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Discovery_HeartBeat_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Discovery_Status_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Discovery_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xc6, 0x49, 0xa6, 0x49, 0xe3, 0xae, 0xa8, 0xb0, 0x02, 0x87, 0xd0, 0x0a, 0xa9,
	0xaa, 0x44, 0x83, 0xda, 0x2b, 0x07, 0x9c, 0xda, 0x69, 0x22, 0xa5, 0x4e, 0xb4, 0x71, 0x91, 0xe0,
	0x62, 0xb9, 0xf1, 0x28, 0xb5, 0xa0, 0x59, 0xe3, 0x5d, 0x57, 0x70, 0x82, 0x9f, 0xe1, 0x4f, 0xf8,
	0x1c, 0x7e, 0x80, 0x1b, 0x8a, 0x63, 0xbb, 0xdb, 0x46, 0xa8, 0x41, 0xaa, 0xc4, 0x2d, 0x33, 0xf3,
	0xfc, 0xf6, 0xcd, 0xcc, 0xd3, 0x04, 0x9a, 0x41, 0xc8, 0xa7, 0xec, 0x06, 0xe3, 0xaf, 0x47, 0x51,
	0xcc, 0x04, 0x23, 0xb5, 0x22, 0xb1, 0xf7, 0x53, 0x81, 0x26, 0xc5, 0x59, 0xc8, 0x05, 0xc6, 0x14,
	0x3f, 0x27, 0xc8, 0x05, 0xd9, 0x05, 0xcd, 0x8f, 0x22, 0x2f, 0x0c, 0x0c, 0xa5, 0xad, 0x1c, 0xd4,
	0x68, 0xd9, 0x8f, 0xa2, 0x41, 0x40, 0xf6, 0xa1, 0x11, 0xce, 0xb9, 0xf0, 0xe7, 0x53, 0xf4, 0x92,
	0x24, 0x0c, 0x0c, 0x35, 0xad, 0xd6, 0xf3, 0xe4, 0x45, 0x12, 0x06, 0xe4, 0x05, 0xd4, 0x39, 0xc6,
	0x37, 0xe1, 0x14, 0xbd, 0x2b, 0xc6, 0x85, 0x51, 0x4a, 0x31, 0x5b, 0x59, 0xae, 0xcf, 0xb8, 0x90,
	0x21, 0x11, 0x8b, 0x85, 0xb1, 0xd9, 0x56, 0x0e, 0xca, 0x05, 0x64, 0xcc, 0x62, 0x41, 0x3a, 0x50,
	0x89, 0xf1, 0x13, 0xfa, 0x1c, 0x8d, 0x72, 0x5b, 0x39, 0xd8, 0x3e, 0xde, 0x3d, 0xba, 0xed, 0xc1,
	0x8c, 0x22, 0xba, 0x2c, 0xd2, 0x1c, 0xb5, 0xf7, 0x5d, 0x81, 0x1d, 0x0b, 0x1f, 0xb3, 0x11, 0x49,
	0x42, 0x69, 0x2d, 0x09, 0xdf, 0x40, 0xef, 0xa3, 0x1f, 0x8b, 0x2e, 0xfa, 0xe2, 0xbf, 0x08, 0x78,
	0x0f, 0xcd, 0x21, 0x9b, 0xfa, 0x22, 0x64, 0xf3, 0x07, 0xde, 0x97, 0xa8, 0xd5, 0xb5, 0xa8, 0x7f,
	0x2b, 0xd0, 0x98, 0x08, 0x5f, 0x24, 0xfc, 0x31, 0x3a, 0x33, 0xa0, 0x72, 0x8d, 0x9c, 0xfb, 0x33,
	0xcc, 0xec, 0x91, 0x87, 0xe4, 0x35, 0x68, 0x3c, 0x7d, 0x26, 0x35, 0xc5, 0xf6, 0xb1, 0x21, 0xe9,
	0x9a, 0x2c, 0xfd, 0x91, 0xc9, 0xc8, 0x70, 0xa4, 0x03, 0x9a, 0xf0, 0xe3, 0x19, 0x8a, 0xcc, 0x28,
	0x4f, 0xe5, 0x2f, 0x52, 0x88, 0x9b, 0x96, 0x69, 0x06, 0x93, 0x7b, 0xd7, 0xd6, 0xea, 0xbd, 0x03,
	0x3b, 0x56, 0x0e, 0xa0, 0xc8, 0x23, 0x36, 0xe7, 0x48, 0x5a, 0x50, 0x8d, 0x71, 0xca, 0xe2, 0x00,
	0x97, 0x03, 0xa8, 0xd2, 0x22, 0xde, 0xfb, 0xa1, 0x40, 0x33, 0x13, 0x9b, 0xef, 0xe3, 0x6f, 0xe3,
	0x32, 0xa0, 0xe2, 0x07, 0x41, 0x8c, 0x9c, 0x1b, 0x6a, 0xbb, 0xb4, 0x98, 0x44, 0x16, 0x4a, 0x93,
	0x28, 0xad, 0x3d, 0x89, 0xa2, 0xb1, 0xcd, 0x75, 0x1a, 0x3b, 0xbc, 0x86, 0xc6, 0x1d, 0x26, 0xb2,
	0x05, 0x95, 0x51, 0xaf, 0x37, 0x1c, 0x38, 0xb6, 0xbe, 0x41, 0x00, 0xb4, 0x91, 0x93, 0xfe, 0x56,
	0x88, 0x01, 0x4f, 0x2c, 0xfb, 0x8c, 0x9a, 0x96, 0x6d, 0x79, 0x63, 0x9b, 0xf6, 0x46, 0xf4, 0xdc,
	0x74, 0x4e, 0x6d, 0x5d, 0x25, 0x04, 0xb6, 0xc7, 0x26, 0x75, 0x07, 0xe6, 0xd0, 0x1b, 0x5d, 0xb8,
	0xe6, 0x99, 0xad, 0x97, 0xc8, 0x2e, 0xec, 0x5c, 0x38, 0x96, 0x4d, 0xbd, 0x73, 0x73, 0xe0, 0xb8,
	0xb6, 0x93, 0x42, 0xcb, 0x87, 0x6f, 0x00, 0x6e, 0x55, 0x2c, 0xe8, 0x27, 0xae, 0xd9, 0x1d, 0x2e,
	0x9e, 0xaa, 0xc2, 0x66, 0xd7, 0x76, 0x4d, 0x5d, 0x21, 0x35, 0x28, 0x9b, 0xc3, 0x71, 0xdf, 0xd4,
	0x55, 0xd2, 0x80, 0xda, 0x98, 0xda, 0xde, 0x32, 0x2c, 0x1d, 0x9e, 0x40, 0x5d, 0x5e, 0x27, 0xa9,
	0x43, 0x75, 0xe0, 0x4c, 0xdc, 0x94, 0x7b, 0x63, 0xa1, 0x7c, 0x62, 0xd3, 0x77, 0x83, 0xd3, 0x85,
	0xda, 0x05, 0xdd, 0xc8, 0xed, 0xeb, 0xea, 0xf1, 0x2f, 0x15, 0x6a, 0xc5, 0xee, 0x88, 0x05, 0xd5,
	0xfc, 0x40, 0x90, 0x96, 0x34, 0x9b, 0x7b, 0x57, 0xa3, 0xf5, 0x5c, 0xaa, 0xad, 0x6e, 0xbe, 0x0f,
	0x70, 0x7b, 0x68, 0xc8, 0x1d, 0x2c, 0xfe, 0x1b, 0x53, 0x0f, 0x6a, 0xc5, 0xc1, 0x20, 0xcf, 0x24,
	0xe8, 0xfd, 0x33, 0xf2, 0x00, 0xcf, 0x5b, 0xd0, 0xb2, 0x05, 0x1a, 0x2b, 0xe6, 0x5f, 0x8f, 0xc1,
	0x86, 0xad, 0x33, 0x14, 0x85, 0x59, 0xe5, 0xe1, 0xdc, 0xbb, 0x28, 0xad, 0xd6, 0xaa, 0x0f, 0x73,
	0x48, 0xf7, 0xe5, 0x87, 0xfd, 0x59, 0x28, 0xae, 0x92, 0xcb, 0xa3, 0x29, 0xbb, 0xee, 0x7c, 0x4c,
	0x2e, 0xf1, 0x4b, 0x27, 0xfd, 0xc3, 0x79, 0x35, 0x63, 0x9d, 0xe2, 0xb3, 0x4b, 0x2d, 0xcd, 0x9d,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x58, 0x6e, 0x28, 0x97, 0x06, 0x00, 0x00,
}
