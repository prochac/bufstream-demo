// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: bufbuild/bufstream_demo/v1/events.proto

package bufstream_demov1

import (
	_ "buf.build/gen/go/bufbuild/confluent/protocolbuffers/go/buf/confluent/v1"
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	OldAddress string `protobuf:"bytes,2,opt,name=old_address,json=oldAddress,proto3" json:"old_address,omitempty"`
	NewAddress string `protobuf:"bytes,3,opt,name=new_address,json=newAddress,proto3" json:"new_address,omitempty"`
}

func (x *EmailUpdated) Reset() {
	*x = EmailUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bufbuild_bufstream_demo_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailUpdated) ProtoMessage() {}

func (x *EmailUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_bufbuild_bufstream_demo_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailUpdated.ProtoReflect.Descriptor instead.
func (*EmailUpdated) Descriptor() ([]byte, []int) {
	return file_bufbuild_bufstream_demo_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *EmailUpdated) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *EmailUpdated) GetOldAddress() string {
	if x != nil {
		return x.OldAddress
	}
	return ""
}

func (x *EmailUpdated) GetNewAddress() string {
	if x != nil {
		return x.NewAddress
	}
	return ""
}

var File_bufbuild_bufstream_demo_v1_events_proto protoreflect.FileDescriptor

var file_bufbuild_bufstream_demo_v1_events_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6c,
	0x75, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x6f, 0x6c, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x07, 0xd8, 0x01,
	0x01, 0x72, 0x02, 0x60, 0x01, 0x80, 0x01, 0x01, 0x52, 0x0a, 0x6f, 0x6c, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02,
	0x60, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x22,
	0xb2, 0x48, 0x1f, 0x0a, 0x0e, 0x3c, 0x43, 0x53, 0x52, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e,
	0x43, 0x45, 0x3e, 0x12, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x87, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x42, 0x58, 0xaa, 0x02,
	0x19, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x42, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5c, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5c, 0x42, 0x75, 0x66, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x6d, 0x6f, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1b, 0x42, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x3a, 0x3a, 0x42, 0x75, 0x66, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x6d, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bufbuild_bufstream_demo_v1_events_proto_rawDescOnce sync.Once
	file_bufbuild_bufstream_demo_v1_events_proto_rawDescData = file_bufbuild_bufstream_demo_v1_events_proto_rawDesc
)

func file_bufbuild_bufstream_demo_v1_events_proto_rawDescGZIP() []byte {
	file_bufbuild_bufstream_demo_v1_events_proto_rawDescOnce.Do(func() {
		file_bufbuild_bufstream_demo_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_bufbuild_bufstream_demo_v1_events_proto_rawDescData)
	})
	return file_bufbuild_bufstream_demo_v1_events_proto_rawDescData
}

var file_bufbuild_bufstream_demo_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bufbuild_bufstream_demo_v1_events_proto_goTypes = []any{
	(*EmailUpdated)(nil), // 0: bufbuild.bufstream_demo.v1.EmailUpdated
}
var file_bufbuild_bufstream_demo_v1_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bufbuild_bufstream_demo_v1_events_proto_init() }
func file_bufbuild_bufstream_demo_v1_events_proto_init() {
	if File_bufbuild_bufstream_demo_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bufbuild_bufstream_demo_v1_events_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmailUpdated); i {
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
			RawDescriptor: file_bufbuild_bufstream_demo_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bufbuild_bufstream_demo_v1_events_proto_goTypes,
		DependencyIndexes: file_bufbuild_bufstream_demo_v1_events_proto_depIdxs,
		MessageInfos:      file_bufbuild_bufstream_demo_v1_events_proto_msgTypes,
	}.Build()
	File_bufbuild_bufstream_demo_v1_events_proto = out.File
	file_bufbuild_bufstream_demo_v1_events_proto_rawDesc = nil
	file_bufbuild_bufstream_demo_v1_events_proto_goTypes = nil
	file_bufbuild_bufstream_demo_v1_events_proto_depIdxs = nil
}