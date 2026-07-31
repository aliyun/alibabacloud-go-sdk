// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelRouterQueryModelGroupListRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryModelGroupListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupListRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryModelGroupListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelGroupListRequest
	GetPageSize() *int32
	SetType(v string) *ModelRouterQueryModelGroupListRequest
	GetType() *string
}

type ModelRouterQueryModelGroupListRequest struct {
	// The keyword for fuzzy match by group name.
	//
	// example:
	//
	// Professional
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// An unused parameter.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The group type filter. Valid values: system, manual, and all. Default value: all.
	//
	// example:
	//
	// all
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelRouterQueryModelGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupListRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryModelGroupListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelGroupListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupListRequest) GetType() *string {
	return s.Type
}

func (s *ModelRouterQueryModelGroupListRequest) SetKeyword(v string) *ModelRouterQueryModelGroupListRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) SetMaxResults(v int32) *ModelRouterQueryModelGroupListRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) SetNextToken(v string) *ModelRouterQueryModelGroupListRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) SetPageIndex(v int32) *ModelRouterQueryModelGroupListRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) SetPageSize(v int32) *ModelRouterQueryModelGroupListRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) SetType(v string) *ModelRouterQueryModelGroupListRequest {
	s.Type = &v
	return s
}

func (s *ModelRouterQueryModelGroupListRequest) Validate() error {
	return dara.Validate(s)
}
