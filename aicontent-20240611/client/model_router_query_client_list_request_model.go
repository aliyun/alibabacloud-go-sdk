// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *ModelRouterQueryClientListRequest
	GetGroupBy() *string
	SetKeyword(v string) *ModelRouterQueryClientListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryClientListRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryClientListRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryClientListRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryClientListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryClientListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryClientListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryClientListRequest
	GetPageSize() *int32
	SetParentId(v int64) *ModelRouterQueryClientListRequest
	GetParentId() *int64
	SetStatus(v int32) *ModelRouterQueryClientListRequest
	GetStatus() *int32
}

type ModelRouterQueryClientListRequest struct {
	// The field to group the results by.
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Specifies whether to return the total count of entries.
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The token for the next page of results. Set this to the nextToken value from the previous response. Omit for the first page.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The field to sort the results by.
	//
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction.
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The page number of the results to return. The first page is 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// Filters the results by status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterQueryClientListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryClientListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryClientListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryClientListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryClientListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryClientListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryClientListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryClientListRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *ModelRouterQueryClientListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterQueryClientListRequest) SetGroupBy(v string) *ModelRouterQueryClientListRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetKeyword(v string) *ModelRouterQueryClientListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetMaxResults(v int32) *ModelRouterQueryClientListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetNeedTotalCount(v bool) *ModelRouterQueryClientListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetNextToken(v string) *ModelRouterQueryClientListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetOrderBy(v string) *ModelRouterQueryClientListRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetOrderDirection(v string) *ModelRouterQueryClientListRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetPageIndex(v int32) *ModelRouterQueryClientListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetPageSize(v int32) *ModelRouterQueryClientListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetParentId(v int64) *ModelRouterQueryClientListRequest {
	s.ParentId = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) SetStatus(v int32) *ModelRouterQueryClientListRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterQueryClientListRequest) Validate() error {
	return dara.Validate(s)
}
