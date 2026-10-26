// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompletion(v string) *GetInstanceInspectionReportResponseBody
	GetCompletion() *string
	SetCoverage(v []*GetInstanceInspectionReportResponseBodyCoverage) *GetInstanceInspectionReportResponseBody
	GetCoverage() []*GetInstanceInspectionReportResponseBodyCoverage
	SetFindings(v []*GetInstanceInspectionReportResponseBodyFindings) *GetInstanceInspectionReportResponseBody
	GetFindings() []*GetInstanceInspectionReportResponseBodyFindings
	SetFinishedAt(v string) *GetInstanceInspectionReportResponseBody
	GetFinishedAt() *string
	SetHtml(v string) *GetInstanceInspectionReportResponseBody
	GetHtml() *string
	SetInspectionRunId(v string) *GetInstanceInspectionReportResponseBody
	GetInspectionRunId() *string
	SetInstanceId(v string) *GetInstanceInspectionReportResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *GetInstanceInspectionReportResponseBody
	GetInstanceName() *string
	SetMarkdown(v string) *GetInstanceInspectionReportResponseBody
	GetMarkdown() *string
	SetOverview(v string) *GetInstanceInspectionReportResponseBody
	GetOverview() *string
	SetRegionId(v string) *GetInstanceInspectionReportResponseBody
	GetRegionId() *string
	SetReportId(v string) *GetInstanceInspectionReportResponseBody
	GetReportId() *string
	SetReportSummary(v *GetInstanceInspectionReportResponseBodyReportSummary) *GetInstanceInspectionReportResponseBody
	GetReportSummary() *GetInstanceInspectionReportResponseBodyReportSummary
	SetTemplateVersion(v string) *GetInstanceInspectionReportResponseBody
	GetTemplateVersion() *string
	SetTimeRange(v *GetInstanceInspectionReportResponseBodyTimeRange) *GetInstanceInspectionReportResponseBody
	GetTimeRange() *GetInstanceInspectionReportResponseBodyTimeRange
}

