// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *RunAgentTaskShrinkRequest
	GetBizRegionId() *string
	SetInstanceIds(v []*string) *RunAgentTaskShrinkRequest
	GetInstanceIds() []*string
	SetMaxSteps(v int32) *RunAgentTaskShrinkRequest
	GetMaxSteps() *int32
	SetRunConfigShrink(v string) *RunAgentTaskShrinkRequest
	GetRunConfigShrink() *string
	SetSaveArtifacts(v bool) *RunAgentTaskShrinkRequest
	GetSaveArtifacts() *bool
	SetScheduleId(v string) *RunAgentTaskShrinkRequest
	GetScheduleId() *string
	SetTargets(v []*RunAgentTaskShrinkRequestTargets) *RunAgentTaskShrinkRequest
	GetTargets() []*RunAgentTaskShrinkRequestTargets
	SetTaskConfigId(v string) *RunAgentTaskShrinkRequest
	GetTaskConfigId() *string
	SetTimeoutSeconds(v int32) *RunAgentTaskShrinkRequest
	GetTimeoutSeconds() *int32
	SetUserPrompt(v string) *RunAgentTaskShrinkRequest
	GetUserPrompt() *string
}

type RunAgentTaskShrinkRequest struct {
	// The ID of the region where the Mobile node resides.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The list of Mobile node IDs. A single request supports up to 100 nodes.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of execution steps for the task to prevent infinite loops. Valid values: 50 to 1000. Default value: 50.
	//
	// example:
	//
	// 50
	MaxSteps *int32 `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	// The runtime configuration that contains the runtime parameters such as skills for the current task.
	//
	// example:
	//
	// {"Skills":["sk-abc","sk-def"]}
	RunConfigShrink *string `json:"RunConfig,omitempty" xml:"RunConfig,omitempty"`
	// Specifies whether to save the task artifacts to your OSS.
	//
	// example:
	//
	// true
	SaveArtifacts *bool `json:"SaveArtifacts,omitempty" xml:"SaveArtifacts,omitempty"`
	// The ID of the scheduling plan. When specified, the execute record is associated with the corresponding scheduled node, which facilitates aggregate query by scheduling aggregation dimension.
	//
	// example:
	//
	// sch-260625-pbj2****
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The list of targets. Each element is an object that contains an InstanceId and a SessionId.
	Targets []*RunAgentTaskShrinkRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The ID of the task configuration. This parameter is used to trigger a task with the specified configuration.
	//
	// example:
	//
	// tsk-260625-49be****
	TaskConfigId *string `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
	// The timeout period of the task in seconds. Valid values: 300 to 3600. Default value: 3600.
	//
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// The user prompt in natural language. The agent performs operations based on this prompt.
	//
	// example:
	//
	// Download DingTalk from the app store
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s RunAgentTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunAgentTaskShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *RunAgentTaskShrinkRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunAgentTaskShrinkRequest) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *RunAgentTaskShrinkRequest) GetRunConfigShrink() *string {
	return s.RunConfigShrink
}

func (s *RunAgentTaskShrinkRequest) GetSaveArtifacts() *bool {
	return s.SaveArtifacts
}

func (s *RunAgentTaskShrinkRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *RunAgentTaskShrinkRequest) GetTargets() []*RunAgentTaskShrinkRequestTargets {
	return s.Targets
}

func (s *RunAgentTaskShrinkRequest) GetTaskConfigId() *string {
	return s.TaskConfigId
}

func (s *RunAgentTaskShrinkRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *RunAgentTaskShrinkRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *RunAgentTaskShrinkRequest) SetBizRegionId(v string) *RunAgentTaskShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetInstanceIds(v []*string) *RunAgentTaskShrinkRequest {
	s.InstanceIds = v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetMaxSteps(v int32) *RunAgentTaskShrinkRequest {
	s.MaxSteps = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetRunConfigShrink(v string) *RunAgentTaskShrinkRequest {
	s.RunConfigShrink = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetSaveArtifacts(v bool) *RunAgentTaskShrinkRequest {
	s.SaveArtifacts = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetScheduleId(v string) *RunAgentTaskShrinkRequest {
	s.ScheduleId = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetTargets(v []*RunAgentTaskShrinkRequestTargets) *RunAgentTaskShrinkRequest {
	s.Targets = v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetTaskConfigId(v string) *RunAgentTaskShrinkRequest {
	s.TaskConfigId = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetTimeoutSeconds(v int32) *RunAgentTaskShrinkRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) SetUserPrompt(v string) *RunAgentTaskShrinkRequest {
	s.UserPrompt = &v
	return s
}

func (s *RunAgentTaskShrinkRequest) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunAgentTaskShrinkRequestTargets struct {
	// The ID of the Mobile node, such as acp-xxx.
	//
	// example:
	//
	// acp-5hh4a31emkt6u****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The session ID. Tasks with the same session ID share the context, such as ses-260702-21b****.
	//
	// example:
	//
	// ses-260702-21bh****。
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunAgentTaskShrinkRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskShrinkRequestTargets) GoString() string {
	return s.String()
}

func (s *RunAgentTaskShrinkRequestTargets) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunAgentTaskShrinkRequestTargets) GetSessionId() *string {
	return s.SessionId
}

func (s *RunAgentTaskShrinkRequestTargets) SetInstanceId(v string) *RunAgentTaskShrinkRequestTargets {
	s.InstanceId = &v
	return s
}

func (s *RunAgentTaskShrinkRequestTargets) SetSessionId(v string) *RunAgentTaskShrinkRequestTargets {
	s.SessionId = &v
	return s
}

func (s *RunAgentTaskShrinkRequestTargets) Validate() error {
	return dara.Validate(s)
}
