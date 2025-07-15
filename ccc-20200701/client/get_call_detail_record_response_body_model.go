// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDetailRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCallDetailRecordResponseBody
	GetCode() *string
	SetData(v *GetCallDetailRecordResponseBodyData) *GetCallDetailRecordResponseBody
	GetData() *GetCallDetailRecordResponseBodyData
	SetHttpStatusCode(v int32) *GetCallDetailRecordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCallDetailRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCallDetailRecordResponseBody
	GetRequestId() *string
}

type GetCallDetailRecordResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCallDetailRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCallDetailRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallDetailRecordResponseBody) GetData() *GetCallDetailRecordResponseBodyData {
	return s.Data
}

func (s *GetCallDetailRecordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCallDetailRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallDetailRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCallDetailRecordResponseBody) SetCode(v string) *GetCallDetailRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetData(v *GetCallDetailRecordResponseBodyData) *GetCallDetailRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetHttpStatusCode(v int32) *GetCallDetailRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetMessage(v string) *GetCallDetailRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetRequestId(v string) *GetCallDetailRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyData struct {
	AgentEvents []*GetCallDetailRecordResponseBodyDataAgentEvents `json:"AgentEvents,omitempty" xml:"AgentEvents,omitempty" type:"Repeated"`
	// example:
	//
	// agent1@ccc-test,agent2@ccc-test
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// example:
	//
	// agent1,agent2
	AgentNames           *string                                             `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	AnalyticsReport      *GetCallDetailRecordResponseBodyDataAnalyticsReport `json:"AnalyticsReport,omitempty" xml:"AnalyticsReport,omitempty" type:"Struct"`
	AnalyticsReportReady *bool                                               `json:"AnalyticsReportReady,omitempty" xml:"AnalyticsReportReady,omitempty"`
	// example:
	//
	// 50
	CallDuration *int64 `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
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
	// job-10963442671187****
	ContactId      *string                                              `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactType    *string                                              `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	CustomerEvents []*GetCallDetailRecordResponseBodyDataCustomerEvents `json:"CustomerEvents,omitempty" xml:"CustomerEvents,omitempty" type:"Repeated"`
	// example:
	//
	// NotConnected
	EarlyMediaState *string `json:"EarlyMediaState,omitempty" xml:"EarlyMediaState,omitempty"`
	// example:
	//
	// 1532458000000
	EstablishedTime *int64 `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId                 *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IvrEvents                  []*GetCallDetailRecordResponseBodyDataIvrEvents   `json:"IvrEvents,omitempty" xml:"IvrEvents,omitempty" type:"Repeated"`
	OutsideNumberReleaseReason *string                                           `json:"OutsideNumberReleaseReason,omitempty" xml:"OutsideNumberReleaseReason,omitempty"`
	QueueEvents                []*GetCallDetailRecordResponseBodyDataQueueEvents `json:"QueueEvents,omitempty" xml:"QueueEvents,omitempty" type:"Repeated"`
	// example:
	//
	// true
	RecordingReady *bool `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	// example:
	//
	// customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// example:
	//
	// 200 - OK
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// example:
	//
	// 1532458000000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// 1
	Satisfaction *int32 `json:"Satisfaction,omitempty" xml:"Satisfaction,omitempty"`
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
	// 1532458000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyData) GetAgentEvents() []*GetCallDetailRecordResponseBodyDataAgentEvents {
	return s.AgentEvents
}

func (s *GetCallDetailRecordResponseBodyData) GetAgentIds() *string {
	return s.AgentIds
}

func (s *GetCallDetailRecordResponseBodyData) GetAgentNames() *string {
	return s.AgentNames
}

func (s *GetCallDetailRecordResponseBodyData) GetAnalyticsReport() *GetCallDetailRecordResponseBodyDataAnalyticsReport {
	return s.AnalyticsReport
}

func (s *GetCallDetailRecordResponseBodyData) GetAnalyticsReportReady() *bool {
	return s.AnalyticsReportReady
}

