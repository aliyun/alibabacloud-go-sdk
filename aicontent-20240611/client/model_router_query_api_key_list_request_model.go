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
	SetIncludeMemberKeys(v bool) *ModelRouterQueryApiKeyListRequest
	GetIncludeMemberKeys() *bool
	SetKeyword(v string) *ModelRouterQueryApiKeyListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryApiKeyListRequest
	GetMaxResults() *int32
	SetMemberUserIds(v string) *ModelRouterQueryApiKeyListRequest
	GetMemberUserIds() *string
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
	// The client ID used to filter the results.
	//
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The field by which to group the results.
	//
	// example:
	//
	// resourceId
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// Optional. If set to true, the keys of members under the department are also included when filtering by department.
	//
	// example:
	//
	// true
	IncludeMemberKeys *bool `json:"includeMemberKeys,omitempty" xml:"includeMemberKeys,omitempty"`
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
	// Optional. Filters by member IDs. Separate multiple member IDs with commas. If this parameter is not specified, the department and all its members are included. If an empty value is specified, only the department is included without members.
	//
	// example:
	//
	// 30001,30002
	MemberUserIds *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// Specifies whether to return the total count.
	//
	// example:
	//
	// true
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The pagination token. An empty value indicates that no more pages are available.
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
	// The sort direction.
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
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The status used to filter the results.
	//
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

func (s *ModelRouterQueryApiKeyListRequest) GetIncludeMemberKeys() *bool {
	return s.IncludeMemberKeys
}

func (s *ModelRouterQueryApiKeyListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryApiKeyListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryApiKeyListRequest) GetMemberUserIds() *string {
	return s.MemberUserIds
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

func (s *ModelRouterQueryApiKeyListRequest) SetIncludeMemberKeys(v bool) *ModelRouterQueryApiKeyListRequest {
	s.IncludeMemberKeys = &v
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

func (s *ModelRouterQueryApiKeyListRequest) SetMemberUserIds(v string) *ModelRouterQueryApiKeyListRequest {
	s.MemberUserIds = &v
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
