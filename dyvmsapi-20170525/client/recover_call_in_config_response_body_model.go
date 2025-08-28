// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverCallInConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RecoverCallInConfigResponseBody
	GetCode() *string
	SetData(v bool) *RecoverCallInConfigResponseBody
	GetData() *bool
	SetMessage(v string) *RecoverCallInConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecoverCallInConfigResponseBody
	GetRequestId() *string
}

type RecoverCallInConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the inbound call was resumed. Valid values:
	//
	// 	- true: The inbound call was resumed.
	//
	// 	- false: The inbound call failed to be resumed.
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

func (s RecoverCallInConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverCallInConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *RecoverCallInConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *RecoverCallInConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecoverCallInConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverCallInConfigResponseBody) SetCode(v string) *RecoverCallInConfigResponseBody {
	s.Code = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetData(v bool) *RecoverCallInConfigResponseBody {
	s.Data = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetMessage(v string) *RecoverCallInConfigResponseBody {
	s.Message = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetRequestId(v string) *RecoverCallInConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
