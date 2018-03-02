// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vip.proto

/*
Package go_micro_srv_vip is a generated protocol buffer package.

It is generated from these files:
	vip.proto

It has these top-level messages:
	VIPRequest
	VIPResponse
*/
package go_micro_srv_vip

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type VIPRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *VIPRequest) Reset()                    { *m = VIPRequest{} }
func (m *VIPRequest) String() string            { return proto.CompactTextString(m) }
func (*VIPRequest) ProtoMessage()               {}
func (*VIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VIPRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VIPResponse struct {
	IsVip bool `protobuf:"varint,1,opt,name=is_vip,json=isVip" json:"is_vip,omitempty"`
}

func (m *VIPResponse) Reset()                    { *m = VIPResponse{} }
func (m *VIPResponse) String() string            { return proto.CompactTextString(m) }
func (*VIPResponse) ProtoMessage()               {}
func (*VIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VIPResponse) GetIsVip() bool {
	if m != nil {
		return m.IsVip
	}
	return false
}

func init() {
	proto.RegisterType((*VIPRequest)(nil), "go.micro.srv.vip.VIPRequest")
	proto.RegisterType((*VIPResponse)(nil), "go.micro.srv.vip.VIPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VIP service

type VIPClient interface {
	CheckVIP(ctx context.Context, in *VIPRequest, opts ...client.CallOption) (*VIPResponse, error)
}

type vIPClient struct {
	c           client.Client
	serviceName string
}

func NewVIPClient(serviceName string, c client.Client) VIPClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vip"
	}
	return &vIPClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vIPClient) CheckVIP(ctx context.Context, in *VIPRequest, opts ...client.CallOption) (*VIPResponse, error) {
	req := c.c.NewRequest(c.serviceName, "VIP.CheckVIP", in)
	out := new(VIPResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VIP service

type VIPHandler interface {
	CheckVIP(context.Context, *VIPRequest, *VIPResponse) error
}

func RegisterVIPHandler(s server.Server, hdlr VIPHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VIP{hdlr}, opts...))
}

type VIP struct {
	VIPHandler
}

func (h *VIP) CheckVIP(ctx context.Context, in *VIPRequest, out *VIPResponse) error {
	return h.VIPHandler.CheckVIP(ctx, in, out)
}

func init() { proto.RegisterFile("vip.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xcb, 0x2c, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c, 0x2e, 0xca, 0xd7,
	0x2b, 0x2e, 0x2a, 0xd3, 0x2b, 0xcb, 0x2c, 0x50, 0x52, 0xe0, 0xe2, 0x0a, 0xf3, 0x0c, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x54, 0xb8, 0xb8, 0xc1, 0x2a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0x44, 0xb9, 0xd8, 0x32, 0x8b, 0xe3, 0xcb, 0x32, 0x0b, 0xc0, 0x8a, 0x38, 0x82,
	0x58, 0x33, 0x8b, 0xc3, 0x32, 0x0b, 0x8c, 0x02, 0xb8, 0x98, 0xc3, 0x3c, 0x03, 0x84, 0x3c, 0xb9,
	0x38, 0x9c, 0x33, 0x52, 0x93, 0xb3, 0x41, 0x6c, 0x19, 0x3d, 0x74, 0xdb, 0xf4, 0x10, 0x56, 0x49,
	0xc9, 0xe2, 0x90, 0x85, 0x58, 0xa3, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xb2, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x0d, 0xb5, 0xa5, 0xbf, 0x00, 0x00, 0x00,
}