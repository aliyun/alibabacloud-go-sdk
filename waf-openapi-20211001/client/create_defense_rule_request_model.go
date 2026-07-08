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
	// The scenario to which the protection rule is applied.
	//
	// If **DefenseType*	- is set to **template**, valid values are:
	//
	// - **waf_group**: basic protection.
	//
	// - **waf_base**: web core protection.
	//
	// - **antiscan**: scan protection.
	//
	// - **ip_blacklist**: IP address blacklist.
	//
	// - **custom_acl**: custom rule.
	//
	// - **whitelist**: whitelist.
	//
	// - **region_block**: geo-blocking.
	//
	// - **custom_response**: custom response.
	//
	// - **cc**: HTTP flood protection.
	//
	// - **tamperproof**: webpage tamper-proofing.
	//
	// - **dlp**: data leakage prevention.
	//
	// - **spike_throttle**: rate limiting for bursts of traffic.
	//
	// - **bot_manager**: bot management.
	//
	// If **DefenseType*	- is set to **resource**, valid values are:
	//
	// - **account_identifier**: account identification.
	//
	// - **custom_response**: custom response.
	//
	// - **waf_codec**: decoding.
	//
	// If **DefenseType*	- is set to **global**, valid values are:
	//
	// - **regular_custom**: custom regular expression.
	//
	// - **address_book**: address book.
	//
	// - **custom_response**: custom response.
	//
	// > You can apply a global custom response to a protected object or a rule. If you configure custom response rules at different levels, the rule with the finest-grained scope takes precedence. The priority is as follows: rule > protected object > default page.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The type of the protection rule. Valid values:
	//
	// - **template*	- (default): a template-based protection rule.
	//
	// - **resource**: a rule for a specific protected object.
	//
	// - **global**: a global protection rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to get the ID of your WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object to which the rule applies.
	//
	// > This parameter is required only when **DefenseType*	- is set to **resource**.
	//
	// example:
	//
	// sec****-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The details of the protection rule. This value is a JSON string.
	//
	// > The parameters in the JSON string vary based on the value of **DefenseScene**. For more information, see **Protection rule parameters**.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection rule template.
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
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
