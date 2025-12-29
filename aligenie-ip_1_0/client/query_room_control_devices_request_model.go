// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomControlDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QueryRoomControlDevicesRequest
	GetHotelId() *string
	SetRoomNo(v string) *QueryRoomControlDevicesRequest
	GetRoomNo() *string
}

type QueryRoomControlDevicesRequest struct {
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
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s QueryRoomControlDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomControlDevicesRequest) GoString() string {
	return s.String()
}

func (s *QueryRoomControlDevicesRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryRoomControlDevicesRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryRoomControlDevicesRequest) SetHotelId(v string) *QueryRoomControlDevicesRequest {
	s.HotelId = &v
	return s
}

func (s *QueryRoomControlDevicesRequest) SetRoomNo(v string) *QueryRoomControlDevicesRequest {
	s.RoomNo = &v
	return s
}

func (s *QueryRoomControlDevicesRequest) Validate() error {
	return dara.Validate(s)
}
