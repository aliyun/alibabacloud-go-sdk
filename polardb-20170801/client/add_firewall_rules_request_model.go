// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddFirewallRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *AddFirewallRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddFirewallRulesRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *AddFirewallRulesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AddFirewallRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddFirewallRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v string) *AddFirewallRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *AddFirewallRulesRequest
	GetRuleName() *string
}

type AddFirewallRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// {"id":"test","enabled":"true","mode":"Collecting","users":{"applies_to":[]},"endpoint":"[{"EndpointName":"pe-************","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************K"}]","type":"WhiteList"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s AddFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *AddFirewallRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddFirewallRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddFirewallRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddFirewallRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddFirewallRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddFirewallRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddFirewallRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *AddFirewallRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddFirewallRulesRequest) SetDBClusterId(v string) *AddFirewallRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddFirewallRulesRequest) SetOwnerAccount(v string) *AddFirewallRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddFirewallRulesRequest) SetOwnerId(v int64) *AddFirewallRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *AddFirewallRulesRequest) SetResourceGroupId(v string) *AddFirewallRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddFirewallRulesRequest) SetResourceOwnerAccount(v string) *AddFirewallRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddFirewallRulesRequest) SetResourceOwnerId(v int64) *AddFirewallRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddFirewallRulesRequest) SetRuleConfig(v string) *AddFirewallRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *AddFirewallRulesRequest) SetRuleName(v string) *AddFirewallRulesRequest {
	s.RuleName = &v
	return s
}

func (s *AddFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}
