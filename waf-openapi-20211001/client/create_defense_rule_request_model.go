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
	// The protection scenario for the rule.
	//
	// When **DefenseType*	- is **template**, valid values are:
	//
	// - **waf_group**: basic protection.
	//
	// - **waf_base**: new version of Web core protection.
	//
	// - **antiscan**: scan protection.
	//
	// - **ip_blacklist**: IP blacklist.
	//
	// - **custom_acl**: custom rules.
	//
	// - **whitelist**: whitelist.
	//
	// - **region_block**: location blacklist.
	//
	// - **custom_response**: legacy custom response.
	//
	// - **cc**: HTTP flood protection.
	//
	// - **tamperproof**: webpage tamper protection.
	//
	// - **dlp**: data leak prevention.
	//
	// - **spike_throttle**: peak traffic throttling.
	//
	// When **DefenseType*	- is **resource**, valid values are:
	//
	// - **account_identifier**: account extraction.
	//
	// - **custom_response**: new version of custom response.
	//
	// - **waf_codec**: decoding.
	//
	// When **DefenseType*	- is **global**, valid values are:
	//
	// - **regular_custom**: custom regex.
	//
	// - **address_book**: address book.
	//
	// - **custom_response**: new version of custom response.
	//
	// > For globally configured custom responses, users can reference them under protected objects or rules. When referenced at different levels, the effective logic follows this order: rule level > protected object level > default page.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The type of the protection rule. Valid values:
	//
	// - **template*	- (default): template-based protection rules.
	//
	// - **resource**: rules applied at the protected object level.
	//
	// - **global**: global-level rules.
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
	// waf_v2_public_cn-****
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
	// The protected object associated with the rule.
	//
	// > Provide this parameter only when **DefenseType*	- is **resource**.
	//
	// example:
	//
	// sec****-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The rule configuration content, specified as a JSON string.
	//
	// > The specific parameters vary based on the specified **DefenseType*	- (**DefenseScene**). For details, see **Protection Rule Parameter Descriptions**.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection template to which the rule belongs.
	//
	// > Provide this parameter only when **DefenseType*	- is **template**.
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
