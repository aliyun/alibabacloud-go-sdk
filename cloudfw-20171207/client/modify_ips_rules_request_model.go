// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallType(v string) *ModifyIpsRulesRequest
	GetFirewallType() *string
	SetLang(v string) *ModifyIpsRulesRequest
	GetLang() *string
	SetRuleAction(v string) *ModifyIpsRulesRequest
	GetRuleAction() *string
	SetRuleType(v string) *ModifyIpsRulesRequest
	GetRuleType() *string
	SetRules(v string) *ModifyIpsRulesRequest
	GetRules() *string
	SetSourceIp(v string) *ModifyIpsRulesRequest
	GetSourceIp() *string
}

type ModifyIpsRulesRequest struct {
	// example:
	//
	// VpcFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alert
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// basicRule
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\\"4664138c-4f81-4650-9c8d-2230ea0d****\\"]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// example:
	//
	// 218.1.147.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyIpsRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *ModifyIpsRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyIpsRulesRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ModifyIpsRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyIpsRulesRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyIpsRulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyIpsRulesRequest) SetFirewallType(v string) *ModifyIpsRulesRequest {
	s.FirewallType = &v
	return s
}

func (s *ModifyIpsRulesRequest) SetLang(v string) *ModifyIpsRulesRequest {
	s.Lang = &v
	return s
}

func (s *ModifyIpsRulesRequest) SetRuleAction(v string) *ModifyIpsRulesRequest {
	s.RuleAction = &v
	return s
}

func (s *ModifyIpsRulesRequest) SetRuleType(v string) *ModifyIpsRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyIpsRulesRequest) SetRules(v string) *ModifyIpsRulesRequest {
	s.Rules = &v
	return s
}

func (s *ModifyIpsRulesRequest) SetSourceIp(v string) *ModifyIpsRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyIpsRulesRequest) Validate() error {
	return dara.Validate(s)
}
