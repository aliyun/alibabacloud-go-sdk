// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SingleCallByVideoResponseBody
	GetCallId() *string
	SetCode(v string) *SingleCallByVideoResponseBody
	GetCode() *string
	SetMessage(v string) *SingleCallByVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SingleCallByVideoResponseBody
	GetRequestId() *string
}

type SingleCallByVideoResponseBody struct {
	// example:
	//
	// 116012354148^10281378****
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// a78278ff-26bb-48ec-805c-26a0f4c102***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoResponseBody) GetCallId() *string {
	return s.CallId
}

func (s *SingleCallByVideoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SingleCallByVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SingleCallByVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SingleCallByVideoResponseBody) SetCallId(v string) *SingleCallByVideoResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetCode(v string) *SingleCallByVideoResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetMessage(v string) *SingleCallByVideoResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetRequestId(v string) *SingleCallByVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleCallByVideoResponseBody) Validate() error {
	return dara.Validate(s)
}
