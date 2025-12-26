// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListKnowledgeBasesRequest
	GetCreator() *string
	SetKnowledgeBaseId(v string) *ListKnowledgeBasesRequest
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseType(v string) *ListKnowledgeBasesRequest
	GetKnowledgeBaseType() *string
	SetMaxResults(v int32) *ListKnowledgeBasesRequest
	GetMaxResults() *int32
	SetName(v string) *ListKnowledgeBasesRequest
	GetName() *string
	SetNextToken(v string) *ListKnowledgeBasesRequest
	GetNextToken() *string
	SetOrder(v string) *ListKnowledgeBasesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListKnowledgeBasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBasesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListKnowledgeBasesRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListKnowledgeBasesRequest
	GetWorkspaceId() *string
}

type ListKnowledgeBasesRequest struct {
	Creator           *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	KnowledgeBaseId   *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	KnowledgeBaseType *string `json:"KnowledgeBaseType,omitempty" xml:"KnowledgeBaseType,omitempty"`
	MaxResults        *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DESC
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListKnowledgeBasesRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBasesRequest) GetKnowledgeBaseType() *string {
	return s.KnowledgeBaseType
}

func (s *ListKnowledgeBasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesRequest) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeBasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListKnowledgeBasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBasesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListKnowledgeBasesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKnowledgeBasesRequest) SetCreator(v string) *ListKnowledgeBasesRequest {
	s.Creator = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBasesRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetKnowledgeBaseType(v string) *ListKnowledgeBasesRequest {
	s.KnowledgeBaseType = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetMaxResults(v int32) *ListKnowledgeBasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetName(v string) *ListKnowledgeBasesRequest {
	s.Name = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNextToken(v string) *ListKnowledgeBasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetOrder(v string) *ListKnowledgeBasesRequest {
	s.Order = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetPageNumber(v int32) *ListKnowledgeBasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetPageSize(v int32) *ListKnowledgeBasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetSortBy(v string) *ListKnowledgeBasesRequest {
	s.SortBy = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetWorkspaceId(v string) *ListKnowledgeBasesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
