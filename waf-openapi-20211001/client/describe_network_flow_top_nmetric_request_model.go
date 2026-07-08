// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTopNMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *DescribeNetworkFlowTopNMetricRequestFilter) *DescribeNetworkFlowTopNMetricRequest
	GetFilter() *DescribeNetworkFlowTopNMetricRequestFilter
	SetInstanceId(v string) *DescribeNetworkFlowTopNMetricRequest
	GetInstanceId() *string
	SetLimit(v int64) *DescribeNetworkFlowTopNMetricRequest
	GetLimit() *int64
	SetMetric(v string) *DescribeNetworkFlowTopNMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeNetworkFlowTopNMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTopNMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeNetworkFlowTopNMetricRequest struct {
	// The filter conditions for the query. If you specify multiple filter conditions, all conditions must be met (logical AND).
	//
	// This parameter is required.
	Filter *DescribeNetworkFlowTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. Results are sorted in descending order. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The metric by which to query and rank data. Valid values:
	//
	// - real_client_ip: Returns the top N source IP addresses of requests that are sent to WAF. The data is aggregated by source IP address and sorted in descending order.
	//
	// - http_user_agent: Returns the top N user agents of requests that are sent to WAF. The data is aggregated by user agent and sorted in descending order.
	//
	// - request_path: Returns the top N request URLs. The data is aggregated by URL and sorted in descending order.
	//
	// - matched_host_by_total_requests: Returns the top N protected objects by total number of requests received.
	//
	// - matched_host_by_qps: Returns the top N protected objects by queries per second (QPS).
	//
	// - matched_host_by_status: Returns the top N protected objects based on a specific WAF-returned HTTP status code. The data is aggregated by protected object and sorted in descending order. If you set Metric to this value, you must specify the status field in the Conditions parameter of Filter. The format is as follows:<br> {"Key":"status","OpValue":"eq","Values":"200"}<br>
	//
	// - matched_host_by_upstream_status: Returns the top N protected objects based on a specific origin server-returned HTTP status code. The data is aggregated by protected object and sorted in descending order. If you set Metric to this value, you must specify the upstream_status field in the Conditions parameter of Filter. The format is as follows:<br> {"Key":"upstream_status","OpValue":"eq","Values":"200"}<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region in which the WAF instance resides. Valid values:
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

func (s DescribeNetworkFlowTopNMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetFilter() *DescribeNetworkFlowTopNMetricRequestFilter {
	return s.Filter
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkFlowTopNMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetFilter(v *DescribeNetworkFlowTopNMetricRequestFilter) *DescribeNetworkFlowTopNMetricRequest {
	s.Filter = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetInstanceId(v string) *DescribeNetworkFlowTopNMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetLimit(v int64) *DescribeNetworkFlowTopNMetricRequest {
	s.Limit = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetMetric(v string) *DescribeNetworkFlowTopNMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetRegionId(v string) *DescribeNetworkFlowTopNMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeNetworkFlowTopNMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkFlowTopNMetricRequestFilter struct {
	// The list of filter conditions.
	Conditions []*DescribeNetworkFlowTopNMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range to query.
	//
	// This parameter is required.
	DateRange *DescribeNetworkFlowTopNMetricRequestFilterDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
}

func (s DescribeNetworkFlowTopNMetricRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricRequestFilter) GetConditions() []*DescribeNetworkFlowTopNMetricRequestFilterConditions {
	return s.Conditions
}

func (s *DescribeNetworkFlowTopNMetricRequestFilter) GetDateRange() *DescribeNetworkFlowTopNMetricRequestFilterDateRange {
	return s.DateRange
}

func (s *DescribeNetworkFlowTopNMetricRequestFilter) SetConditions(v []*DescribeNetworkFlowTopNMetricRequestFilterConditions) *DescribeNetworkFlowTopNMetricRequestFilter {
	s.Conditions = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilter) SetDateRange(v *DescribeNetworkFlowTopNMetricRequestFilterDateRange) *DescribeNetworkFlowTopNMetricRequestFilter {
	s.DateRange = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilter) Validate() error {
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

type DescribeNetworkFlowTopNMetricRequestFilterConditions struct {
	// The field name to use for filtering. Valid values:
	//
	// - matched_host
	//
	// - cluster
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator that is used for the filter condition. For more information about supported operators, see the **Additional information about request parameters*	- section.
	//
	// example:
	//
	// eq
	OpValue *string `json:"OpValue,omitempty" xml:"OpValue,omitempty"`
	// The value to use for the filter condition. The value format depends on the Key and OpValue that you specify.
	//
	// example:
	//
	// test.waf-top
	Values interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) GetOpValue() *string {
	return s.OpValue
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) GetValues() interface{} {
	return s.Values
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) SetKey(v string) *DescribeNetworkFlowTopNMetricRequestFilterConditions {
	s.Key = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) SetOpValue(v string) *DescribeNetworkFlowTopNMetricRequestFilterConditions {
	s.OpValue = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) SetValues(v interface{}) *DescribeNetworkFlowTopNMetricRequestFilterConditions {
	s.Values = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkFlowTopNMetricRequestFilterDateRange struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds. The time range cannot exceed the last 30 days.
	//
	// > The start time must be later than 30 days before the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricRequestFilterDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricRequestFilterDateRange) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterDateRange) SetEndDate(v int64) *DescribeNetworkFlowTopNMetricRequestFilterDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterDateRange) SetStartDate(v int64) *DescribeNetworkFlowTopNMetricRequestFilterDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricRequestFilterDateRange) Validate() error {
	return dara.Validate(s)
}
