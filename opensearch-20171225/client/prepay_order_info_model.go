// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepayOrderInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *PrepayOrderInfo
	GetAutoRenew() *bool
	SetDuration(v int32) *PrepayOrderInfo
	GetDuration() *int32
	SetPricingCycle(v string) *PrepayOrderInfo
	GetPricingCycle() *string
}

type PrepayOrderInfo struct {
	AutoRenew    *bool   `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	Duration     *int32  `json:"duration,omitempty" xml:"duration,omitempty"`
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s PrepayOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s PrepayOrderInfo) GoString() string {
	return s.String()
}

func (s *PrepayOrderInfo) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *PrepayOrderInfo) GetDuration() *int32 {
	return s.Duration
}

func (s *PrepayOrderInfo) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *PrepayOrderInfo) SetAutoRenew(v bool) *PrepayOrderInfo {
	s.AutoRenew = &v
	return s
}

func (s *PrepayOrderInfo) SetDuration(v int32) *PrepayOrderInfo {
	s.Duration = &v
	return s
}

func (s *PrepayOrderInfo) SetPricingCycle(v string) *PrepayOrderInfo {
	s.PricingCycle = &v
	return s
}

func (s *PrepayOrderInfo) Validate() error {
	return dara.Validate(s)
}
