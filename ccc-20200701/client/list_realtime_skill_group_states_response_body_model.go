// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeSkillGroupStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRealtimeSkillGroupStatesResponseBody
	GetCode() *string
	SetData(v *ListRealtimeSkillGroupStatesResponseBodyData) *ListRealtimeSkillGroupStatesResponseBody
	GetData() *ListRealtimeSkillGroupStatesResponseBodyData
	SetHttpStatusCode(v int32) *ListRealtimeSkillGroupStatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRealtimeSkillGroupStatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRealtimeSkillGroupStatesResponseBody
	GetRequestId() *string
}

type ListRealtimeSkillGroupStatesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListRealtimeSkillGroupStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 26A34338-5CD9-4C95-A7A6-5BDCE76C6B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRealtimeSkillGroupStatesResponseBody) GetData() *ListRealtimeSkillGroupStatesResponseBodyData {
	return s.Data
}

func (s *ListRealtimeSkillGroupStatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRealtimeSkillGroupStatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRealtimeSkillGroupStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetCode(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetData(v *ListRealtimeSkillGroupStatesResponseBodyData) *ListRealtimeSkillGroupStatesResponseBody {
	s.Data = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetHttpStatusCode(v int32) *ListRealtimeSkillGroupStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetMessage(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetRequestId(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeSkillGroupStatesResponseBodyData struct {
	List []*ListRealtimeSkillGroupStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) GetList() []*ListRealtimeSkillGroupStatesResponseBodyDataList {
	return s.List
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetList(v []*ListRealtimeSkillGroupStatesResponseBodyDataList) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.List = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetPageNumber(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetPageSize(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetTotalCount(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeSkillGroupStatesResponseBodyDataList struct {
	BreakCodeDetailList []*ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	BreakingAgents *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	LoggedInAgents *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	// example:
	//
	// 0
	LongestWaitingTime *int64 `json:"LongestWaitingTime,omitempty" xml:"LongestWaitingTime,omitempty"`
	// example:
	//
	// 0
	OutboundScenarioReadyAgents *int64 `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	// example:
	//
	// 2
	ReadyAgents *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// skillgroup
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// example:
	//
	// 0
	TalkingAgents *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	TotalAgents   *int64 `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	// example:
	//
	// 0
	WaitingCalls *int64 `json:"WaitingCalls,omitempty" xml:"WaitingCalls,omitempty"`
	// example:
	//
	// 0
	WorkingAgents *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetBreakCodeDetailList() []*ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetBreakingAgents() *int64 {
	return s.BreakingAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetLoggedInAgents() *int64 {
	return s.LoggedInAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetLongestWaitingTime() *int64 {
	return s.LongestWaitingTime
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetOutboundScenarioReadyAgents() *int64 {
	return s.OutboundScenarioReadyAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetReadyAgents() *int64 {
	return s.ReadyAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetTalkingAgents() *int64 {
	return s.TalkingAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetTotalAgents() *int64 {
	return s.TotalAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetWaitingCalls() *int64 {
	return s.WaitingCalls
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) GetWorkingAgents() *int64 {
	return s.WorkingAgents
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetBreakCodeDetailList(v []*ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetBreakingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.BreakingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetInstanceId(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetLoggedInAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.LoggedInAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetLongestWaitingTime(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.LongestWaitingTime = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetOutboundScenarioReadyAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetReadyAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.ReadyAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetSkillGroupId(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetSkillGroupName(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetTalkingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.TalkingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetTotalAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.TotalAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetWaitingCalls(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.WaitingCalls = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetWorkingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.WorkingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) SetBreakCode(v string) *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) SetCount(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataListBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
