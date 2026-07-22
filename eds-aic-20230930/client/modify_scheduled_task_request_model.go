// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyScheduledTaskRequest
	GetCronExpression() *string
	SetInstanceIds(v []*string) *ModifyScheduledTaskRequest
	GetInstanceIds() []*string
	SetRunConfig(v *ModifyScheduledTaskRequestRunConfig) *ModifyScheduledTaskRequest
	GetRunConfig() *ModifyScheduledTaskRequestRunConfig
	SetScheduledId(v string) *ModifyScheduledTaskRequest
	GetScheduledId() *string
	SetStatus(v string) *ModifyScheduledTaskRequest
	GetStatus() *string
	SetTaskName(v string) *ModifyScheduledTaskRequest
	GetTaskName() *string
	SetTaskVersion(v int32) *ModifyScheduledTaskRequest
	GetTaskVersion() *int32
	SetUserPrompt(v string) *ModifyScheduledTaskRequest
	GetUserPrompt() *string
}

type ModifyScheduledTaskRequest struct {
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
	RunConfig *ModifyScheduledTaskRequestRunConfig `json:"RunConfig,omitempty" xml:"RunConfig,omitempty" type:"Struct"`
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

func (s ModifyScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyScheduledTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyScheduledTaskRequest) GetRunConfig() *ModifyScheduledTaskRequestRunConfig {
	return s.RunConfig
}

func (s *ModifyScheduledTaskRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ModifyScheduledTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyScheduledTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyScheduledTaskRequest) GetTaskVersion() *int32 {
	return s.TaskVersion
}

func (s *ModifyScheduledTaskRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *ModifyScheduledTaskRequest) SetCronExpression(v string) *ModifyScheduledTaskRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetInstanceIds(v []*string) *ModifyScheduledTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRunConfig(v *ModifyScheduledTaskRequestRunConfig) *ModifyScheduledTaskRequest {
	s.RunConfig = v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetStatus(v string) *ModifyScheduledTaskRequest {
	s.Status = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTaskName(v string) *ModifyScheduledTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTaskVersion(v int32) *ModifyScheduledTaskRequest {
	s.TaskVersion = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetUserPrompt(v string) *ModifyScheduledTaskRequest {
	s.UserPrompt = &v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	if s.RunConfig != nil {
		if err := s.RunConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyScheduledTaskRequestRunConfig struct {
	// The extended parameters as a JSON string.
	//
	// example:
	//
	// {"batchSize":"1000"}
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// The maximum number of execution steps.
	//
	// example:
	//
	// 10
	MaxSteps *int32 `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	// The timeout in seconds.
	//
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s ModifyScheduledTaskRequestRunConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequestRunConfig) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequestRunConfig) GetExtraParams() *string {
	return s.ExtraParams
}

func (s *ModifyScheduledTaskRequestRunConfig) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *ModifyScheduledTaskRequestRunConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *ModifyScheduledTaskRequestRunConfig) SetExtraParams(v string) *ModifyScheduledTaskRequestRunConfig {
	s.ExtraParams = &v
	return s
}

func (s *ModifyScheduledTaskRequestRunConfig) SetMaxSteps(v int32) *ModifyScheduledTaskRequestRunConfig {
	s.MaxSteps = &v
	return s
}

func (s *ModifyScheduledTaskRequestRunConfig) SetTimeoutSeconds(v int32) *ModifyScheduledTaskRequestRunConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *ModifyScheduledTaskRequestRunConfig) Validate() error {
	return dara.Validate(s)
}
