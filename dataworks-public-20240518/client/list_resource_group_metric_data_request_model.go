// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListResourceGroupMetricDataRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListResourceGroupMetricDataRequest
	GetEndTime() *int64
	SetLength(v int32) *ListResourceGroupMetricDataRequest
	GetLength() *int32
	SetMetricName(v string) *ListResourceGroupMetricDataRequest
	GetMetricName() *string
	SetNextToken(v string) *ListResourceGroupMetricDataRequest
	GetNextToken() *string
	SetPeriod(v string) *ListResourceGroupMetricDataRequest
	GetPeriod() *string
	SetResourceGroupId(v string) *ListResourceGroupMetricDataRequest
	GetResourceGroupId() *string
}

type ListResourceGroupMetricDataRequest struct {
	// Start Time
	//
	// Supported format:
	//
	// 	- Unix timestamp, representing the number of milliseconds that have elapsed since January 1, 1970.
	//
	// The interval between BeginTime and EndTime must be 31 days or less.
	//
	// Default: The current time minus 2 hours, expressed as a millisecond Unix timestamp.
	//
	// example:
	//
	// 1593950832000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// End Time
	//
	// Supported format:
	//
	// 	- Unix timestamp, representing the number of milliseconds that have elapsed since January 1, 1970.
	//
	// The interval between BeginTime and EndTime must be 31 days or less.
	//
	// Default: The current time, expressed as a millisecond Unix timestamp.
	//
	// example:
	//
	// 1750176000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 100
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric name. Available metrics include:
	//
	// 	- CUSpec: Maximum CU capacity of the resource group, in CUs.
	//
	// 	- CUUsage: CU usage of the resource group, in CUs.
	//
	// 	- CUUtilization: CU utilization of the resource group, in %.
	//
	// 	- SlotSpec: Maximum concurrency for resource group scheduling, in slots.
	//
	// 	- SlotUsage: Used concurrency for resource group scheduling, in slots.
	//
	// 	- SchedulerCUMaxSpec: Maximum CU quota for data computing, in CUs.
	//
	// 	- SchedulerCUUsage: CU usage for data computing, in CUs.
	//
	// 	- SchedulerCUMinSpec: Minimum guaranteed CUs for data computing, in CUs.
	//
	// 	- DataIntegrationCUMaxSpec: Maximum CU quota for Data Integration, in CUs.
	//
	// 	- DataIntegrationCUUsage: CU usage for Data Integration, in CUs.
	//
	// 	- DataIntegrationCUMinSpec: Minimum guaranteed CUs for Data Integration, in CUs.
	//
	// 	- DataServiceCUMaxSpec: Maximum CU quota for dataservice, in CUs.
	//
	// 	- DataServiceCUUsage: CU usage for DataService Studio, in CUs.
	//
	// 	- DataServiceCUMinSpec: Minimum guaranteed CUs for DataService Studio, in CUs.
	//
	// 	- ServerIdeCUMaxSpec: Maximum CU quota for personal development environment, in CUs.
	//
	// 	- ServerIdeCUUsage: CU usage for personal development environment, in CUs.
	//
	// 	- ServerIdeCUMinSpec: Minimum guaranteed CUs for personal development environment, in CUs.
	//
	// This parameter is required.
	//
	// example:
	//
	// CUSpec
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// FFqBJBxE8I0PE0IUO6K69k7m2FfyWNNc2qQ9ReUkazhz9VA7dWZKlxBcjUwOV0imSM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period for monitoring data.
	//
	// Value: A multiple of 60.
	//
	// Unit: Seconds.
	//
	// Default: 60
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListResourceGroupMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListResourceGroupMetricDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListResourceGroupMetricDataRequest) GetLength() *int32 {
	return s.Length
}

func (s *ListResourceGroupMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *ListResourceGroupMetricDataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupMetricDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *ListResourceGroupMetricDataRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupMetricDataRequest) SetBeginTime(v int64) *ListResourceGroupMetricDataRequest {
	s.BeginTime = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetEndTime(v int64) *ListResourceGroupMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetLength(v int32) *ListResourceGroupMetricDataRequest {
	s.Length = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetMetricName(v string) *ListResourceGroupMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetNextToken(v string) *ListResourceGroupMetricDataRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetPeriod(v string) *ListResourceGroupMetricDataRequest {
	s.Period = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetResourceGroupId(v string) *ListResourceGroupMetricDataRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
