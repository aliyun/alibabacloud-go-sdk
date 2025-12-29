// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddHotelRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchAddHotelRoomResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchAddHotelRoomResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchAddHotelRoomResponseBody
	GetRequestId() *string
	SetResult(v bool) *BatchAddHotelRoomResponseBody
	GetResult() *bool
}

type BatchAddHotelRoomResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s BatchAddHotelRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddHotelRoomResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchAddHotelRoomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchAddHotelRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddHotelRoomResponseBody) GetResult() *bool {
	return s.Result
}

func (s *BatchAddHotelRoomResponseBody) SetCode(v int32) *BatchAddHotelRoomResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetMessage(v string) *BatchAddHotelRoomResponseBody {
	s.Message = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetRequestId(v string) *BatchAddHotelRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) SetResult(v bool) *BatchAddHotelRoomResponseBody {
	s.Result = &v
	return s
}

func (s *BatchAddHotelRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
