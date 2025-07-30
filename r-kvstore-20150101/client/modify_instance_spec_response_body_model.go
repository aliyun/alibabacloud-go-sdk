// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyInstanceSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyInstanceSpecResponseBody
	GetRequestId() *string
}

type ModifyInstanceSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 20722623431****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0DA1D7EF-C80D-432C-8758-7D225182626B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSpecResponseBody) SetOrderId(v string) *ModifyInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) SetRequestId(v string) *ModifyInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
