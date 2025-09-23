// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLRateLimitingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQLRateLimitingRulesRequest
	GetDBClusterId() *string
	SetMaxResults(v int32) *DescribeSQLRateLimitingRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSQLRateLimitingRulesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSQLRateLimitingRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQLRateLimitingRulesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeSQLRateLimitingRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQLRateLimitingRulesRequest
	GetResourceOwnerId() *int64
	SetRuleNameList(v string) *DescribeSQLRateLimitingRulesRequest
	GetRuleNameList() *string
}

type DescribeSQLRateLimitingRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3QbdIPYBaCDXsvvjLHCQfi
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// testrule
	RuleNameList *string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty"`
}

func (s DescribeSQLRateLimitingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLRateLimitingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLRateLimitingRulesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQLRateLimitingRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSQLRateLimitingRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSQLRateLimitingRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQLRateLimitingRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQLRateLimitingRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQLRateLimitingRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQLRateLimitingRulesRequest) GetRuleNameList() *string {
	return s.RuleNameList
}

func (s *DescribeSQLRateLimitingRulesRequest) SetDBClusterId(v string) *DescribeSQLRateLimitingRulesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetMaxResults(v int32) *DescribeSQLRateLimitingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetNextToken(v string) *DescribeSQLRateLimitingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetOwnerAccount(v string) *DescribeSQLRateLimitingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetOwnerId(v int64) *DescribeSQLRateLimitingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetResourceOwnerAccount(v string) *DescribeSQLRateLimitingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetResourceOwnerId(v int64) *DescribeSQLRateLimitingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) SetRuleNameList(v string) *DescribeSQLRateLimitingRulesRequest {
	s.RuleNameList = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesRequest) Validate() error {
	return dara.Validate(s)
}
