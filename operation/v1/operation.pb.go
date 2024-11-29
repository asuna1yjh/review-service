// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.20.1
// source: operation/v1/operation.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AuditReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewID  int64   `protobuf:"varint,1,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	Status    int32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OpUser    string  `protobuf:"bytes,3,opt,name=opUser,proto3" json:"opUser,omitempty"`
	OpReason  string  `protobuf:"bytes,4,opt,name=opReason,proto3" json:"opReason,omitempty"`
	OpRemarks *string `protobuf:"bytes,5,opt,name=opRemarks,proto3,oneof" json:"opRemarks,omitempty"`
}

func (x *AuditReviewRequest) Reset() {
	*x = AuditReviewRequest{}
	mi := &file_operation_v1_operation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditReviewRequest) ProtoMessage() {}

func (x *AuditReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operation_v1_operation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditReviewRequest.ProtoReflect.Descriptor instead.
func (*AuditReviewRequest) Descriptor() ([]byte, []int) {
	return file_operation_v1_operation_proto_rawDescGZIP(), []int{0}
}

func (x *AuditReviewRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *AuditReviewRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuditReviewRequest) GetOpUser() string {
	if x != nil {
		return x.OpUser
	}
	return ""
}

func (x *AuditReviewRequest) GetOpReason() string {
	if x != nil {
		return x.OpReason
	}
	return ""
}

func (x *AuditReviewRequest) GetOpRemarks() string {
	if x != nil && x.OpRemarks != nil {
		return *x.OpRemarks
	}
	return ""
}

type AuditReviewReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuditReviewReply) Reset() {
	*x = AuditReviewReply{}
	mi := &file_operation_v1_operation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditReviewReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditReviewReply) ProtoMessage() {}

func (x *AuditReviewReply) ProtoReflect() protoreflect.Message {
	mi := &file_operation_v1_operation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditReviewReply.ProtoReflect.Descriptor instead.
func (*AuditReviewReply) Descriptor() ([]byte, []int) {
	return file_operation_v1_operation_proto_rawDescGZIP(), []int{1}
}

type AuditAppealRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppealID  int64   `protobuf:"varint,1,opt,name=appealID,proto3" json:"appealID,omitempty"`
	ReviewID  int64   `protobuf:"varint,2,opt,name=reviewID,proto3" json:"reviewID,omitempty"`
	Status    int32   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	OpUser    string  `protobuf:"bytes,4,opt,name=opUser,proto3" json:"opUser,omitempty"`
	OpReason  string  `protobuf:"bytes,5,opt,name=opReason,proto3" json:"opReason,omitempty"`
	OpRemarks *string `protobuf:"bytes,6,opt,name=opRemarks,proto3,oneof" json:"opRemarks,omitempty"`
}

func (x *AuditAppealRequest) Reset() {
	*x = AuditAppealRequest{}
	mi := &file_operation_v1_operation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditAppealRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditAppealRequest) ProtoMessage() {}

func (x *AuditAppealRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operation_v1_operation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditAppealRequest.ProtoReflect.Descriptor instead.
func (*AuditAppealRequest) Descriptor() ([]byte, []int) {
	return file_operation_v1_operation_proto_rawDescGZIP(), []int{2}
}

func (x *AuditAppealRequest) GetAppealID() int64 {
	if x != nil {
		return x.AppealID
	}
	return 0
}

func (x *AuditAppealRequest) GetReviewID() int64 {
	if x != nil {
		return x.ReviewID
	}
	return 0
}

func (x *AuditAppealRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuditAppealRequest) GetOpUser() string {
	if x != nil {
		return x.OpUser
	}
	return ""
}

func (x *AuditAppealRequest) GetOpReason() string {
	if x != nil {
		return x.OpReason
	}
	return ""
}

func (x *AuditAppealRequest) GetOpRemarks() string {
	if x != nil && x.OpRemarks != nil {
		return *x.OpRemarks
	}
	return ""
}

// 对申诉进行审核的返回值
type AuditAppealReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuditAppealReply) Reset() {
	*x = AuditAppealReply{}
	mi := &file_operation_v1_operation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditAppealReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditAppealReply) ProtoMessage() {}

func (x *AuditAppealReply) ProtoReflect() protoreflect.Message {
	mi := &file_operation_v1_operation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditAppealReply.ProtoReflect.Descriptor instead.
func (*AuditAppealReply) Descriptor() ([]byte, []int) {
	return file_operation_v1_operation_proto_rawDescGZIP(), []int{3}
}

var File_operation_v1_operation_proto protoreflect.FileDescriptor

var file_operation_v1_operation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x06, 0x6f,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02,
	0x52, 0x08, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x6f, 0x70,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x6f, 0x70, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6f, 0x70, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0xf6, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44,
	0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x06, 0x6f, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x08, 0x6f,
	0x70, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x6f, 0x70, 0x52, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x70,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6f,
	0x70, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x41, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf9, 0x01, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x0b, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x12, 0x75, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x70, 0x70, 0x65, 0x61, 0x6c,
	0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41,
	0x70, 0x70, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x65,
	0x61, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x32, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1c,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operation_v1_operation_proto_rawDescOnce sync.Once
	file_operation_v1_operation_proto_rawDescData = file_operation_v1_operation_proto_rawDesc
)

func file_operation_v1_operation_proto_rawDescGZIP() []byte {
	file_operation_v1_operation_proto_rawDescOnce.Do(func() {
		file_operation_v1_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_operation_v1_operation_proto_rawDescData)
	})
	return file_operation_v1_operation_proto_rawDescData
}

var file_operation_v1_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_operation_v1_operation_proto_goTypes = []any{
	(*AuditReviewRequest)(nil), // 0: api.operation.v1.AuditReviewRequest
	(*AuditReviewReply)(nil),   // 1: api.operation.v1.AuditReviewReply
	(*AuditAppealRequest)(nil), // 2: api.operation.v1.AuditAppealRequest
	(*AuditAppealReply)(nil),   // 3: api.operation.v1.AuditAppealReply
}
var file_operation_v1_operation_proto_depIdxs = []int32{
	0, // 0: api.operation.v1.Operation.AuditReview:input_type -> api.operation.v1.AuditReviewRequest
	2, // 1: api.operation.v1.Operation.AuditAppeal:input_type -> api.operation.v1.AuditAppealRequest
	1, // 2: api.operation.v1.Operation.AuditReview:output_type -> api.operation.v1.AuditReviewReply
	3, // 3: api.operation.v1.Operation.AuditAppeal:output_type -> api.operation.v1.AuditAppealReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_operation_v1_operation_proto_init() }
func file_operation_v1_operation_proto_init() {
	if File_operation_v1_operation_proto != nil {
		return
	}
	file_operation_v1_operation_proto_msgTypes[0].OneofWrappers = []any{}
	file_operation_v1_operation_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operation_v1_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operation_v1_operation_proto_goTypes,
		DependencyIndexes: file_operation_v1_operation_proto_depIdxs,
		MessageInfos:      file_operation_v1_operation_proto_msgTypes,
	}.Build()
	File_operation_v1_operation_proto = out.File
	file_operation_v1_operation_proto_rawDesc = nil
	file_operation_v1_operation_proto_goTypes = nil
	file_operation_v1_operation_proto_depIdxs = nil
}
