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
	// The IDs of the MobileAgent resource plans.
	MobileAgentPackageIds []*string `json:"MobileAgentPackageIds,omitempty" xml:"MobileAgentPackageIds,omitempty" type:"Repeated"`
	// The subscription period of the resource. The unit is specified by PeriodUnit. Valid values: 1 month, 2 months, 3 months, 6 months, and 1 year.
	//
	// example:
	//
	// 3
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// - **Month**: Month.
	//
	// - **Year**: Year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
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
