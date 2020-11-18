// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conn.ext.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PackageType int32

const (
	PackageType_PT_UNKNOWN   PackageType = 0
	PackageType_PT_SIGN_IN   PackageType = 1
	PackageType_PT_SYNC      PackageType = 2
	PackageType_PT_HEARTBEAT PackageType = 3
	PackageType_PT_MESSAGE   PackageType = 4
)

var PackageType_name = map[int32]string{
	0: "PT_UNKNOWN",
	1: "PT_SIGN_IN",
	2: "PT_SYNC",
	3: "PT_HEARTBEAT",
	4: "PT_MESSAGE",
}

var PackageType_value = map[string]int32{
	"PT_UNKNOWN":   0,
	"PT_SIGN_IN":   1,
	"PT_SYNC":      2,
	"PT_HEARTBEAT": 3,
	"PT_MESSAGE":   4,
}

func (x PackageType) String() string {
	return proto.EnumName(PackageType_name, int32(x))
}

func (PackageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{0}
}

// 消息类型
type MessageType int32

const (
	MessageType_MT_UNKNOWN  MessageType = 0
	MessageType_MT_TEXT     MessageType = 1
	MessageType_MT_FACE     MessageType = 2
	MessageType_MT_VOICE    MessageType = 3
	MessageType_MT_IMAGE    MessageType = 4
	MessageType_MT_FILE     MessageType = 5
	MessageType_MT_LOCATION MessageType = 6
	MessageType_MT_COMMAND  MessageType = 7
	MessageType_MT_CUSTOM   MessageType = 8
)

var MessageType_name = map[int32]string{
	0: "MT_UNKNOWN",
	1: "MT_TEXT",
	2: "MT_FACE",
	3: "MT_VOICE",
	4: "MT_IMAGE",
	5: "MT_FILE",
	6: "MT_LOCATION",
	7: "MT_COMMAND",
	8: "MT_CUSTOM",
}

var MessageType_value = map[string]int32{
	"MT_UNKNOWN":  0,
	"MT_TEXT":     1,
	"MT_FACE":     2,
	"MT_VOICE":    3,
	"MT_IMAGE":    4,
	"MT_FILE":     5,
	"MT_LOCATION": 6,
	"MT_COMMAND":  7,
	"MT_CUSTOM":   8,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{1}
}

type ReceiverType int32

const (
	ReceiverType_RT_UNKNOWN      ReceiverType = 0
	ReceiverType_RT_USER         ReceiverType = 1
	ReceiverType_RT_NORMAL_GROUP ReceiverType = 2
	ReceiverType_RT_LARGE_GROUP  ReceiverType = 3
)

var ReceiverType_name = map[int32]string{
	0: "RT_UNKNOWN",
	1: "RT_USER",
	2: "RT_NORMAL_GROUP",
	3: "RT_LARGE_GROUP",
}

var ReceiverType_value = map[string]int32{
	"RT_UNKNOWN":      0,
	"RT_USER":         1,
	"RT_NORMAL_GROUP": 2,
	"RT_LARGE_GROUP":  3,
}

func (x ReceiverType) String() string {
	return proto.EnumName(ReceiverType_name, int32(x))
}

func (ReceiverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{2}
}

type SenderType int32

const (
	SenderType_ST_UNKNOWN  SenderType = 0
	SenderType_ST_SYSTEM   SenderType = 1
	SenderType_ST_USER     SenderType = 2
	SenderType_ST_BUSINESS SenderType = 3
)

var SenderType_name = map[int32]string{
	0: "ST_UNKNOWN",
	1: "ST_SYSTEM",
	2: "ST_USER",
	3: "ST_BUSINESS",
}

var SenderType_value = map[string]int32{
	"ST_UNKNOWN":  0,
	"ST_SYSTEM":   1,
	"ST_USER":     2,
	"ST_BUSINESS": 3,
}

func (x SenderType) String() string {
	return proto.EnumName(SenderType_name, int32(x))
}

func (SenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{3}
}

type MessageStatus int32

