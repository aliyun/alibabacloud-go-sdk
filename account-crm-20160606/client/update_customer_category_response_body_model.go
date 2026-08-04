// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCustomerCategoryResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateCustomerCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomerCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomerCategoryResponseBody
	GetSuccess() *bool
}

type UpdateCustomerCategoryResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomerCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomerCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCustomerCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomerCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomerCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomerCategoryResponseBody) SetCode(v string) *UpdateCustomerCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomerCategoryResponseBody) SetMessage(v string) *UpdateCustomerCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomerCategoryResponseBody) SetRequestId(v string) *UpdateCustomerCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomerCategoryResponseBody) SetSuccess(v bool) *UpdateCustomerCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomerCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
