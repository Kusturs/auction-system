// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.27.1
// source: auction.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Auction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LotId        int64                  `protobuf:"varint,2,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	StartPrice   float64                `protobuf:"fixed64,3,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	MinStep      float64                `protobuf:"fixed64,4,opt,name=min_step,json=minStep,proto3" json:"min_step,omitempty"`
	CurrentPrice float64                `protobuf:"fixed64,5,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status       string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	WinnerId     int64                  `protobuf:"varint,9,opt,name=winner_id,json=winnerId,proto3" json:"winner_id,omitempty"`
	WinnerBidId  int64                  `protobuf:"varint,10,opt,name=winner_bid_id,json=winnerBidId,proto3" json:"winner_bid_id,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Auction) Reset() {
	*x = Auction{}
	mi := &file_auction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{0}
}

func (x *Auction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Auction) GetLotId() int64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

func (x *Auction) GetStartPrice() float64 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *Auction) GetMinStep() float64 {
	if x != nil {
		return x.MinStep
	}
	return 0
}

func (x *Auction) GetCurrentPrice() float64 {
	if x != nil {
		return x.CurrentPrice
	}
	return 0
}

func (x *Auction) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Auction) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Auction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Auction) GetWinnerId() int64 {
	if x != nil {
		return x.WinnerId
	}
	return 0
}

func (x *Auction) GetWinnerBidId() int64 {
	if x != nil {
		return x.WinnerBidId
	}
	return 0
}

func (x *Auction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Auction) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotId      int64                  `protobuf:"varint,1,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	StartPrice float64                `protobuf:"fixed64,2,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	MinStep    float64                `protobuf:"fixed64,3,opt,name=min_step,json=minStep,proto3" json:"min_step,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *CreateAuctionRequest) Reset() {
	*x = CreateAuctionRequest{}
	mi := &file_auction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionRequest) ProtoMessage() {}

func (x *CreateAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionRequest.ProtoReflect.Descriptor instead.
func (*CreateAuctionRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuctionRequest) GetLotId() int64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

func (x *CreateAuctionRequest) GetStartPrice() float64 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *CreateAuctionRequest) GetMinStep() float64 {
	if x != nil {
		return x.MinStep
	}
	return 0
}

func (x *CreateAuctionRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CreateAuctionRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type CreateAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
}

func (x *CreateAuctionResponse) Reset() {
	*x = CreateAuctionResponse{}
	mi := &file_auction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionResponse) ProtoMessage() {}

func (x *CreateAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionResponse.ProtoReflect.Descriptor instead.
func (*CreateAuctionResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuctionResponse) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

type GetAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuctionRequest) Reset() {
	*x = GetAuctionRequest{}
	mi := &file_auction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionRequest) ProtoMessage() {}

func (x *GetAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionRequest.ProtoReflect.Descriptor instead.
func (*GetAuctionRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuctionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
}

func (x *GetAuctionResponse) Reset() {
	*x = GetAuctionResponse{}
	mi := &file_auction_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuctionResponse) ProtoMessage() {}

func (x *GetAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuctionResponse.ProtoReflect.Descriptor instead.
func (*GetAuctionResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuctionResponse) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

type UpdateAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartPrice float64                `protobuf:"fixed64,2,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	MinStep    float64                `protobuf:"fixed64,3,opt,name=min_step,json=minStep,proto3" json:"min_step,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status     string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateAuctionRequest) Reset() {
	*x = UpdateAuctionRequest{}
	mi := &file_auction_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuctionRequest) ProtoMessage() {}

func (x *UpdateAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuctionRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuctionRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAuctionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAuctionRequest) GetStartPrice() float64 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *UpdateAuctionRequest) GetMinStep() float64 {
	if x != nil {
		return x.MinStep
	}
	return 0
}

func (x *UpdateAuctionRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *UpdateAuctionRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *UpdateAuctionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auction *Auction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction,omitempty"`
}

func (x *UpdateAuctionResponse) Reset() {
	*x = UpdateAuctionResponse{}
	mi := &file_auction_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuctionResponse) ProtoMessage() {}

func (x *UpdateAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuctionResponse.ProtoReflect.Descriptor instead.
func (*UpdateAuctionResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAuctionResponse) GetAuction() *Auction {
	if x != nil {
		return x.Auction
	}
	return nil
}

type ListAuctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize   int32   `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNumber int32   `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	Status     *string `protobuf:"bytes,3,opt,name=status,proto3,oneof" json:"status,omitempty"`
}

func (x *ListAuctionsRequest) Reset() {
	*x = ListAuctionsRequest{}
	mi := &file_auction_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuctionsRequest) ProtoMessage() {}

func (x *ListAuctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuctionsRequest.ProtoReflect.Descriptor instead.
func (*ListAuctionsRequest) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{7}
}

