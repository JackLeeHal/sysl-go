// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: advancedgrpc.proto

package advancedpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetStuffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerStuff     string `protobuf:"bytes,1,opt,name=innerStuff,proto3" json:"innerStuff,omitempty"`
	SensitiveStuff string `protobuf:"bytes,2,opt,name=sensitiveStuff,proto3" json:"sensitiveStuff,omitempty"`
	TimeStuff      string `protobuf:"bytes,3,opt,name=timeStuff,proto3" json:"timeStuff,omitempty"`
}

func (x *GetStuffRequest) Reset() {
	*x = GetStuffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advancedgrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStuffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffRequest) ProtoMessage() {}

func (x *GetStuffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advancedgrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffRequest.ProtoReflect.Descriptor instead.
func (*GetStuffRequest) Descriptor() ([]byte, []int) {
	return file_advancedgrpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetStuffRequest) GetInnerStuff() string {
	if x != nil {
		return x.InnerStuff
	}
	return ""
}

func (x *GetStuffRequest) GetSensitiveStuff() string {
	if x != nil {
		return x.SensitiveStuff
	}
	return ""
}

func (x *GetStuffRequest) GetTimeStuff() string {
	if x != nil {
		return x.TimeStuff
	}
	return ""
}

type GetStuffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Item `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetStuffResponse) Reset() {
	*x = GetStuffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advancedgrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStuffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffResponse) ProtoMessage() {}

func (x *GetStuffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advancedgrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffResponse.ProtoReflect.Descriptor instead.
func (*GetStuffResponse) Descriptor() ([]byte, []int) {
	return file_advancedgrpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetStuffResponse) GetData() []*Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetStuffTooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerStuff     string `protobuf:"bytes,1,opt,name=innerStuff,proto3" json:"innerStuff,omitempty"`
	SensitiveStuff string `protobuf:"bytes,2,opt,name=sensitiveStuff,proto3" json:"sensitiveStuff,omitempty"`
	TimeStuff      string `protobuf:"bytes,3,opt,name=timeStuff,proto3" json:"timeStuff,omitempty"`
}

func (x *GetStuffTooRequest) Reset() {
	*x = GetStuffTooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advancedgrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStuffTooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffTooRequest) ProtoMessage() {}

func (x *GetStuffTooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advancedgrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffTooRequest.ProtoReflect.Descriptor instead.
func (*GetStuffTooRequest) Descriptor() ([]byte, []int) {
	return file_advancedgrpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetStuffTooRequest) GetInnerStuff() string {
	if x != nil {
		return x.InnerStuff
	}
	return ""
}

func (x *GetStuffTooRequest) GetSensitiveStuff() string {
	if x != nil {
		return x.SensitiveStuff
	}
	return ""
}

func (x *GetStuffTooRequest) GetTimeStuff() string {
	if x != nil {
		return x.TimeStuff
	}
	return ""
}

type GetStuffTooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Item `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetStuffTooResponse) Reset() {
	*x = GetStuffTooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advancedgrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStuffTooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStuffTooResponse) ProtoMessage() {}

