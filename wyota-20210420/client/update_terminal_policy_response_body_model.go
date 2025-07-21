// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTerminalPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTerminalPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTerminalPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTerminalPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTerminalPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTerminalPolicyResponseBody
	GetSuccess() *bool
}

type UpdateTerminalPolicyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTerminalPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTerminalPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTerminalPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTerminalPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTerminalPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTerminalPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTerminalPolicyResponseBody) SetCode(v string) *UpdateTerminalPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateTerminalPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetMessage(v string) *UpdateTerminalPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetRequestId(v string) *UpdateTerminalPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) SetSuccess(v bool) *UpdateTerminalPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTerminalPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
