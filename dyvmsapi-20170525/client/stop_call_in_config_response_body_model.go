// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCallInConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopCallInConfigResponseBody
	GetCode() *string
	SetData(v bool) *StopCallInConfigResponseBody
	GetData() *bool
	SetMessage(v string) *StopCallInConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopCallInConfigResponseBody
	GetRequestId() *string
}

type StopCallInConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the inbound call was stopped. Valid values:
	//
	// 	- true: The inbound call was stopped.
	//
	// 	- false: The inbound call failed to be stopped.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// a78278ff-26bb-48ec-805c-26a0f4c102***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCallInConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCallInConfigResponseBody) GoString() string {
	return s.String()
}

func (s *StopCallInConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopCallInConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopCallInConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopCallInConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCallInConfigResponseBody) SetCode(v string) *StopCallInConfigResponseBody {
	s.Code = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetData(v bool) *StopCallInConfigResponseBody {
	s.Data = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetMessage(v string) *StopCallInConfigResponseBody {
	s.Message = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetRequestId(v string) *StopCallInConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCallInConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
