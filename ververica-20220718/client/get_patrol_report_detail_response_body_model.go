// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolReportDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPatrolReportDetailResponseBodyData) *GetPatrolReportDetailResponseBody
	GetData() *GetPatrolReportDetailResponseBodyData
	SetErrorCode(v string) *GetPatrolReportDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPatrolReportDetailResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetPatrolReportDetailResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetPatrolReportDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPatrolReportDetailResponseBody
	GetSuccess() *bool
}

type GetPatrolReportDetailResponseBody struct {
	// The inspection report details.
	Data *GetPatrolReportDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The business error code. This value is not empty when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This value is not empty when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use success to determine whether the business request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPatrolReportDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBody) GetData() *GetPatrolReportDetailResponseBodyData {
	return s.Data
}

func (s *GetPatrolReportDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPatrolReportDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPatrolReportDetailResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetPatrolReportDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPatrolReportDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPatrolReportDetailResponseBody) SetData(v *GetPatrolReportDetailResponseBodyData) *GetPatrolReportDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetPatrolReportDetailResponseBody) SetErrorCode(v string) *GetPatrolReportDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPatrolReportDetailResponseBody) SetErrorMessage(v string) *GetPatrolReportDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPatrolReportDetailResponseBody) SetHttpCode(v int32) *GetPatrolReportDetailResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPatrolReportDetailResponseBody) SetRequestId(v string) *GetPatrolReportDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPatrolReportDetailResponseBody) SetSuccess(v bool) *GetPatrolReportDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetPatrolReportDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPatrolReportDetailResponseBodyData struct {
	// The completion time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1785981634652
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1772936711518
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The inspection overview.
	Overview *GetPatrolReportDetailResponseBodyDataOverview `json:"overview,omitempty" xml:"overview,omitempty" type:"Struct"`
	// The report ID.
	//
	// example:
	//
	// inspection-cf8f8843-64e4-4b45-9500-06790107130f
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The inspection scope configuration.
	ScopeConfig *GetPatrolReportDetailResponseBodyDataScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type.
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The trigger type.
	//
	// example:
	//
	// CRON
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// The list of unhealthy jobs.
	UnhealthyJobs []*GetPatrolReportDetailResponseBodyDataUnhealthyJobs `json:"unhealthyJobs,omitempty" xml:"unhealthyJobs,omitempty" type:"Repeated"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1758248445816
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPatrolReportDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyData) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *GetPatrolReportDetailResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetPatrolReportDetailResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetPatrolReportDetailResponseBodyData) GetOverview() *GetPatrolReportDetailResponseBodyDataOverview {
	return s.Overview
}

func (s *GetPatrolReportDetailResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *GetPatrolReportDetailResponseBodyData) GetScopeConfig() *GetPatrolReportDetailResponseBodyDataScopeConfig {
	return s.ScopeConfig
}

func (s *GetPatrolReportDetailResponseBodyData) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetPatrolReportDetailResponseBodyData) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetPatrolReportDetailResponseBodyData) GetUnhealthyJobs() []*GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	return s.UnhealthyJobs
}

func (s *GetPatrolReportDetailResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetPatrolReportDetailResponseBodyData) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPatrolReportDetailResponseBodyData) SetCompletedAt(v string) *GetPatrolReportDetailResponseBodyData {
	s.CompletedAt = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetCreatedAt(v string) *GetPatrolReportDetailResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetNamespace(v string) *GetPatrolReportDetailResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetOverview(v *GetPatrolReportDetailResponseBodyDataOverview) *GetPatrolReportDetailResponseBodyData {
	s.Overview = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetReportId(v string) *GetPatrolReportDetailResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetScopeConfig(v *GetPatrolReportDetailResponseBodyDataScopeConfig) *GetPatrolReportDetailResponseBodyData {
	s.ScopeConfig = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetScopeType(v string) *GetPatrolReportDetailResponseBodyData {
	s.ScopeType = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetTriggerType(v string) *GetPatrolReportDetailResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetUnhealthyJobs(v []*GetPatrolReportDetailResponseBodyDataUnhealthyJobs) *GetPatrolReportDetailResponseBodyData {
	s.UnhealthyJobs = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetUpdatedAt(v string) *GetPatrolReportDetailResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) SetWorkspace(v string) *GetPatrolReportDetailResponseBodyData {
	s.Workspace = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyData) Validate() error {
	if s.Overview != nil {
		if err := s.Overview.Validate(); err != nil {
			return err
		}
	}
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.UnhealthyJobs != nil {
		for _, item := range s.UnhealthyJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPatrolReportDetailResponseBodyDataOverview struct {
	// The problem summary.
	ProblemSummary *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary `json:"problemSummary,omitempty" xml:"problemSummary,omitempty" type:"Struct"`
	// The risk summary.
	RiskSummary *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary `json:"riskSummary,omitempty" xml:"riskSummary,omitempty" type:"Struct"`
	// The total number of jobs.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetPatrolReportDetailResponseBodyDataOverview) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyDataOverview) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) GetProblemSummary() *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary {
	return s.ProblemSummary
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) GetRiskSummary() *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary {
	return s.RiskSummary
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) GetTotal() *int32 {
	return s.Total
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) SetProblemSummary(v *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) *GetPatrolReportDetailResponseBodyDataOverview {
	s.ProblemSummary = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) SetRiskSummary(v *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) *GetPatrolReportDetailResponseBodyDataOverview {
	s.RiskSummary = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) SetTotal(v int32) *GetPatrolReportDetailResponseBodyDataOverview {
	s.Total = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverview) Validate() error {
	if s.ProblemSummary != nil {
		if err := s.ProblemSummary.Validate(); err != nil {
			return err
		}
	}
	if s.RiskSummary != nil {
		if err := s.RiskSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPatrolReportDetailResponseBodyDataOverviewProblemSummary struct {
	// The number of jobs with delay and backpressure.
	//
	// example:
	//
	// 0
	DelayAndBackpressure *int32 `json:"delayAndBackpressure,omitempty" xml:"delayAndBackpressure,omitempty"`
	// The number of unhealthy jobs.
	//
	// example:
	//
	// 0
	Unhealthy *int32 `json:"unhealthy,omitempty" xml:"unhealthy,omitempty"`
	// The number of jobs with unhealthy checkpoints.
	//
	// example:
	//
	// 0
	UnhealthyCheckpoints *int32 `json:"unhealthyCheckpoints,omitempty" xml:"unhealthyCheckpoints,omitempty"`
}

func (s GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) GetDelayAndBackpressure() *int32 {
	return s.DelayAndBackpressure
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) GetUnhealthy() *int32 {
	return s.Unhealthy
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) GetUnhealthyCheckpoints() *int32 {
	return s.UnhealthyCheckpoints
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) SetDelayAndBackpressure(v int32) *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary {
	s.DelayAndBackpressure = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) SetUnhealthy(v int32) *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary {
	s.Unhealthy = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) SetUnhealthyCheckpoints(v int32) *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary {
	s.UnhealthyCheckpoints = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewProblemSummary) Validate() error {
	return dara.Validate(s)
}

type GetPatrolReportDetailResponseBodyDataOverviewRiskSummary struct {
	// The number of critical-level jobs.
	//
	// example:
	//
	// 0
	Critical *int32 `json:"critical,omitempty" xml:"critical,omitempty"`
	// The number of info-level jobs.
	//
	// example:
	//
	// 0
	Info *int32 `json:"info,omitempty" xml:"info,omitempty"`
	// The number of warning-level jobs.
	//
	// example:
	//
	// 1
	Warning *int32 `json:"warning,omitempty" xml:"warning,omitempty"`
}

func (s GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) GetCritical() *int32 {
	return s.Critical
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) GetInfo() *int32 {
	return s.Info
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) GetWarning() *int32 {
	return s.Warning
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) SetCritical(v int32) *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary {
	s.Critical = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) SetInfo(v int32) *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary {
	s.Info = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) SetWarning(v int32) *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary {
	s.Warning = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataOverviewRiskSummary) Validate() error {
	return dara.Validate(s)
}

type GetPatrolReportDetailResponseBodyDataScopeConfig struct {
	// The list of deployment IDs. This parameter is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This parameter is valid only when scopeType is set to TAGS. The key is the tag name, and the value is the list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetPatrolReportDetailResponseBodyDataScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyDataScopeConfig) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyDataScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *GetPatrolReportDetailResponseBodyDataScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *GetPatrolReportDetailResponseBodyDataScopeConfig) SetDeploymentIds(v []*string) *GetPatrolReportDetailResponseBodyDataScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataScopeConfig) SetTags(v map[string][]*string) *GetPatrolReportDetailResponseBodyDataScopeConfig {
	s.Tags = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataScopeConfig) Validate() error {
	return dara.Validate(s)
}

type GetPatrolReportDetailResponseBodyDataUnhealthyJobs struct {
	// The analysis.
	//
	// example:
	//
	// The job has experienced continuous backpressure in the last 30 minutes.
	Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// The deployment ID.
	//
	// example:
	//
	// 18b8ceaa-207d-417b-833e-a5845bb31beb
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The deployment name.
	//
	// example:
	//
	// rlt_cust_no_apitime
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// The problem description.
	//
	// example:
	//
	// Job backpressure
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The job ID.
	//
	// example:
	//
	// db201864-e0df-4f8c-81f0-d62103095ff6
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The recommendation.
	//
	// example:
	//
	// Check the processing capacity of the downstream operator.
	Recommendation *string `json:"recommendation,omitempty" xml:"recommendation,omitempty"`
	// The risk level.
	//
	// example:
	//
	// info
	RiskLevel *string `json:"riskLevel,omitempty" xml:"riskLevel,omitempty"`
	// The list of tags diagnosed by AI for the job.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s GetPatrolReportDetailResponseBodyDataUnhealthyJobs) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GoString() string {
	return s.String()
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetAnalysis() *string {
	return s.Analysis
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetDescription() *string {
	return s.Description
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetJobId() *string {
	return s.JobId
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetRecommendation() *string {
	return s.Recommendation
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) GetTags() []*string {
	return s.Tags
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetAnalysis(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.Analysis = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetDeploymentId(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.DeploymentId = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetDeploymentName(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.DeploymentName = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetDescription(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.Description = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetJobId(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.JobId = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetRecommendation(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.Recommendation = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetRiskLevel(v string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.RiskLevel = &v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) SetTags(v []*string) *GetPatrolReportDetailResponseBodyDataUnhealthyJobs {
	s.Tags = v
	return s
}

func (s *GetPatrolReportDetailResponseBodyDataUnhealthyJobs) Validate() error {
	return dara.Validate(s)
}