func (x *ListAuctionsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAuctionsRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListAuctionsRequest) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type ListAuctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auctions   []*Auction `protobuf:"bytes,1,rep,name=auctions,proto3" json:"auctions,omitempty"`
	TotalCount int32      `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListAuctionsResponse) Reset() {
	*x = ListAuctionsResponse{}
	mi := &file_auction_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuctionsResponse) ProtoMessage() {}

func (x *ListAuctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuctionsResponse.ProtoReflect.Descriptor instead.
func (*ListAuctionsResponse) Descriptor() ([]byte, []int) {
	return file_auction_proto_rawDescGZIP(), []int{8}
}

func (x *ListAuctionsResponse) GetAuctions() []*Auction {
	if x != nil {
		return x.Auctions
	}
	return nil
}

func (x *ListAuctionsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_auction_proto protoreflect.FileDescriptor

var file_auction_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x03, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x69, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d,
	0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x42, 0x69, 0x64, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdb, 0x01, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xec, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x65, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xbc,
	0x03, 0x0a, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x64,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x70, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x1a,
	0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x18, 0x5a,
	0x16, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auction_proto_rawDescOnce sync.Once
	file_auction_proto_rawDescData = file_auction_proto_rawDesc
)

func file_auction_proto_rawDescGZIP() []byte {
	file_auction_proto_rawDescOnce.Do(func() {
		file_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_auction_proto_rawDescData)
	})
	return file_auction_proto_rawDescData
}

var file_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auction_proto_goTypes = []any{
	(*Auction)(nil),               // 0: auction.Auction
	(*CreateAuctionRequest)(nil),  // 1: auction.CreateAuctionRequest
	(*CreateAuctionResponse)(nil), // 2: auction.CreateAuctionResponse
	(*GetAuctionRequest)(nil),     // 3: auction.GetAuctionRequest
	(*GetAuctionResponse)(nil),    // 4: auction.GetAuctionResponse
	(*UpdateAuctionRequest)(nil),  // 5: auction.UpdateAuctionRequest
	(*UpdateAuctionResponse)(nil), // 6: auction.UpdateAuctionResponse
	(*ListAuctionsRequest)(nil),   // 7: auction.ListAuctionsRequest
	(*ListAuctionsResponse)(nil),  // 8: auction.ListAuctionsResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_auction_proto_depIdxs = []int32{
	9,  // 0: auction.Auction.start_time:type_name -> google.protobuf.Timestamp
	9,  // 1: auction.Auction.end_time:type_name -> google.protobuf.Timestamp
	9,  // 2: auction.Auction.created_at:type_name -> google.protobuf.Timestamp
	9,  // 3: auction.Auction.updated_at:type_name -> google.protobuf.Timestamp
	9,  // 4: auction.CreateAuctionRequest.start_time:type_name -> google.protobuf.Timestamp
	9,  // 5: auction.CreateAuctionRequest.end_time:type_name -> google.protobuf.Timestamp
	0,  // 6: auction.CreateAuctionResponse.auction:type_name -> auction.Auction
	0,  // 7: auction.GetAuctionResponse.auction:type_name -> auction.Auction
	9,  // 8: auction.UpdateAuctionRequest.start_time:type_name -> google.protobuf.Timestamp
	9,  // 9: auction.UpdateAuctionRequest.end_time:type_name -> google.protobuf.Timestamp
	0,  // 10: auction.UpdateAuctionResponse.auction:type_name -> auction.Auction
	0,  // 11: auction.ListAuctionsResponse.auctions:type_name -> auction.Auction
	1,  // 12: auction.AuctionService.CreateAuction:input_type -> auction.CreateAuctionRequest
	3,  // 13: auction.AuctionService.GetAuction:input_type -> auction.GetAuctionRequest
	5,  // 14: auction.AuctionService.UpdateAuction:input_type -> auction.UpdateAuctionRequest
	7,  // 15: auction.AuctionService.ListAuctions:input_type -> auction.ListAuctionsRequest
	2,  // 16: auction.AuctionService.CreateAuction:output_type -> auction.CreateAuctionResponse
	4,  // 17: auction.AuctionService.GetAuction:output_type -> auction.GetAuctionResponse
	6,  // 18: auction.AuctionService.UpdateAuction:output_type -> auction.UpdateAuctionResponse
	8,  // 19: auction.AuctionService.ListAuctions:output_type -> auction.ListAuctionsResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_auction_proto_init() }
func file_auction_proto_init() {
	if File_auction_proto != nil {
		return
	}
	file_auction_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auction_proto_goTypes,
		DependencyIndexes: file_auction_proto_depIdxs,
		MessageInfos:      file_auction_proto_msgTypes,
	}.Build()
	File_auction_proto = out.File
	file_auction_proto_rawDesc = nil
	file_auction_proto_goTypes = nil
	file_auction_proto_depIdxs = nil
}
