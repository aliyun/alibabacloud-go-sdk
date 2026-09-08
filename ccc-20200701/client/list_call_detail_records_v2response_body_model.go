// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallDetailRecordsV2ResponseBody
	GetCode() *string
	SetData(v *ListCallDetailRecordsV2ResponseBodyData) *ListCallDetailRecordsV2ResponseBody
	GetData() *ListCallDetailRecordsV2ResponseBodyData
	SetHttpStatusCode(v int32) *ListCallDetailRecordsV2ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallDetailRecordsV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCallDetailRecordsV2ResponseBody
	GetRequestId() *string
}

type ListCallDetailRecordsV2ResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *ListCallDetailRecordsV2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01B12EE4-6AF2-4730-8B78-EC15F4E5C025
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallDetailRecordsV2ResponseBody) GetData() *ListCallDetailRecordsV2ResponseBodyData {
	return s.Data
}

func (s *ListCallDetailRecordsV2ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallDetailRecordsV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallDetailRecordsV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallDetailRecordsV2ResponseBody) SetCode(v string) *ListCallDetailRecordsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBody) SetData(v *ListCallDetailRecordsV2ResponseBodyData) *ListCallDetailRecordsV2ResponseBody {
	s.Data = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBody) SetHttpStatusCode(v int32) *ListCallDetailRecordsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBody) SetMessage(v string) *ListCallDetailRecordsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBody) SetRequestId(v string) *ListCallDetailRecordsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCallDetailRecordsV2ResponseBodyData struct {
	// The list of call records.
	List []*ListCallDetailRecordsV2ResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyData) GetList() []*ListCallDetailRecordsV2ResponseBodyDataList {
	return s.List
}

