// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrentRate(v int32) *CreateAiOutboundTaskShrinkRequest
	GetConcurrentRate() *int32
	SetDescription(v string) *CreateAiOutboundTaskShrinkRequest
	GetDescription() *string
	SetExecutionTime(v string) *CreateAiOutboundTaskShrinkRequest
	GetExecutionTime() *string
	SetForecastCallRate(v float32) *CreateAiOutboundTaskShrinkRequest
	GetForecastCallRate() *float32
	SetHandlerId(v int64) *CreateAiOutboundTaskShrinkRequest
	GetHandlerId() *int64
	SetInstanceId(v string) *CreateAiOutboundTaskShrinkRequest
	GetInstanceId() *string
	SetName(v string) *CreateAiOutboundTaskShrinkRequest
	GetName() *string
	SetNumRepeated(v int32) *CreateAiOutboundTaskShrinkRequest
	GetNumRepeated() *int32
	SetOutboundNumsShrink(v string) *CreateAiOutboundTaskShrinkRequest
	GetOutboundNumsShrink() *string
	SetRecallRuleShrink(v string) *CreateAiOutboundTaskShrinkRequest
	GetRecallRuleShrink() *string
	SetType(v int32) *CreateAiOutboundTaskShrinkRequest
	GetType() *int32
}

type CreateAiOutboundTaskShrinkRequest struct {
	// Concurrent call rate for automated outbound calls.
	//
	// example:
	//
	// 10
	ConcurrentRate *int32 `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	// Job description. Length: 0 to 100 characters.
	//
	// example:
	//
	// 房产销售
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Job execution time, in JSON format.
	//
	// > The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"TUESDAY":[{"start":"06:00","end":"06:05"}],"MONDAY":[{"start":"09:00","end":"18:00"},{"start":"20:30","end":"21:45"},{"start":"22:30","end":"22:50"}],"WEDNESDAY":[{"start":"09:00","end":"18:00"}],"THURSDAY":[{"start":"09:00","end":"18:00"}],"FRIDAY":[{"start":"09:00","end":"18:00"}],"SATURDAY":[{"start":"09:00","end":"18:00"}],"SUNDAY":[{"start":"17:00","end":"23:45"}]}
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// Fixed outbound ratio for predictive dialing. Valid values: **≥1**.
	//
	// example:
	//
	// 1.5
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// The skill group ID (for predictive outbound calls) or IVR ID (for automated outbound calls). You can obtain this information in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// AICCS instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// cc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job name. Length: 1 to 15 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Called number deduplication policy. Valid values:
	//
	// - **0**: Remove duplicates within the job.
	//
	// - **1**: Do not remove duplicates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NumRepeated *int32 `json:"NumRepeated,omitempty" xml:"NumRepeated,omitempty"`
	// Outbound caller numbers.
	//
	// > Must be purchased numbers. Separate multiple numbers with commas (,).
	//
	// This parameter is required.
	OutboundNumsShrink *string `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty"`
	// Failed call retry policy.
	//
	// > If empty, no retries are performed.
	RecallRuleShrink *string `json:"RecallRule,omitempty" xml:"RecallRule,omitempty"`
	// Task Type. Valid values:
	//
	// - **2**: Predictive outbound call.
	//
	// - **3**: Automated outbound call.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAiOutboundTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskShrinkRequest) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *CreateAiOutboundTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAiOutboundTaskShrinkRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *CreateAiOutboundTaskShrinkRequest) GetForecastCallRate() *float32 {
	return s.ForecastCallRate
}

func (s *CreateAiOutboundTaskShrinkRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *CreateAiOutboundTaskShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAiOutboundTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAiOutboundTaskShrinkRequest) GetNumRepeated() *int32 {
	return s.NumRepeated
}

func (s *CreateAiOutboundTaskShrinkRequest) GetOutboundNumsShrink() *string {
	return s.OutboundNumsShrink
}

func (s *CreateAiOutboundTaskShrinkRequest) GetRecallRuleShrink() *string {
	return s.RecallRuleShrink
}

func (s *CreateAiOutboundTaskShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAiOutboundTaskShrinkRequest) SetConcurrentRate(v int32) *CreateAiOutboundTaskShrinkRequest {
	s.ConcurrentRate = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetDescription(v string) *CreateAiOutboundTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetExecutionTime(v string) *CreateAiOutboundTaskShrinkRequest {
	s.ExecutionTime = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetForecastCallRate(v float32) *CreateAiOutboundTaskShrinkRequest {
	s.ForecastCallRate = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetHandlerId(v int64) *CreateAiOutboundTaskShrinkRequest {
	s.HandlerId = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetInstanceId(v string) *CreateAiOutboundTaskShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetName(v string) *CreateAiOutboundTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetNumRepeated(v int32) *CreateAiOutboundTaskShrinkRequest {
	s.NumRepeated = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetOutboundNumsShrink(v string) *CreateAiOutboundTaskShrinkRequest {
	s.OutboundNumsShrink = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetRecallRuleShrink(v string) *CreateAiOutboundTaskShrinkRequest {
	s.RecallRuleShrink = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) SetType(v int32) *CreateAiOutboundTaskShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateAiOutboundTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
