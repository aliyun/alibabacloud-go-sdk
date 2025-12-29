// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddHotelRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *BatchAddHotelRoomRequest
	GetHotelId() *string
	SetRoomNoList(v []*string) *BatchAddHotelRoomRequest
	GetRoomNoList() []*string
}

type BatchAddHotelRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoList []*string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty" type:"Repeated"`
}

func (s BatchAddHotelRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddHotelRoomRequest) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *BatchAddHotelRoomRequest) GetRoomNoList() []*string {
	return s.RoomNoList
}

func (s *BatchAddHotelRoomRequest) SetHotelId(v string) *BatchAddHotelRoomRequest {
	s.HotelId = &v
	return s
}

func (s *BatchAddHotelRoomRequest) SetRoomNoList(v []*string) *BatchAddHotelRoomRequest {
	s.RoomNoList = v
	return s
}

func (s *BatchAddHotelRoomRequest) Validate() error {
	return dara.Validate(s)
}
