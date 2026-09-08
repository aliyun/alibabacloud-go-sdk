// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelTypeList(v string) *ListCallDetailRecordsV2Request
	GetAccessChannelTypeList() *string
	SetAgentId(v string) *ListCallDetailRecordsV2Request
	GetAgentId() *string
	SetAnalyticsReportReady(v bool) *ListCallDetailRecordsV2Request
	GetAnalyticsReportReady() *bool
	SetBroker(v string) *ListCallDetailRecordsV2Request
	GetBroker() *string
	SetCalledNumber(v string) *ListCallDetailRecordsV2Request
	GetCalledNumber() *string
	SetCallingNumber(v string) *ListCallDetailRecordsV2Request
	GetCallingNumber() *string
	SetContactDispositionList(v string) *ListCallDetailRecordsV2Request
	GetContactDispositionList() *string
	SetContactIdList(v string) *ListCallDetailRecordsV2Request
	GetContactIdList() *string
	SetContactTypeList(v string) *ListCallDetailRecordsV2Request
	GetContactTypeList() *string
	SetEarlyMediaStateList(v string) *ListCallDetailRecordsV2Request
	GetEarlyMediaStateList() *string
	SetEndTime(v int64) *ListCallDetailRecordsV2Request
	GetEndTime() *int64
	SetFirstAgentId(v string) *ListCallDetailRecordsV2Request
	GetFirstAgentId() *string
	SetInstanceId(v string) *ListCallDetailRecordsV2Request
	GetInstanceId() *string
	SetMediaType(v string) *ListCallDetailRecordsV2Request
	GetMediaType() *string
	SetNumber(v string) *ListCallDetailRecordsV2Request
	GetNumber() *string
	SetOrderByField(v string) *ListCallDetailRecordsV2Request
	GetOrderByField() *string
	SetPageNumber(v int32) *ListCallDetailRecordsV2Request
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallDetailRecordsV2Request
	GetPageSize() *int32
	SetReleaseInitiatorList(v string) *ListCallDetailRecordsV2Request
	GetReleaseInitiatorList() *string
	SetReleaseReasonList(v string) *ListCallDetailRecordsV2Request
	GetReleaseReasonList() *string
	SetSatisfactionDescriptionList(v string) *ListCallDetailRecordsV2Request
	GetSatisfactionDescriptionList() *string
	SetSatisfactionRateList(v string) *ListCallDetailRecordsV2Request
	GetSatisfactionRateList() *string
	SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsV2Request
	GetSatisfactionSurveyChannel() *string
	SetSearchPattern(v string) *ListCallDetailRecordsV2Request
	GetSearchPattern() *string
	SetSkillGroupIdList(v string) *ListCallDetailRecordsV2Request
	GetSkillGroupIdList() *string
	SetSortOrder(v string) *ListCallDetailRecordsV2Request
	GetSortOrder() *string
	SetStartTime(v int64) *ListCallDetailRecordsV2Request
	GetStartTime() *int64
}

