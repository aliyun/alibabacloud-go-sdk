// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHotelRoomRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *BatchDeleteHotelRoomRequest
	GetHotelId() *string
	SetRoomNoList(v []*string) *BatchDeleteHotelRoomRequest
	GetRoomNoList() []*string
}

type BatchDeleteHotelRoomRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	RoomNoList []*string `json:"RoomNoList,omitempty" xml:"RoomNoList,omitempty" type:"Repeated"`
}

func (s BatchDeleteHotelRoomRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHotelRoomRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *BatchDeleteHotelRoomRequest) GetRoomNoList() []*string {
	return s.RoomNoList
}

func (s *BatchDeleteHotelRoomRequest) SetHotelId(v string) *BatchDeleteHotelRoomRequest {
	s.HotelId = &v
	return s
}

func (s *BatchDeleteHotelRoomRequest) SetRoomNoList(v []*string) *BatchDeleteHotelRoomRequest {
	s.RoomNoList = v
	return s
}

func (s *BatchDeleteHotelRoomRequest) Validate() error {
	return dara.Validate(s)
}
