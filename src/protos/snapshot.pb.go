// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: snapshot.proto

package protos

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

type Block_ChecksumAlgorithm int32

const (
	Block_SHA256 Block_ChecksumAlgorithm = 0
)

// Enum value maps for Block_ChecksumAlgorithm.
var (
	Block_ChecksumAlgorithm_name = map[int32]string{
		0: "SHA256",
	}
	Block_ChecksumAlgorithm_value = map[string]int32{
		"SHA256": 0,
	}
)

func (x Block_ChecksumAlgorithm) Enum() *Block_ChecksumAlgorithm {
	p := new(Block_ChecksumAlgorithm)
	*p = x
	return p
}

func (x Block_ChecksumAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Block_ChecksumAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_snapshot_proto_enumTypes[0].Descriptor()
}

func (Block_ChecksumAlgorithm) Type() protoreflect.EnumType {
	return &file_snapshot_proto_enumTypes[0]
}

func (x Block_ChecksumAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Block_ChecksumAlgorithm.Descriptor instead.
func (Block_ChecksumAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{3, 0}
}

type Block_Encoding int32

const (
	Block_RAW  Block_Encoding = 0
	Block_GZIP Block_Encoding = 1
	Block_LZ4  Block_Encoding = 2
)

// Enum value maps for Block_Encoding.
var (
	Block_Encoding_name = map[int32]string{
		0: "RAW",
		1: "GZIP",
		2: "LZ4",
	}
	Block_Encoding_value = map[string]int32{
		"RAW":  0,
		"GZIP": 1,
		"LZ4":  2,
	}
)

func (x Block_Encoding) Enum() *Block_Encoding {
	p := new(Block_Encoding)
	*p = x
	return p
}

func (x Block_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Block_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_snapshot_proto_enumTypes[1].Descriptor()
}

func (Block_Encoding) Type() protoreflect.EnumType {
	return &file_snapshot_proto_enumTypes[1]
}

func (x Block_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Block_Encoding.Descriptor instead.
func (Block_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{3, 1}
}

type SnapshotRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Record:
	//	*SnapshotRecord_Header
	//	*SnapshotRecord_BlockList
	//	*SnapshotRecord_Block
	Record isSnapshotRecord_Record `protobuf_oneof:"record"`
}

func (x *SnapshotRecord) Reset() {
	*x = SnapshotRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRecord) ProtoMessage() {}

func (x *SnapshotRecord) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRecord.ProtoReflect.Descriptor instead.
func (*SnapshotRecord) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{0}
}

func (m *SnapshotRecord) GetRecord() isSnapshotRecord_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (x *SnapshotRecord) GetHeader() *SnapshotHeader {
	if x, ok := x.GetRecord().(*SnapshotRecord_Header); ok {
		return x.Header
	}
	return nil
}

func (x *SnapshotRecord) GetBlockList() *BlockList {
	if x, ok := x.GetRecord().(*SnapshotRecord_BlockList); ok {
		return x.BlockList
	}
	return nil
}

func (x *SnapshotRecord) GetBlock() *Block {
	if x, ok := x.GetRecord().(*SnapshotRecord_Block); ok {
		return x.Block
	}
	return nil
}

type isSnapshotRecord_Record interface {
	isSnapshotRecord_Record()
}

type SnapshotRecord_Header struct {
	Header *SnapshotHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type SnapshotRecord_BlockList struct {
	BlockList *BlockList `protobuf:"bytes,2,opt,name=block_list,json=blockList,proto3,oneof"`
}

type SnapshotRecord_Block struct {
	Block *Block `protobuf:"bytes,3,opt,name=block,proto3,oneof"`
}

func (*SnapshotRecord_Header) isSnapshotRecord_Record() {}

func (*SnapshotRecord_BlockList) isSnapshotRecord_Record() {}

func (*SnapshotRecord_Block) isSnapshotRecord_Record() {}

type SnapshotHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version          int32             `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SnapshotId       string            `protobuf:"bytes,2,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
	ParentSnapshotId string            `protobuf:"bytes,3,opt,name=parent_snapshot_id,json=parentSnapshotId,proto3" json:"parent_snapshot_id,omitempty"`
	VolumeId         string            `protobuf:"bytes,4,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	VolumeSize       int64             `protobuf:"varint,5,opt,name=volume_size,json=volumeSize,proto3" json:"volume_size,omitempty"`
	BlockSize        int32             `protobuf:"varint,6,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Description      string            `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Tags             map[string]string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SnapshotHeader) Reset() {
	*x = SnapshotHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotHeader) ProtoMessage() {}

func (x *SnapshotHeader) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotHeader.ProtoReflect.Descriptor instead.
func (*SnapshotHeader) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotHeader) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SnapshotHeader) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

func (x *SnapshotHeader) GetParentSnapshotId() string {
	if x != nil {
		return x.ParentSnapshotId
	}
	return ""
}

func (x *SnapshotHeader) GetVolumeId() string {
	if x != nil {
		return x.VolumeId
	}
	return ""
}

func (x *SnapshotHeader) GetVolumeSize() int64 {
	if x != nil {
		return x.VolumeSize
	}
	return 0
}

