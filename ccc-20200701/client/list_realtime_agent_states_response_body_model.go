// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeAgentStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRealtimeAgentStatesResponseBody
	GetCode() *string
	SetData(v *ListRealtimeAgentStatesResponseBodyData) *ListRealtimeAgentStatesResponseBody
	GetData() *ListRealtimeAgentStatesResponseBodyData
	SetHttpStatusCode(v int32) *ListRealtimeAgentStatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRealtimeAgentStatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRealtimeAgentStatesResponseBody
	GetRequestId() *string
}

type ListRealtimeAgentStatesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ListRealtimeAgentStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeAgentStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRealtimeAgentStatesResponseBody) GetData() *ListRealtimeAgentStatesResponseBodyData {
	return s.Data
}

func (s *ListRealtimeAgentStatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRealtimeAgentStatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRealtimeAgentStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeAgentStatesResponseBody) SetCode(v string) *ListRealtimeAgentStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetData(v *ListRealtimeAgentStatesResponseBodyData) *ListRealtimeAgentStatesResponseBody {
	s.Data = v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetHttpStatusCode(v int32) *ListRealtimeAgentStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetMessage(v string) *ListRealtimeAgentStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetRequestId(v string) *ListRealtimeAgentStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRealtimeAgentStatesResponseBodyData struct {
	// List of real-time agent status data.
	List []*ListRealtimeAgentStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number, ranging from 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRealtimeAgentStatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBodyData) GetList() []*ListRealtimeAgentStatesResponseBodyDataList {
	return s.List
}

func (s *ListRealtimeAgentStatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeAgentStatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeAgentStatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetList(v []*ListRealtimeAgentStatesResponseBodyDataList) *ListRealtimeAgentStatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetPageNumber(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetPageSize(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetTotalCount(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRealtimeAgentStatesResponseBodyDataList struct {
	// Agent ID.
	//
	// example:
	//
	// agent1@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Agent name.
	//
	// example:
	//
	// 坐席小王
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Break code.
	//
	// **Enumeration values:**
	//
	// - RingingTimeout: Break caused by agent ringing timeout.
	//
	// - RejectCall: Break caused by agent call rejection.
	//
	// - Warm-up: Temporary break state after the agent is published and before becoming idle.
	//
	// example:
	//
	// Warm-up
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Call type.
	//
	// example:
	//
	// Outbound
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// Used in specific three-party scenarios, primarily for listener, coaching, and consultation. In three-party scenarios, it represents the third party—for example, the agent being monitored or coached in a listener or coaching scenario, or the agent or external number to which a call is transferred in a consultation scenario.
	//
	// example:
	//
	// agent@ccc-test
	CounterParty *string `json:"CounterParty,omitempty" xml:"CounterParty,omitempty"`
	// Duration of the current status, in seconds.
	//
	// example:
	//
	// 16
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The agent\\"s extension number.
	//
	// example:
	//
	// 80317391
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The agent\\"s personal phone number.
	//
	// example:
	//
	// 1382114****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Whether the agent is in outbound-only mode.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// List of skill group IDs that the agent has signed into.
	SkillGroupIdList []*string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
	// List of skill group names that the agent has signed into.
	SkillGroupNameList []*string `json:"SkillGroupNameList,omitempty" xml:"SkillGroupNameList,omitempty" type:"Repeated"`
	// Agent status.
	//
	// example:
	//
	// ACW
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Sub-status. In some scenarios, the agent\\"s status cannot be fully represented by the State field alone, so a sub-status is required for clarification. For example, when an agent is being monitored, State=Talking and StateCode=Monitoring.
	//
	// example:
	//
	// Monitored
	StateCode *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	// Time when the status started.
	//
	// example:
	//
	// 1696670640774
	StateTime *int64 `json:"StateTime,omitempty" xml:"StateTime,omitempty"`
	// Work mode.
	//
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ListRealtimeAgentStatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetCallType() *string {
	return s.CallType
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetCounterParty() *string {
	return s.CounterParty
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetExtension() *string {
	return s.Extension
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetMobile() *string {
	return s.Mobile
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetSkillGroupIdList() []*string {
	return s.SkillGroupIdList
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetSkillGroupNameList() []*string {
	return s.SkillGroupNameList
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetStateCode() *string {
	return s.StateCode
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetStateTime() *int64 {
	return s.StateTime
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) GetWorkMode() *string {
	return s.WorkMode
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetAgentId(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetAgentName(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetBreakCode(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.BreakCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetCallType(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.CallType = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetCounterParty(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.CounterParty = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetDuration(v int64) *ListRealtimeAgentStatesResponseBodyDataList {
	s.Duration = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetExtension(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.Extension = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetInstanceId(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetMobile(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetOutboundScenario(v bool) *ListRealtimeAgentStatesResponseBodyDataList {
	s.OutboundScenario = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetSkillGroupIdList(v []*string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.SkillGroupIdList = v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetSkillGroupNameList(v []*string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.SkillGroupNameList = v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetState(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetStateCode(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.StateCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetStateTime(v int64) *ListRealtimeAgentStatesResponseBodyDataList {
	s.StateTime = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetWorkMode(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.WorkMode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
