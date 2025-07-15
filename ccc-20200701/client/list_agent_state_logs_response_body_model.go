// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentStateLogsResponseBody
	GetCode() *string
	SetData(v []*ListAgentStateLogsResponseBodyData) *ListAgentStateLogsResponseBody
	GetData() []*ListAgentStateLogsResponseBodyData
	SetHttpStatusCode(v int32) *ListAgentStateLogsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAgentStateLogsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentStateLogsResponseBody
	GetRequestId() *string
}

type ListAgentStateLogsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAgentStateLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentStateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentStateLogsResponseBody) GetData() []*ListAgentStateLogsResponseBodyData {
	return s.Data
}

func (s *ListAgentStateLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentStateLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentStateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentStateLogsResponseBody) SetCode(v string) *ListAgentStateLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetData(v []*ListAgentStateLogsResponseBodyData) *ListAgentStateLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetHttpStatusCode(v int32) *ListAgentStateLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetMessage(v string) *ListAgentStateLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetRequestId(v string) *ListAgentStateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAgentStateLogsResponseBodyData struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// example:
	//
	// 1620259200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Break
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// CHECK_IN_BREAK
	StateCode *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	WorkMode  *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ListAgentStateLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStateLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponseBodyData) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListAgentStateLogsResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *ListAgentStateLogsResponseBodyData) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListAgentStateLogsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAgentStateLogsResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListAgentStateLogsResponseBodyData) GetStateCode() *string {
	return s.StateCode
}

func (s *ListAgentStateLogsResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ListAgentStateLogsResponseBodyData) SetBreakCode(v string) *ListAgentStateLogsResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetDuration(v int64) *ListAgentStateLogsResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetOutboundScenario(v bool) *ListAgentStateLogsResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetStartTime(v int64) *ListAgentStateLogsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetState(v string) *ListAgentStateLogsResponseBodyData {
	s.State = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetStateCode(v string) *ListAgentStateLogsResponseBodyData {
	s.StateCode = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetWorkMode(v string) *ListAgentStateLogsResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
