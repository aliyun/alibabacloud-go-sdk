// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstancePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateInstancePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstancePolicyResponseBody
	GetRequestId() *string
}

type UpdateInstancePolicyResponseBody struct {
	// The HTTP status code.
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
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 31D8120C-AC52-5CA9-BE4A-E4C6316E19AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateInstancePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstancePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstancePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstancePolicyResponseBody) SetCode(v string) *UpdateInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstancePolicyResponseBody) SetMessage(v string) *UpdateInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstancePolicyResponseBody) SetRequestId(v string) *UpdateInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstancePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
