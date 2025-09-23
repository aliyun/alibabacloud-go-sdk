// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyFirewallRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyFirewallRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyFirewallRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyFirewallRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFirewallRulesRequest
	GetResourceOwnerId() *int64
	SetRuleConfig(v string) *ModifyFirewallRulesRequest
	GetRuleConfig() *string
	SetRuleName(v string) *ModifyFirewallRulesRequest
	GetRuleName() *string
}

type ModifyFirewallRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"id":"test","enabled":"true","mode":"Collecting","users":{"applies_to":[]},"endpoint":"[{"EndpointName":"pe-************","EndpointType":"Cluster","DBEndpointDescription":"Cluster Address"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************"},{"EndpointName":"pe-************","EndpointType":"Custom","DBEndpointDescription":"pc-************K"}]","type":"WhiteList"}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyFirewallRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyFirewallRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFirewallRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFirewallRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFirewallRulesRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyFirewallRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyFirewallRulesRequest) SetDBClusterId(v string) *ModifyFirewallRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetOwnerAccount(v string) *ModifyFirewallRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetOwnerId(v int64) *ModifyFirewallRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetResourceOwnerAccount(v string) *ModifyFirewallRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetResourceOwnerId(v int64) *ModifyFirewallRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetRuleConfig(v string) *ModifyFirewallRulesRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyFirewallRulesRequest) SetRuleName(v string) *ModifyFirewallRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}
