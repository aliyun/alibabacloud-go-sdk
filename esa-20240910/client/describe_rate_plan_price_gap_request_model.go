// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanPriceGapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRatePlanPriceGapRequest
	GetInstanceId() *string
	SetOrderType(v string) *DescribeRatePlanPriceGapRequest
	GetOrderType() *string
	SetTargetPlanCode(v string) *DescribeRatePlanPriceGapRequest
	GetTargetPlanCode() *string
	SetTargetPlanName(v string) *DescribeRatePlanPriceGapRequest
	GetTargetPlanName() *string
}

type DescribeRatePlanPriceGapRequest struct {
	// The target plan name. You can obtain this value from the [DescribeRatePlanPrice](~~DescribeRatePlanPrice~~) operation.
	//
	// example:
	//
	// entranceplan
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The specification change type. Valid values:
	//
	// - DOWNGRADE: downgrade.
	//
	// - UPGRADE: upgrade.
	//
	// example:
	//
	// UPGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The specification change type. Valid values:
	//
	// - DOWNGRADE: downgrade.
	//
	// - UPGRADE: upgrade.
	//
	// example:
	//
	// UPGRADE
	TargetPlanCode *string `json:"TargetPlanCode,omitempty" xml:"TargetPlanCode,omitempty"`
	// The name of the target plan for the specification change.
	//
	// example:
	//
	// entranceplan
	TargetPlanName *string `json:"TargetPlanName,omitempty" xml:"TargetPlanName,omitempty"`
}

func (s DescribeRatePlanPriceGapRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceGapRequest) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceGapRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRatePlanPriceGapRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeRatePlanPriceGapRequest) GetTargetPlanCode() *string {
	return s.TargetPlanCode
}

func (s *DescribeRatePlanPriceGapRequest) GetTargetPlanName() *string {
	return s.TargetPlanName
}

func (s *DescribeRatePlanPriceGapRequest) SetInstanceId(v string) *DescribeRatePlanPriceGapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRatePlanPriceGapRequest) SetOrderType(v string) *DescribeRatePlanPriceGapRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeRatePlanPriceGapRequest) SetTargetPlanCode(v string) *DescribeRatePlanPriceGapRequest {
	s.TargetPlanCode = &v
	return s
}

func (s *DescribeRatePlanPriceGapRequest) SetTargetPlanName(v string) *DescribeRatePlanPriceGapRequest {
	s.TargetPlanName = &v
	return s
}

func (s *DescribeRatePlanPriceGapRequest) Validate() error {
	return dara.Validate(s)
}
