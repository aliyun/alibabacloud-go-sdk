// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListManualTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListManualTasksRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListManualTasksRequest
	GetSearchVal() *string
	SetTaskType(v string) *ListManualTasksRequest
	GetTaskType() *string
	SetWorkspaceId(v string) *ListManualTasksRequest
	GetWorkspaceId() *string
}

type ListManualTasksRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListManualTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManualTasksRequest) GoString() string {
	return s.String()
}

func (s *ListManualTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManualTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManualTasksRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListManualTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListManualTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListManualTasksRequest) SetMaxResults(v int32) *ListManualTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManualTasksRequest) SetNextToken(v string) *ListManualTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListManualTasksRequest) SetSearchVal(v string) *ListManualTasksRequest {
	s.SearchVal = &v
	return s
}

func (s *ListManualTasksRequest) SetTaskType(v string) *ListManualTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListManualTasksRequest) SetWorkspaceId(v string) *ListManualTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListManualTasksRequest) Validate() error {
	return dara.Validate(s)
}
