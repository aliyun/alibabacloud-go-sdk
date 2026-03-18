// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *ModelRouterQueryModelListRequest
	GetGroupBy() *string
	SetKeyword(v string) *ModelRouterQueryModelListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryModelListRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryModelListRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryModelListRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryModelListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryModelListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryModelListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelListRequest
	GetPageSize() *int32
	SetStatus(v int32) *ModelRouterQueryModelListRequest
	GetStatus() *int32
}

type ModelRouterQueryModelListRequest struct {
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterQueryModelListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryModelListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryModelListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryModelListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryModelListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryModelListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterQueryModelListRequest) SetGroupBy(v string) *ModelRouterQueryModelListRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetKeyword(v string) *ModelRouterQueryModelListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetMaxResults(v int32) *ModelRouterQueryModelListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetNeedTotalCount(v bool) *ModelRouterQueryModelListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetNextToken(v string) *ModelRouterQueryModelListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetOrderBy(v string) *ModelRouterQueryModelListRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetOrderDirection(v string) *ModelRouterQueryModelListRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetPageIndex(v int32) *ModelRouterQueryModelListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetPageSize(v int32) *ModelRouterQueryModelListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) SetStatus(v int32) *ModelRouterQueryModelListRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterQueryModelListRequest) Validate() error {
	return dara.Validate(s)
}
