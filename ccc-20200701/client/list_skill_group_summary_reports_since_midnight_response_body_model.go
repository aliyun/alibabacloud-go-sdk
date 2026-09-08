// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupSummaryReportsSinceMidnightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetMessage() *string
	SetPagedSkillGroupSummaryReport(v *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetPagedSkillGroupSummaryReport() *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport
	SetRequestId(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSkillGroupSummaryReportsSinceMidnightResponseBody
	GetSuccess() *bool
}

type ListSkillGroupSummaryReportsSinceMidnightResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	PagedSkillGroupSummaryReport *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport `json:"PagedSkillGroupSummaryReport,omitempty" xml:"PagedSkillGroupSummaryReport,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2B36CEBC-6D11-5451-9E6B-C6D1927841C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetPagedSkillGroupSummaryReport() *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport {
	return s.PagedSkillGroupSummaryReport
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetCode(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetHttpStatusCode(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetMessage(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetPagedSkillGroupSummaryReport(v *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.PagedSkillGroupSummaryReport = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetRequestId(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) SetSuccess(v bool) *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	s.Success = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBody) Validate() error {
	if s.PagedSkillGroupSummaryReport != nil {
		if err := s.PagedSkillGroupSummaryReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport struct {
	// A list of skill group summary reports.
	List []*ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number. The valid range is 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The valid range is 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) GetList() []*ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	return s.List
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) SetList(v []*ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport {
	s.List = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) SetPageNumber(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) SetPageSize(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) SetTotalCount(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport {
	s.TotalCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReport) Validate() error {
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

type ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList struct {
	// Inbound call metrics.
	Inbound *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Outbound call metrics.
	Outbound *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	// Overall metrics.
	Overall *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	// The skill group ID.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The name of the skill group.
	//
	// example:
	//
	// 测试技能组
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	// The timestamp when the metrics were collected.
	//
	// example:
	//
	// 2018-09-13 00:00:00
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetInbound() *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	return s.Inbound
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetOutbound() *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	return s.Outbound
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetOverall() *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	return s.Overall
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetInbound(v *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.Inbound = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetInstanceId(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetOutbound(v *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.Outbound = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetOverall(v *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.Overall = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetSkillGroupId(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetSkillGroupName(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) SetTimestamp(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList {
	s.Timestamp = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportList) Validate() error {
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

type ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound struct {
	// The number of calls abandoned while waiting in the skill group queue.
	//
	// example:
	//
	// 0
	AbandonedInQueueOfQueueCount *int64 `json:"AbandonedInQueueOfQueueCount,omitempty" xml:"AbandonedInQueueOfQueueCount,omitempty"`
	// The number of calls routed to the skill group and answered by an agent.
	//
	// example:
	//
	// 0
	AnsweredByAgentOfQueueCount *int64 `json:"AnsweredByAgentOfQueueCount,omitempty" xml:"AnsweredByAgentOfQueueCount,omitempty"`
	// The maximum queue wait time for a call answered by an agent.
	//
	// example:
	//
	// 0
	AnsweredByAgentOfQueueMaxWaitTimeDuration *int64 `json:"AnsweredByAgentOfQueueMaxWaitTimeDuration,omitempty" xml:"AnsweredByAgentOfQueueMaxWaitTimeDuration,omitempty"`
	// The total queue wait time for calls answered by an agent, in seconds.
	//
	// example:
	//
	// 0
	AnsweredByAgentOfQueueWaitTimeDuration *int64 `json:"AnsweredByAgentOfQueueWaitTimeDuration,omitempty" xml:"AnsweredByAgentOfQueueWaitTimeDuration,omitempty"`
	// The average ring time, in seconds.
	//
	// example:
	//
	// 0
	AverageRingTime *int64 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	// The average talk time, in seconds.
	//
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// The average after-call work time, in seconds (`TotalWorkTime` / `CallsHandled`).
	//
	// example:
	//
	// 0
	AverageWorkTime *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAbandoned  *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	// The number of attended transfers initiated. If a call is transferred multiple times, each transfer is counted separately.
	//
	// example:
	//
	// 0
	CallsAttendedTransferOut *int64 `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	// The number of blind transfers initiated. If a call is transferred multiple times, each transfer is counted separately.
	//
	// example:
	//
	// 0
	CallsBlindTransferOut *int64 `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	// The number of handled inbound calls. Each call is counted once per queue entry, regardless of how many agents answer it.
	//
	// example:
	//
	// 0
	CallsHandled *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	// The number of calls offered to the skill group.
	//
	// example:
	//
	// 0
	CallsOffered         *int64  `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsOverflow        *string `json:"CallsOverflow,omitempty" xml:"CallsOverflow,omitempty"`
	CallsQueuingCanceled *string `json:"CallsQueuingCanceled,omitempty" xml:"CallsQueuingCanceled,omitempty"`
	CallsQueuingFailure  *string `json:"CallsQueuingFailure,omitempty" xml:"CallsQueuingFailure,omitempty"`
	CallsQueuingRerouted *string `json:"CallsQueuingRerouted,omitempty" xml:"CallsQueuingRerouted,omitempty"`
	CallsQueuingTimeout  *int64  `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	// The number of calls answered by an agent within 10 seconds.
	//
	// example:
	//
	// 0
	CallsServiceLevel10 *int64 `json:"CallsServiceLevel10,omitempty" xml:"CallsServiceLevel10,omitempty"`
	// The number of calls answered by an agent within 20 seconds.
	//
	// example:
	//
	// 0
	CallsServiceLevel20 *int64 `json:"CallsServiceLevel20,omitempty" xml:"CallsServiceLevel20,omitempty"`
	// The number of calls answered by an agent within 30 seconds.
	//
	// example:
	//
	// 0
	CallsServiceLevel30 *int64 `json:"CallsServiceLevel30,omitempty" xml:"CallsServiceLevel30,omitempty"`
	CallsTimeout        *int64 `json:"CallsTimeout,omitempty" xml:"CallsTimeout,omitempty"`
	// The number of calls that were abandoned while an agent in the skill group was being alerted.
	//
	// example:
	//
	// 0
	GiveUpByAgentOfQueueCount *int64 `json:"GiveUpByAgentOfQueueCount,omitempty" xml:"GiveUpByAgentOfQueueCount,omitempty"`
	// The handle rate (`CallsHandled` / `CallsOffered`). The value can exceed 100% because the handle and offer events might occur in different time windows.
	//
	// example:
	//
	// 0
	HandleRate *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	// The number of calls that entered the skill group queue.
	//
	// example:
	//
	// 0
	InComingQueueOfQueueCount *int64 `json:"InComingQueueOfQueueCount,omitempty" xml:"InComingQueueOfQueueCount,omitempty"`
	// The maximum ring time, in seconds.
	//
	// example:
	//
	// 0
	MaxRingTime *int64 `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	// The maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *string `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// The maximum after-call work time, in seconds.
	//
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// The number of calls that overflowed from the skill group queue.
	//
	// example:
	//
	// 0
	OverFlowInQueueOfQueueCount *int64 `json:"OverFlowInQueueOfQueueCount,omitempty" xml:"OverFlowInQueueOfQueueCount,omitempty"`
	// The maximum queue wait time, in seconds.
	//
	// example:
	//
	// 0
	QueueMaxWaitTimeDuration *int64 `json:"QueueMaxWaitTimeDuration,omitempty" xml:"QueueMaxWaitTimeDuration,omitempty"`
	// The queue wait time, in seconds.
	//
	// example:
	//
	// 0
	QueueWaitTimeDuration *int64 `json:"QueueWaitTimeDuration,omitempty" xml:"QueueWaitTimeDuration,omitempty"`
	// The satisfaction index, calculated as the average of single-digit keypad inputs from satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// The number of satisfaction surveys offered.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// The number of satisfaction surveys responded to.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// The percentage of calls answered within 20 seconds.
	//
	// example:
	//
	// 0
	ServiceLevel20 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	// The total ring time, in seconds.
	//
	// example:
	//
	// 0
	TotalRingTime *int64 `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	// The total talk time for inbound calls, in seconds.
	//
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// The total after-call work time, in seconds.
	//
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAbandonedInQueueOfQueueCount() *int64 {
	return s.AbandonedInQueueOfQueueCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAnsweredByAgentOfQueueCount() *int64 {
	return s.AnsweredByAgentOfQueueCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAnsweredByAgentOfQueueMaxWaitTimeDuration() *int64 {
	return s.AnsweredByAgentOfQueueMaxWaitTimeDuration
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAnsweredByAgentOfQueueWaitTimeDuration() *int64 {
	return s.AnsweredByAgentOfQueueWaitTimeDuration
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAverageRingTime() *int64 {
	return s.AverageRingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsAttendedTransferOut() *int64 {
	return s.CallsAttendedTransferOut
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsBlindTransferOut() *int64 {
	return s.CallsBlindTransferOut
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsHandled() *int64 {
	return s.CallsHandled
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsOverflow() *string {
	return s.CallsOverflow
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsQueuingCanceled() *string {
	return s.CallsQueuingCanceled
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsQueuingFailure() *string {
	return s.CallsQueuingFailure
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsQueuingRerouted() *string {
	return s.CallsQueuingRerouted
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsServiceLevel10() *int64 {
	return s.CallsServiceLevel10
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsServiceLevel20() *int64 {
	return s.CallsServiceLevel20
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsServiceLevel30() *int64 {
	return s.CallsServiceLevel30
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetCallsTimeout() *int64 {
	return s.CallsTimeout
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetGiveUpByAgentOfQueueCount() *int64 {
	return s.GiveUpByAgentOfQueueCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetHandleRate() *float32 {
	return s.HandleRate
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetInComingQueueOfQueueCount() *int64 {
	return s.InComingQueueOfQueueCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetMaxRingTime() *int64 {
	return s.MaxRingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetMaxTalkTime() *string {
	return s.MaxTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetOverFlowInQueueOfQueueCount() *int64 {
	return s.OverFlowInQueueOfQueueCount
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetQueueMaxWaitTimeDuration() *int64 {
	return s.QueueMaxWaitTimeDuration
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetQueueWaitTimeDuration() *int64 {
	return s.QueueWaitTimeDuration
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetServiceLevel20() *float32 {
	return s.ServiceLevel20
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetTotalRingTime() *int64 {
	return s.TotalRingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAbandonedInQueueOfQueueCount(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AbandonedInQueueOfQueueCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAnsweredByAgentOfQueueCount(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AnsweredByAgentOfQueueCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAnsweredByAgentOfQueueMaxWaitTimeDuration(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AnsweredByAgentOfQueueMaxWaitTimeDuration = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAnsweredByAgentOfQueueWaitTimeDuration(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AnsweredByAgentOfQueueWaitTimeDuration = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAverageRingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAverageTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetAverageWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsAbandoned(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsAttendedTransferOut(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsBlindTransferOut(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsHandled(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsOffered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsOverflow(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsOverflow = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsQueuingCanceled(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsQueuingCanceled = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsQueuingFailure(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsQueuingFailure = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsQueuingRerouted(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsQueuingRerouted = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsQueuingTimeout(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsServiceLevel10(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsServiceLevel10 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsServiceLevel20(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsServiceLevel20 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsServiceLevel30(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsServiceLevel30 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetCallsTimeout(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.CallsTimeout = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetGiveUpByAgentOfQueueCount(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.GiveUpByAgentOfQueueCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetHandleRate(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetInComingQueueOfQueueCount(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.InComingQueueOfQueueCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetMaxRingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetMaxTalkTime(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetMaxWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetOverFlowInQueueOfQueueCount(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.OverFlowInQueueOfQueueCount = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetQueueMaxWaitTimeDuration(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.QueueMaxWaitTimeDuration = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetQueueWaitTimeDuration(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.QueueWaitTimeDuration = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetSatisfactionIndex(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetSatisfactionSurveysOffered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetSatisfactionSurveysResponded(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetServiceLevel20(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetTotalRingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetTotalTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) SetTotalWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListInbound) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound struct {
	// The answer rate, expressed as a percentage (%).
	//
	// example:
	//
	// 0
	AnswerRate *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// The average dialing time, in seconds.
	//
	// example:
	//
	// 0
	AverageDialingTime *int64 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	// The average talk time, in seconds.
	//
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// The average after-call work time, in seconds.
	//
	// example:
	//
	// 0
	AverageWorkTime   *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAbandoned    *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	CallsAgentHandled *int64 `json:"CallsAgentHandled,omitempty" xml:"CallsAgentHandled,omitempty"`
	// The number of answered outbound calls.
	//
	// example:
	//
	// 0
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	// The number of outbound calls dialed.
	//
	// example:
	//
	// 0
	CallsDialed           *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsOffered          *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsQueuingCancelled *int64 `json:"CallsQueuingCancelled,omitempty" xml:"CallsQueuingCancelled,omitempty"`
	CallsQueuingFailed    *int64 `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	CallsQueuingFailure   *int64 `json:"CallsQueuingFailure,omitempty" xml:"CallsQueuingFailure,omitempty"`
	CallsQueuingOverflow  *int64 `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	CallsQueuingRerouted  *int64 `json:"CallsQueuingRerouted,omitempty" xml:"CallsQueuingRerouted,omitempty"`
	CallsQueuingTimeout   *int64 `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	// The number of calls answered by an agent within 30 seconds.
	//
	// example:
	//
	// 0
	CallsServiceLevel30   *string `json:"CallsServiceLevel30,omitempty" xml:"CallsServiceLevel30,omitempty"`
	CallsServiceLevel30V2 *int64  `json:"CallsServiceLevel30V2,omitempty" xml:"CallsServiceLevel30V2,omitempty"`
	// The maximum dialing time, in seconds.
	//
	// example:
	//
	// 0
	MaxDialingTime *int64 `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	// The maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// The maximum after-call work time, in seconds.
	//
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// The satisfaction index, calculated as the average of single-digit keypad inputs from satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// The number of satisfaction surveys offered.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// The number of satisfaction surveys responded to.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// The total dialing time, in seconds.
	//
	// example:
	//
	// 0
	TotalDialingTime *int64 `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	// The total talk time for outbound calls, in seconds.
	//
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWaitTime *int64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	// The total after-call work time, in seconds.
	//
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetAverageDialingTime() *int64 {
	return s.AverageDialingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsAgentHandled() *int64 {
	return s.CallsAgentHandled
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsAnswered() *int64 {
	return s.CallsAnswered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsOffered() *int64 {
	return s.CallsOffered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingCancelled() *int64 {
	return s.CallsQueuingCancelled
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingFailed() *int64 {
	return s.CallsQueuingFailed
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingFailure() *int64 {
	return s.CallsQueuingFailure
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingOverflow() *int64 {
	return s.CallsQueuingOverflow
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingRerouted() *int64 {
	return s.CallsQueuingRerouted
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsQueuingTimeout() *int64 {
	return s.CallsQueuingTimeout
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsServiceLevel30() *string {
	return s.CallsServiceLevel30
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetCallsServiceLevel30V2() *int64 {
	return s.CallsServiceLevel30V2
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetMaxDialingTime() *int64 {
	return s.MaxDialingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetTotalDialingTime() *int64 {
	return s.TotalDialingTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetTotalWaitTime() *int64 {
	return s.TotalWaitTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetAnswerRate(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetAverageDialingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetAverageTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetAverageWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsAbandoned(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsAgentHandled(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsAgentHandled = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsAnswered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsDialed(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsOffered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsOffered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingCancelled(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingCancelled = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingFailed(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingFailure(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingFailure = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingOverflow(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingRerouted(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingRerouted = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsQueuingTimeout(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsServiceLevel30(v string) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsServiceLevel30 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetCallsServiceLevel30V2(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.CallsServiceLevel30V2 = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetMaxDialingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetMaxTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetMaxWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetSatisfactionIndex(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetSatisfactionSurveysOffered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetSatisfactionSurveysResponded(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetTotalDialingTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetTotalTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetTotalWaitTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) SetTotalWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOutbound) Validate() error {
	return dara.Validate(s)
}

type ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall struct {
	// The average ready time, in seconds.
	//
	// example:
	//
	// 0
	AverageReadyTime *int64 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	// The average talk time, in seconds.
	//
	// example:
	//
	// 0
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	// The average after-call work time, in seconds.
	//
	// example:
	//
	// 0
	AverageWorkTime *int64 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	// The maximum ready time, in seconds.
	//
	// example:
	//
	// 0
	MaxReadyTime *int64 `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	// The maximum talk time, in seconds.
	//
	// example:
	//
	// 0
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// The maximum after-call work time, in seconds.
	//
	// example:
	//
	// 0
	MaxWorkTime *int64 `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	// The agent occupancy rate, expressed as a percentage (%).
	//
	// example:
	//
	// 0
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	// The satisfaction index, calculated as the average of single-digit keypad inputs from satisfaction surveys.
	//
	// example:
	//
	// 0
	SatisfactionIndex *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	// The number of satisfaction surveys offered.
	//
	// example:
	//
	// 0
	SatisfactionSurveysOffered *int64 `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	// The number of satisfaction surveys responded to.
	//
	// example:
	//
	// 0
	SatisfactionSurveysResponded *int64 `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	// The total break time, in seconds.
	//
	// example:
	//
	// 0
	TotalBreakTime *int64 `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	// The total number of calls (`CallsOffered` + `CallsDialed`).
	//
	// example:
	//
	// 0
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	// The total logged-in time, in seconds.
	//
	// example:
	//
	// 0
	TotalLoggedInTime *int64 `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	// The total ready time, in seconds.
	//
	// example:
	//
	// 0
	TotalReadyTime *int64 `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	// The total talk time, in seconds.
	//
	// example:
	//
	// 0
	TotalTalkTime *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	// The total after-call work time, in seconds.
	//
	// example:
	//
	// 0
	TotalWorkTime *int64 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetAverageReadyTime() *int64 {
	return s.AverageReadyTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetAverageTalkTime() *int64 {
	return s.AverageTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetAverageWorkTime() *int64 {
	return s.AverageWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetMaxReadyTime() *int64 {
	return s.MaxReadyTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetMaxWorkTime() *int64 {
	return s.MaxWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetSatisfactionIndex() *float32 {
	return s.SatisfactionIndex
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetSatisfactionSurveysOffered() *int64 {
	return s.SatisfactionSurveysOffered
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetSatisfactionSurveysResponded() *int64 {
	return s.SatisfactionSurveysResponded
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalBreakTime() *int64 {
	return s.TotalBreakTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalLoggedInTime() *int64 {
	return s.TotalLoggedInTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalReadyTime() *int64 {
	return s.TotalReadyTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalTalkTime() *int64 {
	return s.TotalTalkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) GetTotalWorkTime() *int64 {
	return s.TotalWorkTime
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetAverageReadyTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetAverageTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetAverageWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetMaxReadyTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetMaxTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetMaxWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetOccupancyRate(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetSatisfactionIndex(v float32) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetSatisfactionSurveysOffered(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetSatisfactionSurveysResponded(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalBreakTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalCalls(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalLoggedInTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalReadyTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalTalkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) SetTotalWorkTime(v int64) *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponseBodyPagedSkillGroupSummaryReportListOverall) Validate() error {
	return dara.Validate(s)
}
