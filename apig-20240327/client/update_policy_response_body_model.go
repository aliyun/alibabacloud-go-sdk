// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePolicyResponseBody
	GetRequestId() *string
}

type UpdatePolicyResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854C***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyResponseBody) SetCode(v string) *UpdatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetMessage(v string) *UpdatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetRequestId(v string) *UpdatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
