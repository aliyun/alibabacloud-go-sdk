// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountAddressInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAccountAddressInfoResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAccountAddressInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAccountAddressInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAccountAddressInfoResponseBody
	GetSuccess() *bool
}

type UpdateAccountAddressInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAccountAddressInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountAddressInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountAddressInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAccountAddressInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAccountAddressInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountAddressInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAccountAddressInfoResponseBody) SetCode(v string) *UpdateAccountAddressInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountAddressInfoResponseBody) SetMessage(v string) *UpdateAccountAddressInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountAddressInfoResponseBody) SetRequestId(v string) *UpdateAccountAddressInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountAddressInfoResponseBody) SetSuccess(v bool) *UpdateAccountAddressInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAccountAddressInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
