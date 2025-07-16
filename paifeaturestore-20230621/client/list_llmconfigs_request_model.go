// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListLLMConfigsRequest
	GetMaxResults() *int32
	SetName(v string) *ListLLMConfigsRequest
	GetName() *string
	SetNextToken(v string) *ListLLMConfigsRequest
	GetNextToken() *string
	SetOrder(v string) *ListLLMConfigsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListLLMConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLLMConfigsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLLMConfigsRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListLLMConfigsRequest
	GetWorkspaceId() *string
}

type ListLLMConfigsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// llm_config_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListLLMConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLLMConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListLLMConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLLMConfigsRequest) GetName() *string {
	return s.Name
}

func (s *ListLLMConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLLMConfigsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListLLMConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLLMConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLLMConfigsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLLMConfigsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListLLMConfigsRequest) SetMaxResults(v int32) *ListLLMConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLLMConfigsRequest) SetName(v string) *ListLLMConfigsRequest {
	s.Name = &v
	return s
}

func (s *ListLLMConfigsRequest) SetNextToken(v string) *ListLLMConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLLMConfigsRequest) SetOrder(v string) *ListLLMConfigsRequest {
	s.Order = &v
	return s
}

func (s *ListLLMConfigsRequest) SetPageNumber(v int32) *ListLLMConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLLMConfigsRequest) SetPageSize(v int32) *ListLLMConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLLMConfigsRequest) SetSortBy(v string) *ListLLMConfigsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLLMConfigsRequest) SetWorkspaceId(v string) *ListLLMConfigsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListLLMConfigsRequest) Validate() error {
	return dara.Validate(s)
}
