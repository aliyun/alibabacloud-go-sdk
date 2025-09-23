// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteFirewallRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteFirewallRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteFirewallRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteFirewallRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteFirewallRulesRequest
	GetResourceOwnerId() *int64
	SetRuleNameList(v string) *DeleteFirewallRulesRequest
	GetRuleNameList() *string
}

type DeleteFirewallRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DeleteFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteFirewallRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteFirewallRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteFirewallRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteFirewallRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteFirewallRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DeleteFirewallRulesRequest) SetDBClusterId(v string) *DeleteFirewallRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetOwnerAccount(v string) *DeleteFirewallRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetOwnerId(v int64) *DeleteFirewallRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetResourceOwnerAccount(v string) *DeleteFirewallRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetResourceOwnerId(v int64) *DeleteFirewallRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetRuleNameList(v string) *DeleteFirewallRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DeleteFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}
