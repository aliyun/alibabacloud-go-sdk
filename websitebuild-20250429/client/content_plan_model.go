// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentPlan interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalRequirements(v string) *ContentPlan
	GetAdditionalRequirements() *string
	SetAnalysisTaskId(v string) *ContentPlan
	GetAnalysisTaskId() *string
	SetConfigRevision(v int64) *ContentPlan
	GetConfigRevision() *int64
	SetDiagnosisCompletedAt(v int64) *ContentPlan
	GetDiagnosisCompletedAt() *int64
	SetDiagnosisTaskId(v string) *ContentPlan
	GetDiagnosisTaskId() *string
	SetDiagnosisType(v string) *ContentPlan
	GetDiagnosisType() *string
	SetEndDate(v string) *ContentPlan
	GetEndDate() *string
	SetGenerationTaskId(v string) *ContentPlan
	GetGenerationTaskId() *string
	SetPlanId(v string) *ContentPlan
	GetPlanId() *string
	SetProjectId(v string) *ContentPlan
	GetProjectId() *string
	SetPublishFrequency(v int32) *ContentPlan
	GetPublishFrequency() *int32
	SetReportConfigRevision(v int64) *ContentPlan
	GetReportConfigRevision() *int64
	SetStartDate(v string) *ContentPlan
	GetStartDate() *string
	SetStatus(v string) *ContentPlan
	GetStatus() *string
	SetSummary(v string) *ContentPlan
	GetSummary() *string
	SetSupersededBy(v string) *ContentPlan
	GetSupersededBy() *string
	SetTimezone(v string) *ContentPlan
	GetTimezone() *string
	SetWeekdays(v []*int32) *ContentPlan
	GetWeekdays() []*int32
}

type ContentPlan struct {
	// The additional user comments.
	//
	// example:
	//
	// Highlight the differentiated capabilities of the product and avoid exaggerated claims
	AdditionalRequirements *string `json:"AdditionalRequirements,omitempty" xml:"AdditionalRequirements,omitempty"`
	// The ID of the analysis task used for scheduling.
	//
	// example:
	//
	// cpt_1a2b3c4d5e6f708192a3b4c5d6e7f809
	AnalysisTaskId *string `json:"AnalysisTaskId,omitempty" xml:"AnalysisTaskId,omitempty"`
	// The current configuration version of the workspace when the analysis is initiated.
	//
	// example:
	//
	// 12
	ConfigRevision *int64 `json:"ConfigRevision,omitempty" xml:"ConfigRevision,omitempty"`
	// The completion time of the GEO diagnosis, in milliseconds.
	//
	// example:
	//
	// 1735660800000
	DiagnosisCompletedAt *int64 `json:"DiagnosisCompletedAt,omitempty" xml:"DiagnosisCompletedAt,omitempty"`
	// The ID of the GEO diagnostic node.
	//
	// example:
	//
	// 1f2e3d4c-5b6a-4798-8a9b-c0d1e2f3a4b5
	DiagnosisTaskId *string `json:"DiagnosisTaskId,omitempty" xml:"DiagnosisTaskId,omitempty"`
	// The type of the GEO diagnosis.
	//
	// example:
	//
	// GEO_BRAND
	DiagnosisType *string `json:"DiagnosisType,omitempty" xml:"DiagnosisType,omitempty"`
	// The end date of the plan in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-10-18
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the generation task.
	//
	// example:
	//
	// cpt_9b7d4a5f8c0e1d2b3a4c5d6e3f7a1c2e
	GenerationTaskId *string `json:"GenerationTaskId,omitempty" xml:"GenerationTaskId,omitempty"`
	// The ID of the content plan.
	//
	// example:
	//
	// plan_3f7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the workspace, which is the GEO workspace for brand marketing and corresponds to the appWorkspaceId.
	//
	// example:
	//
	// d4e5f60718293a4b5c6d7e8f90a1b2c3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The number of scheduled articles, which equals the number of matched generation days.
	//
	// example:
	//
	// 3
	PublishFrequency *int32 `json:"PublishFrequency,omitempty" xml:"PublishFrequency,omitempty"`
	// The workspace configuration version used for the diagnostic report, which is strictly read during article generation.
	//
	// example:
	//
	// 12
	ReportConfigRevision *int64 `json:"ReportConfigRevision,omitempty" xml:"ReportConfigRevision,omitempty"`
	// The start date of the plan in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-10-12
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the plan. Valid values: ACTIVE, SUPERSEDED, COMPLETED, and FAILED.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The summary of the overall strategy for the plan.
	//
	// example:
	//
	// This week focuses on two themes: brand awareness and product capabilities, with a total of three articles
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The ID of the new plan that supersedes the current plan.
	//
	// example:
	//
	// plan_7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e3f
	SupersededBy *string `json:"SupersededBy,omitempty" xml:"SupersededBy,omitempty"`
	// The IANA time zone used for scheduling.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The weekly generation days, where ISO weekday 1 represents Monday and 7 represents Sunday.
	Weekdays []*int32 `json:"Weekdays,omitempty" xml:"Weekdays,omitempty" type:"Repeated"`
}

