// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryPaidResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetGroupBy() *string
	SetMaxResults(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetNextToken() *string
	SetOrderBy(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetPageSize() *int32
	SetResourceType(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest
	GetResourceType() *string
}

type AliyunConsoleOpenApiQueryPaidResourceRequest struct {
	// groupBy
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// needTotalCount
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// orderBy
	//
	// example:
	//
	// resourceId
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// orderDirection
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// pageIndex
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// pageSize
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// resourceType
	//
	// example:
	//
	// ALL
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AliyunConsoleOpenApiQueryPaidResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryPaidResourceRequest) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetGroupBy(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.GroupBy = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetMaxResults(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetNeedTotalCount(v bool) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetNextToken(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.NextToken = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetOrderBy(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.OrderBy = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetOrderDirection(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.OrderDirection = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetPageIndex(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.PageIndex = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetPageSize(v int32) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.PageSize = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) SetResourceType(v string) *AliyunConsoleOpenApiQueryPaidResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceRequest) Validate() error {
	return dara.Validate(s)
}
