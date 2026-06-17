// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPSRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackApp(v string) *DescribeIPSRulesRequest
	GetAttackApp() *string
	SetAttackAppCategory(v []*string) *DescribeIPSRulesRequest
	GetAttackAppCategory() []*string
	SetAttackApps(v []*string) *DescribeIPSRulesRequest
	GetAttackApps() []*string
	SetAttackType(v string) *DescribeIPSRulesRequest
	GetAttackType() *string
	SetCve(v string) *DescribeIPSRulesRequest
	GetCve() *string
	SetDefaultAction(v string) *DescribeIPSRulesRequest
	GetDefaultAction() *string
	SetFirewallType(v string) *DescribeIPSRulesRequest
	GetFirewallType() *string
	SetLang(v string) *DescribeIPSRulesRequest
	GetLang() *string
	SetOrder(v string) *DescribeIPSRulesRequest
	GetOrder() *string
	SetPageNo(v int64) *DescribeIPSRulesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeIPSRulesRequest
	GetPageSize() *int64
	SetQueryModify(v string) *DescribeIPSRulesRequest
	GetQueryModify() *string
	SetRuleAction(v string) *DescribeIPSRulesRequest
	GetRuleAction() *string
	SetRuleClass(v string) *DescribeIPSRulesRequest
	GetRuleClass() *string
	SetRuleId(v string) *DescribeIPSRulesRequest
	GetRuleId() *string
	SetRuleLevel(v int64) *DescribeIPSRulesRequest
	GetRuleLevel() *int64
	SetRuleName(v string) *DescribeIPSRulesRequest
	GetRuleName() *string
	SetRuleType(v string) *DescribeIPSRulesRequest
	GetRuleType() *string
	SetSort(v string) *DescribeIPSRulesRequest
	GetSort() *string
	SetSourceIp(v string) *DescribeIPSRulesRequest
	GetSourceIp() *string
	SetVpcFirewallId(v string) *DescribeIPSRulesRequest
	GetVpcFirewallId() *string
}

type DescribeIPSRulesRequest struct {
	// The application targeted by the attack.
	//
	// example:
	//
	// SMB
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The categories of applications targeted by attacks.
	AttackAppCategory []*string `json:"AttackAppCategory,omitempty" xml:"AttackAppCategory,omitempty" type:"Repeated"`
	// The applications targeted by attacks.
	AttackApps []*string `json:"AttackApps,omitempty" xml:"AttackApps,omitempty" type:"Repeated"`
	// The attack type.
	//
	// example:
	//
	// Web attack
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The CVE ID.
	//
	// example:
	//
	// CVE-2026-25253
	Cve *string `json:"Cve,omitempty" xml:"Cve,omitempty"`
	// The status of the rule.
	//
	// example:
	//
	// 1
	DefaultAction *string `json:"DefaultAction,omitempty" xml:"DefaultAction,omitempty"`
	// The type of firewall.
	//
	// example:
	//
	// VpcFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The language of the request and response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether to retrieve only modified rules.
	//
	// example:
	//
	// true
	QueryModify *string `json:"QueryModify,omitempty" xml:"QueryModify,omitempty"`
	// The rule action.
	//
	// example:
	//
	// alert
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The inspection mode.
	//
	// example:
	//
	// dropStrict
	RuleClass *string `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// f835533b-01ef-49f4-b172-85bbfd0e****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule level.
	//
	// example:
	//
	// 1
	RuleLevel *int64 `json:"RuleLevel,omitempty" xml:"RuleLevel,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Nmap scan detection
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type.
	//
	// This parameter is required.
	//
	// example:
	//
	// patchRule
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The field to sort by.
	//
	// example:
	//
	// UpdateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 8.139.222.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// cen-rnbkqx4a8er21a****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeIPSRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPSRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIPSRulesRequest) GetAttackApp() *string {
	return s.AttackApp
}

func (s *DescribeIPSRulesRequest) GetAttackAppCategory() []*string {
	return s.AttackAppCategory
}

func (s *DescribeIPSRulesRequest) GetAttackApps() []*string {
	return s.AttackApps
}

func (s *DescribeIPSRulesRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeIPSRulesRequest) GetCve() *string {
	return s.Cve
}

func (s *DescribeIPSRulesRequest) GetDefaultAction() *string {
	return s.DefaultAction
}

func (s *DescribeIPSRulesRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeIPSRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIPSRulesRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeIPSRulesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeIPSRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeIPSRulesRequest) GetQueryModify() *string {
	return s.QueryModify
}

func (s *DescribeIPSRulesRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *DescribeIPSRulesRequest) GetRuleClass() *string {
	return s.RuleClass
}

func (s *DescribeIPSRulesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeIPSRulesRequest) GetRuleLevel() *int64 {
	return s.RuleLevel
}

func (s *DescribeIPSRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeIPSRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeIPSRulesRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeIPSRulesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeIPSRulesRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeIPSRulesRequest) SetAttackApp(v string) *DescribeIPSRulesRequest {
	s.AttackApp = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetAttackAppCategory(v []*string) *DescribeIPSRulesRequest {
	s.AttackAppCategory = v
	return s
}

func (s *DescribeIPSRulesRequest) SetAttackApps(v []*string) *DescribeIPSRulesRequest {
	s.AttackApps = v
	return s
}

func (s *DescribeIPSRulesRequest) SetAttackType(v string) *DescribeIPSRulesRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetCve(v string) *DescribeIPSRulesRequest {
	s.Cve = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetDefaultAction(v string) *DescribeIPSRulesRequest {
	s.DefaultAction = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetFirewallType(v string) *DescribeIPSRulesRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetLang(v string) *DescribeIPSRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetOrder(v string) *DescribeIPSRulesRequest {
	s.Order = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetPageNo(v int64) *DescribeIPSRulesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetPageSize(v int64) *DescribeIPSRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetQueryModify(v string) *DescribeIPSRulesRequest {
	s.QueryModify = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleAction(v string) *DescribeIPSRulesRequest {
	s.RuleAction = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleClass(v string) *DescribeIPSRulesRequest {
	s.RuleClass = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleId(v string) *DescribeIPSRulesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleLevel(v int64) *DescribeIPSRulesRequest {
	s.RuleLevel = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleName(v string) *DescribeIPSRulesRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetRuleType(v string) *DescribeIPSRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetSort(v string) *DescribeIPSRulesRequest {
	s.Sort = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetSourceIp(v string) *DescribeIPSRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIPSRulesRequest) SetVpcFirewallId(v string) *DescribeIPSRulesRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeIPSRulesRequest) Validate() error {
	return dara.Validate(s)
}
