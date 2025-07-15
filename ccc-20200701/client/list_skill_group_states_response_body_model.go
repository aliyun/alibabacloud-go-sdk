// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillGroupStatesResponseBody
	GetCode() *string
	SetData(v *ListSkillGroupStatesResponseBodyData) *ListSkillGroupStatesResponseBody
	GetData() *ListSkillGroupStatesResponseBodyData
	SetHttpStatusCode(v int32) *ListSkillGroupStatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSkillGroupStatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSkillGroupStatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSkillGroupStatesResponseBody
	GetSuccess() *bool
}

type ListSkillGroupStatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSkillGroupStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1F69EBB0-63E9-5DDE-887F-9FC040ADF309
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupStatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillGroupStatesResponseBody) GetData() *ListSkillGroupStatesResponseBodyData {
	return s.Data
}

func (s *ListSkillGroupStatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSkillGroupStatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillGroupStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillGroupStatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSkillGroupStatesResponseBody) SetCode(v string) *ListSkillGroupStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupStatesResponseBody) SetData(v *ListSkillGroupStatesResponseBodyData) *ListSkillGroupStatesResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupStatesResponseBody) SetHttpStatusCode(v int32) *ListSkillGroupStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillGroupStatesResponseBody) SetMessage(v string) *ListSkillGroupStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupStatesResponseBody) SetRequestId(v string) *ListSkillGroupStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupStatesResponseBody) SetSuccess(v bool) *ListSkillGroupStatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSkillGroupStatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupStatesResponseBodyData struct {
	List []*ListSkillGroupStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillGroupStatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupStatesResponseBodyData) GetList() []*ListSkillGroupStatesResponseBodyDataList {
	return s.List
}

func (s *ListSkillGroupStatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupStatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupStatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillGroupStatesResponseBodyData) SetList(v []*ListSkillGroupStatesResponseBodyDataList) *ListSkillGroupStatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListSkillGroupStatesResponseBodyData) SetPageNumber(v int32) *ListSkillGroupStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyData) SetPageSize(v int32) *ListSkillGroupStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyData) SetTotalCount(v int32) *ListSkillGroupStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupStatesResponseBodyDataList struct {
	// example:
	//
	// 3
	AverageWaitingTime *int64 `json:"AverageWaitingTime,omitempty" xml:"AverageWaitingTime,omitempty"`
	// example:
	//
	// 0
	BreakingAgents       *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	InboundTalkingAgents *int64 `json:"InboundTalkingAgents,omitempty" xml:"InboundTalkingAgents,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0
	LoggedInAgents *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	// example:
	//
	// 0
	LongestCall *int64 `json:"LongestCall,omitempty" xml:"LongestCall,omitempty"`
	// example:
	//
	// 0
	OutboundScenarioReadyAgents *int64 `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundTalkingAgents       *int64 `json:"OutboundTalkingAgents,omitempty" xml:"OutboundTalkingAgents,omitempty"`
	// example:
	//
	// 1
	ReadyAgents *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	// example:
	//
	// skillgroup1@ccc-test
	SkillGroupId   *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 0
	TalkingAgents *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	// example:
	//
	// 0
	WaitingCalls *int64 `json:"WaitingCalls,omitempty" xml:"WaitingCalls,omitempty"`
	// example:
	//
	// 0
	WaitingCallsLevel10 *int64 `json:"WaitingCallsLevel10,omitempty" xml:"WaitingCallsLevel10,omitempty"`
	// example:
	//
	// 0
	WaitingCallsLevel20 *int64 `json:"WaitingCallsLevel20,omitempty" xml:"WaitingCallsLevel20,omitempty"`
	// example:
	//
	// 0
	WaitingCallsLevel30 *int64 `json:"WaitingCallsLevel30,omitempty" xml:"WaitingCallsLevel30,omitempty"`
	// example:
	//
	// 0
	WorkingAgents *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s ListSkillGroupStatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetAverageWaitingTime() *int64 {
	return s.AverageWaitingTime
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetBreakingAgents() *int64 {
	return s.BreakingAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetInboundTalkingAgents() *int64 {
	return s.InboundTalkingAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetLongestCall() *int64 {
	return s.LongestCall
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetOutboundScenarioReadyAgents() *int64 {
	return s.OutboundScenarioReadyAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetOutboundTalkingAgents() *int64 {
	return s.OutboundTalkingAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetTalkingAgents() *int64 {
	return s.TalkingAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetWaitingCalls() *int64 {
	return s.WaitingCalls
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetWaitingCallsLevel10() *int64 {
	return s.WaitingCallsLevel10
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetWaitingCallsLevel20() *int64 {
	return s.WaitingCallsLevel20
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetWaitingCallsLevel30() *int64 {
	return s.WaitingCallsLevel30
}

func (s *ListSkillGroupStatesResponseBodyDataList) GetWorkingAgents() *int64 {
	return s.WorkingAgents
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetAverageWaitingTime(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.AverageWaitingTime = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetBreakingAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.BreakingAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetInboundTalkingAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.InboundTalkingAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetInstanceId(v string) *ListSkillGroupStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetLoggedInAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.LoggedInAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetLongestCall(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.LongestCall = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetOutboundScenarioReadyAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetOutboundTalkingAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.OutboundTalkingAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetReadyAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.ReadyAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetSkillGroupId(v string) *ListSkillGroupStatesResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetSkillGroupName(v string) *ListSkillGroupStatesResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetTalkingAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.TalkingAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetWaitingCalls(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.WaitingCalls = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetWaitingCallsLevel10(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.WaitingCallsLevel10 = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetWaitingCallsLevel20(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.WaitingCallsLevel20 = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetWaitingCallsLevel30(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.WaitingCallsLevel30 = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) SetWorkingAgents(v int64) *ListSkillGroupStatesResponseBodyDataList {
	s.WorkingAgents = &v
	return s
}

func (s *ListSkillGroupStatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
