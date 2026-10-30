// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *DescribeMobileAgentPackagePriceRequest
	GetAmount() *string
	SetBizRegionId(v string) *DescribeMobileAgentPackagePriceRequest
	GetBizRegionId() *string
	SetCreditAmount(v int32) *DescribeMobileAgentPackagePriceRequest
	GetCreditAmount() *int32
	SetMobileAgentPackageSpec(v string) *DescribeMobileAgentPackagePriceRequest
	GetMobileAgentPackageSpec() *string
	SetPeriod(v string) *DescribeMobileAgentPackagePriceRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *DescribeMobileAgentPackagePriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeMobileAgentPackagePriceRequest
	GetPromotionId() *string
}

type DescribeMobileAgentPackagePriceRequest struct {
	// The number of MobileAgent packages.
	//
	// example:
	//
	// 1
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The region ID for the purchase.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The number of credits.
	//
	// example:
	//
	// 1000
	CreditAmount *int32 `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// The specifications of the MobileAgent package.
	//
	// example:
	//
	// advanced
	MobileAgentPackageSpec *string `json:"MobileAgentPackageSpec,omitempty" xml:"MobileAgentPackageSpec,omitempty"`
	// The subscription duration of the purchased resource. The unit is specified by PeriodUnit. Valid value: 1 month.
	//
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The duration unit of the subscription billing method. Valid values:
	//
	// - **Year**: Year.
	//
	// - **Month*	- (default): Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotion.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s DescribeMobileAgentPackagePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceRequest) GetAmount() *string {
	return s.Amount
}

func (s *DescribeMobileAgentPackagePriceRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeMobileAgentPackagePriceRequest) GetCreditAmount() *int32 {
	return s.CreditAmount
}

func (s *DescribeMobileAgentPackagePriceRequest) GetMobileAgentPackageSpec() *string {
	return s.MobileAgentPackageSpec
}

func (s *DescribeMobileAgentPackagePriceRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMobileAgentPackagePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeMobileAgentPackagePriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMobileAgentPackagePriceRequest) SetAmount(v string) *DescribeMobileAgentPackagePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetBizRegionId(v string) *DescribeMobileAgentPackagePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetCreditAmount(v int32) *DescribeMobileAgentPackagePriceRequest {
	s.CreditAmount = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetMobileAgentPackageSpec(v string) *DescribeMobileAgentPackagePriceRequest {
	s.MobileAgentPackageSpec = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetPeriod(v string) *DescribeMobileAgentPackagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetPeriodUnit(v string) *DescribeMobileAgentPackagePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) SetPromotionId(v string) *DescribeMobileAgentPackagePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceRequest) Validate() error {
	return dara.Validate(s)
}
