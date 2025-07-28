// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *CreateDefenseRuleRequest
	GetDefenseScene() *string
	SetDefenseType(v string) *CreateDefenseRuleRequest
	GetDefenseType() *string
	SetInstanceId(v string) *CreateDefenseRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateDefenseRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateDefenseRuleRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetRules(v string) *CreateDefenseRuleRequest
	GetRules() *string
	SetTemplateId(v int64) *CreateDefenseRuleRequest
	GetTemplateId() *int64
}

type CreateDefenseRuleRequest struct {
	// The module to which the protection rule that you want to create belongs.
	//
	// 	- **waf_group:*	- the basic protection rule module.
	//
	// 	- **antiscan:*	- the scan protection module.
	//
	// 	- **ip_blacklist:*	- the IP address blacklist module.
	//
	// 	- **custom_acl:*	- the custom rule module.
	//
	// 	- **whitelist:*	- the whitelist module.
	//
	// 	- **region_block:*	- the region blacklist module.
	//
	// 	- **custom_response:*	- the custom response module.
	//
	// 	- **cc:*	- the HTTP flood protection module.
	//
	// 	- **tamperproof:*	- the website tamper-proofing module.
	//
	// 	- **dlp:*	- the data leakage prevention module.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	DefenseType  *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-****
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
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the rule template for which you want to create a protection rule.
	//
	// example:
	//
	// 1122
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *CreateDefenseRuleRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *CreateDefenseRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDefenseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefenseRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateDefenseRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDefenseRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *CreateDefenseRuleRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateDefenseRuleRequest) SetDefenseScene(v string) *CreateDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetDefenseType(v string) *CreateDefenseRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetInstanceId(v string) *CreateDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetRegionId(v string) *CreateDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetResource(v string) *CreateDefenseRuleRequest {
	s.Resource = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetRules(v string) *CreateDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetTemplateId(v int64) *CreateDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
