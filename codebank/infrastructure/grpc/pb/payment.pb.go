// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: protofile/payment.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreditCard  *PaymentRequest_CreditCard `protobuf:"bytes,1,opt,name=creditCard,proto3" json:"creditCard,omitempty"`
	Amount      float64                    `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Store       string                     `protobuf:"bytes,3,opt,name=store,proto3" json:"store,omitempty"`
	Description string                     `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofile_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protofile_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_protofile_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetCreditCard() *PaymentRequest_CreditCard {
	if x != nil {
		return x.CreditCard
	}
	return nil
}

func (x *PaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *PaymentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PaymentRequest_CreditCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number          string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	ExpirationMonth int32  `protobuf:"varint,3,opt,name=expirationMonth,proto3" json:"expirationMonth,omitempty"`
	ExpirationYear  int32  `protobuf:"varint,4,opt,name=expirationYear,proto3" json:"expirationYear,omitempty"`
	Cvv             int32  `protobuf:"varint,5,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *PaymentRequest_CreditCard) Reset() {
	*x = PaymentRequest_CreditCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofile_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest_CreditCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest_CreditCard) ProtoMessage() {}

func (x *PaymentRequest_CreditCard) ProtoReflect() protoreflect.Message {
	mi := &file_protofile_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest_CreditCard.ProtoReflect.Descriptor instead.
func (*PaymentRequest_CreditCard) Descriptor() ([]byte, []int) {
	return file_protofile_payment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PaymentRequest_CreditCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaymentRequest_CreditCard) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PaymentRequest_CreditCard) GetExpirationMonth() int32 {
	if x != nil {
		return x.ExpirationMonth
	}
	return 0
}

func (x *PaymentRequest_CreditCard) GetExpirationYear() int32 {
	if x != nil {
		return x.ExpirationYear
	}
	return 0
}

func (x *PaymentRequest_CreditCard) GetCvv() int32 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

var File_protofile_payment_proto protoreflect.FileDescriptor

var file_protofile_payment_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x9c, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x63, 0x76, 0x76, 0x32, 0x4e, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofile_payment_proto_rawDescOnce sync.Once
	file_protofile_payment_proto_rawDescData = file_protofile_payment_proto_rawDesc
)

func file_protofile_payment_proto_rawDescGZIP() []byte {
	file_protofile_payment_proto_rawDescOnce.Do(func() {
		file_protofile_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofile_payment_proto_rawDescData)
	})
	return file_protofile_payment_proto_rawDescData
}

var file_protofile_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protofile_payment_proto_goTypes = []interface{}{
	(*PaymentRequest)(nil),            // 0: payment.PaymentRequest
	(*PaymentRequest_CreditCard)(nil), // 1: payment.PaymentRequest.CreditCard
	(*empty.Empty)(nil),               // 2: google.protobuf.Empty
}
var file_protofile_payment_proto_depIdxs = []int32{
	1, // 0: payment.PaymentRequest.creditCard:type_name -> payment.PaymentRequest.CreditCard
	0, // 1: payment.PaymentService.Payment:input_type -> payment.PaymentRequest
	2, // 2: payment.PaymentService.Payment:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protofile_payment_proto_init() }
func file_protofile_payment_proto_init() {
	if File_protofile_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofile_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
		file_protofile_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest_CreditCard); i {
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
			RawDescriptor: file_protofile_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofile_payment_proto_goTypes,
		DependencyIndexes: file_protofile_payment_proto_depIdxs,
		MessageInfos:      file_protofile_payment_proto_msgTypes,
	}.Build()
	File_protofile_payment_proto = out.File
	file_protofile_payment_proto_rawDesc = nil
	file_protofile_payment_proto_goTypes = nil
	file_protofile_payment_proto_depIdxs = nil
}
