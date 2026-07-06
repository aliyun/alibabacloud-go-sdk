// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentStoragePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckAgentStoragePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *CheckAgentStoragePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckAgentStoragePolicyResponseBody
	GetRequestId() *string
}

type CheckAgentStoragePolicyResponseBody struct {
	// The response status code.
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
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// E734979F-5A44-5993-9CE5-C23103576923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAgentStoragePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentStoragePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAgentStoragePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckAgentStoragePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckAgentStoragePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckAgentStoragePolicyResponseBody) SetCode(v string) *CheckAgentStoragePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAgentStoragePolicyResponseBody) SetMessage(v string) *CheckAgentStoragePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAgentStoragePolicyResponseBody) SetRequestId(v string) *CheckAgentStoragePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAgentStoragePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
