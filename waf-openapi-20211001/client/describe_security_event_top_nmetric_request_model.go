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
	// The filter conditions for the query. A logical AND operator is used between multiple filter conditions.
	//
	// This parameter is required.
	Filter *DescribeSecurityEventTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return. The entries are sorted in descending order. Maximum value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Specifies the content of the returned data. Different metrics correspond to different data content. This operation supports the following metrics:
	//
	// > For the definition of an attack request, see the Description section of this topic. The following descriptions use this definition.
	//
	// - real_client_ip: Aggregates the source IP addresses of attack requests, sorts them in descending order, and returns the top N results.
	//
	// - http_user_agent: Aggregates the User-Agent headers of attack requests, sorts them in descending order, and returns the top N results.
	//
	// - matched_host: Aggregates the protected objects that are hit by attack requests, sorts them in descending order, and returns the top N results.
	//
	// - remote_region_id: Aggregates the countries of origin for the source IP addresses of attack requests, sorts them in descending order, and returns the top N results.
	//
	// - request_path: Aggregates the URLs of attack requests, excluding query strings, sorts them in descending order, and returns the top N results.
	//
	// - block_defense_scene: Aggregates the final protection modules that handled blocked requests, sorts them in descending order, and returns the top N results. Blocked requests are requests whose action is not Monitor.
	//
	// - defense_scene: Aggregates all protection modules that are hit by attack requests, sorts them in descending order, and returns the top N results.
	//
	// - defense_scene_rule_id: Queries the top rule IDs of hit rules that are not in Monitor mode and their corresponding protection modules. This query returns statistics only for rules that are not in Monitor mode. The data is returned in the following format:<br>
	//
	//   `{ "Attribute": "waf_base", "Value": 140, "Name": "111034" }`<br>
	//
	// - defense_scene_with_rule_id: Returns the top N rule IDs based on the number of hits and their corresponding protection modules. The rule ID and module are connected by a hyphen (-). This query includes rules in both Monitor and Block modes. The data is returned in the following format:<br>
	//
	//   `{ "Attribute": "", "Value": 1, "Name": "120075-waf_base" }`<br>
	//
	// - defense_scene_top_rule_id: Queries the top rules that are hit in a specific protection module. Specify a filter condition in the Conditions field of the Filter parameter. For example, to query the top rules that are hit in the custom access control list (ACL) module, set the Conditions field as follows:<br>
	//
	//   `{ "Key": "defense_scene_map", "OpValue": "contain", "Values": "custom_acl" }`<br>
	//
	// - defense_scene_rule_type: Queries the top rule types that are hit in the core web protection module. Only the core web protection module supports this query because it is the only module that has rule subtypes. Specify a filter condition in the Conditions field of the Filter parameter as follows:<br>
	//
	//   `{ "Key": "defense_scene", "OpValue": "eq", "Values": "waf_base" }`<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// real_client_ip
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
	// The ID of the resource group.
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
	// A list of filter conditions. Each object in the array represents a filter condition.
	Conditions []*DescribeSecurityEventTopNMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range to query.
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
	// The field to filter by. This operation supports all fields.
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
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start of the time range to query. The value is a UNIX timestamp. Unit: seconds. You can query data within the last 30 days.
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
