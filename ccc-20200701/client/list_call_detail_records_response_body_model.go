// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallDetailRecordsResponseBody
	GetCode() *string
	SetData(v *ListCallDetailRecordsResponseBodyData) *ListCallDetailRecordsResponseBody
	GetData() *ListCallDetailRecordsResponseBodyData
	SetHttpStatusCode(v int32) *ListCallDetailRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallDetailRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCallDetailRecordsResponseBody
	GetRequestId() *string
}

type ListCallDetailRecordsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCallDetailRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallDetailRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallDetailRecordsResponseBody) GetData() *ListCallDetailRecordsResponseBodyData {
	return s.Data
}

func (s *ListCallDetailRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallDetailRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallDetailRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallDetailRecordsResponseBody) SetCode(v string) *ListCallDetailRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetData(v *ListCallDetailRecordsResponseBodyData) *ListCallDetailRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetHttpStatusCode(v int32) *ListCallDetailRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetMessage(v string) *ListCallDetailRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetRequestId(v string) *ListCallDetailRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsResponseBodyData struct {
	List []*ListCallDetailRecordsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallDetailRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyData) GetList() []*ListCallDetailRecordsResponseBodyDataList {
	return s.List
}

func (s *ListCallDetailRecordsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCallDetailRecordsResponseBodyData) SetList(v []*ListCallDetailRecordsResponseBodyDataList) *ListCallDetailRecordsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageNumber(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageSize(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetTotalCount(v int32) *ListCallDetailRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsResponseBodyDataList struct {
	// example:
	//
	// 0533128****
	AdditionalBroker *string `json:"AdditionalBroker,omitempty" xml:"AdditionalBroker,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentIds   *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	AgentNames *string `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	// example:
	//
	// 0533127****
	Broker *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	// example:
	//
	// 30
	CallDuration *string `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	CallIds      *string `json:"CallIds,omitempty" xml:"CallIds,omitempty"`
	// example:
	//
	// 1332315****
	CalledNumber   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalleeLocation *string `json:"CalleeLocation,omitempty" xml:"CalleeLocation,omitempty"`
	CallerLocation *string `json:"CallerLocation,omitempty" xml:"CallerLocation,omitempty"`
	// example:
	//
	// 0533128****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// Success
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	// example:
	//
	// job-12515239414412****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// Outbound
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	DialingTime *int64  `json:"DialingTime,omitempty" xml:"DialingTime,omitempty"`
	// example:
	//
	// NotConnected
	EarlyMediaState *string `json:"EarlyMediaState,omitempty" xml:"EarlyMediaState,omitempty"`
	// example:
	//
	// 1532448000000
	EstablishedTime *int64 `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	HeldTime        *int64 `json:"HeldTime,omitempty" xml:"HeldTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 8
	IvrTime *int64 `json:"IvrTime,omitempty" xml:"IvrTime,omitempty"`
	// example:
	//
	// 0
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// example:
	//
	// 10
	RecordingDuration *int64 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// example:
	//
	// true
	RecordingReady *bool `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	// example:
	//
	// customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ReleaseReason    *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// 1532707199000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// 5
	RingTime                *int64  `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	SatisfactionDescription *string `json:"SatisfactionDescription,omitempty" xml:"SatisfactionDescription,omitempty"`
	// example:
	//
	// 1
	SatisfactionIndex *int32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// IVR
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	// example:
	//
	// true
	SatisfactionSurveyOffered *bool `json:"SatisfactionSurveyOffered,omitempty" xml:"SatisfactionSurveyOffered,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupIds   *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	SkillGroupNames *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TalkTime  *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// example:
	//
	// 5
	WaitTime *int64 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s ListCallDetailRecordsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetAdditionalBroker() *string {
	return s.AdditionalBroker
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetAgentIds() *string {
	return s.AgentIds
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetAgentNames() *string {
	return s.AgentNames
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetBroker() *string {
	return s.Broker
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCallDuration() *string {
	return s.CallDuration
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCallIds() *string {
	return s.CallIds
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCalleeLocation() *string {
	return s.CalleeLocation
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCallerLocation() *string {
	return s.CallerLocation
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetContactDisposition() *string {
	return s.ContactDisposition
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetContactType() *string {
	return s.ContactType
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetDialingTime() *int64 {
	return s.DialingTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetEarlyMediaState() *string {
	return s.EarlyMediaState
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetEstablishedTime() *int64 {
	return s.EstablishedTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetHeldTime() *int64 {
	return s.HeldTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetIvrTime() *int64 {
	return s.IvrTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetQueueTime() *int64 {
	return s.QueueTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetRecordingDuration() *int64 {
	return s.RecordingDuration
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetRecordingReady() *bool {
	return s.RecordingReady
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetRingTime() *int64 {
	return s.RingTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSatisfactionDescription() *string {
	return s.SatisfactionDescription
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSatisfactionIndex() *int32 {
	return s.SatisfactionIndex
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSatisfactionSurveyChannel() *string {
	return s.SatisfactionSurveyChannel
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSatisfactionSurveyOffered() *bool {
	return s.SatisfactionSurveyOffered
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetSkillGroupNames() *string {
	return s.SkillGroupNames
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) GetWaitTime() *int64 {
	return s.WaitTime
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAdditionalBroker(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AdditionalBroker = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAgentIds(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AgentIds = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAgentNames(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AgentNames = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetBroker(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.Broker = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallDuration(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallDuration = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallIds(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallIds = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCalledNumber(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCalleeLocation(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CalleeLocation = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallerLocation(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallerLocation = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallingNumber(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactDisposition(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactDisposition = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactId(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactType(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetDialingTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.DialingTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetEarlyMediaState(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.EarlyMediaState = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetEstablishedTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.EstablishedTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetHeldTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.HeldTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetInstanceId(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetIvrTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.IvrTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetQueueTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.QueueTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRecordingDuration(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.RecordingDuration = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRecordingReady(v bool) *ListCallDetailRecordsResponseBodyDataList {
	s.RecordingReady = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetReleaseInitiator(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ReleaseInitiator = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetReleaseReason(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ReleaseReason = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetReleaseTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRingTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.RingTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionDescription(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionDescription = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionIndex(v int32) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionSurveyOffered(v bool) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionSurveyOffered = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSkillGroupIds(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSkillGroupNames(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SkillGroupNames = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetStartTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetTalkTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.TalkTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetWaitTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.WaitTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
