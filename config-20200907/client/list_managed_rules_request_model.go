// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterType(v string) *ListManagedRulesRequest
	GetFilterType() *string
	SetKeyword(v string) *ListManagedRulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListManagedRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListManagedRulesRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListManagedRulesRequest
	GetResourceTypes() *string
	SetRiskLevel(v int32) *ListManagedRulesRequest
	GetRiskLevel() *int32
}

type ListManagedRulesRequest struct {
	// The scope for filtering managed rules allows you to filter out managed rules without resource coverage. The possible values are:
	//
	//  - ALL: All rules.
	//
	//  - UNCOVERED_RESOURCE: Filters managed rules where some resources are not covered.
	//
	// example:
	//
	// ALL
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The keyword of the managed rule.
	//
	// example:
	//
	// CDN
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number of the page to return.
	//
	// Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resources to be evaluated based on the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the managed rule. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListManagedRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListManagedRulesRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *ListManagedRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListManagedRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListManagedRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListManagedRulesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListManagedRulesRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListManagedRulesRequest) SetFilterType(v string) *ListManagedRulesRequest {
	s.FilterType = &v
	return s
}

func (s *ListManagedRulesRequest) SetKeyword(v string) *ListManagedRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListManagedRulesRequest) SetPageNumber(v int32) *ListManagedRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesRequest) SetPageSize(v int32) *ListManagedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesRequest) SetResourceTypes(v string) *ListManagedRulesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListManagedRulesRequest) SetRiskLevel(v int32) *ListManagedRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListManagedRulesRequest) Validate() error {
	return dara.Validate(s)
}
