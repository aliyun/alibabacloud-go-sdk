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
	SetTaskId(v string) *DescribeComfyTasksRequest
	GetTaskId() *string
	SetTaskState(v string) *DescribeComfyTasksRequest
	GetTaskState() *string
	SetWorkflowId(v string) *DescribeComfyTasksRequest
	GetWorkflowId() *string
}

type DescribeComfyTasksRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to display per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Comfy workflow ID used as a filter condition.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status used as a filter condition.
	//
	// example:
	//
	// QUEUED
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The Comfy workflow ID used as a filter condition.
	//
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

func (s *DescribeComfyTasksRequest) GetTaskId() *string {
	return s.TaskId
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

func (s *DescribeComfyTasksRequest) SetTaskId(v string) *DescribeComfyTasksRequest {
	s.TaskId = &v
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
