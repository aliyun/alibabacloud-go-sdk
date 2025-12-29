// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelAlarmRequest
	GetHotelId() *string
	SetRooms(v []*string) *ListHotelAlarmRequest
	GetRooms() []*string
}

type ListHotelAlarmRequest struct {
	// example:
	//
	// a7a3***013
	HotelId *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Rooms   []*string `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
}

func (s ListHotelAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmRequest) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelAlarmRequest) GetRooms() []*string {
	return s.Rooms
}

func (s *ListHotelAlarmRequest) SetHotelId(v string) *ListHotelAlarmRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelAlarmRequest) SetRooms(v []*string) *ListHotelAlarmRequest {
	s.Rooms = v
	return s
}

func (s *ListHotelAlarmRequest) Validate() error {
	return dara.Validate(s)
}
