// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalInstanceReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntervalInstanceReportResponseBody
	GetCode() *string
	SetData(v []*ListIntervalInstanceReportResponseBodyData) *ListIntervalInstanceReportResponseBody
	GetData() []*ListIntervalInstanceReportResponseBodyData
	SetHttpStatusCode(v int32) *ListIntervalInstanceReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntervalInstanceReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntervalInstanceReportResponseBody
	GetRequestId() *string
}

type ListIntervalInstanceReportResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of instance segment statistics data.
	Data []*ListIntervalInstanceReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntervalInstanceReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntervalInstanceReportResponseBody) GetData() []*ListIntervalInstanceReportResponseBodyData {
	return s.Data
}

func (s *ListIntervalInstanceReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntervalInstanceReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntervalInstanceReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntervalInstanceReportResponseBody) SetCode(v string) *ListIntervalInstanceReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetData(v []*ListIntervalInstanceReportResponseBodyData) *ListIntervalInstanceReportResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalInstanceReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetMessage(v string) *ListIntervalInstanceReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetRequestId(v string) *ListIntervalInstanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntervalInstanceReportResponseBodyData struct {
	// Inbound metrics.
	Inbound *ListIntervalInstanceReportResponseBodyDataInbound `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	// Outbound metrics.
	Outbound *ListIntervalInstanceReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	// Overall metrics.
	Overall *ListIntervalInstanceReportResponseBodyDataOverall `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// Start Time, formatted as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1620230400000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyData) GetInbound() *ListIntervalInstanceReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *ListIntervalInstanceReportResponseBodyData) GetOutbound() *ListIntervalInstanceReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *ListIntervalInstanceReportResponseBodyData) GetOverall() *ListIntervalInstanceReportResponseBodyDataOverall {
	return s.Overall
}

