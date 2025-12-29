// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoomCheckOutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RoomCheckOutResponseBody
	GetCode() *int32
	SetMessage(v string) *RoomCheckOutResponseBody
	GetMessage() *string
	SetRequestId(v string) *RoomCheckOutResponseBody
	GetRequestId() *string
	SetResult(v bool) *RoomCheckOutResponseBody
	GetResult() *bool
}

type RoomCheckOutResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dsvrevd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RoomCheckOutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RoomCheckOutResponseBody) GoString() string {
	return s.String()
}

func (s *RoomCheckOutResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RoomCheckOutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RoomCheckOutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RoomCheckOutResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RoomCheckOutResponseBody) SetCode(v int32) *RoomCheckOutResponseBody {
	s.Code = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetMessage(v string) *RoomCheckOutResponseBody {
	s.Message = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetRequestId(v string) *RoomCheckOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoomCheckOutResponseBody) SetResult(v bool) *RoomCheckOutResponseBody {
	s.Result = &v
	return s
}

func (s *RoomCheckOutResponseBody) Validate() error {
	return dara.Validate(s)
}
