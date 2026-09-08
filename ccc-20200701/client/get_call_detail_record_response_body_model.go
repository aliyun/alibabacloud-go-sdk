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
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *GetCallDetailRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyData struct {
	// The list of agent events.
	AgentEvents []*GetCallDetailRecordResponseBodyDataAgentEvents `json:"AgentEvents,omitempty" xml:"AgentEvents,omitempty" type:"Repeated"`
	// The IDs of the agents who are involved in the call. Multiple IDs are separated by commas.
	//
	// example:
	//
	// agent1@ccc-test,agent2@ccc-test
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// The names of the agents who are involved in the call. Multiple names are separated by commas.
	//
	// example:
	//
	// agent1,agent2
	AgentNames           *string                                             `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	AnalyticsReport      *GetCallDetailRecordResponseBodyDataAnalyticsReport `json:"AnalyticsReport,omitempty" xml:"AnalyticsReport,omitempty" type:"Struct"`
	AnalyticsReportReady *bool                                               `json:"AnalyticsReportReady,omitempty" xml:"AnalyticsReportReady,omitempty"`
	// The call duration, in seconds.
	//
	// example:
	//
	// 50
	CallDuration *int64 `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// The called number.
	//
	// example:
	//
	// 1332315****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The location of the called number.
	//
	// example:
	//
	// 河北省-唐山
	CalleeLocation *string `json:"CalleeLocation,omitempty" xml:"CalleeLocation,omitempty"`
	// The location of the calling number.
	//
	// example:
	//
	// 山东省-淄博
	CallerLocation *string `json:"CallerLocation,omitempty" xml:"CallerLocation,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 0533128****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The reason why the call ended. Note: The \\`Voicemail\\`, \\`QueuingFailed\\`, \\`QueuingTimeout\\`, \\`QueuingOverflow\\`, and \\`IVRException\\` reasons are returned only if you configure the hang-up reason node. If you do not configure this node and the IVR flow does not include a module to transfer the call to an agent, the default reason is \\`AbandonedInIVR\\`.
	//
	// example:
	//
	// Success
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-10963442671187****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The call type.
	//
	// example:
	//
	// OUTBOUND
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// The list of customer events.
	CustomerEvents []*GetCallDetailRecordResponseBodyDataCustomerEvents `json:"CustomerEvents,omitempty" xml:"CustomerEvents,omitempty" type:"Repeated"`
	// The state of the early media. An exception occurred during the early media phase, which is when the customer is being called. An exception at this stage can cause the call to fail. This parameter provides possible reasons for the connection failure based on an analysis of the early media state.
	//
	// example:
	//
	// NotConnected
	EarlyMediaState *string `json:"EarlyMediaState,omitempty" xml:"EarlyMediaState,omitempty"`
	// The time when the call was connected. This parameter is empty if the call was not connected. The value is a UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1532458000000
	EstablishedTime *int64 `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IVR events.
	IvrEvents                  []*GetCallDetailRecordResponseBodyDataIvrEvents `json:"IvrEvents,omitempty" xml:"IvrEvents,omitempty" type:"Repeated"`
	OutsideNumberReleaseReason *string                                         `json:"OutsideNumberReleaseReason,omitempty" xml:"OutsideNumberReleaseReason,omitempty"`
	// The list of queue events.
	QueueEvents []*GetCallDetailRecordResponseBodyDataQueueEvents `json:"QueueEvents,omitempty" xml:"QueueEvents,omitempty" type:"Repeated"`
	// Indicates whether the recording was generated. A value of \\`false\\` is returned if the call was not connected.
	//
	// example:
	//
	// true
	RecordingReady *bool `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	// The release initiator.
	//
	// example:
	//
	// customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The reason why the call ended. The value is usually the SIP code followed by a text description.
	//
	// example:
	//
	// 200 - OK
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The time when the call ended. This is the time when the last party of the call hangs up. The value is a UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1532458000000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The satisfaction score. The value and its meaning are defined by you.
	//
	// example:
	//
	// 1
	Satisfaction *int32 `json:"Satisfaction,omitempty" xml:"Satisfaction,omitempty"`
	// The channel through which the satisfaction survey was initiated.
	//
	// example:
	//
	// IVR
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	// Indicates whether a satisfaction survey was initiated.
	//
	// example:
	//
	// true
	SatisfactionSurveyOffered *bool `json:"SatisfactionSurveyOffered,omitempty" xml:"SatisfactionSurveyOffered,omitempty"`
	// The IDs of the skill groups to which the agents involved in the call belong. Multiple IDs are separated by commas.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupIds *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	// The names of the skill groups to which the agents involved in the call belong. Multiple names are separated by commas.
	//
	// example:
	//
	// 测试技能组
	SkillGroupNames *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	// The time when the call started. For an inbound call, this is the time when the call enters the IVR. For an outbound call, this is the time when the call is initiated. The value is a UNIX timestamp, in milliseconds.
	//
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
	if s.AgentEvents != nil {
		for _, item := range s.AgentEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AnalyticsReport != nil {
		if err := s.AnalyticsReport.Validate(); err != nil {
			return err
		}
	}
	if s.CustomerEvents != nil {
		for _, item := range s.CustomerEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IvrEvents != nil {
		for _, item := range s.IvrEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueueEvents != nil {
		for _, item := range s.QueueEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataAgentEvents struct {
	// The agent ID.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// 坐席小王
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The sequence of events.
	EventSequence []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// The skill group ID.
	//
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
	if s.EventSequence != nil {
		for _, item := range s.EventSequence {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataAgentEventsEventSequence struct {
	// The event duration, in seconds.
	//
	// example:
	//
	// 3
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The event type.
	//
	// example:
	//
	// Dialing
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The time when the event occurred. The value is a UNIX timestamp, in milliseconds.
	//
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
	if s.Emotion != nil {
		if err := s.Emotion.Validate(); err != nil {
			return err
		}
	}
	if s.ProblemSolving != nil {
		if err := s.ProblemSolving.Validate(); err != nil {
			return err
		}
	}
	if s.Satisfaction != nil {
		if err := s.Satisfaction.Validate(); err != nil {
			return err
		}
	}
	if s.TodoList != nil {
		if err := s.TodoList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// The customer ID. This is usually the customer\\"s phone number.
	//
	// example:
	//
	// 1332315****
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The sequence of events.
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
	if s.EventSequence != nil {
		for _, item := range s.EventSequence {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataCustomerEventsEventSequence struct {
	// The event type.
	//
	// example:
	//
	// Released
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The time when the event occurred. The value is a UNIX timestamp, in milliseconds.
	//
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
	// The sequence of events.
	EventSequence []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// The ID of the IVR contact flow.
	//
	// example:
	//
	// edaf2eaa-8f88-44ca-812e-41b3cd2b7a90
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The type of the contact flow.
	//
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
	if s.EventSequence != nil {
		for _, item := range s.EventSequence {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataIvrEventsEventSequence struct {
	// The event type.
	//
	// example:
	//
	// Route2IVR
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The time when the event occurred. The value is a UNIX timestamp, in milliseconds.
	//
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
	// The sequence of events.
	EventSequence []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
	// The contact flow ID.
	//
	// example:
	//
	// edaf2eaa-8f88-44ca-812e-41b3cd2b7a90
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The queue ID. If the call is routed to a skill group, this is the skill group ID. If the call is routed to an agent, this is the agent ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The queue name.
	//
	// example:
	//
	// 测试技能组
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The queue type.
	//
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
	if s.EventSequence != nil {
		for _, item := range s.EventSequence {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataQueueEventsEventSequence struct {
	// The event type.
	//
	// example:
	//
	// Enqueue
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The time when the event occurred. The value is a UNIX timestamp, in milliseconds.
	//
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
