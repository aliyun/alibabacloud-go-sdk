// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrentRate(v int32) *CreateAiOutboundTaskRequest
	GetConcurrentRate() *int32
	SetDescription(v string) *CreateAiOutboundTaskRequest
	GetDescription() *string
	SetExecutionTime(v string) *CreateAiOutboundTaskRequest
	GetExecutionTime() *string
	SetForecastCallRate(v float32) *CreateAiOutboundTaskRequest
	GetForecastCallRate() *float32
	SetHandlerId(v int64) *CreateAiOutboundTaskRequest
	GetHandlerId() *int64
	SetInstanceId(v string) *CreateAiOutboundTaskRequest
	GetInstanceId() *string
	SetName(v string) *CreateAiOutboundTaskRequest
	GetName() *string
	SetNumRepeated(v int32) *CreateAiOutboundTaskRequest
	GetNumRepeated() *int32
	SetOutboundNums(v []*string) *CreateAiOutboundTaskRequest
	GetOutboundNums() []*string
	SetRecallRule(v *CreateAiOutboundTaskRequestRecallRule) *CreateAiOutboundTaskRequest
	GetRecallRule() *CreateAiOutboundTaskRequestRecallRule
	SetType(v int32) *CreateAiOutboundTaskRequest
	GetType() *int32
}

type CreateAiOutboundTaskRequest struct {
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
	OutboundNums []*string `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty" type:"Repeated"`
	// Failed call retry policy.
	//
	// > If empty, no retries are performed.
	RecallRule *CreateAiOutboundTaskRequestRecallRule `json:"RecallRule,omitempty" xml:"RecallRule,omitempty" type:"Struct"`
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

func (s CreateAiOutboundTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskRequest) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *CreateAiOutboundTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAiOutboundTaskRequest) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *CreateAiOutboundTaskRequest) GetForecastCallRate() *float32 {
	return s.ForecastCallRate
}

func (s *CreateAiOutboundTaskRequest) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *CreateAiOutboundTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAiOutboundTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateAiOutboundTaskRequest) GetNumRepeated() *int32 {
	return s.NumRepeated
}

func (s *CreateAiOutboundTaskRequest) GetOutboundNums() []*string {
	return s.OutboundNums
}

func (s *CreateAiOutboundTaskRequest) GetRecallRule() *CreateAiOutboundTaskRequestRecallRule {
	return s.RecallRule
}

func (s *CreateAiOutboundTaskRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAiOutboundTaskRequest) SetConcurrentRate(v int32) *CreateAiOutboundTaskRequest {
	s.ConcurrentRate = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetDescription(v string) *CreateAiOutboundTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetExecutionTime(v string) *CreateAiOutboundTaskRequest {
	s.ExecutionTime = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetForecastCallRate(v float32) *CreateAiOutboundTaskRequest {
	s.ForecastCallRate = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetHandlerId(v int64) *CreateAiOutboundTaskRequest {
	s.HandlerId = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetInstanceId(v string) *CreateAiOutboundTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetName(v string) *CreateAiOutboundTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetNumRepeated(v int32) *CreateAiOutboundTaskRequest {
	s.NumRepeated = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetOutboundNums(v []*string) *CreateAiOutboundTaskRequest {
	s.OutboundNums = v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetRecallRule(v *CreateAiOutboundTaskRequestRecallRule) *CreateAiOutboundTaskRequest {
	s.RecallRule = v
	return s
}

func (s *CreateAiOutboundTaskRequest) SetType(v int32) *CreateAiOutboundTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateAiOutboundTaskRequest) Validate() error {
	if s.RecallRule != nil {
		if err := s.RecallRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAiOutboundTaskRequestRecallRule struct {
	// Number of retries. Valid values: **1 to 3**.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Retry interval. Valid values: **1 to 60**, unit: minutes.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s CreateAiOutboundTaskRequestRecallRule) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskRequestRecallRule) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskRequestRecallRule) GetCount() *int32 {
	return s.Count
}

func (s *CreateAiOutboundTaskRequestRecallRule) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateAiOutboundTaskRequestRecallRule) SetCount(v int32) *CreateAiOutboundTaskRequestRecallRule {
	s.Count = &v
	return s
}

func (s *CreateAiOutboundTaskRequestRecallRule) SetInterval(v int32) *CreateAiOutboundTaskRequestRecallRule {
	s.Interval = &v
	return s
}

func (s *CreateAiOutboundTaskRequestRecallRule) Validate() error {
	return dara.Validate(s)
}
