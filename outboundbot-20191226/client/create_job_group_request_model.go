// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *CreateJobGroupRequest
	GetCallingNumber() []*string
	SetFlashSmsExtras(v string) *CreateJobGroupRequest
	GetFlashSmsExtras() *string
	SetInstanceId(v string) *CreateJobGroupRequest
	GetInstanceId() *string
	SetJobGroupDescription(v string) *CreateJobGroupRequest
	GetJobGroupDescription() *string
	SetJobGroupName(v string) *CreateJobGroupRequest
	GetJobGroupName() *string
	SetMinConcurrency(v int64) *CreateJobGroupRequest
	GetMinConcurrency() *int64
	SetPriority(v string) *CreateJobGroupRequest
	GetPriority() *string
	SetRecallCallingNumber(v []*string) *CreateJobGroupRequest
	GetRecallCallingNumber() []*string
	SetRecallStrategyJson(v string) *CreateJobGroupRequest
	GetRecallStrategyJson() *string
	SetRingingDuration(v int64) *CreateJobGroupRequest
	GetRingingDuration() *int64
	SetScenarioId(v string) *CreateJobGroupRequest
	GetScenarioId() *string
	SetScriptId(v string) *CreateJobGroupRequest
	GetScriptId() *string
	SetStrategyJson(v string) *CreateJobGroupRequest
	GetStrategyJson() *string
}

type CreateJobGroupRequest struct {
	CallingNumber  []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	FlashSmsExtras *string   `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// This parameter is required.
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// example:
	//
	// 1
	MinConcurrency      *int64    `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	Priority            *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"emptyNumberIgnore\\":true,\\"inArrearsIgnore\\":true,\\"outOfServiceIgnore\\":true}
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
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// {\\"maxAttemptsPerDay\\":1,\\"name\\":\\"fa16dc2e-f778-44ab-8f25-54b7901df82a\\",\\"startTime\\":1640157314127,\\"endTime\\":1640160914127,\\"minAttemptInterval\\":10}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s CreateJobGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateJobGroupRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *CreateJobGroupRequest) GetFlashSmsExtras() *string {
	return s.FlashSmsExtras
}

func (s *CreateJobGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateJobGroupRequest) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *CreateJobGroupRequest) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *CreateJobGroupRequest) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *CreateJobGroupRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateJobGroupRequest) GetRecallCallingNumber() []*string {
	return s.RecallCallingNumber
}

func (s *CreateJobGroupRequest) GetRecallStrategyJson() *string {
	return s.RecallStrategyJson
}

func (s *CreateJobGroupRequest) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *CreateJobGroupRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CreateJobGroupRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateJobGroupRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *CreateJobGroupRequest) SetCallingNumber(v []*string) *CreateJobGroupRequest {
	s.CallingNumber = v
	return s
}

func (s *CreateJobGroupRequest) SetFlashSmsExtras(v string) *CreateJobGroupRequest {
	s.FlashSmsExtras = &v
	return s
}

func (s *CreateJobGroupRequest) SetInstanceId(v string) *CreateJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateJobGroupRequest) SetJobGroupDescription(v string) *CreateJobGroupRequest {
	s.JobGroupDescription = &v
	return s
}

func (s *CreateJobGroupRequest) SetJobGroupName(v string) *CreateJobGroupRequest {
	s.JobGroupName = &v
	return s
}

func (s *CreateJobGroupRequest) SetMinConcurrency(v int64) *CreateJobGroupRequest {
	s.MinConcurrency = &v
	return s
}

func (s *CreateJobGroupRequest) SetPriority(v string) *CreateJobGroupRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobGroupRequest) SetRecallCallingNumber(v []*string) *CreateJobGroupRequest {
	s.RecallCallingNumber = v
	return s
}

func (s *CreateJobGroupRequest) SetRecallStrategyJson(v string) *CreateJobGroupRequest {
	s.RecallStrategyJson = &v
	return s
}

func (s *CreateJobGroupRequest) SetRingingDuration(v int64) *CreateJobGroupRequest {
	s.RingingDuration = &v
	return s
}

func (s *CreateJobGroupRequest) SetScenarioId(v string) *CreateJobGroupRequest {
	s.ScenarioId = &v
	return s
}

func (s *CreateJobGroupRequest) SetScriptId(v string) *CreateJobGroupRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateJobGroupRequest) SetStrategyJson(v string) *CreateJobGroupRequest {
	s.StrategyJson = &v
	return s
}

func (s *CreateJobGroupRequest) Validate() error {
	return dara.Validate(s)
}
