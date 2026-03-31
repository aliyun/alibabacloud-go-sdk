// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTimeSeriesMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *DescribeSecurityEventTimeSeriesMetricRequestFilter) *DescribeSecurityEventTimeSeriesMetricRequest
	GetFilter() *DescribeSecurityEventTimeSeriesMetricRequestFilter
	SetInstanceId(v string) *DescribeSecurityEventTimeSeriesMetricRequest
	GetInstanceId() *string
	SetMetric(v string) *DescribeSecurityEventTimeSeriesMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeSecurityEventTimeSeriesMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTimeSeriesMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventTimeSeriesMetricRequest struct {
	// The filter conditions for the query. Multiple conditions are evaluated by using a logical AND.
	//
	// This parameter is required.
	Filter *DescribeSecurityEventTimeSeriesMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric whose time series data you want to return. The following metrics are supported:
	//
	// 	- mitigated_requests: The system returns the time series data of requests that are blocked.
	//
	// 	- monitored_requests: The system returns the time series data of requests that match Monitor protection rules.
	//
	// 	- mitigated_requests_group_by_defense_scene: The system returns the number of requests that match each protection module. The returned results are grouped by protection module and can be used to generate time series charts. A request can match multiple protection modules. Therefore, the total number of matched requests is inconsistent with the total number of requests.
	//
	// 	- mitigated_requests_group_by_block_defense_scene: The system returns the number of requests that are blocked by each protection module. The returned results are grouped by protection module and can be used to generate time series charts. A request can be blocked by only one protection module. Therefore, the total number of blocked requests is consistent with the total number of requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// mitigated_requests
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: The Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) GetFilter() *DescribeSecurityEventTimeSeriesMetricRequestFilter {
	return s.Filter
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) SetFilter(v *DescribeSecurityEventTimeSeriesMetricRequestFilter) *DescribeSecurityEventTimeSeriesMetricRequest {
	s.Filter = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) SetInstanceId(v string) *DescribeSecurityEventTimeSeriesMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) SetMetric(v string) *DescribeSecurityEventTimeSeriesMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) SetRegionId(v string) *DescribeSecurityEventTimeSeriesMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTimeSeriesMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventTimeSeriesMetricRequestFilter struct {
	// The filter conditions. Each object describes a filter condition.
	Conditions []*DescribeSecurityEventTimeSeriesMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range for the query.
	//
	// This parameter is required.
	DateRange *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilter) GetConditions() []*DescribeSecurityEventTimeSeriesMetricRequestFilterConditions {
	return s.Conditions
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilter) GetDateRange() *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilter) SetConditions(v []*DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) *DescribeSecurityEventTimeSeriesMetricRequestFilter {
	s.Conditions = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilter) SetDateRange(v *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) *DescribeSecurityEventTimeSeriesMetricRequestFilter {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilter) Validate() error {
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

type DescribeSecurityEventTimeSeriesMetricRequestFilterConditions struct {
	// The field name. This operation supports all fields. For details, see the **Supported field names*	- section below.
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator. For details, see the **Supported operators*	- section below.
	//
	// example:
	//
	// eq
	OpValue *string `json:"OpValue,omitempty" xml:"OpValue,omitempty"`
	// The field content.
	//
	// example:
	//
	// test.waf-top
	Values interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) GetKey() *string {
	return s.Key
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) GetOpValue() *string {
	return s.OpValue
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) GetValues() interface{} {
	return s.Values
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) SetKey(v string) *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions {
	s.Key = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) SetOpValue(v string) *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions {
	s.OpValue = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) SetValues(v interface{}) *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions {
	s.Values = v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange struct {
	// The end of the time range to query. The value is a Unix timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. The value is a Unix timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) SetEndDate(v int64) *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) SetStartDate(v int64) *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventTimeSeriesMetricRequestFilterDateRange) Validate() error {
	return dara.Validate(s)
}