type ListCallDetailRecordsV2Request struct {
	// The list of visitor channel types.
	//
	// Valid values:
	//
	// Web: web page.
	//
	// DingTalkServiceWindow: DingTalk service window.
	//
	// AliMe: chatbot.
	//
	// DingTalkRobot: DingTalk chatbot.
	//
	// example:
	//
	// ["Web","AliMe"]
	AccessChannelTypeList *string `json:"AccessChannelTypeList,omitempty" xml:"AccessChannelTypeList,omitempty"`
	// The ID of the agent.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Indicates whether the AI post-call analytics report is ready.
	//
	// example:
	//
	// true
	AnalyticsReportReady *bool `json:"AnalyticsReportReady,omitempty" xml:"AnalyticsReportReady,omitempty"`
	// The intermediate number. The number must be an available outbound number under the instance. In a back-to-back call, this number is used to call the caller and the callee sequentially.
	//
	// example:
	//
	// 021****4972
	Broker *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	// The called number.
	//
	// example:
	//
	// 191***9993
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 191***9993
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Filters results by hangup cause.
	//
	// Valid values:
	//
	// AbandonedInQueue: Abandoned while queuing.
	//
	// NoAnswer: Customer did not answer.
	//
	// QueuingTimeout: Queuing timed out.
	//
	// Voicemail: Transferred to voice mail.
	//
	// QueuingFailed: Transfer to agent failed.
	//
	// QueuingOverflow: Queue overflow.
	//
	// AbandonedInVoiceNavigator: Abandoned during intelligent navigation.
	//
	// Success: Ended normally.
	//
	// IVRException: IVR exception.
	//
	// AbandonedInRing: Abandoned while ringing.
	//
	// AbandonedInIVR: Abandoned in IVR.
	//
	// Reject: Customer rejected the call.
	//
	// ForwardToOutsideNumber: Forwarded to an external number.
	//
	// example:
	//
	// ["Success","NoAnswer"]
	ContactDispositionList *string `json:"ContactDispositionList,omitempty" xml:"ContactDispositionList,omitempty"`
	// The list of call IDs.
	//
	// example:
	//
	// ["job-123456789","job-234567891"]
	ContactIdList *string `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty"`
	// Filters results by call type.
	//
	// Valid values:
	//
	// OUTBOUND: outbound call.
	//
	// BACK2BACK: back-to-back call.
	//
	// PRIVACY_DIAL: encrypted call.
	//
	// INTERNAL: internal call.
	//
	// PREDICTIVE: predictive outbound call.
	//
	// INBOUND: inbound call.
	//
	// CONFERENCE: conference call.
	//
	// example:
	//
	// ["INBOUND","OUTBOUND"]
	ContactTypeList *string `json:"ContactTypeList,omitempty" xml:"ContactTypeList,omitempty"`
	// Filters results by the reason for unanswered calls.
	//
	// Valid values:
	//
	// NoAnswer: No answer.
	//
	// OutOfService: Out of service.
	//
	// NotExist: Number does not exist.
	//
	// Restricted: Call restricted.
	//
	// Busy: Line busy.
	//
	// NotConnected: Cannot be connected.
	//
	// PowerOff: Phone powered off.
	//
	// example:
	//
	// ["NotConnected","NoAnswer"]
	EarlyMediaStateList *string `json:"EarlyMediaStateList,omitempty" xml:"EarlyMediaStateList,omitempty"`
	// The end time of the historical data to retrieve. The default value is the current time. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1657879880010
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the first agent who participated in the call.
	//
	// example:
	//
	// agent@ccc-test
	FirstAgentId *string `json:"FirstAgentId,omitempty" xml:"FirstAgentId,omitempty"`
	// The instance ID of the call center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The media type.
	//
	// Valid values:
	//
	// AUDIO: voice call.
	//
	// VIDEO: video call.
	//
	// CHAT: message.
	//
	// ALL: all types.
	//
	// example:
	//
	// AUDIO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 191***9993
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// The field by which to sort the results. This parameter is optional. Default value: startTime (call start time).
	//
	// Valid values:
	//
	// startTime: call start time.
	//
	// example:
	//
	// startTime
	OrderByField *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	// The page number. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The party that hung up.
	//
	// example:
	//
	// [\\"AGENT\\"]
	ReleaseInitiatorList *string `json:"ReleaseInitiatorList,omitempty" xml:"ReleaseInitiatorList,omitempty"`
	// The hangup reason.
	//
	// example:
	//
	// [\\"VisitorInitiatedClosure\\",\\"AgentInitiatedClosure\\"]
	ReleaseReasonList *string `json:"ReleaseReasonList,omitempty" xml:"ReleaseReasonList,omitempty"`
	// Filters results by satisfaction description. The description content is defined by the customer.
	//
	// example:
	//
	// ["Satisfied","Average"]
	SatisfactionDescriptionList *string `json:"SatisfactionDescriptionList,omitempty" xml:"SatisfactionDescriptionList,omitempty"`
	// Filters results by satisfaction rating. Separate multiple satisfaction ratings with commas (,).
	//
	// Valid values:
	//
	// -2: Not sent.
	//
	// -1: Not rated.
	//
	// Other positive numbers: custom satisfaction rating values.
	//
	// example:
	//
	// ["-1","3"]
	//
	// [">2"]
	//
	// ["<3"]
	SatisfactionRateList *string `json:"SatisfactionRateList,omitempty" xml:"SatisfactionRateList,omitempty"`
	// The satisfaction survey channel.
	//
	// Valid values:
	//
	// IVR: voice-based satisfaction survey.
	//
	// SMS: SMS-based satisfaction survey.
	//
	// CHAT: message-based satisfaction survey.
	//
	// example:
	//
	// IVR
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	// The custom query text that follows the Lucene query syntax.
	//
	// The following table describes the supported query fields and query types:
	//
	// accessChannelUserId: the visitor ID. Supports exact match and fuzzy match.
	//
	// accessChannelName: the channel name. Supports exact match and fuzzy match.
	//
	// text: the session text. Supports fuzzy match.
	//
	// example:
	//
	// text:please wait AND accessChannelUserId:af1a0-afaa-5086e2946e0b	- AND accessChannelName:dedicated development channel*
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// The list of skill group IDs to which the agent belongs.
	//
	// example:
	//
	// [
	//
	//       "skg1@ccc-test",
	//
	//       "skg2@ccc-test"
	//
	// ]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// The sort order. Default value: DESC.
	//
	// Valid values:
	//
	// ASC: ascending order.
	//
	// DESC: descending order.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The start time of the historical data to retrieve. The default value is 00:00 of the current day. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1657853640015
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCallDetailRecordsV2Request) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2Request) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2Request) GetAccessChannelTypeList() *string {
	return s.AccessChannelTypeList
}

