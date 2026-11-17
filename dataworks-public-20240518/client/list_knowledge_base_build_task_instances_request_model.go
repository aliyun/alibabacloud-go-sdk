// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTaskInstancesRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *ListKnowledgeBaseBuildTaskInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListKnowledgeBaseBuildTaskInstancesRequest
	GetNextToken() *string
	SetTaskId(v string) *ListKnowledgeBaseBuildTaskInstancesRequest
	GetTaskId() *string
}

type ListKnowledgeBaseBuildTaskInstancesRequest struct {
	// The target knowledge base.
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
	// The pagination token.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The task ID used to filter results.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTaskInstancesRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) SetMaxResults(v int32) *ListKnowledgeBaseBuildTaskInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) SetNextToken(v string) *ListKnowledgeBaseBuildTaskInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) SetTaskId(v string) *ListKnowledgeBaseBuildTaskInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
