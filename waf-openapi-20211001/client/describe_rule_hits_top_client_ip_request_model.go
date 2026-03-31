// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopClientIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopClientIpRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopClientIpRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRuleHitsTopClientIpRequest
	GetRegionId() *string
	SetResource(v string) *DescribeRuleHitsTopClientIpRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopClientIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeRuleHitsTopClientIpRequest
	GetRuleType() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopClientIpRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopClientIpRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
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
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
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

func (s DescribeRuleHitsTopClientIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopClientIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopClientIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopClientIpRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopClientIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopClientIpRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleHitsTopClientIpRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopClientIpRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetInstanceId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetRegionId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetResource(v string) *DescribeRuleHitsTopClientIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetRuleType(v string) *DescribeRuleHitsTopClientIpRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) Validate() error {
	return dara.Validate(s)
}
