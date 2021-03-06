package face

type IRoomManager interface {
	CreateRoom(roleId int64) IRoom
	GetRoomByRoomId(roomId int32) IRoom
	GetAllRoom()map[int32]IRoom
	CheckAllRoomMemberLoadFinished(roomId int32) bool
}
