// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndAttachPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAndAttachPolicyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAndAttachPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAndAttachPolicyResponseBody
	GetRequestId() *string
}

type UpdateAndAttachPolicyResponseBody struct {
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
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAndAttachPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndAttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAndAttachPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAndAttachPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAndAttachPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAndAttachPolicyResponseBody) SetCode(v string) *UpdateAndAttachPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAndAttachPolicyResponseBody) SetMessage(v string) *UpdateAndAttachPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAndAttachPolicyResponseBody) SetRequestId(v string) *UpdateAndAttachPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAndAttachPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
