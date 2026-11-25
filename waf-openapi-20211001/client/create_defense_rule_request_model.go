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
	SetDryRun(v bool) *CreateDefenseRuleRequest
	GetDryRun() *bool
	SetInstanceId(v string) *CreateDefenseRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateDefenseRuleRequest
	GetRegionId() *string
	SetResource(v string) *CreateDefenseRuleRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *CreateDefenseRuleRequest
	GetResourceType() *string
	SetRules(v string) *CreateDefenseRuleRequest
	GetRules() *string
	SetTemplateId(v int64) *CreateDefenseRuleRequest
	GetTemplateId() *int64
}

type CreateDefenseRuleRequest struct {
	// The WAF protection scenario to be created.
	//
	// Valid values when the protection rule type **DefenseType*	- is set to **template**:
	//
	// - **waf_group**: basic protection.
	//
	// - **waf_base**: new version of web core protection.
	//
	// - **antiscan**: scan protection.
	//
	// - **ip_blacklist**: IP blacklist.
	//
	// - **custom_acl**: custom rules.
	//
	// - **whitelist**: whitelist.
	//
	// - **region_block**: area blocking.
	//
	// - **custom_response**: custom response of the old version.
	//
	// - **cc**: HTTP flood protection.
	//
	// - **tamperproof**: web tamper-proofing.
	//
	// - **dlp**: data leak prevention.
	//
	// - **spike_throttle**: peak traffic throttling.
	//
	// - **bot_manager**: bot management.
	//
	//
	// Valid values when the protection rule type **DefenseType*	- is set to **resource**:
	//
	// -  **account_identifier**: account extraction.
	//
	// - **custom_response**: custom response of the new version.
	//
	// - **waf_codec**: decoding.
	//
	// - **websdk**: WebSDK integration.
	//
	// Valid values when the protection rule type **DefenseType*	- is set to **global**:
	//
	// -  **regular_custom**: custom regular expressions.
	//
	// - **address_book**: address book.
	//
	// - **custom_response**: custom response of the new version.
	//
	// > For globally configured custom responses, you can reference them under protected objects or rules. When custom response rules are referenced at different dimensions, the actual effective priority is: rule level > protected object level > default page.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The type of the protection rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// Specifies whether to enable the dry run mode. If this parameter is not specified, a normal request is initiated. Valid values:
	//
	// - **true**: initiates a dry run request. The system only checks whether the request meets the execution conditions and does not perform the operations specified in the request. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the error code Defense.Control.DryRunOperation is returned.
	//
	// - **false**: initiates a normal request. The system performs the operations specified in the request after the check succeeds.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of your WAF instance.
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
	// The protected object corresponding to the created rule.
	//
	// > This parameter is required only when **DefenseType*	- is set to **resource**.
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
	// The type of the protected target. This parameter applies to rules where **DefenseType*	- is set to **resource**. Valid values:
	//
	// - **single*	- (default): a single protected object.
	//
	// - **group**: a protection group.
	//
	// example:
	//
	// single
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The rule configuration content, which is converted into a string in JSON format constructed by a series of parameters.
	//
	// > The specific parameters involved vary depending on the specified **protection rule type*	- (**DefenseScene**). For more information, see **Protection rule parameter description**.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// The ID of the protection template for the protection rule to be created.
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
	//
	// > The number of rules that can be created in a single protection template is limited. For more information, see **Rule quantity limits**. If the limit is reached, you can call the [CreateDefenseTemplate](https://help.aliyun.com/document_detail/461613.html) operation to create a new protection template, or call the [ModifyDefenseRule](https://help.aliyun.com/document_detail/461422.html) operation to modify an existing rule.
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

func (s *CreateDefenseRuleRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateDefenseRuleRequest) GetResourceType() *string {
	return s.ResourceType
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

func (s *CreateDefenseRuleRequest) SetDryRun(v bool) *CreateDefenseRuleRequest {
	s.DryRun = &v
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

func (s *CreateDefenseRuleRequest) SetResourceType(v string) *CreateDefenseRuleRequest {
	s.ResourceType = &v
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
