// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAutoGroupingRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoGroupingRulesRequest
	GetNextToken() *string
	SetRuleId(v string) *ListAutoGroupingRulesRequest
	GetRuleId() *string
	SetRuleName(v string) *ListAutoGroupingRulesRequest
	GetRuleName() *string
	SetRuleType(v string) *ListAutoGroupingRulesRequest
	GetRuleType() *string
}

type ListAutoGroupingRulesRequest struct {
	// The maximum number of entries to return in a single request. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, you can use NextToken to send another request to retrieve the content after the current truncation point.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Auto group transfer rule for Project A online resources
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type. Valid values:
	//
	// - custom_condition: custom rule-based group transfer.
	//
	// - associated_transfer: attached-resource allocation group transfer.
	//
	// example:
	//
	// custom_condition
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListAutoGroupingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoGroupingRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoGroupingRulesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ListAutoGroupingRulesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoGroupingRulesRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ListAutoGroupingRulesRequest) SetMaxResults(v int32) *ListAutoGroupingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoGroupingRulesRequest) SetNextToken(v string) *ListAutoGroupingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoGroupingRulesRequest) SetRuleId(v string) *ListAutoGroupingRulesRequest {
	s.RuleId = &v
	return s
}

func (s *ListAutoGroupingRulesRequest) SetRuleName(v string) *ListAutoGroupingRulesRequest {
	s.RuleName = &v
	return s
}

func (s *ListAutoGroupingRulesRequest) SetRuleType(v string) *ListAutoGroupingRulesRequest {
	s.RuleType = &v
	return s
}

func (s *ListAutoGroupingRulesRequest) Validate() error {
	return dara.Validate(s)
}
