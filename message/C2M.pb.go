// Code generated by protoc-gen-go. DO NOT EDIT.
// source: C2M.proto

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Room struct {
	GameType             int32    `protobuf:"varint,1,opt,name=gameType,proto3" json:"gameType,omitempty"`
	RoomName             string   `protobuf:"bytes,2,opt,name=roomName,proto3" json:"roomName,omitempty"`
	MaxPlayerNum         int32    `protobuf:"varint,3,opt,name=maxPlayerNum,proto3" json:"maxPlayerNum,omitempty"`
	MapId                int32    `protobuf:"varint,4,opt,name=mapId,proto3" json:"mapId,omitempty"`
	RoomOwnerName        string   `protobuf:"bytes,5,opt,name=roomOwnerName,proto3" json:"roomOwnerName,omitempty"`
	RoomOwnerId          int64    `protobuf:"varint,6,opt,name=roomOwnerId,proto3" json:"roomOwnerId,omitempty"`
	RoomId               int32    `protobuf:"varint,7,opt,name=roomId,proto3" json:"roomId,omitempty"`
	CurPlayeNum          int32    `protobuf:"varint,8,opt,name=curPlayeNum,proto3" json:"curPlayeNum,omitempty"`
	RoomOwnerGroupId     int32    `protobuf:"varint,9,opt,name=roomOwnerGroupId,proto3" json:"roomOwnerGroupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{0}
}
func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (dst *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(dst, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetGameType() int32 {
	if m != nil {
		return m.GameType
	}
	return 0
}

func (m *Room) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *Room) GetMaxPlayerNum() int32 {
	if m != nil {
		return m.MaxPlayerNum
	}
	return 0
}

func (m *Room) GetMapId() int32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *Room) GetRoomOwnerName() string {
	if m != nil {
		return m.RoomOwnerName
	}
	return ""
}

func (m *Room) GetRoomOwnerId() int64 {
	if m != nil {
		return m.RoomOwnerId
	}
	return 0
}

func (m *Room) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *Room) GetCurPlayeNum() int32 {
	if m != nil {
		return m.CurPlayeNum
	}
	return 0
}

func (m *Room) GetRoomOwnerGroupId() int32 {
	if m != nil {
		return m.RoomOwnerGroupId
	}
	return 0
}

// 操作指令
type Command struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	CommandType          int32    `protobuf:"varint,2,opt,name=commandType,proto3" json:"commandType,omitempty"`
	Param                string   `protobuf:"bytes,3,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{1}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Command) GetCommandType() int32 {
	if m != nil {
		return m.CommandType
	}
	return 0
}

func (m *Command) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

// 创建房间
type C2M_CreateRoom struct {
	Room                 *Room    `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_CreateRoom) Reset()         { *m = C2M_CreateRoom{} }
func (m *C2M_CreateRoom) String() string { return proto.CompactTextString(m) }
func (*C2M_CreateRoom) ProtoMessage()    {}
func (*C2M_CreateRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{2}
}
func (m *C2M_CreateRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_CreateRoom.Unmarshal(m, b)
}
func (m *C2M_CreateRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_CreateRoom.Marshal(b, m, deterministic)
}
func (dst *C2M_CreateRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_CreateRoom.Merge(dst, src)
}
func (m *C2M_CreateRoom) XXX_Size() int {
	return xxx_messageInfo_C2M_CreateRoom.Size(m)
}
func (m *C2M_CreateRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_CreateRoom.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_CreateRoom proto.InternalMessageInfo

func (m *C2M_CreateRoom) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

// 请求刷新房间列表
type C2M_ReqRefreshRoomList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_ReqRefreshRoomList) Reset()         { *m = C2M_ReqRefreshRoomList{} }
func (m *C2M_ReqRefreshRoomList) String() string { return proto.CompactTextString(m) }
func (*C2M_ReqRefreshRoomList) ProtoMessage()    {}
func (*C2M_ReqRefreshRoomList) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{3}
}
func (m *C2M_ReqRefreshRoomList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_ReqRefreshRoomList.Unmarshal(m, b)
}
func (m *C2M_ReqRefreshRoomList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_ReqRefreshRoomList.Marshal(b, m, deterministic)
}
func (dst *C2M_ReqRefreshRoomList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_ReqRefreshRoomList.Merge(dst, src)
}
func (m *C2M_ReqRefreshRoomList) XXX_Size() int {
	return xxx_messageInfo_C2M_ReqRefreshRoomList.Size(m)
}
func (m *C2M_ReqRefreshRoomList) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_ReqRefreshRoomList.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_ReqRefreshRoomList proto.InternalMessageInfo

// 请求加入房间
type C2M_ReqJoinRoom struct {
	RoomId               int32    `protobuf:"varint,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_ReqJoinRoom) Reset()         { *m = C2M_ReqJoinRoom{} }
func (m *C2M_ReqJoinRoom) String() string { return proto.CompactTextString(m) }
func (*C2M_ReqJoinRoom) ProtoMessage()    {}
func (*C2M_ReqJoinRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{4}
}
func (m *C2M_ReqJoinRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_ReqJoinRoom.Unmarshal(m, b)
}
func (m *C2M_ReqJoinRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_ReqJoinRoom.Marshal(b, m, deterministic)
}
func (dst *C2M_ReqJoinRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_ReqJoinRoom.Merge(dst, src)
}
func (m *C2M_ReqJoinRoom) XXX_Size() int {
	return xxx_messageInfo_C2M_ReqJoinRoom.Size(m)
}
func (m *C2M_ReqJoinRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_ReqJoinRoom.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_ReqJoinRoom proto.InternalMessageInfo

func (m *C2M_ReqJoinRoom) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

// 请求开始战斗
type C2M_StartBattle struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_StartBattle) Reset()         { *m = C2M_StartBattle{} }
func (m *C2M_StartBattle) String() string { return proto.CompactTextString(m) }
func (*C2M_StartBattle) ProtoMessage()    {}
func (*C2M_StartBattle) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{5}
}
func (m *C2M_StartBattle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_StartBattle.Unmarshal(m, b)
}
func (m *C2M_StartBattle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_StartBattle.Marshal(b, m, deterministic)
}
func (dst *C2M_StartBattle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_StartBattle.Merge(dst, src)
}
func (m *C2M_StartBattle) XXX_Size() int {
	return xxx_messageInfo_C2M_StartBattle.Size(m)
}
func (m *C2M_StartBattle) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_StartBattle.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_StartBattle proto.InternalMessageInfo

