// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAndDeletePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachAndDeletePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *DetachAndDeletePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachAndDeletePolicyResponseBody
	GetRequestId() *string
}

type DetachAndDeletePolicyResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to trace the call link.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetachAndDeletePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachAndDeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAndDeletePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachAndDeletePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachAndDeletePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachAndDeletePolicyResponseBody) SetCode(v string) *DetachAndDeletePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DetachAndDeletePolicyResponseBody) SetMessage(v string) *DetachAndDeletePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DetachAndDeletePolicyResponseBody) SetRequestId(v string) *DetachAndDeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachAndDeletePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
