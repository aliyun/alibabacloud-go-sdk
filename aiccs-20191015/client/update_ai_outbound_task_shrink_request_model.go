// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiOutboundTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrentRate(v int32) *UpdateAiOutboundTaskShrinkRequest
	GetConcurrentRate() *int32
	SetDescription(v string) *UpdateAiOutboundTaskShrinkRequest
	GetDescription() *string
	SetExecutionTime(v string) *UpdateAiOutboundTaskShrinkRequest
	GetExecutionTime() *string
	SetForecastCallRate(v float32) *UpdateAiOutboundTaskShrinkRequest
	GetForecastCallRate() *float32
	SetHandlerId(v int64) *UpdateAiOutboundTaskShrinkRequest
	GetHandlerId() *int64
	SetInstanceId(v string) *UpdateAiOutboundTaskShrinkRequest
	GetInstanceId() *string
	SetName(v string) *UpdateAiOutboundTaskShrinkRequest
	GetName() *string
	SetNumRepeated(v int32) *UpdateAiOutboundTaskShrinkRequest
	GetNumRepeated() *int32
	SetOutboundNumsShrink(v string) *UpdateAiOutboundTaskShrinkRequest
	GetOutboundNumsShrink() *string
	SetRecallRuleShrink(v string) *UpdateAiOutboundTaskShrinkRequest
	GetRecallRuleShrink() *string
	SetTaskId(v int64) *UpdateAiOutboundTaskShrinkRequest
	GetTaskId() *int64
}

type UpdateAiOutboundTaskShrinkRequest struct {
	// The concurrent rate for automated outbound calls.
	//
	// example:
	//
	// 10
	ConcurrentRate *int32 `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	// The job description. It can contain 0 to 100 characters.
	//
	// example:
	//
	// 房产销售
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job execution time.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"TUESDAY":[{"start":"06:00","end":"06:05"}],"MONDAY":[{"start":"09:00","end":"18:00"},{"start":"20:30","end":"21:45"},{"start":"22:30","end":"22:50"}],"WEDNESDAY":[{"start":"09:00","end":"18:00"}],"THURSDAY":[{"start":"09:00","end":"18:00"}],"FRIDAY":[{"start":"09:00","end":"18:00"}],"SATURDAY":[{"start":"09:00","end":"18:00"}],"SUNDAY":[{"start":"17:00","end":"23:45"}]}
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// Fixed dialing ratio for predictive outbound calls. Valid values: **≥1**.
	//
	// example:
	//
	// 1.2
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// The skill group ID (for predictive outbound calls) or IVR ID (for automated outbound calls).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job name. Length: 1 to 15 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx外呼
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy for handling duplicate callee numbers.
	//
	// - **0**: Remove duplicates within the job.
	//
	// - **1**: Do not remove duplicates within the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NumRepeated *int32 `json:"NumRepeated,omitempty" xml:"NumRepeated,omitempty"`
	// Outbound caller numbers.
	//
	// This parameter is required.
	OutboundNumsShrink *string `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty"`
	// Failed-call retry policy.
	//
	// > If empty, no retry is performed when an outbound call fails.
	RecallRuleShrink *string `json:"RecallRule,omitempty" xml:"RecallRule,omitempty"`
	// The job ID.
	//
	// You can invoke the [CreateAiOutboundTask](https://help.aliyun.com/document_detail/312260.html) API and check the **Data*	- field in the response, or invoke the [GetAiOutboundTaskList](https://help.aliyun.com/document_detail/2718026.html) API and check the **TaskId*	- field in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateAiOutboundTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiOutboundTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetForecastCallRate() *float32 {
	return s.ForecastCallRate
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetNumRepeated() *int32 {
	return s.NumRepeated
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetOutboundNumsShrink() *string {
	return s.OutboundNumsShrink
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetRecallRuleShrink() *string {
	return s.RecallRuleShrink
}

func (s *UpdateAiOutboundTaskShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetConcurrentRate(v int32) *UpdateAiOutboundTaskShrinkRequest {
	s.ConcurrentRate = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetDescription(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetExecutionTime(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.ExecutionTime = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetForecastCallRate(v float32) *UpdateAiOutboundTaskShrinkRequest {
	s.ForecastCallRate = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetHandlerId(v int64) *UpdateAiOutboundTaskShrinkRequest {
	s.HandlerId = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetInstanceId(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetName(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetNumRepeated(v int32) *UpdateAiOutboundTaskShrinkRequest {
	s.NumRepeated = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetOutboundNumsShrink(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.OutboundNumsShrink = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetRecallRuleShrink(v string) *UpdateAiOutboundTaskShrinkRequest {
	s.RecallRuleShrink = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) SetTaskId(v int64) *UpdateAiOutboundTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAiOutboundTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