func (s *ListCallDetailRecordsV2ResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsV2ResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCallDetailRecordsV2ResponseBodyData) SetList(v []*ListCallDetailRecordsV2ResponseBodyDataList) *ListCallDetailRecordsV2ResponseBodyData {
	s.List = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyData) SetPageNumber(v int32) *ListCallDetailRecordsV2ResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyData) SetPageSize(v int32) *ListCallDetailRecordsV2ResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyData) SetTotalCount(v int32) *ListCallDetailRecordsV2ResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyData) Validate() error {
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

type ListCallDetailRecordsV2ResponseBodyDataList struct {
	// The channel ID.
	//
	// example:
	//
	// 8f6bcbad-6e2d-4ca3-9b65-91cf0aafdf00
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// The channel name. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// TestChannel
	AccessChannelName *string `json:"AccessChannelName,omitempty" xml:"AccessChannelName,omitempty"`
	// The channel type. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// Web
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// The visitor ID. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// test-user-id
	AccessChannelUserId *string `json:"AccessChannelUserId,omitempty" xml:"AccessChannelUserId,omitempty"`
	// The visitor name. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// TestVisitor
	AccessChannelUserName *string `json:"AccessChannelUserName,omitempty" xml:"AccessChannelUserName,omitempty"`
	// The additional intermediate number, which may be used in dual-call scenarios.
	//
	// example:
	//
	// 0533128****
	AdditionalBroker *string `json:"AdditionalBroker,omitempty" xml:"AdditionalBroker,omitempty"`
	// The list of agent IDs. Multiple values are separated by commas.
	//
	// example:
	//
	// agent@ccc-test
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// The list of agent names involved in the call. Multiple agents are separated by commas.
	//
	// example:
	//
	// CloudCallCenterTestAgent
	AgentNames *string `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	// The post-call analytics report.
	AnalyticsReport *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport `json:"AnalyticsReport,omitempty" xml:"AnalyticsReport,omitempty" type:"Struct"`
	// Indicates whether the AI post-call analytics report is complete.
	//
	// example:
	//
	// true
	AnalyticsReportReady *bool `json:"AnalyticsReportReady,omitempty" xml:"AnalyticsReportReady,omitempty"`
	// The intermediate number.
	//
	// example:
	//
	// 053xxxx3127
	Broker *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	// The call duration. For inbound calls, timing starts when the call enters the IVR. For outbound calls, timing starts when the call is connected. Unit: seconds.
	//
	// example:
	//
	// 16
	CallDuration *string `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// The list of SIP CallIds.
	//
	// example:
	//
	// dxxx1sdf,xkkwwwa
	CallIds *string `json:"CallIds,omitempty" xml:"CallIds,omitempty"`
	// The called number.
	//
	// example:
	//
	// 1332315****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The location of the callee.
	//
	// example:
	//
	// Shandong-Jinan
	CalleeLocation *string `json:"CalleeLocation,omitempty" xml:"CalleeLocation,omitempty"`
	// The location of the caller.
	//
	// example:
	//
	// Beijing-Beijing
	CallerLocation *string `json:"CallerLocation,omitempty" xml:"CallerLocation,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 0533128****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The visitor client application name. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// Unknown
	ClientAppName *string `json:"ClientAppName,omitempty" xml:"ClientAppName,omitempty"`
	// The visitor client IP address. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// 10.100.2.1
	ClientIpAddress *string `json:"ClientIpAddress,omitempty" xml:"ClientIpAddress,omitempty"`
	// The visitor client location. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// ---
	ClientLocation *string `json:"ClientLocation,omitempty" xml:"ClientLocation,omitempty"`
	// The visitor client user agent information. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36
	ClientUserAgent *string `json:"ClientUserAgent,omitempty" xml:"ClientUserAgent,omitempty"`
	// The call end reason. For the enumeration list, refer to the ContactDispositionList query parameter.
	//
	// example:
	//
	// Success
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-2255019651513856
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The call type. For the enumeration list, refer to the ContactTypeList query parameter.
	//
	// example:
	//
	// OUTBOUND
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// The dialing duration. Unit: seconds.
	//
	// example:
	//
	// 0
	DialingTime *int64 `json:"DialingTime,omitempty" xml:"DialingTime,omitempty"`
	// The early media state.
	//
	// example:
	//
	// NotConnected
	EarlyMediaState *string `json:"EarlyMediaState,omitempty" xml:"EarlyMediaState,omitempty"`
	// The early media text.
	//
	// example:
	//
	// Currently on a call
	EarlyMediaText *string `json:"EarlyMediaText,omitempty" xml:"EarlyMediaText,omitempty"`
	// The time when the call was established. If the call was not established, this value is empty. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1532448000000
	EstablishedTime *int64 `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	// The time taken for the agent to first respond to the customer message. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// 10
	FirstResponseTime *int64 `json:"FirstResponseTime,omitempty" xml:"FirstResponseTime,omitempty"`
	// The hold time. Unit: seconds.
	//
	// example:
	//
	// 12
	HeldTime *int64 `json:"HeldTime,omitempty" xml:"HeldTime,omitempty"`
	// The call center instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IVR duration, which is the time from when the call enters the IVR to when it starts entering the manual service queue (or the customer abandons the call). Unit: seconds.
	//
	// example:
	//
	// 8
	IvrTime *int64 `json:"IvrTime,omitempty" xml:"IvrTime,omitempty"`
	// The media type.
	//
	// Valid values:
	//
	// - AUDIO: Voice.
	//
	// - CHAT: Message.
	//
	// - VIDEO: Video.
	//
	// example:
	//
	// CHAT
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The total number of messages. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// 10
	MessagesSent *int64 `json:"MessagesSent,omitempty" xml:"MessagesSent,omitempty"`
	// The total number of messages sent by the agent. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// 5
	MessagesSentByAgent *int64 `json:"MessagesSentByAgent,omitempty" xml:"MessagesSentByAgent,omitempty"`
	// The total number of messages sent by the customer. This parameter has a value only when MediaType is CHAT.
	//
	// example:
	//
	// 5
	MessagesSentByCustomer *int64 `json:"MessagesSentByCustomer,omitempty" xml:"MessagesSentByCustomer,omitempty"`
	// The list of off-site agent IDs involved in the call. Multiple agents are separated by commas.
	//
	// example:
	//
	// skg-default@ccc-test
	OffSiteAgentIds *string `json:"OffSiteAgentIds,omitempty" xml:"OffSiteAgentIds,omitempty"`
	// The off-site agent number that is called when the assigned agent is an off-site agent.
	//
	// example:
	//
	// 80312348
	OffsiteAgentDestinationNumbers *string `json:"OffsiteAgentDestinationNumbers,omitempty" xml:"OffsiteAgentDestinationNumbers,omitempty"`
	// The number that initiates the call to the off-site agent when the assigned agent is an off-site agent.
	//
	// example:
	//
	// 0101257****
	OffsiteAgentOriginatorNumbers *string `json:"OffsiteAgentOriginatorNumbers,omitempty" xml:"OffsiteAgentOriginatorNumbers,omitempty"`
	// The hangup reason of the last assigned offsite agent.
	//
	// example:
	//
	// IVRException
	OffsiteAgentReleaseReason *string `json:"OffsiteAgentReleaseReason,omitempty" xml:"OffsiteAgentReleaseReason,omitempty"`
	// The called number when the call is transferred to an outside line.
	//
	// example:
	//
	// 134xxxxxx
	OutsideNumberDestinationNumber *string `json:"OutsideNumberDestinationNumber,omitempty" xml:"OutsideNumberDestinationNumber,omitempty"`
	// The hangup reason when the call is transferred to an outside line.
	//
	// example:
	//
	// NoAnswer
	OutsideNumberReleaseReason *string `json:"OutsideNumberReleaseReason,omitempty" xml:"OutsideNumberReleaseReason,omitempty"`
	// The queue duration, which is the time from when the call enters the manual service queue to when the agent starts ringing (or the customer abandons the call, the queue times out, or the queue overflows). Unit: seconds.
	//
	// example:
	//
	// 0
	QueueTime *int64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	// The recording duration. Unit: seconds.
	//
	// example:
	//
	// 10
	RecordingDuration *int64 `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	// Indicates whether the recording has been generated. If the call was not established, false is returned.
	//
	// example:
	//
	// true
	RecordingReady *bool `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	// The party that hung up.
	//
	// Valid values:
	//
	// - agent: The agent.
	//
	// - customer: The customer.
	//
	// example:
	//
	// customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The hangup reason of the call channel, indicating why the current call channel was hung up. The value is a response code defined in the SIP protocol. Refer to the SIP protocol to analyze the hangup reason.
	//
	// example:
	//
	// 486:USER_BUSY
	ReleaseReason *string `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	// The time when the call ended, in Unix timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1532707199000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The ringing duration, which is the time from when the agent starts ringing to when the agent answers (or the customer abandons the call). Unit: seconds.
	//
	// example:
	//
	// 5
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// The satisfaction description, which corresponds to the configuration of the satisfaction module in the satisfaction IVR and is defined by the customer.
	//
	// example:
	//
	// Satisfied
	SatisfactionDescription *string `json:"SatisfactionDescription,omitempty" xml:"SatisfactionDescription,omitempty"`
	// The satisfaction rating, which is the value of the satisfaction key digit (a single digit).
	//
	// example:
	//
	// 1
	SatisfactionIndex *int32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// The satisfaction survey channel.
	//
	// example:
	//
	// IVR
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	// Indicates whether a satisfaction survey was sent.
	//
	// example:
	//
	// true
	SatisfactionSurveyOffered *bool `json:"SatisfactionSurveyOffered,omitempty" xml:"SatisfactionSurveyOffered,omitempty"`
	// The skill group IDs of the agents who participated in the call. Multiple skill groups are separated by commas.
	//
	// example:
	//
	// skg-default@ccc-test
	SkillGroupIds *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	// The skill group names of the agents who participated in the call. Multiple skill groups are separated by commas.
	//
	// example:
	//
	// TestSkillGroup1,TestSkillGroup2
	SkillGroupNames *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	// The time when the call started. For inbound calls, this is counted from when the call enters the IVR. For outbound calls, this is counted from when dialing begins. The value is in Unix timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1631440860000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The summary index.
	SummaryIndex *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex `json:"SummaryIndex,omitempty" xml:"SummaryIndex,omitempty" type:"Struct"`
	// The talk time. Unit: seconds.
	//
	// example:
	//
	// 0
	TalkTime *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// The number of transfers.
	//
	// example:
	//
	// 1
	TransferCount *int64 `json:"TransferCount,omitempty" xml:"TransferCount,omitempty"`
	// The outbound called number carried when the call is transferred from an intelligent outbound call.
	//
	// example:
	//
	// 134xxxxxx
	VoicebotDestinationNumber *string `json:"VoicebotDestinationNumber,omitempty" xml:"VoicebotDestinationNumber,omitempty"`
	// The outbound caller number carried when the call is transferred from an intelligent outbound call.
	//
	// example:
	//
	// 021xxxxxxx
	VoicebotOriginatorNumber *string `json:"VoicebotOriginatorNumber,omitempty" xml:"VoicebotOriginatorNumber,omitempty"`
	// The customer wait time, which equals QueueTime + RingTime. If the customer abandons the call during the waiting period, the value is 0. Unit: seconds.
	//
	// example:
	//
	// 5
	WaitTime *int64 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAccessChannelName() *string {
	return s.AccessChannelName
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAccessChannelUserId() *string {
	return s.AccessChannelUserId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAccessChannelUserName() *string {
	return s.AccessChannelUserName
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAdditionalBroker() *string {
	return s.AdditionalBroker
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAgentIds() *string {
	return s.AgentIds
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAgentNames() *string {
	return s.AgentNames
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAnalyticsReport() *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport {
	return s.AnalyticsReport
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetAnalyticsReportReady() *bool {
	return s.AnalyticsReportReady
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetBroker() *string {
	return s.Broker
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCallDuration() *string {
	return s.CallDuration
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCallIds() *string {
	return s.CallIds
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCalleeLocation() *string {
	return s.CalleeLocation
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCallerLocation() *string {
	return s.CallerLocation
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetClientAppName() *string {
	return s.ClientAppName
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetClientIpAddress() *string {
	return s.ClientIpAddress
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetClientLocation() *string {
	return s.ClientLocation
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetClientUserAgent() *string {
	return s.ClientUserAgent
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetContactDisposition() *string {
	return s.ContactDisposition
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetContactId() *string {
	return s.ContactId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetContactType() *string {
	return s.ContactType
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetDialingTime() *int64 {
	return s.DialingTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetEarlyMediaState() *string {
	return s.EarlyMediaState
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetEarlyMediaText() *string {
	return s.EarlyMediaText
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetEstablishedTime() *int64 {
	return s.EstablishedTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetFirstResponseTime() *int64 {
	return s.FirstResponseTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetHeldTime() *int64 {
	return s.HeldTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetIvrTime() *int64 {
	return s.IvrTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetMediaType() *string {
	return s.MediaType
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetMessagesSent() *int64 {
	return s.MessagesSent
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetMessagesSentByAgent() *int64 {
	return s.MessagesSentByAgent
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetMessagesSentByCustomer() *int64 {
	return s.MessagesSentByCustomer
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOffSiteAgentIds() *string {
	return s.OffSiteAgentIds
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOffsiteAgentDestinationNumbers() *string {
	return s.OffsiteAgentDestinationNumbers
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOffsiteAgentOriginatorNumbers() *string {
	return s.OffsiteAgentOriginatorNumbers
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOffsiteAgentReleaseReason() *string {
	return s.OffsiteAgentReleaseReason
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOutsideNumberDestinationNumber() *string {
	return s.OutsideNumberDestinationNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetOutsideNumberReleaseReason() *string {
	return s.OutsideNumberReleaseReason
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetQueueTime() *int64 {
	return s.QueueTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetRecordingDuration() *int64 {
	return s.RecordingDuration
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetRecordingReady() *bool {
	return s.RecordingReady
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetReleaseReason() *string {
	return s.ReleaseReason
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetRingTime() *int64 {
	return s.RingTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSatisfactionDescription() *string {
	return s.SatisfactionDescription
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSatisfactionIndex() *int32 {
	return s.SatisfactionIndex
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSatisfactionSurveyChannel() *string {
	return s.SatisfactionSurveyChannel
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSatisfactionSurveyOffered() *bool {
	return s.SatisfactionSurveyOffered
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSkillGroupNames() *string {
	return s.SkillGroupNames
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetSummaryIndex() *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex {
	return s.SummaryIndex
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetTransferCount() *int64 {
	return s.TransferCount
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetVoicebotDestinationNumber() *string {
	return s.VoicebotDestinationNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetVoicebotOriginatorNumber() *string {
	return s.VoicebotOriginatorNumber
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) GetWaitTime() *int64 {
	return s.WaitTime
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAccessChannelId(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AccessChannelId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAccessChannelName(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AccessChannelName = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAccessChannelType(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AccessChannelType = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAccessChannelUserId(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AccessChannelUserId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAccessChannelUserName(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AccessChannelUserName = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAdditionalBroker(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AdditionalBroker = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAgentIds(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AgentIds = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAgentNames(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AgentNames = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAnalyticsReport(v *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AnalyticsReport = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetAnalyticsReportReady(v bool) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.AnalyticsReportReady = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetBroker(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.Broker = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCallDuration(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CallDuration = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCallIds(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CallIds = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCalledNumber(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCalleeLocation(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CalleeLocation = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCallerLocation(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CallerLocation = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetCallingNumber(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetClientAppName(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ClientAppName = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetClientIpAddress(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ClientIpAddress = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetClientLocation(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ClientLocation = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetClientUserAgent(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ClientUserAgent = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetContactDisposition(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ContactDisposition = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetContactId(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetContactType(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetDialingTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.DialingTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetEarlyMediaState(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.EarlyMediaState = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetEarlyMediaText(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.EarlyMediaText = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetEstablishedTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.EstablishedTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetFirstResponseTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.FirstResponseTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetHeldTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.HeldTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetInstanceId(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetIvrTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.IvrTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetMediaType(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.MediaType = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetMessagesSent(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.MessagesSent = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetMessagesSentByAgent(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.MessagesSentByAgent = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetMessagesSentByCustomer(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.MessagesSentByCustomer = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOffSiteAgentIds(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OffSiteAgentIds = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOffsiteAgentDestinationNumbers(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OffsiteAgentDestinationNumbers = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOffsiteAgentOriginatorNumbers(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OffsiteAgentOriginatorNumbers = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOffsiteAgentReleaseReason(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OffsiteAgentReleaseReason = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOutsideNumberDestinationNumber(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OutsideNumberDestinationNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetOutsideNumberReleaseReason(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.OutsideNumberReleaseReason = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetQueueTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.QueueTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetRecordingDuration(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.RecordingDuration = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetRecordingReady(v bool) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.RecordingReady = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetReleaseInitiator(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ReleaseInitiator = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetReleaseReason(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ReleaseReason = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetReleaseTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetRingTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.RingTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSatisfactionDescription(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SatisfactionDescription = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSatisfactionIndex(v int32) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSatisfactionSurveyOffered(v bool) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SatisfactionSurveyOffered = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSkillGroupIds(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSkillGroupNames(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SkillGroupNames = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetStartTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetSummaryIndex(v *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.SummaryIndex = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetTalkTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.TalkTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetTransferCount(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.TransferCount = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetVoicebotDestinationNumber(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.VoicebotDestinationNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetVoicebotOriginatorNumber(v string) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.VoicebotOriginatorNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) SetWaitTime(v int64) *ListCallDetailRecordsV2ResponseBodyDataList {
	s.WaitTime = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataList) Validate() error {
	if s.AnalyticsReport != nil {
		if err := s.AnalyticsReport.Validate(); err != nil {
			return err
		}
	}
	if s.SummaryIndex != nil {
		if err := s.SummaryIndex.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport struct {
	// The emotion analysis.
	Emotion *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion `json:"Emotion,omitempty" xml:"Emotion,omitempty" type:"Struct"`
	// The problem resolution status.
	ProblemSolving *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving `json:"ProblemSolving,omitempty" xml:"ProblemSolving,omitempty" type:"Struct"`
	// The satisfaction analysis.
	Satisfaction *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction `json:"Satisfaction,omitempty" xml:"Satisfaction,omitempty" type:"Struct"`
	// The to-do list.
	TodoList *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Struct"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) GetEmotion() *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	return s.Emotion
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) GetProblemSolving() *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	return s.ProblemSolving
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) GetSatisfaction() *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction {
	return s.Satisfaction
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) GetTodoList() *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList {
	return s.TodoList
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) SetEmotion(v *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport {
	s.Emotion = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) SetProblemSolving(v *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport {
	s.ProblemSolving = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) SetSatisfaction(v *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport {
	s.Satisfaction = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) SetTodoList(v *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport {
	s.TodoList = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReport) Validate() error {
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

type ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion struct {
	// The confidence level.
	//
	// example:
	//
	// 50
	Confidence *int32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The detailed remark of the emotion analysis.
	//
	// example:
	//
	// No emotional change from the customer
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 0ff07fe35670423089dbdf12766d962f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The emotion type.
	//
	// example:
	//
	// Neutral
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GetConfidence() *int32 {
	return s.Confidence
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GetRemark() *string {
	return s.Remark
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GetSuccess() *bool {
	return s.Success
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) GetType() *string {
	return s.Type
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) SetConfidence(v int32) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	s.Confidence = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) SetRemark(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	s.Remark = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) SetSuccess(v bool) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	s.Success = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) SetTaskId(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	s.TaskId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) SetType(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion {
	s.Type = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportEmotion) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving struct {
	// The problem description.
	//
	// example:
	//
	// Alert issue
	Problem *string `json:"Problem,omitempty" xml:"Problem,omitempty"`
	// The Solutions.
	//
	// example:
	//
	// The enrichment service automatically closes the original alert
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// Indicates whether the problem is resolved.
	//
	// example:
	//
	// true
	Solved *bool `json:"Solved,omitempty" xml:"Solved,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8bf18b7a10064b29a75946a8d5b8469a
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GetProblem() *string {
	return s.Problem
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GetSolution() *string {
	return s.Solution
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GetSolved() *bool {
	return s.Solved
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GetSuccess() *bool {
	return s.Success
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) SetProblem(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	s.Problem = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) SetSolution(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	s.Solution = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) SetSolved(v bool) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	s.Solved = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) SetSuccess(v bool) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	s.Success = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) SetTaskId(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving {
	s.TaskId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportProblemSolving) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction struct {
	// The detailed remark of the satisfaction analysis.
	//
	// example:
	//
	// The customer expressed satisfaction
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The satisfaction description.
	//
	// example:
	//
	// Satisfied
	SatisfactionDescription *string `json:"SatisfactionDescription,omitempty" xml:"SatisfactionDescription,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// cb67479ce28243b28ff39948feaa0806
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) GetRemark() *string {
	return s.Remark
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) GetSatisfactionDescription() *string {
	return s.SatisfactionDescription
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) GetSuccess() *bool {
	return s.Success
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) SetRemark(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction {
	s.Remark = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) SetSatisfactionDescription(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction {
	s.SatisfactionDescription = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) SetSuccess(v bool) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction {
	s.Success = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) SetTaskId(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction {
	s.TaskId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportSatisfaction) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList struct {
	// Indicates whether the call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID.
	//
	// example:
	//
	// cb67479ce28243b28ff39948feaa0806
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The to-do items.
	Tasks []*string `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) GetSuccess() *bool {
	return s.Success
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) GetTasks() []*string {
	return s.Tasks
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) SetSuccess(v bool) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList {
	s.Success = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) SetTaskId(v string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList {
	s.TaskId = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) SetTasks(v []*string) *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList {
	s.Tasks = v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListAnalyticsReportTodoList) Validate() error {
	return dara.Validate(s)
}

type ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex struct {
	// The keywords.
	//
	// example:
	//
	// CustomerService
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
}

func (s ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) GetKeywords() *string {
	return s.Keywords
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) SetKeywords(v string) *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex {
	s.Keywords = &v
	return s
}

func (s *ListCallDetailRecordsV2ResponseBodyDataListSummaryIndex) Validate() error {
	return dara.Validate(s)
}
