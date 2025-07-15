// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBandwidthResourcePackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateBandwidthResourcePackagesRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateBandwidthResourcePackagesRequest
	GetAutoPay() *bool
	SetPackageSize(v int32) *CreateBandwidthResourcePackagesRequest
	GetPackageSize() *int32
	SetPeriod(v int32) *CreateBandwidthResourcePackagesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateBandwidthResourcePackagesRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateBandwidthResourcePackagesRequest
	GetPromotionId() *string
	SetRegionId(v string) *CreateBandwidthResourcePackagesRequest
	GetRegionId() *string
}

type CreateBandwidthResourcePackagesRequest struct {
	// The number of the data transfer plans that you want to create at the same time. Valid values: 1 to 20. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable the auto-payment feature.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The size of the data transfer plan. Valid values: 10 to 1000. Unit: GiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PackageSize *int32 `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	// The subscription duration. The valid values of this parameter vary based on the value of `PeriodUnit`.
	//
	// 	- If `PeriodUnit` is set to `Month`, the valid values of Period are 1, 3, and 6.
	//
	// 	- If `PeriodUnit` is set to `Year`, the valid value of Period is 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotional activity.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBandwidthResourcePackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBandwidthResourcePackagesRequest) GoString() string {
	return s.String()
}

func (s *CreateBandwidthResourcePackagesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateBandwidthResourcePackagesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateBandwidthResourcePackagesRequest) GetPackageSize() *int32 {
	return s.PackageSize
}

func (s *CreateBandwidthResourcePackagesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateBandwidthResourcePackagesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateBandwidthResourcePackagesRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateBandwidthResourcePackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBandwidthResourcePackagesRequest) SetAmount(v int32) *CreateBandwidthResourcePackagesRequest {
	s.Amount = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetAutoPay(v bool) *CreateBandwidthResourcePackagesRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPackageSize(v int32) *CreateBandwidthResourcePackagesRequest {
	s.PackageSize = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPeriod(v int32) *CreateBandwidthResourcePackagesRequest {
	s.Period = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPeriodUnit(v string) *CreateBandwidthResourcePackagesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetPromotionId(v string) *CreateBandwidthResourcePackagesRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) SetRegionId(v string) *CreateBandwidthResourcePackagesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBandwidthResourcePackagesRequest) Validate() error {
	return dara.Validate(s)
}
