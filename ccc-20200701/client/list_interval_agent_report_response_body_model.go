// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntervalAgentReportResponseBody
	GetCode() *string
	SetData(v []*ListIntervalAgentReportResponseBodyData) *ListIntervalAgentReportResponseBody
	GetData() []*ListIntervalAgentReportResponseBodyData
	SetHttpStatusCode(v int32) *ListIntervalAgentReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIntervalAgentReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntervalAgentReportResponseBody
	GetRequestId() *string
}

type ListIntervalAgentReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListIntervalAgentReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListIntervalAgentReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntervalAgentReportResponseBody) GetData() []*ListIntervalAgentReportResponseBodyData {
	return s.Data
}

func (s *ListIntervalAgentReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntervalAgentReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntervalAgentReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntervalAgentReportResponseBody) SetCode(v string) *ListIntervalAgentReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetData(v []*ListIntervalAgentReportResponseBodyData) *ListIntervalAgentReportResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalAgentReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetMessage(v string) *ListIntervalAgentReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetRequestId(v string) *ListIntervalAgentReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyData struct {
	Back2Back *ListIntervalAgentReportResponseBodyDataBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	Inbound   *ListIntervalAgentReportResponseBodyDataInbound   `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal  *ListIntervalAgentReportResponseBodyDataInternal  `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound  *ListIntervalAgentReportResponseBodyDataOutbound  `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalAgentReportResponseBodyDataOverall   `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// 1620291600000
	StatsTime *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyData) GetBack2Back() *ListIntervalAgentReportResponseBodyDataBack2Back {
	return s.Back2Back
}

func (s *ListIntervalAgentReportResponseBodyData) GetInbound() *ListIntervalAgentReportResponseBodyDataInbound {
	return s.Inbound
}

func (s *ListIntervalAgentReportResponseBodyData) GetInternal() *ListIntervalAgentReportResponseBodyDataInternal {
	return s.Internal
}

func (s *ListIntervalAgentReportResponseBodyData) GetOutbound() *ListIntervalAgentReportResponseBodyDataOutbound {
	return s.Outbound
}

func (s *ListIntervalAgentReportResponseBodyData) GetOverall() *ListIntervalAgentReportResponseBodyDataOverall {
	return s.Overall
}

func (s *ListIntervalAgentReportResponseBodyData) GetStatsTime() *int64 {
	return s.StatsTime
}

