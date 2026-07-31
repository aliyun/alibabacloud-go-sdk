// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ModelRouterQueryModelGroupClientsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupClientsRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryModelGroupClientsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelGroupClientsRequest
	GetPageSize() *int32
}

type ModelRouterQueryModelGroupClientsRequest struct {
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This field is not used.
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

func (s ModelRouterQueryModelGroupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupClientsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupClientsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupClientsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupClientsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelGroupClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupClientsRequest) SetMaxResults(v int32) *ModelRouterQueryModelGroupClientsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsRequest) SetNextToken(v string) *ModelRouterQueryModelGroupClientsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsRequest) SetPageIndex(v int32) *ModelRouterQueryModelGroupClientsRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsRequest) SetPageSize(v int32) *ModelRouterQueryModelGroupClientsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupClientsRequest) Validate() error {
	return dara.Validate(s)
}
