// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRealtimeInstanceStatesResponseBody
	GetCode() *string
	SetData(v *GetRealtimeInstanceStatesResponseBodyData) *GetRealtimeInstanceStatesResponseBody
	GetData() *GetRealtimeInstanceStatesResponseBodyData
	SetHttpStatusCode(v int32) *GetRealtimeInstanceStatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRealtimeInstanceStatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRealtimeInstanceStatesResponseBody
	GetRequestId() *string
}

type GetRealtimeInstanceStatesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *GetRealtimeInstanceStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeInstanceStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRealtimeInstanceStatesResponseBody) GetData() *GetRealtimeInstanceStatesResponseBodyData {
	return s.Data
}

func (s *GetRealtimeInstanceStatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRealtimeInstanceStatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRealtimeInstanceStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRealtimeInstanceStatesResponseBody) SetCode(v string) *GetRealtimeInstanceStatesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetData(v *GetRealtimeInstanceStatesResponseBodyData) *GetRealtimeInstanceStatesResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetHttpStatusCode(v int32) *GetRealtimeInstanceStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetMessage(v string) *GetRealtimeInstanceStatesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetRequestId(v string) *GetRealtimeInstanceStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRealtimeInstanceStatesResponseBodyData struct {
	// List of distributions of break code counts.
	BreakCodeDetailList []*GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// Number of agents on break.
	//
	// example:
	//
	// 0
	BreakingAgents *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Number of calls interacting within IVR.
	//
	// example:
	//
	// 0
	InteractiveCalls *int64 `json:"InteractiveCalls,omitempty" xml:"InteractiveCalls,omitempty"`
	// Number of logged-in agents (including agents in ready, on break, on call, or other non-unpublished statuses).
	//
	// example:
	//
	// 0
	LoggedInAgents *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	// Maximum queue waiting time among currently queued incoming calls.
	//
	// example:
	//
	// 0
	LongestWaitingTime *int64 `json:"LongestWaitingTime,omitempty" xml:"LongestWaitingTime,omitempty"`
	// Number of agents in ready status.
	//
	// example:
	//
	// 0
	ReadyAgents *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	// Number of agents on calls.
	//
	// example:
	//
	// 0
	TalkingAgents *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	// Total number of agents.
	//
	// example:
	//
	// 0
	TotalAgents *int64 `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	// Number of incoming calls currently queued.
	//
	// example:
	//
	// 0
	WaitingCalls *int64 `json:"WaitingCalls,omitempty" xml:"WaitingCalls,omitempty"`
	// Number of agents in post-processing status.
	//
	// example:
	//
	// 0
	WorkingAgents *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s GetRealtimeInstanceStatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetBreakCodeDetailList() []*GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetBreakingAgents() *int64 {
	return s.BreakingAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetInteractiveCalls() *int64 {
	return s.InteractiveCalls
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetLongestWaitingTime() *int64 {
	return s.LongestWaitingTime
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetTalkingAgents() *int64 {
	return s.TalkingAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetTotalAgents() *int64 {
	return s.TotalAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetWaitingCalls() *int64 {
	return s.WaitingCalls
}

func (s *GetRealtimeInstanceStatesResponseBodyData) GetWorkingAgents() *int64 {
	return s.WorkingAgents
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetBreakCodeDetailList(v []*GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) *GetRealtimeInstanceStatesResponseBodyData {
	s.BreakCodeDetailList = v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetBreakingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.BreakingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetInstanceId(v string) *GetRealtimeInstanceStatesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetInteractiveCalls(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.InteractiveCalls = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetLoggedInAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetLongestWaitingTime(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.LongestWaitingTime = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetReadyAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetTalkingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.TalkingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetTotalAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.TotalAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetWaitingCalls(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.WaitingCalls = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetWorkingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.WorkingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) Validate() error {
	if s.BreakCodeDetailList != nil {
		for _, item := range s.BreakCodeDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList struct {
	// Break code.
	//
	// example:
	//
	// 默认
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Break count.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) SetBreakCode(v string) *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) SetCount(v int64) *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyDataBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
