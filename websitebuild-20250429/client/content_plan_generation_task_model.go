// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentPlanGenerationTask interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysis(v *ContentPlanGenerationTaskAnalysis) *ContentPlanGenerationTask
	GetAnalysis() *ContentPlanGenerationTaskAnalysis
	SetAnalysisTaskId(v string) *ContentPlanGenerationTask
	GetAnalysisTaskId() *string
	SetConfigRevision(v int64) *ContentPlanGenerationTask
	GetConfigRevision() *int64
	SetDiagnosisTaskId(v string) *ContentPlanGenerationTask
	GetDiagnosisTaskId() *string
	SetErrorCode(v string) *ContentPlanGenerationTask
	GetErrorCode() *string
	SetErrorMessage(v string) *ContentPlanGenerationTask
	GetErrorMessage() *string
	SetPlanId(v string) *ContentPlanGenerationTask
	GetPlanId() *string
	SetProgress(v int32) *ContentPlanGenerationTask
	GetProgress() *int32
	SetProjectId(v string) *ContentPlanGenerationTask
	GetProjectId() *string
	SetReportConfigRevision(v int64) *ContentPlanGenerationTask
	GetReportConfigRevision() *int64
	SetStaleConfig(v bool) *ContentPlanGenerationTask
	GetStaleConfig() *bool
	SetStatus(v string) *ContentPlanGenerationTask
	GetStatus() *string
	SetSubmittedAt(v int64) *ContentPlanGenerationTask
	GetSubmittedAt() *int64
	SetTaskId(v string) *ContentPlanGenerationTask
	GetTaskId() *string
	SetTaskType(v string) *ContentPlanGenerationTask
	GetTaskType() *string
	SetUpdatedAt(v int64) *ContentPlanGenerationTask
	GetUpdatedAt() *int64
}

type ContentPlanGenerationTask struct {
	// The analysis conclusions, returned only after the analysis task succeeds.
	Analysis *ContentPlanGenerationTaskAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The ID of the analysis task on which the scheduling task is based.
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
	// The ID of the GEO diagnostic task used by the analysis node.
	//
	// example:
	//
	// 1f2e3d4c-5b6a-4798-8a9b-c0d1e2f3a4b5
	DiagnosisTaskId *string `json:"DiagnosisTaskId,omitempty" xml:"DiagnosisTaskId,omitempty"`
	// The error code for the failure.
	//
	// example:
	//
	// CONTENT_PLAN_AGENT_TIMEOUT
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for the failure.
	//
	// example:
	//
	// Content plan generation timed out. Please try again later
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the content plan generated after the scheduling task succeeds.
	//
	// example:
	//
	// plan_3f7a1c2e9b7d4a5f8c0e1d2b3a4c5d6e
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The task progress, ranging from 0 to 10000.
	//
	// example:
	//
	// 10000
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the workspace, which refers to the GEO workspace for brand marketing, namely the appWorkspaceId.
	//
	// example:
	//
	// d4e5f60718293a4b5c6d7e8f90a1b2c3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The workspace configuration version on which the diagnostic report is based. The analysis, scheduling, and article generation processes always read this version.
	//
	// example:
	//
	// 12
	ReportConfigRevision *int64 `json:"ReportConfigRevision,omitempty" xml:"ReportConfigRevision,omitempty"`
	// Indicates whether the diagnostic report is generated based on an outdated configuration. If the value is true, the frontend should prompt the user.
	//
	// example:
	//
	// false
	StaleConfig *bool `json:"StaleConfig,omitempty" xml:"StaleConfig,omitempty"`
	// The task status.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submission time, in milliseconds.
	//
	// example:
	//
	// 1735660800000
	SubmittedAt *int64 `json:"SubmittedAt,omitempty" xml:"SubmittedAt,omitempty"`
	// The ID of the generation task.
	//
	// example:
	//
	// cpt_9b7d4a5f8c0e1d2b3a4c5d6e3f7a1c2e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values: ANALYSIS and SCHEDULE.
	//
	// example:
	//
	// SCHEDULE
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The last update time, in milliseconds.
	//
	// example:
	//
	// 1735660900000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ContentPlanGenerationTask) String() string {
	return dara.Prettify(s)
}