const (
	MessageStatus_MS_UNKNOWN MessageStatus = 0
	MessageStatus_MS_NORMAL  MessageStatus = 1
	MessageStatus_MS_RECALL  MessageStatus = 2
)

var MessageStatus_name = map[int32]string{
	0: "MS_UNKNOWN",
	1: "MS_NORMAL",
	2: "MS_RECALL",
}

var MessageStatus_value = map[string]int32{
	"MS_UNKNOWN": 0,
	"MS_NORMAL":  1,
	"MS_RECALL":  2,
}

func (x MessageStatus) String() string {
	return proto.EnumName(MessageStatus_name, int32(x))
}

func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{4}
}

// 单条消息投递内容（估算大约100个字节）,todo 通知栏提醒
type Message struct {
	Sender               *Sender       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	ReceiverType         ReceiverType  `protobuf:"varint,2,opt,name=receiver_type,json=receiverType,proto3,enum=pb.ReceiverType" json:"receiver_type,omitempty"`
	ReceiverId           int64         `protobuf:"varint,3,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	ToUserIds            []int64       `protobuf:"varint,4,rep,packed,name=to_user_ids,json=toUserIds,proto3" json:"to_user_ids,omitempty"`
	MessageType          MessageType   `protobuf:"varint,5,opt,name=message_type,json=messageType,proto3,enum=pb.MessageType" json:"message_type,omitempty"`
	MessageContent       []byte        `protobuf:"bytes,6,opt,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"`
	Seq                  int64         `protobuf:"varint,7,opt,name=seq,proto3" json:"seq,omitempty"`
	SendTime             int64         `protobuf:"varint,8,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Status               MessageStatus `protobuf:"varint,9,opt,name=status,proto3,enum=pb.MessageStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSender() *Sender {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Message) GetReceiverType() ReceiverType {
	if m != nil {
		return m.ReceiverType
	}
	return ReceiverType_RT_UNKNOWN
}

func (m *Message) GetReceiverId() int64 {
	if m != nil {
		return m.ReceiverId
	}
	return 0
}

func (m *Message) GetToUserIds() []int64 {
	if m != nil {
		return m.ToUserIds
	}
	return nil
}

func (m *Message) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_MT_UNKNOWN
}

func (m *Message) GetMessageContent() []byte {
	if m != nil {
		return m.MessageContent
	}
	return nil
}

func (m *Message) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Message) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *Message) GetStatus() MessageStatus {
	if m != nil {
		return m.Status
	}
	return MessageStatus_MS_UNKNOWN
}

type Sender struct {
	SenderType           SenderType `protobuf:"varint,1,opt,name=sender_type,json=senderType,proto3,enum=pb.SenderType" json:"sender_type,omitempty"`
	SenderId             int64      `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	AvatarUrl            string     `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Nickname             string     `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Extra                string     `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Sender) Reset()         { *m = Sender{} }
func (m *Sender) String() string { return proto.CompactTextString(m) }
func (*Sender) ProtoMessage()    {}
func (*Sender) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{1}
}

func (m *Sender) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sender.Unmarshal(m, b)
}
func (m *Sender) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sender.Marshal(b, m, deterministic)
}
func (m *Sender) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sender.Merge(m, src)
}
func (m *Sender) XXX_Size() int {
	return xxx_messageInfo_Sender.Size(m)
}
func (m *Sender) XXX_DiscardUnknown() {
	xxx_messageInfo_Sender.DiscardUnknown(m)
}

var xxx_messageInfo_Sender proto.InternalMessageInfo

func (m *Sender) GetSenderType() SenderType {
	if m != nil {
		return m.SenderType
	}
	return SenderType_ST_UNKNOWN
}

func (m *Sender) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Sender) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Sender) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Sender) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

