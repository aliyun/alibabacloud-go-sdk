// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopRuleIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopRuleIdRequest
	GetInstanceId() *string
	SetIsGroupResource(v string) *DescribeRuleHitsTopRuleIdRequest
	GetIsGroupResource() *string
	SetRegionId(v string) *DescribeRuleHitsTopRuleIdRequest
	GetRegionId() *string
	SetResource(v string) *DescribeRuleHitsTopRuleIdRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopRuleIdRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeRuleHitsTopRuleIdRequest
	GetRuleType() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopRuleIdRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
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
	// Specifies whether to aggregate rule hits by protected object.
	//
	// - **true (default)**: Aggregates rule hits by protected object. In this case, rule hits are counted based on individual rules only, and the Resource value in the response is empty.
	//
	// - **false**: Does not aggregate rule hits by protected object. In this case, rule hits are counted based on both protected objects and rules.
	//
	// example:
	//
	// true
	IsGroupResource *string `json:"IsGroupResource,omitempty" xml:"IsGroupResource,omitempty"`
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
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of rule hit by the protected object. By default, this parameter is not specified, which indicates that all rules are displayed.
	//
	// - **blacklist**: IP blacklist.
	//
	// - **custom**: Custom rule.
	//
	// - **antiscan**: Scan protection rule.
	//
	// - **cc_system**: HTTP flood mitigation rule.
	//
	// - **region_block**: Location blacklist.
	//
	// example:
	//
	// blacklist
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetIsGroupResource() *string {
	return s.IsGroupResource
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleHitsTopRuleIdRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetInstanceId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetIsGroupResource(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.IsGroupResource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetRegionId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetResource(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetRuleType(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) Validate() error {
	return dara.Validate(s)
}
