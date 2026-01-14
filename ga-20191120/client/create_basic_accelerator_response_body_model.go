// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicAcceleratorResponseBody
	GetAcceleratorId() *string
	SetOrderId(v string) *CreateBasicAcceleratorResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateBasicAcceleratorResponseBody
	GetRequestId() *string
}

type CreateBasicAcceleratorResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The order ID.
	//
	// This parameter is returned only if ChargeType is set to PREPAY.
	//
	// If **AutoPay*	- is set to **false**, go to [Order Center](https://usercenter2-intl.aliyun.com/order/list) to complete the payment after an order is generated.
	//
	// example:
	//
	// 2082574365
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAcceleratorResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateBasicAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicAcceleratorResponseBody) SetAcceleratorId(v string) *CreateBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAcceleratorResponseBody) SetOrderId(v string) *CreateBasicAcceleratorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateBasicAcceleratorResponseBody) SetRequestId(v string) *CreateBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
