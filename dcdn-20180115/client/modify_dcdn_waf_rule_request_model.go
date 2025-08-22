// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleConfig(v string) *ModifyDcdnWafRuleRequest
	GetRuleConfig() *string
	SetRuleId(v int64) *ModifyDcdnWafRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifyDcdnWafRuleRequest
	GetRuleName() *string
	SetRuleStatus(v string) *ModifyDcdnWafRuleRequest
	GetRuleStatus() *string
}

type ModifyDcdnWafRuleRequest struct {
	// The new configurations of the protection rule.
	//
	// > After you modify the configurations of the protection rule, the previous configurations are overwritten.
	//
	// example:
	//
	// {\\"origin\\":\\"custom\\",\\"conditions\\":[{\\"opValue\\":\\"eq\\",\\"key\\":\\"URL\\",\\"values\\":\\"/example\\"},{\\"opValue\\":\\"eq\\",\\"key\\":\\"Header\\",\\"values\\":\\"3333\\",\\"subKey\\":\\"trt\\"}],\\"actionExternal\\":{},\\"action\\":\\"monitor\\",\\"ccStatus\\":1,\\"ratelimit\\":{\\"target\\":\\"remote_addr\\",\\"interval\\":\\"5\\",\\"threshold\\":\\"2\\",\\"effect\\":\\"rule\\",\\"status\\":{\\"code\\":\\"404\\",\\"count\\":\\"2\\"},\\"ttl\\":\\"1800\\"}}\\"
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The ID of the protection rule. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The new name of the protection rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The new status of the protection rule. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ModifyDcdnWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafRuleRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyDcdnWafRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyDcdnWafRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyDcdnWafRuleRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ModifyDcdnWafRuleRequest) SetRuleConfig(v string) *ModifyDcdnWafRuleRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyDcdnWafRuleRequest) SetRuleId(v int64) *ModifyDcdnWafRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDcdnWafRuleRequest) SetRuleName(v string) *ModifyDcdnWafRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyDcdnWafRuleRequest) SetRuleStatus(v string) *ModifyDcdnWafRuleRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyDcdnWafRuleRequest) Validate() error {
	return dara.Validate(s)
}
