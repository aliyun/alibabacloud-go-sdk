// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *DescribeInstanceGroupPriceRequest
	GetAmount() *int32
	SetBandwidth(v int32) *DescribeInstanceGroupPriceRequest
	GetBandwidth() *int32
	SetBizRegionId(v string) *DescribeInstanceGroupPriceRequest
	GetBizRegionId() *string
	SetImageId(v string) *DescribeInstanceGroupPriceRequest
	GetImageId() *string
	SetInstanceGroupSpec(v string) *DescribeInstanceGroupPriceRequest
	GetInstanceGroupSpec() *string
	SetInstanceVersion(v string) *DescribeInstanceGroupPriceRequest
	GetInstanceVersion() *string
	SetNumberOfInstances(v int32) *DescribeInstanceGroupPriceRequest
	GetNumberOfInstances() *int32
	SetOrderType(v string) *DescribeInstanceGroupPriceRequest
	GetOrderType() *string
	SetPeriod(v int32) *DescribeInstanceGroupPriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeInstanceGroupPriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeInstanceGroupPriceRequest
	GetPromotionId() *string
	SetSaleMode(v string) *DescribeInstanceGroupPriceRequest
	GetSaleMode() *string
}

type DescribeInstanceGroupPriceRequest struct {
	// The number of instance groups.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// np-shdbbkdubb****
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-0aae4rgm5umzo****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance group specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// acp.basic.small
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// The instance version. You do not need to specify this parameter. The system automatically matches the value.
	//
	// example:
	//
	// basic
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The number of instances in the instance group.
	//
	// example:
	//
	// 10
	NumberOfInstances *int32 `json:"NumberOfInstances,omitempty" xml:"NumberOfInstances,omitempty"`
	// The order type.
	//
	// example:
	//
	// CREATE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The subscription period. The unit is specified by PeriodUnit. Default value: 1.
	//
	// example:
	//
	// 2
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotion.
	//
	// example:
	//
	// 88436689****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The sales mode. You do not need to specify this parameter. The system automatically matches the value.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
}

func (s DescribeInstanceGroupPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeInstanceGroupPriceRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeInstanceGroupPriceRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeInstanceGroupPriceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstanceGroupPriceRequest) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *DescribeInstanceGroupPriceRequest) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *DescribeInstanceGroupPriceRequest) GetNumberOfInstances() *int32 {
	return s.NumberOfInstances
}

func (s *DescribeInstanceGroupPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeInstanceGroupPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeInstanceGroupPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeInstanceGroupPriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeInstanceGroupPriceRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeInstanceGroupPriceRequest) SetAmount(v int32) *DescribeInstanceGroupPriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetBandwidth(v int32) *DescribeInstanceGroupPriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetBizRegionId(v string) *DescribeInstanceGroupPriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetImageId(v string) *DescribeInstanceGroupPriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetInstanceGroupSpec(v string) *DescribeInstanceGroupPriceRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetInstanceVersion(v string) *DescribeInstanceGroupPriceRequest {
	s.InstanceVersion = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetNumberOfInstances(v int32) *DescribeInstanceGroupPriceRequest {
	s.NumberOfInstances = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetOrderType(v string) *DescribeInstanceGroupPriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetPeriod(v int32) *DescribeInstanceGroupPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetPeriodUnit(v string) *DescribeInstanceGroupPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetPromotionId(v string) *DescribeInstanceGroupPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) SetSaleMode(v string) *DescribeInstanceGroupPriceRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeInstanceGroupPriceRequest) Validate() error {
	return dara.Validate(s)
}
