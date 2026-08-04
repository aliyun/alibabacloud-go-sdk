// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomerLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddCustomerLabelResponseBody
	GetCode() *string
	SetMessage(v string) *AddCustomerLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomerLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCustomerLabelResponseBody
	GetSuccess() *bool
}

type AddCustomerLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCustomerLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomerLabelResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomerLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddCustomerLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomerLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomerLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCustomerLabelResponseBody) SetCode(v string) *AddCustomerLabelResponseBody {
	s.Code = &v
	return s
}

func (s *AddCustomerLabelResponseBody) SetMessage(v string) *AddCustomerLabelResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomerLabelResponseBody) SetRequestId(v string) *AddCustomerLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomerLabelResponseBody) SetSuccess(v bool) *AddCustomerLabelResponseBody {
	s.Success = &v
	return s
}

func (s *AddCustomerLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
