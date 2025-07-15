// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalInstanceReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHistoricalInstanceReportResponseBody
	GetCode() *string
	SetData(v *GetHistoricalInstanceReportResponseBodyData) *GetHistoricalInstanceReportResponseBody
	GetData() *GetHistoricalInstanceReportResponseBodyData
	SetHttpStatusCode(v int32) *GetHistoricalInstanceReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHistoricalInstanceReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHistoricalInstanceReportResponseBody
	GetRequestId() *string
}

type GetHistoricalInstanceReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHistoricalInstanceReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 943D8EF3-3321-471F-A104-51C96FCA94D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHistoricalInstanceReportResponseBody) GetData() *GetHistoricalInstanceReportResponseBodyData {
	return s.Data
}

func (s *GetHistoricalInstanceReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHistoricalInstanceReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHistoricalInstanceReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHistoricalInstanceReportResponseBody) SetCode(v string) *GetHistoricalInstanceReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetData(v *GetHistoricalInstanceReportResponseBodyData) *GetHistoricalInstanceReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalInstanceReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetMessage(v string) *GetHistoricalInstanceReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetRequestId(v string) *GetHistoricalInstanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyData struct {
	Inbound  *GetHistoricalInstanceReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal *GetHistoricalInstanceReportResponseBodyDataInternal `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound *GetHistoricalInstanceReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall  *GetHistoricalInstanceReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s GetHistoricalInstanceReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyData) GetInbound() *GetHistoricalInstanceReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *GetHistoricalInstanceReportResponseBodyData) GetInternal() *GetHistoricalInstanceReportResponseBodyDataInternal {
	return s.Internal
}

func (s *GetHistoricalInstanceReportResponseBodyData) GetOutbound() *GetHistoricalInstanceReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *GetHistoricalInstanceReportResponseBodyData) GetOverall() *GetHistoricalInstanceReportResponseBodyDataOverall {
	return s.Overall
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetInbound(v *GetHistoricalInstanceReportResponseBodyDataInbound) *GetHistoricalInstanceReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetInternal(v *GetHistoricalInstanceReportResponseBodyDataInternal) *GetHistoricalInstanceReportResponseBodyData {
	s.Internal = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetOutbound(v *GetHistoricalInstanceReportResponseBodyDataOutbound) *GetHistoricalInstanceReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetOverall(v *GetHistoricalInstanceReportResponseBodyDataOverall) *GetHistoricalInstanceReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyDataInbound struct {
	// example:
	//
	// 0
	AbandonRate                 *float32                                                                         `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	AccessChannelTypeDetailList []*GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList `json:"AccessChannelTypeDetailList,omitempty" xml:"AccessChannelTypeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	AverageAbandonTime *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	// example:
	//
	// 0
	AverageAbandonedInIVRTime *float32 `json:"AverageAbandonedInIVRTime,omitempty" xml:"AverageAbandonedInIVRTime,omitempty"`
	// example:
	//
	// 0
	AverageAbandonedInQueueTime *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	// example:
	//
	// 0
	AverageAbandonedInRingTime *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	AverageFirstResponseTime   *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	// example:
	//
	// 0
	AverageHoldTime     *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageResponseTime *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
	// example:
	//
	// 0
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWaitTime *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	CallsAbandoned *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInIVR *int64 `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInQueue *int64 `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInRing *int64 `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	// example:
	//
	// 0
	CallsAbandonedInVoiceNavigator *int64 `json:"CallsAbandonedInVoiceNavigator,omitempty" xml:"CallsAbandonedInVoiceNavigator,omitempty"`
	// example:
	//
	// 0
	CallsAttendedTransferred *int64 `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferred   *int64 `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	CallsCausedIVRException *int64 `json:"CallsCausedIVRException,omitempty" xml:"CallsCausedIVRException,omitempty"`
	// example:
	//
	// 0
	CallsForwardToOutsideNumber *int64 `json:"CallsForwardToOutsideNumber,omitempty" xml:"CallsForwardToOutsideNumber,omitempty"`
	// example:
	//
	// 0
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 0
	CallsIVRException *int64 `json:"CallsIVRException,omitempty" xml:"CallsIVRException,omitempty"`
	// example:
	//
	// 0
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 0
	CallsQueued *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	// example:
	//
	// 0
	CallsQueuingFailed *int64 `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	// example:
	//
	// 0
	CallsQueuingOverflow *int64 `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	// example:
	//
	// 0
	CallsQueuingTimeout *int64 `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	// example:
	//
	// 0
	CallsRinged      *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsToVoicemail *int64 `json:"CallsToVoicemail,omitempty" xml:"CallsToVoicemail,omitempty"`
	// example:
	//
	// 0
	CallsVoicemail *int64 `json:"CallsVoicemail,omitempty" xml:"CallsVoicemail,omitempty"`
	// example:
	//
	// 0
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// example:
	//
	// 0
	MaxAbandonTime *int64 `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	// example:
	//
	// 100
	MaxAbandonedInIVRTime *int64 `json:"MaxAbandonedInIVRTime,omitempty" xml:"MaxAbandonedInIVRTime,omitempty"`
	// example:
	//
	// 0
	MaxAbandonedInQueueTime *int64 `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	// example:
	//
	// 0
	MaxAbandonedInRingTime *int64 `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 0
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 0
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	ServiceLevel15               *float32 `json:"ServiceLevel15,omitempty" xml:"ServiceLevel15,omitempty"`
	// example:
	//
	// 0
	ServiceLevel20 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	ServiceLevel30 *float32 `json:"ServiceLevel30,omitempty" xml:"ServiceLevel30,omitempty"`
	// example:
	//
	// 0
	TotalAbandonTime *int64 `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	// example:
	//
	// 0
	TotalAbandonedInIVRTime *int64 `json:"TotalAbandonedInIVRTime,omitempty" xml:"TotalAbandonedInIVRTime,omitempty"`
	// example:
	//
	// 0
	TotalAbandonedInQueueTime *int64 `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
	// example:
	//
	// 0
	TotalAbandonedInRingTime *int64 `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	// example:
	//
	// 0
	TotalHoldTime               *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalMessagesSent           *int64 `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	TotalMessagesSentByAgent    *int64 `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	TotalMessagesSentByCustomer *int64 `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	// example:
	//
	// 0
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 0
	TotalWaitTime *int64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAbandonRate() *float32 {
	return s.AbandonRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAccessChannelTypeDetailList() []*GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList {
	return s.AccessChannelTypeDetailList
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageAbandonTime() *float32 {
	return s.AverageAbandonTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInIVRTime() *float32 {
	return s.AverageAbandonedInIVRTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInQueueTime() *float32 {
	return s.AverageAbandonedInQueueTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageAbandonedInRingTime() *float32 {
	return s.AverageAbandonedInRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageWaitTime() *float32 {
	return s.AverageWaitTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInIVR() *int64 {
	return s.CallsAbandonedInIVR
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInQueue() *int64 {
	return s.CallsAbandonedInQueue
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInRing() *int64 {
	return s.CallsAbandonedInRing
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAbandonedInVoiceNavigator() *int64 {
	return s.CallsAbandonedInVoiceNavigator
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsAttendedTransferred() *int64 {
	return s.CallsAttendedTransferred
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsBlindTransferred() *int64 {
	return s.CallsBlindTransferred
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsCausedIVRException() *int64 {
	return s.CallsCausedIVRException
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsForwardToOutsideNumber() *int64 {
	return s.CallsForwardToOutsideNumber
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsIVRException() *int64 {
	return s.CallsIVRException
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsQueued() *int64 {
	return s.CallsQueued
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsQueuingFailed() *int64 {
	return s.CallsQueuingFailed
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsQueuingOverflow() *int64 {
	return s.CallsQueuingOverflow
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsToVoicemail() *int64 {
	return s.CallsToVoicemail
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetCallsVoicemail() *int64 {
	return s.CallsVoicemail
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxAbandonTime() *int64 {
	return s.MaxAbandonTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInIVRTime() *int64 {
	return s.MaxAbandonedInIVRTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInQueueTime() *int64 {
	return s.MaxAbandonedInQueueTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxAbandonedInRingTime() *int64 {
	return s.MaxAbandonedInRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetServiceLevel15() *float32 {
	return s.ServiceLevel15
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetServiceLevel30() *float32 {
	return s.ServiceLevel30
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalAbandonTime() *int64 {
	return s.TotalAbandonTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInIVRTime() *int64 {
	return s.TotalAbandonedInIVRTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInQueueTime() *int64 {
	return s.TotalAbandonedInQueueTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalAbandonedInRingTime() *int64 {
	return s.TotalAbandonedInRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalWaitTime() *int64 {
	return s.TotalWaitTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAbandonRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AbandonRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAccessChannelTypeDetailList(v []*GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AccessChannelTypeDetailList = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInIVRTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInVoiceNavigator(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInVoiceNavigator = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAttendedTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsBlindTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsCausedIVRException(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsCausedIVRException = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsForwardToOutsideNumber(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsForwardToOutsideNumber = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsHandled(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsHold(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsIVRException(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsIVRException = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueued(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingFailed(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingOverflow(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingTimeout(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsRinged(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsToVoicemail(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsToVoicemail = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsVoicemail(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsVoicemail = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetHandleRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInIVRTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetServiceLevel15(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel15 = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetServiceLevel20(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetServiceLevel30(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel30 = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInIVRTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList struct {
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	CallsOffered      *int64  `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) SetAccessChannelType(v string) *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList {
	s.AccessChannelType = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) SetCallsOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList {
	s.CallsOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInboundAccessChannelTypeDetailList) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyDataInternal struct {
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsDialed   *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataInternal) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataInternal) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataInternal) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *GetHistoricalInstanceReportResponseBodyDataInternal) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *GetHistoricalInstanceReportResponseBodyDataInternal) SetCallsAnswered(v int64) *GetHistoricalInstanceReportResponseBodyDataInternal {
	s.CallsAnswered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInternal) SetCallsDialed(v int64) *GetHistoricalInstanceReportResponseBodyDataInternal {
	s.CallsDialed = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInternal) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyDataOutbound struct {
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 0
	AverageDialingTime *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// example:
	//
	// 0
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 0
	CallsAttendedTransferred *int64 `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferred *int64 `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	// example:
	//
	// 0
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 0
	CallsHold *int32 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 0
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 0
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 0
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// example:
	//
	// 0
	TotalDialingTime *int64 `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// example:
	//
	// 0
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsAttendedTransferred() *int64 {
	return s.CallsAttendedTransferred
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsBlindTransferred() *int64 {
	return s.CallsBlindTransferred
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsHold() *int32 {
	return s.CallsHold
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAnswerRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsAttendedTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsBlindTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsDialed(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsHold(v int32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsRinged(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalInstanceReportResponseBodyDataOverall struct {
	// example:
	//
	// 0
	AverageBreakTime *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	// example:
	//
	// 0
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// example:
	//
	// 0
	AverageReadyTime *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// example:
	//
	// 0
	MaxHoldTime       *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxLoggedInAgents *int64 `json:"MaxLoggedInAgents,omitempty" xml:"MaxLoggedInAgents,omitempty"`
	// example:
	//
	// 0
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 0
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// example:
	//
	// 0
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// example:
	//
	// 0
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// example:
	//
	// 0
	TotalLoggedInTime *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	// example:
	//
	// 0
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxLoggedInAgents() *int64 {
	return s.MaxLoggedInAgents
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxLoggedInAgents(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxLoggedInAgents = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetOccupancyRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalCalls(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}