type GetInstanceInspectionReportResponseBody struct {
	// example:
	//
	// complete
	Completion *string `json:"completion,omitempty" xml:"completion,omitempty"`
	// example:
	//
	// []
	Coverage []*GetInstanceInspectionReportResponseBodyCoverage `json:"coverage,omitempty" xml:"coverage,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Findings []*GetInstanceInspectionReportResponseBodyFindings `json:"findings,omitempty" xml:"findings,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-18T08:02:00Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// <h1>实例巡检诊断报告</h1>
	Html *string `json:"html,omitempty" xml:"html,omitempty"`
	// example:
	//
	// irun-example
	InspectionRunId *string `json:"inspectionRunId,omitempty" xml:"inspectionRunId,omitempty"`
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
	// # 实例巡检诊断报告
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// example:
	//
	// 消费延迟与容量限制可能相关
	Overview *string `json:"overview,omitempty" xml:"overview,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// irpt-example
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// example:
	//
	// {}
	ReportSummary *GetInstanceInspectionReportResponseBodyReportSummary `json:"reportSummary,omitempty" xml:"reportSummary,omitempty" type:"Struct"`
	// example:
	//
	// inspection-v2.1
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionReportResponseBodyTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBody) GetCompletion() *string {
	return s.Completion
}

func (s *GetInstanceInspectionReportResponseBody) GetCoverage() []*GetInstanceInspectionReportResponseBodyCoverage {
	return s.Coverage
}

func (s *GetInstanceInspectionReportResponseBody) GetFindings() []*GetInstanceInspectionReportResponseBodyFindings {
	return s.Findings
}

func (s *GetInstanceInspectionReportResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetInstanceInspectionReportResponseBody) GetHtml() *string {
	return s.Html
}

func (s *GetInstanceInspectionReportResponseBody) GetInspectionRunId() *string {
	return s.InspectionRunId
}

func (s *GetInstanceInspectionReportResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceInspectionReportResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceInspectionReportResponseBody) GetMarkdown() *string {
	return s.Markdown
}

func (s *GetInstanceInspectionReportResponseBody) GetOverview() *string {
	return s.Overview
}

func (s *GetInstanceInspectionReportResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceInspectionReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GetInstanceInspectionReportResponseBody) GetReportSummary() *GetInstanceInspectionReportResponseBodyReportSummary {
	return s.ReportSummary
}

func (s *GetInstanceInspectionReportResponseBody) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetInstanceInspectionReportResponseBody) GetTimeRange() *GetInstanceInspectionReportResponseBodyTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionReportResponseBody) SetCompletion(v string) *GetInstanceInspectionReportResponseBody {
	s.Completion = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetCoverage(v []*GetInstanceInspectionReportResponseBodyCoverage) *GetInstanceInspectionReportResponseBody {
	s.Coverage = v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetFindings(v []*GetInstanceInspectionReportResponseBodyFindings) *GetInstanceInspectionReportResponseBody {
	s.Findings = v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetFinishedAt(v string) *GetInstanceInspectionReportResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetHtml(v string) *GetInstanceInspectionReportResponseBody {
	s.Html = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetInspectionRunId(v string) *GetInstanceInspectionReportResponseBody {
	s.InspectionRunId = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetInstanceId(v string) *GetInstanceInspectionReportResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetInstanceName(v string) *GetInstanceInspectionReportResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetMarkdown(v string) *GetInstanceInspectionReportResponseBody {
	s.Markdown = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetOverview(v string) *GetInstanceInspectionReportResponseBody {
	s.Overview = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetRegionId(v string) *GetInstanceInspectionReportResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetReportId(v string) *GetInstanceInspectionReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetReportSummary(v *GetInstanceInspectionReportResponseBodyReportSummary) *GetInstanceInspectionReportResponseBody {
	s.ReportSummary = v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetTemplateVersion(v string) *GetInstanceInspectionReportResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) SetTimeRange(v *GetInstanceInspectionReportResponseBodyTimeRange) *GetInstanceInspectionReportResponseBody {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionReportResponseBody) Validate() error {
	if s.Coverage != nil {
		for _, item := range s.Coverage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Findings != nil {
		for _, item := range s.Findings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReportSummary != nil {
		if err := s.ReportSummary.Validate(); err != nil {
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

type GetInstanceInspectionReportResponseBodyCoverage struct {
	// example:
	//
	// Succeeded
	DiagnosisStatus *string `json:"diagnosisStatus,omitempty" xml:"diagnosisStatus,omitempty"`
	// example:
	//
	// 消费健康
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2026-09-01T01:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// example:
	//
	// risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// {"startTime":"2026-09-01T00:00:00Z","endTime":"2026-09-01T01:00:00Z"}
	TimeRange *GetInstanceInspectionReportResponseBodyCoverageTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
}

func (s GetInstanceInspectionReportResponseBodyCoverage) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyCoverage) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetDiagnosisStatus() *string {
	return s.DiagnosisStatus
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetResult() *string {
	return s.Result
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) GetTimeRange() *GetInstanceInspectionReportResponseBodyCoverageTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetDiagnosisStatus(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.DiagnosisStatus = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetName(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetObservedAt(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetResult(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.Result = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetStatus(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.Status = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetTemplateCode(v string) *GetInstanceInspectionReportResponseBodyCoverage {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) SetTimeRange(v *GetInstanceInspectionReportResponseBodyCoverageTimeRange) *GetInstanceInspectionReportResponseBodyCoverage {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverage) Validate() error {
	if s.TimeRange != nil {
		if err := s.TimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceInspectionReportResponseBodyCoverageTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionReportResponseBodyCoverageTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyCoverageTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyCoverageTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionReportResponseBodyCoverageTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionReportResponseBodyCoverageTimeRange) SetEndTime(v string) *GetInstanceInspectionReportResponseBodyCoverageTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverageTimeRange) SetStartTime(v string) *GetInstanceInspectionReportResponseBodyCoverageTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyCoverageTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionReportResponseBodyFindings struct {
	// example:
	//
	// []
	Alternatives []*string `json:"alternatives,omitempty" xml:"alternatives,omitempty" type:"Repeated"`
	// example:
	//
	// 消费环节
	Boundary *string `json:"boundary,omitempty" xml:"boundary,omitempty"`
	// example:
	//
	// 消费者处理速度可能不足
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// likely
	Confidence *string `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// example:
	//
	// ["消费延迟 60 秒"]
	Evidence []*string `json:"evidence,omitempty" xml:"evidence,omitempty" type:"Repeated"`
	// example:
	//
	// 消息处理延迟升高
	Impact *string `json:"impact,omitempty" xml:"impact,omitempty"`
	// example:
	//
	// 消费组监控
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// []
	MissingEvidence []*string `json:"missingEvidence,omitempty" xml:"missingEvidence,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Observations []*string `json:"observations,omitempty" xml:"observations,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-01T01:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
	// example:
	//
	// ["检查消费者处理耗时"]
	Recommendations []*string `json:"recommendations,omitempty" xml:"recommendations,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Resource *GetInstanceInspectionReportResponseBodyFindingsResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	// example:
	//
	// 消费延迟
	RiskType *string `json:"riskType,omitempty" xml:"riskType,omitempty"`
	// example:
	//
	// script
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// example:
	//
	// {}
	TimeRange *GetInstanceInspectionReportResponseBodyFindingsTimeRange `json:"timeRange,omitempty" xml:"timeRange,omitempty" type:"Struct"`
	// example:
	//
	// []
	Verification []*string `json:"verification,omitempty" xml:"verification,omitempty" type:"Repeated"`
}

func (s GetInstanceInspectionReportResponseBodyFindings) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyFindings) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetAlternatives() []*string {
	return s.Alternatives
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetBoundary() *string {
	return s.Boundary
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetCause() *string {
	return s.Cause
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetConfidence() *string {
	return s.Confidence
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetEvidence() []*string {
	return s.Evidence
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetImpact() *string {
	return s.Impact
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetLocation() *string {
	return s.Location
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetMissingEvidence() []*string {
	return s.MissingEvidence
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetObservations() []*string {
	return s.Observations
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetRecommendations() []*string {
	return s.Recommendations
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetResource() *GetInstanceInspectionReportResponseBodyFindingsResource {
	return s.Resource
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetRiskType() *string {
	return s.RiskType
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetSource() *string {
	return s.Source
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetTimeRange() *GetInstanceInspectionReportResponseBodyFindingsTimeRange {
	return s.TimeRange
}

func (s *GetInstanceInspectionReportResponseBodyFindings) GetVerification() []*string {
	return s.Verification
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetAlternatives(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Alternatives = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetBoundary(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Boundary = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetCause(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Cause = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetConfidence(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Confidence = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetEvidence(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Evidence = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetImpact(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Impact = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetLocation(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Location = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetMissingEvidence(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.MissingEvidence = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetObservations(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Observations = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetObservedAt(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.ObservedAt = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetRecommendations(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Recommendations = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetResource(v *GetInstanceInspectionReportResponseBodyFindingsResource) *GetInstanceInspectionReportResponseBodyFindings {
	s.Resource = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetRiskType(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.RiskType = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetSource(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Source = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetTemplateCode(v string) *GetInstanceInspectionReportResponseBodyFindings {
	s.TemplateCode = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetTimeRange(v *GetInstanceInspectionReportResponseBodyFindingsTimeRange) *GetInstanceInspectionReportResponseBodyFindings {
	s.TimeRange = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) SetVerification(v []*string) *GetInstanceInspectionReportResponseBodyFindings {
	s.Verification = v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindings) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
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

type GetInstanceInspectionReportResponseBodyFindingsResource struct {
	// example:
	//
	// GID-orders
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 订单消费组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// group
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceInspectionReportResponseBodyFindingsResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyFindingsResource) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) GetId() *string {
	return s.Id
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) GetName() *string {
	return s.Name
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) GetType() *string {
	return s.Type
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) SetId(v string) *GetInstanceInspectionReportResponseBodyFindingsResource {
	s.Id = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) SetName(v string) *GetInstanceInspectionReportResponseBodyFindingsResource {
	s.Name = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) SetType(v string) *GetInstanceInspectionReportResponseBodyFindingsResource {
	s.Type = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindingsResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionReportResponseBodyFindingsTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionReportResponseBodyFindingsTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyFindingsTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyFindingsTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionReportResponseBodyFindingsTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionReportResponseBodyFindingsTimeRange) SetEndTime(v string) *GetInstanceInspectionReportResponseBodyFindingsTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindingsTimeRange) SetStartTime(v string) *GetInstanceInspectionReportResponseBodyFindingsTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyFindingsTimeRange) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionReportResponseBodyReportSummary struct {
	// example:
	//
	// 4
	InfoCount *int32 `json:"infoCount,omitempty" xml:"infoCount,omitempty"`
	// example:
	//
	// 1
	RiskCount *int32 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s GetInstanceInspectionReportResponseBodyReportSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyReportSummary) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyReportSummary) GetInfoCount() *int32 {
	return s.InfoCount
}

func (s *GetInstanceInspectionReportResponseBodyReportSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *GetInstanceInspectionReportResponseBodyReportSummary) SetInfoCount(v int32) *GetInstanceInspectionReportResponseBodyReportSummary {
	s.InfoCount = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyReportSummary) SetRiskCount(v int32) *GetInstanceInspectionReportResponseBodyReportSummary {
	s.RiskCount = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyReportSummary) Validate() error {
	return dara.Validate(s)
}

type GetInstanceInspectionReportResponseBodyTimeRange struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetInstanceInspectionReportResponseBodyTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionReportResponseBodyTimeRange) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionReportResponseBodyTimeRange) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionReportResponseBodyTimeRange) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionReportResponseBodyTimeRange) SetEndTime(v string) *GetInstanceInspectionReportResponseBodyTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyTimeRange) SetStartTime(v string) *GetInstanceInspectionReportResponseBodyTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionReportResponseBodyTimeRange) Validate() error {
	return dara.Validate(s)
}
