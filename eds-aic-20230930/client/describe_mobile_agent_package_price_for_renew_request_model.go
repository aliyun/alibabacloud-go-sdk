// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceForRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobileAgentPackageIds(v []*string) *DescribeMobileAgentPackagePriceForRenewRequest
	GetMobileAgentPackageIds() []*string
	SetPeriod(v string) *DescribeMobileAgentPackagePriceForRenewRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *DescribeMobileAgentPackagePriceForRenewRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeMobileAgentPackagePriceForRenewRequest
	GetPromotionId() *string
}

type DescribeMobileAgentPackagePriceForRenewRequest struct {
	MobileAgentPackageIds []*string `json:"MobileAgentPackageIds,omitempty" xml:"MobileAgentPackageIds,omitempty" type:"Repeated"`
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

func (s DescribeMobileAgentPackagePriceForRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewRequest) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) GetMobileAgentPackageIds() []*string {
	return s.MobileAgentPackageIds
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) SetMobileAgentPackageIds(v []*string) *DescribeMobileAgentPackagePriceForRenewRequest {
	s.MobileAgentPackageIds = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) SetPeriod(v string) *DescribeMobileAgentPackagePriceForRenewRequest {
	s.Period = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) SetPeriodUnit(v string) *DescribeMobileAgentPackagePriceForRenewRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) SetPromotionId(v string) *DescribeMobileAgentPackagePriceForRenewRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewRequest) Validate() error {
	return dara.Validate(s)
}
