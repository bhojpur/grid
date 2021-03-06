// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GridUIClient is the client API for GridUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GridUIClient interface {
	// ListStationSpecs returns a list of Station that can be started through the UI.
	ListStationSpecs(ctx context.Context, in *ListStationSpecsRequest, opts ...grpc.CallOption) (GridUI_ListStationSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type gridUIClient struct {
	cc grpc.ClientConnInterface
}

func NewGridUIClient(cc grpc.ClientConnInterface) GridUIClient {
	return &gridUIClient{cc}
}

func (c *gridUIClient) ListStationSpecs(ctx context.Context, in *ListStationSpecsRequest, opts ...grpc.CallOption) (GridUI_ListStationSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GridUI_ServiceDesc.Streams[0], "/v1.GridUI/ListStationSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &gridUIListStationSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GridUI_ListStationSpecsClient interface {
	Recv() (*ListStationSpecsResponse, error)
	grpc.ClientStream
}

type gridUIListStationSpecsClient struct {
	grpc.ClientStream
}

func (x *gridUIListStationSpecsClient) Recv() (*ListStationSpecsResponse, error) {
	m := new(ListStationSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gridUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.GridUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GridUIServer is the server API for GridUI service.
// All implementations must embed UnimplementedGridUIServer
// for forward compatibility
type GridUIServer interface {
	// ListStationSpecs returns a list of Station that can be started through the UI.
	ListStationSpecs(*ListStationSpecsRequest, GridUI_ListStationSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedGridUIServer()
}

// UnimplementedGridUIServer must be embedded to have forward compatible implementations.
type UnimplementedGridUIServer struct {
}

func (UnimplementedGridUIServer) ListStationSpecs(*ListStationSpecsRequest, GridUI_ListStationSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStationSpecs not implemented")
}
func (UnimplementedGridUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedGridUIServer) mustEmbedUnimplementedGridUIServer() {}

// UnsafeGridUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GridUIServer will
// result in compilation errors.
type UnsafeGridUIServer interface {
	mustEmbedUnimplementedGridUIServer()
}

func RegisterGridUIServer(s grpc.ServiceRegistrar, srv GridUIServer) {
	s.RegisterService(&GridUI_ServiceDesc, srv)
}

func _GridUI_ListStationSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListStationSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GridUIServer).ListStationSpecs(m, &gridUIListStationSpecsServer{stream})
}

type GridUI_ListStationSpecsServer interface {
	Send(*ListStationSpecsResponse) error
	grpc.ServerStream
}

type gridUIListStationSpecsServer struct {
	grpc.ServerStream
}

func (x *gridUIListStationSpecsServer) Send(m *ListStationSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GridUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GridUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.GridUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GridUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GridUI_ServiceDesc is the grpc.ServiceDesc for GridUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GridUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GridUI",
	HandlerType: (*GridUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _GridUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStationSpecs",
			Handler:       _GridUI_ListStationSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grid-ui.proto",
}
