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
	// Calling number.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// Job description.
	//
	// example:
	//
	// 修改后的作业组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Configuration parameters for flash SMS push, in JSON format, containing third-party flash SMS configuration information.
	//
	// templateId: Flash SMS template ID.
	//
	// configId: Flash SMS configuration ID.
	//
	// example:
	//
	// {"templateId":"10471","configId":"8037f524-6ff2-4dbe-bb28-f59234ea7a64"}
	FlashSmsExtras *string `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3edc0260-6f7c-4de4-8535-09372240618b
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Task status. Valid values:
	//
	// - Draft: Draft.
	//
	// - Paused: Paused.
	//
	// example:
	//
	// Draft
	JobGroupStatus *string `json:"JobGroupStatus,omitempty" xml:"JobGroupStatus,omitempty"`
	// Minimum concurrency guarantee
	//
	// When the job starts, at least N concurrent calls are guaranteed.
	//
	// The sum of minimum concurrency guarantees for jobs with the same priority must not exceed the instance\\"s maximum concurrency.
	//
	// If the minimum concurrency guarantee is set to 0, the system intelligently assigns available idle concurrency.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// Job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 修改后的作业组
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Job group priority. Valid values:
	//
	// - **Urgent**: Urgent job.
	//
	// - **Daily**: Daily job.
	//
	// example:
	//
	// Daily
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// List of redial calling numbers
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// Redial strategy JSON
	//
	// example:
	//
	// {"emptyNumberIgnore":false,"inArrearsIgnore":false,"outOfServiceIgnore":false}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// Optimal ringing duration.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// Scenario ID. (This is a legacy parameter; use ScriptId going forward.)
	//
	// > This parameter is now deprecated. You can pass any value, but it is recommended to keep it consistent with scriptId.
	//
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Scenario script ID.
	//
	// example:
	//
	// 5a3940ce-a12f-4222-9f0f-605a9b89ea7c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Task execution policy.
	//
	// - repeatBy: Recurrence type. Once (no recurrence), Day (daily recurrence), Week (weekly recurrence), Month (monthly recurrence).
	//
	// - startTime: Start time of the policy.
	//
	// - endTime: End time of the policy.
	//
	// - workingTime: Allowed outbound calling time segment.
	//
	// - maxAttemptsPerDay: Maximum number of call attempts per day for each number under this job.
	//
	// - minAttemptInterval: Minimum interval between retry calls for a number, in minutes.
	//
	// - routingStrategy: Number routing strategy. None (not specified), LocalFirst (prioritize numbers from the same city), LocalProvinceFirst (prioritize numbers from the same province).
	//
	// - repeatDays: Execution dates corresponding to the recurrence type. If repeatBy is Week, 0 represents Sunday and 1–6 represent Monday through Saturday respectively. If repeatBy is Month, values 1–31 represent the 1st through the 31st day of the month; months without the specified date will skip execution (for example, if the 30th is selected, February will not execute).
	//
	// - repeatable: Whether loop task is enabled, true/false.
	//
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
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
