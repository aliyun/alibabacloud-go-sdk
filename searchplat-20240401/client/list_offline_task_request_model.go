// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*string) *ListOfflineTaskRequest
	GetLabels() []*string
	SetPageNumber(v int32) *ListOfflineTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOfflineTaskRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListOfflineTaskRequest
	GetRegionId() *string
	SetTaskName(v string) *ListOfflineTaskRequest
	GetTaskName() *string
	SetTaskStatus(v []*string) *ListOfflineTaskRequest
	GetTaskStatus() []*string
}

type ListOfflineTaskRequest struct {
	// The list of task labels.
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 0
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The task name.
	//
	// example:
	//
	// syh
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The task status.
	TaskStatus []*string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty" type:"Repeated"`
}

func (s ListOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskRequest) GetLabels() []*string {
	return s.Labels
}

func (s *ListOfflineTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOfflineTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfflineTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListOfflineTaskRequest) GetTaskStatus() []*string {
	return s.TaskStatus
}

func (s *ListOfflineTaskRequest) SetLabels(v []*string) *ListOfflineTaskRequest {
	s.Labels = v
	return s
}

func (s *ListOfflineTaskRequest) SetPageNumber(v int32) *ListOfflineTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOfflineTaskRequest) SetPageSize(v int32) *ListOfflineTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListOfflineTaskRequest) SetRegionId(v string) *ListOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfflineTaskRequest) SetTaskName(v string) *ListOfflineTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListOfflineTaskRequest) SetTaskStatus(v []*string) *ListOfflineTaskRequest {
	s.TaskStatus = v
	return s
}

func (s *ListOfflineTaskRequest) Validate() error {
	return dara.Validate(s)
}