func (x *GetStuffTooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advancedgrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStuffTooResponse.ProtoReflect.Descriptor instead.
func (*GetStuffTooResponse) Descriptor() ([]byte, []int) {
	return file_advancedgrpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetStuffTooResponse) GetData() []*Item {
	if x != nil {
		return x.Data
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	A1   string `protobuf:"bytes,2,opt,name=A1,proto3" json:"A1,omitempty"`
	A2   string `protobuf:"bytes,3,opt,name=A2,proto3" json:"A2,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advancedgrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_advancedgrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_advancedgrpc_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetA1() string {
	if x != nil {
		return x.A1
	}
	return ""
}

func (x *Item) GetA2() string {
	if x != nil {
		return x.A2
	}
	return ""
}

var File_advancedgrpc_proto protoreflect.FileDescriptor

var file_advancedgrpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70, 0x62,
	0x22, 0x77, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x75, 0x66,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x75, 0x66, 0x66, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x54,
	0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x53, 0x74, 0x75, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x75, 0x66,
	0x66, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x75, 0x66, 0x66, 0x22,
	0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x54, 0x6f, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70,
	0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x41, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x41, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x41, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x41, 0x32, 0x32, 0x9f, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x64, 0x47, 0x72, 0x70, 0x63, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x66, 0x66, 0x54, 0x6f, 0x6f, 0x12,
	0x1b, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_advancedgrpc_proto_rawDescOnce sync.Once
	file_advancedgrpc_proto_rawDescData = file_advancedgrpc_proto_rawDesc
)

func file_advancedgrpc_proto_rawDescGZIP() []byte {
	file_advancedgrpc_proto_rawDescOnce.Do(func() {
		file_advancedgrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_advancedgrpc_proto_rawDescData)
	})
	return file_advancedgrpc_proto_rawDescData
}

var file_advancedgrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_advancedgrpc_proto_goTypes = []interface{}{
	(*GetStuffRequest)(nil),     // 0: advancedpb.GetStuffRequest
	(*GetStuffResponse)(nil),    // 1: advancedpb.GetStuffResponse
	(*GetStuffTooRequest)(nil),  // 2: advancedpb.GetStuffTooRequest
	(*GetStuffTooResponse)(nil), // 3: advancedpb.GetStuffTooResponse
	(*Item)(nil),                // 4: advancedpb.Item
}
var file_advancedgrpc_proto_depIdxs = []int32{
	4, // 0: advancedpb.GetStuffResponse.Data:type_name -> advancedpb.Item
	4, // 1: advancedpb.GetStuffTooResponse.Data:type_name -> advancedpb.Item
	0, // 2: advancedpb.AdvancedGrpc.GetStuff:input_type -> advancedpb.GetStuffRequest
	0, // 3: advancedpb.AdvancedGrpc.GetStuffToo:input_type -> advancedpb.GetStuffRequest
	1, // 4: advancedpb.AdvancedGrpc.GetStuff:output_type -> advancedpb.GetStuffResponse
	1, // 5: advancedpb.AdvancedGrpc.GetStuffToo:output_type -> advancedpb.GetStuffResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_advancedgrpc_proto_init() }
func file_advancedgrpc_proto_init() {
	if File_advancedgrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_advancedgrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStuffRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advancedgrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStuffResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advancedgrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStuffTooRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advancedgrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStuffTooResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_advancedgrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_advancedgrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_advancedgrpc_proto_goTypes,
		DependencyIndexes: file_advancedgrpc_proto_depIdxs,
		MessageInfos:      file_advancedgrpc_proto_msgTypes,
	}.Build()
	File_advancedgrpc_proto = out.File
	file_advancedgrpc_proto_rawDesc = nil
	file_advancedgrpc_proto_goTypes = nil
	file_advancedgrpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdvancedGrpcClient is the client API for AdvancedGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdvancedGrpcClient interface {
	GetStuff(ctx context.Context, in *GetStuffRequest, opts ...grpc.CallOption) (*GetStuffResponse, error)
	GetStuffToo(ctx context.Context, in *GetStuffRequest, opts ...grpc.CallOption) (*GetStuffResponse, error)
}

type advancedGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvancedGrpcClient(cc grpc.ClientConnInterface) AdvancedGrpcClient {
	return &advancedGrpcClient{cc}
}

func (c *advancedGrpcClient) GetStuff(ctx context.Context, in *GetStuffRequest, opts ...grpc.CallOption) (*GetStuffResponse, error) {
	out := new(GetStuffResponse)
	err := c.cc.Invoke(ctx, "/advancedpb.AdvancedGrpc/GetStuff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advancedGrpcClient) GetStuffToo(ctx context.Context, in *GetStuffRequest, opts ...grpc.CallOption) (*GetStuffResponse, error) {
	out := new(GetStuffResponse)
	err := c.cc.Invoke(ctx, "/advancedpb.AdvancedGrpc/GetStuffToo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedGrpcServer is the server API for AdvancedGrpc service.
type AdvancedGrpcServer interface {
	GetStuff(context.Context, *GetStuffRequest) (*GetStuffResponse, error)
	GetStuffToo(context.Context, *GetStuffRequest) (*GetStuffResponse, error)
}

// UnimplementedAdvancedGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedAdvancedGrpcServer struct {
}

func (*UnimplementedAdvancedGrpcServer) GetStuff(context.Context, *GetStuffRequest) (*GetStuffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStuff not implemented")
}
func (*UnimplementedAdvancedGrpcServer) GetStuffToo(context.Context, *GetStuffRequest) (*GetStuffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStuffToo not implemented")
}

func RegisterAdvancedGrpcServer(s *grpc.Server, srv AdvancedGrpcServer) {
	s.RegisterService(&_AdvancedGrpc_serviceDesc, srv)
}

func _AdvancedGrpc_GetStuff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStuffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvancedGrpcServer).GetStuff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advancedpb.AdvancedGrpc/GetStuff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvancedGrpcServer).GetStuff(ctx, req.(*GetStuffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvancedGrpc_GetStuffToo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStuffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvancedGrpcServer).GetStuffToo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advancedpb.AdvancedGrpc/GetStuffToo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvancedGrpcServer).GetStuffToo(ctx, req.(*GetStuffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdvancedGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advancedpb.AdvancedGrpc",
	HandlerType: (*AdvancedGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStuff",
			Handler:    _AdvancedGrpc_GetStuff_Handler,
		},
		{
			MethodName: "GetStuffToo",
			Handler:    _AdvancedGrpc_GetStuffToo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advancedgrpc.proto",
}
