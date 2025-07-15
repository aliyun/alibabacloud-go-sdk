// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSummaryReportsSinceMidnightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetMessage() *string
	SetPagedAgentSummaryReport(v *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetPagedAgentSummaryReport() *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport
	SetRequestId(v string) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentSummaryReportsSinceMidnightResponseBody
	GetSuccess() *bool
}

type ListAgentSummaryReportsSinceMidnightResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode          *int32                                                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message                 *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PagedAgentSummaryReport *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport `json:"PagedAgentSummaryReport,omitempty" xml:"PagedAgentSummaryReport,omitempty" type:"Struct"`
	// example:
	//
	// 27DD30C4-CAE2-481A-97CC-D3C54625341D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetPagedAgentSummaryReport() *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport {
	return s.PagedAgentSummaryReport
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetCode(v string) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetHttpStatusCode(v int32) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetMessage(v string) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetPagedAgentSummaryReport(v *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.PagedAgentSummaryReport = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetRequestId(v string) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) SetSuccess(v bool) *ListAgentSummaryReportsSinceMidnightResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport struct {
	List []*ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) GetList() []*ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	return s.List
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) SetList(v []*ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport {
	s.List = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) SetPageNumber(v int32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport {
	s.PageNumber = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) SetPageSize(v int32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport {
	s.PageSize = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) SetTotalCount(v int32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport {
	s.TotalCount = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReport) Validate() error {
	return dara.Validate(s)
}

type ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList struct {
	// example:
	//
	// agent@ccc-test
	AgentId   *string                                                                             `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName *string                                                                             `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	Inbound   *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent
	LoginName *string                                                                              `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Outbound  *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// example:
	//
	// ["skillgroup1@ccc-test","skillgroup2@ccc-test"]
	SkillGroupIds   *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	SkillGroupNames *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	// example:
	//
	// 2018-09-13 00:00:00
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetInbound() *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	return s.Inbound
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetLoginName() *string {
	return s.LoginName
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetOutbound() *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	return s.Outbound
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetOverall() *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	return s.Overall
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetSkillGroupIds() *string {
	return s.SkillGroupIds
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetSkillGroupNames() *string {
	return s.SkillGroupNames
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetAgentId(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.AgentId = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetAgentName(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.AgentName = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetInbound(v *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.Inbound = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetInstanceId(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.InstanceId = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetLoginName(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.LoginName = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetOutbound(v *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.Outbound = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetOverall(v *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.Overall = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetSkillGroupIds(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetSkillGroupNames(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.SkillGroupNames = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) SetTimestamp(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList {
	s.Timestamp = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportList) Validate() error {
	return dara.Validate(s)
}

type ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound struct {
	// example:
	//
	// 0
	AverageRingTime *int64 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// example:
	//
	// 0
	CallsOffered *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	// example:
	//
	// 0
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
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

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetAverageRingTime() *int64 {
	return s.AverageRingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetAverageRingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetAverageTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetAverageWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetCallsHandled(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetCallsOffered(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetHandleRate(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetMaxRingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetMaxTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetMaxWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetSatisfactionIndex(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetSatisfactionSurveysOffered(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetSatisfactionSurveysResponded(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetServiceLevel20(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetTotalRingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetTotalTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) SetTotalWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListInbound) Validate() error {
	return dara.Validate(s)
}

type ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound struct {
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 0
	AverageDialingTime *int64 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 0
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// example:
	//
	// 0
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 0
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// example:
	//
	// 0
	MaxWorkTime *string `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
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
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetAverageDialingTime() *int64 {
	return s.AverageDialingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetMaxWorkTime() *string {
	return s.MaxWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetAnswerRate(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetAverageDialingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetAverageTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetAverageWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetCallsAnswered(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetCallsDialed(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetMaxDialingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetMaxTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetMaxWorkTime(v string) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetSatisfactionIndex(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetSatisfactionSurveysOffered(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetSatisfactionSurveysResponded(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetTotalDialingTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetTotalTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) SetTotalWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOutbound) Validate() error {
	return dara.Validate(s)
}

type ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall struct {
	// example:
	//
	// 0
	AverageReadyTime *int64 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// example:
	//
	// 0
	AverageWorkTime *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// example:
	//
	// 37
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
	OneTransferCalls *int64 `json:"OneTransferCalls,omitempty" xml:"OneTransferCalls,omitempty"`
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
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

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetAverageReadyTime() *int64 {
	return s.AverageReadyTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetOneTransferCalls() *int64 {
	return s.OneTransferCalls
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetAverageReadyTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetAverageTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetAverageWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetMaxReadyTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetMaxTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetMaxWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetOccupancyRate(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetOneTransferCalls(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.OneTransferCalls = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetSatisfactionIndex(v float32) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetSatisfactionSurveysOffered(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetSatisfactionSurveysResponded(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalBreakTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalCalls(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalLoggedInTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalReadyTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalTalkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) SetTotalWorkTime(v int64) *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponseBodyPagedAgentSummaryReportListOverall) Validate() error {
	return dara.Validate(s)
}
