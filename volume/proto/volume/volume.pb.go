// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/volume/volume.proto

package volume

import (
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

type VolumeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VolumeName                 string  `protobuf:"bytes,2,opt,name=volume_name,json=volumeName,proto3" json:"volume_name,omitempty"`
	VolumeNamespace            string  `protobuf:"bytes,3,opt,name=volume_namespace,json=volumeNamespace,proto3" json:"volume_namespace,omitempty"`
	VolumeAccessMode           string  `protobuf:"bytes,4,opt,name=volume_access_mode,json=volumeAccessMode,proto3" json:"volume_access_mode,omitempty"`
	VolumeStorageClassName     string  `protobuf:"bytes,5,opt,name=volume_storage_class_name,json=volumeStorageClassName,proto3" json:"volume_storage_class_name,omitempty"`
	VolumeRequest              float32 `protobuf:"fixed32,6,opt,name=volume_request,json=volumeRequest,proto3" json:"volume_request,omitempty"`
	VolumePersistentVolumeMode string  `protobuf:"bytes,7,opt,name=volume_persistent_volume_mode,json=volumePersistentVolumeMode,proto3" json:"volume_persistent_volume_mode,omitempty"`
}

func (x *VolumeInfo) Reset() {
	*x = VolumeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeInfo) ProtoMessage() {}

func (x *VolumeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeInfo.ProtoReflect.Descriptor instead.
func (*VolumeInfo) Descriptor() ([]byte, []int) {
	return file_proto_volume_volume_proto_rawDescGZIP(), []int{0}
}

func (x *VolumeInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VolumeInfo) GetVolumeName() string {
	if x != nil {
		return x.VolumeName
	}
	return ""
}

func (x *VolumeInfo) GetVolumeNamespace() string {
	if x != nil {
		return x.VolumeNamespace
	}
	return ""
}

func (x *VolumeInfo) GetVolumeAccessMode() string {
	if x != nil {
		return x.VolumeAccessMode
	}
	return ""
}

func (x *VolumeInfo) GetVolumeStorageClassName() string {
	if x != nil {
		return x.VolumeStorageClassName
	}
	return ""
}

func (x *VolumeInfo) GetVolumeRequest() float32 {
	if x != nil {
		return x.VolumeRequest
	}
	return 0
}

func (x *VolumeInfo) GetVolumePersistentVolumeMode() string {
	if x != nil {
		return x.VolumePersistentVolumeMode
	}
	return ""
}

type VolumeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VolumeId) Reset() {
	*x = VolumeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeId) ProtoMessage() {}

func (x *VolumeId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeId.ProtoReflect.Descriptor instead.
func (*VolumeId) Descriptor() ([]byte, []int) {
	return file_proto_volume_volume_proto_rawDescGZIP(), []int{1}
}

func (x *VolumeId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAll) Reset() {
	*x = FindAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAll) ProtoMessage() {}

func (x *FindAll) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAll.ProtoReflect.Descriptor instead.
func (*FindAll) Descriptor() ([]byte, []int) {
	return file_proto_volume_volume_proto_rawDescGZIP(), []int{2}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_volume_volume_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AllVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VolumeInfo []*VolumeInfo `protobuf:"bytes,1,rep,name=volume_info,json=volumeInfo,proto3" json:"volume_info,omitempty"`
}

func (x *AllVolume) Reset() {
	*x = AllVolume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_volume_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllVolume) ProtoMessage() {}

func (x *AllVolume) ProtoReflect() protoreflect.Message {
	mi := &file_proto_volume_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllVolume.ProtoReflect.Descriptor instead.
func (*AllVolume) Descriptor() ([]byte, []int) {
	return file_proto_volume_volume_proto_rawDescGZIP(), []int{4}
}

func (x *AllVolume) GetVolumeInfo() []*VolumeInfo {
	if x != nil {
		return x.VolumeInfo
	}
	return nil
}

var File_proto_volume_volume_proto protoreflect.FileDescriptor

var file_proto_volume_volume_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x22, 0xbb, 0x02, 0x0a, 0x0a, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x19,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x16, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41,
	0x0a, 0x1d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x22, 0x1a, 0x0a, 0x08, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x09, 0x0a,
	0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x40, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x9c, 0x02, 0x0a, 0x06, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x12, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x12,
	0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x10, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x0f, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x1a, 0x11, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x3b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_volume_volume_proto_rawDescOnce sync.Once
	file_proto_volume_volume_proto_rawDescData = file_proto_volume_volume_proto_rawDesc
)

func file_proto_volume_volume_proto_rawDescGZIP() []byte {
	file_proto_volume_volume_proto_rawDescOnce.Do(func() {
		file_proto_volume_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_volume_volume_proto_rawDescData)
	})
	return file_proto_volume_volume_proto_rawDescData
}

var file_proto_volume_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_volume_volume_proto_goTypes = []interface{}{
	(*VolumeInfo)(nil), // 0: volume.VolumeInfo
	(*VolumeId)(nil),   // 1: volume.VolumeId
	(*FindAll)(nil),    // 2: volume.FindAll
	(*Response)(nil),   // 3: volume.Response
	(*AllVolume)(nil),  // 4: volume.AllVolume
}
var file_proto_volume_volume_proto_depIdxs = []int32{
	0, // 0: volume.AllVolume.volume_info:type_name -> volume.VolumeInfo
	0, // 1: volume.Volume.AddVolume:input_type -> volume.VolumeInfo
	1, // 2: volume.Volume.DeleteVolume:input_type -> volume.VolumeId
	0, // 3: volume.Volume.UpdateVolume:input_type -> volume.VolumeInfo
	1, // 4: volume.Volume.FindVolumeByID:input_type -> volume.VolumeId
	2, // 5: volume.Volume.FindAllVolume:input_type -> volume.FindAll
	3, // 6: volume.Volume.AddVolume:output_type -> volume.Response
	3, // 7: volume.Volume.DeleteVolume:output_type -> volume.Response
	3, // 8: volume.Volume.UpdateVolume:output_type -> volume.Response
	0, // 9: volume.Volume.FindVolumeByID:output_type -> volume.VolumeInfo
	4, // 10: volume.Volume.FindAllVolume:output_type -> volume.AllVolume
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_volume_volume_proto_init() }
func file_proto_volume_volume_proto_init() {
	if File_proto_volume_volume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_volume_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeInfo); i {
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
		file_proto_volume_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeId); i {
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
		file_proto_volume_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAll); i {
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
		file_proto_volume_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_volume_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllVolume); i {
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
			RawDescriptor: file_proto_volume_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_volume_volume_proto_goTypes,
		DependencyIndexes: file_proto_volume_volume_proto_depIdxs,
		MessageInfos:      file_proto_volume_volume_proto_msgTypes,
	}.Build()
	File_proto_volume_volume_proto = out.File
	file_proto_volume_volume_proto_rawDesc = nil
	file_proto_volume_volume_proto_goTypes = nil
	file_proto_volume_volume_proto_depIdxs = nil
}