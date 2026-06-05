// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeComfyTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyTasksRequest
	GetPageSize() *int32
	SetTaskState(v string) *DescribeComfyTasksRequest
	GetTaskState() *string
	SetWorkflowId(v string) *DescribeComfyTasksRequest
	GetWorkflowId() *string
}

type DescribeComfyTasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// QUEUED
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DescribeComfyTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyTasksRequest) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeComfyTasksRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeComfyTasksRequest) SetPageNumber(v int32) *DescribeComfyTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyTasksRequest) SetPageSize(v int32) *DescribeComfyTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyTasksRequest) SetTaskState(v string) *DescribeComfyTasksRequest {
	s.TaskState = &v
	return s
}

func (s *DescribeComfyTasksRequest) SetWorkflowId(v string) *DescribeComfyTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *DescribeComfyTasksRequest) Validate() error {
	return dara.Validate(s)
}
