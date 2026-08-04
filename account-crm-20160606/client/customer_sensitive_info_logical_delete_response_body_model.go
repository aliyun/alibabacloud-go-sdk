// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoLogicalDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CustomerSensitiveInfoLogicalDeleteResponseBody
	GetSuccess() *bool
}

type CustomerSensitiveInfoLogicalDeleteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CustomerSensitiveInfoLogicalDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoLogicalDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) SetCode(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) SetMessage(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) SetRequestId(v string) *CustomerSensitiveInfoLogicalDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) SetSuccess(v bool) *CustomerSensitiveInfoLogicalDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
