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
	// The ID of the order.
	//
	// This parameter is returned only when the InstanceChargeType parameter is set to PREPAY. If AutoPay is set to false, you must manually complete the payment in the [Order Center](https://usercenter2-intl.aliyun.com/order/list).
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
