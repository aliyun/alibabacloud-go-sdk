// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrentRate(v int32) *UpdateAiOutboundTaskRequest
	GetConcurrentRate() *int32
	SetDescription(v string) *UpdateAiOutboundTaskRequest
	GetDescription() *string
	SetExecutionTime(v string) *UpdateAiOutboundTaskRequest
	GetExecutionTime() *string
	SetForecastCallRate(v float32) *UpdateAiOutboundTaskRequest
	GetForecastCallRate() *float32
	SetHandlerId(v int64) *UpdateAiOutboundTaskRequest
	GetHandlerId() *int64
	SetInstanceId(v string) *UpdateAiOutboundTaskRequest
	GetInstanceId() *string
	SetName(v string) *UpdateAiOutboundTaskRequest
	GetName() *string
	SetNumRepeated(v int32) *UpdateAiOutboundTaskRequest
	GetNumRepeated() *int32
	SetOutboundNums(v []*string) *UpdateAiOutboundTaskRequest
	GetOutboundNums() []*string
	SetRecallRule(v *UpdateAiOutboundTaskRequestRecallRule) *UpdateAiOutboundTaskRequest
	GetRecallRule() *UpdateAiOutboundTaskRequestRecallRule
	SetTaskId(v int64) *UpdateAiOutboundTaskRequest
	GetTaskId() *int64
}

type UpdateAiOutboundTaskRequest struct {
	// The concurrency for automatic outbound calls.
	//
	// example:
	//
	// 10
	ConcurrentRate *int32 `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	// The task description. The description can be up to 100 characters in length.
	//
	// example:
	//
	// PropertySales
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The task execution time. Use a JSON object to configure time periods by day of the week. The start and end values must be in the HH:mm format. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"TUESDAY":[{"start":"06:00","end":"06:05"}],"MONDAY":[{"start":"09:00","end":"18:00"},{"start":"20:30","end":"21:45"},{"start":"22:30","end":"22:50"}],"WEDNESDAY":[{"start":"09:00","end":"18:00"}],"THURSDAY":[{"start":"09:00","end":"18:00"}],"FRIDAY":[{"start":"09:00","end":"18:00"}],"SATURDAY":[{"start":"09:00","end":"18:00"}],"SUNDAY":[{"start":"17:00","end":"23:45"}]}
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The fixed call ratio for predictive outbound calls. Valid values: **≥ 1**.
	//
	// example:
	//
	// 1.2
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// The skill group ID for predictive outbound calls, or the IVR ID for automatic outbound calls.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// The ID of the Artificial Intelligence Cloud Call Service (AICCS) instance.
	//
	// You can obtain the instance ID from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task name. The name must be 1 to 15 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleOutboundCall
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The called number deduplication policy. Valid values:
	//
	// - **0**: Deduplicate within the task.
	//
	// - **1**: Do not deduplicate within the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NumRepeated *int32 `json:"NumRepeated,omitempty" xml:"NumRepeated,omitempty"`
	// The outbound caller numbers.
	//
	// This parameter is required.
	OutboundNums []*string `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty" type:"Repeated"`
	// The failed recall policy.
	//
	// > If this parameter is left empty, the system does not recall when an outbound call fails.
	RecallRule *UpdateAiOutboundTaskRequestRecallRule `json:"RecallRule,omitempty" xml:"RecallRule,omitempty" type:"Struct"`
	// The task ID.
	//
	// You can call the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) operation and check the **Data*	- parameter in the response, or call the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) operation and check the **TaskId*	- parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiOutboundTaskRequest) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *UpdateAiOutboundTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAiOutboundTaskRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *UpdateAiOutboundTaskRequest) GetForecastCallRate() *float32 {
	return s.ForecastCallRate
}

func (s *UpdateAiOutboundTaskRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *UpdateAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAiOutboundTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAiOutboundTaskRequest) GetNumRepeated() *int32 {
	return s.NumRepeated
}

func (s *UpdateAiOutboundTaskRequest) GetOutboundNums() []*string {
	return s.OutboundNums
}

func (s *UpdateAiOutboundTaskRequest) GetRecallRule() *UpdateAiOutboundTaskRequestRecallRule {
	return s.RecallRule
}

func (s *UpdateAiOutboundTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateAiOutboundTaskRequest) SetConcurrentRate(v int32) *UpdateAiOutboundTaskRequest {
	s.ConcurrentRate = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetDescription(v string) *UpdateAiOutboundTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetExecutionTime(v string) *UpdateAiOutboundTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetForecastCallRate(v float32) *UpdateAiOutboundTaskRequest {
	s.ForecastCallRate = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetHandlerId(v int64) *UpdateAiOutboundTaskRequest {
	s.HandlerId = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetInstanceId(v string) *UpdateAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetName(v string) *UpdateAiOutboundTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetNumRepeated(v int32) *UpdateAiOutboundTaskRequest {
	s.NumRepeated = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetOutboundNums(v []*string) *UpdateAiOutboundTaskRequest {
	s.OutboundNums = v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetRecallRule(v *UpdateAiOutboundTaskRequestRecallRule) *UpdateAiOutboundTaskRequest {
	s.RecallRule = v
	return s
}

func (s *UpdateAiOutboundTaskRequest) SetTaskId(v int64) *UpdateAiOutboundTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAiOutboundTaskRequest) Validate() error {
	if s.RecallRule != nil {
		if err := s.RecallRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAiOutboundTaskRequestRecallRule struct {
	// The number of failed recall attempts. Valid values: **1 to 3**.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The interval between failed recall attempts. Valid values: **1 to 60**. Unit: minutes.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s UpdateAiOutboundTaskRequestRecallRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiOutboundTaskRequestRecallRule) GoString() string {
	return s.String()
}

func (s *UpdateAiOutboundTaskRequestRecallRule) GetCount() *int32 {
	return s.Count
}

func (s *UpdateAiOutboundTaskRequestRecallRule) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateAiOutboundTaskRequestRecallRule) SetCount(v int32) *UpdateAiOutboundTaskRequestRecallRule {
	s.Count = &v
	return s
}

func (s *UpdateAiOutboundTaskRequestRecallRule) SetInterval(v int32) *UpdateAiOutboundTaskRequestRecallRule {
	s.Interval = &v
	return s
}

func (s *UpdateAiOutboundTaskRequestRecallRule) Validate() error {
	return dara.Validate(s)
}
