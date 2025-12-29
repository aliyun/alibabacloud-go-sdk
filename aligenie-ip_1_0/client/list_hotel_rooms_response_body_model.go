// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelRoomsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelRoomsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelRoomsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelRoomsResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelRoomsResponseBodyResult) *ListHotelRoomsResponseBody
	GetResult() []*ListHotelRoomsResponseBodyResult
}

type ListHotelRoomsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelRoomsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelRoomsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelRoomsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelRoomsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelRoomsResponseBody) GetResult() []*ListHotelRoomsResponseBodyResult {
	return s.Result
}

func (s *ListHotelRoomsResponseBody) SetCode(v int32) *ListHotelRoomsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetMessage(v string) *ListHotelRoomsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetRequestId(v string) *ListHotelRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelRoomsResponseBody) SetResult(v []*ListHotelRoomsResponseBodyResult) *ListHotelRoomsResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelRoomsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelRoomsResponseBodyResult struct {
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 102
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
}

func (s ListHotelRoomsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelRoomsResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListHotelRoomsResponseBodyResult) SetHotelId(v string) *ListHotelRoomsResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListHotelRoomsResponseBodyResult) SetRoomNo(v string) *ListHotelRoomsResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListHotelRoomsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