func (x *SnapshotHeader) GetBlockSize() int32 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *SnapshotHeader) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SnapshotHeader) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type BlockList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []int32 `protobuf:"varint,1,rep,packed,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlockList) Reset() {
	*x = BlockList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockList) ProtoMessage() {}

func (x *BlockList) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockList.ProtoReflect.Descriptor instead.
func (*BlockList) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *BlockList) GetBlocks() []int32 {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockIndex        int32                   `protobuf:"varint,1,opt,name=block_index,json=blockIndex,proto3" json:"block_index,omitempty"`
	ChecksumAlgorithm Block_ChecksumAlgorithm `protobuf:"varint,2,opt,name=checksum_algorithm,json=checksumAlgorithm,proto3,enum=ebsnapshot.Block_ChecksumAlgorithm" json:"checksum_algorithm,omitempty"`
	DataChecksum      []byte                  `protobuf:"bytes,3,opt,name=data_checksum,json=dataChecksum,proto3" json:"data_checksum,omitempty"`
	DataEncoding      Block_Encoding          `protobuf:"varint,4,opt,name=data_encoding,json=dataEncoding,proto3,enum=ebsnapshot.Block_Encoding" json:"data_encoding,omitempty"`
	Data              []byte                  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetBlockIndex() int32 {
	if x != nil {
		return x.BlockIndex
	}
	return 0
}

func (x *Block) GetChecksumAlgorithm() Block_ChecksumAlgorithm {
	if x != nil {
		return x.ChecksumAlgorithm
	}
	return Block_SHA256
}

func (x *Block) GetDataChecksum() []byte {
	if x != nil {
		return x.DataChecksum
	}
	return nil
}

func (x *Block) GetDataEncoding() Block_Encoding {
	if x != nil {
		return x.DataEncoding
	}
	return Block_RAW
}

func (x *Block) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_snapshot_proto protoreflect.FileDescriptor

var file_snapshot_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x67, 0x6f, 0x73, 0x6e, 0x61, 0x70, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x73, 0x6e, 0x61, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x73, 0x6e, 0x61, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x67, 0x6f, 0x73, 0x6e, 0x61, 0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0xe7, 0x02, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f,
	0x73, 0x6e, 0x61, 0x70, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x09,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x22, 0xb7, 0x02, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4e, 0x0a, 0x12,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x73, 0x6e, 0x61,
	0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75,
	0x6d, 0x12, 0x3b, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x73, 0x6e, 0x61,
	0x70, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x1f, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35,
	0x36, 0x10, 0x00, 0x22, 0x26, 0x0a, 0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x5a, 0x49, 0x50,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x5a, 0x34, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d, 0x67,
	0x6f, 0x73, 0x6e, 0x61, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_snapshot_proto_rawDescOnce sync.Once
	file_snapshot_proto_rawDescData = file_snapshot_proto_rawDesc
)

func file_snapshot_proto_rawDescGZIP() []byte {
	file_snapshot_proto_rawDescOnce.Do(func() {
		file_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_snapshot_proto_rawDescData)
	})
	return file_snapshot_proto_rawDescData
}

var file_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_snapshot_proto_goTypes = []interface{}{
	(Block_ChecksumAlgorithm)(0), // 0: ebsnapshot.Block.ChecksumAlgorithm
	(Block_Encoding)(0),          // 1: ebsnapshot.Block.Encoding
	(*SnapshotRecord)(nil),       // 2: ebsnapshot.SnapshotRecord
	(*SnapshotHeader)(nil),       // 3: ebsnapshot.SnapshotHeader
	(*BlockList)(nil),            // 4: ebsnapshot.BlockList
	(*Block)(nil),                // 5: ebsnapshot.Block
	nil,                          // 6: ebsnapshot.SnapshotHeader.TagsEntry
}
var file_snapshot_proto_depIdxs = []int32{
	3, // 0: ebsnapshot.SnapshotRecord.header:type_name -> ebsnapshot.SnapshotHeader
	4, // 1: ebsnapshot.SnapshotRecord.block_list:type_name -> ebsnapshot.BlockList
	5, // 2: ebsnapshot.SnapshotRecord.block:type_name -> ebsnapshot.Block
	6, // 3: ebsnapshot.SnapshotHeader.tags:type_name -> ebsnapshot.SnapshotHeader.TagsEntry
	0, // 4: ebsnapshot.Block.checksum_algorithm:type_name -> ebsnapshot.Block.ChecksumAlgorithm
	1, // 5: ebsnapshot.Block.data_encoding:type_name -> ebsnapshot.Block.Encoding
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_snapshot_proto_init() }
func file_snapshot_proto_init() {
	if File_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRecord); i {
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
		file_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotHeader); i {
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
		file_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockList); i {
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
		file_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
	file_snapshot_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SnapshotRecord_Header)(nil),
		(*SnapshotRecord_BlockList)(nil),
		(*SnapshotRecord_Block)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_snapshot_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_snapshot_proto_goTypes,
		DependencyIndexes: file_snapshot_proto_depIdxs,
		EnumInfos:         file_snapshot_proto_enumTypes,
		MessageInfos:      file_snapshot_proto_msgTypes,
	}.Build()
	File_snapshot_proto = out.File
	file_snapshot_proto_rawDesc = nil
	file_snapshot_proto_goTypes = nil
	file_snapshot_proto_depIdxs = nil
}
