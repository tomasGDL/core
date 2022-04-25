// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/core/v1/rawdata.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRawdataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId     string            `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	StartTime    int64             `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      int64             `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Path         string            `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	PageNum      int32             `protobuf:"varint,5,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize     int32             `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	IsDescending bool              `protobuf:"varint,7,opt,name=is_descending,json=isDescending,proto3" json:"is_descending,omitempty"`
	Filters      map[string]string `protobuf:"bytes,8,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetRawdataRequest) Reset() {
	*x = GetRawdataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_rawdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawdataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawdataRequest) ProtoMessage() {}

func (x *GetRawdataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_rawdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawdataRequest.ProtoReflect.Descriptor instead.
func (*GetRawdataRequest) Descriptor() ([]byte, []int) {
	return file_api_core_v1_rawdata_proto_rawDescGZIP(), []int{0}
}

func (x *GetRawdataRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *GetRawdataRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetRawdataRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GetRawdataRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetRawdataRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetRawdataRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetRawdataRequest) GetIsDescending() bool {
	if x != nil {
		return x.IsDescending
	}
	return false
}

func (x *GetRawdataRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

type GetRawdataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int32              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageNum  int32              `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int32              `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Items    []*RawdataResponse `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetRawdataResponse) Reset() {
	*x = GetRawdataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_rawdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawdataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawdataResponse) ProtoMessage() {}

func (x *GetRawdataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_rawdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawdataResponse.ProtoReflect.Descriptor instead.
func (*GetRawdataResponse) Descriptor() ([]byte, []int) {
	return file_api_core_v1_rawdata_proto_rawDescGZIP(), []int{1}
}

func (x *GetRawdataResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetRawdataResponse) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetRawdataResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetRawdataResponse) GetItems() []*RawdataResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type RawdataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	EntityId  string `protobuf:"bytes,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Values    string `protobuf:"bytes,5,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *RawdataResponse) Reset() {
	*x = RawdataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_core_v1_rawdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawdataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawdataResponse) ProtoMessage() {}

func (x *RawdataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_core_v1_rawdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawdataResponse.ProtoReflect.Descriptor instead.
func (*RawdataResponse) Descriptor() ([]byte, []int) {
	return file_api_core_v1_rawdata_proto_rawDescGZIP(), []int{2}
}

func (x *RawdataResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *RawdataResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RawdataResponse) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *RawdataResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RawdataResponse) GetValues() string {
	if x != nil {
		return x.Values
	}
	return ""
}

var File_api_core_v1_rawdata_proto protoreflect.FileDescriptor

