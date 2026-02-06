// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *ModifyJobGroupRequest
	GetCallingNumber() []*string
	SetDescription(v string) *ModifyJobGroupRequest
	GetDescription() *string
	SetFlashSmsExtras(v string) *ModifyJobGroupRequest
	GetFlashSmsExtras() *string
	SetInstanceId(v string) *ModifyJobGroupRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *ModifyJobGroupRequest
	GetJobGroupId() *string
	SetJobGroupStatus(v string) *ModifyJobGroupRequest
	GetJobGroupStatus() *string
	SetMinConcurrency(v int64) *ModifyJobGroupRequest
	GetMinConcurrency() *int64
	SetName(v string) *ModifyJobGroupRequest
	GetName() *string
	SetPriority(v string) *ModifyJobGroupRequest
	GetPriority() *string
	SetRecallCallingNumber(v []*string) *ModifyJobGroupRequest
	GetRecallCallingNumber() []*string
	SetRecallStrategyJson(v string) *ModifyJobGroupRequest
	GetRecallStrategyJson() *string
	SetRingingDuration(v int64) *ModifyJobGroupRequest
	GetRingingDuration() *int64
	SetScenarioId(v string) *ModifyJobGroupRequest
	GetScenarioId() *string
	SetScriptId(v string) *ModifyJobGroupRequest
	GetScriptId() *string
	SetStrategyJson(v string) *ModifyJobGroupRequest
	GetStrategyJson() *string
}

type ModifyJobGroupRequest struct {
	CallingNumber  []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	FlashSmsExtras *string   `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3edc0260-6f7c-4de4-8535-09372240618b
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// Draft
	JobGroupStatus *string `json:"JobGroupStatus,omitempty" xml:"JobGroupStatus,omitempty"`
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// This parameter is required.
	Name                *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority            *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"emptyNumberIgnore\\":false,\\"inArrearsIgnore\\":false,\\"outOfServiceIgnore\\":false}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// example:
	//
	// 5a3940ce-a12f-4222-9f0f-605a9b89ea7c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// {\\"maxAttemptsPerDay\\":\\"0\\",\\"minAttemptInterval\\":\\"5\\",\\"Id\\":\\"689fc584-7f9f-4dc2-933d-8711beef7b15\\"}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s ModifyJobGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *ModifyJobGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyJobGroupRequest) GetFlashSmsExtras() *string {
	return s.FlashSmsExtras
}

func (s *ModifyJobGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyJobGroupRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ModifyJobGroupRequest) GetJobGroupStatus() *string {
	return s.JobGroupStatus
}

func (s *ModifyJobGroupRequest) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *ModifyJobGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyJobGroupRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifyJobGroupRequest) GetRecallCallingNumber() []*string {
	return s.RecallCallingNumber
}

func (s *ModifyJobGroupRequest) GetRecallStrategyJson() *string {
	return s.RecallStrategyJson
}

func (s *ModifyJobGroupRequest) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *ModifyJobGroupRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ModifyJobGroupRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyJobGroupRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *ModifyJobGroupRequest) SetCallingNumber(v []*string) *ModifyJobGroupRequest {
	s.CallingNumber = v
	return s
}

func (s *ModifyJobGroupRequest) SetDescription(v string) *ModifyJobGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyJobGroupRequest) SetFlashSmsExtras(v string) *ModifyJobGroupRequest {
	s.FlashSmsExtras = &v
	return s
}

func (s *ModifyJobGroupRequest) SetInstanceId(v string) *ModifyJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetJobGroupId(v string) *ModifyJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetJobGroupStatus(v string) *ModifyJobGroupRequest {
	s.JobGroupStatus = &v
	return s
}

func (s *ModifyJobGroupRequest) SetMinConcurrency(v int64) *ModifyJobGroupRequest {
	s.MinConcurrency = &v
	return s
}

func (s *ModifyJobGroupRequest) SetName(v string) *ModifyJobGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyJobGroupRequest) SetPriority(v string) *ModifyJobGroupRequest {
	s.Priority = &v
	return s
}

func (s *ModifyJobGroupRequest) SetRecallCallingNumber(v []*string) *ModifyJobGroupRequest {
	s.RecallCallingNumber = v
	return s
}

func (s *ModifyJobGroupRequest) SetRecallStrategyJson(v string) *ModifyJobGroupRequest {
	s.RecallStrategyJson = &v
	return s
}

func (s *ModifyJobGroupRequest) SetRingingDuration(v int64) *ModifyJobGroupRequest {
	s.RingingDuration = &v
	return s
}

func (s *ModifyJobGroupRequest) SetScenarioId(v string) *ModifyJobGroupRequest {
	s.ScenarioId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetScriptId(v string) *ModifyJobGroupRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyJobGroupRequest) SetStrategyJson(v string) *ModifyJobGroupRequest {
	s.StrategyJson = &v
	return s
}

func (s *ModifyJobGroupRequest) Validate() error {
	return dara.Validate(s)
}