func (s *ListIntervalInstanceReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListIntervalInstanceReportResponseBodyData) SetInbound(v *ListIntervalInstanceReportResponseBodyDataInbound) *ListIntervalInstanceReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetOutbound(v *ListIntervalInstanceReportResponseBodyDataOutbound) *ListIntervalInstanceReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetOverall(v *ListIntervalInstanceReportResponseBodyDataOverall) *ListIntervalInstanceReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetStatsTime(v int64) *ListIntervalInstanceReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) Validate() error {
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

type ListIntervalInstanceReportResponseBodyDataInbound struct {
	// Abandon rate. Calculation Formula: CallsAbandoned / CallsOffered (because management events related to abandonment and assignment may fall into different Time Ranges, the Result may exceed 100% in certain cases).
	//
	// example:
	//
	// 0
	AbandonRate *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	// Deprecated. Refer to the AbandonRate field instead.
	//
	// example:
	//
	// 0
	AbandonedRate *float32 `json:"AbandonedRate,omitempty" xml:"AbandonedRate,omitempty"`
	// The average abandon time, in seconds. Calculation Formula: TotalAbandonTime / CallsAbandoned.
	//
	// example:
	//
	// 0
	AverageAbandonTime *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	// Average IVR abandonment time, in seconds. Calculation Formula: TotalAbandonedInIVRTime / CallsAbandonedInIVR.
	//
	// example:
	//
	// 0
	AverageAbandonedInIVRTime *float32 `json:"AverageAbandonedInIVRTime,omitempty" xml:"AverageAbandonedInIVRTime,omitempty"`
	// Average queue abandonment duration, in seconds. Calculation Formula: TotalAbandonedInQueueTime / CallsAbandonedInQueue.
	//
	// example:
	//
	// 0
	AverageAbandonedInQueueTime *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	// Average abandoned-in-ring time, in seconds. Calculation Formula: TotalAbandonedInRingTime / CallsAbandonedInRing.
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
	// Average response time (RT) of chat sessions.
	//
	// example:
	//
	// 15
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
	// Average wait time, which is the average duration callers wait before an agent answers a call. Calculation Formula: TotalWaitTime / CallsHandled.
	//
	// example:
	//
	// 5
	AverageWaitTime *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	// Average post-processing time, in seconds. Calculation Formula: TotalWorkTime / CallsHandled.
	//
	// example:
	//
	// 13
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// Total abandoned calls. Calculation Formula: CallsAbandonedInIVR + CallsAbandonedInQueue + CallsAbandonedInRing.
	//
	// example:
	//
	// 0
	CallsAbandoned *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	// Number of calls abandoned in IVR, which refers to the count of calls where the Customer hung up during the IVR flow after entering the IVR process.
	//
	// example:
	//
	// 0
	CallsAbandonedInIVR *int64 `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	// The number of calls abandoned in the queue, which refers to calls where the customer hung up while waiting in the queue after entering it.
	//
	// example:
	//
	// 0
	CallsAbandonedInQueue *int64 `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	// Number of calls abandoned during ringing, which refers to the Quantity of calls where the Customer hung up while the agent\\"s phone was ringing.
	//
	// example:
	//
	// 0
	CallsAbandonedInRing *int64 `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	// The number of calls abandoned in the Intelligent Navigation module.
	//
	// example:
	//
	// 0
	CallsAbandonedInVoiceNavigator *int64 `json:"CallsAbandonedInVoiceNavigator,omitempty" xml:"CallsAbandonedInVoiceNavigator,omitempty"`
	// Number of consult transfers, which refers to the Count of calls that involved a consult transfer. If a single call initiated multiple transfers, it is counted as one.
	//
	// example:
	//
	// 0
	CallsAttendedTransferred *int64 `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	// Quantity of blind transfers, which is the number of calls that were directly transferred. If a single call was transferred multiple times, it is counted as one.
	//
	// example:
	//
	// 0
	CallsBlindTransferred *int64 `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	// Number of calls that caused IVR exceptions.
	//
	// example:
	//
	// 0
	CallsCausedIVRException *int64 `json:"CallsCausedIVRException,omitempty" xml:"CallsCausedIVRException,omitempty"`
	// The number of calls forwarded to an external number.
	//
	// example:
	//
	// 0
	CallsForwardToOutsideNumber *int64 `json:"CallsForwardToOutsideNumber,omitempty" xml:"CallsForwardToOutsideNumber,omitempty"`
	// The acknowledgement count, which refers to the number of calls answered by agents. If a single call is answered by multiple agents, it is counted as one.
	//
	// example:
	//
	// 7
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// Number of calls placed on hold. If a single call is placed on hold multiple times, it is counted as one.
	//
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// The number of calls with IVR exceptions. A call is counted when the IVR enters a hang-up reason node and the hang-up reason configured for that node is "failed transfer to agent." In this case, the count increases by 1.
	//
	// example:
	//
	// 0
	CallsIVRException *int64 `json:"CallsIVRException,omitempty" xml:"CallsIVRException,omitempty"`
	// Number of calls offered to Cloud Contact Center.
	//
	// example:
	//
	// 7
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// Number of calls that entered the queue. If a single call entered the queue multiple times, it is counted as one.
	//
	// example:
	//
	// 7
	CallsQueued *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	// Queue Failure quantity, which refers to the number of calls where the customer hung up during queuing after the call entered the queue.
	//
	// example:
	//
	// 0
	CallsQueuingFailed *int64 `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	// Quantity of calls that overflowed from the queue, meaning calls that encountered queue overflow while waiting in the IVR queue.
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
	// Number of calls that rang to agents. If a single call was assigned to multiple agents and rang for each, it is counted as one.
	//
	// example:
	//
	// 7
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// Quantity of calls transferred to voicemail.
	//
	// example:
	//
	// 1
	CallsToVoicemail *int64 `json:"CallsToVoicemail,omitempty" xml:"CallsToVoicemail,omitempty"`
	// The number of calls directed to voicemail.
	//
	// example:
	//
	// 0
	CallsVoicemail *int64 `json:"CallsVoicemail,omitempty" xml:"CallsVoicemail,omitempty"`
	// Acknowledgement rate. Calculation Formula: CallsHandled / CallsOffered (because acknowledgement events and assign events may fall into different time ranges, the result may exceed 100% in certain cases).
	//
	// example:
	//
	// 1
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// The maximum abandon time, in seconds.
	//
	// example:
	//
	// 0
	MaxAbandonTime *int64 `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	// Maximum IVR abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	MaxAbandonedInIVRTime *int64 `json:"MaxAbandonedInIVRTime,omitempty" xml:"MaxAbandonedInIVRTime,omitempty"`
	// Maximum abandoned-in-queue time, in seconds.
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
	// Maximum hold time, in seconds.
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
	// Maximum talk time, in seconds.
	//
	// example:
	//
	// 219
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// The maximum wait time, in seconds.
	//
	// example:
	//
	// 13
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// Maximum post-processing time, in seconds.
	//
	// example:
	//
	// 17
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// Satisfaction index, which is the average value of the satisfaction keypress digits (single-digit numbers).
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// Satisfaction rate. Calculation Formula: Number of evaluations marked as satisfied / Count of satisfaction survey responses.
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
	// The count of satisfaction survey responses.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// Service level within 20 seconds, calculated as the number of calls with wait time less than or equal to 20 seconds divided by CallsQueued.
	//
	// example:
	//
	// 1
	ServiceLevel20 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	// Total abandon time, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonTime *int64 `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	// Total IVR abandonment duration, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonedInIVRTime *int64 `json:"TotalAbandonedInIVRTime,omitempty" xml:"TotalAbandonedInIVRTime,omitempty"`
	// Total abandoned-in-queue time, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonedInQueueTime *int64 `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
	// Total abandoned-in-ring time, in seconds.
	//
	// example:
	//
	// 0
	TotalAbandonedInRingTime *int64 `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	// Total hold time during calls, in seconds.
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
	// 8
	TotalMessagesSentByAgent *int64 `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	// Total number of messages sent by the Customer in chat sessions.
	//
	// example:
	//
	// 4
	TotalMessagesSentByCustomer *int64 `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	// Total ring time, in seconds.
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
	// The total wait time, in seconds.
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

func (s ListIntervalInstanceReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAbandonRate() *float32 {
	return s.AbandonRate
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAbandonedRate() *float32 {
	return s.AbandonedRate
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageAbandonTime() *float32 {
	return s.AverageAbandonTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInIVRTime() *float32 {
	return s.AverageAbandonedInIVRTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInQueueTime() *float32 {
	return s.AverageAbandonedInQueueTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInRingTime() *float32 {
	return s.AverageAbandonedInRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageWaitTime() *float32 {
	return s.AverageWaitTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInIVR() *int64 {
	return s.CallsAbandonedInIVR
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInQueue() *int64 {
	return s.CallsAbandonedInQueue
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInRing() *int64 {
	return s.CallsAbandonedInRing
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInVoiceNavigator() *int64 {
	return s.CallsAbandonedInVoiceNavigator
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsAttendedTransferred() *int64 {
	return s.CallsAttendedTransferred
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsBlindTransferred() *int64 {
	return s.CallsBlindTransferred
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsCausedIVRException() *int64 {
	return s.CallsCausedIVRException
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsForwardToOutsideNumber() *int64 {
	return s.CallsForwardToOutsideNumber
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsIVRException() *int64 {
	return s.CallsIVRException
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsQueued() *int64 {
	return s.CallsQueued
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsQueuingFailed() *int64 {
	return s.CallsQueuingFailed
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsQueuingOverflow() *int64 {
	return s.CallsQueuingOverflow
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsToVoicemail() *int64 {
	return s.CallsToVoicemail
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetCallsVoicemail() *int64 {
	return s.CallsVoicemail
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxAbandonTime() *int64 {
	return s.MaxAbandonTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInIVRTime() *int64 {
	return s.MaxAbandonedInIVRTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInQueueTime() *int64 {
	return s.MaxAbandonedInQueueTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInRingTime() *int64 {
	return s.MaxAbandonedInRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalAbandonTime() *int64 {
	return s.TotalAbandonTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInIVRTime() *int64 {
	return s.TotalAbandonedInIVRTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInQueueTime() *int64 {
	return s.TotalAbandonedInQueueTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInRingTime() *int64 {
	return s.TotalAbandonedInRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalWaitTime() *int64 {
	return s.TotalWaitTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAbandonRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AbandonRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAbandonedRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AbandonedRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInIVRTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInVoiceNavigator(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInVoiceNavigator = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAttendedTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsBlindTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsCausedIVRException(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsCausedIVRException = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsForwardToOutsideNumber(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsForwardToOutsideNumber = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsIVRException(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsIVRException = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueued(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingFailed(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingOverflow(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingTimeout(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsToVoicemail(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsToVoicemail = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsVoicemail(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsVoicemail = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInIVRTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetServiceLevel20(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInIVRTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalInstanceReportResponseBodyDataOutbound struct {
	// Answer rate. Calculation Formula: CallsAnswered / CallsDialed (because management events for answering and acknowledgement may fall into different Time Ranges, the Result may exceed 100% in certain cases).
	//
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// Average dial-up duration, in seconds. Calculation Formula: TotalDialingTime / CallsDialed.
	//
	// example:
	//
	// 0
	AverageDialingTime *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// Average hold duration, in seconds. Calculation formula: TotalHoldTime / CallsHold.
	//
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// Average ring time in seconds. Calculation Formula: TotalRingTime / CallsRinged.
	//
	// example:
	//
	// 0
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// Average talk time, in seconds. Calculation Formula: TotalTalkTime / CallsAnswered.
	//
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Average post-processing time in seconds. Calculation Formula: TotalWorkTime / CallsDialed.
	//
	// example:
	//
	// 0
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// Number of answered calls.
	//
	// example:
	//
	// 0
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// The number of calls transferred via consultation, which refers to the quantity of calls that underwent consultation-based transfer. If a single call was transferred multiple times, it is counted as one.
	//
	// example:
	//
	// 0
	CallsAttendedTransferred *int64 `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	// Quantity of blind transfers, which refers to the number of calls that were directly transferred. If a single call is transferred multiple times, it counts as one.
	//
	// example:
	//
	// 0
	CallsBlindTransferred *int64 `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	// Number of dialed calls.
	//
	// example:
	//
	// 0
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// Number of calls placed on hold. If a single call is placed on hold multiple times, it is counted as one.
	//
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// Number of calls that rang for agents. If a single call is assigned to multiple agents and rings for each, it is counted as one.
	//
	// example:
	//
	// 0
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// Maximum dial-up duration, in seconds.
	//
	// example:
	//
	// 0
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// Maximum hold time during calls, in seconds.
	//
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// Maximum ring time, in seconds.
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
	// Maximum post-processing time in seconds.
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
	// Satisfaction rate. Calculation Formula: Quantity of evaluations marked as satisfied / Count of satisfaction survey responses.
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
	// Total dial-up duration, in seconds.
	//
	// example:
	//
	// 0
	TotalDialingTime *int64 `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	// Total hold duration, in seconds.
	//
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// Total ring time, in seconds.
	//
	// example:
	//
	// 0
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// Total post-processing duration, in seconds.
	//
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsAttendedTransferred() *int64 {
	return s.CallsAttendedTransferred
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsBlindTransferred() *int64 {
	return s.CallsBlindTransferred
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsAttendedTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsBlindTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalInstanceReportResponseBodyDataOverall struct {
	// Average break time, in seconds. Calculation Formula: TotalBreakTime / Break Count. Break Count is not an API statistics field.
	//
	// example:
	//
	// 0
	AverageBreakTime *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	// Average call hold time, in seconds. Calculation Formula: TotalHoldTime / (Inbound Calls on Hold + Outbound Calls on Hold).
	//
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// Average ready time, in seconds. Calculation Formula: TotalReadyTime / Ready Count. Ready Count is not an API statistics field.
	//
	// example:
	//
	// 0
	AverageReadyTime *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// Average talk time, in seconds. Calculation Formula: TotalTalkTime / (CallsAnswered + CallsHandled).
	//
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// Average post-processing time per call, in seconds. Calculation Formula: TotalWorkTime / TotalCalls.
	//
	// example:
	//
	// 0
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// Maximum break time, in seconds.
	//
	// example:
	//
	// 0
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// Maximum call hold time, in seconds.
	//
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// Maximum ready time, in seconds.
	//
	// example:
	//
	// 0
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// Maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// Maximum post-processing time, in seconds.
	//
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// Agent occupancy rate. Calculation Formula: (TotalWorkTime + TotalTalkTime) / TotalLoggedInTime.
	//
	// example:
	//
	// 0
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// Satisfaction index, which is the average value of the satisfaction keypress digits (single-digit numbers).
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
	// Response Count of satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// Total break time in seconds.
	//
	// example:
	//
	// 0
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// Total call volume. Calculation Formula: CallsOffered + CallsDialed.
	//
	// example:
	//
	// 0
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// Total hold time, in seconds.
	//
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// Total logon duration, in seconds. Exclude break time.
	//
	// example:
	//
	// 0
	TotalLoggedInTime *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	// Total ready time in seconds.
	//
	// example:
	//
	// 0
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// Total talk time, in seconds.
	//
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// Total post-processing time in seconds.
	//
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}
