// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.2
// source: phrase.proto

package phrasev1

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

// PhraseServiceClient is the client API for PhraseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhraseServiceClient interface {
	GetPhrase(ctx context.Context, in *PhraseRequest, opts ...grpc.CallOption) (*Phrase, error)
	GetPhrasesByStory(ctx context.Context, in *PhrasesByStoryIdRequest, opts ...grpc.CallOption) (*PhrasesByStoryIdResponse, error)
	AddToStory(ctx context.Context, in *AddPhraseToStoryRequest, opts ...grpc.CallOption) (*Phrase, error)
	GetPhrasesStream(ctx context.Context, in *PhrasesByStoryIdRequest, opts ...grpc.CallOption) (PhraseService_GetPhrasesStreamClient, error)
}

type phraseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPhraseServiceClient(cc grpc.ClientConnInterface) PhraseServiceClient {
	return &phraseServiceClient{cc}
}

func (c *phraseServiceClient) GetPhrase(ctx context.Context, in *PhraseRequest, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/phrase.PhraseService/getPhrase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phraseServiceClient) GetPhrasesByStory(ctx context.Context, in *PhrasesByStoryIdRequest, opts ...grpc.CallOption) (*PhrasesByStoryIdResponse, error) {
	out := new(PhrasesByStoryIdResponse)
	err := c.cc.Invoke(ctx, "/phrase.PhraseService/getPhrasesByStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phraseServiceClient) AddToStory(ctx context.Context, in *AddPhraseToStoryRequest, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/phrase.PhraseService/addToStory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phraseServiceClient) GetPhrasesStream(ctx context.Context, in *PhrasesByStoryIdRequest, opts ...grpc.CallOption) (PhraseService_GetPhrasesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PhraseService_ServiceDesc.Streams[0], "/phrase.PhraseService/getPhrasesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &phraseServiceGetPhrasesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PhraseService_GetPhrasesStreamClient interface {
	Recv() (*Phrase, error)
	grpc.ClientStream
}

type phraseServiceGetPhrasesStreamClient struct {
	grpc.ClientStream
}

func (x *phraseServiceGetPhrasesStreamClient) Recv() (*Phrase, error) {
	m := new(Phrase)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PhraseServiceServer is the server API for PhraseService service.
// All implementations must embed UnimplementedPhraseServiceServer
// for forward compatibility
type PhraseServiceServer interface {
	GetPhrase(context.Context, *PhraseRequest) (*Phrase, error)
	GetPhrasesByStory(context.Context, *PhrasesByStoryIdRequest) (*PhrasesByStoryIdResponse, error)
	AddToStory(context.Context, *AddPhraseToStoryRequest) (*Phrase, error)
	GetPhrasesStream(*PhrasesByStoryIdRequest, PhraseService_GetPhrasesStreamServer) error
	mustEmbedUnimplementedPhraseServiceServer()
}

// UnimplementedPhraseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPhraseServiceServer struct {
}

func (UnimplementedPhraseServiceServer) GetPhrase(context.Context, *PhraseRequest) (*Phrase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhrase not implemented")
}
func (UnimplementedPhraseServiceServer) GetPhrasesByStory(context.Context, *PhrasesByStoryIdRequest) (*PhrasesByStoryIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhrasesByStory not implemented")
}
func (UnimplementedPhraseServiceServer) AddToStory(context.Context, *AddPhraseToStoryRequest) (*Phrase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToStory not implemented")
}
func (UnimplementedPhraseServiceServer) GetPhrasesStream(*PhrasesByStoryIdRequest, PhraseService_GetPhrasesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPhrasesStream not implemented")
}
func (UnimplementedPhraseServiceServer) mustEmbedUnimplementedPhraseServiceServer() {}

// UnsafePhraseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhraseServiceServer will
// result in compilation errors.
type UnsafePhraseServiceServer interface {
	mustEmbedUnimplementedPhraseServiceServer()
}

func RegisterPhraseServiceServer(s grpc.ServiceRegistrar, srv PhraseServiceServer) {
	s.RegisterService(&PhraseService_ServiceDesc, srv)
}

func _PhraseService_GetPhrase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhraseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhraseServiceServer).GetPhrase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phrase.PhraseService/getPhrase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhraseServiceServer).GetPhrase(ctx, req.(*PhraseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhraseService_GetPhrasesByStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhrasesByStoryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhraseServiceServer).GetPhrasesByStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phrase.PhraseService/getPhrasesByStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhraseServiceServer).GetPhrasesByStory(ctx, req.(*PhrasesByStoryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhraseService_AddToStory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhraseToStoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhraseServiceServer).AddToStory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phrase.PhraseService/addToStory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhraseServiceServer).AddToStory(ctx, req.(*AddPhraseToStoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhraseService_GetPhrasesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PhrasesByStoryIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhraseServiceServer).GetPhrasesStream(m, &phraseServiceGetPhrasesStreamServer{stream})
}

type PhraseService_GetPhrasesStreamServer interface {
	Send(*Phrase) error
	grpc.ServerStream
}

type phraseServiceGetPhrasesStreamServer struct {
	grpc.ServerStream
}

func (x *phraseServiceGetPhrasesStreamServer) Send(m *Phrase) error {
	return x.ServerStream.SendMsg(m)
}

// PhraseService_ServiceDesc is the grpc.ServiceDesc for PhraseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhraseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phrase.PhraseService",
	HandlerType: (*PhraseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPhrase",
			Handler:    _PhraseService_GetPhrase_Handler,
		},
		{
			MethodName: "getPhrasesByStory",
			Handler:    _PhraseService_GetPhrasesByStory_Handler,
		},
		{
			MethodName: "addToStory",
			Handler:    _PhraseService_AddToStory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getPhrasesStream",
			Handler:       _PhraseService_GetPhrasesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "phrase.proto",
}
