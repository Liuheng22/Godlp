// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: dlp.proto

package rpcpb

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

// DlpServiceClient is the client API for DlpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DlpServiceClient interface {
	Dlp(ctx context.Context, in *Intext, opts ...grpc.CallOption) (*Outtext, error)
}

type dlpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDlpServiceClient(cc grpc.ClientConnInterface) DlpServiceClient {
	return &dlpServiceClient{cc}
}

func (c *dlpServiceClient) Dlp(ctx context.Context, in *Intext, opts ...grpc.CallOption) (*Outtext, error) {
	out := new(Outtext)
	err := c.cc.Invoke(ctx, "/rpcpb.DlpService/Dlp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DlpServiceServer is the server API for DlpService service.
// All implementations must embed UnimplementedDlpServiceServer
// for forward compatibility
type DlpServiceServer interface {
	Dlp(context.Context, *Intext) (*Outtext, error)
	mustEmbedUnimplementedDlpServiceServer()
}

// UnimplementedDlpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDlpServiceServer struct {
}

func (UnimplementedDlpServiceServer) Dlp(context.Context, *Intext) (*Outtext, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dlp not implemented")
}
func (UnimplementedDlpServiceServer) mustEmbedUnimplementedDlpServiceServer() {}

// UnsafeDlpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DlpServiceServer will
// result in compilation errors.
type UnsafeDlpServiceServer interface {
	mustEmbedUnimplementedDlpServiceServer()
}

func RegisterDlpServiceServer(s grpc.ServiceRegistrar, srv DlpServiceServer) {
	s.RegisterService(&DlpService_ServiceDesc, srv)
}

func _DlpService_Dlp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Intext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlpServiceServer).Dlp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.DlpService/Dlp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlpServiceServer).Dlp(ctx, req.(*Intext))
	}
	return interceptor(ctx, in, info, handler)
}

// DlpService_ServiceDesc is the grpc.ServiceDesc for DlpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DlpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.DlpService",
	HandlerType: (*DlpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dlp",
			Handler:    _DlpService_Dlp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dlp.proto",
}