// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAsyncTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAsyncTasksRequest
	GetPageSize() *int32
	SetResourceIds(v string) *ListAsyncTasksRequest
	GetResourceIds() *string
	SetResourceType(v string) *ListAsyncTasksRequest
	GetResourceType() *string
	SetTaskStatus(v string) *ListAsyncTasksRequest
	GetTaskStatus() *string
	SetTaskType(v string) *ListAsyncTasksRequest
	GetTaskType() *string
}

type ListAsyncTasksRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tasks to display per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource ID to which the task belongs, such as a site ID. You can obtain the site ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// example:
	//
	// 4080****3752512
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type.
	//
	// example:
	//
	// site
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The current status of the task. Valid values:
	//
	// - in_progress: in progress.
	//
	// - success: completed.
	//
	// - fail: failed.
	//
	// example:
	//
	// success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The task type. For example, the task type for applying for a free certificate is free_cert.
	//
	// example:
	//
	// free_cert
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListAsyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAsyncTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAsyncTasksRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *ListAsyncTasksRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAsyncTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListAsyncTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAsyncTasksRequest) SetPageNumber(v int32) *ListAsyncTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAsyncTasksRequest) SetPageSize(v int32) *ListAsyncTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListAsyncTasksRequest) SetResourceIds(v string) *ListAsyncTasksRequest {
	s.ResourceIds = &v
	return s
}

func (s *ListAsyncTasksRequest) SetResourceType(v string) *ListAsyncTasksRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatus(v string) *ListAsyncTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskType(v string) *ListAsyncTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
