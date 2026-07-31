// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateScheduledTaskShrinkRequest
	GetCronExpression() *string
	SetInstanceIds(v []*string) *CreateScheduledTaskShrinkRequest
	GetInstanceIds() []*string
	SetMaxExecutions(v int32) *CreateScheduledTaskShrinkRequest
	GetMaxExecutions() *int32
	SetRunConfigShrink(v string) *CreateScheduledTaskShrinkRequest
	GetRunConfigShrink() *string
	SetTaskName(v string) *CreateScheduledTaskShrinkRequest
	GetTaskName() *string
	SetUserPrompt(v string) *CreateScheduledTaskShrinkRequest
	GetUserPrompt() *string
}

type CreateScheduledTaskShrinkRequest struct {
	// The cron expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 0 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The list of instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["ai-instance-001"]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of executions.
	//
	// example:
	//
	// 1
	MaxExecutions *int32 `json:"MaxExecutions,omitempty" xml:"MaxExecutions,omitempty"`
	// The run configuration.
	//
	// example:
	//
	// {"maxSteps":10,"timeoutSeconds":3600}
	RunConfigShrink *string `json:"RunConfig,omitempty" xml:"RunConfig,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Auto-reply to DingTalk messages.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The user prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// Open DingTalk and reply to the first 5 unread messages.
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s CreateScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskShrinkRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateScheduledTaskShrinkRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateScheduledTaskShrinkRequest) GetMaxExecutions() *int32 {
	return s.MaxExecutions
}

func (s *CreateScheduledTaskShrinkRequest) GetRunConfigShrink() *string {
	return s.RunConfigShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateScheduledTaskShrinkRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *CreateScheduledTaskShrinkRequest) SetCronExpression(v string) *CreateScheduledTaskShrinkRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetInstanceIds(v []*string) *CreateScheduledTaskShrinkRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetMaxExecutions(v int32) *CreateScheduledTaskShrinkRequest {
	s.MaxExecutions = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetRunConfigShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.RunConfigShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTaskName(v string) *CreateScheduledTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetUserPrompt(v string) *CreateScheduledTaskShrinkRequest {
	s.UserPrompt = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