func (s *ListCallDetailRecordsV2Request) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCallDetailRecordsV2Request) GetAnalyticsReportReady() *bool {
	return s.AnalyticsReportReady
}

func (s *ListCallDetailRecordsV2Request) GetBroker() *string {
	return s.Broker
}

func (s *ListCallDetailRecordsV2Request) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListCallDetailRecordsV2Request) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListCallDetailRecordsV2Request) GetContactDispositionList() *string {
	return s.ContactDispositionList
}

func (s *ListCallDetailRecordsV2Request) GetContactIdList() *string {
	return s.ContactIdList
}

func (s *ListCallDetailRecordsV2Request) GetContactTypeList() *string {
	return s.ContactTypeList
}

func (s *ListCallDetailRecordsV2Request) GetEarlyMediaStateList() *string {
	return s.EarlyMediaStateList
}

func (s *ListCallDetailRecordsV2Request) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCallDetailRecordsV2Request) GetFirstAgentId() *string {
	return s.FirstAgentId
}

func (s *ListCallDetailRecordsV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsV2Request) GetMediaType() *string {
	return s.MediaType
}

func (s *ListCallDetailRecordsV2Request) GetNumber() *string {
	return s.Number
}

func (s *ListCallDetailRecordsV2Request) GetOrderByField() *string {
	return s.OrderByField
}

func (s *ListCallDetailRecordsV2Request) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsV2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsV2Request) GetReleaseInitiatorList() *string {
	return s.ReleaseInitiatorList
}

func (s *ListCallDetailRecordsV2Request) GetReleaseReasonList() *string {
	return s.ReleaseReasonList
}

func (s *ListCallDetailRecordsV2Request) GetSatisfactionDescriptionList() *string {
	return s.SatisfactionDescriptionList
}

func (s *ListCallDetailRecordsV2Request) GetSatisfactionRateList() *string {
	return s.SatisfactionRateList
}

func (s *ListCallDetailRecordsV2Request) GetSatisfactionSurveyChannel() *string {
	return s.SatisfactionSurveyChannel
}

func (s *ListCallDetailRecordsV2Request) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListCallDetailRecordsV2Request) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListCallDetailRecordsV2Request) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListCallDetailRecordsV2Request) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCallDetailRecordsV2Request) SetAccessChannelTypeList(v string) *ListCallDetailRecordsV2Request {
	s.AccessChannelTypeList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetAgentId(v string) *ListCallDetailRecordsV2Request {
	s.AgentId = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetAnalyticsReportReady(v bool) *ListCallDetailRecordsV2Request {
	s.AnalyticsReportReady = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetBroker(v string) *ListCallDetailRecordsV2Request {
	s.Broker = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetCalledNumber(v string) *ListCallDetailRecordsV2Request {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetCallingNumber(v string) *ListCallDetailRecordsV2Request {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetContactDispositionList(v string) *ListCallDetailRecordsV2Request {
	s.ContactDispositionList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetContactIdList(v string) *ListCallDetailRecordsV2Request {
	s.ContactIdList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetContactTypeList(v string) *ListCallDetailRecordsV2Request {
	s.ContactTypeList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetEarlyMediaStateList(v string) *ListCallDetailRecordsV2Request {
	s.EarlyMediaStateList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetEndTime(v int64) *ListCallDetailRecordsV2Request {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetFirstAgentId(v string) *ListCallDetailRecordsV2Request {
	s.FirstAgentId = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetInstanceId(v string) *ListCallDetailRecordsV2Request {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetMediaType(v string) *ListCallDetailRecordsV2Request {
	s.MediaType = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetNumber(v string) *ListCallDetailRecordsV2Request {
	s.Number = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetOrderByField(v string) *ListCallDetailRecordsV2Request {
	s.OrderByField = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetPageNumber(v int32) *ListCallDetailRecordsV2Request {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetPageSize(v int32) *ListCallDetailRecordsV2Request {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetReleaseInitiatorList(v string) *ListCallDetailRecordsV2Request {
	s.ReleaseInitiatorList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetReleaseReasonList(v string) *ListCallDetailRecordsV2Request {
	s.ReleaseReasonList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSatisfactionDescriptionList(v string) *ListCallDetailRecordsV2Request {
	s.SatisfactionDescriptionList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSatisfactionRateList(v string) *ListCallDetailRecordsV2Request {
	s.SatisfactionRateList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsV2Request {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSearchPattern(v string) *ListCallDetailRecordsV2Request {
	s.SearchPattern = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSkillGroupIdList(v string) *ListCallDetailRecordsV2Request {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetSortOrder(v string) *ListCallDetailRecordsV2Request {
	s.SortOrder = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) SetStartTime(v int64) *ListCallDetailRecordsV2Request {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsV2Request) Validate() error {
	return dara.Validate(s)
}
