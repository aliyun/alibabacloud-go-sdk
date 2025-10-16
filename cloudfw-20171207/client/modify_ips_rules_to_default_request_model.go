// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpsRulesToDefaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v string) *ModifyIpsRulesToDefaultRequest
	GetAttackApp() *string
	SetFirewallType(v string) *ModifyIpsRulesToDefaultRequest
	GetFirewallType() *string
	SetLang(v string) *ModifyIpsRulesToDefaultRequest
	GetLang() *string
	SetRuleType(v string) *ModifyIpsRulesToDefaultRequest
	GetRuleType() *string
	SetRules(v string) *ModifyIpsRulesToDefaultRequest
	GetRules() *string
	SetSourceIp(v string) *ModifyIpsRulesToDefaultRequest
	GetSourceIp() *string
}

type ModifyIpsRulesToDefaultRequest struct {
	// example:
	//
	// PHP
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// example:
	//
	// InternetFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// [
	//
	//   "uuid_1",
	//
	//   "uuid_2"
	//
	// ]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// example:
	//
	// 140.205.118.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyIpsRulesToDefaultRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpsRulesToDefaultRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpsRulesToDefaultRequest) GetAttackApp() *string {
	return s.AttackApp
}

func (s *ModifyIpsRulesToDefaultRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *ModifyIpsRulesToDefaultRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyIpsRulesToDefaultRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyIpsRulesToDefaultRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyIpsRulesToDefaultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyIpsRulesToDefaultRequest) SetAttackApp(v string) *ModifyIpsRulesToDefaultRequest {
	s.AttackApp = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) SetFirewallType(v string) *ModifyIpsRulesToDefaultRequest {
	s.FirewallType = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) SetLang(v string) *ModifyIpsRulesToDefaultRequest {
	s.Lang = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) SetRuleType(v string) *ModifyIpsRulesToDefaultRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) SetRules(v string) *ModifyIpsRulesToDefaultRequest {
	s.Rules = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) SetSourceIp(v string) *ModifyIpsRulesToDefaultRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyIpsRulesToDefaultRequest) Validate() error {
	return dara.Validate(s)
}
