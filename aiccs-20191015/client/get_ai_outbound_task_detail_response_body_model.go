// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiOutboundTaskDetailResponseBody
	GetCode() *string
	SetData(v *GetAiOutboundTaskDetailResponseBodyData) *GetAiOutboundTaskDetailResponseBody
	GetData() *GetAiOutboundTaskDetailResponseBodyData
	SetMessage(v string) *GetAiOutboundTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiOutboundTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiOutboundTaskDetailResponseBody
	GetSuccess() *bool
}

type GetAiOutboundTaskDetailResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAiOutboundTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAiOutboundTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiOutboundTaskDetailResponseBody) GetData() *GetAiOutboundTaskDetailResponseBodyData {
	return s.Data
}

func (s *GetAiOutboundTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiOutboundTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiOutboundTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiOutboundTaskDetailResponseBody) SetCode(v string) *GetAiOutboundTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBody) SetData(v *GetAiOutboundTaskDetailResponseBodyData) *GetAiOutboundTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBody) SetMessage(v string) *GetAiOutboundTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBody) SetRequestId(v string) *GetAiOutboundTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBody) SetSuccess(v bool) *GetAiOutboundTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskDetailResponseBodyData struct {
	// example:
	//
	// 10
	ConcurrentRate *int32  `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {"TUESDAY":[{"start":"06:00","end":"06:05"}],"MONDAY":[{"start":"09:00","end":"18:00"},{"start":"20:30","end":"21:45"},{"start":"22:30","end":"22:50"}],"WEDNESDAY":[{"start":"09:00","end":"18:00"}],"THURSDAY":[{"start":"09:00","end":"18:00"}],"FRIDAY":[{"start":"09:00","end":"18:00"}],"SATURDAY":[{"start":"09:00","end":"18:00"}],"SUNDAY":[{"start":"17:00","end":"23:45"}]}
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// example:
	//
	// 1.2
	ForecastCallRate *float32 `json:"ForecastCallRate,omitempty" xml:"ForecastCallRate,omitempty"`
	// example:
	//
	// 123456
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// example:
	//
	// 热线技能组
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// example:
	//
	// xx外呼任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	NumRepeated  *int32                                             `json:"NumRepeated,omitempty" xml:"NumRepeated,omitempty"`
	OutboundNums []*string                                          `json:"OutboundNums,omitempty" xml:"OutboundNums,omitempty" type:"Repeated"`
	RecallRule   *GetAiOutboundTaskDetailResponseBodyDataRecallRule `json:"RecallRule,omitempty" xml:"RecallRule,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 未开始
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAiOutboundTaskDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetForecastCallRate() *float32 {
	return s.ForecastCallRate
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetHandlerName() *string {
	return s.HandlerName
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetNumRepeated() *int32 {
	return s.NumRepeated
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetOutboundNums() []*string {
	return s.OutboundNums
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetRecallRule() *GetAiOutboundTaskDetailResponseBodyDataRecallRule {
	return s.RecallRule
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskDetailResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetConcurrentRate(v int32) *GetAiOutboundTaskDetailResponseBodyData {
	s.ConcurrentRate = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetDescription(v string) *GetAiOutboundTaskDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetExecutionTime(v string) *GetAiOutboundTaskDetailResponseBodyData {
	s.ExecutionTime = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetForecastCallRate(v float32) *GetAiOutboundTaskDetailResponseBodyData {
	s.ForecastCallRate = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetHandlerId(v int64) *GetAiOutboundTaskDetailResponseBodyData {
	s.HandlerId = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetHandlerName(v string) *GetAiOutboundTaskDetailResponseBodyData {
	s.HandlerName = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetName(v string) *GetAiOutboundTaskDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetNumRepeated(v int32) *GetAiOutboundTaskDetailResponseBodyData {
	s.NumRepeated = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetOutboundNums(v []*string) *GetAiOutboundTaskDetailResponseBodyData {
	s.OutboundNums = v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetRecallRule(v *GetAiOutboundTaskDetailResponseBodyDataRecallRule) *GetAiOutboundTaskDetailResponseBodyData {
	s.RecallRule = v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetStatus(v int32) *GetAiOutboundTaskDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetStatusDesc(v string) *GetAiOutboundTaskDetailResponseBodyData {
	s.StatusDesc = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetTaskId(v int64) *GetAiOutboundTaskDetailResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) SetType(v int32) *GetAiOutboundTaskDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyData) Validate() error {
	if s.RecallRule != nil {
		if err := s.RecallRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskDetailResponseBodyDataRecallRule struct {
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s GetAiOutboundTaskDetailResponseBodyDataRecallRule) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskDetailResponseBodyDataRecallRule) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskDetailResponseBodyDataRecallRule) GetCount() *int32 {
	return s.Count
}

func (s *GetAiOutboundTaskDetailResponseBodyDataRecallRule) GetInterval() *int32 {
	return s.Interval
}

func (s *GetAiOutboundTaskDetailResponseBodyDataRecallRule) SetCount(v int32) *GetAiOutboundTaskDetailResponseBodyDataRecallRule {
	s.Count = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyDataRecallRule) SetInterval(v int32) *GetAiOutboundTaskDetailResponseBodyDataRecallRule {
	s.Interval = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponseBodyDataRecallRule) Validate() error {
	return dara.Validate(s)
}
