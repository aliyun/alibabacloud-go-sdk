// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountScope(v string) *DescribePoliciesV2ShrinkRequest
	GetAccountScope() *string
	SetAccountsShrink(v string) *DescribePoliciesV2ShrinkRequest
	GetAccountsShrink() *string
	SetMaxResults(v int32) *DescribePoliciesV2ShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePoliciesV2ShrinkRequest
	GetNextToken() *string
	SetPolicyId(v string) *DescribePoliciesV2ShrinkRequest
	GetPolicyId() *string
	SetRuleScope(v string) *DescribePoliciesV2ShrinkRequest
	GetRuleScope() *string
}

type DescribePoliciesV2ShrinkRequest struct {
	AccountScope   *string `json:"AccountScope,omitempty" xml:"AccountScope,omitempty"`
	AccountsShrink *string `json:"Accounts,omitempty" xml:"Accounts,omitempty"`
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

func (s DescribePoliciesV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ShrinkRequest) GetAccountScope() *string {
	return s.AccountScope
}

func (s *DescribePoliciesV2ShrinkRequest) GetAccountsShrink() *string {
	return s.AccountsShrink
}

func (s *DescribePoliciesV2ShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePoliciesV2ShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePoliciesV2ShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePoliciesV2ShrinkRequest) GetRuleScope() *string {
	return s.RuleScope
}

func (s *DescribePoliciesV2ShrinkRequest) SetAccountScope(v string) *DescribePoliciesV2ShrinkRequest {
	s.AccountScope = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) SetAccountsShrink(v string) *DescribePoliciesV2ShrinkRequest {
	s.AccountsShrink = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) SetMaxResults(v int32) *DescribePoliciesV2ShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) SetNextToken(v string) *DescribePoliciesV2ShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) SetPolicyId(v string) *DescribePoliciesV2ShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) SetRuleScope(v string) *DescribePoliciesV2ShrinkRequest {
	s.RuleScope = &v
	return s
}

func (s *DescribePoliciesV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
