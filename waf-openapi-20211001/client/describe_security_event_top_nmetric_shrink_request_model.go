// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventTopNMetricShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *DescribeSecurityEventTopNMetricShrinkRequest
	GetFilterShrink() *string
	SetInstanceId(v string) *DescribeSecurityEventTopNMetricShrinkRequest
	GetInstanceId() *string
	SetLimit(v int64) *DescribeSecurityEventTopNMetricShrinkRequest
	GetLimit() *int64
	SetMetric(v string) *DescribeSecurityEventTopNMetricShrinkRequest
	GetMetric() *string
	SetRegionId(v string) *DescribeSecurityEventTopNMetricShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTopNMetricShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventTopNMetricShrinkRequest struct {
	// The query filter conditions. Multiple filter conditions are evaluated using a logical AND.
	//
	// This parameter is required.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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

func (s DescribeSecurityEventTopNMetricShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventTopNMetricShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetMetric() *string {
	return s.Metric
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetFilterShrink(v string) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetInstanceId(v string) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetLimit(v int64) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.Limit = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetMetric(v string) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.Metric = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetRegionId(v string) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventTopNMetricShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventTopNMetricShrinkRequest) Validate() error {
	return dara.Validate(s)
}
