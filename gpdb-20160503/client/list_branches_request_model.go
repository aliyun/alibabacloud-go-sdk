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
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The cursor for the paged query. You do not need to specify this parameter for the first request. For subsequent requests, use the NextToken value returned in the previous response for paging.
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
	// The parent branch ID. This parameter specifies the parent branch for a new branch or a query filter.
	//
	// example:
	//
	// br-main
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The region ID. This parameter is required when you create a primary branch. When you create a sub-branch, the region is inherited from the primary branch by default.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The search keyword. Fuzzy search by branch ID or branch name is supported.
	//
	// example:
	//
	// main
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// The field by which to sort the results.
	//
	// Valid values:
	//
	// - BranchName: sorts by branch name.
	//
	// - CreateTime: sorts by creation time.
	//
	// - LastRunTime: sorts by last run time.
	//
	// Default value: CreateTime.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order.
	//
	// Valid values:
	//
	// - Asc: ascending order.
	//
	// - Desc: descending order.
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
