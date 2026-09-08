// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalSkillGroupReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHistoricalSkillGroupReportResponseBody
	GetCode() *string
	SetData(v *ListHistoricalSkillGroupReportResponseBodyData) *ListHistoricalSkillGroupReportResponseBody
	GetData() *ListHistoricalSkillGroupReportResponseBodyData
	SetHttpStatusCode(v int32) *ListHistoricalSkillGroupReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHistoricalSkillGroupReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHistoricalSkillGroupReportResponseBody
	GetRequestId() *string
}

type ListHistoricalSkillGroupReportResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ListHistoricalSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 26A34338-5CD9-4C95-A7A6-5BDCE76C6B94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHistoricalSkillGroupReportResponseBody) GetData() *ListHistoricalSkillGroupReportResponseBodyData {
	return s.Data
}

func (s *ListHistoricalSkillGroupReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHistoricalSkillGroupReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoricalSkillGroupReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetCode(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetData(v *ListHistoricalSkillGroupReportResponseBodyData) *ListHistoricalSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetMessage(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetRequestId(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHistoricalSkillGroupReportResponseBodyData struct {
	// List of historical data for skill groups.
	List []*ListHistoricalSkillGroupReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) GetList() []*ListHistoricalSkillGroupReportResponseBodyDataList {
	return s.List
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetList(v []*ListHistoricalSkillGroupReportResponseBodyDataList) *ListHistoricalSkillGroupReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetPageSize(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) Validate() error {
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

type ListHistoricalSkillGroupReportResponseBodyDataList struct {
	// Back-to-back metric.
	Back2Back *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	// Inbound metrics.
	Inbound *ListHistoricalSkillGroupReportResponseBodyDataListInbound `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	// Outbound metrics.
	Outbound *ListHistoricalSkillGroupReportResponseBodyDataListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	// Overall metrics.
	Overall *ListHistoricalSkillGroupReportResponseBodyDataListOverall `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// Skill group ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Skill group name.
	//
	// example:
	//
	// skillgroup
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetBack2Back() *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	return s.Back2Back
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetInbound() *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	return s.Inbound
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetOutbound() *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	return s.Outbound
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetOverall() *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	return s.Overall
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetBack2Back(v *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Back2Back = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetInbound(v *ListHistoricalSkillGroupReportResponseBodyDataListInbound) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetOutbound(v *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetOverall(v *ListHistoricalSkillGroupReportResponseBodyDataListOverall) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Overall = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetSkillGroupId(v string) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetSkillGroupName(v string) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) Validate() error {
	if s.Back2Back != nil {
		if err := s.Back2Back.Validate(); err != nil {
			return err
		}
	}
	if s.Inbound != nil {
		if err := s.Inbound.Validate(); err != nil {
			return err
		}
	}
	if s.Outbound != nil {
		if err := s.Outbound.Validate(); err != nil {
			return err
		}
	}
	if s.Overall != nil {
		if err := s.Overall.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHistoricalSkillGroupReportResponseBodyDataListBack2Back struct {
	// Agent acknowledgement rate.
	//
	// example:
	//
	// 1
	AgentHandleRate *float32 `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	// Answer rate. Calculation Formula: CallsAnswered / CallsDialed. (Because acknowledgement events and answer events may fall into different time ranges, the result may exceed 100% in certain cases.)
	//
	// example:
	//
	// 0.6
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// Average customer-side ring time, in seconds.
	//
	// example:
	//
	// 100
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	// Average ring time, in seconds.
	//
	// example:
	//
	// 100
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// Average talk time, in seconds.
	//
	// example:
	//
	// 100
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Number of answered calls.
	//
	// example:
	//
	// 100
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// Number of calls answered by the customer.
	//
	// example:
	//
	// 8
	CallsCustomerAnswered *int64 `json:"CallsCustomerAnswered,omitempty" xml:"CallsCustomerAnswered,omitempty"`
	// Number of dial-up calls.
	//
	// example:
	//
	// 100
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// Customer answer rate.
	//
	// example:
	//
	// 0.8
	CustomerAnswerRate *float32 `json:"CustomerAnswerRate,omitempty" xml:"CustomerAnswerRate,omitempty"`
	// Maximum Customer-side ring time, in seconds.
	//
	// example:
	//
	// 100
	MaxCustomerRingTime *int64 `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	// Maximum ring time, in seconds.
	//
	// example:
	//
	// 100
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// Maximum talk time, in seconds.
	//
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// Total Customer-side ring time, in seconds.
	//
	// example:
	//
	// 100
	TotalCustomerRingTime *int64 `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	// Total ring time, in seconds.
	//
	// example:
	//
	// 100
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetAgentHandleRate() *float32 {
	return s.AgentHandleRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetAverageCustomerRingTime() *float32 {
	return s.AverageCustomerRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetCallsCustomerAnswered() *int64 {
	return s.CallsCustomerAnswered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetCustomerAnswerRate() *float32 {
	return s.CustomerAnswerRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetMaxCustomerRingTime() *int64 {
	return s.MaxCustomerRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetTotalCustomerRingTime() *int64 {
	return s.TotalCustomerRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetAgentHandleRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetAnswerRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetAverageCustomerRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetAverageRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetCallsAnswered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetCallsCustomerAnswered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetCallsDialed(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetCustomerAnswerRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetMaxCustomerRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetMaxRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetTotalCustomerRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetTotalRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalSkillGroupReportResponseBodyDataListInbound struct {
	// Abandon rate. Calculation Formula: CallsAbandoned / CallsOffered (Because abandonment events and assignment events may fall into different time ranges, the result may exceed 100% in certain cases).
	//
	// example:
	//
	// 0
	AbandonRate *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	// Statistics for each channel.
	AccessChannelTypeDetails []*ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails `json:"AccessChannelTypeDetails,omitempty" xml:"AccessChannelTypeDetails,omitempty" type:"Repeated"`
	// Average abandonment duration, in seconds. Calculation Formula: TotalAbandonTime / CallsAbandoned.
	//
	// example:
	//
	// 0
	AverageAbandonTime *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	// Average queue abandonment duration, in seconds. Calculation Formula: TotalAbandonedInQueueTime / CallsAbandonedInQueue.
	//
	// example:
	//
	// 0
	AverageAbandonedInQueueTime *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	// Average ringing abandonment duration, in seconds. Calculation Formula: TotalAbandonedInRingTime / CallsAbandonedInRing.
	//
	// example:
	//
	// 0
	AverageAbandonedInRingTime *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	// Average first response time for chat sessions, in seconds.
	//
	// example:
	//
	// 6
	AverageFirstResponseTime *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	// Average call hold duration, in seconds. Calculation Formula: TotalHoldTime / CallsHold.
	//
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// Average response time for chat sessions.
	//
	// example:
	//
	// 8
	AverageResponseTime *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
	// Average ring time, in seconds. Calculation Formula: TotalRingTime / CallsRinged.
	//
	// example:
	//
	// 5
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// Average talk time, in seconds. Calculation Formula: TotalTalkTime / CallsHandled.
	//
	// example:
	//
	// 64
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Average wait time, which is the average duration a caller waits before an agent answers the call. Calculation Formula: TotalWaitTime / CallsHandled.
	//
	// example:
	//
	// 5
	AverageWaitTime *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	// Average post-processing duration, in seconds. Calculation Formula: TotalWorkTime / CallsHandled.
	//
	// example:
	//
	// 13
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// Quantity of abandoned calls. Calculation Formula: CallsAbandonedInQueue + CallsAbandonedInRing.
	//
	// example:
	//
	// 0
	CallsAbandoned *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	// Number of calls abandoned in queue, which refers to the number of calls where the customer hung up after entering the queue but before being answered.
	//
	// example:
	//
	// 0
	CallsAbandonedInQueue *int64 `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	// Ring abandonment count, which is the number of calls where the customer hung up while the agent\\"s phone was ringing.
	//
	// example:
	//
	// 0
	CallsAbandonedInRing *int64 `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	// Transfer-in volume, which refers to the number of calls transferred to this skill group from other skill groups. Transfers between agents within the same skill group are not counted. If an agent is signed into multiple skill groups simultaneously, the call is attributed to the first skill group the agent signed into. If a single call is transferred multiple times from other skill groups to this skill group, each transfer is counted separately. The same rule applies below.
	//
	// example:
	//
	// 0
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// Quantity of attended transfer-out calls, which refers to the number of calls transferred from this skill group to another skill group via consultation. Transfers between agents within the same skill group are not counted.
	//
	// example:
	//
	// 0
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// Number of blind transfer-in calls, which refers to the number of calls directly transferred to this skill group from other skill groups. Transfers between agents within the same skill group are not counted. If an agent is signed into multiple skill groups simultaneously, the call is attributed to the first skill group the agent signed into. If a single call is transferred multiple times from other skill groups to this skill group, each transfer is counted separately. The same rule applies below.
	//
	// example:
	//
	// 0
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// Number of blind transfer-out calls, which refers to the number of calls directly transferred from this skill group to another skill group. Transfers between agents within the same skill group are not counted.
	//
	// example:
	//
	// 0
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// Acknowledgement count, which is the number of times agents answered calls. For a single call that enters a queue multiple times, if it is answered by multiple agents after one queue entry, it is counted as one.
	//
	// example:
	//
	// 7
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// Hold count, which is the number of times calls were placed on hold. Each time a call enters the queue and experiences multiple holds, it counts as one.
	//
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// Assigned call volume, which is the number of calls assigned to this skill group, including calls assigned through queues and calls assigned via transfers (consultation transfers and direct transfers). Calculation Formula: CallsQueued + CallsBlindTransferIn + CallsAttendedTransferIn.
	//
	// example:
	//
	// 7
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// Overflow count, which is the number of calls that experienced queue (skill group) overflow. If a single call enters the same queue multiple times, each overflow is counted separately.
	//
	// example:
	//
	// 0
	CallsOverflow *int64 `json:"CallsOverflow,omitempty" xml:"CallsOverflow,omitempty"`
	// Number of inbound calls entering a queue (skill group). If a single call enters the same queue multiple times, each entry is counted separately.
	//
	// example:
	//
	// 7
	CallsQueued *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	// Queue Failure Quantity, which is the number of calls where the customer hung up after entering the queue but before being answered.
	//
	// example:
	//
	// 0
	CallsQueuingFailed *int64 `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	// Quantity of calls that overflowed from the queue. Queue overflow refers to calls that overflow while queuing in IVR.
	//
	// example:
	//
	// 0
	CallsQueuingOverflow *int64 `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	// Number of calls that timed out during the queuing phase.
	//
	// example:
	//
	// 0
	CallsQueuingTimeout *int64 `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	// Number of calls that rang to agents. Each time a call enters the queue and is assigned to multiple agents, resulting in ringing, it counts as one.
	//
	// example:
	//
	// 7
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// Timeout count, which is the number of calls that experienced queue (skill group) timeout. If a single call enters the same queue multiple times, each timeout is counted separately.
	//
	// example:
	//
	// 0
	CallsTimeout *int64 `json:"CallsTimeout,omitempty" xml:"CallsTimeout,omitempty"`
	// Acknowledgement rate. Calculation Formula: CallsHandled / CallsOffered (because acknowledgement events and assign events may fall into different time ranges, the result may exceed 100% in certain cases).
	//
	// example:
	//
	// 1
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// Maximum abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	MaxAbandonTime *int64 `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	// Maximum queue abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	MaxAbandonedInQueueTime *int64 `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	// Maximum ring abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	MaxAbandonedInRingTime *int64 `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	// Maximum call hold time, in seconds.
	//
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// Maximum ring duration, in seconds.
	//
	// example:
	//
	// 12
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// Maximum talk duration, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// Maximum wait time, in seconds.
	//
	// example:
	//
	// 13
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// Maximum post-processing duration, in seconds.
	//
	// example:
	//
	// 12
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// Satisfaction index, which is the average of the satisfaction keypress digits (single-digit numbers).
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// Satisfaction rate. Calculation Formula: Count of evaluations marked as satisfied / Count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// Sending Count of satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// Count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// Service level within 15 seconds.
	//
	// example:
	//
	// 0.7
	ServiceLevel15 *float32 `json:"ServiceLevel15,omitempty" xml:"ServiceLevel15,omitempty"`
	// Service level within 20 seconds: number of calls with wait time less than or equal to 20 seconds divided by CallsQueued.
	//
	// example:
	//
	// 0
	ServiceLevel20 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	// Service level within 30 seconds.
	//
	// example:
	//
	// 0.9
	ServiceLevel30 *float32 `json:"ServiceLevel30,omitempty" xml:"ServiceLevel30,omitempty"`
	// Total abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonTime *int64 `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	// Total queue abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonedInQueueTime *int64 `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
	// Total ring abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonedInRingTime *int64 `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	// Total call hold duration, in seconds.
	//
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// Total number of messages sent in chat sessions.
	//
	// example:
	//
	// 12
	TotalMessagesSent *int64 `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	// Total number of messages sent by agents in chat sessions.
	//
	// example:
	//
	// 9
	TotalMessagesSentByAgent *int64 `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	// Total number of messages sent by the customer in chat sessions.
	//
	// example:
	//
	// 3
	TotalMessagesSentByCustomer *int64 `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	// Total ringing duration, in seconds.
	//
	// example:
	//
	// 32
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 447
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// Total waiting duration, in seconds.
	//
	// example:
	//
	// 34
	TotalWaitTime *int64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// Total post-processing time, in seconds.
	//
	// example:
	//
	// 85
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAbandonRate() *float32 {
	return s.AbandonRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAccessChannelTypeDetails() []*ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails {
	return s.AccessChannelTypeDetails
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageAbandonTime() *float32 {
	return s.AverageAbandonTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageAbandonedInQueueTime() *float32 {
	return s.AverageAbandonedInQueueTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageAbandonedInRingTime() *float32 {
	return s.AverageAbandonedInRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageWaitTime() *float32 {
	return s.AverageWaitTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsAbandonedInQueue() *int64 {
	return s.CallsAbandonedInQueue
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsAbandonedInRing() *int64 {
	return s.CallsAbandonedInRing
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsOverflow() *int64 {
	return s.CallsOverflow
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsQueued() *int64 {
	return s.CallsQueued
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsQueuingFailed() *int64 {
	return s.CallsQueuingFailed
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsQueuingOverflow() *int64 {
	return s.CallsQueuingOverflow
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetCallsTimeout() *int64 {
	return s.CallsTimeout
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxAbandonTime() *int64 {
	return s.MaxAbandonTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxAbandonedInQueueTime() *int64 {
	return s.MaxAbandonedInQueueTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxAbandonedInRingTime() *int64 {
	return s.MaxAbandonedInRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetServiceLevel15() *float32 {
	return s.ServiceLevel15
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetServiceLevel30() *float32 {
	return s.ServiceLevel30
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalAbandonTime() *int64 {
	return s.TotalAbandonTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalAbandonedInQueueTime() *int64 {
	return s.TotalAbandonedInQueueTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalAbandonedInRingTime() *int64 {
	return s.TotalAbandonedInRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalWaitTime() *int64 {
	return s.TotalWaitTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAbandonRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AbandonRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAccessChannelTypeDetails(v []*ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AccessChannelTypeDetails = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonedInQueueTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonedInRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageFirstResponseTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageResponseTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageWaitTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandoned(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandonedInQueue(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandonedInRing(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsOverflow(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsOverflow = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsQueued(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsQueuingFailed(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsQueuingOverflow(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsQueuingTimeout(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsTimeout(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsTimeout = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonedInQueueTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonedInRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxWaitTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetServiceLevel15(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.ServiceLevel15 = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetServiceLevel20(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetServiceLevel30(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.ServiceLevel30 = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonedInQueueTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonedInRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSent(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByAgent(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByCustomer(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalWaitTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) Validate() error {
	if s.AccessChannelTypeDetails != nil {
		for _, item := range s.AccessChannelTypeDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails struct {
	// Channel Type.
	//
	// example:
	//
	// Web
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// Quantity of assigned sessions.
	//
	// example:
	//
	// 2
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) SetAccessChannelType(v string) *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails {
	s.AccessChannelType = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) SetCallsOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInboundAccessChannelTypeDetails) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalSkillGroupReportResponseBodyDataListOutbound struct {
	// Answer rate. Calculation Formula: CallsAnswered / CallsDialed. (Because the call answering event and the acknowledgement event may fall into different time ranges, the result may exceed 100% in certain cases.)
	//
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// Average dial-up duration, in seconds. Calculation Formula: TotalDialingTime / CallsDialed.
	//
	// example:
	//
	// 37
	AverageDialingTime *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// Average call hold duration, in seconds. Calculation Formula: TotalHoldTime / CallsHold.
	//
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// Average ring time, in seconds. Calculation Formula: TotalRingTime / CallsRinged.
	//
	// example:
	//
	// 0
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// Average talk time, in seconds. Calculation Formula: TotalTalkTime / CallsAnswered.
	//
	// example:
	//
	// 3
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Average post-processing duration per call, in seconds. Calculation Formula: TotalWorkTime / CallsDialed
	//
	// example:
	//
	// 2
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// Number of answered calls.
	//
	// example:
	//
	// 1
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// Transfer-in volume for consultation, which refers to the number of calls transferred to this skill group from other skill groups for consultation. Transfers between agents within the same skill group are not counted. If an agent joins multiple skill groups simultaneously, the call is attributed to the first skill group the agent signed into. If a single call is transferred multiple times from other skill groups to this skill group, each transfer is counted separately. The same rule applies below.
	//
	// example:
	//
	// 0
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// Quantity of attended transfer-out calls, which refers to the number of calls transferred from this skill group to another skill group for consultation. Transfers between agents within the same skill group are not counted.
	//
	// example:
	//
	// 0
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// Quantity of direct transfer-in calls, which refers to the number of calls directly transferred to this skill group from other skill groups. Transfers between agents within the same skill group are not counted. If an agent is signed into multiple skill groups simultaneously, the call is attributed to the first skill group the agent signed into. If a single call is transferred multiple times from other skill groups to this skill group, each transfer is counted separately. The same rule applies below.
	//
	// example:
	//
	// 0
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// Quantity of direct transfer-out calls, which refers to the number of calls directly transferred from this skill group to other skill groups. Transfers between agents within the same skill group are not counted.
	//
	// example:
	//
	// 0
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// Number of dialed calls.
	//
	// example:
	//
	// 6
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// Number of calls placed on hold. If a call is placed on hold multiple times before being transfer-out from the current skill group, it counts as one occurrence.
	//
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// Number of calls that rang to agents. Each time a call enters the queue and is assigned to multiple agents, resulting in ringing, it counts as one occurrence.
	//
	// example:
	//
	// 0
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// Maximum dialing time, in seconds.
	//
	// example:
	//
	// 12
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// Maximum hold time during calls, in seconds.
	//
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// Maximum ring duration, in seconds.
	//
	// example:
	//
	// 0
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// Maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// Maximum post-processing duration per call, in seconds.
	//
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// Satisfaction index, which is the average value of the single-digit satisfaction key presses.
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// Satisfaction rate. Calculation Formula: Quantity of evaluations marked as satisfied divided by the Count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// Sending Count of satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// Response Count of satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// Total dial-up duration, in seconds.
	//
	// example:
	//
	// 218
	TotalDialingTime *int64 `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	// Total call hold duration, in seconds.
	//
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// Total ring duration, in seconds.
	//
	// example:
	//
	// 0
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 3
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// Total post-processing duration, in seconds.
	//
	// example:
	//
	// 9
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalSkillGroupReportResponseBodyDataListOverall struct {
	// Average break duration, in seconds. Calculation Formula: TotalBreakTime / Break Count. Break Count is a non-API statistical field.
	//
	// example:
	//
	// 0
	AverageBreakTime *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	// Average call hold duration, in seconds. Calculation Formula: TotalHoldTime / (Inbound CallsHold + Outbound CallsHold).
	//
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// Average ready time, in seconds. Calculation Formula: TotalReadyTime / Count of ready events. The count of ready events is not an API statistics field.
	//
	// example:
	//
	// 0
	AverageReadyTime *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// Average talk time, in seconds. Calculation formula: TotalTalkTime / (CallsAnswered + CallsHandled).
	//
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Average post-processing time, in seconds. Calculation Formula: TotalWorkTime / TotalCalls.
	//
	// example:
	//
	// 8
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// List of break details.
	BreakCodeDetailList []*ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// Maximum break duration, in seconds.
	//
	// example:
	//
	// 1
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// Maximum call hold duration, in seconds.
	//
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// Maximum ready time, in seconds.
	//
	// example:
	//
	// 19328
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// Maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// Maximum post-processing duration, in seconds.
	//
	// example:
	//
	// 12
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// Agent occupancy rate. Calculation formula: (TotalWorkTime + TotalTalkTime) / TotalLoggedInTime.
	//
	// example:
	//
	// 0.02332222293912065
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// Satisfaction index, which is the average value of the satisfaction keypress digits (single-digit numbers).
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// Satisfaction rate. Calculation Formula: Number of responses marked as satisfied / Count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// Sending Count of satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// Count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// Total break time, in seconds.
	//
	// example:
	//
	// 3
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// Total call volume. Calculation Formula: CallsOffered + CallsDialed.
	//
	// example:
	//
	// 13
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// Total hold duration, in seconds.
	//
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// Total logon time, in seconds.
	//
	// _Note: Excludes offline and short break durations._
	//
	// example:
	//
	// 23218
	TotalLoggedInTime *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	// Total ready time, in seconds.
	//
	// example:
	//
	// 22428
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 449
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// Total post-processing duration, in seconds.
	//
	// example:
	//
	// 94
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverall) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetBreakCodeDetailList() []*ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetBreakCodeDetailList(v []*ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) Validate() error {
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

type ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList struct {
	// Break type code.
	//
	// example:
	//
	// 会议
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// Number of occurrences of this break type.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Total duration of this break type, in seconds.
	//
	// example:
	//
	// 3600
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetBreakCode(v string) *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetCount(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetDuration(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
