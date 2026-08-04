// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgServiceStatusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAgServiceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgServiceStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgServiceStatusResponseBody
	GetSuccess() *bool
}

type UpdateAgServiceStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAgServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgServiceStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgServiceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgServiceStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgServiceStatusResponseBody) SetCode(v string) *UpdateAgServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgServiceStatusResponseBody) SetMessage(v string) *UpdateAgServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgServiceStatusResponseBody) SetRequestId(v string) *UpdateAgServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgServiceStatusResponseBody) SetSuccess(v bool) *UpdateAgServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
