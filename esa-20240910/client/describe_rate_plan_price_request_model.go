// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *DescribeRatePlanPriceRequest
	GetAmount() *int32
	SetPeriod(v int32) *DescribeRatePlanPriceRequest
	GetPeriod() *int32
	SetPlanName(v string) *DescribeRatePlanPriceRequest
	GetPlanName() *string
}

type DescribeRatePlanPriceRequest struct {
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 1
	Period   *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
}

func (s DescribeRatePlanPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribeRatePlanPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeRatePlanPriceRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeRatePlanPriceRequest) SetAmount(v int32) *DescribeRatePlanPriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeRatePlanPriceRequest) SetPeriod(v int32) *DescribeRatePlanPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeRatePlanPriceRequest) SetPlanName(v string) *DescribeRatePlanPriceRequest {
	s.PlanName = &v
	return s
}

func (s *DescribeRatePlanPriceRequest) Validate() error {
	return dara.Validate(s)
}
