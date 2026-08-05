// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelsShrink(v string) *ListOfflineTaskShrinkRequest
	GetLabelsShrink() *string
	SetPageNumber(v int32) *ListOfflineTaskShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOfflineTaskShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListOfflineTaskShrinkRequest
	GetRegionId() *string
	SetTaskName(v string) *ListOfflineTaskShrinkRequest
	GetTaskName() *string
	SetTaskStatusShrink(v string) *ListOfflineTaskShrinkRequest
	GetTaskStatusShrink() *string
}

type ListOfflineTaskShrinkRequest struct {
	// The list of task labels.
	LabelsShrink *string `json:"labels,omitempty" xml:"labels,omitempty"`
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
	TaskStatusShrink *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s ListOfflineTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *ListOfflineTaskShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOfflineTaskShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOfflineTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfflineTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListOfflineTaskShrinkRequest) GetTaskStatusShrink() *string {
	return s.TaskStatusShrink
}

func (s *ListOfflineTaskShrinkRequest) SetLabelsShrink(v string) *ListOfflineTaskShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) SetPageNumber(v int32) *ListOfflineTaskShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) SetPageSize(v int32) *ListOfflineTaskShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) SetRegionId(v string) *ListOfflineTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) SetTaskName(v string) *ListOfflineTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) SetTaskStatusShrink(v string) *ListOfflineTaskShrinkRequest {
	s.TaskStatusShrink = &v
	return s
}

func (s *ListOfflineTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
