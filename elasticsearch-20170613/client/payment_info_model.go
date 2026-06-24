// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaymentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewDuration(v int64) *PaymentInfo
	GetAutoRenewDuration() *int64
	SetDuration(v int64) *PaymentInfo
	GetDuration() *int64
	SetIsAutoRenew(v bool) *PaymentInfo
	GetIsAutoRenew() *bool
	SetPricingCycle(v string) *PaymentInfo
	GetPricingCycle() *string
}

type PaymentInfo struct {
	// The auto-renewal cycle. Unit: month. This parameter is required when **isAutoRenew*	- is set to **true**. The valid values are the same as those on the buy page.
	//
	// example:
	//
	// 3
	AutoRenewDuration *int64 `json:"autoRenewDuration,omitempty" xml:"autoRenewDuration,omitempty"`
	// The subscription duration. This parameter is required. You can specify the duration in months or years.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: enabled.
	//
	// - false (default): disabled.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"isAutoRenew,omitempty" xml:"isAutoRenew,omitempty"`
	// The unit of the subscription duration. This parameter is required. Valid values:
	//
	// - Year: year
	//
	// - Month: month.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s PaymentInfo) String() string {
	return dara.Prettify(s)
}

func (s PaymentInfo) GoString() string {
	return s.String()
}

func (s *PaymentInfo) GetAutoRenewDuration() *int64 {
	return s.AutoRenewDuration
}

func (s *PaymentInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *PaymentInfo) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *PaymentInfo) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *PaymentInfo) SetAutoRenewDuration(v int64) *PaymentInfo {
	s.AutoRenewDuration = &v
	return s
}

func (s *PaymentInfo) SetDuration(v int64) *PaymentInfo {
	s.Duration = &v
	return s
}

func (s *PaymentInfo) SetIsAutoRenew(v bool) *PaymentInfo {
	s.IsAutoRenew = &v
	return s
}

func (s *PaymentInfo) SetPricingCycle(v string) *PaymentInfo {
	s.PricingCycle = &v
	return s
}

func (s *PaymentInfo) Validate() error {
	return dara.Validate(s)
}
