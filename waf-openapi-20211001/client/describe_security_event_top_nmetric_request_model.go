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
	// The query filter conditions. Multiple filter conditions are evaluated using a logical AND.
	//
	// This parameter is required.
	Filter *DescribeSecurityEventTopNMetricRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
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
	// Specifies the type of data to return. Different Metric values correspond to different data content. The following Metric values are supported by this API operation:
	//
	// > The definition of "attack request" is described in the API operation description. The following descriptions reference this concept.
	//
	// - real_client_ip: performs aggregation and sorting of the source IP addresses of attack requests in descending order, and returns the top N entries.
	//
	// - http_user_agent: performs aggregation and sorting of the User-Agent values of attack requests in descending order, and returns the top N entries.
	//
	// - matched_host: performs aggregation and sorting of the protected objects hit by attack requests in descending order, and returns the top N entries.
	//
	// - remote_region_id: performs aggregation and sorting of the countries to which the source IP addresses of attack requests belong in descending order, and returns the top N entries.
	//
	// - request_path: performs aggregation and sorting of the URLs (excluding query strings) of attack requests in descending order, and returns the top N entries.
	//
	// - block_defense_scene: performs aggregation and sorting of the final action modules of blocked requests (whose action is not "monitor") in descending order, and returns the top N entries.
	//
	// - defense_scene: performs aggregation and sorting of all protection modules hit by attack requests in descending order, and returns the top N entries.
	//
	// - defense_scene_rule_id: queries the top rule IDs of hit non-monitor rules and the protection modules to which the rules belong. This query returns statistics only for non-monitor mode rules. The returned data format is as follows:<br>
	//
	//  `{ "Attribute": "waf_base", "Value": 140, "Name": "111034" }`
	//
	// - defense_scene_with_rule_id: returns the top N rule IDs ranked by the number of hit requests and the protection modules to which the rules belong, connected by "-". This query does not distinguish between rule actions and includes both monitor rules and block rules. The returned format is as follows:<br>
	//
	//  `{ "Attribute": "",  "Value": 1,  "Name": "120075-waf_base" }`
	//
	// - defense_scene_top_rule_id: queries the top rule hits of a specific protection module. Specify filter conditions in the Conditions field of Filter. For example, to query the top rule hits of the "custom ACL" module, set the Conditions field as follows:<br>
	//
	//    `{ "Key": "defense_scene_map", "OpValue": "contain", "Values": "custom_acl" }`
	//
	// - defense_scene_rule_type: queries the top hit rule types of the web core protection module. Only the web core protection module supports this query because only web core protection has rule child classes. Specify filter conditions in the Conditions field of Filter. The format is as follows:<br>
	//
	// `    { "Key": "defense_scene", "OpValue": "eq", "Values": "waf_base" }`
	//
	// This parameter is required.
	//
	// example:
	//
	// real_client_ip
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
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
	// The list of filter conditions. Each node describes a filter condition.
	Conditions []*DescribeSecurityEventTopNMetricRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The query time range.
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
	// The field name on which the filter operation is performed. This operation supports all fields.
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
	// The end time used for querying data, expressed as a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The query data range cannot exceed the past 30 days. The start time used for querying data, expressed as a UNIX timestamp. Unit: seconds.
	//
	//
	// > The start time must be later than the current time minus 30 days.
	//
	// > -
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