// 文本消息
type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{2}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// 表情消息
type Face struct {
	FaceId               int64    `protobuf:"varint,1,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`
	FaceUrl              string   `protobuf:"bytes,2,opt,name=face_url,json=faceUrl,proto3" json:"face_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Face) Reset()         { *m = Face{} }
func (m *Face) String() string { return proto.CompactTextString(m) }
func (*Face) ProtoMessage()    {}
func (*Face) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{3}
}

func (m *Face) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Face.Unmarshal(m, b)
}
func (m *Face) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Face.Marshal(b, m, deterministic)
}
func (m *Face) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Face.Merge(m, src)
}
func (m *Face) XXX_Size() int {
	return xxx_messageInfo_Face.Size(m)
}
func (m *Face) XXX_DiscardUnknown() {
	xxx_messageInfo_Face.DiscardUnknown(m)
}

var xxx_messageInfo_Face proto.InternalMessageInfo

func (m *Face) GetFaceId() int64 {
	if m != nil {
		return m.FaceId
	}
	return 0
}

func (m *Face) GetFaceUrl() string {
	if m != nil {
		return m.FaceUrl
	}
	return ""
}

// 语音消息
type Voice struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Duration             int32    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Voice) Reset()         { *m = Voice{} }
func (m *Voice) String() string { return proto.CompactTextString(m) }
func (*Voice) ProtoMessage()    {}
func (*Voice) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{4}
}

func (m *Voice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voice.Unmarshal(m, b)
}
func (m *Voice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voice.Marshal(b, m, deterministic)
}
func (m *Voice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voice.Merge(m, src)
}
func (m *Voice) XXX_Size() int {
	return xxx_messageInfo_Voice.Size(m)
}
func (m *Voice) XXX_DiscardUnknown() {
	xxx_messageInfo_Voice.DiscardUnknown(m)
}

var xxx_messageInfo_Voice proto.InternalMessageInfo

func (m *Voice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Voice) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Voice) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Voice) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// 图片消息
type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailUrl         string   `protobuf:"bytes,5,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{5}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

// 文件消息
type File struct {
	Id                   int64    `protobuf:"varint,12,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,14,opt,name=size,proto3" json:"size,omitempty"`
	Url                  string   `protobuf:"bytes,15,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{6}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// 地理位置消息
type Location struct {
	Desc                 string   `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{7}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// Command 指令推送，1000以下，IM内部用，1000以上，留给业务用
type Command struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{8}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Command) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 自定义消息
type Custom struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Custom) Reset()         { *m = Custom{} }
func (m *Custom) String() string { return proto.CompactTextString(m) }
func (*Custom) ProtoMessage()    {}
func (*Custom) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{9}
}

func (m *Custom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Custom.Unmarshal(m, b)
}
func (m *Custom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Custom.Marshal(b, m, deterministic)
}
func (m *Custom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Custom.Merge(m, src)
}
func (m *Custom) XXX_Size() int {
	return xxx_messageInfo_Custom.Size(m)
}
func (m *Custom) XXX_DiscardUnknown() {
	xxx_messageInfo_Custom.DiscardUnknown(m)
}

var xxx_messageInfo_Custom proto.InternalMessageInfo

func (m *Custom) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// 上行数据
type Input struct {
	Type                 PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PackageType" json:"type,omitempty"`
	RequestId            int64       `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Data                 []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{10}
}

func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetType() PackageType {
	if m != nil {
		return m.Type
	}
	return PackageType_PT_UNKNOWN
}

func (m *Input) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Input) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 下行数据
type Output struct {
	Type                 PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PackageType" json:"type,omitempty"`
	RequestId            int64       `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Code                 int32       `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Message              string      `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []byte      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{11}
}

func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetType() PackageType {
	if m != nil {
		return m.Type
	}
	return PackageType_PT_UNKNOWN
}

func (m *Output) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Output) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Output) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Output) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 设备登录,package_type:1
type SignInInput struct {
	DeviceId             int64    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInInput) Reset()         { *m = SignInInput{} }
func (m *SignInInput) String() string { return proto.CompactTextString(m) }
func (*SignInInput) ProtoMessage()    {}
func (*SignInInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{12}
}

func (m *SignInInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInInput.Unmarshal(m, b)
}
func (m *SignInInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInInput.Marshal(b, m, deterministic)
}
func (m *SignInInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInInput.Merge(m, src)
}
func (m *SignInInput) XXX_Size() int {
	return xxx_messageInfo_SignInInput.Size(m)
}
func (m *SignInInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInInput.DiscardUnknown(m)
}

var xxx_messageInfo_SignInInput proto.InternalMessageInfo

func (m *SignInInput) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *SignInInput) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SignInInput) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 消息同步请求,package_type:2
type SyncInput struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncInput) Reset()         { *m = SyncInput{} }
func (m *SyncInput) String() string { return proto.CompactTextString(m) }
func (*SyncInput) ProtoMessage()    {}
func (*SyncInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{13}
}

func (m *SyncInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncInput.Unmarshal(m, b)
}
func (m *SyncInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncInput.Marshal(b, m, deterministic)
}
func (m *SyncInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncInput.Merge(m, src)
}
func (m *SyncInput) XXX_Size() int {
	return xxx_messageInfo_SyncInput.Size(m)
}
func (m *SyncInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncInput.DiscardUnknown(m)
}

var xxx_messageInfo_SyncInput proto.InternalMessageInfo

func (m *SyncInput) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

// 消息同步响应,package_type:2
type SyncOutput struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	HasMore              bool       `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SyncOutput) Reset()         { *m = SyncOutput{} }
func (m *SyncOutput) String() string { return proto.CompactTextString(m) }
func (*SyncOutput) ProtoMessage()    {}
func (*SyncOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{14}
}