func (s ContentPlan) String() string {
	return dara.Prettify(s)
}

func (s ContentPlan) GoString() string {
	return s.String()
}

func (s *ContentPlan) GetAdditionalRequirements() *string {
	return s.AdditionalRequirements
}

func (s *ContentPlan) GetAnalysisTaskId() *string {
	return s.AnalysisTaskId
}

func (s *ContentPlan) GetConfigRevision() *int64 {
	return s.ConfigRevision
}

func (s *ContentPlan) GetDiagnosisCompletedAt() *int64 {
	return s.DiagnosisCompletedAt
}

func (s *ContentPlan) GetDiagnosisTaskId() *string {
	return s.DiagnosisTaskId
}

func (s *ContentPlan) GetDiagnosisType() *string {
	return s.DiagnosisType
}

func (s *ContentPlan) GetEndDate() *string {
	return s.EndDate
}

func (s *ContentPlan) GetGenerationTaskId() *string {
	return s.GenerationTaskId
}

func (s *ContentPlan) GetPlanId() *string {
	return s.PlanId
}

func (s *ContentPlan) GetProjectId() *string {
	return s.ProjectId
}

func (s *ContentPlan) GetPublishFrequency() *int32 {
	return s.PublishFrequency
}

func (s *ContentPlan) GetReportConfigRevision() *int64 {
	return s.ReportConfigRevision
}

func (s *ContentPlan) GetStartDate() *string {
	return s.StartDate
}

func (s *ContentPlan) GetStatus() *string {
	return s.Status
}

func (s *ContentPlan) GetSummary() *string {
	return s.Summary
}

func (s *ContentPlan) GetSupersededBy() *string {
	return s.SupersededBy
}

func (s *ContentPlan) GetTimezone() *string {
	return s.Timezone
}

func (s *ContentPlan) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *ContentPlan) SetAdditionalRequirements(v string) *ContentPlan {
	s.AdditionalRequirements = &v
	return s
}

func (s *ContentPlan) SetAnalysisTaskId(v string) *ContentPlan {
	s.AnalysisTaskId = &v
	return s
}

func (s *ContentPlan) SetConfigRevision(v int64) *ContentPlan {
	s.ConfigRevision = &v
	return s
}

func (s *ContentPlan) SetDiagnosisCompletedAt(v int64) *ContentPlan {
	s.DiagnosisCompletedAt = &v
	return s
}

func (s *ContentPlan) SetDiagnosisTaskId(v string) *ContentPlan {
	s.DiagnosisTaskId = &v
	return s
}

func (s *ContentPlan) SetDiagnosisType(v string) *ContentPlan {
	s.DiagnosisType = &v
	return s
}

func (s *ContentPlan) SetEndDate(v string) *ContentPlan {
	s.EndDate = &v
	return s
}

func (s *ContentPlan) SetGenerationTaskId(v string) *ContentPlan {
	s.GenerationTaskId = &v
	return s
}

func (s *ContentPlan) SetPlanId(v string) *ContentPlan {
	s.PlanId = &v
	return s
}

func (s *ContentPlan) SetProjectId(v string) *ContentPlan {
	s.ProjectId = &v
	return s
}

func (s *ContentPlan) SetPublishFrequency(v int32) *ContentPlan {
	s.PublishFrequency = &v
	return s
}

func (s *ContentPlan) SetReportConfigRevision(v int64) *ContentPlan {
	s.ReportConfigRevision = &v
	return s
}

func (s *ContentPlan) SetStartDate(v string) *ContentPlan {
	s.StartDate = &v
	return s
}

func (s *ContentPlan) SetStatus(v string) *ContentPlan {
	s.Status = &v
	return s
}

func (s *ContentPlan) SetSummary(v string) *ContentPlan {
	s.Summary = &v
	return s
}

func (s *ContentPlan) SetSupersededBy(v string) *ContentPlan {
	s.SupersededBy = &v
	return s
}

func (s *ContentPlan) SetTimezone(v string) *ContentPlan {
	s.Timezone = &v
	return s
}

func (s *ContentPlan) SetWeekdays(v []*int32) *ContentPlan {
	s.Weekdays = v
	return s
}

func (s *ContentPlan) Validate() error {
	return dara.Validate(s)
}
