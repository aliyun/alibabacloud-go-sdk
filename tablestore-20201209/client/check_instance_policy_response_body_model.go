// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstancePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckInstancePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *CheckInstancePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckInstancePolicyResponseBody
	GetRequestId() *string
}

type CheckInstancePolicyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Verification passed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 757E172A-F94B-5E78-8A23-D9068E42F2E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstancePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckInstancePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckInstancePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstancePolicyResponseBody) SetCode(v string) *CheckInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckInstancePolicyResponseBody) SetMessage(v string) *CheckInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckInstancePolicyResponseBody) SetRequestId(v string) *CheckInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstancePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
