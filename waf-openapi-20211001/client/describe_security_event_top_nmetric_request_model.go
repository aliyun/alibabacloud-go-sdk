// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTopNMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *DescribeSecurityEventTopNMetricRequestFilter) *DescribeSecurityEventTopNMetricRequest
	GetFilter() *DescribeSecurityEventTopNMetricRequestFilter
	SetInstanceId(v string) *DescribeSecurityEventTopNMetricRequest
	GetInstanceId() *string
	SetLimit(v int64) *DescribeSecurityEventTopNMetricRequest
	GetLimit() *int64
	SetMetric(v string) *DescribeSecurityEventTopNMetricRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeSecurityEventTopNMetricRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTopNMetricRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventTopNMetricRequest struct {
	// The filter conditions for the query. Multiple conditions are evaluated by using a logical AND.
	//
	// This parameter is required.
	Filter *DescribeSecurityEventTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of data entries that can be returned. Data entries are sorted in descending order before they are returned. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The metric whose top N data entries you want to return. The following metrics are supported:
	//
	// >  For more information about attack requests, see the "Operation description" section of this topic.
	//
	// 	- real_client_ip: The system aggregates the source IP addresses of attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries.
	//
	// 	- http_user_agent: The system aggregates the User-Agent header field of attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries.
	//
	// 	- matched_host: The system aggregates the protected objects that are matched by attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries.
	//
	// 	- remote_region_id: The system aggregates the countries to which the source IP addresses of attack requests belong to collect statistics, sorts the statistical results in descending order, and returns top N data entries.
	//
	// 	- request_path: The system aggregates the URLs of attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries. The URLs exclude query strings.
	//
	// 	- block_defense_scene: The system aggregates the protection modules that block attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries. The requests match protection rules whose actions are not set to Monitor.
	//
	// 	- defense_scene: The system aggregates the protection modules that are matched by attack requests to collect statistics, sorts the statistical results in descending order, and returns top N data entries.
	//
	// 	- defense_scene_rule_id: The system returns the IDs of top N protection rules that are matched by attack requests and also the related protection modules. Only protection rules whose actions are not set to Monitor are counted. The system returns the value in the following format:\\
	//
	//     `{ "Attribute": "waf_base", "Value": 140, "Name": "111034" }`
	//
	// 	- defense_scene_with_rule_id: The system returns the IDs of top N protection rules that are matched by attack requests and also the related protection modules. The IDs and protection modules are connected by using hyphens (-). Protection rules whose actions are set to Monitor and Block are counted. The system returns the value in the following format:\\
	//
	//     `{ "Attribute": "", "Value": 1, "Name": "120075-waf_base" }`
	//
	// 	- defense_scene_top_rule_id: The system returns top N matched protection rules of a specific protection module. You can specify Conditions in Filter to configure filter conditions. For example, you can use the following condition to query top N matched protection rules of the custom rule module:\\
	//
	//     `{ "Key": "defense_scene_map", "OpValue": "contain", "Values": "custom_acl" }`
	//
	// 	- defense_scene_rule_type: The system returns top N matched protection rules of the core web protection module. This metric is supported only by the core web protection module because only this module supports subtypes of protection rules. You must specify Conditions in Filter to configure filter conditions. Example:\\
	//
	//     `{ "Key": "defense_scene", "OpValue": "eq", "Values": "waf_base" }`
	//
	// This parameter is required.
	//
	// example:
	//
	// real_client_ip
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

func (s DescribeSecurityEventTopNMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricRequest) GetFilter() *DescribeSecurityEventTopNMetricRequestFilter {
	return s.Filter
}

func (s *DescribeSecurityEventTopNMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventTopNMetricRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeSecurityEventTopNMetricRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeSecurityEventTopNMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventTopNMetricRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventTopNMetricRequest) SetFilter(v *DescribeSecurityEventTopNMetricRequestFilter) *DescribeSecurityEventTopNMetricRequest {
	s.Filter = v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) SetInstanceId(v string) *DescribeSecurityEventTopNMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) SetLimit(v int64) *DescribeSecurityEventTopNMetricRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) SetMetric(v string) *DescribeSecurityEventTopNMetricRequest {
	s.Metric = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) SetRegionId(v string) *DescribeSecurityEventTopNMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTopNMetricRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventTopNMetricRequestFilter struct {
	// The filter conditions. Each object describes a filter condition.
	Conditions []*DescribeSecurityEventTopNMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range for the query.
	//
	// This parameter is required.
	DateRange *DescribeSecurityEventTopNMetricRequestFilterDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventTopNMetricRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricRequestFilter) GetConditions() []*DescribeSecurityEventTopNMetricRequestFilterConditions {
	return s.Conditions
}

func (s *DescribeSecurityEventTopNMetricRequestFilter) GetDateRange() *DescribeSecurityEventTopNMetricRequestFilterDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventTopNMetricRequestFilter) SetConditions(v []*DescribeSecurityEventTopNMetricRequestFilterConditions) *DescribeSecurityEventTopNMetricRequestFilter {
	s.Conditions = v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilter) SetDateRange(v *DescribeSecurityEventTopNMetricRequestFilterDateRange) *DescribeSecurityEventTopNMetricRequestFilter {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilter) Validate() error {
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

type DescribeSecurityEventTopNMetricRequestFilterConditions struct {
	// The field name. This operation supports all fields. For more information, see the **Supported field names*	- section below.
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator. For more information, see the **Supported operators*	- section below.
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

func (s DescribeSecurityEventTopNMetricRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) GetKey() *string {
	return s.Key
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) GetOpValue() *string {
	return s.OpValue
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) GetValues() interface{} {
	return s.Values
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) SetKey(v string) *DescribeSecurityEventTopNMetricRequestFilterConditions {
	s.Key = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) SetOpValue(v string) *DescribeSecurityEventTopNMetricRequestFilterConditions {
	s.OpValue = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) SetValues(v interface{}) *DescribeSecurityEventTopNMetricRequestFilterConditions {
	s.Values = v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventTopNMetricRequestFilterDateRange struct {
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

func (s DescribeSecurityEventTopNMetricRequestFilterDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricRequestFilterDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricRequestFilterDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventTopNMetricRequestFilterDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventTopNMetricRequestFilterDateRange) SetEndDate(v int64) *DescribeSecurityEventTopNMetricRequestFilterDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilterDateRange) SetStartDate(v int64) *DescribeSecurityEventTopNMetricRequestFilterDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricRequestFilterDateRange) Validate() error {
	return dara.Validate(s)
}
