// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountProfileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAccountProfileInfoResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAccountProfileInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAccountProfileInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAccountProfileInfoResponseBody
	GetSuccess() *bool
}

type UpdateAccountProfileInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAccountProfileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountProfileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountProfileInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAccountProfileInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAccountProfileInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountProfileInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAccountProfileInfoResponseBody) SetCode(v string) *UpdateAccountProfileInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountProfileInfoResponseBody) SetMessage(v string) *UpdateAccountProfileInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountProfileInfoResponseBody) SetRequestId(v string) *UpdateAccountProfileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountProfileInfoResponseBody) SetSuccess(v bool) *UpdateAccountProfileInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAccountProfileInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
