// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInstanceInspectionRunsResponseBodyData) *ListInstanceInspectionRunsResponseBody
	GetData() []*ListInstanceInspectionRunsResponseBodyData
	SetMaxResults(v int32) *ListInstanceInspectionRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionRunsResponseBody
	GetNextToken() *string
	SetTotal(v int64) *ListInstanceInspectionRunsResponseBody
	GetTotal() *int64
}

type ListInstanceInspectionRunsResponseBody struct {
	// example:
	//
	// []
	Data []*ListInstanceInspectionRunsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBody) GetData() []*ListInstanceInspectionRunsResponseBodyData {
	return s.Data
}

func (s *ListInstanceInspectionRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionRunsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListInstanceInspectionRunsResponseBody) SetData(v []*ListInstanceInspectionRunsResponseBodyData) *ListInstanceInspectionRunsResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceInspectionRunsResponseBody) SetMaxResults(v int32) *ListInstanceInspectionRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBody) SetNextToken(v string) *ListInstanceInspectionRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBody) SetTotal(v int64) *ListInstanceInspectionRunsResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBody) Validate() error {
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

type ListInstanceInspectionRunsResponseBodyData struct {
	// 每次巡检结果由 LLM 总结；包含生成状态及成功后的原文，未生成时为空。
	//
	// example:
	//
	// {}
	AiSummary *ListInstanceInspectionRunsResponseBodyDataAiSummary `json:"aiSummary,omitempty" xml:"aiSummary,omitempty" type:"Struct"`
	// example:
	//
	// 2026-09-08T00:01:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// rmq-cn-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 订单实例
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// {}
	Notification *ListInstanceInspectionRunsResponseBodyDataNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 2026-10-15T10:00:00Z
	ReportFinishedAt *string `json:"reportFinishedAt,omitempty" xml:"reportFinishedAt,omitempty"`
	// example:
	//
	// irpt-one
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// Succeeded
	ReportStatus *string `json:"reportStatus,omitempty" xml:"reportStatus,omitempty"`
	// example:
	//
	// 2026-09-08T00:00:00Z
	RequestedAt *string `json:"requestedAt,omitempty" xml:"requestedAt,omitempty"`
	// example:
	//
	// inspection-run-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// {}
	Summary *ListInstanceInspectionRunsResponseBodyDataSummary `json:"summary,omitempty" xml:"summary,omitempty" type:"Struct"`
	// example:
	//
	// {}
	TimeRange *ListInstanceInspectionRunsResponseBodyDataTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// manual
	TriggerSource *string `json:"triggerSource,omitempty" xml:"triggerSource,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetAiSummary() *ListInstanceInspectionRunsResponseBodyDataAiSummary {
	return s.AiSummary
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetNotification() *ListInstanceInspectionRunsResponseBodyDataNotification {
	return s.Notification
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetReportFinishedAt() *string {
	return s.ReportFinishedAt
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetRequestedAt() *string {
	return s.RequestedAt
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetRunId() *string {
	return s.RunId
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetSummary() *ListInstanceInspectionRunsResponseBodyDataSummary {
	return s.Summary
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetTimeRange() *ListInstanceInspectionRunsResponseBodyDataTimeRange {
	return s.TimeRange
}

func (s *ListInstanceInspectionRunsResponseBodyData) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetAiSummary(v *ListInstanceInspectionRunsResponseBodyDataAiSummary) *ListInstanceInspectionRunsResponseBodyData {
	s.AiSummary = v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetFinishedAt(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.FinishedAt = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetInstanceId(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetInstanceName(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetNotification(v *ListInstanceInspectionRunsResponseBodyDataNotification) *ListInstanceInspectionRunsResponseBodyData {
	s.Notification = v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetRegionId(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetReportFinishedAt(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.ReportFinishedAt = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetReportId(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetReportStatus(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetRequestedAt(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.RequestedAt = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetRunId(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.RunId = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetStatus(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetSummary(v *ListInstanceInspectionRunsResponseBodyDataSummary) *ListInstanceInspectionRunsResponseBodyData {
	s.Summary = v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetTimeRange(v *ListInstanceInspectionRunsResponseBodyDataTimeRange) *ListInstanceInspectionRunsResponseBodyData {
	s.TimeRange = v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) SetTriggerSource(v string) *ListInstanceInspectionRunsResponseBodyData {
	s.TriggerSource = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyData) Validate() error {
	if s.AiSummary != nil {
		if err := s.AiSummary.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceInspectionRunsResponseBodyDataAiSummary struct {
	// Pending 等待、Running 生成中、Succeeded 已生成、Failed 生成失败；不改变巡检结果。
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 仅生成成功时返回，最多 200 字，不由前端拼接或截取。
	//
	// example:
	//
	// 本次巡检未发现风险，实例运行正常。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBodyDataAiSummary) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBodyDataAiSummary) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBodyDataAiSummary) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceInspectionRunsResponseBodyDataAiSummary) GetText() *string {
	return s.Text
}

func (s *ListInstanceInspectionRunsResponseBodyDataAiSummary) SetStatus(v string) *ListInstanceInspectionRunsResponseBodyDataAiSummary {
	s.Status = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataAiSummary) SetText(v string) *ListInstanceInspectionRunsResponseBodyDataAiSummary {
	s.Text = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataAiSummary) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionRunsResponseBodyDataNotification struct {
	// example:
	//
	// inspection-event-001
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Sent
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBodyDataNotification) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBodyDataNotification) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBodyDataNotification) GetEventId() *string {
	return s.EventId
}

func (s *ListInstanceInspectionRunsResponseBodyDataNotification) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceInspectionRunsResponseBodyDataNotification) SetEventId(v string) *ListInstanceInspectionRunsResponseBodyDataNotification {
	s.EventId = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataNotification) SetStatus(v string) *ListInstanceInspectionRunsResponseBodyDataNotification {
	s.Status = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataNotification) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionRunsResponseBodyDataSummary struct {
	// example:
	//
	// 2
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBodyDataSummary) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) GetResult() *string {
	return s.Result
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) SetInfoCount(v int32) *ListInstanceInspectionRunsResponseBodyDataSummary {
	s.InfoCount = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) SetResult(v string) *ListInstanceInspectionRunsResponseBodyDataSummary {
	s.Result = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) SetRiskCount(v int32) *ListInstanceInspectionRunsResponseBodyDataSummary {
	s.RiskCount = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataSummary) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionRunsResponseBodyDataTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListInstanceInspectionRunsResponseBodyDataTimeRange) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsResponseBodyDataTimeRange) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsResponseBodyDataTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInstanceInspectionRunsResponseBodyDataTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInstanceInspectionRunsResponseBodyDataTimeRange) SetEndTime(v string) *ListInstanceInspectionRunsResponseBodyDataTimeRange {
	s.EndTime = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataTimeRange) SetStartTime(v string) *ListInstanceInspectionRunsResponseBodyDataTimeRange {
	s.StartTime = &v
	return s
}

func (s *ListInstanceInspectionRunsResponseBodyDataTimeRange) Validate() error {
	return dara.Validate(s)
}
