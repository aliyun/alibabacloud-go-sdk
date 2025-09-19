// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCycleTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *DescribeCycleTaskListRequest
	GetConfigId() *string
	SetCurrentPage(v int32) *DescribeCycleTaskListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeCycleTaskListRequest
	GetPageSize() *int32
	SetTaskName(v string) *DescribeCycleTaskListRequest
	GetTaskName() *string
	SetTaskType(v string) *DescribeCycleTaskListRequest
	GetTaskType() *string
}

type DescribeCycleTaskListRequest struct {
	// The ID of the task configuration.
	//
	// >  You can call the [CreateCycleTask](~~CreateCycleTask~~) operation to query the IDs of task configurations.
	//
	// example:
	//
	// f93b6ee24cfd0aad44b897ad5051****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus scan task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// example:
	//
	// IMAGE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus scan task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// example:
	//
	// IMAGE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeCycleTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCycleTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCycleTaskListRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCycleTaskListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCycleTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCycleTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCycleTaskListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCycleTaskListRequest) SetConfigId(v string) *DescribeCycleTaskListRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeCycleTaskListRequest) SetCurrentPage(v int32) *DescribeCycleTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCycleTaskListRequest) SetPageSize(v int32) *DescribeCycleTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCycleTaskListRequest) SetTaskName(v string) *DescribeCycleTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeCycleTaskListRequest) SetTaskType(v string) *DescribeCycleTaskListRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeCycleTaskListRequest) Validate() error {
	return dara.Validate(s)
}
