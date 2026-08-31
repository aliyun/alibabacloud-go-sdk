// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeComfyTasksResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribeComfyTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeComfyTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeComfyTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeComfyTasksResponseBodyTasks) *DescribeComfyTasksResponseBody
	GetTasks() []*DescribeComfyTasksResponseBodyTasks
	SetTotal(v int32) *DescribeComfyTasksResponseBody
	GetTotal() *int32
}

type DescribeComfyTasksResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// > Only applicable to recording queries.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task list.
	Tasks []*DescribeComfyTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The number of tasks.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeComfyTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComfyTasksResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeComfyTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComfyTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComfyTasksResponseBody) GetTasks() []*DescribeComfyTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeComfyTasksResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeComfyTasksResponseBody) SetCode(v int64) *DescribeComfyTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetMessage(v string) *DescribeComfyTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetPageNumber(v int32) *DescribeComfyTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetPageSize(v int32) *DescribeComfyTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetRequestId(v string) *DescribeComfyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetTasks(v []*DescribeComfyTasksResponseBodyTasks) *DescribeComfyTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeComfyTasksResponseBody) SetTotal(v int32) *DescribeComfyTasksResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeComfyTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComfyTasksResponseBodyTasks struct {
	// The creation time.
	//
	// example:
	//
	// 1780387493000
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The task end time.
	//
	// example:
	//
	// 1780387494000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The resource pool ID used by the task.
	//
	// example:
	//
	// hive-26cd567b35c04a0a90f017388207b2
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// The task status.
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 6c8234f4-d1e1-4cea-b08b-7926fbdea144
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The task status.
	TaskStateMessage *string `json:"TaskStateMessage,omitempty" xml:"TaskStateMessage,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 1780387494000
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The Comfy workflow ID used by the task.
	//
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DescribeComfyTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeComfyTasksResponseBodyTasks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeComfyTasksResponseBodyTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeComfyTasksResponseBodyTasks) GetHiveId() *string {
	return s.HiveId
}

func (s *DescribeComfyTasksResponseBodyTasks) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *DescribeComfyTasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeComfyTasksResponseBodyTasks) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeComfyTasksResponseBodyTasks) GetTaskStateMessage() *string {
	return s.TaskStateMessage
}

func (s *DescribeComfyTasksResponseBodyTasks) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *DescribeComfyTasksResponseBodyTasks) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DescribeComfyTasksResponseBodyTasks) SetCreationTime(v string) *DescribeComfyTasksResponseBodyTasks {
	s.CreationTime = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetEndTime(v string) *DescribeComfyTasksResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetHiveId(v string) *DescribeComfyTasksResponseBodyTasks {
	s.HiveId = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetTaskErrorMessage(v string) *DescribeComfyTasksResponseBodyTasks {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetTaskId(v string) *DescribeComfyTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetTaskState(v string) *DescribeComfyTasksResponseBodyTasks {
	s.TaskState = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetTaskStateMessage(v string) *DescribeComfyTasksResponseBodyTasks {
	s.TaskStateMessage = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetUpdatedTime(v string) *DescribeComfyTasksResponseBodyTasks {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) SetWorkflowId(v string) *DescribeComfyTasksResponseBodyTasks {
	s.WorkflowId = &v
	return s
}

func (s *DescribeComfyTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
