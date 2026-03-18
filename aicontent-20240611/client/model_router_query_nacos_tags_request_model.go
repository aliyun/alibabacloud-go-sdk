// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *ModelRouterQueryNacosTagsRequest
	GetGroupBy() *string
	SetMaxResults(v int32) *ModelRouterQueryNacosTagsRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryNacosTagsRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryNacosTagsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryNacosTagsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryNacosTagsRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryNacosTagsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryNacosTagsRequest
	GetPageSize() *int32
}

type ModelRouterQueryNacosTagsRequest struct {
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
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
}

func (s ModelRouterQueryNacosTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosTagsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosTagsRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryNacosTagsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryNacosTagsRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryNacosTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryNacosTagsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryNacosTagsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryNacosTagsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryNacosTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryNacosTagsRequest) SetGroupBy(v string) *ModelRouterQueryNacosTagsRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetMaxResults(v int32) *ModelRouterQueryNacosTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetNeedTotalCount(v bool) *ModelRouterQueryNacosTagsRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetNextToken(v string) *ModelRouterQueryNacosTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetOrderBy(v string) *ModelRouterQueryNacosTagsRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetOrderDirection(v string) *ModelRouterQueryNacosTagsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetPageIndex(v int32) *ModelRouterQueryNacosTagsRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) SetPageSize(v int32) *ModelRouterQueryNacosTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryNacosTagsRequest) Validate() error {
	return dara.Validate(s)
}
