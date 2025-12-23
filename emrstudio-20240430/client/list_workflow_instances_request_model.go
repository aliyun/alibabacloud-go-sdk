// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListWorkflowInstancesRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListWorkflowInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowInstancesRequest
	GetNextToken() *string
	SetStartTime(v string) *ListWorkflowInstancesRequest
	GetStartTime() *string
	SetStatus(v string) *ListWorkflowInstancesRequest
	GetStatus() *string
	SetWorkflowId(v string) *ListWorkflowInstancesRequest
	GetWorkflowId() *string
	SetWorkspaceId(v string) *ListWorkflowInstancesRequest
	GetWorkspaceId() *string
}

type ListWorkflowInstancesRequest struct {
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
	// 2024-03-27 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowInstancesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowInstancesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkflowInstancesRequest) SetEndTime(v string) *ListWorkflowInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetMaxResults(v int32) *ListWorkflowInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetNextToken(v string) *ListWorkflowInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStartTime(v string) *ListWorkflowInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStatus(v string) *ListWorkflowInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkflowId(v string) *ListWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkspaceId(v string) *ListWorkflowInstancesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}
