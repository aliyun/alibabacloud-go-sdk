// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecommendManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeRegionIdsScope(v string) *ListRecommendManagedRulesRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *ListRecommendManagedRulesRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *ListRecommendManagedRulesRequest
	GetExcludeResourceIdsScope() *string
	SetMaxResults(v int32) *ListRecommendManagedRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRecommendManagedRulesRequest
	GetNextToken() *string
	SetRegionIdsScope(v string) *ListRecommendManagedRulesRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *ListRecommendManagedRulesRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *ListRecommendManagedRulesRequest
	GetResourceIdsScope() *string
	SetSelectedManagedRuleIdentifiers(v string) *ListRecommendManagedRulesRequest
	GetSelectedManagedRuleIdentifiers() *string
}

type ListRecommendManagedRulesRequest struct {
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
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
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
	// ram-user-last-login-expired-check
	SelectedManagedRuleIdentifiers *string `json:"SelectedManagedRuleIdentifiers,omitempty" xml:"SelectedManagedRuleIdentifiers,omitempty"`
}

func (s ListRecommendManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecommendManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRecommendManagedRulesRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRecommendManagedRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecommendManagedRulesRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *ListRecommendManagedRulesRequest) GetSelectedManagedRuleIdentifiers() *string {
	return s.SelectedManagedRuleIdentifiers
}

func (s *ListRecommendManagedRulesRequest) SetExcludeRegionIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetExcludeResourceGroupIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetExcludeResourceIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetMaxResults(v int32) *ListRecommendManagedRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetNextToken(v string) *ListRecommendManagedRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetRegionIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetResourceGroupIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetResourceIdsScope(v string) *ListRecommendManagedRulesRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) SetSelectedManagedRuleIdentifiers(v string) *ListRecommendManagedRulesRequest {
	s.SelectedManagedRuleIdentifiers = &v
	return s
}

func (s *ListRecommendManagedRulesRequest) Validate() error {
	return dara.Validate(s)
}
