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
	OutboundNums []*string                              `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty" type:"Repeated"`
	RecallRule   *UpdateAiOutboundTaskRequestRecallRule `json:"RecallRule,omitempty" xml:"RecallRule,omitempty" type:"Struct"`
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
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
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
