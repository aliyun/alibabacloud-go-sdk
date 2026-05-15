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
	// 1.5
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc_xp_pre-cn-***
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
