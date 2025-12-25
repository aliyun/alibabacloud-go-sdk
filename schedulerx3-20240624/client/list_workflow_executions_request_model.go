// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListWorkflowExecutionsRequest
	GetAppName() *string
	SetClusterId(v string) *ListWorkflowExecutionsRequest
	GetClusterId() *string
	SetEndTime(v string) *ListWorkflowExecutionsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListWorkflowExecutionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowExecutionsRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListWorkflowExecutionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListWorkflowExecutionsRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListWorkflowExecutionsRequest
	GetStartTime() *string
	SetStatus(v int32) *ListWorkflowExecutionsRequest
	GetStatus() *int32
	SetWorkflowExecutionId(v int64) *ListWorkflowExecutionsRequest
	GetWorkflowExecutionId() *int64
	SetWorkflowId(v int64) *ListWorkflowExecutionsRequest
	GetWorkflowId() *int64
	SetWorkflowName(v string) *ListWorkflowExecutionsRequest
	GetWorkflowName() *string
}

type ListWorkflowExecutionsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2025-10-13 16:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2025-10-27 02:15:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// example:
	//
	// myWorkflow
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListWorkflowExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowExecutionsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListWorkflowExecutionsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListWorkflowExecutionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowExecutionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowExecutionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListWorkflowExecutionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowExecutionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowExecutionsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkflowExecutionsRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *ListWorkflowExecutionsRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowExecutionsRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListWorkflowExecutionsRequest) SetAppName(v string) *ListWorkflowExecutionsRequest {
	s.AppName = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetClusterId(v string) *ListWorkflowExecutionsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetEndTime(v string) *ListWorkflowExecutionsRequest {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetMaxResults(v int32) *ListWorkflowExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetNextToken(v string) *ListWorkflowExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetPageNum(v int32) *ListWorkflowExecutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetPageSize(v int32) *ListWorkflowExecutionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetStartTime(v string) *ListWorkflowExecutionsRequest {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetStatus(v int32) *ListWorkflowExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetWorkflowExecutionId(v int64) *ListWorkflowExecutionsRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetWorkflowId(v int64) *ListWorkflowExecutionsRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) SetWorkflowName(v string) *ListWorkflowExecutionsRequest {
	s.WorkflowName = &v
	return s
}

func (s *ListWorkflowExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
