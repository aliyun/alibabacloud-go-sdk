// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductionStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDeductionStatisticRequest
	GetEndTime() *int64
	SetInstanceIds(v []*string) *DescribeDeductionStatisticRequest
	GetInstanceIds() []*string
	SetPeriods(v []*DescribeDeductionStatisticRequestPeriods) *DescribeDeductionStatisticRequest
	GetPeriods() []*DescribeDeductionStatisticRequestPeriods
	SetResourceTypes(v []*string) *DescribeDeductionStatisticRequest
	GetResourceTypes() []*string
	SetStartTime(v int64) *DescribeDeductionStatisticRequest
	GetStartTime() *int64
}

type DescribeDeductionStatisticRequest struct {
	EndTime     *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	Periods       []*DescribeDeductionStatisticRequestPeriods `json:"Periods,omitempty" xml:"Periods,omitempty" type:"Repeated"`
	ResourceTypes []*string                                   `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	StartTime     *int64                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDeductionStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDeductionStatisticRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDeductionStatisticRequest) GetPeriods() []*DescribeDeductionStatisticRequestPeriods {
	return s.Periods
}

func (s *DescribeDeductionStatisticRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeDeductionStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDeductionStatisticRequest) SetEndTime(v int64) *DescribeDeductionStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDeductionStatisticRequest) SetInstanceIds(v []*string) *DescribeDeductionStatisticRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDeductionStatisticRequest) SetPeriods(v []*DescribeDeductionStatisticRequestPeriods) *DescribeDeductionStatisticRequest {
	s.Periods = v
	return s
}

func (s *DescribeDeductionStatisticRequest) SetResourceTypes(v []*string) *DescribeDeductionStatisticRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeDeductionStatisticRequest) SetStartTime(v int64) *DescribeDeductionStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDeductionStatisticRequest) Validate() error {
	if s.Periods != nil {
		for _, item := range s.Periods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeductionStatisticRequestPeriods struct {
	// example:
	//
	// 2024-01-01
	BaseTime   *string `json:"BaseTime,omitempty" xml:"BaseTime,omitempty"`
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s DescribeDeductionStatisticRequestPeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticRequestPeriods) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticRequestPeriods) GetBaseTime() *string {
	return s.BaseTime
}

func (s *DescribeDeductionStatisticRequestPeriods) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeDeductionStatisticRequestPeriods) SetBaseTime(v string) *DescribeDeductionStatisticRequestPeriods {
	s.BaseTime = &v
	return s
}

func (s *DescribeDeductionStatisticRequestPeriods) SetPeriodUnit(v string) *DescribeDeductionStatisticRequestPeriods {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeDeductionStatisticRequestPeriods) Validate() error {
	return dara.Validate(s)
}
