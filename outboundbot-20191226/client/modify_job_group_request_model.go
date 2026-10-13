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
	// The caller number.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// The job description.
	//
	// example:
	//
	// Modified job group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration parameters for flash message push in JSON format, including the configuration information of third-party flash messages.
	//
	// templateId: the flash message template ID.
	//
	// configId: the flash message configuration ID.
	//
	// example:
	//
	// {"templateId":"10471","configId":"8037f524-6ff2-4dbe-bb28-f59234ea7a64"}
	FlashSmsExtras *string `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3edc0260-6f7c-4de4-8535-09372240618b
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The job status. Valid values:
	//
	// - Draft: draft.
	//
	// - Paused: paused.
	//
	// example:
	//
	// Draft
	JobGroupStatus *string `json:"JobGroupStatus,omitempty" xml:"JobGroupStatus,omitempty"`
	// The guaranteed concurrency.
	//
	// When a job starts, a minimum of N concurrent calls are guaranteed.
	//
	// The sum of guaranteed concurrency values for jobs of the same priority cannot exceed the instance concurrency.
	//
	// If the guaranteed concurrency is set to 0, the system intelligently allocates idle concurrency.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// The job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Modified job group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The job group priority. Valid values:
	//
	// - **Urgent**: urgent job.
	//
	// - **Daily**: daily job.
	//
	// example:
	//
	// Daily
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The list of caller numbers for recall.
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// The JSON string of the recall strategy.
	//
	// example:
	//
	// {"emptyNumberIgnore":false,"inArrearsIgnore":false,"outOfServiceIgnore":false}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// The optimal ringing duration.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// The scenario ID. This is a legacy parameter. Use ScriptId instead.
	//
	// > This parameter is a legacy parameter. You can pass any value. Keep it consistent with the ScriptId.
	//
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 5a3940ce-a12f-4222-9f0f-605a9b89ea7c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The job execution strategy.
	//
	// - id: the job group strategy ID. You can obtain it from the StrategyId response parameter of the DescribeJobGroup operation.
	//
	// - repeatBy: the recurrence type. Valid values: Once (no recurrence), Day (daily), Week (weekly), and Month (monthly).
	//
	// - startTime: the start time of the strategy.
	//
	// - endTime: the end time of the strategy.
	//
	// - workingTime: the time period during which outbound calls are allowed.
	//
	// - maxAttemptsPerDay: the maximum number of call attempts per day for a phone number under this job.
	//
	// - minAttemptInterval: the retry interval for a phone number, in minutes.
	//
	// - routingStrategy: the phone number strategy. Valid values: None (not specified), LocalFirst (local numbers first), and LocalProvinceFirst (local province numbers first).
	//
	// - repeatDays: the execution days corresponding to the recurrence type. If RepeatBy is Week, 0 indicates Sunday, and 1 to 6 indicate Monday to Saturday. If RepeatBy is Month, 1 to 31 indicate the 1st to the 31st day of the month. If a month does not have the specified day, the job is not executed in that month. For example, if the 30th is selected, the job is not executed in February.
	//
	// - repeatable: specifies whether to enable recurring jobs. Valid values: true and false.
	//
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},"id":"cdca9ed7-6470-42d0-afb3-a41e08b41383",{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
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
