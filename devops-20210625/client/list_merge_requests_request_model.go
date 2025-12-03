// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListMergeRequestsRequest
	GetAccessToken() *string
	SetAuthorIds(v string) *ListMergeRequestsRequest
	GetAuthorIds() *string
	SetCreatedAfter(v string) *ListMergeRequestsRequest
	GetCreatedAfter() *string
	SetCreatedBefore(v string) *ListMergeRequestsRequest
	GetCreatedBefore() *string
	SetFilter(v string) *ListMergeRequestsRequest
	GetFilter() *string
	SetGroupIds(v string) *ListMergeRequestsRequest
	GetGroupIds() *string
	SetLabelIds(v string) *ListMergeRequestsRequest
	GetLabelIds() *string
	SetOrderBy(v string) *ListMergeRequestsRequest
	GetOrderBy() *string
	SetOrganizationId(v string) *ListMergeRequestsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListMergeRequestsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListMergeRequestsRequest
	GetPageSize() *int64
	SetProjectIds(v string) *ListMergeRequestsRequest
	GetProjectIds() *string
	SetReviewerIds(v string) *ListMergeRequestsRequest
	GetReviewerIds() *string
	SetSearch(v string) *ListMergeRequestsRequest
	GetSearch() *string
	SetSort(v string) *ListMergeRequestsRequest
	GetSort() *string
	SetState(v string) *ListMergeRequestsRequest
	GetState() *string
}

type ListMergeRequestsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 1234567890
	AuthorIds     *string `json:"authorIds,omitempty" xml:"authorIds,omitempty"`
	CreatedAfter  *string `json:"createdAfter,omitempty" xml:"createdAfter,omitempty"`
	CreatedBefore *string `json:"createdBefore,omitempty" xml:"createdBefore,omitempty"`
	// example:
	//
	// new
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// 889910, 889911
	GroupIds *string `json:"groupIds,omitempty" xml:"groupIds,omitempty"`
	LabelIds *string `json:"labelIds,omitempty" xml:"labelIds,omitempty"`
	// example:
	//
	// updated_at
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2308912, 2308913
	ProjectIds *string `json:"projectIds,omitempty" xml:"projectIds,omitempty"`
	// example:
	//
	// 1234567890123
	ReviewerIds *string `json:"reviewerIds,omitempty" xml:"reviewerIds,omitempty"`
	// example:
	//
	// test-search
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// example:
	//
	// opened
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListMergeRequestsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListMergeRequestsRequest) GetAuthorIds() *string {
	return s.AuthorIds
}

func (s *ListMergeRequestsRequest) GetCreatedAfter() *string {
	return s.CreatedAfter
}

func (s *ListMergeRequestsRequest) GetCreatedBefore() *string {
	return s.CreatedBefore
}

func (s *ListMergeRequestsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListMergeRequestsRequest) GetGroupIds() *string {
	return s.GroupIds
}

func (s *ListMergeRequestsRequest) GetLabelIds() *string {
	return s.LabelIds
}

func (s *ListMergeRequestsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMergeRequestsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMergeRequestsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListMergeRequestsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMergeRequestsRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *ListMergeRequestsRequest) GetReviewerIds() *string {
	return s.ReviewerIds
}

func (s *ListMergeRequestsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListMergeRequestsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListMergeRequestsRequest) GetState() *string {
	return s.State
}

func (s *ListMergeRequestsRequest) SetAccessToken(v string) *ListMergeRequestsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestsRequest) SetAuthorIds(v string) *ListMergeRequestsRequest {
	s.AuthorIds = &v
	return s
}

func (s *ListMergeRequestsRequest) SetCreatedAfter(v string) *ListMergeRequestsRequest {
	s.CreatedAfter = &v
	return s
}

func (s *ListMergeRequestsRequest) SetCreatedBefore(v string) *ListMergeRequestsRequest {
	s.CreatedBefore = &v
	return s
}

func (s *ListMergeRequestsRequest) SetFilter(v string) *ListMergeRequestsRequest {
	s.Filter = &v
	return s
}

func (s *ListMergeRequestsRequest) SetGroupIds(v string) *ListMergeRequestsRequest {
	s.GroupIds = &v
	return s
}

func (s *ListMergeRequestsRequest) SetLabelIds(v string) *ListMergeRequestsRequest {
	s.LabelIds = &v
	return s
}

func (s *ListMergeRequestsRequest) SetOrderBy(v string) *ListMergeRequestsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMergeRequestsRequest) SetOrganizationId(v string) *ListMergeRequestsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestsRequest) SetPage(v int64) *ListMergeRequestsRequest {
	s.Page = &v
	return s
}

func (s *ListMergeRequestsRequest) SetPageSize(v int64) *ListMergeRequestsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMergeRequestsRequest) SetProjectIds(v string) *ListMergeRequestsRequest {
	s.ProjectIds = &v
	return s
}

func (s *ListMergeRequestsRequest) SetReviewerIds(v string) *ListMergeRequestsRequest {
	s.ReviewerIds = &v
	return s
}

func (s *ListMergeRequestsRequest) SetSearch(v string) *ListMergeRequestsRequest {
	s.Search = &v
	return s
}

func (s *ListMergeRequestsRequest) SetSort(v string) *ListMergeRequestsRequest {
	s.Sort = &v
	return s
}

func (s *ListMergeRequestsRequest) SetState(v string) *ListMergeRequestsRequest {
	s.State = &v
	return s
}

func (s *ListMergeRequestsRequest) Validate() error {
	return dara.Validate(s)
}
