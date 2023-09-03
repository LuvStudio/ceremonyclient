// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: keys.proto

package protobufs

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

// Describes a raw Ed448 public key
type Ed448PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 57 byte value
}

func (x *Ed448PublicKey) Reset() {
	*x = Ed448PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed448PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed448PublicKey) ProtoMessage() {}

func (x *Ed448PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed448PublicKey.ProtoReflect.Descriptor instead.
func (*Ed448PublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{0}
}

func (x *Ed448PublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// Describes a raw Ed448 private key – notably this is post-derivation,
// not the seed.
type Ed448PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  []byte          `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 57 byte value
	PublicKey *Ed448PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Ed448PrivateKey) Reset() {
	*x = Ed448PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed448PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed448PrivateKey) ProtoMessage() {}

func (x *Ed448PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed448PrivateKey.ProtoReflect.Descriptor instead.
func (*Ed448PrivateKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{1}
}

func (x *Ed448PrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *Ed448PrivateKey) GetPublicKey() *Ed448PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Describes a raw Ed448 signature
type Ed448Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte          `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"` // 114 byte value
	PublicKey *Ed448PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Ed448Signature) Reset() {
	*x = Ed448Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ed448Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ed448Signature) ProtoMessage() {}

func (x *Ed448Signature) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ed448Signature.ProtoReflect.Descriptor instead.
func (*Ed448Signature) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{2}
}

func (x *Ed448Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Ed448Signature) GetPublicKey() *Ed448PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Describes a raw X448 public key
type X448PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 57 byte value
}

func (x *X448PublicKey) Reset() {
	*x = X448PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X448PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X448PublicKey) ProtoMessage() {}

