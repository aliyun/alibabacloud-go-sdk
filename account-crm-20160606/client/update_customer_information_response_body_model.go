// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCustomerInformationResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateCustomerInformationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomerInformationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomerInformationResponseBody
	GetSuccess() *bool
}

type UpdateCustomerInformationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomerInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomerInformationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomerInformationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomerInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomerInformationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomerInformationResponseBody) SetCode(v string) *UpdateCustomerInformationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomerInformationResponseBody) SetMessage(v string) *UpdateCustomerInformationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomerInformationResponseBody) SetRequestId(v string) *UpdateCustomerInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomerInformationResponseBody) SetSuccess(v bool) *UpdateCustomerInformationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomerInformationResponseBody) Validate() error {
	return dara.Validate(s)
}
