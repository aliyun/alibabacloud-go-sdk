// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspirationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmountspec(v string) *CreateInspirationRequest
	GetAmountspec() *string
	SetClientToken(v string) *CreateInspirationRequest
	GetClientToken() *string
	SetDuration(v int32) *CreateInspirationRequest
	GetDuration() *int32
	SetExtend(v string) *CreateInspirationRequest
	GetExtend() *string
	SetPaymentType(v string) *CreateInspirationRequest
	GetPaymentType() *string
	SetPricingCycle(v string) *CreateInspirationRequest
	GetPricingCycle() *string
	SetQuantity(v int32) *CreateInspirationRequest
	GetQuantity() *int32
}

type CreateInspirationRequest struct {
	// The specification of the resource plan.
	//
	// example:
	//
	// 1000
	Amountspec *string `json:"Amountspec,omitempty" xml:"Amountspec,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// 111
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The extended information in JSON format.
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The payment type.
	//
	// example:
	//
	// AUTO_PAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The unit of the subscription cycle. Valid values:
	//
	// - Year: year
	//
	// - Month: month
	//
	// - Day: day
	//
	// - Hour: hour
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The number of instances to subscribe to.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s CreateInspirationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInspirationRequest) GoString() string {
	return s.String()
}

func (s *CreateInspirationRequest) GetAmountspec() *string {
	return s.Amountspec
}

func (s *CreateInspirationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInspirationRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInspirationRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateInspirationRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInspirationRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInspirationRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateInspirationRequest) SetAmountspec(v string) *CreateInspirationRequest {
	s.Amountspec = &v
	return s
}

func (s *CreateInspirationRequest) SetClientToken(v string) *CreateInspirationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInspirationRequest) SetDuration(v int32) *CreateInspirationRequest {
	s.Duration = &v
	return s
}

func (s *CreateInspirationRequest) SetExtend(v string) *CreateInspirationRequest {
	s.Extend = &v
	return s
}

func (s *CreateInspirationRequest) SetPaymentType(v string) *CreateInspirationRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInspirationRequest) SetPricingCycle(v string) *CreateInspirationRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInspirationRequest) SetQuantity(v int32) *CreateInspirationRequest {
	s.Quantity = &v
	return s
}

func (s *CreateInspirationRequest) Validate() error {
	return dara.Validate(s)
}
