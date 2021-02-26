// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TermproxyServiceClient is the client API for TermproxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TermproxyServiceClient interface {
	OpenTerminal(ctx context.Context, opts ...grpc.CallOption) (TermproxyService_OpenTerminalClient, error)
}

type termproxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTermproxyServiceClient(cc grpc.ClientConnInterface) TermproxyServiceClient {
	return &termproxyServiceClient{cc}
}

func (c *termproxyServiceClient) OpenTerminal(ctx context.Context, opts ...grpc.CallOption) (TermproxyService_OpenTerminalClient, error) {
	stream, err := c.cc.NewStream(ctx, &TermproxyService_ServiceDesc.Streams[0], "/TermproxyService/OpenTerminal", opts...)
	if err != nil {
		return nil, err
	}
	x := &termproxyServiceOpenTerminalClient{stream}
	return x, nil
}

type TermproxyService_OpenTerminalClient interface {
	Send(*TerminalBytes) error
	Recv() (*TerminalBytes, error)
	grpc.ClientStream
}

type termproxyServiceOpenTerminalClient struct {
	grpc.ClientStream
}

func (x *termproxyServiceOpenTerminalClient) Send(m *TerminalBytes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *termproxyServiceOpenTerminalClient) Recv() (*TerminalBytes, error) {
	m := new(TerminalBytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TermproxyServiceServer is the server API for TermproxyService service.
// All implementations must embed UnimplementedTermproxyServiceServer
// for forward compatibility
type TermproxyServiceServer interface {
	OpenTerminal(TermproxyService_OpenTerminalServer) error
	mustEmbedUnimplementedTermproxyServiceServer()
}

// UnimplementedTermproxyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTermproxyServiceServer struct {
}

func (UnimplementedTermproxyServiceServer) OpenTerminal(TermproxyService_OpenTerminalServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenTerminal not implemented")
}
func (UnimplementedTermproxyServiceServer) mustEmbedUnimplementedTermproxyServiceServer() {}

// UnsafeTermproxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TermproxyServiceServer will
// result in compilation errors.
type UnsafeTermproxyServiceServer interface {
	mustEmbedUnimplementedTermproxyServiceServer()
}

func RegisterTermproxyServiceServer(s grpc.ServiceRegistrar, srv TermproxyServiceServer) {
	s.RegisterService(&TermproxyService_ServiceDesc, srv)
}

func _TermproxyService_OpenTerminal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TermproxyServiceServer).OpenTerminal(&termproxyServiceOpenTerminalServer{stream})
}

type TermproxyService_OpenTerminalServer interface {
	Send(*TerminalBytes) error
	Recv() (*TerminalBytes, error)
	grpc.ServerStream
}

type termproxyServiceOpenTerminalServer struct {
	grpc.ServerStream
}

func (x *termproxyServiceOpenTerminalServer) Send(m *TerminalBytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *termproxyServiceOpenTerminalServer) Recv() (*TerminalBytes, error) {
	m := new(TerminalBytes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TermproxyService_ServiceDesc is the grpc.ServiceDesc for TermproxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TermproxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TermproxyService",
	HandlerType: (*TermproxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenTerminal",
			Handler:       _TermproxyService_OpenTerminal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "termproxy.proto",
}