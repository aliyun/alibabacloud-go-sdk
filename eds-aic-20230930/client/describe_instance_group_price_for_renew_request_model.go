// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceForRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupIds(v []*string) *DescribeInstanceGroupPriceForRenewRequest
	GetInstanceGroupIds() []*string
	SetPeriod(v int32) *DescribeInstanceGroupPriceForRenewRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeInstanceGroupPriceForRenewRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeInstanceGroupPriceForRenewRequest
	GetPromotionId() *string
}

type DescribeInstanceGroupPriceForRenewRequest struct {
	// The list of instance group IDs.
	//
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The subscription period of the resource. The unit is specified by PeriodUnit. Valid values: 1 month, 2 months, 3 months, 6 months, and 1 year.
	//
	// example:
	//
	// 3
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// - **Month**: month
	//
	// - **Year**: year
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

func (s DescribeInstanceGroupPriceForRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *DescribeInstanceGroupPriceForRenewRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeInstanceGroupPriceForRenewRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeInstanceGroupPriceForRenewRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeInstanceGroupPriceForRenewRequest) SetInstanceGroupIds(v []*string) *DescribeInstanceGroupPriceForRenewRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewRequest) SetPeriod(v int32) *DescribeInstanceGroupPriceForRenewRequest {
	s.Period = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewRequest) SetPeriodUnit(v string) *DescribeInstanceGroupPriceForRenewRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewRequest) SetPromotionId(v string) *DescribeInstanceGroupPriceForRenewRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewRequest) Validate() error {
	return dara.Validate(s)
}
