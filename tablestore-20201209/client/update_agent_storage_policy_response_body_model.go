// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStoragePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentStoragePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAgentStoragePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentStoragePolicyResponseBody
	GetRequestId() *string
}

type UpdateAgentStoragePolicyResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// B37BBA04-D827-55C8-B901-5264B904E8C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAgentStoragePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStoragePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentStoragePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentStoragePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentStoragePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentStoragePolicyResponseBody) SetCode(v string) *UpdateAgentStoragePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentStoragePolicyResponseBody) SetMessage(v string) *UpdateAgentStoragePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentStoragePolicyResponseBody) SetRequestId(v string) *UpdateAgentStoragePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentStoragePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
