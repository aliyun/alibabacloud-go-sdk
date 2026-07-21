// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *ModelRouterQueryNacosTagsRequest
	GetConfigType() *string
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
	// providers
	ConfigType *string `json:"configType,omitempty" xml:"configType,omitempty"`
	// The field by which to group the results.
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Specifies whether to return the total count of results.
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The token from a previous response to retrieve the next page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort order.
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
	// The number of results to return per page.
	//
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

func (s *ModelRouterQueryNacosTagsRequest) GetConfigType() *string {
	return s.ConfigType
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

func (s *ModelRouterQueryNacosTagsRequest) SetConfigType(v string) *ModelRouterQueryNacosTagsRequest {
	s.ConfigType = &v
	return s
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
