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
	// An array of filter conditions. Multiple filter parameters use AND logic.
	//
	// This parameter is required.
	Filter *DescribeNetworkFlowTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The Web Application Firewall (WAF) instance ID.
	//
	// >  Call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to retrieve the WAF instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Returns up to 10 data entries, sorted in descending order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Specifies the data type to be returned. Valid values:
	//
	// 	- real_client_ip: The top N requests, sorted in descending order by source IP address, aggregated from all the current user\\"s WAF requests.
	//
	// 	- request_path: The top N requests, sorted in descending order by user-agent, aggregated from all the current user\\"s WAF requests.
	//
	// 	- request_path: The top N requests, sorted in descending order by request URL, aggregated from all the current user\\"s WAF requests.
	//
	// 	- matched_host_by_total_requests: The top N protected objects and their request counts for the current user.
	//
	// 	- matched_host_by_qps: The top N protected objects and their queries per second (QPS) values.
	//
	// 	- matched_host_by_status: When using it, you must specify status in the Conditions field of the Filter parameter. If the HTTP response code returned by WAF matches the status specified in the Conditions, then the top N data is returned, sorted in descending order by protected objects. The format for specifying the status is as follows:\\
	//
	//     {"Key":"status","OpValue":"eq","Values":"200"}
	//
	// 	- matched_host_by_upstream_status: When using it, you must specify upstream_status in the Conditions field of the Filter parameter. If the HTTP response code returned by the origin server matches the upstream_status specified, the top N data is returned, sorted in descending order by protected objects. The format for specifying the upstream_status is as follows:\\
	//
	//     {"Key":"upstream_status","OpValue":"eq","Values":"200"}
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
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
	// The resource group ID.
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
	// The list of filter conditions. Each node describes a filter condition.
	Conditions []*DescribeNetworkFlowTopNMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// Specifies the date range for the query.
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
	// The filter fields. Valid values:
	//
	// 	- matched_host
	//
	// 	- cluster
	//
	// For details, see the **Filter fields (Key)*	- section below.
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter operator.
	//
	// For details, see the **Filter operators (OpValue)*	- section below.
	//
	// example:
	//
	// eq
	OpValue *string `json:"OpValue,omitempty" xml:"OpValue,omitempty"`
	// The filter content.
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
	// End time of the query range (Unix timestamp, seconds).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time of the query range (Unix timestamp, seconds).
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
