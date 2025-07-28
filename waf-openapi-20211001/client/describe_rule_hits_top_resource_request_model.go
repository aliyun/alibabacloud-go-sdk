// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopResourceRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopResourceRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRuleHitsTopResourceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeRuleHitsTopResourceRequest
	GetRuleType() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopResourceRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopResourceRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386340
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
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
	// The type of rules that are triggered by the protected object. By default, this parameter is not specified and all types of rules are queried.
	//
	// 	- **blacklist:*	- IP address blacklist rules.
	//
	// 	- **custom:*	- custom rules.
	//
	// 	- **antiscan:*	- scan protection rules.
	//
	// 	- **cc_system:*	- HTTP flood protection rules.
	//
	// 	- **region_block:*	- region blacklist rules.
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

func (s DescribeRuleHitsTopResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopResourceRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleHitsTopResourceRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopResourceRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetInstanceId(v string) *DescribeRuleHitsTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetRegionId(v string) *DescribeRuleHitsTopResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetRuleType(v string) *DescribeRuleHitsTopResourceRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) Validate() error {
	return dara.Validate(s)
}
