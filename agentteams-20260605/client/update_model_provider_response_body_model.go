// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelProviderResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateModelProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelProviderResponseBody
	GetSuccess() *bool
}

type UpdateModelProviderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModelProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelProviderResponseBody) SetCode(v string) *UpdateModelProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelProviderResponseBody) SetMessage(v string) *UpdateModelProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelProviderResponseBody) SetRequestId(v string) *UpdateModelProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelProviderResponseBody) SetSuccess(v bool) *UpdateModelProviderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