func (m *SyncOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncOutput.Unmarshal(m, b)
}
func (m *SyncOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncOutput.Marshal(b, m, deterministic)
}
func (m *SyncOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncOutput.Merge(m, src)
}
func (m *SyncOutput) XXX_Size() int {
	return xxx_messageInfo_SyncOutput.Size(m)
}
func (m *SyncOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncOutput.DiscardUnknown(m)
}

var xxx_messageInfo_SyncOutput proto.InternalMessageInfo

func (m *SyncOutput) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *SyncOutput) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

// 消息投递,package_type:4
type MessageSend struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageSend) Reset()         { *m = MessageSend{} }
func (m *MessageSend) String() string { return proto.CompactTextString(m) }
func (*MessageSend) ProtoMessage()    {}
func (*MessageSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{15}
}

func (m *MessageSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSend.Unmarshal(m, b)
}
func (m *MessageSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSend.Marshal(b, m, deterministic)
}
func (m *MessageSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSend.Merge(m, src)
}
func (m *MessageSend) XXX_Size() int {
	return xxx_messageInfo_MessageSend.Size(m)
}
func (m *MessageSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSend.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSend proto.InternalMessageInfo

func (m *MessageSend) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

// 投递消息回执,package_type:4
type MessageACK struct {
	DeviceAck            int64    `protobuf:"varint,2,opt,name=device_ack,json=deviceAck,proto3" json:"device_ack,omitempty"`
	ReceiveTime          int64    `protobuf:"varint,3,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageACK) Reset()         { *m = MessageACK{} }
func (m *MessageACK) String() string { return proto.CompactTextString(m) }
func (*MessageACK) ProtoMessage()    {}
func (*MessageACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_17200129a46574b9, []int{16}
}

func (m *MessageACK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageACK.Unmarshal(m, b)
}
func (m *MessageACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageACK.Marshal(b, m, deterministic)
}
func (m *MessageACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageACK.Merge(m, src)
}
func (m *MessageACK) XXX_Size() int {
	return xxx_messageInfo_MessageACK.Size(m)
}
func (m *MessageACK) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageACK.DiscardUnknown(m)
}

var xxx_messageInfo_MessageACK proto.InternalMessageInfo

func (m *MessageACK) GetDeviceAck() int64 {
	if m != nil {
		return m.DeviceAck
	}
	return 0
}

func (m *MessageACK) GetReceiveTime() int64 {
	if m != nil {
		return m.ReceiveTime
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.PackageType", PackageType_name, PackageType_value)
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("pb.ReceiverType", ReceiverType_name, ReceiverType_value)
	proto.RegisterEnum("pb.SenderType", SenderType_name, SenderType_value)
	proto.RegisterEnum("pb.MessageStatus", MessageStatus_name, MessageStatus_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Sender)(nil), "pb.Sender")
	proto.RegisterType((*Text)(nil), "pb.Text")
	proto.RegisterType((*Face)(nil), "pb.Face")
	proto.RegisterType((*Voice)(nil), "pb.Voice")
	proto.RegisterType((*Image)(nil), "pb.Image")
	proto.RegisterType((*File)(nil), "pb.File")
	proto.RegisterType((*Location)(nil), "pb.Location")
	proto.RegisterType((*Command)(nil), "pb.Command")
	proto.RegisterType((*Custom)(nil), "pb.Custom")
	proto.RegisterType((*Input)(nil), "pb.Input")
	proto.RegisterType((*Output)(nil), "pb.Output")
	proto.RegisterType((*SignInInput)(nil), "pb.SignInInput")
	proto.RegisterType((*SyncInput)(nil), "pb.SyncInput")
	proto.RegisterType((*SyncOutput)(nil), "pb.SyncOutput")
	proto.RegisterType((*MessageSend)(nil), "pb.MessageSend")
	proto.RegisterType((*MessageACK)(nil), "pb.MessageACK")
}

func init() { proto.RegisterFile("conn.ext.proto", fileDescriptor_17200129a46574b9) }

var fileDescriptor_17200129a46574b9 = []byte{
	// 1086 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x8e, 0xe3, 0x44,
	0x13, 0xfe, 0x6d, 0xe7, 0xe4, 0x72, 0x92, 0xf1, 0xdf, 0xac, 0x20, 0xec, 0x01, 0x82, 0x57, 0x68,
	0xc3, 0x5c, 0x0c, 0x62, 0x80, 0x1b, 0x24, 0x2e, 0xb2, 0xc1, 0x3b, 0x58, 0x1b, 0x27, 0x51, 0xdb,
	0x59, 0x76, 0x24, 0x24, 0xcb, 0x63, 0x37, 0x13, 0x6b, 0x62, 0x3b, 0x6b, 0xb7, 0x87, 0x19, 0xc4,
	0x33, 0x20, 0xde, 0x82, 0x97, 0xe2, 0x61, 0x50, 0x1f, 0xec, 0x78, 0x96, 0x5b, 0xee, 0xba, 0xaa,
	0xba, 0xab, 0xbe, 0xef, 0x73, 0x75, 0xb5, 0x61, 0x1c, 0xe5, 0x59, 0x76, 0x46, 0xee, 0xe8, 0xd9,
	0xa1, 0xc8, 0x69, 0x8e, 0xd4, 0xc3, 0x95, 0xf5, 0xb7, 0x0a, 0x7d, 0x97, 0x94, 0x65, 0x78, 0x4d,
	0x90, 0x05, 0xbd, 0x92, 0x64, 0x31, 0x29, 0x26, 0xca, 0x54, 0x99, 0x19, 0xe7, 0x70, 0x76, 0xb8,
	0x3a, 0xf3, 0xb8, 0x07, 0xcb, 0x08, 0xfa, 0x16, 0x46, 0x05, 0x89, 0x48, 0x72, 0x4b, 0x8a, 0x80,
	0xde, 0x1f, 0xc8, 0x44, 0x9d, 0x2a, 0xb3, 0xf1, 0xb9, 0xc9, 0xb6, 0x62, 0x19, 0xf0, 0xef, 0x0f,
	0x04, 0x0f, 0x8b, 0x96, 0x85, 0x3e, 0x05, 0xa3, 0x39, 0x96, 0xc4, 0x13, 0x6d, 0xaa, 0xcc, 0x34,
	0x0c, 0xb5, 0xcb, 0x89, 0xd1, 0x27, 0x60, 0xd0, 0x3c, 0xa8, 0x4a, 0x1e, 0x2f, 0x27, 0x9d, 0xa9,
	0x36, 0xd3, 0xb0, 0x4e, 0xf3, 0x6d, 0xc9, 0xc2, 0x25, 0x3a, 0x87, 0x61, 0x2a, 0x60, 0x8a, 0xb2,
	0x5d, 0x5e, 0xf6, 0x84, 0x95, 0x95, 0xf0, 0x79, 0x55, 0x23, 0x3d, 0x1a, 0xe8, 0x05, 0x9c, 0xd4,
	0x67, 0xa2, 0x3c, 0xa3, 0x24, 0xa3, 0x93, 0xde, 0x54, 0x99, 0x0d, 0xf1, 0x58, 0xba, 0x17, 0xc2,
	0x8b, 0x4c, 0xd0, 0x4a, 0xf2, 0x6e, 0xd2, 0xe7, 0xa8, 0xd8, 0x12, 0x3d, 0x01, 0x9d, 0x11, 0x0e,
	0x68, 0x92, 0x92, 0xc9, 0x80, 0xfb, 0x07, 0xcc, 0xe1, 0x27, 0x29, 0x41, 0x5f, 0x40, 0xaf, 0xa4,
	0x21, 0xad, 0xca, 0x89, 0xce, 0x51, 0xfc, 0xbf, 0x85, 0xc2, 0xe3, 0x01, 0x2c, 0x37, 0x58, 0x7f,
	0x29, 0xd0, 0x13, 0x0a, 0xa2, 0x2f, 0xc1, 0x10, 0x1a, 0x0a, 0x02, 0x0a, 0x3f, 0x3a, 0x3e, 0x4a,
	0xcc, 0xf1, 0x43, 0xd9, 0xac, 0x6b, 0x0c, 0x42, 0x31, 0xf5, 0x88, 0x81, 0xeb, 0xf5, 0x0c, 0x20,
	0xbc, 0x0d, 0x69, 0x58, 0x04, 0x55, 0xb1, 0xe7, 0x7a, 0xea, 0x58, 0x17, 0x9e, 0x6d, 0xb1, 0x47,
	0x8f, 0x61, 0x90, 0x25, 0xd1, 0x4d, 0x16, 0xa6, 0x64, 0xd2, 0xe1, 0xc1, 0xc6, 0x46, 0x8f, 0xa0,
	0x4b, 0xee, 0x68, 0x11, 0x72, 0x0d, 0x75, 0x2c, 0x0c, 0xeb, 0x31, 0x74, 0x7c, 0x72, 0x47, 0x11,
	0x82, 0x0e, 0x25, 0x77, 0x94, 0xe3, 0xd3, 0x31, 0x5f, 0x5b, 0xdf, 0x41, 0xe7, 0x55, 0x18, 0x11,
	0xf4, 0x11, 0xf4, 0x7f, 0x09, 0x23, 0xc2, 0xf0, 0x28, 0x1c, 0x4f, 0x8f, 0x99, 0x4e, 0x8c, 0x3e,
	0x86, 0x01, 0x0f, 0x30, 0x2c, 0x2a, 0x3f, 0xc8, 0x37, 0x6e, 0x8b, 0xbd, 0x75, 0x09, 0xdd, 0x37,
	0x79, 0x12, 0x11, 0x34, 0x06, 0x55, 0x9e, 0xd3, 0xb1, 0x9a, 0xc4, 0xac, 0x50, 0x99, 0xfc, 0x26,
	0x1a, 0xa8, 0x8b, 0xf9, 0x9a, 0xc1, 0x8e, 0xab, 0x22, 0xa4, 0x49, 0x9e, 0x71, 0x4e, 0x5d, 0xdc,
	0xd8, 0xec, 0x23, 0xb1, 0xf4, 0x82, 0x0d, 0x5b, 0x5a, 0xbf, 0x43, 0xd7, 0x49, 0x59, 0xe3, 0xbe,
	0x9f, 0xfa, 0x11, 0x74, 0x7f, 0x4d, 0x62, 0xba, 0x93, 0xb9, 0x85, 0x81, 0x3e, 0x84, 0xde, 0x8e,
	0x24, 0xd7, 0x3b, 0x2a, 0x53, 0x4b, 0xeb, 0xdf, 0x89, 0xd1, 0x73, 0x18, 0xd1, 0x5d, 0x95, 0x5e,
	0x65, 0x61, 0xb2, 0xe7, 0x9c, 0x84, 0x52, 0xc3, 0xc6, 0xc9, 0x88, 0x6d, 0xa0, 0xf3, 0x2a, 0xd9,
	0xd7, 0xc5, 0x87, 0x5c, 0x0f, 0xc9, 0x8b, 0xcb, 0x3e, 0x12, 0x02, 0x72, 0xc9, 0x6b, 0xae, 0x63,
	0xbe, 0x4b, 0x70, 0x95, 0x65, 0x4f, 0x8e, 0x7c, 0xde, 0xc2, 0x60, 0x99, 0x47, 0x82, 0x2d, 0x82,
	0x4e, 0x4c, 0xca, 0xa8, 0xfe, 0x0c, 0x6c, 0xcd, 0xd4, 0xd9, 0x87, 0x34, 0xa1, 0x55, 0x2c, 0x54,
	0x53, 0x70, 0x63, 0xa3, 0xa7, 0xa0, 0xef, 0xf3, 0xec, 0x5a, 0x04, 0x35, 0x1e, 0x3c, 0x3a, 0xac,
	0xaf, 0xa0, 0xbf, 0xc8, 0xd3, 0x34, 0xcc, 0x38, 0xbc, 0x28, 0x8f, 0x45, 0xff, 0x75, 0x31, 0x5f,
	0xf3, 0x62, 0x21, 0x0d, 0x79, 0xd2, 0x21, 0xe6, 0x6b, 0xeb, 0x29, 0xf4, 0x16, 0x55, 0x49, 0xf3,
	0xb4, 0x89, 0xd6, 0x50, 0x58, 0x34, 0x80, 0xae, 0x93, 0x1d, 0x2a, 0x8a, 0x9e, 0x43, 0xa7, 0xd5,
	0xce, 0xfc, 0x3e, 0x6e, 0xc2, 0xe8, 0xa6, 0xbe, 0x8f, 0x3c, 0xc8, 0x9a, 0xb5, 0x20, 0xef, 0x2a,
	0x52, 0xd2, 0x63, 0x2b, 0xeb, 0xd2, 0xe3, 0xc4, 0x4d, 0x01, 0xad, 0x55, 0xfe, 0x0f, 0x05, 0x7a,
	0xeb, 0x8a, 0xfe, 0x87, 0x25, 0x38, 0x6b, 0xad, 0xc5, 0x7a, 0x02, 0x7d, 0x39, 0x07, 0xe4, 0xb7,
	0xaf, 0xcd, 0x06, 0x50, 0xb7, 0x05, 0xe8, 0x12, 0x0c, 0x2f, 0xb9, 0xce, 0x9c, 0x4c, 0xf0, 0x7e,
	0x02, 0x7a, 0x4c, 0x6e, 0x93, 0xf6, 0x65, 0x18, 0x08, 0x87, 0x13, 0xb3, 0x7b, 0x22, 0x27, 0x99,
	0x44, 0xd2, 0xab, 0xf8, 0x18, 0x63, 0x8d, 0x49, 0xf3, 0x1b, 0x92, 0xc9, 0x0b, 0x2b, 0x0c, 0xeb,
	0x19, 0xe8, 0xde, 0x7d, 0x16, 0x89, 0xc4, 0x72, 0x16, 0x29, 0xcd, 0x2c, 0xb2, 0x36, 0x00, 0x2c,
	0x2c, 0xd5, 0x78, 0x01, 0x03, 0x09, 0xb3, 0x9c, 0x28, 0x53, 0x6d, 0x66, 0x9c, 0x1b, 0xad, 0xf1,
	0x83, 0x9b, 0x20, 0xbb, 0x93, 0xbb, 0xb0, 0x0c, 0xd2, 0xbc, 0x10, 0xdd, 0x32, 0xc0, 0xfd, 0x5d,
	0x58, 0xba, 0x79, 0x41, 0xac, 0x6f, 0xc0, 0xa8, 0xc7, 0x15, 0xc9, 0x62, 0xf4, 0xf9, 0x51, 0x08,
	0x31, 0xf8, 0x1f, 0x64, 0xac, 0x63, 0xd6, 0x0a, 0x40, 0xfa, 0xe6, 0x8b, 0xd7, 0x4c, 0x70, 0x29,
	0x40, 0x18, 0xdd, 0xd4, 0x82, 0x0b, 0xcf, 0x3c, 0xba, 0x41, 0x9f, 0x41, 0xfd, 0x00, 0x88, 0x19,
	0x2a, 0x26, 0x7e, 0xfd, 0x08, 0xb0, 0x31, 0x7a, 0xfa, 0x33, 0x18, 0xad, 0xef, 0x88, 0xc6, 0x00,
	0x1b, 0x3f, 0xd8, 0xae, 0x5e, 0xaf, 0xd6, 0x3f, 0xad, 0xcc, 0xff, 0x49, 0xdb, 0x73, 0x2e, 0x56,
	0x81, 0xb3, 0x32, 0x15, 0x64, 0x40, 0x9f, 0xd9, 0x97, 0xab, 0x85, 0xa9, 0x22, 0x13, 0x86, 0x1b,
	0x3f, 0xf8, 0xd1, 0x9e, 0x63, 0xff, 0xa5, 0x3d, 0xf7, 0x4d, 0x4d, 0x6e, 0x77, 0x6d, 0xcf, 0x9b,
	0x5f, 0xd8, 0x66, 0xe7, 0xf4, 0x4f, 0xa5, 0x21, 0x59, 0xa7, 0x77, 0xdb, 0xe9, 0x0d, 0xe8, 0xbb,
	0x7e, 0xe0, 0xdb, 0x6f, 0x7d, 0x91, 0xdb, 0xf5, 0x83, 0x57, 0xf3, 0x85, 0x6d, 0xaa, 0x68, 0x08,
	0x03, 0xd7, 0x0f, 0xde, 0xac, 0x9d, 0x85, 0x6d, 0x6a, 0xd2, 0x72, 0x5c, 0x9e, 0xb5, 0xde, 0xe8,
	0x2c, 0x6d, 0xb3, 0x8b, 0x4e, 0xc0, 0x70, 0xfd, 0x60, 0xb9, 0x5e, 0xcc, 0x7d, 0x67, 0xbd, 0x32,
	0x7b, 0xb2, 0xc6, 0x62, 0xed, 0xba, 0xf3, 0xd5, 0x0f, 0x66, 0x1f, 0x8d, 0x40, 0x67, 0xf6, 0xd6,
	0xf3, 0xd7, 0xae, 0x39, 0x38, 0xf5, 0x61, 0xd8, 0x7e, 0x22, 0xd9, 0x76, 0xfc, 0x1e, 0x24, 0x66,
	0x7b, 0x36, 0x36, 0x15, 0xf4, 0x01, 0x9c, 0x60, 0x3f, 0x58, 0xad, 0xb1, 0x3b, 0x5f, 0x06, 0x17,
	0x78, 0xbd, 0xdd, 0x98, 0x2a, 0x42, 0x30, 0xc6, 0x7e, 0xb0, 0x9c, 0xe3, 0x0b, 0x5b, 0xfa, 0xb4,
	0x53, 0x07, 0xe0, 0xf8, 0x80, 0xb0, 0x9c, 0x5e, 0x3b, 0xe7, 0x08, 0x74, 0x8f, 0xa9, 0xe6, 0xf9,
	0xb6, 0x2b, 0x88, 0x7a, 0xb2, 0x84, 0xca, 0xf0, 0x7b, 0x7e, 0xf0, 0x72, 0xeb, 0x39, 0x2b, 0xdb,
	0xf3, 0x4c, 0xed, 0xf4, 0x7b, 0x18, 0x3d, 0x78, 0xc6, 0x38, 0x21, 0xef, 0x61, 0x36, 0xd7, 0x93,
	0xa0, 0x4c, 0x45, 0x9a, 0xd8, 0x5e, 0xcc, 0x97, 0x4b, 0x53, 0xbd, 0xea, 0xf1, 0xdf, 0x8a, 0xaf,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xb3, 0x77, 0x30, 0x68, 0x08, 0x00, 0x00,
}
