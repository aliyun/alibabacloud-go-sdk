// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ModelRouterQueryModelGroupUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupUsersRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryModelGroupUsersRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelGroupUsersRequest
	GetPageSize() *int32
}

type ModelRouterQueryModelGroupUsersRequest struct {
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
}

func (s ModelRouterQueryModelGroupUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupUsersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupUsersRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelGroupUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupUsersRequest) SetMaxResults(v int32) *ModelRouterQueryModelGroupUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersRequest) SetNextToken(v string) *ModelRouterQueryModelGroupUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersRequest) SetPageIndex(v int32) *ModelRouterQueryModelGroupUsersRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersRequest) SetPageSize(v int32) *ModelRouterQueryModelGroupUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupUsersRequest) Validate() error {
	return dara.Validate(s)
}
