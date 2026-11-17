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
	SetProxyId(v string) *ModifyIpsRulesToDefaultRequest
	GetProxyId() *string
	SetRuleType(v string) *ModifyIpsRulesToDefaultRequest
	GetRuleType() *string
	SetRules(v string) *ModifyIpsRulesToDefaultRequest
	GetRules() *string
	SetSourceIp(v string) *ModifyIpsRulesToDefaultRequest
	GetSourceIp() *string
}

type ModifyIpsRulesToDefaultRequest struct {
	// The attack application.
	//
	// example:
	//
	// PHP
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The Cloud Firewall type. Valid values:
	//
	// - **VpcFirewall**: virtual private cloud (VPC) firewall
	//
	// - **InternetFirewall*	- (default): the Internet firewall
	//
	// example:
	//
	// InternetFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The language type. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Instance ID of the NAT firewall instance. This parameter is required only when FirewallType is set to **NatFirewall**. In this case, the scope is the rules under this instance. This parameter is ignored for Internet and VPC firewalls. You can call the [ListNatFirewallIpsConfigs](https://help.aliyun.com/document_detail/2879296.html) operation to query instance ID.
	//
	// example:
	//
	// proxy-nat53f4bbc233234271a1d7
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The rule type. Valid values:
	//
	// - **basicRule**
	//
	// - **customize**
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The list of rules. This parameter is conditionally required: when RuleType is set to customize, you must specify at least one of this parameter or AttackApp. If you specify only RuleType without Rules or AttackApp, the API returns an ErrorParameters error.
	//
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
	// The source IP address of the requester.
	//
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

func (s *ModifyIpsRulesToDefaultRequest) GetProxyId() *string {
	return s.ProxyId
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

func (s *ModifyIpsRulesToDefaultRequest) SetProxyId(v string) *ModifyIpsRulesToDefaultRequest {
	s.ProxyId = &v
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
