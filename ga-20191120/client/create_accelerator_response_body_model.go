// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateAcceleratorResponseBody
	GetAcceleratorId() *string
	SetOrderId(v string) *CreateAcceleratorResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateAcceleratorResponseBody
	GetRequestId() *string
}

type CreateAcceleratorResponseBody struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Order ID.
	//
	// <props="china">
	//
	// The order ID. This parameter is returned when \\`InstanceChargeType\\` is set to \\`PREPAY\\` (subscription). If \\`AutoPay\\` is set to \\`false\\`, go to the [Order Hub](https://usercenter2.aliyun.com/order/list) on the Alibaba Cloud China site to complete the payment.
	//
	//
	//
	// <props="intl">
	//
	// If you are using the Alibaba Cloud International site and \\`AutoPay\\` is set to \\`false\\`, go to the [Order Hub](https://usercenter2-intl.aliyun.com/order/list) to complete the payment.
	//
	// example:
	//
	// 208257****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateAcceleratorResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAcceleratorResponseBody) SetAcceleratorId(v string) *CreateAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetOrderId(v string) *CreateAcceleratorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) SetRequestId(v string) *CreateAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
