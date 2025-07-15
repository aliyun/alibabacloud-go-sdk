// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReleaseChatResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ReleaseChatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ReleaseChatResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseChatResponseBody
	GetRequestId() *string
}

type ReleaseChatResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B06B3244-1B44-481B-90C4-F2F92E59D6B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseChatResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReleaseChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReleaseChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseChatResponseBody) SetCode(v string) *ReleaseChatResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseChatResponseBody) SetHttpStatusCode(v int32) *ReleaseChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseChatResponseBody) SetMessage(v string) *ReleaseChatResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseChatResponseBody) SetRequestId(v string) *ReleaseChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseChatResponseBody) Validate() error {
	return dara.Validate(s)
}
