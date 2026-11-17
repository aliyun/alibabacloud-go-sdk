// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTasksRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeBaseBuildTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseBuildTasksRequest
	GetNextToken() *string
}

type ListKnowledgeBaseBuildTasksRequest struct {
	// The ID of the target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseBuildTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseBuildTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseBuildTasksRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTasksRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksRequest) SetMaxResults(v int32) *ListKnowledgeBaseBuildTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksRequest) SetNextToken(v string) *ListKnowledgeBaseBuildTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksRequest) Validate() error {
	return dara.Validate(s)
}
