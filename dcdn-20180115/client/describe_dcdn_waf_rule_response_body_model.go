// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnWafRuleResponseBody
	GetRequestId() *string
	SetRule(v *DescribeDcdnWafRuleResponseBodyRule) *DescribeDcdnWafRuleResponseBody
	GetRule() *DescribeDcdnWafRuleResponseBodyRule
}

type DescribeDcdnWafRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the protection rule.
	Rule *DescribeDcdnWafRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s DescribeDcdnWafRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafRuleResponseBody) GetRule() *DescribeDcdnWafRuleResponseBodyRule {
	return s.Rule
}

func (s *DescribeDcdnWafRuleResponseBody) SetRequestId(v string) *DescribeDcdnWafRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBody) SetRule(v *DescribeDcdnWafRuleResponseBodyRule) *DescribeDcdnWafRuleResponseBody {
	s.Rule = v
	return s
}

func (s *DescribeDcdnWafRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafRuleResponseBodyRule struct {
	// The type of the protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: IP address whitelist
	//
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The time when the scaling group was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-29T17:08:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the protection policy.
	//
	// example:
	//
	// 200001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The configurations of the protection rule.
	//
	// example:
	//
	// {\\"action\\":\\"monitor\\",\\"actionExternal\\":\\"{}\\",\\"ccStatus\\":1,\\"conditions\\":[{\\"key\\":\\"URL\\",\\"opValue\\":\\"eq\\",\\"targetKey\\":\\"request_uri\\",\\"values\\":\\"/example\\"},{\\"key\\":\\"Header\\",\\"opValue\\":\\"eq\\",\\"subKey\\":\\"trt\\",\\"targetKey\\":\\"header.trt\\",\\"values\\":\\"3333\\"}],\\"effect\\":\\"service\\",\\"name\\":\\"aaa333\\",\\"origin\\":\\"custom\\",\\"ratelimit\\":{\\"interval\\":5,\\"status\\":{\\"code\\":404,\\"count\\":2,\\"stat\\":{\\"mode\\":\\"count\\",\\"value\\":2.0}},\\"target\\":\\"remote_addr\\",\\"threshold\\":2,\\"ttl\\":1800}}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The ID of the protection rule.
	//
	// example:
	//
	// 100001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the protection rule.
	//
	// example:
	//
	// rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the protection rule. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s DescribeDcdnWafRuleResponseBodyRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeDcdnWafRuleResponseBodyRule) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetDefenseScene(v string) *DescribeDcdnWafRuleResponseBodyRule {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetGmtModified(v string) *DescribeDcdnWafRuleResponseBodyRule {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetPolicyId(v int64) *DescribeDcdnWafRuleResponseBodyRule {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetRuleConfig(v string) *DescribeDcdnWafRuleResponseBodyRule {
	s.RuleConfig = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetRuleId(v int64) *DescribeDcdnWafRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetRuleName(v string) *DescribeDcdnWafRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) SetRuleStatus(v string) *DescribeDcdnWafRuleResponseBodyRule {
	s.RuleStatus = &v
	return s
}

func (s *DescribeDcdnWafRuleResponseBodyRule) Validate() error {
	return dara.Validate(s)
}
