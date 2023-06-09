package server

import (
	"fmt"
	"main/src/zdemo/Server/siface"
	"main/src/zdemo/Server/utils"
	"main/src/zinx/ziface"
	"strings"
)

// Boradcast
type BroadcastRouter struct {
	Router
}

func NewBroadcastRouter() (broadcast_rt *BroadcastRouter) {
	broadcast_rt = &BroadcastRouter{}
	return
}

func (this *BroadcastRouter) Handle(request ziface.IRequest) {
	iuser, err := request.GetConn().GetProperty("user")
	if err != nil {
		panic(err.Error())
	}
	user := iuser.(siface.IUser)
	room := user.GetRoom()

	msg := fmt.Sprintf("[%s]:%s", user.GetName(), request.GetData())
	room.BroadCastMsg([]byte(msg))
}

// Private chat
type PrivateChatRouter struct {
	Router
}

func NewPrivateChatRouter() (private_chat_rt *PrivateChatRouter) {
	private_chat_rt = &PrivateChatRouter{}
	return
}

func (this *PrivateChatRouter) Handle(request ziface.IRequest) {
	iuser, err := request.GetConn().GetProperty("user")
	if err != nil {
		panic(err.Error())
	}
	user := iuser.(siface.IUser)
	room := user.GetRoom()

	data := string(request.GetData())
	strs := strings.SplitN(data, ":", 2)
	name := strs[0]
	msg := strs[1]

	user2chat, err := room.GetUser(name)
	if err != nil {
		msg = fmt.Sprintf("User %s is not exist", name)
		request.GetConn().SendMsg(utils.NCmdResponse, []byte(msg))
		return
	}

	msg = fmt.Sprintf("[%s]:%s", user.GetName(), msg)
	user.AddMsg([]byte(msg))
	user2chat.AddMsg([]byte(msg))
}
