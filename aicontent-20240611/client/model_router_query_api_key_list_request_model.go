// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ModelRouterQueryApiKeyListRequest
	GetClientId() *int64
	SetGroupBy(v string) *ModelRouterQueryApiKeyListRequest
	GetGroupBy() *string
	SetKeyword(v string) *ModelRouterQueryApiKeyListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryApiKeyListRequest
	GetMaxResults() *int32
	SetNeedTotalCount(v bool) *ModelRouterQueryApiKeyListRequest
	GetNeedTotalCount() *bool
	SetNextToken(v string) *ModelRouterQueryApiKeyListRequest
	GetNextToken() *string
	SetOrderBy(v string) *ModelRouterQueryApiKeyListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ModelRouterQueryApiKeyListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ModelRouterQueryApiKeyListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryApiKeyListRequest
	GetPageSize() *int32
	SetStatus(v int32) *ModelRouterQueryApiKeyListRequest
	GetStatus() *int32
}

type ModelRouterQueryApiKeyListRequest struct {
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
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

func (s ModelRouterQueryApiKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyListRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *ModelRouterQueryApiKeyListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ModelRouterQueryApiKeyListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryApiKeyListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryApiKeyListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ModelRouterQueryApiKeyListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryApiKeyListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ModelRouterQueryApiKeyListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ModelRouterQueryApiKeyListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryApiKeyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryApiKeyListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterQueryApiKeyListRequest) SetClientId(v int64) *ModelRouterQueryApiKeyListRequest {
	s.ClientId = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetGroupBy(v string) *ModelRouterQueryApiKeyListRequest {
	s.GroupBy = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetKeyword(v string) *ModelRouterQueryApiKeyListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetMaxResults(v int32) *ModelRouterQueryApiKeyListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetNeedTotalCount(v bool) *ModelRouterQueryApiKeyListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetNextToken(v string) *ModelRouterQueryApiKeyListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetOrderBy(v string) *ModelRouterQueryApiKeyListRequest {
	s.OrderBy = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetOrderDirection(v string) *ModelRouterQueryApiKeyListRequest {
	s.OrderDirection = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetPageIndex(v int32) *ModelRouterQueryApiKeyListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetPageSize(v int32) *ModelRouterQueryApiKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) SetStatus(v int32) *ModelRouterQueryApiKeyListRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterQueryApiKeyListRequest) Validate() error {
	return dara.Validate(s)
}
