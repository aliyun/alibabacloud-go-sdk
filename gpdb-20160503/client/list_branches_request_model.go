// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListBranchesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBranchesRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListBranchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBranchesRequest
	GetPageSize() *int32
	SetParentBranchId(v string) *ListBranchesRequest
	GetParentBranchId() *string
	SetRegionId(v string) *ListBranchesRequest
	GetRegionId() *string
	SetSearch(v string) *ListBranchesRequest
	GetSearch() *string
	SetSortBy(v string) *ListBranchesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListBranchesRequest
	GetSortOrder() *string
}

type ListBranchesRequest struct {
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. It is not required for the first query. For subsequent queries, use the NextToken returned from the previous query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// Valid values:
	//
	// - 10
	//
	// - 20
	//
	// - 50
	//
	// - 100
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent branch ID, used to specify the parent branch for a new branch or as a query filter condition.
	//
	// example:
	//
	// br-main
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The region ID. Must be specified when creating a primary branch. When creating a sub-branch, it inherits the region of the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search keyword. Supports fuzzy search by branch ID or branch name.
	//
	// example:
	//
	// main
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// The sort field.
	//
	// Valid values:
	//
	// - BranchName: Sort by branch name.
	//
	// - CreateTime: Sort by creation time.
	//
	// - LastRunTime: Sort by last run time.
	//
	// Default value: CreateTime.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort direction.
	//
	// Valid values:
	//
	// - Asc: Ascending order.
	//
	// - Desc: Descending order.
	//
	// Default value: Desc.
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListBranchesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBranchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBranchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBranchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBranchesRequest) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *ListBranchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBranchesRequest) GetSearch() *string {
	return s.Search
}

func (s *ListBranchesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListBranchesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListBranchesRequest) SetMaxResults(v int32) *ListBranchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBranchesRequest) SetNextToken(v string) *ListBranchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBranchesRequest) SetPageNumber(v int32) *ListBranchesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBranchesRequest) SetPageSize(v int32) *ListBranchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListBranchesRequest) SetParentBranchId(v string) *ListBranchesRequest {
	s.ParentBranchId = &v
	return s
}

func (s *ListBranchesRequest) SetRegionId(v string) *ListBranchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListBranchesRequest) SetSearch(v string) *ListBranchesRequest {
	s.Search = &v
	return s
}

func (s *ListBranchesRequest) SetSortBy(v string) *ListBranchesRequest {
	s.SortBy = &v
	return s
}

func (s *ListBranchesRequest) SetSortOrder(v string) *ListBranchesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListBranchesRequest) Validate() error {
	return dara.Validate(s)
}
