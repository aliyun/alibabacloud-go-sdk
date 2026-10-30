// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodeRenewPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIds(v []*string) *DescribeCloudPhoneNodeRenewPriceRequest
	GetNodeIds() []*string
	SetPeriod(v int32) *DescribeCloudPhoneNodeRenewPriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeCloudPhoneNodeRenewPriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeCloudPhoneNodeRenewPriceRequest
	GetPromotionId() *string
}

type DescribeCloudPhoneNodeRenewPriceRequest struct {
	// The list of cloud phone matrix IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The subscription period. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
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
	// The ID of the promotion.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) SetNodeIds(v []*string) *DescribeCloudPhoneNodeRenewPriceRequest {
	s.NodeIds = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) SetPeriod(v int32) *DescribeCloudPhoneNodeRenewPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) SetPeriodUnit(v string) *DescribeCloudPhoneNodeRenewPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) SetPromotionId(v string) *DescribeCloudPhoneNodeRenewPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceRequest) Validate() error {
	return dara.Validate(s)
}
