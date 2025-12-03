// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyInstanceTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyInstanceTypeResponseBody
	GetRequestId() *string
}

type ModifyInstanceTypeResponseBody struct {
	// example:
	//
	// 123412341234123
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3E19E345-101D-4014-946C-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceTypeResponseBody) SetOrderId(v string) *ModifyInstanceTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceTypeResponseBody) SetRequestId(v string) *ModifyInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
