// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *SubscriptionConfig
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *SubscriptionConfig
	GetAutoRenewDuration() *int32
	SetAutoRenewDurationUnit(v string) *SubscriptionConfig
	GetAutoRenewDurationUnit() *string
	SetPaymentDuration(v int32) *SubscriptionConfig
	GetPaymentDuration() *int32
	SetPaymentDurationUnit(v string) *SubscriptionConfig
	GetPaymentDurationUnit() *string
}

type SubscriptionConfig struct {
	// Specifies whether auto-renewal is enabled. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled (default).
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. This parameter takes effect only when AutoRenew is set to true. When AutoRenewDurationUnit is Month, valid values are: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 12
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The auto-renewal duration unit. Valid value:
	//
	// - Month
	//
	// example:
	//
	// Month
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	// The payment duration. When PaymentDurationUnit is Month, valid values are: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// The payment duration unit. Valid value:
	//
	// - Month
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
}

func (s SubscriptionConfig) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionConfig) GoString() string {
	return s.String()
}

func (s *SubscriptionConfig) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *SubscriptionConfig) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *SubscriptionConfig) GetAutoRenewDurationUnit() *string {
	return s.AutoRenewDurationUnit
}

func (s *SubscriptionConfig) GetPaymentDuration() *int32 {
	return s.PaymentDuration
}

func (s *SubscriptionConfig) GetPaymentDurationUnit() *string {
	return s.PaymentDurationUnit
}

func (s *SubscriptionConfig) SetAutoRenew(v bool) *SubscriptionConfig {
	s.AutoRenew = &v
	return s
}

func (s *SubscriptionConfig) SetAutoRenewDuration(v int32) *SubscriptionConfig {
	s.AutoRenewDuration = &v
	return s
}

func (s *SubscriptionConfig) SetAutoRenewDurationUnit(v string) *SubscriptionConfig {
	s.AutoRenewDurationUnit = &v
	return s
}

func (s *SubscriptionConfig) SetPaymentDuration(v int32) *SubscriptionConfig {
	s.PaymentDuration = &v
	return s
}

func (s *SubscriptionConfig) SetPaymentDurationUnit(v string) *SubscriptionConfig {
	s.PaymentDurationUnit = &v
	return s
}

func (s *SubscriptionConfig) Validate() error {
	return dara.Validate(s)
}
