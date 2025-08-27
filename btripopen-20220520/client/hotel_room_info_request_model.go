// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelRoomInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoomIds(v []*int64) *HotelRoomInfoRequest
	GetRoomIds() []*int64
}

type HotelRoomInfoRequest struct {
	// This parameter is required.
	RoomIds []*int64 `json:"room_ids,omitempty" xml:"room_ids,omitempty" type:"Repeated"`
}

func (s HotelRoomInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoRequest) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoRequest) GetRoomIds() []*int64 {
	return s.RoomIds
}

func (s *HotelRoomInfoRequest) SetRoomIds(v []*int64) *HotelRoomInfoRequest {
	s.RoomIds = v
	return s
}

func (s *HotelRoomInfoRequest) Validate() error {
	return dara.Validate(s)
}
