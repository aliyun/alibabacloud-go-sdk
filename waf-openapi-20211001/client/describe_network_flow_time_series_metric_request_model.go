// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTimeSeriesMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *DescribeNetworkFlowTimeSeriesMetricRequestFilter) *DescribeNetworkFlowTimeSeriesMetricRequest
	GetFilter() *DescribeNetworkFlowTimeSeriesMetricRequestFilter
	SetInstanceId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeNetworkFlowTimeSeriesMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeNetworkFlowTimeSeriesMetricRequest struct {
	// The filter conditions for the query. Multiple filter conditions are combined with a logical AND.
	//
	// This parameter is required.
	Filter *DescribeNetworkFlowTimeSeriesMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the type of data to return. Different values for Metric correspond to different data. This API operation supports the following values:
	//
	// - qps: The number of requests that WAF processes per second. A queries per second (QPS) value is calculated every 10 seconds. The peak QPS value within the specified time granularity is returned.
	//
	// - total_requests: The total number of requests processed by WAF.
	//
	// - top5_status: The top five response status codes that WAF returns to the client, and the corresponding time series statistics.
	//
	// - top 5_upstream_status: The top five response status codes that the origin server returns to the client, and the corresponding time series statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// total_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) GetFilter() *DescribeNetworkFlowTimeSeriesMetricRequestFilter {
	return s.Filter
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) SetFilter(v *DescribeNetworkFlowTimeSeriesMetricRequestFilter) *DescribeNetworkFlowTimeSeriesMetricRequest {
	s.Filter = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) SetInstanceId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) SetMetric(v string) *DescribeNetworkFlowTimeSeriesMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) SetRegionId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTimeSeriesMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkFlowTimeSeriesMetricRequestFilter struct {
	// A list of filter conditions. Each item in the list describes a filter condition.
	Conditions []*DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range to query.
	//
	// This parameter is required.
	DateRange *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilter) GetConditions() []*DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions {
	return s.Conditions
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilter) GetDateRange() *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange {
	return s.DateRange
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilter) SetConditions(v []*DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) *DescribeNetworkFlowTimeSeriesMetricRequestFilter {
	s.Conditions = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilter) SetDateRange(v *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) *DescribeNetworkFlowTimeSeriesMetricRequestFilter {
	s.DateRange = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilter) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DateRange != nil {
		if err := s.DateRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions struct {
	// The field to use for the filter operation. This API operation supports the following values for Key:
	//
	// - matched_host
	//
	// - cluster
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator.
	//
	// example:
	//
	// eq
	OpValue *string `json:"OpValue,omitempty" xml:"OpValue,omitempty"`
	// The filter value.
	//
	// example:
	//
	// test.waf-top
	Values interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) GetOpValue() *string {
	return s.OpValue
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) GetValues() interface{} {
	return s.Values
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) SetKey(v string) *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions {
	s.Key = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) SetOpValue(v string) *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions {
	s.OpValue = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) SetValues(v interface{}) *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions {
	s.Values = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The query period cannot exceed the last 30 days. The start of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// > ## The start time must be within the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) SetEndDate(v int64) *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) SetStartDate(v int64) *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricRequestFilterDateRange) Validate() error {
	return dara.Validate(s)
}
