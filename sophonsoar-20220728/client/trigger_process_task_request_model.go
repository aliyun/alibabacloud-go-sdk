// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerProcessTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *TriggerProcessTaskRequest
	GetActionType() *string
	SetTaskId(v string) *TriggerProcessTaskRequest
	GetTaskId() *string
}

type TriggerProcessTaskRequest struct {
	// The type of the action. Valid values:
	//
	// 	- **remove**: cancels blocking or isolation.
	//
	// 	- **retry**: submits the task again.
	//
	// This parameter is required.
	//
	// example:
	//
	// remove
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the handling task.
	//
	// >  You can call the [DescribeProcessTasks](~~DescribeProcessTasks~~) operation to query the IDs of handling tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15355xxxxxx82894882
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TriggerProcessTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerProcessTaskRequest) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskRequest) GetActionType() *string {
	return s.ActionType
}

func (s *TriggerProcessTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *TriggerProcessTaskRequest) SetActionType(v string) *TriggerProcessTaskRequest {
	s.ActionType = &v
	return s
}

func (s *TriggerProcessTaskRequest) SetTaskId(v string) *TriggerProcessTaskRequest {
	s.TaskId = &v
	return s
}

func (s *TriggerProcessTaskRequest) Validate() error {
	return dara.Validate(s)
}
