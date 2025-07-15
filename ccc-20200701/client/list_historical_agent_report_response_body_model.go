// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHistoricalAgentReportResponseBody
	GetCode() *string
	SetData(v *ListHistoricalAgentReportResponseBodyData) *ListHistoricalAgentReportResponseBody
	GetData() *ListHistoricalAgentReportResponseBodyData
	SetHttpStatusCode(v int32) *ListHistoricalAgentReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHistoricalAgentReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHistoricalAgentReportResponseBody
	GetRequestId() *string
}

type ListHistoricalAgentReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListHistoricalAgentReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalAgentReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHistoricalAgentReportResponseBody) GetData() *ListHistoricalAgentReportResponseBodyData {
	return s.Data
}

func (s *ListHistoricalAgentReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHistoricalAgentReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoricalAgentReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoricalAgentReportResponseBody) SetCode(v string) *ListHistoricalAgentReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetData(v *ListHistoricalAgentReportResponseBodyData) *ListHistoricalAgentReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalAgentReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetMessage(v string) *ListHistoricalAgentReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetRequestId(v string) *ListHistoricalAgentReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyData struct {
	List []*ListHistoricalAgentReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyData) GetList() []*ListHistoricalAgentReportResponseBodyDataList {
	return s.List
}

func (s *ListHistoricalAgentReportResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalAgentReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalAgentReportResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHistoricalAgentReportResponseBodyData) SetList(v []*ListHistoricalAgentReportResponseBodyDataList) *ListHistoricalAgentReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetPageSize(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataList struct {
	// example:
	//
	// agent1@ccc-test
	AgentId   *string                                                 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string                                                 `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	Back2Back *ListHistoricalAgentReportResponseBodyDataListBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	// example:
	//
	// 001
	DisplayId *string                                                `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	Inbound   *ListHistoricalAgentReportResponseBodyDataListInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal  *ListHistoricalAgentReportResponseBodyDataListInternal `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound  *ListHistoricalAgentReportResponseBodyDataListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListHistoricalAgentReportResponseBodyDataListOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIds   *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	SkillGroupNames *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetBack2Back() *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	return s.Back2Back
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetDisplayId() *string {
	return s.DisplayId
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetInbound() *ListHistoricalAgentReportResponseBodyDataListInbound {
	return s.Inbound
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetInternal() *ListHistoricalAgentReportResponseBodyDataListInternal {
	return s.Internal
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetOutbound() *ListHistoricalAgentReportResponseBodyDataListOutbound {
	return s.Outbound
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetOverall() *ListHistoricalAgentReportResponseBodyDataListOverall {
	return s.Overall
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListHistoricalAgentReportResponseBodyDataList) GetSkillGroupNames() *string {
	return s.SkillGroupNames
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetAgentId(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetAgentName(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetBack2Back(v *ListHistoricalAgentReportResponseBodyDataListBack2Back) *ListHistoricalAgentReportResponseBodyDataList {
	s.Back2Back = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetDisplayId(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.DisplayId = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetInbound(v *ListHistoricalAgentReportResponseBodyDataListInbound) *ListHistoricalAgentReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetInternal(v *ListHistoricalAgentReportResponseBodyDataListInternal) *ListHistoricalAgentReportResponseBodyDataList {
	s.Internal = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetOutbound(v *ListHistoricalAgentReportResponseBodyDataListOutbound) *ListHistoricalAgentReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetOverall(v *ListHistoricalAgentReportResponseBodyDataListOverall) *ListHistoricalAgentReportResponseBodyDataList {
	s.Overall = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetSkillGroupIds(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetSkillGroupNames(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.SkillGroupNames = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListBack2Back struct {
	AgentHandleRate         *string `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	AnswerRate              *string `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageCustomerRingTime *string `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	AverageRingTime         *string `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime         *string `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAgentHandled       *string `json:"CallsAgentHandled,omitempty" xml:"CallsAgentHandled,omitempty"`
	CallsAnswered           *string `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerAnswered   *string `json:"CallsCustomerAnswered,omitempty" xml:"CallsCustomerAnswered,omitempty"`
	CallsDialed             *string `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerAnswerRate      *string `json:"CustomerAnswerRate,omitempty" xml:"CustomerAnswerRate,omitempty"`
	MaxCustomerRingTime     *string `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	MaxRingTime             *string `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime             *string `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCustomerRingTime   *string `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	TotalRingTime           *string `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime           *string `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListBack2Back) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetAgentHandleRate() *string {
	return s.AgentHandleRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetAnswerRate() *string {
	return s.AnswerRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetAverageCustomerRingTime() *string {
	return s.AverageCustomerRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetAverageRingTime() *string {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetAverageTalkTime() *string {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetCallsAgentHandled() *string {
	return s.CallsAgentHandled
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetCallsAnswered() *string {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetCallsCustomerAnswered() *string {
	return s.CallsCustomerAnswered
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetCallsDialed() *string {
	return s.CallsDialed
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetCustomerAnswerRate() *string {
	return s.CustomerAnswerRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetMaxCustomerRingTime() *string {
	return s.MaxCustomerRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetMaxRingTime() *string {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetMaxTalkTime() *string {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetTotalCustomerRingTime() *string {
	return s.TotalCustomerRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetTotalRingTime() *string {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) GetTotalTalkTime() *string {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetAgentHandleRate(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetAnswerRate(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetAverageCustomerRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetAverageRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetAverageTalkTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetCallsAgentHandled(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetCallsAnswered(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetCallsCustomerAnswered(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetCallsDialed(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetCustomerAnswerRate(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetMaxCustomerRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetMaxRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetMaxTalkTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetTotalCustomerRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetTotalRingTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) SetTotalTalkTime(v string) *ListHistoricalAgentReportResponseBodyDataListBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListInbound struct {
	AccessChannelTypeDetails []*ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails `json:"AccessChannelTypeDetails,omitempty" xml:"AccessChannelTypeDetails,omitempty" type:"Repeated"`
	AverageFirstResponseTime *float32                                                                        `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
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
	TotalHoldTime               *int64  `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalMessagesSent           *int64  `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	TotalMessagesSentByAgent    *int64  `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	TotalMessagesSentByCustomer *string `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
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

func (s ListHistoricalAgentReportResponseBodyDataListInbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAccessChannelTypeDetails() []*ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails {
	return s.AccessChannelTypeDetails
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetServiceLevel15() *float32 {
	return s.ServiceLevel15
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalMessagesSentByCustomer() *string {
	return s.TotalMessagesSentByCustomer
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAccessChannelTypeDetails(v []*ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AccessChannelTypeDetails = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageFirstResponseTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageResponseTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetServiceLevel15(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.ServiceLevel15 = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalMessagesSent(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalMessagesSentByAgent(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalMessagesSentByCustomer(v string) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails struct {
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	CallsOffered      *int64  `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) SetAccessChannelType(v string) *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails {
	s.AccessChannelType = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) SetCallsOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInboundAccessChannelTypeDetails) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListInternal struct {
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered   *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsDialed     *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHandled    *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsOffered    *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsTalked     *int64   `json:"CallsTalked,omitempty" xml:"CallsTalked,omitempty"`
	MaxTalkTime     *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalTalkTime   *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListInternal) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListInternal) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetCallsTalked() *int64 {
	return s.CallsTalked
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetCallsAnswered(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetCallsDialed(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetCallsHandled(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetCallsOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetCallsTalked(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.CallsTalked = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInternal {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInternal) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListOutbound struct {
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

func (s ListHistoricalAgentReportResponseBodyDataListOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListOverall struct {
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
	AverageWorkTime     *float32                                                                   `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList []*ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	FirstCheckInTime    *int64                                                                     `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	LastCheckOutTime    *int64                                                                     `json:"LastCheckOutTime,omitempty" xml:"LastCheckOutTime,omitempty"`
	// example:
	//
	// 0
	MaxBreakTime *int64 `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	// example:
	//
	// 0
	MaxHoldTime *int64 `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
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

func (s ListHistoricalAgentReportResponseBodyDataListOverall) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetBreakCodeDetailList() []*ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetFirstCheckInTime() *int64 {
	return s.FirstCheckInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetLastCheckOutTime() *int64 {
	return s.LastCheckOutTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOffSiteLoggedInTime() *int64 {
	return s.TotalOffSiteLoggedInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOffSiteOnlineTime() *int64 {
	return s.TotalOffSiteOnlineTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOfficePhoneLoggedInTime() *int64 {
	return s.TotalOfficePhoneLoggedInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOfficePhoneOnlineTime() *int64 {
	return s.TotalOfficePhoneOnlineTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOnSiteLoggedInTime() *int64 {
	return s.TotalOnSiteLoggedInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOnSiteOnlineTime() *int64 {
	return s.TotalOnSiteOnlineTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOutboundScenarioLoggedInTime() *int64 {
	return s.TotalOutboundScenarioLoggedInTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOutboundScenarioReadyTime() *int64 {
	return s.TotalOutboundScenarioReadyTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalOutboundScenarioTime() *int64 {
	return s.TotalOutboundScenarioTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetBreakCodeDetailList(v []*ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetFirstCheckInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetLastCheckOutTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.LastCheckOutTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOffSiteLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOffSiteLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOffSiteOnlineTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOffSiteOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOfficePhoneLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOfficePhoneLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOfficePhoneOnlineTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOfficePhoneOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOnSiteLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOnSiteLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOnSiteOnlineTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOnSiteOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOutboundScenarioLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalOutboundScenarioTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) SetBreakCode(v string) *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) SetCount(v int64) *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) SetDuration(v int64) *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
