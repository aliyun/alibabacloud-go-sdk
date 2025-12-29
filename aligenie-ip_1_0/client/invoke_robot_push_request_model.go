// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeRobotPushRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *InvokeRobotPushRequest
	GetHotelId() *string
	SetPushType(v string) *InvokeRobotPushRequest
	GetPushType() *string
	SetRoomName(v string) *InvokeRobotPushRequest
	GetRoomName() *string
	SetRoomNo(v string) *InvokeRobotPushRequest
	GetRoomNo() *string
}

type InvokeRobotPushRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GET
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s InvokeRobotPushRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeRobotPushRequest) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *InvokeRobotPushRequest) GetPushType() *string {
	return s.PushType
}

func (s *InvokeRobotPushRequest) GetRoomName() *string {
	return s.RoomName
}

func (s *InvokeRobotPushRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *InvokeRobotPushRequest) SetHotelId(v string) *InvokeRobotPushRequest {
	s.HotelId = &v
	return s
}

func (s *InvokeRobotPushRequest) SetPushType(v string) *InvokeRobotPushRequest {
	s.PushType = &v
	return s
}

func (s *InvokeRobotPushRequest) SetRoomName(v string) *InvokeRobotPushRequest {
	s.RoomName = &v
	return s
}

func (s *InvokeRobotPushRequest) SetRoomNo(v string) *InvokeRobotPushRequest {
	s.RoomNo = &v
	return s
}

func (s *InvokeRobotPushRequest) Validate() error {
	return dara.Validate(s)
}
