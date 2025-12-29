// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushHotelMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PushHotelMessageResponseBody
	GetCode() *int32
	SetMessage(v string) *PushHotelMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushHotelMessageResponseBody
	GetRequestId() *string
	SetResult(v bool) *PushHotelMessageResponseBody
	GetResult() *bool
}

type PushHotelMessageResponseBody struct {
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

func (s PushHotelMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushHotelMessageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PushHotelMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushHotelMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushHotelMessageResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PushHotelMessageResponseBody) SetCode(v int32) *PushHotelMessageResponseBody {
	s.Code = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetMessage(v string) *PushHotelMessageResponseBody {
	s.Message = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetRequestId(v string) *PushHotelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushHotelMessageResponseBody) SetResult(v bool) *PushHotelMessageResponseBody {
	s.Result = &v
	return s
}

func (s *PushHotelMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
