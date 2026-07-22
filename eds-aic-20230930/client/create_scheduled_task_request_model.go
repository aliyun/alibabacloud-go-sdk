// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateScheduledTaskRequest
	GetCronExpression() *string
	SetInstanceIds(v []*string) *CreateScheduledTaskRequest
	GetInstanceIds() []*string
	SetMaxExecutions(v int32) *CreateScheduledTaskRequest
	GetMaxExecutions() *int32
	SetRunConfig(v *CreateScheduledTaskRequestRunConfig) *CreateScheduledTaskRequest
	GetRunConfig() *CreateScheduledTaskRequestRunConfig
	SetTaskName(v string) *CreateScheduledTaskRequest
	GetTaskName() *string
	SetUserPrompt(v string) *CreateScheduledTaskRequest
	GetUserPrompt() *string
}

type CreateScheduledTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0 0 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["ai-instance-001"]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaxExecutions *int32 `json:"MaxExecutions,omitempty" xml:"MaxExecutions,omitempty"`
	// example:
	//
	// {"maxSteps":10,"timeoutSeconds":3600}
	RunConfig *CreateScheduledTaskRequestRunConfig `json:"RunConfig,omitempty" xml:"RunConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 自动回复钉钉消息
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 打开钉钉，回复前5个未读消息
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateScheduledTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateScheduledTaskRequest) GetMaxExecutions() *int32 {
	return s.MaxExecutions
}

func (s *CreateScheduledTaskRequest) GetRunConfig() *CreateScheduledTaskRequestRunConfig {
	return s.RunConfig
}

func (s *CreateScheduledTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateScheduledTaskRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *CreateScheduledTaskRequest) SetCronExpression(v string) *CreateScheduledTaskRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetInstanceIds(v []*string) *CreateScheduledTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateScheduledTaskRequest) SetMaxExecutions(v int32) *CreateScheduledTaskRequest {
	s.MaxExecutions = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRunConfig(v *CreateScheduledTaskRequestRunConfig) *CreateScheduledTaskRequest {
	s.RunConfig = v
	return s
}

func (s *CreateScheduledTaskRequest) SetTaskName(v string) *CreateScheduledTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetUserPrompt(v string) *CreateScheduledTaskRequest {
	s.UserPrompt = &v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	if s.RunConfig != nil {
		if err := s.RunConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScheduledTaskRequestRunConfig struct {
	// example:
	//
	// {"batchSize":"1000"}
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	// example:
	//
	// 10
	MaxSteps *int32 `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateScheduledTaskRequestRunConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestRunConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestRunConfig) GetExtraParams() *string {
	return s.ExtraParams
}

func (s *CreateScheduledTaskRequestRunConfig) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *CreateScheduledTaskRequestRunConfig) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateScheduledTaskRequestRunConfig) SetExtraParams(v string) *CreateScheduledTaskRequestRunConfig {
	s.ExtraParams = &v
	return s
}

func (s *CreateScheduledTaskRequestRunConfig) SetMaxSteps(v int32) *CreateScheduledTaskRequestRunConfig {
	s.MaxSteps = &v
	return s
}

func (s *CreateScheduledTaskRequestRunConfig) SetTimeoutSeconds(v int32) *CreateScheduledTaskRequestRunConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateScheduledTaskRequestRunConfig) Validate() error {
	return dara.Validate(s)
}