func (s ContentPlanGenerationTask) GoString() string {
	return s.String()
}

func (s *ContentPlanGenerationTask) GetAnalysis() *ContentPlanGenerationTaskAnalysis {
	return s.Analysis
}

func (s *ContentPlanGenerationTask) GetAnalysisTaskId() *string {
	return s.AnalysisTaskId
}

func (s *ContentPlanGenerationTask) GetConfigRevision() *int64 {
	return s.ConfigRevision
}

func (s *ContentPlanGenerationTask) GetDiagnosisTaskId() *string {
	return s.DiagnosisTaskId
}

func (s *ContentPlanGenerationTask) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ContentPlanGenerationTask) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ContentPlanGenerationTask) GetPlanId() *string {
	return s.PlanId
}

func (s *ContentPlanGenerationTask) GetProgress() *int32 {
	return s.Progress
}

func (s *ContentPlanGenerationTask) GetProjectId() *string {
	return s.ProjectId
}

func (s *ContentPlanGenerationTask) GetReportConfigRevision() *int64 {
	return s.ReportConfigRevision
}

func (s *ContentPlanGenerationTask) GetStaleConfig() *bool {
	return s.StaleConfig
}

func (s *ContentPlanGenerationTask) GetStatus() *string {
	return s.Status
}

func (s *ContentPlanGenerationTask) GetSubmittedAt() *int64 {
	return s.SubmittedAt
}

func (s *ContentPlanGenerationTask) GetTaskId() *string {
	return s.TaskId
}

func (s *ContentPlanGenerationTask) GetTaskType() *string {
	return s.TaskType
}

func (s *ContentPlanGenerationTask) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ContentPlanGenerationTask) SetAnalysis(v *ContentPlanGenerationTaskAnalysis) *ContentPlanGenerationTask {
	s.Analysis = v
	return s
}

func (s *ContentPlanGenerationTask) SetAnalysisTaskId(v string) *ContentPlanGenerationTask {
	s.AnalysisTaskId = &v
	return s
}

func (s *ContentPlanGenerationTask) SetConfigRevision(v int64) *ContentPlanGenerationTask {
	s.ConfigRevision = &v
	return s
}

func (s *ContentPlanGenerationTask) SetDiagnosisTaskId(v string) *ContentPlanGenerationTask {
	s.DiagnosisTaskId = &v
	return s
}

func (s *ContentPlanGenerationTask) SetErrorCode(v string) *ContentPlanGenerationTask {
	s.ErrorCode = &v
	return s
}

func (s *ContentPlanGenerationTask) SetErrorMessage(v string) *ContentPlanGenerationTask {
	s.ErrorMessage = &v
	return s
}

func (s *ContentPlanGenerationTask) SetPlanId(v string) *ContentPlanGenerationTask {
	s.PlanId = &v
	return s
}

func (s *ContentPlanGenerationTask) SetProgress(v int32) *ContentPlanGenerationTask {
	s.Progress = &v
	return s
}

func (s *ContentPlanGenerationTask) SetProjectId(v string) *ContentPlanGenerationTask {
	s.ProjectId = &v
	return s
}

func (s *ContentPlanGenerationTask) SetReportConfigRevision(v int64) *ContentPlanGenerationTask {
	s.ReportConfigRevision = &v
	return s
}

func (s *ContentPlanGenerationTask) SetStaleConfig(v bool) *ContentPlanGenerationTask {
	s.StaleConfig = &v
	return s
}

func (s *ContentPlanGenerationTask) SetStatus(v string) *ContentPlanGenerationTask {
	s.Status = &v
	return s
}

func (s *ContentPlanGenerationTask) SetSubmittedAt(v int64) *ContentPlanGenerationTask {
	s.SubmittedAt = &v
	return s
}

func (s *ContentPlanGenerationTask) SetTaskId(v string) *ContentPlanGenerationTask {
	s.TaskId = &v
	return s
}

