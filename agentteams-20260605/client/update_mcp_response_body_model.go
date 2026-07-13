// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMcpResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateMcpResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMcpResponseBody
	GetSuccess() *bool
}

type UpdateMcpResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMcpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMcpResponseBody) SetCode(v string) *UpdateMcpResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMcpResponseBody) SetMessage(v string) *UpdateMcpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMcpResponseBody) SetRequestId(v string) *UpdateMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpResponseBody) SetSuccess(v bool) *UpdateMcpResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
