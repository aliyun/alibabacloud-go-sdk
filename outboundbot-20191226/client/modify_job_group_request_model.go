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
	// The calling numbers for the job group.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// The description of the job group.
	//
	// example:
	//
	// 修改后的作业组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The flash SMS configuration, specified as a JSON string. This may include settings for third-party flash SMS services.
	//
	// `templateId`: The flash SMS template ID.<br>
	//
	// `configId`: The flash SMS configuration ID.<br>
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
	// The job group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3edc0260-6f7c-4de4-8535-09372240618b
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The status of the job group. Valid values:
	//
	// - `Draft`: The job group is in a draft state.
	//
	// - `Paused`: The job group is paused.
	//
	// example:
	//
	// Draft
	JobGroupStatus *string `json:"JobGroupStatus,omitempty" xml:"JobGroupStatus,omitempty"`
	// The guaranteed minimum number of concurrent calls for the job group. The sum of this value for all job groups with the same priority cannot exceed the instance\\"s total concurrency. If you set this parameter to `0`, the system dynamically allocates available lines from a shared pool.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// The name of the job group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 修改后的作业组
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the job group. Valid values:
	//
	// - **Urgent**: An urgent job.
	//
	// - **Daily**: A routine job.
	//
	// example:
	//
	// Daily
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The recall calling numbers.
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// A JSON string that defines the recall strategy.
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
	// The scenario ID. (This is a legacy parameter. Use `ScriptId` instead.)
	//
	// > You can pass any value for this legacy parameter, but we recommend using the same value as `ScriptId` for consistency.
	//
	// example:
	//
	// c6a668d1-3145-4048-9101-cb3678bb8884
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the script for the scenario.
	//
	// example:
	//
	// 5a3940ce-a12f-4222-9f0f-605a9b89ea7c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// A JSON string that defines the execution strategy for the job group.
	//
	// - `id`: The ID of the job group strategy. You can obtain this ID from the `StrategyId` parameter returned by the `DescribeJobGroup` operation.
	//
	// - `repeatBy`: The frequency of the job. Valid values: `Once` (does not repeat), `Day` (repeats daily), `Week` (repeats weekly), and `Month` (repeats monthly).
	//
	// - `startTime`: The start time of the strategy.
	//
	// - `endTime`: The end time of the strategy.
	//
	// - `workingTime`: The time windows for making outbound calls.
	//
	// - `maxAttemptsPerDay`: The maximum number of call attempts per day for a number in this job group.
	//
	// - `minAttemptInterval`: The minimum interval, in minutes, between call retries to the same number.
	//
	// - `routingStrategy`: The number routing strategy. Valid values: `None` (not specified), `LocalFirst` (prioritizes numbers in the same city), and `LocalProvinceFirst` (prioritizes numbers in the same province).
	//
	// - `repeatDays`: The specific days to run the job, based on the `repeatBy` value. If `repeatBy` is `Week`, `0` indicates Sunday, and `1` through `6` indicate Monday through Saturday. If `repeatBy` is `Month`, valid values are `1` through `31`. If a month does not have the specified day (for example, day 30 in February), the job does not run on that day.
	//
	// - `repeatable`: Whether the job is recurring. Valid values are `true` and `false`.
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
