// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *ModelRouterQueryNacosProvidersRequest
	GetGroupBy() *string
	SetMaxResults(v int32) *ModelRouterQueryNacosProvidersRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryNacosProvidersRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryNacosProvidersRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryNacosProvidersRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryNacosProvidersRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryNacosProvidersRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryNacosProvidersRequest
	GetPageSize() *int32
}

type ModelRouterQueryNacosProvidersRequest struct {
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

func (s ModelRouterQueryNacosProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryNacosProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryNacosProvidersRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryNacosProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryNacosProvidersRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryNacosProvidersRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryNacosProvidersRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryNacosProvidersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryNacosProvidersRequest) SetGroupBy(v string) *ModelRouterQueryNacosProvidersRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetMaxResults(v int32) *ModelRouterQueryNacosProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetNeedTotalCount(v bool) *ModelRouterQueryNacosProvidersRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetNextToken(v string) *ModelRouterQueryNacosProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetOrderBy(v string) *ModelRouterQueryNacosProvidersRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetOrderDirection(v string) *ModelRouterQueryNacosProvidersRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetPageIndex(v int32) *ModelRouterQueryNacosProvidersRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) SetPageSize(v int32) *ModelRouterQueryNacosProvidersRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersRequest) Validate() error {
	return dara.Validate(s)
}
