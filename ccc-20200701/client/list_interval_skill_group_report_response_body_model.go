// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalSkillGroupReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntervalSkillGroupReportResponseBody
	GetCode() *string
	SetData(v []*ListIntervalSkillGroupReportResponseBodyData) *ListIntervalSkillGroupReportResponseBody
	GetData() []*ListIntervalSkillGroupReportResponseBodyData
	SetHttpStatusCode(v int32) *ListIntervalSkillGroupReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntervalSkillGroupReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntervalSkillGroupReportResponseBody
	GetRequestId() *string
}

type ListIntervalSkillGroupReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListIntervalSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListIntervalSkillGroupReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntervalSkillGroupReportResponseBody) GetData() []*ListIntervalSkillGroupReportResponseBodyData {
	return s.Data
}

func (s *ListIntervalSkillGroupReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntervalSkillGroupReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntervalSkillGroupReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntervalSkillGroupReportResponseBody) SetCode(v string) *ListIntervalSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetData(v []*ListIntervalSkillGroupReportResponseBodyData) *ListIntervalSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetMessage(v string) *ListIntervalSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetRequestId(v string) *ListIntervalSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyData struct {
	Back2Back *ListIntervalSkillGroupReportResponseBodyDataBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	Inbound   *ListIntervalSkillGroupReportResponseBodyDataInbound   `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound  *ListIntervalSkillGroupReportResponseBodyDataOutbound  `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalSkillGroupReportResponseBodyDataOverall   `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// 1604639129000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyData) GetBack2Back() *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	return s.Back2Back
}

func (s *ListIntervalSkillGroupReportResponseBodyData) GetInbound() *ListIntervalSkillGroupReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *ListIntervalSkillGroupReportResponseBodyData) GetOutbound() *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *ListIntervalSkillGroupReportResponseBodyData) GetOverall() *ListIntervalSkillGroupReportResponseBodyDataOverall {
	return s.Overall
}

func (s *ListIntervalSkillGroupReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetBack2Back(v *ListIntervalSkillGroupReportResponseBodyDataBack2Back) *ListIntervalSkillGroupReportResponseBodyData {
	s.Back2Back = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetInbound(v *ListIntervalSkillGroupReportResponseBodyDataInbound) *ListIntervalSkillGroupReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetOutbound(v *ListIntervalSkillGroupReportResponseBodyDataOutbound) *ListIntervalSkillGroupReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetOverall(v *ListIntervalSkillGroupReportResponseBodyDataOverall) *ListIntervalSkillGroupReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetStatsTime(v int64) *ListIntervalSkillGroupReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyDataBack2Back struct {
	AgentHandleRate         *float32 `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	AnswerRate              *string  `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	AverageRingTime         *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime         *string  `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAgentHandled       *int64   `json:"CallsAgentHandled,omitempty" xml:"CallsAgentHandled,omitempty"`
	CallsAnswered           *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerAnswered   *int64   `json:"CallsCustomerAnswered,omitempty" xml:"CallsCustomerAnswered,omitempty"`
	CallsDialed             *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerAnswerRate      *float32 `json:"CustomerAnswerRate,omitempty" xml:"CustomerAnswerRate,omitempty"`
	MaxCustomerRingTime     *int64   `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	MaxRingTime             *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime             *string  `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCustomerRingTime   *int64   `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	TotalRingTime           *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime           *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataBack2Back) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetAgentHandleRate() *float32 {
	return s.AgentHandleRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetAnswerRate() *string {
	return s.AnswerRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetAverageCustomerRingTime() *float32 {
	return s.AverageCustomerRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetAverageTalkTime() *string {
	return s.AverageTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetCallsAgentHandled() *int64 {
	return s.CallsAgentHandled
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetCallsCustomerAnswered() *int64 {
	return s.CallsCustomerAnswered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetCustomerAnswerRate() *float32 {
	return s.CustomerAnswerRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetMaxCustomerRingTime() *int64 {
	return s.MaxCustomerRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetMaxTalkTime() *string {
	return s.MaxTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetTotalCustomerRingTime() *int64 {
	return s.TotalCustomerRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetAgentHandleRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetAnswerRate(v string) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetAverageCustomerRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetAverageRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetAverageTalkTime(v string) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetCallsAgentHandled(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetCallsAnswered(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetCallsCustomerAnswered(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetCallsDialed(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetCustomerAnswerRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetMaxCustomerRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetMaxRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetMaxTalkTime(v string) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetTotalCustomerRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetTotalRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyDataInbound struct {
	// example:
	//
	// 0
	AbandonRate *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	// example:
	//
	// 0
	AverageAbandonTime *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
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
	// 11
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 5
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 11
	AverageWaitTime *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	// example:
	//
	// 8
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	CallsAbandoned *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
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
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// example:
	//
	// 0
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// example:
	//
	// 2
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 3
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 0
	CallsOverflow *int64 `json:"CallsOverflow,omitempty" xml:"CallsOverflow,omitempty"`
	// example:
	//
	// 3
	CallsQueued          *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	CallsQueuingOverflow *int64 `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	CallsQueuingTimeout  *int64 `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	// example:
	//
	// 3
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 0
	CallsTimeout *int64 `json:"CallsTimeout,omitempty" xml:"CallsTimeout,omitempty"`
	// example:
	//
	// 0.6666666666666666
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// example:
	//
	// 0
	MaxAbandonTime *int64 `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
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
	// 18
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 6
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 18
	MaxWaitTime *int64 `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// example:
	//
	// 19
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
	ServiceLevel20 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	// example:
	//
	// 0
	TotalAbandonTime *int64 `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
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
	// 33
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 9
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 33
	TotalWaitTime *int64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// example:
	//
	// 23
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAbandonRate() *float32 {
	return s.AbandonRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageAbandonTime() *float32 {
	return s.AverageAbandonTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageAbandonedInQueueTime() *float32 {
	return s.AverageAbandonedInQueueTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageAbandonedInRingTime() *float32 {
	return s.AverageAbandonedInRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageWaitTime() *float32 {
	return s.AverageWaitTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsAbandonedInQueue() *int64 {
	return s.CallsAbandonedInQueue
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsAbandonedInRing() *int64 {
	return s.CallsAbandonedInRing
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsOverflow() *int64 {
	return s.CallsOverflow
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsQueued() *int64 {
	return s.CallsQueued
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsQueuingOverflow() *int64 {
	return s.CallsQueuingOverflow
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetCallsTimeout() *int64 {
	return s.CallsTimeout
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxAbandonTime() *int64 {
	return s.MaxAbandonTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxAbandonedInQueueTime() *int64 {
	return s.MaxAbandonedInQueueTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxAbandonedInRingTime() *int64 {
	return s.MaxAbandonedInRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxWaitTime() *int64 {
	return s.MaxWaitTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalAbandonTime() *int64 {
	return s.TotalAbandonTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalAbandonedInQueueTime() *int64 {
	return s.TotalAbandonedInQueueTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalAbandonedInRingTime() *int64 {
	return s.TotalAbandonedInRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalWaitTime() *int64 {
	return s.TotalWaitTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAbandonRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AbandonRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsOverflow(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsOverflow = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsQueued(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsQueuingOverflow(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsQueuingTimeout(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsTimeout(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsTimeout = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetServiceLevel20(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyDataOutbound struct {
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 30
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
	// 5
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 3
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 1
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 0
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// example:
	//
	// 0
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// example:
	//
	// 0
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// example:
	//
	// 2
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 0
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 49
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
	// 5
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 4
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
	// 60
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
	// 5
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 5
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyDataOverall struct {
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
	// 6
	AverageWorkTime     *float32                                                                  `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList []*ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 4927
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// example:
	//
	// 6
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 19
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0.00422315148470254
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
	// 5
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// example:
	//
	// 5
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// example:
	//
	// 0
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// example:
	//
	// 9236
	TotalLoggedInTime *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	// example:
	//
	// 9106
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// example:
	//
	// 13
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 27
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetBreakCodeDetailList() []*ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetBreakCodeDetailList(v []*ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}

type ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetBreakCode(v string) *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetCount(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetDuration(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
