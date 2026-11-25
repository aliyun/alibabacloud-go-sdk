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
	// The query filter conditions. Multiple filter conditions are connected by the logical AND operator.
	//
	// This parameter is required.
	Filter *DescribeNetworkFlowTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The WAF instance ID.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of data entries to return after the statistics are sorted in descending order. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The metric used to distinguish the returned data content. Different metrics correspond to different data content. The metrics supported by this operation are as follows:
	//
	// - real_client_ip: aggregates and sorts all requests of the current user that pass through WAF by the source IP address of the request in descending order, and returns the top N data entries.
	//
	// - http_user_agent: aggregates and sorts all requests of the current user that pass through WAF by the User-Agent of the request in descending order, and returns the top N data entries.
	//
	// - request_path: aggregates and sorts all requests of the current user that pass through WAF by the URL of the request in descending order, and returns the top N data entries.
	//
	// - matched_host_by_total_requests: returns the top N protected objects with the largest number of requests for the current user and the corresponding number of requests.
	//
	// - matched_host_by_qps: returns the top N protected objects with the highest QPS for the current user and the corresponding QPS data.
	//
	// - matched_host_by_status: when this metric is used, you must specify status in the Conditions field of Filter. WAF aggregates and sorts the data whose HTTP response code returned by WAF is equal to the status specified in Conditions by protected object, and returns the top N data entries. The format for specifying status is as follows:<br>
	//
	// {"Key":"status","OpValue":"eq","Values":"200"}
	//
	// - matched_host_by_upstream_status: when this metric is used, you must specify upstream_status in the Conditions field of Filter. WAF aggregates and sorts the data whose HTTP response code returned by the origin server is equal to the upstream_status specified in Conditions by protected object, and returns the top N data entries. The format for specifying upstream_status is as follows:<br>
	//
	// {"Key":"upstream_status","OpValue":"eq","Values":"200"}
	//
	// This parameter is required.
	//
	// example:
	//
	// matched_host_by_upstream_status
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
	// The query time range.
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
	// The name of the field on which the filter operation is performed. The Key parameter of the current operation supports only the following two fields:
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
	// The filter values.
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
	// The end time of the query. The end time is represented by a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query. The data query range cannot exceed the past 30 days. The start time is represented by a UNIX timestamp in seconds.
	//
	// >
	//
	// > The start time must be later than the current time minus 30 days.
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
