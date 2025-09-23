// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeFirewallRulesRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeFirewallRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFirewallRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeFirewallRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeFirewallRulesRequest
	GetResourceOwnerId() *int64
	SetRuleNameList(v string) *DescribeFirewallRulesRequest
	GetRuleNameList() *string
}

type DescribeFirewallRulesRequest struct {
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
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DescribeFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeFirewallRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFirewallRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFirewallRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFirewallRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeFirewallRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DescribeFirewallRulesRequest) SetDBClusterId(v string) *DescribeFirewallRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeFirewallRulesRequest) SetOwnerAccount(v string) *DescribeFirewallRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFirewallRulesRequest) SetOwnerId(v int64) *DescribeFirewallRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFirewallRulesRequest) SetResourceOwnerAccount(v string) *DescribeFirewallRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFirewallRulesRequest) SetResourceOwnerId(v int64) *DescribeFirewallRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFirewallRulesRequest) SetRuleNameList(v string) *DescribeFirewallRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DescribeFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}
