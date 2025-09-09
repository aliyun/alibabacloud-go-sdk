// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceLevelTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDrdsInstanceLevelTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsInstanceLevelTasksResponseBody
	GetSuccess() *bool
	SetTasks(v *DescribeDrdsInstanceLevelTasksResponseBodyTasks) *DescribeDrdsInstanceLevelTasksResponseBody
	GetTasks() *DescribeDrdsInstanceLevelTasksResponseBodyTasks
}

type DescribeDrdsInstanceLevelTasksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of returned unfinished tasks.
	Tasks *DescribeDrdsInstanceLevelTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) GetTasks() *DescribeDrdsInstanceLevelTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetRequestId(v string) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) SetTasks(v *DescribeDrdsInstanceLevelTasksResponseBodyTasks) *DescribeDrdsInstanceLevelTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasks struct {
	Task []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasks) GetTask() []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasks) SetTask(v []*DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) *DescribeDrdsInstanceLevelTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceLevelTasksResponseBodyTasksTask struct {
	// Indicates whether the task can be canceled.
	//
	// example:
	//
	// false
	AllowCancel *bool `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// The error message returned for the task.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The timestamp when the task is created.
	//
	// example:
	//
	// 1568705520000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The progress of the task. Valid values: 0 to 100.
	//
	// example:
	//
	// 99
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The description of the task progress.
	ProgressDescription *string `json:"ProgressDescription,omitempty" xml:"ProgressDescription,omitempty"`
	// Indicates whether the progress of the task is displayed.
	//
	// example:
	//
	// true
	ShowProgress *bool `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 12312
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// upgrade_instance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The phase of the task.
	//
	// example:
	//
	// 1
	TaskPhase *string `json:"TaskPhase,omitempty" xml:"TaskPhase,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- 0: The task is being executed.
	//
	// 	- 1: The task is executed.
	//
	// 	- 2: The task failed to be executed.
	//
	// 	- 3: The task is canceled.
	//
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// 11
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetAllowCancel() *bool {
	return s.AllowCancel
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetProgressDescription() *string {
	return s.ProgressDescription
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetShowProgress() *bool {
	return s.ShowProgress
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetTargetId() *int64 {
	return s.TargetId
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetTaskPhase() *string {
	return s.TaskPhase
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) GetTaskType() *int32 {
	return s.TaskType
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetAllowCancel(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.AllowCancel = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetErrMsg(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetGmtCreate(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetProgressDescription(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ProgressDescription = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetShowProgress(v bool) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.ShowProgress = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTargetId(v int64) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TargetId = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskName(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskName = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskPhase(v string) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskPhase = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskStatus(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) SetTaskType(v int32) *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

func (s *DescribeDrdsInstanceLevelTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
