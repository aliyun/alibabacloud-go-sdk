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
	// The list of calling numbers. If not specified, all numbers bound to the instance are selected by default.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// The configuration parameters for flash SMS in JSON format, including third-party flash SMS configuration information.
	//
	// - templateId: the flash SMS template ID.
	//
	// - configId: the flash SMS configuration ID.
	//
	// - templateContent: the flash SMS content.
	//
	// > Obtain the value of templateContent from the corresponding flash SMS capability provider.
	//
	// example:
	//
	// {"templateId":"104xx","configId":"8037f524-6fxxxxx", "templateContent": "【智能外呼机器人】给您来电，敬请接听！"}
	FlashSmsExtras *string `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task description.
	//
	// example:
	//
	// 任务描述
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 第一个任务
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// The guaranteed concurrency value.
	//
	// - When the task starts, a minimum of N concurrent calls are guaranteed.
	//
	// - The sum of guaranteed concurrency values for tasks with the same priority cannot exceed the instance concurrency.
	//
	// - If the guaranteed concurrency value is set to 0, the system intelligently allocates idle concurrency.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// The job group priority. Valid values:
	//
	// - **Urgent**: urgent task.
	//
	// - **Daily**: daily task.
	//
	// example:
	//
	// Daily
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The list of redial calling numbers.
	RecallCallingNumber []*string `json:"RecallCallingNumber,omitempty" xml:"RecallCallingNumber,omitempty" type:"Repeated"`
	// The redial strategy in JSON format. Parameter values default to false.
	//
	// - **emptyNumberIgnore**: does not call nonexistent numbers.
	//
	// - **inArrearsIgnore**: does not call numbers with overdue payments.
	//
	// - **outOfServiceIgnore**: does not call numbers that are out of service.
	//
	// example:
	//
	// {"emptyNumberIgnore":true,"inArrearsIgnore":true,"outOfServiceIgnore":true}
	RecallStrategyJson *string `json:"RecallStrategyJson,omitempty" xml:"RecallStrategyJson,omitempty"`
	// The optimal ringing duration. Default value: 25.
	//
	// example:
	//
	// 25
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// b9ff4e88-65f9-4eb3-987c-11ba51f3f24d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The task execution strategy.
	//
	// - repeatBy: the repeat type. Valid values: Once (no repeat), Week (repeat weekly), and Month (repeat monthly).
	//
	// - startTime: the strategy start time for time-based execution.
	//
	// - endTime: the strategy end time for time-based execution.
	//
	// > The execution mode is determined as follows:
	//
	// > - If no strategy start time or end time is specified, the task is executed immediately.
	//
	// > - If a strategy time is specified, the task is executed based on the schedule. You must also specify the repeat type repeatBy.
	//
	// - workingTime: the time window during which outbound calls can be made.
	//
	// - maxAttemptsPerDay: the maximum number of call attempts per day for each number in the task.
	//
	// - minAttemptInterval: the retry interval for a number, in minutes.
	//
	// - routingStrategy: the number strategy. Valid values: None (not specified), LocalFirst (local city numbers preferred), and LocalProvinceFirst (local province numbers preferred).
	//
	// - repeatDays: the execution days corresponding to the repeat type. If RepeatBy is set to Week, 0 indicates Sunday and 1-6 indicate Monday through Saturday. If RepeatBy is set to Month, 1-31 indicate the 1st through 31st day. The task is not executed in months that do not have the specified day. For example, if the 30th is selected, the task is not executed in February.
	//
	// - repeatable: specifies whether to enable cyclic tasks. Valid values: true and false.
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
