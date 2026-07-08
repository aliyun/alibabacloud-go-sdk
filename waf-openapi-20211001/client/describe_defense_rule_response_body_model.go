// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseRuleResponseBody
	GetRequestId() *string
	SetRule(v *DescribeDefenseRuleResponseBodyRule) *DescribeDefenseRuleResponseBody
	GetRule() *DescribeDefenseRuleResponseBodyRule
}

type DescribeDefenseRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the protection rule.
	Rule *DescribeDefenseRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s DescribeDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseRuleResponseBody) GetRule() *DescribeDefenseRuleResponseBodyRule {
	return s.Rule
}

func (s *DescribeDefenseRuleResponseBody) SetRequestId(v string) *DescribeDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBody) SetRule(v *DescribeDefenseRuleResponseBodyRule) *DescribeDefenseRuleResponseBody {
	s.Rule = v
	return s
}

func (s *DescribeDefenseRuleResponseBody) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDefenseRuleResponseBodyRule struct {
	// The configuration of the protection rule, returned as a JSON string. For more information, see the **Protection rule parameters*	- section in [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html).
	//
	// example:
	//
	// {\\"status\\":1,\\"policyId\\":1012,\\"action\\":\\"block\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The origin of the protection rule. This parameter indicates whether the rule was created by the user or by the system. Valid values:
	//
	// - **custom**: a custom rule.
	//
	// - **system**: a system-generated rule.
	//
	// example:
	//
	// custom
	DefenseOrigin *string `json:"DefenseOrigin,omitempty" xml:"DefenseOrigin,omitempty"`
	// The scenario to which the protection rule applies.
	//
	// If the **DefenseType*	- parameter is set to **template**, the valid values are:
	//
	// - **waf_group**: basic protection with regular expression rules.
	//
	// - **waf_base**: web core protection.
	//
	// - **waf_base_compliance**: basic protection with protocol compliance rules.
	//
	// - **waf_base_sema**: basic protection with semantic analysis rules.
	//
	// - **cc**: HTTP flood protection.
	//
	// - **antiscan_dirscan**: directory traversal blocking in scan protection.
	//
	// - **antiscan_highfreq**: high-frequency scan blocking in scan protection.
	//
	// - **antiscan_scantools**: scanning tool blocking in scan protection.
	//
	// - **ip_blacklist**: an IP address blacklist.
	//
	// - **custom_acl**: a custom rule.
	//
	// - **region_block**: a location blacklist.
	//
	// - **tamperproof**: webpage tamper protection.
	//
	// - **dlp**: data leakage prevention.
	//
	// - **custom_response_block**: a custom response.
	//
	// - **spike_throttle**: peak traffic throttling.
	//
	// If the **DefenseType*	- parameter is set to **resource**, the valid values are:
	//
	// - **account_identifier**: account identification.
	//
	// - **custom_response**: a custom response.
	//
	// - **waf_codec**: data decoding settings.
	//
	// If the **DefenseType*	- parameter is set to **global**, the valid values are:
	//
	// - **regular_custom**: a custom regular expression.
	//
	// - **address_book**: an IP address book.
	//
	// - **custom_response**: a custom response.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
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
	// The time when the protection rule was modified. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665196746000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The protected object to which the protection rule applies.
	//
	// > This parameter is returned only if the **DefenseType*	- parameter is set to **resource**.
	//
	// example:
	//
	// rencs***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the protection rule.
	//
	// example:
	//
	// 2732975
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the protection rule.
	//
	// example:
	//
	// test1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the protection rule. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the protection rule template.
	//
	// > This parameter is returned only if the **DefenseType*	- parameter is set to **template**.
	//
	// example:
	//
	// 9114
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleResponseBodyRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBodyRule) GetConfig() *string {
	return s.Config
}

func (s *DescribeDefenseRuleResponseBodyRule) GetDefenseOrigin() *string {
	return s.DefenseOrigin
}

func (s *DescribeDefenseRuleResponseBodyRule) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseRuleResponseBodyRule) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DescribeDefenseRuleResponseBodyRule) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDefenseRuleResponseBodyRule) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseRuleResponseBodyRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDefenseRuleResponseBodyRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDefenseRuleResponseBodyRule) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDefenseRuleResponseBodyRule) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseRuleResponseBodyRule) SetConfig(v string) *DescribeDefenseRuleResponseBodyRule {
	s.Config = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseOrigin(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseOrigin = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseScene(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseType(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseType = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetGmtModified(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetResource(v string) *DescribeDefenseRuleResponseBodyRule {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleName(v string) *DescribeDefenseRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetStatus(v int32) *DescribeDefenseRuleResponseBodyRule {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetTemplateId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) Validate() error {
	return dara.Validate(s)
}
