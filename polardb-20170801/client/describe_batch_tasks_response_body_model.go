// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeBatchTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBatchTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBatchTasksResponseBody
	GetRequestId() *string
	SetTaskList(v []*DescribeBatchTasksResponseBodyTaskList) *DescribeBatchTasksResponseBody
	GetTaskList() []*DescribeBatchTasksResponseBodyTaskList
	SetTotalCount(v int32) *DescribeBatchTasksResponseBody
	GetTotalCount() *int32
}

type DescribeBatchTasksResponseBody struct {
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task list.
	TaskList []*DescribeBatchTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of tasks that match the query, ignoring pagination.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBatchTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBatchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchTasksResponseBody) GetTaskList() []*DescribeBatchTasksResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeBatchTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBatchTasksResponseBody) SetPageNumber(v int32) *DescribeBatchTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchTasksResponseBody) SetPageSize(v int32) *DescribeBatchTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchTasksResponseBody) SetRequestId(v string) *DescribeBatchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchTasksResponseBody) SetTaskList(v []*DescribeBatchTasksResponseBodyTaskList) *DescribeBatchTasksResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeBatchTasksResponseBody) SetTotalCount(v int32) *DescribeBatchTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchTasksResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBatchTasksResponseBodyTaskList struct {
	// The batch ID.
	//
	// example:
	//
	// pcb-xxx
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of successfully completed subtasks in the batch.
	//
	// example:
	//
	// 4
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The task start time.
	//
	// example:
	//
	// 2026-04-06 20:51:44
	TaskBegin *string `json:"TaskBegin,omitempty" xml:"TaskBegin,omitempty"`
	// The task end time.
	//
	// example:
	//
	// 2026-04-06 22:43:26
	TaskEnd *string `json:"TaskEnd,omitempty" xml:"TaskEnd,omitempty"`
	// The task name.
	//
	// example:
	//
	// batch_task_test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task type.
	//
	// example:
	//
	// polarclaw_install_skills
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of subtasks in the batch.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchTasksResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetStatus() *string {
	return s.Status
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetTaskBegin() *string {
	return s.TaskBegin
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetTaskEnd() *string {
	return s.TaskEnd
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeBatchTasksResponseBodyTaskList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetBatchId(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.BatchId = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetStatus(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetSuccessCount(v int32) *DescribeBatchTasksResponseBodyTaskList {
	s.SuccessCount = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetTaskBegin(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.TaskBegin = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetTaskEnd(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.TaskEnd = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetTaskName(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.TaskName = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetTaskType(v string) *DescribeBatchTasksResponseBodyTaskList {
	s.TaskType = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) SetTotalCount(v int32) *DescribeBatchTasksResponseBodyTaskList {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchTasksResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
