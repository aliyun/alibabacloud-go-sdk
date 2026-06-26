// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *RunAgentTaskRequest
	GetBizRegionId() *string
	SetInstanceIds(v []*string) *RunAgentTaskRequest
	GetInstanceIds() []*string
	SetMaxSteps(v int32) *RunAgentTaskRequest
	GetMaxSteps() *int32
	SetScheduleId(v string) *RunAgentTaskRequest
	GetScheduleId() *string
	SetTaskConfigId(v string) *RunAgentTaskRequest
	GetTaskConfigId() *string
	SetTimeoutSeconds(v int32) *RunAgentTaskRequest
	GetTimeoutSeconds() *int32
	SetUserPrompt(v string) *RunAgentTaskRequest
	GetUserPrompt() *string
}

type RunAgentTaskRequest struct {
	// The region ID of the Mobile node.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The list of Mobile node IDs. A maximum of 100 nodes are supported per request.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of execution steps for the task to prevent infinite loops. Valid values: 30 to 1000. Default value: 1000.
	//
	// example:
	//
	// 30
	MaxSteps     *int32  `json:"MaxSteps,omitempty" xml:"MaxSteps,omitempty"`
	ScheduleId   *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	TaskConfigId *string `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
	// The task timeout period in seconds. Valid values: 300 to 3600. Default value: 3600.
	//
	// example:
	//
	// 3600
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// The user instruction in natural language. The Agent performs operations based on this instruction.
	//
	// example:
	//
	// Download DingTalk from App Store
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s RunAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *RunAgentTaskRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *RunAgentTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RunAgentTaskRequest) GetMaxSteps() *int32 {
	return s.MaxSteps
}

func (s *RunAgentTaskRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *RunAgentTaskRequest) GetTaskConfigId() *string {
	return s.TaskConfigId
}

func (s *RunAgentTaskRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *RunAgentTaskRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *RunAgentTaskRequest) SetBizRegionId(v string) *RunAgentTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *RunAgentTaskRequest) SetInstanceIds(v []*string) *RunAgentTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *RunAgentTaskRequest) SetMaxSteps(v int32) *RunAgentTaskRequest {
	s.MaxSteps = &v
	return s
}

func (s *RunAgentTaskRequest) SetScheduleId(v string) *RunAgentTaskRequest {
	s.ScheduleId = &v
	return s
}

func (s *RunAgentTaskRequest) SetTaskConfigId(v string) *RunAgentTaskRequest {
	s.TaskConfigId = &v
	return s
}

func (s *RunAgentTaskRequest) SetTimeoutSeconds(v int32) *RunAgentTaskRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *RunAgentTaskRequest) SetUserPrompt(v string) *RunAgentTaskRequest {
	s.UserPrompt = &v
	return s
}

func (s *RunAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
