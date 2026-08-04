// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgAccountAddressInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgAccountAddressInfoResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAgAccountAddressInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgAccountAddressInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgAccountAddressInfoResponseBody
	GetSuccess() *bool
}

type UpdateAgAccountAddressInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAgAccountAddressInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgAccountAddressInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgAccountAddressInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgAccountAddressInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgAccountAddressInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgAccountAddressInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgAccountAddressInfoResponseBody) SetCode(v string) *UpdateAgAccountAddressInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgAccountAddressInfoResponseBody) SetMessage(v string) *UpdateAgAccountAddressInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgAccountAddressInfoResponseBody) SetRequestId(v string) *UpdateAgAccountAddressInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgAccountAddressInfoResponseBody) SetSuccess(v bool) *UpdateAgAccountAddressInfoResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgAccountAddressInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
