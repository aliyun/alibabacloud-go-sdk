// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewMobileAgentPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewMobileAgentPackageRequest
	GetAutoRenew() *bool
	SetMobileAgentPackageIds(v []*string) *RenewMobileAgentPackageRequest
	GetMobileAgentPackageIds() []*string
	SetPaidCallbackUrl(v string) *RenewMobileAgentPackageRequest
	GetPaidCallbackUrl() *string
	SetPeriod(v int32) *RenewMobileAgentPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewMobileAgentPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewMobileAgentPackageRequest
	GetPromotionId() *string
}

type RenewMobileAgentPackageRequest struct {
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew             *bool     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	MobileAgentPackageIds []*string `json:"MobileAgentPackageIds,omitempty" xml:"MobileAgentPackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// https://aim.wuying.aliyun.com/nodes
	PaidCallbackUrl *string `json:"PaidCallbackUrl,omitempty" xml:"PaidCallbackUrl,omitempty"`
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s RenewMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *RenewMobileAgentPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewMobileAgentPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewMobileAgentPackageRequest) GetMobileAgentPackageIds() []*string {
	return s.MobileAgentPackageIds
}

func (s *RenewMobileAgentPackageRequest) GetPaidCallbackUrl() *string {
	return s.PaidCallbackUrl
}

func (s *RenewMobileAgentPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewMobileAgentPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewMobileAgentPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewMobileAgentPackageRequest) SetAutoPay(v bool) *RenewMobileAgentPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetAutoRenew(v bool) *RenewMobileAgentPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetMobileAgentPackageIds(v []*string) *RenewMobileAgentPackageRequest {
	s.MobileAgentPackageIds = v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetPaidCallbackUrl(v string) *RenewMobileAgentPackageRequest {
	s.PaidCallbackUrl = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetPeriod(v int32) *RenewMobileAgentPackageRequest {
	s.Period = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetPeriodUnit(v string) *RenewMobileAgentPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) SetPromotionId(v string) *RenewMobileAgentPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewMobileAgentPackageRequest) Validate() error {
	return dara.Validate(s)
}
