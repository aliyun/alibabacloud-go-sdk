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
	// The rule does not take effect on resources in the specified regions. The resources in the specified regions are not evaluated. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The rule does not take effect on resources in the specified resource groups. The resources in the specified resource groups are not evaluated. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The rule does not take effect on the specified resources. The specified resources are not evaluated. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The maximum number of entries to return on each page. Default value: 200.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which to start the query. If this parameter is left empty, the query starts from the beginning.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The scope of region IDs. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The rule takes effect only on resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rule takes effect on the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The managed rules that have been selected. Separate multiple rule identifiers with commas (,).
	//
	// The system does not recommend managed rules that are of the same resource type as the selected managed rules.
	//
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