func (x *X448PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X448PublicKey.ProtoReflect.Descriptor instead.
func (*X448PublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{3}
}

func (x *X448PublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// Describes a raw X448 private key – notably this is post-derivation,
// not the seed.
type X448PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  []byte         `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 57 byte value
	PublicKey *X448PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *X448PrivateKey) Reset() {
	*x = X448PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *X448PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X448PrivateKey) ProtoMessage() {}

func (x *X448PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X448PrivateKey.ProtoReflect.Descriptor instead.
func (*X448PrivateKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{4}
}

func (x *X448PrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *X448PrivateKey) GetPublicKey() *X448PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Describes a raw PCAS public key
type PCASPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 256 kilobyte value
}

func (x *PCASPublicKey) Reset() {
	*x = PCASPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCASPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCASPublicKey) ProtoMessage() {}

func (x *PCASPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCASPublicKey.ProtoReflect.Descriptor instead.
func (*PCASPublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{5}
}

func (x *PCASPublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// Describes a raw PCAS private key
type PCASPrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  []byte         `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 256 byte value
	PublicKey *PCASPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *PCASPrivateKey) Reset() {
	*x = PCASPrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PCASPrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PCASPrivateKey) ProtoMessage() {}

func (x *PCASPrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PCASPrivateKey.ProtoReflect.Descriptor instead.
func (*PCASPrivateKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{6}
}

func (x *PCASPrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *PCASPrivateKey) GetPublicKey() *PCASPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Describes a raw compressed BLS48-581 G1 public key
type BLS48581G1PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 74 byte value
}

func (x *BLS48581G1PublicKey) Reset() {
	*x = BLS48581G1PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS48581G1PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS48581G1PublicKey) ProtoMessage() {}

func (x *BLS48581G1PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS48581G1PublicKey.ProtoReflect.Descriptor instead.
func (*BLS48581G1PublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{7}
}

func (x *BLS48581G1PublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// Describes a raw BLS48-581 private key, with corresponding G1 public key
type BLS48581G1PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  []byte               `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 73 byte value
	PublicKey *BLS48581G1PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *BLS48581G1PrivateKey) Reset() {
	*x = BLS48581G1PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS48581G1PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS48581G1PrivateKey) ProtoMessage() {}

func (x *BLS48581G1PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS48581G1PrivateKey.ProtoReflect.Descriptor instead.
func (*BLS48581G1PrivateKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{8}
}

func (x *BLS48581G1PrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *BLS48581G1PrivateKey) GetPublicKey() *BLS48581G1PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

// Describes a raw compressed BLS48-581 G2 public key
type BLS48581G2PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 585 byte value
}

func (x *BLS48581G2PublicKey) Reset() {
	*x = BLS48581G2PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS48581G2PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS48581G2PublicKey) ProtoMessage() {}

func (x *BLS48581G2PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS48581G2PublicKey.ProtoReflect.Descriptor instead.
func (*BLS48581G2PublicKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{9}
}

func (x *BLS48581G2PublicKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

// Describes a raw BLS48-581 private key, with corresponding G2 public key
type BLS48581G2PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyValue  []byte               `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // 73 byte value
	PublicKey *BLS48581G2PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *BLS48581G2PrivateKey) Reset() {
	*x = BLS48581G2PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BLS48581G2PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BLS48581G2PrivateKey) ProtoMessage() {}

func (x *BLS48581G2PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BLS48581G2PrivateKey.ProtoReflect.Descriptor instead.
func (*BLS48581G2PrivateKey) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{10}
}

func (x *BLS48581G2PrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

func (x *BLS48581G2PrivateKey) GetPublicKey() *BLS48581G2PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

var File_keys_proto protoreflect.FileDescriptor

var file_keys_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x71, 0x75,
	0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65,
	0x79, 0x73, 0x2e, 0x70, 0x62, 0x22, 0x2d, 0x0a, 0x0e, 0x45, 0x64, 0x34, 0x34, 0x38, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x76, 0x0a, 0x0f, 0x45, 0x64, 0x34, 0x34, 0x38, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69,
	0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x64, 0x34, 0x34, 0x38, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x76, 0x0a, 0x0e,
	0x45, 0x64, 0x34, 0x34, 0x38, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x46, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x34, 0x34, 0x38,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x0d, 0x58, 0x34, 0x34, 0x38, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x74, 0x0a, 0x0e, 0x58, 0x34, 0x34, 0x38, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69,
	0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e,
	0x58, 0x34, 0x34, 0x38, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x0d, 0x50, 0x43, 0x41, 0x53,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x74, 0x0a, 0x0e, 0x50, 0x43, 0x41, 0x53, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x69, 0x6c,
	0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x43, 0x41, 0x53, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x13,
	0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47, 0x31, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x80, 0x01, 0x0a, 0x14, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47, 0x31, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x75, 0x69,
	0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47, 0x31, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x32, 0x0a, 0x13, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47,
	0x32, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65,
	0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x42, 0x4c, 0x53, 0x34,
	0x38, 0x35, 0x38, 0x31, 0x47, 0x32, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x34,
	0x38, 0x35, 0x38, 0x31, 0x47, 0x32, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x3a, 0x5a, 0x38, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2f, 0x6d,
	0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keys_proto_rawDescOnce sync.Once
	file_keys_proto_rawDescData = file_keys_proto_rawDesc
)

func file_keys_proto_rawDescGZIP() []byte {
	file_keys_proto_rawDescOnce.Do(func() {
		file_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_keys_proto_rawDescData)
	})
	return file_keys_proto_rawDescData
}

var file_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_keys_proto_goTypes = []interface{}{
	(*Ed448PublicKey)(nil),       // 0: quilibrium.node.keys.pb.Ed448PublicKey
	(*Ed448PrivateKey)(nil),      // 1: quilibrium.node.keys.pb.Ed448PrivateKey
	(*Ed448Signature)(nil),       // 2: quilibrium.node.keys.pb.Ed448Signature
	(*X448PublicKey)(nil),        // 3: quilibrium.node.keys.pb.X448PublicKey
	(*X448PrivateKey)(nil),       // 4: quilibrium.node.keys.pb.X448PrivateKey
	(*PCASPublicKey)(nil),        // 5: quilibrium.node.keys.pb.PCASPublicKey
	(*PCASPrivateKey)(nil),       // 6: quilibrium.node.keys.pb.PCASPrivateKey
	(*BLS48581G1PublicKey)(nil),  // 7: quilibrium.node.keys.pb.BLS48581G1PublicKey
	(*BLS48581G1PrivateKey)(nil), // 8: quilibrium.node.keys.pb.BLS48581G1PrivateKey
	(*BLS48581G2PublicKey)(nil),  // 9: quilibrium.node.keys.pb.BLS48581G2PublicKey
	(*BLS48581G2PrivateKey)(nil), // 10: quilibrium.node.keys.pb.BLS48581G2PrivateKey
}
var file_keys_proto_depIdxs = []int32{
	0, // 0: quilibrium.node.keys.pb.Ed448PrivateKey.public_key:type_name -> quilibrium.node.keys.pb.Ed448PublicKey
	0, // 1: quilibrium.node.keys.pb.Ed448Signature.public_key:type_name -> quilibrium.node.keys.pb.Ed448PublicKey
	3, // 2: quilibrium.node.keys.pb.X448PrivateKey.public_key:type_name -> quilibrium.node.keys.pb.X448PublicKey
	5, // 3: quilibrium.node.keys.pb.PCASPrivateKey.public_key:type_name -> quilibrium.node.keys.pb.PCASPublicKey
	7, // 4: quilibrium.node.keys.pb.BLS48581G1PrivateKey.public_key:type_name -> quilibrium.node.keys.pb.BLS48581G1PublicKey
	9, // 5: quilibrium.node.keys.pb.BLS48581G2PrivateKey.public_key:type_name -> quilibrium.node.keys.pb.BLS48581G2PublicKey
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_keys_proto_init() }
func file_keys_proto_init() {
	if File_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed448PublicKey); i {
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
		file_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed448PrivateKey); i {
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
		file_keys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ed448Signature); i {
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
		file_keys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X448PublicKey); i {
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
		file_keys_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*X448PrivateKey); i {
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
		file_keys_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCASPublicKey); i {
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
		file_keys_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PCASPrivateKey); i {
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
		file_keys_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS48581G1PublicKey); i {
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
		file_keys_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS48581G1PrivateKey); i {
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
		file_keys_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS48581G2PublicKey); i {
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
		file_keys_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BLS48581G2PrivateKey); i {
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
			RawDescriptor: file_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keys_proto_goTypes,
		DependencyIndexes: file_keys_proto_depIdxs,
		MessageInfos:      file_keys_proto_msgTypes,
	}.Build()
	File_keys_proto = out.File
	file_keys_proto_rawDesc = nil
	file_keys_proto_goTypes = nil
	file_keys_proto_depIdxs = nil
}
