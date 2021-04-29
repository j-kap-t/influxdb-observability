// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package experimental

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

// MetricConfigClient is the client API for MetricConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricConfigClient interface {
	GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error)
}

type metricConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricConfigClient(cc grpc.ClientConnInterface) MetricConfigClient {
	return &metricConfigClient{cc}
}

func (c *metricConfigClient) GetMetricConfig(ctx context.Context, in *MetricConfigRequest, opts ...grpc.CallOption) (*MetricConfigResponse, error) {
	out := new(MetricConfigResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricConfigServer is the server API for MetricConfig service.
// All implementations must embed UnimplementedMetricConfigServer
// for forward compatibility
type MetricConfigServer interface {
	GetMetricConfig(context.Context, *MetricConfigRequest) (*MetricConfigResponse, error)
	mustEmbedUnimplementedMetricConfigServer()
}

// UnimplementedMetricConfigServer must be embedded to have forward compatible implementations.
type UnimplementedMetricConfigServer struct {
}

func (UnimplementedMetricConfigServer) GetMetricConfig(context.Context, *MetricConfigRequest) (*MetricConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricConfig not implemented")
}
func (UnimplementedMetricConfigServer) mustEmbedUnimplementedMetricConfigServer() {}

// UnsafeMetricConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricConfigServer will
// result in compilation errors.
type UnsafeMetricConfigServer interface {
	mustEmbedUnimplementedMetricConfigServer()
}

func RegisterMetricConfigServer(s grpc.ServiceRegistrar, srv MetricConfigServer) {
	s.RegisterService(&MetricConfig_ServiceDesc, srv)
}

func _MetricConfig_GetMetricConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.metrics.experimental.MetricConfig/GetMetricConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricConfigServer).GetMetricConfig(ctx, req.(*MetricConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricConfig_ServiceDesc is the grpc.ServiceDesc for MetricConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.metrics.experimental.MetricConfig",
	HandlerType: (*MetricConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricConfig",
			Handler:    _MetricConfig_GetMetricConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/metrics/experimental/configservice.proto",
}
