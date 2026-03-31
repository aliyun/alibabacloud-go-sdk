// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseResourceTemplatesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseResourceTemplatesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseResourceTemplatesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceTemplatesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *DescribeDefenseResourceTemplatesRequest
	GetResourceType() *string
	SetRuleId(v int64) *DescribeDefenseResourceTemplatesRequest
	GetRuleId() *int64
	SetRuleName(v string) *DescribeDefenseResourceTemplatesRequest
	GetRuleName() *string
	SetRuleType(v string) *DescribeDefenseResourceTemplatesRequest
	GetRuleType() *string
	SetTemplateName(v string) *DescribeDefenseResourceTemplatesRequest
	GetTemplateName() *string
}

type DescribeDefenseResourceTemplatesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-9lb36****0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object or protected object group that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxhemicals.cn-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2ax2y5****pi
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protected resource. Valid values:
	//
	// 	- **single**: protected object. This is the default value.
	//
	// 	- **group**: protected object group.
	//
	// example:
	//
	// single
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the protection rule.
	//
	// example:
	//
	// 20111098
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// demoRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the protection rule. Valid values:
	//
	// 	- **defense**: defense rule. This is the default value.
	//
	// 	- **whitelist**: whitelist rule.
	//
	// example:
	//
	// whitelist
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The name of the protection rule template.
	//
	// example:
	//
	// test221
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeDefenseResourceTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceTemplatesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseResourceTemplatesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceTemplatesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDefenseResourceTemplatesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDefenseResourceTemplatesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDefenseResourceTemplatesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeDefenseResourceTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDefenseResourceTemplatesRequest) SetInstanceId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRegionId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResource(v string) *DescribeDefenseResourceTemplatesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceTemplatesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetResourceType(v string) *DescribeDefenseResourceTemplatesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRuleId(v int64) *DescribeDefenseResourceTemplatesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRuleName(v string) *DescribeDefenseResourceTemplatesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetRuleType(v string) *DescribeDefenseResourceTemplatesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) SetTemplateName(v string) *DescribeDefenseResourceTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseResourceTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
