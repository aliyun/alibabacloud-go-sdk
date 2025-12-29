// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHotelRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchDeleteHotelRoomResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchDeleteHotelRoomResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteHotelRoomResponseBody
	GetRequestId() *string
	SetResult(v bool) *BatchDeleteHotelRoomResponseBody
	GetResult() *bool
}

type BatchDeleteHotelRoomResponseBody struct {
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

func (s BatchDeleteHotelRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHotelRoomResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchDeleteHotelRoomResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteHotelRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteHotelRoomResponseBody) GetResult() *bool {
	return s.Result
}

func (s *BatchDeleteHotelRoomResponseBody) SetCode(v int32) *BatchDeleteHotelRoomResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetMessage(v string) *BatchDeleteHotelRoomResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetRequestId(v string) *BatchDeleteHotelRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) SetResult(v bool) *BatchDeleteHotelRoomResponseBody {
	s.Result = &v
	return s
}

func (s *BatchDeleteHotelRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