func (s *GetCallDetailRecordResponseBodyData) GetCallDuration() *int64 {
	return s.CallDuration
}

func (s *GetCallDetailRecordResponseBodyData) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *GetCallDetailRecordResponseBodyData) GetCalleeLocation() *string {
	return s.CalleeLocation
}

func (s *GetCallDetailRecordResponseBodyData) GetCallerLocation() *string {
	return s.CallerLocation
}

func (s *GetCallDetailRecordResponseBodyData) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *GetCallDetailRecordResponseBodyData) GetContactDisposition() *string {
	return s.ContactDisposition
}

func (s *GetCallDetailRecordResponseBodyData) GetContactId() *string {
	return s.ContactId
}

func (s *GetCallDetailRecordResponseBodyData) GetContactType() *string {
	return s.ContactType
}

func (s *GetCallDetailRecordResponseBodyData) GetCustomerEvents() []*GetCallDetailRecordResponseBodyDataCustomerEvents {
	return s.CustomerEvents
}

func (s *GetCallDetailRecordResponseBodyData) GetEarlyMediaState() *string {
	return s.EarlyMediaState
}

func (s *GetCallDetailRecordResponseBodyData) GetEstablishedTime() *int64 {
	return s.EstablishedTime
}

func (s *GetCallDetailRecordResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCallDetailRecordResponseBodyData) GetIvrEvents() []*GetCallDetailRecordResponseBodyDataIvrEvents {
	return s.IvrEvents
}

func (s *GetCallDetailRecordResponseBodyData) GetOutsideNumberReleaseReason() *string {
	return s.OutsideNumberReleaseReason
}

func (s *GetCallDetailRecordResponseBodyData) GetQueueEvents() []*GetCallDetailRecordResponseBodyDataQueueEvents {
	return s.QueueEvents
}

func (s *GetCallDetailRecordResponseBodyData) GetRecordingReady() *bool {
	return s.RecordingReady
}

func (s *GetCallDetailRecordResponseBodyData) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *GetCallDetailRecordResponseBodyData) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *GetCallDetailRecordResponseBodyData) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *GetCallDetailRecordResponseBodyData) GetSatisfaction() *int32 {
	return s.Satisfaction
}

func (s *GetCallDetailRecordResponseBodyData) GetSatisfactionSurveyChannel() *string {
	return s.SatisfactionSurveyChannel
}

func (s *GetCallDetailRecordResponseBodyData) GetSatisfactionSurveyOffered() *bool {
	return s.SatisfactionSurveyOffered
}

func (s *GetCallDetailRecordResponseBodyData) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *GetCallDetailRecordResponseBodyData) GetSkillGroupNames() *string {
	return s.SkillGroupNames
}

