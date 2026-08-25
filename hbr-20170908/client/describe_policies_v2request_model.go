// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAccountScope(v string) *DescribePoliciesV2Request
	GetAccountScope() *string
	SetAccounts(v []*DescribePoliciesV2RequestAccounts) *DescribePoliciesV2Request
	GetAccounts() []*DescribePoliciesV2RequestAccounts
	SetMaxResults(v int32) *DescribePoliciesV2Request
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePoliciesV2Request
	GetNextToken() *string
	SetPolicyId(v string) *DescribePoliciesV2Request
	GetPolicyId() *string
	SetRuleScope(v string) *DescribePoliciesV2Request
	GetRuleScope() *string
}

type DescribePoliciesV2Request struct {
	AccountScope *string                              `json:"AccountScope,omitempty" xml:"AccountScope,omitempty"`
	Accounts     []*DescribePoliciesV2RequestAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token required to retrieve the next page of policies.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// po-000************2l6
	PolicyId  *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RuleScope *string `json:"RuleScope,omitempty" xml:"RuleScope,omitempty"`
}

func (s DescribePoliciesV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2Request) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2Request) GetAccountScope() *string {
	return s.AccountScope
}

func (s *DescribePoliciesV2Request) GetAccounts() []*DescribePoliciesV2RequestAccounts {
	return s.Accounts
}

func (s *DescribePoliciesV2Request) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePoliciesV2Request) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePoliciesV2Request) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePoliciesV2Request) GetRuleScope() *string {
	return s.RuleScope
}

func (s *DescribePoliciesV2Request) SetAccountScope(v string) *DescribePoliciesV2Request {
	s.AccountScope = &v
	return s
}

func (s *DescribePoliciesV2Request) SetAccounts(v []*DescribePoliciesV2RequestAccounts) *DescribePoliciesV2Request {
	s.Accounts = v
	return s
}

func (s *DescribePoliciesV2Request) SetMaxResults(v int32) *DescribePoliciesV2Request {
	s.MaxResults = &v
	return s
}

func (s *DescribePoliciesV2Request) SetNextToken(v string) *DescribePoliciesV2Request {
	s.NextToken = &v
	return s
}

func (s *DescribePoliciesV2Request) SetPolicyId(v string) *DescribePoliciesV2Request {
	s.PolicyId = &v
	return s
}

func (s *DescribePoliciesV2Request) SetRuleScope(v string) *DescribePoliciesV2Request {
	s.RuleScope = &v
	return s
}

func (s *DescribePoliciesV2Request) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePoliciesV2RequestAccounts struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s DescribePoliciesV2RequestAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2RequestAccounts) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2RequestAccounts) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribePoliciesV2RequestAccounts) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribePoliciesV2RequestAccounts) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribePoliciesV2RequestAccounts) SetCrossAccountRoleName(v string) *DescribePoliciesV2RequestAccounts {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribePoliciesV2RequestAccounts) SetCrossAccountType(v string) *DescribePoliciesV2RequestAccounts {
	s.CrossAccountType = &v
	return s
}

func (s *DescribePoliciesV2RequestAccounts) SetCrossAccountUserId(v int64) *DescribePoliciesV2RequestAccounts {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribePoliciesV2RequestAccounts) Validate() error {
	return dara.Validate(s)
}
