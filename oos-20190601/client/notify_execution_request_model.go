// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *NotifyExecutionRequest
	GetExecutionId() *string
	SetExecutionStatus(v string) *NotifyExecutionRequest
	GetExecutionStatus() *string
	SetLoopItem(v string) *NotifyExecutionRequest
	GetLoopItem() *string
	SetNotifyNote(v string) *NotifyExecutionRequest
	GetNotifyNote() *string
	SetNotifyType(v string) *NotifyExecutionRequest
	GetNotifyType() *string
	SetParameters(v string) *NotifyExecutionRequest
	GetParameters() *string
	SetRegionId(v string) *NotifyExecutionRequest
	GetRegionId() *string
	SetTaskExecutionId(v string) *NotifyExecutionRequest
	GetTaskExecutionId() *string
	SetTaskExecutionIds(v string) *NotifyExecutionRequest
	GetTaskExecutionIds() *string
	SetTaskName(v string) *NotifyExecutionRequest
	GetTaskName() *string
}

type NotifyExecutionRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The state of the terminated execution. This parameter is valid if you set the NotifyType parameter to CompleteExecution.
	//
	// example:
	//
	// Success
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The items of the child node in the loop task.
	//
	// example:
	//
	// i-xxx
	LoopItem *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	// The description for the notification.
	//
	// example:
	//
	// Note
	NotifyNote *string `json:"NotifyNote,omitempty" xml:"NotifyNote,omitempty"`
	// The type of the notification. Valid values:
	//
	// 	- **ExecuteTask**: starts to run a specific task. This value is used if you perform debugging in the Debug mode. If you set this parameter to ExecuteTask, you also need to set the Parameters parameter.
	//
	// 	- **CancelTask**: cancels a current task. This value is used if you perform debugging in the Debug mode.
	//
	// 	- **CompleteExecution**: manually terminates an execution if you perform debugging in the Debug mode. You can specify the state of the terminated execution by using the **ExecutionStatus*	- parameter.
	//
	// 	- **Approve**: approves an execution. For example, you are aware of the risks of an operation task and agree to approve the execution.
	//
	// 	- **Reject**: rejects an execution. For example, you want to reject the execution of a high-risk operation task.
	//
	// 	- **RetryTask**: retries a failed task whose execution mode is Suspend upon Failure.
	//
	// 	- **SkipTask**: skips a failed task whose execution mode is Suspend upon Failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approve
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The parameters of the subsequent task. This parameter is valid if you set the NotifyType parameter to ExecuteTask.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region in which the execution resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The execution IDs of the tasks.
	//
	// example:
	//
	// ["exec-79c321c11003a97c","exec-79c321c11003aqw97cz"]
	TaskExecutionIds *string `json:"TaskExecutionIds,omitempty" xml:"TaskExecutionIds,omitempty"`
	// The name of the subsequent task.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s NotifyExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s NotifyExecutionRequest) GoString() string {
	return s.String()
}

func (s *NotifyExecutionRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *NotifyExecutionRequest) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *NotifyExecutionRequest) GetLoopItem() *string {
	return s.LoopItem
}

func (s *NotifyExecutionRequest) GetNotifyNote() *string {
	return s.NotifyNote
}

func (s *NotifyExecutionRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *NotifyExecutionRequest) GetParameters() *string {
	return s.Parameters
}

func (s *NotifyExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *NotifyExecutionRequest) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *NotifyExecutionRequest) GetTaskExecutionIds() *string {
	return s.TaskExecutionIds
}

func (s *NotifyExecutionRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *NotifyExecutionRequest) SetExecutionId(v string) *NotifyExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetExecutionStatus(v string) *NotifyExecutionRequest {
	s.ExecutionStatus = &v
	return s
}

func (s *NotifyExecutionRequest) SetLoopItem(v string) *NotifyExecutionRequest {
	s.LoopItem = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyNote(v string) *NotifyExecutionRequest {
	s.NotifyNote = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyType(v string) *NotifyExecutionRequest {
	s.NotifyType = &v
	return s
}

func (s *NotifyExecutionRequest) SetParameters(v string) *NotifyExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *NotifyExecutionRequest) SetRegionId(v string) *NotifyExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionId(v string) *NotifyExecutionRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionIds(v string) *NotifyExecutionRequest {
	s.TaskExecutionIds = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskName(v string) *NotifyExecutionRequest {
	s.TaskName = &v
	return s
}

func (s *NotifyExecutionRequest) Validate() error {
	return dara.Validate(s)
}
