// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoPhysicalDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CustomerSensitiveInfoPhysicalDeleteResponseBody
	GetSuccess() *bool
}

type CustomerSensitiveInfoPhysicalDeleteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CustomerSensitiveInfoPhysicalDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoPhysicalDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) SetCode(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) SetMessage(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) SetRequestId(v string) *CustomerSensitiveInfoPhysicalDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) SetSuccess(v bool) *CustomerSensitiveInfoPhysicalDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