func (s *ContentPlanGenerationTask) SetTaskType(v string) *ContentPlanGenerationTask {
	s.TaskType = &v
	return s
}

func (s *ContentPlanGenerationTask) SetUpdatedAt(v int64) *ContentPlanGenerationTask {
	s.UpdatedAt = &v
	return s
}

func (s *ContentPlanGenerationTask) Validate() error {
	if s.Analysis != nil {
		if err := s.Analysis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ContentPlanGenerationTaskAnalysis struct {
	// The analysis modules.
	Modules []*ContentPlanGenerationTaskAnalysisModules `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Repeated"`
	// The reference article writing guidelines, which include only structure and perspective, excluding third-party facts.
	ReferenceWritingNotes []*string `json:"ReferenceWritingNotes,omitempty" xml:"ReferenceWritingNotes,omitempty" type:"Repeated"`
	// The analysis summary.
	//
	// example:
	//
	// Insufficient brand awareness is the current main shortcoming. It is recommended to prioritize adding content related to product capabilities
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ContentPlanGenerationTaskAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ContentPlanGenerationTaskAnalysis) GoString() string {
	return s.String()
}

func (s *ContentPlanGenerationTaskAnalysis) GetModules() []*ContentPlanGenerationTaskAnalysisModules {
	return s.Modules
}

func (s *ContentPlanGenerationTaskAnalysis) GetReferenceWritingNotes() []*string {
	return s.ReferenceWritingNotes
}

func (s *ContentPlanGenerationTaskAnalysis) GetSummary() *string {
	return s.Summary
}

func (s *ContentPlanGenerationTaskAnalysis) SetModules(v []*ContentPlanGenerationTaskAnalysisModules) *ContentPlanGenerationTaskAnalysis {
	s.Modules = v
	return s
}

func (s *ContentPlanGenerationTaskAnalysis) SetReferenceWritingNotes(v []*string) *ContentPlanGenerationTaskAnalysis {
	s.ReferenceWritingNotes = v
	return s
}

func (s *ContentPlanGenerationTaskAnalysis) SetSummary(v string) *ContentPlanGenerationTaskAnalysis {
	s.Summary = &v
	return s
}

func (s *ContentPlanGenerationTaskAnalysis) Validate() error {
	if s.Modules != nil {
		for _, item := range s.Modules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ContentPlanGenerationTaskAnalysisModules struct {
	// The module content.
	//
	// example:
	//
	// The current brand mention rate in the AI website builder category is lower than the industry average, and user perception is mostly focused on pricing
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The key questions.
	KeyQuestions []*string `json:"KeyQuestions,omitempty" xml:"KeyQuestions,omitempty" type:"Repeated"`
	// The module title.
	//
	// example:
	//
	// Current status of brand awareness
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ContentPlanGenerationTaskAnalysisModules) String() string {
	return dara.Prettify(s)
}

func (s ContentPlanGenerationTaskAnalysisModules) GoString() string {
	return s.String()
}

func (s *ContentPlanGenerationTaskAnalysisModules) GetContent() *string {
	return s.Content
}

func (s *ContentPlanGenerationTaskAnalysisModules) GetKeyQuestions() []*string {
	return s.KeyQuestions
}

func (s *ContentPlanGenerationTaskAnalysisModules) GetTitle() *string {
	return s.Title
}

func (s *ContentPlanGenerationTaskAnalysisModules) SetContent(v string) *ContentPlanGenerationTaskAnalysisModules {
	s.Content = &v
	return s
}

func (s *ContentPlanGenerationTaskAnalysisModules) SetKeyQuestions(v []*string) *ContentPlanGenerationTaskAnalysisModules {
	s.KeyQuestions = v
	return s
}

func (s *ContentPlanGenerationTaskAnalysisModules) SetTitle(v string) *ContentPlanGenerationTaskAnalysisModules {
	s.Title = &v
	return s
}

func (s *ContentPlanGenerationTaskAnalysisModules) Validate() error {
	return dara.Validate(s)
}