// 请求准备
type C2M_ReqReady struct {
	Ready                bool     `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_ReqReady) Reset()         { *m = C2M_ReqReady{} }
func (m *C2M_ReqReady) String() string { return proto.CompactTextString(m) }
func (*C2M_ReqReady) ProtoMessage()    {}
func (*C2M_ReqReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{6}
}
func (m *C2M_ReqReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_ReqReady.Unmarshal(m, b)
}
func (m *C2M_ReqReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_ReqReady.Marshal(b, m, deterministic)
}
func (dst *C2M_ReqReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_ReqReady.Merge(dst, src)
}
func (m *C2M_ReqReady) XXX_Size() int {
	return xxx_messageInfo_C2M_ReqReady.Size(m)
}
func (m *C2M_ReqReady) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_ReqReady.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_ReqReady proto.InternalMessageInfo

func (m *C2M_ReqReady) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

// 加载完成
type C2M_LoadFinished struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_LoadFinished) Reset()         { *m = C2M_LoadFinished{} }
func (m *C2M_LoadFinished) String() string { return proto.CompactTextString(m) }
func (*C2M_LoadFinished) ProtoMessage()    {}
func (*C2M_LoadFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{7}
}
func (m *C2M_LoadFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_LoadFinished.Unmarshal(m, b)
}
func (m *C2M_LoadFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_LoadFinished.Marshal(b, m, deterministic)
}
func (dst *C2M_LoadFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_LoadFinished.Merge(dst, src)
}
func (m *C2M_LoadFinished) XXX_Size() int {
	return xxx_messageInfo_C2M_LoadFinished.Size(m)
}
func (m *C2M_LoadFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_LoadFinished.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_LoadFinished proto.InternalMessageInfo

// 玩家的操作指令
type C2M_Command struct {
	Cmd                  *Command `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_Command) Reset()         { *m = C2M_Command{} }
func (m *C2M_Command) String() string { return proto.CompactTextString(m) }
func (*C2M_Command) ProtoMessage()    {}
func (*C2M_Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2M_92d5198a279382cf, []int{8}
}
func (m *C2M_Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_Command.Unmarshal(m, b)
}
func (m *C2M_Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_Command.Marshal(b, m, deterministic)
}
func (dst *C2M_Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_Command.Merge(dst, src)
}
func (m *C2M_Command) XXX_Size() int {
	return xxx_messageInfo_C2M_Command.Size(m)
}
func (m *C2M_Command) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_Command.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_Command proto.InternalMessageInfo

