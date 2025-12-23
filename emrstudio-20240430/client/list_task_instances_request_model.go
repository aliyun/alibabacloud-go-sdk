// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTaskInstancesRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListTaskInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTaskInstancesRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListTaskInstancesRequest
	GetSearchVal() *string
	SetStartTime(v string) *ListTaskInstancesRequest
	GetStartTime() *string
	SetStatus(v string) *ListTaskInstancesRequest
	GetStatus() *string
	SetWorkflowInstanceId(v string) *ListTaskInstancesRequest
	GetWorkflowInstanceId() *string
	SetWorkspaceId(v string) *ListTaskInstancesRequest
	GetWorkspaceId() *string
}

type ListTaskInstancesRequest struct {
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
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
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTaskInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTaskInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTaskInstancesRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListTaskInstancesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTaskInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskInstancesRequest) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *ListTaskInstancesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTaskInstancesRequest) SetEndTime(v string) *ListTaskInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListTaskInstancesRequest) SetMaxResults(v int32) *ListTaskInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTaskInstancesRequest) SetNextToken(v string) *ListTaskInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaskInstancesRequest) SetSearchVal(v string) *ListTaskInstancesRequest {
	s.SearchVal = &v
	return s
}

func (s *ListTaskInstancesRequest) SetStartTime(v string) *ListTaskInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListTaskInstancesRequest) SetStatus(v string) *ListTaskInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkflowInstanceId(v string) *ListTaskInstancesRequest {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListTaskInstancesRequest) SetWorkspaceId(v string) *ListTaskInstancesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
