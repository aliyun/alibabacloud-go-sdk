// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelGroupModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ModelRouterQueryModelGroupModelsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ModelRouterQueryModelGroupModelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelGroupModelsRequest
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryModelGroupModelsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelGroupModelsRequest
	GetPageSize() *int32
}

type ModelRouterQueryModelGroupModelsRequest struct {
	// Searches by model name or identifier.
	//
	// example:
	//
	// qwen
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 20
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

func (s ModelRouterQueryModelGroupModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelGroupModelsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelGroupModelsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ModelRouterQueryModelGroupModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelGroupModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelGroupModelsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelGroupModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelGroupModelsRequest) SetKeyword(v string) *ModelRouterQueryModelGroupModelsRequest {
	s.Keyword = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsRequest) SetMaxResults(v int32) *ModelRouterQueryModelGroupModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsRequest) SetNextToken(v string) *ModelRouterQueryModelGroupModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsRequest) SetPageIndex(v int32) *ModelRouterQueryModelGroupModelsRequest {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsRequest) SetPageSize(v int32) *ModelRouterQueryModelGroupModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelGroupModelsRequest) Validate() error {
	return dara.Validate(s)
}
