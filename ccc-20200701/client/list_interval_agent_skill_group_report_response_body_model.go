// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentSkillGroupReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntervalAgentSkillGroupReportResponseBody
	GetCode() *string
	SetData(v []*ListIntervalAgentSkillGroupReportResponseBodyData) *ListIntervalAgentSkillGroupReportResponseBody
	GetData() []*ListIntervalAgentSkillGroupReportResponseBodyData
	SetHttpStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntervalAgentSkillGroupReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntervalAgentSkillGroupReportResponseBody
	GetRequestId() *string
}

type ListIntervalAgentSkillGroupReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListIntervalAgentSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2778FA12-EDD6-42AA-9B15-AF855072E5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) GetData() []*ListIntervalAgentSkillGroupReportResponseBodyData {
	return s.Data
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetCode(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetData(v []*ListIntervalAgentSkillGroupReportResponseBodyData) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetMessage(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetRequestId(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyData struct {
	Back2Back *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	Inbound   *ListIntervalAgentSkillGroupReportResponseBodyDataInbound   `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal  *ListIntervalAgentSkillGroupReportResponseBodyDataInternal  `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound  *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound  `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalAgentSkillGroupReportResponseBodyDataOverall   `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// 1620291600000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetBack2Back() *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	return s.Back2Back
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetInbound() *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetInternal() *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	return s.Internal
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetOutbound() *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetOverall() *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	return s.Overall
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetBack2Back(v *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Back2Back = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetInbound(v *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetInternal(v *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Internal = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetOutbound(v *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetOverall(v *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetStatsTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back struct {
	// example:
	//
	// 100
	AgentAnswerRate *float32 `json:"AgentAnswerRate,omitempty" xml:"AgentAnswerRate,omitempty"`
	AgentHandleRate *float32 `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	// example:
	//
	// 0.5
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 100
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	// example:
	//
	// 100
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 100
	AverageTalkTime   *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAgentHandled *int64 `json:"CallsAgentHandled,omitempty" xml:"CallsAgentHandled,omitempty"`
	// example:
	//
	// 100
	CallsAnswered         *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerAnswered *int64 `json:"CallsCustomerAnswered,omitempty" xml:"CallsCustomerAnswered,omitempty"`
	// example:
	//
	// 100
	CallsCustomerHandled *int64 `json:"CallsCustomerHandled,omitempty" xml:"CallsCustomerHandled,omitempty"`
	// example:
	//
	// 100
	CallsDialed        *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerAnswerRate *float32 `json:"CustomerAnswerRate,omitempty" xml:"CustomerAnswerRate,omitempty"`
	// example:
	//
	// 100
	CustomerHandleRate *float32 `json:"CustomerHandleRate,omitempty" xml:"CustomerHandleRate,omitempty"`
	// example:
	//
	// 100
	MaxCustomerRingTime *int64 `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	// example:
	//
	// 100
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	TotalCustomerRingTime *int64 `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	// example:
	//
	// 100
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAgentAnswerRate() *float32 {
	return s.AgentAnswerRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAgentHandleRate() *float32 {
	return s.AgentHandleRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAverageCustomerRingTime() *float32 {
	return s.AverageCustomerRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCallsAgentHandled() *int64 {
	return s.CallsAgentHandled
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCallsCustomerAnswered() *int64 {
	return s.CallsCustomerAnswered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCallsCustomerHandled() *int64 {
	return s.CallsCustomerHandled
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCustomerAnswerRate() *float32 {
	return s.CustomerAnswerRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetCustomerHandleRate() *float32 {
	return s.CustomerHandleRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetMaxCustomerRingTime() *int64 {
	return s.MaxCustomerRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetTotalCustomerRingTime() *int64 {
	return s.TotalCustomerRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAgentAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AgentAnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAgentHandleRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageCustomerRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsAgentHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsCustomerAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsCustomerHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsCustomerHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCustomerAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCustomerHandleRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CustomerHandleRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxCustomerRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalCustomerRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataInbound struct {
	AverageFirstResponseTime *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	// example:
	//
	// 100
	AverageHoldTime     *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageResponseTime *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
	// example:
	//
	// 11
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 100
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 100
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 100
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// example:
	//
	// 10
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// example:
	//
	// 100
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// example:
	//
	// 100
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// example:
	//
	// 7
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 100
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 100
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 100
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 0.5
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// example:
	//
	// 100
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 100
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 100
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 0.5
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// example:
	//
	// 100
	TotalHoldTime               *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalMessagesSent           *int64 `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	TotalMessagesSentByAgent    *int64 `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	TotalMessagesSentByCustomer *int64 `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	// example:
	//
	// 100
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 85
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataInternal struct {
	// example:
	//
	// 100
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 100
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 100
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 100
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 100
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 100
	CallsTalk   *int64 `json:"CallsTalk,omitempty" xml:"CallsTalk,omitempty"`
	CallsTalked *int64 `json:"CallsTalked,omitempty" xml:"CallsTalked,omitempty"`
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInternal) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsTalk() *int64 {
	return s.CallsTalk
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetCallsTalked() *int64 {
	return s.CallsTalked
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsTalk(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsTalk = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsTalked(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsTalked = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOutbound struct {
	// example:
	//
	// 0.5
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 100
	AverageDialingTime *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// example:
	//
	// 100
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// example:
	//
	// 100
	AverageRingTime *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 100
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 100
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 100
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 100
	CallsAttendedTransferIn *int64 `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	// example:
	//
	// 100
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// example:
	//
	// 100
	CallsBlindTransferIn *int64 `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	// example:
	//
	// 100
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// example:
	//
	// 100
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 100
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 100
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 100
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// example:
	//
	// 100
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 100
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 100
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 100
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// example:
	//
	// 100
	TotalDialingTime *int64 `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	// example:
	//
	// 100
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// example:
	//
	// 100
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 100
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOverall struct {
	// example:
	//
	// 100
	AverageBreakTime *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	// example:
	//
	// 100
	AverageHoldTime *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	// example:
	//
	// 100
	AverageReadyTime *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// example:
	//
	// 100
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 100
	AverageWorkTime     *float32                                                                       `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList []*ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	FirstCheckInTime *int64 `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	LastCheckOutTime *int64 `json:"LastCheckOutTime,omitempty" xml:"LastCheckOutTime,omitempty"`
	// example:
	//
	// 100
	LastCheckoutTime *int64 `json:"LastCheckoutTime,omitempty" xml:"LastCheckoutTime,omitempty"`
	// example:
	//
	// 100
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// example:
	//
	// 100
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	// example:
	//
	// 100
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0.4
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// example:
	//
	// 1.4
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// example:
	//
	// 100
	SatisfactionRate *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// example:
	//
	// 100
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// example:
	//
	// 100
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// example:
	//
	// 100
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// example:
	//
	// 100
	TotalHoldTime *int64 `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	// example:
	//
	// 100
	TotalLoggedInTime                 *int64  `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	TotalOffSiteLoggedInTime          *string `json:"TotalOffSiteLoggedInTime,omitempty" xml:"TotalOffSiteLoggedInTime,omitempty"`
	TotalOfficePhoneLoggedInTime      *string `json:"TotalOfficePhoneLoggedInTime,omitempty" xml:"TotalOfficePhoneLoggedInTime,omitempty"`
	TotalOnSiteLoggedInTime           *string `json:"TotalOnSiteLoggedInTime,omitempty" xml:"TotalOnSiteLoggedInTime,omitempty"`
	TotalOutboundScenarioLoggedInTime *int64  `json:"TotalOutboundScenarioLoggedInTime,omitempty" xml:"TotalOutboundScenarioLoggedInTime,omitempty"`
	// example:
	//
	// 100
	TotalOutboundScenarioReadyTime *int64 `json:"TotalOutboundScenarioReadyTime,omitempty" xml:"TotalOutboundScenarioReadyTime,omitempty"`
	// example:
	//
	// 100
	TotalOutboundScenarioTime *int64 `json:"TotalOutboundScenarioTime,omitempty" xml:"TotalOutboundScenarioTime,omitempty"`
	// example:
	//
	// 100
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// example:
	//
	// 100
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 100
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetBreakCodeDetailList() []*ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetFirstCheckInTime() *int64 {
	return s.FirstCheckInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetLastCheckOutTime() *int64 {
	return s.LastCheckOutTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetLastCheckoutTime() *int64 {
	return s.LastCheckoutTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOffSiteLoggedInTime() *string {
	return s.TotalOffSiteLoggedInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOfficePhoneLoggedInTime() *string {
	return s.TotalOfficePhoneLoggedInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOnSiteLoggedInTime() *string {
	return s.TotalOnSiteLoggedInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOutboundScenarioLoggedInTime() *int64 {
	return s.TotalOutboundScenarioLoggedInTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOutboundScenarioReadyTime() *int64 {
	return s.TotalOutboundScenarioReadyTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalOutboundScenarioTime() *int64 {
	return s.TotalOutboundScenarioTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetBreakCodeDetailList(v []*ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetFirstCheckInTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetLastCheckOutTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.LastCheckOutTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetLastCheckoutTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.LastCheckoutTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOffSiteLoggedInTime(v string) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOffSiteLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOfficePhoneLoggedInTime(v string) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOfficePhoneLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOnSiteLoggedInTime(v string) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOnSiteLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOutboundScenarioLoggedInTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOutboundScenarioLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOutboundScenarioTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetBreakCode(v string) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetCount(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetDuration(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
