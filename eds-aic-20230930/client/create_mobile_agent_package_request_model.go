// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateMobileAgentPackageRequest
	GetAmount() *string
	SetAutoPay(v bool) *CreateMobileAgentPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateMobileAgentPackageRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateMobileAgentPackageRequest
	GetBizRegionId() *string
	SetCreditAmount(v string) *CreateMobileAgentPackageRequest
	GetCreditAmount() *string
	SetCreditConfig(v string) *CreateMobileAgentPackageRequest
	GetCreditConfig() *string
	SetImageId(v string) *CreateMobileAgentPackageRequest
	GetImageId() *string
	SetInstanceName(v string) *CreateMobileAgentPackageRequest
	GetInstanceName() *string
	SetMobileAgentPackageSpec(v string) *CreateMobileAgentPackageRequest
	GetMobileAgentPackageSpec() *string
	SetPackageSpecId(v int64) *CreateMobileAgentPackageRequest
	GetPackageSpecId() *int64
	SetPaidCallbackUrl(v string) *CreateMobileAgentPackageRequest
	GetPaidCallbackUrl() *string
	SetPeriod(v int32) *CreateMobileAgentPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateMobileAgentPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateMobileAgentPackageRequest
	GetPromotionId() *string
}

type CreateMobileAgentPackageRequest struct {
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 10000.0
	CreditAmount *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// example:
	//
	// 1
	CreditConfig *string `json:"CreditConfig,omitempty" xml:"CreditConfig,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// CloudPhone
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// advanced
	MobileAgentPackageSpec *string `json:"MobileAgentPackageSpec,omitempty" xml:"MobileAgentPackageSpec,omitempty"`
	PackageSpecId          *int64  `json:"PackageSpecId,omitempty" xml:"PackageSpecId,omitempty"`
	// example:
	//
	// https://aim.wuying.aliyun.com/nodes
	PaidCallbackUrl *string `json:"PaidCallbackUrl,omitempty" xml:"PaidCallbackUrl,omitempty"`
	// example:
	//
	// 1
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

func (s CreateMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentPackageRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateMobileAgentPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateMobileAgentPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateMobileAgentPackageRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateMobileAgentPackageRequest) GetCreditAmount() *string {
	return s.CreditAmount
}

func (s *CreateMobileAgentPackageRequest) GetCreditConfig() *string {
	return s.CreditConfig
}

func (s *CreateMobileAgentPackageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateMobileAgentPackageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateMobileAgentPackageRequest) GetMobileAgentPackageSpec() *string {
	return s.MobileAgentPackageSpec
}

func (s *CreateMobileAgentPackageRequest) GetPackageSpecId() *int64 {
	return s.PackageSpecId
}

func (s *CreateMobileAgentPackageRequest) GetPaidCallbackUrl() *string {
	return s.PaidCallbackUrl
}

func (s *CreateMobileAgentPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateMobileAgentPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateMobileAgentPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateMobileAgentPackageRequest) SetAmount(v string) *CreateMobileAgentPackageRequest {
	s.Amount = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetAutoPay(v bool) *CreateMobileAgentPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetAutoRenew(v bool) *CreateMobileAgentPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetBizRegionId(v string) *CreateMobileAgentPackageRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetCreditAmount(v string) *CreateMobileAgentPackageRequest {
	s.CreditAmount = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetCreditConfig(v string) *CreateMobileAgentPackageRequest {
	s.CreditConfig = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetImageId(v string) *CreateMobileAgentPackageRequest {
	s.ImageId = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetInstanceName(v string) *CreateMobileAgentPackageRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetMobileAgentPackageSpec(v string) *CreateMobileAgentPackageRequest {
	s.MobileAgentPackageSpec = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetPackageSpecId(v int64) *CreateMobileAgentPackageRequest {
	s.PackageSpecId = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetPaidCallbackUrl(v string) *CreateMobileAgentPackageRequest {
	s.PaidCallbackUrl = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetPeriod(v int32) *CreateMobileAgentPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetPeriodUnit(v string) *CreateMobileAgentPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) SetPromotionId(v string) *CreateMobileAgentPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateMobileAgentPackageRequest) Validate() error {
	return dara.Validate(s)
}
