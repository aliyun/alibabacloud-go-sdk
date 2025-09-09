// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRTCLiveRoomsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRTCLiveRoomsResponseBody
	GetRequestId() *string
	SetRooms(v []*string) *ListRTCLiveRoomsResponseBody
	GetRooms() []*string
	SetTotal(v int32) *ListRTCLiveRoomsResponseBody
	GetTotal() *int32
}

type ListRTCLiveRoomsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b96****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rooms     []*string `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListRTCLiveRoomsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRTCLiveRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRTCLiveRoomsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRTCLiveRoomsResponseBody) GetRooms() []*string {
	return s.Rooms
}

func (s *ListRTCLiveRoomsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListRTCLiveRoomsResponseBody) SetRequestId(v string) *ListRTCLiveRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRTCLiveRoomsResponseBody) SetRooms(v []*string) *ListRTCLiveRoomsResponseBody {
	s.Rooms = v
	return s
}

func (s *ListRTCLiveRoomsResponseBody) SetTotal(v int32) *ListRTCLiveRoomsResponseBody {
	s.Total = &v
	return s
}

func (s *ListRTCLiveRoomsResponseBody) Validate() error {
	return dara.Validate(s)
}
