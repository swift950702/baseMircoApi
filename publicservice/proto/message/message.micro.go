// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/message/message.proto

package go_micro_messageservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PushMessageToUserService service

type PushMessageToUserService interface {
	PushMessageToUser(ctx context.Context, in *PushMessageToUserParam, opts ...client.CallOption) (*PushMessageToUserResponse, error)
}

type pushMessageToUserService struct {
	c    client.Client
	name string
}

func NewPushMessageToUserService(name string, c client.Client) PushMessageToUserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.messageservice"
	}
	return &pushMessageToUserService{
		c:    c,
		name: name,
	}
}

func (c *pushMessageToUserService) PushMessageToUser(ctx context.Context, in *PushMessageToUserParam, opts ...client.CallOption) (*PushMessageToUserResponse, error) {
	req := c.c.NewRequest(c.name, "PushMessageToUserService.PushMessageToUser", in)
	out := new(PushMessageToUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PushMessageToUserService service

type PushMessageToUserServiceHandler interface {
	PushMessageToUser(context.Context, *PushMessageToUserParam, *PushMessageToUserResponse) error
}

func RegisterPushMessageToUserServiceHandler(s server.Server, hdlr PushMessageToUserServiceHandler, opts ...server.HandlerOption) error {
	type pushMessageToUserService interface {
		PushMessageToUser(ctx context.Context, in *PushMessageToUserParam, out *PushMessageToUserResponse) error
	}
	type PushMessageToUserService struct {
		pushMessageToUserService
	}
	h := &pushMessageToUserServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PushMessageToUserService{h}, opts...))
}

type pushMessageToUserServiceHandler struct {
	PushMessageToUserServiceHandler
}

func (h *pushMessageToUserServiceHandler) PushMessageToUser(ctx context.Context, in *PushMessageToUserParam, out *PushMessageToUserResponse) error {
	return h.PushMessageToUserServiceHandler.PushMessageToUser(ctx, in, out)
}

// Client API for JumpUrlWithKeyAndParamsService service

type JumpUrlWithKeyAndParamsService interface {
	JumpUrlWithKeyAndParams(ctx context.Context, in *JumpUrlWithKeyAndParamsParam, opts ...client.CallOption) (*JumpUrlWithKeyAndParamsResponse, error)
}

type jumpUrlWithKeyAndParamsService struct {
	c    client.Client
	name string
}

func NewJumpUrlWithKeyAndParamsService(name string, c client.Client) JumpUrlWithKeyAndParamsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.messageservice"
	}
	return &jumpUrlWithKeyAndParamsService{
		c:    c,
		name: name,
	}
}

func (c *jumpUrlWithKeyAndParamsService) JumpUrlWithKeyAndParams(ctx context.Context, in *JumpUrlWithKeyAndParamsParam, opts ...client.CallOption) (*JumpUrlWithKeyAndParamsResponse, error) {
	req := c.c.NewRequest(c.name, "JumpUrlWithKeyAndParamsService.JumpUrlWithKeyAndParams", in)
	out := new(JumpUrlWithKeyAndParamsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JumpUrlWithKeyAndParamsService service

type JumpUrlWithKeyAndParamsServiceHandler interface {
	JumpUrlWithKeyAndParams(context.Context, *JumpUrlWithKeyAndParamsParam, *JumpUrlWithKeyAndParamsResponse) error
}

func RegisterJumpUrlWithKeyAndParamsServiceHandler(s server.Server, hdlr JumpUrlWithKeyAndParamsServiceHandler, opts ...server.HandlerOption) error {
	type jumpUrlWithKeyAndParamsService interface {
		JumpUrlWithKeyAndParams(ctx context.Context, in *JumpUrlWithKeyAndParamsParam, out *JumpUrlWithKeyAndParamsResponse) error
	}
	type JumpUrlWithKeyAndParamsService struct {
		jumpUrlWithKeyAndParamsService
	}
	h := &jumpUrlWithKeyAndParamsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&JumpUrlWithKeyAndParamsService{h}, opts...))
}

type jumpUrlWithKeyAndParamsServiceHandler struct {
	JumpUrlWithKeyAndParamsServiceHandler
}

func (h *jumpUrlWithKeyAndParamsServiceHandler) JumpUrlWithKeyAndParams(ctx context.Context, in *JumpUrlWithKeyAndParamsParam, out *JumpUrlWithKeyAndParamsResponse) error {
	return h.JumpUrlWithKeyAndParamsServiceHandler.JumpUrlWithKeyAndParams(ctx, in, out)
}
