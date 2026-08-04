// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCustomerInfoResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyCustomerInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCustomerInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCustomerInfoResponseBody
	GetSuccess() *bool
}

type ModifyCustomerInfoResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCustomerInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomerInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCustomerInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCustomerInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomerInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCustomerInfoResponseBody) SetCode(v string) *ModifyCustomerInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCustomerInfoResponseBody) SetMessage(v string) *ModifyCustomerInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCustomerInfoResponseBody) SetRequestId(v string) *ModifyCustomerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomerInfoResponseBody) SetSuccess(v bool) *ModifyCustomerInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCustomerInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
