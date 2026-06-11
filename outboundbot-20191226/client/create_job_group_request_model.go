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
	// List of caller numbers. If not specified, all numbers attached to the instance are selected by default.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// Configuration parameters for flash SMS push, in JSON format, containing third-party flash SMS configuration information.
	//
	// - templateId: Flash SMS Template ID.
	//
	// - configId: Flash SMS configuration ID.
	//
	// - templateContent: Flash SMS Content.
	//
	// > Obtain the value of templateContent from the partner providing the flash SMS capability.
	//
	// example:
	//
	// {"templateId":"104xx","configId":"8037f524-6fxxxxx", "templateContent": "【智能外呼机器人】给您来电，敬请接听！"}
	FlashSmsExtras *string `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task description.
	//
	// example:
	//
	// 任务描述
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// Task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 第一个任务
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// Concurrent guarantee value.
	//
	// - When a job starts, it is guaranteed a minimum of N concurrent executions.
	//
	// - The sum of concurrent guarantee values for jobs with the same priority must not exceed the instance concurrency limit.
	//
	// - If the concurrent guarantee value is configured as 0, the system intelligently assigns idle concurrency resources.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
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
	// List of caller numbers for redial.
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// Redial policy in JSON format. The default value of parameters in the JSON is false.
	//
	// - **emptyNumberIgnore**: Do not make outbound calls to nonexistent numbers.
	//
	// - **inArrearsIgnore**: Do not make outbound calls for overdue payments.
	//
	// - **outOfServiceIgnore**: Do not make outbound calls to out-of-service numbers.
	//
	// example:
	//
	// {"emptyNumberIgnore":true,"inArrearsIgnore":true,"outOfServiceIgnore":true}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// Optimal ringing duration. Default value is 25.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// Deprecated
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Scenario ID.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Job execution policy.
	//
	// - repeatBy: Recurrence type. Options are Once (no recurrence), Week (weekly recurrence), and Month (monthly recurrence).
	//
	// - startTime: Policy start time for time-based execution.
	//
	// - endTime: Policy end time for time-based execution.
	//
	// > Execution modes are as follows:
	//
	// > - If no start or end time is specified, the job executes immediately.
	//
	// > - If start and end times are provided, the job executes based on the schedule, and a recurrence type (repeatBy) must be selected.
	//
	// - workingTime: Allowed outbound calling time segment.
	//
	// - maxAttemptsPerDay: Maximum number of call attempts per day for numbers under this job.
	//
	// - minAttemptInterval: Minimum time interval between retry calls for a number, in minutes.
	//
	// - routingStrategy: Number routing strategy. Options are None (not specified), LocalFirst (local city numbers prioritized), and LocalProvinceFirst (local province numbers prioritized).
	//
	// - repeatDays: Execution dates corresponding to the recurrence type. If repeatBy is Week, 0 represents Sunday and 1–6 represent Monday through Saturday. If repeatBy is Month, values 1–31 represent the 1st through the 31st day of the month; months without the specified date skip execution (for example, if the 30th is selected, February skips execution).
	//
	// - repeatable: Whether loop task is enabled, true/false.
	//
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
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
