// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateOrderRequest
	GetClientToken() *string
	SetCommodity(v string) *CreateOrderRequest
	GetCommodity() *string
	SetOrderSouce(v string) *CreateOrderRequest
	GetOrderSouce() *string
	SetOrderType(v string) *CreateOrderRequest
	GetOrderType() *string
	SetOwnerId(v string) *CreateOrderRequest
	GetOwnerId() *string
	SetPaymentType(v string) *CreateOrderRequest
	GetPaymentType() *string
}

type CreateOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dhjggh7
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"productCode\\":\\"cmgj00063418\\",\\"skuCode\\":\\"postpay\\",\\"pricingCycle\\":\\"YEAR\\"}
	Commodity *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// example:
	//
	// ros
	OrderSouce *string `json:"OrderSouce,omitempty" xml:"OrderSouce,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE_BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOrderRequest) GetCommodity() *string {
	return s.Commodity
}

func (s *CreateOrderRequest) GetOrderSouce() *string {
	return s.OrderSouce
}

func (s *CreateOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateOrderRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateOrderRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateOrderRequest) SetClientToken(v string) *CreateOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOrderRequest) SetCommodity(v string) *CreateOrderRequest {
	s.Commodity = &v
	return s
}

func (s *CreateOrderRequest) SetOrderSouce(v string) *CreateOrderRequest {
	s.OrderSouce = &v
	return s
}

func (s *CreateOrderRequest) SetOrderType(v string) *CreateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateOrderRequest) SetOwnerId(v string) *CreateOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrderRequest) SetPaymentType(v string) *CreateOrderRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	return dara.Validate(s)
}
