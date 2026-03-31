// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRecommendManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateRecommendManagedRulesRequest
	GetAggregatorId() *string
	SetExcludeRegionIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetExcludeResourceIdsScope() *string
	SetMaxResults(v int32) *ListAggregateRecommendManagedRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregateRecommendManagedRulesRequest
	GetNextToken() *string
	SetRegionIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *ListAggregateRecommendManagedRulesRequest
	GetResourceIdsScope() *string
	SetSelectedManagedRuleIdentifiers(v string) *ListAggregateRecommendManagedRulesRequest
	GetSelectedManagedRuleIdentifiers() *string
}

type ListAggregateRecommendManagedRulesRequest struct {
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// example:
	//
	// ram-user-mfa-check-v2,ram-user-last-login-expired-check
	SelectedManagedRuleIdentifiers *string `json:"SelectedManagedRuleIdentifiers,omitempty" xml:"SelectedManagedRuleIdentifiers,omitempty"`
}

func (s ListAggregateRecommendManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRecommendManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateRecommendManagedRulesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateRecommendManagedRulesRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateRecommendManagedRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateRecommendManagedRulesRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *ListAggregateRecommendManagedRulesRequest) GetSelectedManagedRuleIdentifiers() *string {
	return s.SelectedManagedRuleIdentifiers
}

func (s *ListAggregateRecommendManagedRulesRequest) SetAggregatorId(v string) *ListAggregateRecommendManagedRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetExcludeRegionIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetExcludeResourceGroupIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetExcludeResourceIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetMaxResults(v int32) *ListAggregateRecommendManagedRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetNextToken(v string) *ListAggregateRecommendManagedRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetRegionIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetResourceGroupIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetResourceIdsScope(v string) *ListAggregateRecommendManagedRulesRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) SetSelectedManagedRuleIdentifiers(v string) *ListAggregateRecommendManagedRulesRequest {
	s.SelectedManagedRuleIdentifiers = &v
	return s
}

func (s *ListAggregateRecommendManagedRulesRequest) Validate() error {
	return dara.Validate(s)
}
