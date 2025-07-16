// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifySecurityIpsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySecurityIpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifySecurityIpsResponseBody
	GetSuccess() *bool
}

type ModifySecurityIpsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifySecurityIpsResponseBody) SetMessage(v string) *ModifySecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetSuccess(v bool) *ModifySecurityIpsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
