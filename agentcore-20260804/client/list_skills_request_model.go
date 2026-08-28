// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListSkillsRequest
	GetOrderBy() *string
	SetOwner(v string) *ListSkillsRequest
	GetOwner() *string
	SetPageNo(v int32) *ListSkillsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSkillsRequest
	GetPageSize() *int32
	SetScope(v string) *ListSkillsRequest
	GetScope() *string
	SetSearch(v string) *ListSkillsRequest
	GetSearch() *string
	SetSkillName(v string) *ListSkillsRequest
	GetSkillName() *string
}

type ListSkillsRequest struct {
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The sort field. The value download_count is supported. Default value: gmt_modified.
	//
	// example:
	//
	// download_count
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// Filters results by owner.
	//
	// example:
	//
	// user123
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters results by visibility. Valid values:
	//
	// - PUBLIC
	//
	// - PRIVATE
	//
	// example:
	//
	// PUBLIC
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The search mode. Valid values:
	//
	// - accurate: exact match.
	//
	// - blur: fuzzy match.
	//
	// example:
	//
	// blur
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// The filter keyword.
	//
	// example:
	//
	// customer
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSkillsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListSkillsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSkillsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListSkillsRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ListSkillsRequest) SetMaxResults(v int32) *ListSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsRequest) SetNextToken(v string) *ListSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillsRequest) SetOrderBy(v string) *ListSkillsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSkillsRequest) SetOwner(v string) *ListSkillsRequest {
	s.Owner = &v
	return s
}

func (s *ListSkillsRequest) SetPageNo(v int32) *ListSkillsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetScope(v string) *ListSkillsRequest {
	s.Scope = &v
	return s
}

func (s *ListSkillsRequest) SetSearch(v string) *ListSkillsRequest {
	s.Search = &v
	return s
}

func (s *ListSkillsRequest) SetSkillName(v string) *ListSkillsRequest {
	s.SkillName = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
