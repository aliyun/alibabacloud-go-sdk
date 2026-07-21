// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *ModelRouterQueryConversationListRequest
	GetGroupBy() *string
	SetKeyword(v string) *ModelRouterQueryConversationListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryConversationListRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryConversationListRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryConversationListRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryConversationListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryConversationListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryConversationListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryConversationListRequest
	GetPageSize() *int32
	SetStatus(v int32) *ModelRouterQueryConversationListRequest
	GetStatus() *int32
}

type ModelRouterQueryConversationListRequest struct {
	// The grouping field.
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// A keyword to filter the results.
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
	// Specifies whether to return the total count of items. Set to `true` to return the total count.
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The pagination token for the next page of results. To retrieve the first page, do not specify this parameter.
	//
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The sort field.
	//
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort order. Valid values are `ASC` (ascending) and `DESC` (descending).
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of items to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters the results by status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterQueryConversationListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryConversationListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryConversationListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryConversationListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryConversationListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryConversationListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryConversationListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryConversationListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryConversationListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryConversationListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterQueryConversationListRequest) SetGroupBy(v string) *ModelRouterQueryConversationListRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetKeyword(v string) *ModelRouterQueryConversationListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetMaxResults(v int32) *ModelRouterQueryConversationListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetNeedTotalCount(v bool) *ModelRouterQueryConversationListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetNextToken(v string) *ModelRouterQueryConversationListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetOrderBy(v string) *ModelRouterQueryConversationListRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetOrderDirection(v string) *ModelRouterQueryConversationListRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetPageIndex(v int32) *ModelRouterQueryConversationListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetPageSize(v int32) *ModelRouterQueryConversationListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) SetStatus(v int32) *ModelRouterQueryConversationListRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterQueryConversationListRequest) Validate() error {
	return dara.Validate(s)
}
