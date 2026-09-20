// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConvertInstanceRequest
	GetClusterId() *string
	SetDuration(v int32) *ConvertInstanceRequest
	GetDuration() *int32
	SetPayType(v string) *ConvertInstanceRequest
	GetPayType() *string
	SetPricingCycle(v string) *ConvertInstanceRequest
	GetPricingCycle() *string
}

type ConvertInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The subscription duration. Valid values:
	//
	// - If PricingCycle is set to year, the value ranges from 1 to 3.
	//
	// - If PricingCycle is set to month, the value ranges from 1 to 9.
	//
	// example:
	//
	// 7
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Prepaid**: subscription.
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// - year: year.
	//
	// - month: month.
	//
	// example:
	//
	// month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s ConvertInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConvertInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *ConvertInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConvertInstanceRequest) SetClusterId(v string) *ConvertInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ConvertInstanceRequest) SetDuration(v int32) *ConvertInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequest) SetPayType(v string) *ConvertInstanceRequest {
	s.PayType = &v
	return s
}

func (s *ConvertInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequest) Validate() error {
	return dara.Validate(s)
}
