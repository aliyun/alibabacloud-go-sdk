// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyScheduledTaskShrinkRequest
	GetCronExpression() *string
	SetInstanceIds(v []*string) *ModifyScheduledTaskShrinkRequest
	GetInstanceIds() []*string
	SetRunConfigShrink(v string) *ModifyScheduledTaskShrinkRequest
	GetRunConfigShrink() *string
	SetScheduledId(v string) *ModifyScheduledTaskShrinkRequest
	GetScheduledId() *string
	SetStatus(v string) *ModifyScheduledTaskShrinkRequest
	GetStatus() *string
	SetTaskName(v string) *ModifyScheduledTaskShrinkRequest
	GetTaskName() *string
	SetTaskVersion(v int32) *ModifyScheduledTaskShrinkRequest
	GetTaskVersion() *int32
	SetUserPrompt(v string) *ModifyScheduledTaskShrinkRequest
	GetUserPrompt() *string
}

type ModifyScheduledTaskShrinkRequest struct {
	// The cron expression.
	//
	// example:
	//
	// 0 30 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The list of instance IDs.
	//
	// example:
	//
	// ["acp-5hh431emkt6u*****"]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The run configuration.
	//
	// example:
	//
	// {"maxSteps":10,"timeoutSeconds":3600}
	RunConfigShrink *string `json:"RunConfig,omitempty" xml:"RunConfig,omitempty"`
	// The scheduled task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sch-260705-agb*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// The status switch: ACTIVE/DISABLED.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	//
	// example:
	//
	// NewTaskName.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The CAS version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskVersion *int32 `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	// The user prompt.
	//
	// example:
	//
	// Execute daily data synchronization task.
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s ModifyScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskShrinkRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyScheduledTaskShrinkRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyScheduledTaskShrinkRequest) GetRunConfigShrink() *string {
	return s.RunConfigShrink
}

func (s *ModifyScheduledTaskShrinkRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ModifyScheduledTaskShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyScheduledTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyScheduledTaskShrinkRequest) GetTaskVersion() *int32 {
	return s.TaskVersion
}

func (s *ModifyScheduledTaskShrinkRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *ModifyScheduledTaskShrinkRequest) SetCronExpression(v string) *ModifyScheduledTaskShrinkRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetInstanceIds(v []*string) *ModifyScheduledTaskShrinkRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetRunConfigShrink(v string) *ModifyScheduledTaskShrinkRequest {
	s.RunConfigShrink = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetScheduledId(v string) *ModifyScheduledTaskShrinkRequest {
	s.ScheduledId = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetStatus(v string) *ModifyScheduledTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetTaskName(v string) *ModifyScheduledTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetTaskVersion(v int32) *ModifyScheduledTaskShrinkRequest {
	s.TaskVersion = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) SetUserPrompt(v string) *ModifyScheduledTaskShrinkRequest {
	s.UserPrompt = &v
	return s
}

func (s *ModifyScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
