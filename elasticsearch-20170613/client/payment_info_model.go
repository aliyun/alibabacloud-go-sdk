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
	AutoRenewDuration *int64  `json:"autoRenewDuration,omitempty" xml:"autoRenewDuration,omitempty"`
	Duration          *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	IsAutoRenew       *bool   `json:"isAutoRenew,omitempty" xml:"isAutoRenew,omitempty"`
	PricingCycle      *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
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
