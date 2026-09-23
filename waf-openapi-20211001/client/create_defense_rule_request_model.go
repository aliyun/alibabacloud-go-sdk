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
	SetRules(v string) *CreateDefenseRuleRequest
	GetRules() *string
	SetTemplateId(v int64) *CreateDefenseRuleRequest
	GetTemplateId() *int64
}

type CreateDefenseRuleRequest struct {
	// The WAF protection scenario to create.
	//
	// When the protection rule type **DefenseType*	- is set to **template**, valid values:
	//
	// - **waf_group**: basic protection.
	//
	// - **waf_base**: new version of Web Core Protection.
	//
	// - **antiscan**: Scan Protection.
	//
	// - **ip_blacklist**: IP Blacklist.
	//
	// - **custom_acl**: Custom Rule.
	//
	// - **whitelist**: Whitelist.
	//
	// - **region_block**: Location Blacklist.
	//
	// - **custom_response**: legacy Custom Response.
	//
	// - **cc**: HTTP Flood Protection.
	//
	// - **tamperproof**: web tamper proofing.
	//
	// - **dlp**: Information Leak Prevention.
	//
	// - **spike_throttle**: peak traffic throttling.
	//
	// - **bot_manager**: BOT Management.
	//
	//
	// When the protection rule type **DefenseType*	- is set to **resource**, valid values:
	//
	// - **account_identifier**: Account Extraction.
	//
	// - **custom_response**: new version of Custom Response.
	//
	// - **waf_codec**: Decoding.
	//
	// - **websdk**: WebSDK Integration.
	//
	// When the protection rule type **DefenseType*	- is set to **global**, valid values:
	//
	// - **regular_custom**: Custom Regex.
	//
	// - **address_book**: Address Book.
	//
	// - **custom_response**: new version of Custom Response.
	//
	// > For the custom response in global configuration, users can reference it at the protected object or rule level. When custom response rules are referenced at different dimensions, the actual effective logic is: rule level > protected object level > default page.
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
	// Specifies whether to enable the dry run mode. If you do not specify this parameter, a normal request is sent. Valid values:
	//
	// - **true**: A dry run request is sent. The system only checks whether the request meets the execution conditions without performing the specified operation. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the error code Defense.Control.DryRunOperation is returned.
	//
	// - **false**: A normal request is sent. The specified operation is performed after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of your current WAF instance.
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
	// The protected object associated with the rule to be created.
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
	// The rule configuration content, which is a string converted from a JSON-formatted array of parameters.
	//
	// > The specific parameters vary depending on the specified **protection rule type*	- (**DefenseScene**). For more information, refer to **Protection rule parameter descriptions**.
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
	// > There is an upper limit on the number of rules that can be created within the same protection template. For specific limits, refer to **Rule quantity limits**. When the rule quantity has reached the upper limit, you can call the [CreateDefenseTemplate](https://help.aliyun.com/document_detail/461613.html) operation to create a new protection template. You can also call the [ModifyDefenseRule](https://help.aliyun.com/document_detail/461422.html) operation to modify an existing rule.
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
