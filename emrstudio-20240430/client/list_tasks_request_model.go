// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListTasksRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListTasksRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListTasksRequest
	GetSearchVal() *string
	SetTaskType(v string) *ListTasksRequest
	GetTaskType() *string
	SetWorkflowId(v string) *ListTasksRequest
	GetWorkflowId() *string
	SetWorkspaceId(v string) *ListTasksRequest
	GetWorkspaceId() *string
}

type ListTasksRequest struct {
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// example:
	//
	// w-n72kong0832****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTasksRequest) SetMaxResults(v string) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetSearchVal(v string) *ListTasksRequest {
	s.SearchVal = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListTasksRequest) SetWorkflowId(v string) *ListTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTasksRequest) SetWorkspaceId(v string) *ListTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