func (m *C2M_Command) GetCmd() *Command {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func init() {
	proto.RegisterType((*Room)(nil), "message.Room")
	proto.RegisterType((*Command)(nil), "message.Command")
	proto.RegisterType((*C2M_CreateRoom)(nil), "message.C2M_CreateRoom")
	proto.RegisterType((*C2M_ReqRefreshRoomList)(nil), "message.C2M_ReqRefreshRoomList")
	proto.RegisterType((*C2M_ReqJoinRoom)(nil), "message.C2M_ReqJoinRoom")
	proto.RegisterType((*C2M_StartBattle)(nil), "message.C2M_StartBattle")
	proto.RegisterType((*C2M_ReqReady)(nil), "message.C2M_ReqReady")
	proto.RegisterType((*C2M_LoadFinished)(nil), "message.C2M_LoadFinished")
	proto.RegisterType((*C2M_Command)(nil), "message.C2M_Command")
}

func init() { proto.RegisterFile("C2M.proto", fileDescriptor_C2M_92d5198a279382cf) }

var fileDescriptor_C2M_92d5198a279382cf = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x95, 0xf9, 0xf6, 0x00, 0x2d, 0x5d, 0x55, 0xc8, 0xea, 0xc9, 0x5d, 0x71, 0xa0, 0x3d, 0x20,
	0x15, 0xfe, 0x41, 0x2d, 0xb5, 0x72, 0x05, 0x34, 0xda, 0xe4, 0x1a, 0x45, 0x1b, 0x76, 0x03, 0x96,
	0x58, 0xaf, 0xb3, 0x36, 0x4a, 0xf8, 0x43, 0xf9, 0x9d, 0xd1, 0x8c, 0x8d, 0x01, 0xe5, 0xe6, 0xf7,
	0xde, 0xcc, 0x9b, 0x99, 0xb7, 0x06, 0x3f, 0x9a, 0xaf, 0x66, 0x99, 0xb3, 0x85, 0x65, 0x5d, 0xa3,
	0xf3, 0x5c, 0x6e, 0x35, 0x7f, 0x6b, 0x40, 0x4b, 0x58, 0x6b, 0xd8, 0x37, 0xe8, 0x6d, 0xa5, 0xd1,
	0x77, 0xc7, 0x4c, 0x07, 0x5e, 0xe8, 0x4d, 0xdb, 0xa2, 0xc6, 0xa8, 0x39, 0x6b, 0xcd, 0x5a, 0x1a,
	0x1d, 0x34, 0x42, 0x6f, 0xea, 0x8b, 0x1a, 0x33, 0x0e, 0x03, 0x23, 0x5f, 0x6f, 0xf6, 0xf2, 0xa8,
	0xdd, 0xfa, 0x60, 0x82, 0x26, 0xf5, 0x5e, 0x71, 0xec, 0x2b, 0xb4, 0x8d, 0xcc, 0x62, 0x15, 0xb4,
	0x48, 0x2c, 0x01, 0x9b, 0xc0, 0x10, 0x5d, 0xfe, 0xbf, 0xa4, 0xda, 0x91, 0x75, 0x9b, 0xac, 0xaf,
	0x49, 0x16, 0x42, 0xbf, 0x26, 0x62, 0x15, 0x74, 0x42, 0x6f, 0xda, 0x14, 0x97, 0x14, 0x1b, 0x43,
	0x07, 0x61, 0xac, 0x82, 0x2e, 0xd9, 0x57, 0x08, 0x3b, 0x37, 0x07, 0x47, 0x5b, 0xe0, 0x62, 0x3d,
	0x12, 0x2f, 0x29, 0xf6, 0x13, 0x46, 0xb5, 0xd1, 0x5f, 0x67, 0x0f, 0xb8, 0xa2, 0x4f, 0x65, 0x1f,
	0x78, 0x7e, 0x0f, 0xdd, 0xc8, 0x1a, 0x23, 0x53, 0x85, 0x71, 0x64, 0x74, 0x5b, 0xac, 0x4e, 0x51,
	0x9d, 0x30, 0x0d, 0x2d, 0xcb, 0x28, 0xc9, 0x46, 0x35, 0xf4, 0x4c, 0x61, 0x18, 0x99, 0x74, 0xb2,
	0x4c, 0xca, 0x17, 0x25, 0xe0, 0x0b, 0xf8, 0x14, 0xcd, 0x57, 0x0f, 0x91, 0xd3, 0xb2, 0xd0, 0xf4,
	0x20, 0xdf, 0xa1, 0x85, 0x4b, 0xd0, 0x84, 0xfe, 0x7c, 0x38, 0xab, 0x5e, 0x6c, 0x86, 0xa2, 0x20,
	0x89, 0x07, 0x30, 0xc6, 0x26, 0xa1, 0x9f, 0x85, 0x7e, 0x72, 0x3a, 0xdf, 0xa1, 0xb6, 0x4c, 0xf2,
	0x82, 0xff, 0x80, 0xcf, 0x95, 0xf2, 0xcf, 0x26, 0x29, 0xf9, 0x9d, 0x63, 0xf2, 0x2e, 0x63, 0xe2,
	0x5f, 0xca, 0xd2, 0xdb, 0x42, 0xba, 0xe2, 0xb7, 0x2c, 0x8a, 0xbd, 0xe6, 0x13, 0x18, 0xd4, 0xbe,
	0x52, 0x1d, 0x71, 0x65, 0x87, 0x1f, 0xd4, 0xd9, 0x13, 0x25, 0xe0, 0x0c, 0x46, 0x58, 0xb5, 0xb4,
	0x52, 0xfd, 0x49, 0xd2, 0x24, 0xdf, 0x69, 0xc5, 0x7f, 0x41, 0x9f, 0xce, 0xa8, 0x92, 0xe2, 0xd0,
	0xdc, 0x18, 0x55, 0x9d, 0x30, 0xaa, 0x4f, 0xa8, 0x64, 0x81, 0xe2, 0x63, 0x87, 0xfe, 0xc8, 0xc5,
	0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x77, 0x05, 0x07, 0x9e, 0x02, 0x00, 0x00,
}