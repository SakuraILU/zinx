package siface

type IRoom interface {
	StartRoom()
	StopRoom()

	GetName() string

	GetUser(name string) (IUser, error)
	AddUser(IUser) error
	RemoveUser(IUser)
	GetAllUserMsg() string
	ClearAll()

	GetUserNum() uint32
	GetCapacity() uint32

	BroadCastMsg([]byte)
}