var file_api_core_v1_rawdata_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61,
	0x77, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x04, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0x92, 0x41, 0x0b, 0x32, 0x09, 0xe5, 0xae, 0x9e, 0xe4, 0xbd, 0x93, 0x20, 0x69, 0x64, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x11, 0x92,
	0x41, 0x0e, 0x32, 0x0c, 0xe8, 0xb5, 0xb7, 0xe5, 0xa7, 0x8b, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x11, 0x92,
	0x41, 0x0e, 0x32, 0x0c, 0xe7, 0xbb, 0x88, 0xe6, 0xad, 0xa2, 0xe6, 0x97, 0xb6, 0xe9, 0x97, 0xb4,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x92, 0x41, 0x11, 0x32, 0x0f, 0xe5, 0xb1,
	0x9e, 0xe6, 0x80, 0xa7, 0xe7, 0x9a, 0x84, 0xe4, 0xbd, 0x8d, 0xe7, 0xbd, 0xae, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0xe8, 0xae, 0xb0, 0xe5,
	0xbd, 0x95, 0xe5, 0xbc, 0x80, 0xe5, 0xa7, 0x8b, 0xe4, 0xbd, 0x8d, 0xe7, 0xbd, 0xae, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32,
	0x12, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5, 0xe9, 0x99, 0x90, 0xe5, 0x88, 0xb6, 0xe6, 0x9d, 0xa1,
	0xe6, 0x95, 0xb0, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x59, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x34, 0x92, 0x41, 0x31, 0x32, 0x2f, 0xe6, 0x98, 0xaf, 0xe5, 0x90,
	0xa6, 0xe9, 0x80, 0x86, 0xe5, 0xba, 0x8f, 0xef, 0xbc, 0x8c, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0xef, 0xbc, 0x9a, 0xe4, 0xb8, 0x8d, 0xe9, 0x80, 0x86, 0xe5, 0xba, 0x8f, 0xef, 0xbc, 0x8c, 0x74,
	0x72, 0x75, 0x65, 0x3a, 0xe9, 0x80, 0x86, 0xe5, 0xba, 0x8f, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x65,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x58, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x11, 0x92, 0x41, 0x0e, 0x32, 0x0c, 0xe8, 0xbf, 0x87,
	0xe6, 0xbb, 0xa4, 0xe6, 0x9d, 0xa1, 0xe4, 0xbb, 0xb6, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x27,
	0x92, 0x41, 0x24, 0x0a, 0x22, 0x2a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x14, 0x47, 0x65, 0x74, 0x20, 0x72, 0x61, 0x77, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x20, 0x92,
	0x41, 0x1d, 0x32, 0x1b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92, 0x41, 0x14, 0x32, 0x12, 0xe8,
	0xae, 0xb0, 0xe5, 0xbd, 0x95, 0xe5, 0xbc, 0x80, 0xe5, 0xa7, 0x8b, 0xe4, 0xbd, 0x8d, 0xe7, 0xbd,
	0xae, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x92,
	0x41, 0x14, 0x32, 0x12, 0xe6, 0xaf, 0x8f, 0xe9, 0xa1, 0xb5, 0xe9, 0x99, 0x90, 0xe5, 0x88, 0xb6,
	0xe6, 0x9d, 0xa1, 0xe6, 0x95, 0xb0, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x77, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x92,
	0x41, 0x0b, 0x32, 0x09, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32,
	0xb3, 0x01, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x12, 0xa7, 0x01, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x7b, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x36, 0x0a,
	0x07, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x47, 0x65, 0x74, 0x20, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x20, 0x72, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x77, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12,
	0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x42, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_core_v1_rawdata_proto_rawDescOnce sync.Once
	file_api_core_v1_rawdata_proto_rawDescData = file_api_core_v1_rawdata_proto_rawDesc
)

func file_api_core_v1_rawdata_proto_rawDescGZIP() []byte {
	file_api_core_v1_rawdata_proto_rawDescOnce.Do(func() {
		file_api_core_v1_rawdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_core_v1_rawdata_proto_rawDescData)
	})
	return file_api_core_v1_rawdata_proto_rawDescData
}

var file_api_core_v1_rawdata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_core_v1_rawdata_proto_goTypes = []interface{}{
	(*GetRawdataRequest)(nil),  // 0: api.core.v1.GetRawdataRequest
	(*GetRawdataResponse)(nil), // 1: api.core.v1.GetRawdataResponse
	(*RawdataResponse)(nil),    // 2: api.core.v1.RawdataResponse
	nil,                        // 3: api.core.v1.GetRawdataRequest.FiltersEntry
}
var file_api_core_v1_rawdata_proto_depIdxs = []int32{
	3, // 0: api.core.v1.GetRawdataRequest.filters:type_name -> api.core.v1.GetRawdataRequest.FiltersEntry
	2, // 1: api.core.v1.GetRawdataResponse.items:type_name -> api.core.v1.RawdataResponse
	0, // 2: api.core.v1.Rawdata.GetRawdata:input_type -> api.core.v1.GetRawdataRequest
	1, // 3: api.core.v1.Rawdata.GetRawdata:output_type -> api.core.v1.GetRawdataResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_core_v1_rawdata_proto_init() }
func file_api_core_v1_rawdata_proto_init() {
	if File_api_core_v1_rawdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_core_v1_rawdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawdataRequest); i {
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
		file_api_core_v1_rawdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawdataResponse); i {
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
		file_api_core_v1_rawdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawdataResponse); i {
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
			RawDescriptor: file_api_core_v1_rawdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_core_v1_rawdata_proto_goTypes,
		DependencyIndexes: file_api_core_v1_rawdata_proto_depIdxs,
		MessageInfos:      file_api_core_v1_rawdata_proto_msgTypes,
	}.Build()
	File_api_core_v1_rawdata_proto = out.File
	file_api_core_v1_rawdata_proto_rawDesc = nil
	file_api_core_v1_rawdata_proto_goTypes = nil
	file_api_core_v1_rawdata_proto_depIdxs = nil
}
