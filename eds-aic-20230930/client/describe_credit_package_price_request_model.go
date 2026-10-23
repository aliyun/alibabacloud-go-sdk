// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackagePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreditAmount(v int32) *DescribeCreditPackagePriceRequest
	GetCreditAmount() *int32
	SetPackageAmount(v string) *DescribeCreditPackagePriceRequest
	GetPackageAmount() *string
	SetPeriod(v string) *DescribeCreditPackagePriceRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *DescribeCreditPackagePriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeCreditPackagePriceRequest
	GetPromotionId() *string
}

type DescribeCreditPackagePriceRequest struct {
	// example:
	//
	// 1000
	CreditAmount *int32 `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	// example:
	//
	// 1
	PackageAmount *string `json:"PackageAmount,omitempty" xml:"PackageAmount,omitempty"`
	// example:
	//
	// 3
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s DescribeCreditPackagePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceRequest) GetCreditAmount() *int32 {
	return s.CreditAmount
}

func (s *DescribeCreditPackagePriceRequest) GetPackageAmount() *string {
	return s.PackageAmount
}

func (s *DescribeCreditPackagePriceRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeCreditPackagePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeCreditPackagePriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCreditPackagePriceRequest) SetCreditAmount(v int32) *DescribeCreditPackagePriceRequest {
	s.CreditAmount = &v
	return s
}

func (s *DescribeCreditPackagePriceRequest) SetPackageAmount(v string) *DescribeCreditPackagePriceRequest {
	s.PackageAmount = &v
	return s
}

func (s *DescribeCreditPackagePriceRequest) SetPeriod(v string) *DescribeCreditPackagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeCreditPackagePriceRequest) SetPeriodUnit(v string) *DescribeCreditPackagePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeCreditPackagePriceRequest) SetPromotionId(v string) *DescribeCreditPackagePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeCreditPackagePriceRequest) Validate() error {
	return dara.Validate(s)
}
