// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobAction(v string) *ListKnowledgeBaseJobsRequest
	GetJobAction() *string
	SetKnowledgeBaseJobId(v string) *ListKnowledgeBaseJobsRequest
	GetKnowledgeBaseJobId() *string
	SetMaxResults(v int32) *ListKnowledgeBaseJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseJobsRequest
	GetNextToken() *string
	SetOrder(v string) *ListKnowledgeBaseJobsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListKnowledgeBaseJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKnowledgeBaseJobsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListKnowledgeBaseJobsRequest
	GetSortBy() *string
	SetStatus(v string) *ListKnowledgeBaseJobsRequest
	GetStatus() *string
	SetWorkspaceId(v string) *ListKnowledgeBaseJobsRequest
	GetWorkspaceId() *string
}

type ListKnowledgeBaseJobsRequest struct {
	// example:
	//
	// SyncIndex
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// example:
	//
	// kbjob-9m******54
	KnowledgeBaseJobId *string `json:"KnowledgeBaseJobId,omitempty" xml:"KnowledgeBaseJobId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DESC
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
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListKnowledgeBaseJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseJobsRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseJobsRequest) GetJobAction() *string {
	return s.JobAction
}

func (s *ListKnowledgeBaseJobsRequest) GetKnowledgeBaseJobId() *string {
	return s.KnowledgeBaseJobId
}

func (s *ListKnowledgeBaseJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListKnowledgeBaseJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKnowledgeBaseJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKnowledgeBaseJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListKnowledgeBaseJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListKnowledgeBaseJobsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKnowledgeBaseJobsRequest) SetJobAction(v string) *ListKnowledgeBaseJobsRequest {
	s.JobAction = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetKnowledgeBaseJobId(v string) *ListKnowledgeBaseJobsRequest {
	s.KnowledgeBaseJobId = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetMaxResults(v int32) *ListKnowledgeBaseJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetNextToken(v string) *ListKnowledgeBaseJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetOrder(v string) *ListKnowledgeBaseJobsRequest {
	s.Order = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetPageNumber(v int32) *ListKnowledgeBaseJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetPageSize(v int32) *ListKnowledgeBaseJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetSortBy(v string) *ListKnowledgeBaseJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetStatus(v string) *ListKnowledgeBaseJobsRequest {
	s.Status = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) SetWorkspaceId(v string) *ListKnowledgeBaseJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKnowledgeBaseJobsRequest) Validate() error {
	return dara.Validate(s)
}
