// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelResponseBody
	GetSuccess() *bool
}

type UpdateModelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelResponseBody) SetCode(v string) *UpdateModelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelResponseBody) SetMessage(v string) *UpdateModelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelResponseBody) SetRequestId(v string) *UpdateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelResponseBody) SetSuccess(v bool) *UpdateModelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelResponseBody) Validate() error {
	return dara.Validate(s)
}
