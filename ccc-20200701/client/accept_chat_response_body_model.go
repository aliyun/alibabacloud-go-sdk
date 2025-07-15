// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AcceptChatResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AcceptChatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AcceptChatResponseBody
	GetMessage() *string
	SetRequestId(v string) *AcceptChatResponseBody
	GetRequestId() *string
}

type AcceptChatResponseBody struct {
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
	// 2263B273-AC1B-44EB-BA98-87F2322C6780
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptChatResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *AcceptChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AcceptChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AcceptChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptChatResponseBody) SetCode(v string) *AcceptChatResponseBody {
	s.Code = &v
	return s
}

func (s *AcceptChatResponseBody) SetHttpStatusCode(v int32) *AcceptChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AcceptChatResponseBody) SetMessage(v string) *AcceptChatResponseBody {
	s.Message = &v
	return s
}

func (s *AcceptChatResponseBody) SetRequestId(v string) *AcceptChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptChatResponseBody) Validate() error {
	return dara.Validate(s)
}
