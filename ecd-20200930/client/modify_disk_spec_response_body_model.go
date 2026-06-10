// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyDiskSpecResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDiskSpecResponseBody
	GetRequestId() *string
}

type ModifyDiskSpecResponseBody struct {
	// The order ID. You can find the order ID in <props="china">[Expenses and Costs > My Orders](https://billing-cost.console.aliyun.com/order/list). <props="intl">[Expenses and Costs > Order Management](https://usercenter2-intl.aliyun.com/order/list).
	//
	// example:
	//
	// 219861020660568
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F89BBB13-8B3B-5C8A-A700-EEFDC17B8227
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDiskSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskSpecResponseBody) SetOrderId(v string) *ModifyDiskSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDiskSpecResponseBody) SetRequestId(v string) *ModifyDiskSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