func (s *ListIntervalAgentReportResponseBodyData) SetBack2Back(v *ListIntervalAgentReportResponseBodyDataBack2Back) *ListIntervalAgentReportResponseBodyData {
	s.Back2Back = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetInbound(v *ListIntervalAgentReportResponseBodyDataInbound) *ListIntervalAgentReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetInternal(v *ListIntervalAgentReportResponseBodyDataInternal) *ListIntervalAgentReportResponseBodyData {
	s.Internal = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetOutbound(v *ListIntervalAgentReportResponseBodyDataOutbound) *ListIntervalAgentReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetOverall(v *ListIntervalAgentReportResponseBodyDataOverall) *ListIntervalAgentReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetStatsTime(v int64) *ListIntervalAgentReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataBack2Back struct {
	AgentHandleRate         *float32 `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	AnswerRate              *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	AverageRingTime         *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime         *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAgentHandled       *int64   `json:"CallsAgentHandled,omitempty" xml:"CallsAgentHandled,omitempty"`
	CallsAnswered           *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerAnswered   *int64   `json:"CallsCustomerAnswered,omitempty" xml:"CallsCustomerAnswered,omitempty"`
	CallsDialed             *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerAnswerRate      *float32 `json:"CustomerAnswerRate,omitempty" xml:"CustomerAnswerRate,omitempty"`
	MaxCustomerRingTime     *int64   `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	MaxRingTime             *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime             *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCustomerRingTime   *int64   `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	TotalRingTime           *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime           *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataBack2Back) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetAgentHandleRate() *float32 {
	return s.AgentHandleRate
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetAverageCustomerRingTime() *float32 {
	return s.AverageCustomerRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetCallsAgentHandled() *int64 {
	return s.CallsAgentHandled
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetCallsCustomerAnswered() *int64 {
	return s.CallsCustomerAnswered
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetCustomerAnswerRate() *float32 {
	return s.CustomerAnswerRate
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetMaxCustomerRingTime() *int64 {
	return s.MaxCustomerRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetTotalCustomerRingTime() *int64 {
	return s.TotalCustomerRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetAgentHandleRate(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetAnswerRate(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetAverageCustomerRingTime(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetAverageRingTime(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetCallsAgentHandled(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetCallsAnswered(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetCallsCustomerAnswered(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetCallsDialed(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetCustomerAnswerRate(v float32) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetMaxCustomerRingTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetMaxRingTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetTotalCustomerRingTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetTotalRingTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataInbound struct {
	AccessChannelTypeDetails []*ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails `json:"AccessChannelTypeDetails,omitempty" xml:"AccessChannelTypeDetails,omitempty" type:"Repeated"`
	AverageFirstResponseTime *float32                                                                  `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
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
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
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
	// 0
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 0
	CallsHold *int64 `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	// example:
	//
	// 0
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 0
	CallsRinged *int64 `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	// example:
	//
	// 0
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
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
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	ServiceLevel15               *float32 `json:"ServiceLevel15,omitempty" xml:"ServiceLevel15,omitempty"`
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
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataInbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAccessChannelTypeDetails() []*ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails {
	return s.AccessChannelTypeDetails
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetServiceLevel15() *float32 {
	return s.ServiceLevel15
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAccessChannelTypeDetails(v []*ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AccessChannelTypeDetails = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetServiceLevel15(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.ServiceLevel15 = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails struct {
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	CallsOffered      *int64  `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) SetAccessChannelType(v string) *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails {
	s.AccessChannelType = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) SetCallsOffered(v int64) *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInboundAccessChannelTypeDetails) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataInternal struct {
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered   *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsDialed     *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHandled    *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsOffered    *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsTalked     *int64   `json:"CallsTalked,omitempty" xml:"CallsTalked,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataInternal) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataInternal) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) GetCallsTalked() *int64 {
	return s.CallsTalked
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetCallsAnswered(v int64) *ListIntervalAgentReportResponseBodyDataInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetCallsDialed(v int64) *ListIntervalAgentReportResponseBodyDataInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetCallsHandled(v int64) *ListIntervalAgentReportResponseBodyDataInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetCallsOffered(v int64) *ListIntervalAgentReportResponseBodyDataInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) SetCallsTalked(v int64) *ListIntervalAgentReportResponseBodyDataInternal {
	s.CallsTalked = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInternal) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataOutbound struct {
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
	// 0
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 1
	AverageWorkTime *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
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
	// 5
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
	// 60
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
	// 2
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
	// 148
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
	// 4
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataOverall struct {
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
	// 1
	AverageWorkTime     *float32                                                             `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList []*ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	FirstCheckInTime *int64 `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	LastCheckOutTime *int64 `json:"LastCheckOutTime,omitempty" xml:"LastCheckOutTime,omitempty"`
	// example:
	//
	// 0
	LastCheckoutTime *int64 `json:"LastCheckoutTime,omitempty" xml:"LastCheckoutTime,omitempty"`
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
	// 435
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 2
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
	// 1
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
	// 914
	TotalLoggedInTime        *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	TotalOffSiteLoggedInTime *int64 `json:"TotalOffSiteLoggedInTime,omitempty" xml:"TotalOffSiteLoggedInTime,omitempty"`
	// example:
	//
	// 已弃用，请使用TotalOffSiteLoggedInTime代替此参数
	TotalOffSiteOnlineTime       *int64 `json:"TotalOffSiteOnlineTime,omitempty" xml:"TotalOffSiteOnlineTime,omitempty"`
	TotalOfficePhoneLoggedInTime *int64 `json:"TotalOfficePhoneLoggedInTime,omitempty" xml:"TotalOfficePhoneLoggedInTime,omitempty"`
	// example:
	//
	// 已弃用，请使用TotalOfficePhoneLoggedInTime代替此参数
	TotalOfficePhoneOnlineTime *int64 `json:"TotalOfficePhoneOnlineTime,omitempty" xml:"TotalOfficePhoneOnlineTime,omitempty"`
	TotalOnSiteLoggedInTime    *int64 `json:"TotalOnSiteLoggedInTime,omitempty" xml:"TotalOnSiteLoggedInTime,omitempty"`
	// example:
	//
	// 已弃用，请使用TotalOnSiteLoggedInTime代替此参数
	TotalOnSiteOnlineTime             *int64 `json:"TotalOnSiteOnlineTime,omitempty" xml:"TotalOnSiteOnlineTime,omitempty"`
	TotalOutboundScenarioLoggedInTime *int64 `json:"TotalOutboundScenarioLoggedInTime,omitempty" xml:"TotalOutboundScenarioLoggedInTime,omitempty"`
	TotalOutboundScenarioReadyTime    *int64 `json:"TotalOutboundScenarioReadyTime,omitempty" xml:"TotalOutboundScenarioReadyTime,omitempty"`
	// example:
	//
	// 已弃用，请使用TotalOutboundScenarioLoggedInTime代替此参数
	TotalOutboundScenarioTime *int64 `json:"TotalOutboundScenarioTime,omitempty" xml:"TotalOutboundScenarioTime,omitempty"`
	// example:
	//
	// 763
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 4
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataOverall) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetBreakCodeDetailList() []*ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetFirstCheckInTime() *int64 {
	return s.FirstCheckInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetLastCheckOutTime() *int64 {
	return s.LastCheckOutTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetLastCheckoutTime() *int64 {
	return s.LastCheckoutTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOffSiteLoggedInTime() *int64 {
	return s.TotalOffSiteLoggedInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOffSiteOnlineTime() *int64 {
	return s.TotalOffSiteOnlineTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOfficePhoneLoggedInTime() *int64 {
	return s.TotalOfficePhoneLoggedInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOfficePhoneOnlineTime() *int64 {
	return s.TotalOfficePhoneOnlineTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOnSiteLoggedInTime() *int64 {
	return s.TotalOnSiteLoggedInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOnSiteOnlineTime() *int64 {
	return s.TotalOnSiteOnlineTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOutboundScenarioLoggedInTime() *int64 {
	return s.TotalOutboundScenarioLoggedInTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOutboundScenarioReadyTime() *int64 {
	return s.TotalOutboundScenarioReadyTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalOutboundScenarioTime() *int64 {
	return s.TotalOutboundScenarioTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetBreakCodeDetailList(v []*ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) *ListIntervalAgentReportResponseBodyDataOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetFirstCheckInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetLastCheckOutTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.LastCheckOutTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetLastCheckoutTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.LastCheckoutTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOffSiteLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOffSiteLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOffSiteOnlineTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOffSiteOnlineTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOfficePhoneLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOfficePhoneLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOfficePhoneOnlineTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOfficePhoneOnlineTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOnSiteLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOnSiteLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOnSiteOnlineTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOnSiteOnlineTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOutboundScenarioLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOutboundScenarioLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalOutboundScenarioTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) Validate() error {
	return dara.Validate(s)
}

type ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) SetBreakCode(v string) *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) SetCount(v int64) *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) SetDuration(v int64) *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
