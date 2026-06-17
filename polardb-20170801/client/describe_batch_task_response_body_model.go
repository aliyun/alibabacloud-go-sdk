// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *DescribeBatchTaskResponseBody
	GetApplicationType() *string
	SetBatchId(v string) *DescribeBatchTaskResponseBody
	GetBatchId() *string
	SetRequestId(v string) *DescribeBatchTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeBatchTaskResponseBody
	GetStatus() *string
	SetSubTasks(v []*DescribeBatchTaskResponseBodySubTasks) *DescribeBatchTaskResponseBody
	GetSubTasks() []*DescribeBatchTaskResponseBodySubTasks
	SetSuccessCount(v int32) *DescribeBatchTaskResponseBody
	GetSuccessCount() *int32
	SetTaskBegin(v string) *DescribeBatchTaskResponseBody
	GetTaskBegin() *string
	SetTaskEnd(v string) *DescribeBatchTaskResponseBody
	GetTaskEnd() *string
	SetTaskName(v string) *DescribeBatchTaskResponseBody
	GetTaskName() *string
	SetTaskType(v string) *DescribeBatchTaskResponseBody
	GetTaskType() *string
	SetTotalCount(v int32) *DescribeBatchTaskResponseBody
	GetTotalCount() *int32
}

type DescribeBatchTaskResponseBody struct {
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// pcb-xxx
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25C70FF3-D49B-594D-BECE-0DE2BA1D8BBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of subtasks.
	SubTasks []*DescribeBatchTaskResponseBodySubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	// The number of successful subtasks.
	//
	// example:
	//
	// 8
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2026/05/28T16:38:20Z
	TaskBegin *string `json:"TaskBegin,omitempty" xml:"TaskBegin,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2026/05/28T15:23:47Z
	TaskEnd *string `json:"TaskEnd,omitempty" xml:"TaskEnd,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// batch_task_test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// polarclaw_install_skills
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of subtasks.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchTaskResponseBody) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *DescribeBatchTaskResponseBody) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeBatchTaskResponseBody) GetSubTasks() []*DescribeBatchTaskResponseBodySubTasks {
	return s.SubTasks
}

func (s *DescribeBatchTaskResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeBatchTaskResponseBody) GetTaskBegin() *string {
	return s.TaskBegin
}

func (s *DescribeBatchTaskResponseBody) GetTaskEnd() *string {
	return s.TaskEnd
}

func (s *DescribeBatchTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeBatchTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeBatchTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBatchTaskResponseBody) SetApplicationType(v string) *DescribeBatchTaskResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetBatchId(v string) *DescribeBatchTaskResponseBody {
	s.BatchId = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetRequestId(v string) *DescribeBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetStatus(v string) *DescribeBatchTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetSubTasks(v []*DescribeBatchTaskResponseBodySubTasks) *DescribeBatchTaskResponseBody {
	s.SubTasks = v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetSuccessCount(v int32) *DescribeBatchTaskResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetTaskBegin(v string) *DescribeBatchTaskResponseBody {
	s.TaskBegin = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetTaskEnd(v string) *DescribeBatchTaskResponseBody {
	s.TaskEnd = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetTaskName(v string) *DescribeBatchTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetTaskType(v string) *DescribeBatchTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) SetTotalCount(v int32) *DescribeBatchTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchTaskResponseBody) Validate() error {
	if s.SubTasks != nil {
		for _, item := range s.SubTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBatchTaskResponseBodySubTasks struct {
	// The error message.
	//
	// example:
	//
	// aliuid:1422133474238823 assumeOssRole not exist,serviceName:aliyunesarealtimelogpushossrole
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pc-pz5f6mvi1p84t35d7
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the subtask.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the subtask started.
	//
	// example:
	//
	// 2026/05/28T17:38:20Z
	TaskBegin *string `json:"TaskBegin,omitempty" xml:"TaskBegin,omitempty"`
	// The time when the subtask ended.
	//
	// example:
	//
	// 2026/05/28T20:38:20Z
	TaskEnd *string `json:"TaskEnd,omitempty" xml:"TaskEnd,omitempty"`
	// The ID of the subtask.
	//
	// example:
	//
	// 629271331
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeBatchTaskResponseBodySubTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTaskResponseBodySubTasks) GoString() string {
	return s.String()
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetTaskBegin() *string {
	return s.TaskBegin
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetTaskEnd() *string {
	return s.TaskEnd
}

func (s *DescribeBatchTaskResponseBodySubTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetErrorMsg(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetInstanceId(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.InstanceId = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetStatus(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.Status = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetTaskBegin(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.TaskBegin = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetTaskEnd(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.TaskEnd = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) SetTaskId(v string) *DescribeBatchTaskResponseBodySubTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchTaskResponseBodySubTasks) Validate() error {
	return dara.Validate(s)
}
