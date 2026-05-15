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
	// example:
	//
	// 10
	ConcurrentRate *int32  `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"TUESDAY":[{"start":"06:00","end":"06:05"}],"MONDAY":[{"start":"09:00","end":"18:00"},{"start":"20:30","end":"21:45"},{"start":"22:30","end":"22:50"}],"WEDNESDAY":[{"start":"09:00","end":"18:00"}],"THURSDAY":[{"start":"09:00","end":"18:00"}],"FRIDAY":[{"start":"09:00","end":"18:00"}],"SATURDAY":[{"start":"09:00","end":"18:00"}],"SUNDAY":[{"start":"17:00","end":"23:45"}]}
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// 1.2
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx外呼任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	NumRepeated *int32 `json:"NumRepeated,omitempty" xml:"NumRepeated,omitempty"`
	// This parameter is required.
	OutboundNumsShrink *string `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty"`
	RecallRuleShrink   *string `json:"RecallRule,omitempty" xml:"RecallRule,omitempty"`
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