func (s *GetCallDetailRecordResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentEvents(v []*GetCallDetailRecordResponseBodyDataAgentEvents) *GetCallDetailRecordResponseBodyData {
	s.AgentEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentIds(v string) *GetCallDetailRecordResponseBodyData {
	s.AgentIds = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentNames(v string) *GetCallDetailRecordResponseBodyData {
	s.AgentNames = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAnalyticsReport(v *GetCallDetailRecordResponseBodyDataAnalyticsReport) *GetCallDetailRecordResponseBodyData {
	s.AnalyticsReport = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAnalyticsReportReady(v bool) *GetCallDetailRecordResponseBodyData {
	s.AnalyticsReportReady = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallDuration(v int64) *GetCallDetailRecordResponseBodyData {
	s.CallDuration = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCalledNumber(v string) *GetCallDetailRecordResponseBodyData {
	s.CalledNumber = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCalleeLocation(v string) *GetCallDetailRecordResponseBodyData {
	s.CalleeLocation = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallerLocation(v string) *GetCallDetailRecordResponseBodyData {
	s.CallerLocation = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallingNumber(v string) *GetCallDetailRecordResponseBodyData {
	s.CallingNumber = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactDisposition(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactDisposition = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactId(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactType(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCustomerEvents(v []*GetCallDetailRecordResponseBodyDataCustomerEvents) *GetCallDetailRecordResponseBodyData {
	s.CustomerEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetEarlyMediaState(v string) *GetCallDetailRecordResponseBodyData {
	s.EarlyMediaState = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetEstablishedTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.EstablishedTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetInstanceId(v string) *GetCallDetailRecordResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetIvrEvents(v []*GetCallDetailRecordResponseBodyDataIvrEvents) *GetCallDetailRecordResponseBodyData {
	s.IvrEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetOutsideNumberReleaseReason(v string) *GetCallDetailRecordResponseBodyData {
	s.OutsideNumberReleaseReason = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetQueueEvents(v []*GetCallDetailRecordResponseBodyDataQueueEvents) *GetCallDetailRecordResponseBodyData {
	s.QueueEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetRecordingReady(v bool) *GetCallDetailRecordResponseBodyData {
	s.RecordingReady = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseInitiator(v string) *GetCallDetailRecordResponseBodyData {
	s.ReleaseInitiator = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseReason(v string) *GetCallDetailRecordResponseBodyData {
	s.ReleaseReason = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfaction(v int32) *GetCallDetailRecordResponseBodyData {
	s.Satisfaction = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfactionSurveyChannel(v string) *GetCallDetailRecordResponseBodyData {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfactionSurveyOffered(v bool) *GetCallDetailRecordResponseBodyData {
	s.SatisfactionSurveyOffered = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSkillGroupIds(v string) *GetCallDetailRecordResponseBodyData {
	s.SkillGroupIds = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSkillGroupNames(v string) *GetCallDetailRecordResponseBodyData {
	s.SkillGroupNames = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetStartTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAgentEvents struct {
	// example:
	//
	// agent@ccc-test
	AgentId       *string                                                        `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName     *string                                                        `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	EventSequence []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAgentEvents) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAgentEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) GetAgentId() *string {
	return s.AgentId
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) GetAgentName() *string {
	return s.AgentName
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) GetEventSequence() []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	return s.EventSequence
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetAgentId(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.AgentId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetAgentName(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.AgentName = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.EventSequence = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetSkillGroupId(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.SkillGroupId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAgentEventsEventSequence struct {
	// example:
	//
	// 3
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Dialing
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1604639129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) GetDuration() *int64 {
	return s.Duration
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) GetEvent() *string {
	return s.Event
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) SetDuration(v int64) *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	s.Duration = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAnalyticsReport struct {
	Emotion        *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion        `json:"Emotion,omitempty" xml:"Emotion,omitempty" type:"Struct"`
	ProblemSolving *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving `json:"ProblemSolving,omitempty" xml:"ProblemSolving,omitempty" type:"Struct"`
	Satisfaction   *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction   `json:"Satisfaction,omitempty" xml:"Satisfaction,omitempty" type:"Struct"`
	TodoList       *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList       `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Struct"`
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReport) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReport) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) GetEmotion() *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	return s.Emotion
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) GetProblemSolving() *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	return s.ProblemSolving
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) GetSatisfaction() *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction {
	return s.Satisfaction
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) GetTodoList() *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList {
	return s.TodoList
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) SetEmotion(v *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) *GetCallDetailRecordResponseBodyDataAnalyticsReport {
	s.Emotion = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) SetProblemSolving(v *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) *GetCallDetailRecordResponseBodyDataAnalyticsReport {
	s.ProblemSolving = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) SetSatisfaction(v *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) *GetCallDetailRecordResponseBodyDataAnalyticsReport {
	s.Satisfaction = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) SetTodoList(v *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) *GetCallDetailRecordResponseBodyDataAnalyticsReport {
	s.TodoList = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReport) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion struct {
	Confidence *int32  `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GetConfidence() *int32 {
	return s.Confidence
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GetRemark() *string {
	return s.Remark
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) GetType() *string {
	return s.Type
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) SetConfidence(v int32) *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	s.Confidence = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) SetRemark(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	s.Remark = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) SetSuccess(v bool) *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	s.Success = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) SetTaskId(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	s.TaskId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) SetType(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion {
	s.Type = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportEmotion) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving struct {
	Problem  *string `json:"Problem,omitempty" xml:"Problem,omitempty"`
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	Solved   *bool   `json:"Solved,omitempty" xml:"Solved,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GetProblem() *string {
	return s.Problem
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GetSolution() *string {
	return s.Solution
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GetSolved() *bool {
	return s.Solved
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) SetProblem(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	s.Problem = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) SetSolution(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	s.Solution = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) SetSolved(v bool) *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	s.Solved = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) SetSuccess(v bool) *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	s.Success = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) SetTaskId(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving {
	s.TaskId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportProblemSolving) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction struct {
	Remark                  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SatisfactionDescription *string `json:"SatisfactionDescription,omitempty" xml:"SatisfactionDescription,omitempty"`
	Success                 *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId                  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) GetRemark() *string {
	return s.Remark
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) GetSatisfactionDescription() *string {
	return s.SatisfactionDescription
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) SetRemark(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction {
	s.Remark = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) SetSatisfactionDescription(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction {
	s.SatisfactionDescription = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) SetSuccess(v bool) *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction {
	s.Success = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) SetTaskId(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction {
	s.TaskId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportSatisfaction) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList struct {
	Success *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId  *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Tasks   []*string `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) GetTasks() []*string {
	return s.Tasks
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) SetSuccess(v bool) *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList {
	s.Success = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) SetTaskId(v string) *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList {
	s.TaskId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) SetTasks(v []*string) *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList {
	s.Tasks = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAnalyticsReportTodoList) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataCustomerEvents struct {
	// example:
	//
	// 1332315****
	CustomerId    *string                                                           `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	EventSequence []*GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
}

func (s GetCallDetailRecordResponseBodyDataCustomerEvents) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataCustomerEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEvents) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEvents) GetEventSequence() []*GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence {
	return s.EventSequence
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEvents) SetCustomerId(v string) *GetCallDetailRecordResponseBodyDataCustomerEvents {
	s.CustomerId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) *GetCallDetailRecordResponseBodyDataCustomerEvents {
	s.EventSequence = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEvents) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence struct {
	// example:
	//
	// Released
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1532458000000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) GetEvent() *string {
	return s.Event
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataIvrEvents struct {
	EventSequence []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// example:
	//
	// edaf2eaa-8f88-44ca-812e-41b3cd2b7a90
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// MAIN_FLOW
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataIvrEvents) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataIvrEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) GetEventSequence() []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence {
	return s.EventSequence
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) GetFlowId() *string {
	return s.FlowId
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) GetFlowType() *string {
	return s.FlowType
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) *GetCallDetailRecordResponseBodyDataIvrEvents {
	s.EventSequence = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) SetFlowId(v string) *GetCallDetailRecordResponseBodyDataIvrEvents {
	s.FlowId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) SetFlowType(v string) *GetCallDetailRecordResponseBodyDataIvrEvents {
	s.FlowType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataIvrEventsEventSequence struct {
	// example:
	//
	// Route2IVR
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1604639129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) GetEvent() *string {
	return s.Event
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataQueueEvents struct {
	EventSequence []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// example:
	//
	// edaf2eaa-8f88-44ca-812e-41b3cd2b7a90
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	QueueId   *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// 1
	QueueType *int32 `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataQueueEvents) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataQueueEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) GetEventSequence() []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence {
	return s.EventSequence
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) GetFlowId() *string {
	return s.FlowId
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) GetQueueId() *string {
	return s.QueueId
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) GetQueueName() *string {
	return s.QueueName
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) GetQueueType() *int32 {
	return s.QueueType
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.EventSequence = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetFlowId(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.FlowId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueId(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueName(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueName = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueType(v int32) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataQueueEventsEventSequence struct {
	// example:
	//
	// Enqueue
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1604639129000
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) GetEvent() *string {
	return s.Event
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) Validate() error {
	return dara.Validate(s)
}
