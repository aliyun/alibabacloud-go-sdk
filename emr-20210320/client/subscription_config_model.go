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
	// 自动续费。取值范围：
	//
	// - true：开启启动续费。
	//
	// - false：不开启自动续费。
	//
	// 默认值：false。
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// 自动续费时长。当AutoRenew取值为true时生效。当AutoRenewDurationUnit取值为Month时，取值：1、2、3、4、5、6、7、8、9、12、24、36、48、60。
	//
	// example:
	//
	// 12
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// - Month：月。
	//
	// example:
	//
	// Month
	AutoRenewDurationUnit *string `json:"AutoRenewDurationUnit,omitempty" xml:"AutoRenewDurationUnit,omitempty"`
	// 付费时长。PaymentDurationUnit取值为Month时，取值：1、2、3、4、5、6、7、8、9、12、24、36、48、60。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// 付费时长单位。取值范围：
	//
	// - Month：月。
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
