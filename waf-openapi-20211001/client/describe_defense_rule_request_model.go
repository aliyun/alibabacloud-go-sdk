// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseType(v string) *DescribeDefenseRuleRequest
	GetDefenseType() *string
	SetInstanceId(v string) *DescribeDefenseRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseRuleRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseRuleRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DescribeDefenseRuleRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *DescribeDefenseRuleRequest
	GetTemplateId() *int64
}

type DescribeDefenseRuleRequest struct {
	// The type of the protection rule. Valid values:
	//
	// - **template*	- (default): a protection rule template.
	//
	// - **resource**: a rule for a protected object.
	//
	// - **global**: a global rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The protected object to which the protection rule applies.
	//
	// > This parameter is required only when **DefenseType*	- is set to **resource**.
	//
	// example:
	//
	// rencs***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20026192
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the protection rule template.
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
	//
	// example:
	//
	// 10318
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DescribeDefenseRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDefenseRuleRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseRuleRequest) SetDefenseType(v string) *DescribeDefenseRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetInstanceId(v string) *DescribeDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetRegionId(v string) *DescribeDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetResource(v string) *DescribeDefenseRuleRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetRuleId(v int64) *DescribeDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetTemplateId(v int64) *DescribeDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
