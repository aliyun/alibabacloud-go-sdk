// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentSkillGroupReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHistoricalAgentSkillGroupReportResponseBody
	GetCode() *string
	SetData(v *ListHistoricalAgentSkillGroupReportResponseBodyData) *ListHistoricalAgentSkillGroupReportResponseBody
	GetData() *ListHistoricalAgentSkillGroupReportResponseBodyData
	SetHttpStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListHistoricalAgentSkillGroupReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHistoricalAgentSkillGroupReportResponseBody
	GetRequestId() *string
}

type ListHistoricalAgentSkillGroupReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListHistoricalAgentSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) GetData() *ListHistoricalAgentSkillGroupReportResponseBodyData {
	return s.Data
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetCode(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetData(v *ListHistoricalAgentSkillGroupReportResponseBodyData) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetMessage(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetRequestId(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyData struct {
	List []*ListHistoricalAgentSkillGroupReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) GetList() []*ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	return s.List
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetList(v []*ListHistoricalAgentSkillGroupReportResponseBodyDataList) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetPageSize(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataList struct {
	// example:
	//
	// agent1@ccc-test
	AgentId   *string                                                           `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string                                                           `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	Back2Back *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	// example:
	//
	// 1001
	DisplayId *string                                                          `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	Inbound   *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal  *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound  *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// skillgroup1@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// Default
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetBack2Back() *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	return s.Back2Back
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetDisplayId() *string {
	return s.DisplayId
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetInbound() *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	return s.Inbound
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetInternal() *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	return s.Internal
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetOutbound() *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	return s.Outbound
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetOverall() *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	return s.Overall
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetAgentId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetAgentName(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetBack2Back(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Back2Back = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetDisplayId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.DisplayId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetInbound(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetInternal(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Internal = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetOutbound(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetOverall(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Overall = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetSkillGroupId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetSkillGroupName(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back struct {
	// example:
	//
	// 0.5
	AgentAnswerRate *float32 `json:"AgentAnswerRate,omitempty" xml:"AgentAnswerRate,omitempty"`
	AgentHandleRate *float32 `json:"AgentHandleRate,omitempty" xml:"AgentHandleRate,omitempty"`
	// example:
	//
	// 0.6
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
	// 0.5
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

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAgentAnswerRate() *float32 {
	return s.AgentAnswerRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAgentHandleRate() *float32 {
	return s.AgentHandleRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAverageCustomerRingTime() *float32 {
	return s.AverageCustomerRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCallsAgentHandled() *int64 {
	return s.CallsAgentHandled
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCallsCustomerAnswered() *int64 {
	return s.CallsCustomerAnswered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCallsCustomerHandled() *int64 {
	return s.CallsCustomerHandled
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCustomerAnswerRate() *float32 {
	return s.CustomerAnswerRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetCustomerHandleRate() *float32 {
	return s.CustomerHandleRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetMaxCustomerRingTime() *int64 {
	return s.MaxCustomerRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetTotalCustomerRingTime() *int64 {
	return s.TotalCustomerRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAgentAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AgentAnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAgentHandleRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AgentHandleRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageCustomerRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsAgentHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsCustomerAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsCustomerAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsCustomerHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsCustomerHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCustomerAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CustomerAnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCustomerHandleRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CustomerHandleRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxCustomerRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalCustomerRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound struct {
	AverageFirstResponseTime *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	// example:
	//
	// 100
	AverageHoldTime     *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageResponseTime *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
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
	// 100
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
	// 100
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageFirstResponseTime() *float32 {
	return s.AverageFirstResponseTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageResponseTime() *float32 {
	return s.AverageResponseTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSent() *int64 {
	return s.TotalMessagesSent
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSentByAgent() *int64 {
	return s.TotalMessagesSentByAgent
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalMessagesSentByCustomer() *int64 {
	return s.TotalMessagesSentByCustomer
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageFirstResponseTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageResponseTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSent(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByAgent(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByCustomer(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal struct {
	// example:
	//
	// 100
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
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

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsTalk() *int64 {
	return s.CallsTalk
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetCallsTalked() *int64 {
	return s.CallsTalked
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetAverageTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsTalk(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsTalk = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsTalked(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsTalked = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound struct {
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
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 100
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 1.4
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

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAverageDialingTime() *float32 {
	return s.AverageDialingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAverageRingTime() *float32 {
	return s.AverageRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsAttendedTransferIn() *int64 {
	return s.CallsAttendedTransferIn
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsBlindTransferIn() *int64 {
	return s.CallsBlindTransferIn
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsHold() *int64 {
	return s.CallsHold
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetCallsRinged() *int64 {
	return s.CallsRinged
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall struct {
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
	AverageWorkTime     *float32                                                                             `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList []*ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 1686030515000
	FirstCheckInTime *int64 `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	// example:
	//
	// 1686030515000
	LastCheckOutTime *int64 `json:"LastCheckOutTime,omitempty" xml:"LastCheckOutTime,omitempty"`
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
	// 0.5
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// example:
	//
	// 1.4
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
	TotalOffSiteLggedInTime           *int64  `json:"TotalOffSiteLggedInTime,omitempty" xml:"TotalOffSiteLggedInTime,omitempty"`
	TotalOfficePhoneLoggedInTime      *int64  `json:"TotalOfficePhoneLoggedInTime,omitempty" xml:"TotalOfficePhoneLoggedInTime,omitempty"`
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

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetAverageBreakTime() *float32 {
	return s.AverageBreakTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetAverageHoldTime() *float32 {
	return s.AverageHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetAverageReadyTime() *float32 {
	return s.AverageReadyTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetAverageTalkTime() *float32 {
	return s.AverageTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetAverageWorkTime() *float32 {
	return s.AverageWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetBreakCodeDetailList() []*ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	return s.BreakCodeDetailList
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetFirstCheckInTime() *int64 {
	return s.FirstCheckInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetLastCheckOutTime() *int64 {
	return s.LastCheckOutTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetMaxBreakTime() *int64 {
	return s.MaxBreakTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetMaxHoldTime() *int64 {
	return s.MaxHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetSatisfactionRate() *float32 {
	return s.SatisfactionRate
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalHoldTime() *int64 {
	return s.TotalHoldTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOffSiteLggedInTime() *int64 {
	return s.TotalOffSiteLggedInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOfficePhoneLoggedInTime() *int64 {
	return s.TotalOfficePhoneLoggedInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOnSiteLoggedInTime() *string {
	return s.TotalOnSiteLoggedInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOutboundScenarioLoggedInTime() *int64 {
	return s.TotalOutboundScenarioLoggedInTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOutboundScenarioReadyTime() *int64 {
	return s.TotalOutboundScenarioReadyTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalOutboundScenarioTime() *int64 {
	return s.TotalOutboundScenarioTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetBreakCodeDetailList(v []*ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetFirstCheckInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetLastCheckOutTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.LastCheckOutTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOffSiteLggedInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOffSiteLggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOfficePhoneLoggedInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOfficePhoneLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOnSiteLoggedInTime(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOnSiteLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOutboundScenarioLoggedInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOutboundScenarioTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) Validate() error {
	return dara.Validate(s)
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList struct {
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

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetBreakCode() *string {
	return s.BreakCode
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetCount() *int64 {
	return s.Count
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GetDuration() *int64 {
	return s.Duration
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetBreakCode(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetCount(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetDuration(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) Validate() error {
	return dara.Validate(s)
}
