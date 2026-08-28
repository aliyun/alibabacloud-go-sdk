// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatrolReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListPatrolReportsResponseBodyData) *ListPatrolReportsResponseBody
	GetData() *ListPatrolReportsResponseBodyData
	SetErrorCode(v string) *ListPatrolReportsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPatrolReportsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListPatrolReportsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListPatrolReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPatrolReportsResponseBody
	GetSuccess() *bool
}

type ListPatrolReportsResponseBody struct {
	// The paginated data of inspection reports.
	Data *ListPatrolReportsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The business status code, which is always 200. Use the success field to determine whether the business request was successful.
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

func (s ListPatrolReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsResponseBody) GetData() *ListPatrolReportsResponseBodyData {
	return s.Data
}

func (s *ListPatrolReportsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPatrolReportsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPatrolReportsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListPatrolReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPatrolReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPatrolReportsResponseBody) SetData(v *ListPatrolReportsResponseBodyData) *ListPatrolReportsResponseBody {
	s.Data = v
	return s
}

func (s *ListPatrolReportsResponseBody) SetErrorCode(v string) *ListPatrolReportsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPatrolReportsResponseBody) SetErrorMessage(v string) *ListPatrolReportsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPatrolReportsResponseBody) SetHttpCode(v int32) *ListPatrolReportsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListPatrolReportsResponseBody) SetRequestId(v string) *ListPatrolReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPatrolReportsResponseBody) SetSuccess(v bool) *ListPatrolReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPatrolReportsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPatrolReportsResponseBodyData struct {
	// The list of inspection reports.
	Items []*ListPatrolReportsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 5
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPatrolReportsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsResponseBodyData) GetItems() []*ListPatrolReportsResponseBodyDataItems {
	return s.Items
}

func (s *ListPatrolReportsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListPatrolReportsResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *ListPatrolReportsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListPatrolReportsResponseBodyData) SetItems(v []*ListPatrolReportsResponseBodyDataItems) *ListPatrolReportsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPatrolReportsResponseBodyData) SetPage(v int32) *ListPatrolReportsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListPatrolReportsResponseBodyData) SetSize(v int32) *ListPatrolReportsResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListPatrolReportsResponseBodyData) SetTotal(v int32) *ListPatrolReportsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListPatrolReportsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPatrolReportsResponseBodyDataItems struct {
	// The inspection completion time.
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
	// 1755158793586
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The error message.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The number of inspected jobs.
	//
	// example:
	//
	// 2
	InspectedJobs *int32 `json:"inspectedJobs,omitempty" xml:"inspectedJobs,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The number of jobs with delay and backpressure.
	//
	// example:
	//
	// 1
	ProblemDelayBackpressureCount *int32 `json:"problemDelayBackpressureCount,omitempty" xml:"problemDelayBackpressureCount,omitempty"`
	// The number of jobs with unhealthy checkpoints.
	//
	// example:
	//
	// 0
	ProblemUnhealthyCheckpointCount *int32 `json:"problemUnhealthyCheckpointCount,omitempty" xml:"problemUnhealthyCheckpointCount,omitempty"`
	// The number of unhealthy jobs.
	//
	// example:
	//
	// 1
	ProblemUnhealthyCount *int32 `json:"problemUnhealthyCount,omitempty" xml:"problemUnhealthyCount,omitempty"`
	// The report ID.
	//
	// example:
	//
	// e7fd2601edc24a37baaba7eec5f64312
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The number of jobs with a Critical risk level.
	//
	// example:
	//
	// 0
	RiskCriticalCount *int32 `json:"riskCriticalCount,omitempty" xml:"riskCriticalCount,omitempty"`
	// The number of jobs with an Info risk level.
	//
	// example:
	//
	// 1
	RiskInfoCount *int32 `json:"riskInfoCount,omitempty" xml:"riskInfoCount,omitempty"`
	// The number of jobs with a Warning risk level.
	//
	// example:
	//
	// 1
	RiskWarningCount *int32 `json:"riskWarningCount,omitempty" xml:"riskWarningCount,omitempty"`
	// The inspection scope configuration.
	ScopeConfig *ListPatrolReportsResponseBodyDataItemsScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type.
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The inspection start time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1755158793586
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The report status.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total number of jobs.
	//
	// example:
	//
	// 5
	TotalJobs *int32 `json:"totalJobs,omitempty" xml:"totalJobs,omitempty"`
	// The trigger type.
	//
	// example:
	//
	// CRON
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1784168829417
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPatrolReportsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsResponseBodyDataItems) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *ListPatrolReportsResponseBodyDataItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListPatrolReportsResponseBodyDataItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPatrolReportsResponseBodyDataItems) GetInspectedJobs() *int32 {
	return s.InspectedJobs
}

func (s *ListPatrolReportsResponseBodyDataItems) GetNamespace() *string {
	return s.Namespace
}

func (s *ListPatrolReportsResponseBodyDataItems) GetProblemDelayBackpressureCount() *int32 {
	return s.ProblemDelayBackpressureCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetProblemUnhealthyCheckpointCount() *int32 {
	return s.ProblemUnhealthyCheckpointCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetProblemUnhealthyCount() *int32 {
	return s.ProblemUnhealthyCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetReportId() *string {
	return s.ReportId
}

func (s *ListPatrolReportsResponseBodyDataItems) GetRiskCriticalCount() *int32 {
	return s.RiskCriticalCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetRiskInfoCount() *int32 {
	return s.RiskInfoCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetRiskWarningCount() *int32 {
	return s.RiskWarningCount
}

func (s *ListPatrolReportsResponseBodyDataItems) GetScopeConfig() *ListPatrolReportsResponseBodyDataItemsScopeConfig {
	return s.ScopeConfig
}

func (s *ListPatrolReportsResponseBodyDataItems) GetScopeType() *string {
	return s.ScopeType
}

func (s *ListPatrolReportsResponseBodyDataItems) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListPatrolReportsResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListPatrolReportsResponseBodyDataItems) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *ListPatrolReportsResponseBodyDataItems) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListPatrolReportsResponseBodyDataItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListPatrolReportsResponseBodyDataItems) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPatrolReportsResponseBodyDataItems) SetCompletedAt(v string) *ListPatrolReportsResponseBodyDataItems {
	s.CompletedAt = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetCreatedAt(v string) *ListPatrolReportsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetErrorMessage(v string) *ListPatrolReportsResponseBodyDataItems {
	s.ErrorMessage = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetInspectedJobs(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.InspectedJobs = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetNamespace(v string) *ListPatrolReportsResponseBodyDataItems {
	s.Namespace = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetProblemDelayBackpressureCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.ProblemDelayBackpressureCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetProblemUnhealthyCheckpointCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.ProblemUnhealthyCheckpointCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetProblemUnhealthyCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.ProblemUnhealthyCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetReportId(v string) *ListPatrolReportsResponseBodyDataItems {
	s.ReportId = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetRiskCriticalCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.RiskCriticalCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetRiskInfoCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.RiskInfoCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetRiskWarningCount(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.RiskWarningCount = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetScopeConfig(v *ListPatrolReportsResponseBodyDataItemsScopeConfig) *ListPatrolReportsResponseBodyDataItems {
	s.ScopeConfig = v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetScopeType(v string) *ListPatrolReportsResponseBodyDataItems {
	s.ScopeType = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetStartedAt(v string) *ListPatrolReportsResponseBodyDataItems {
	s.StartedAt = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetStatus(v string) *ListPatrolReportsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetTotalJobs(v int32) *ListPatrolReportsResponseBodyDataItems {
	s.TotalJobs = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetTriggerType(v string) *ListPatrolReportsResponseBodyDataItems {
	s.TriggerType = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetUpdatedAt(v string) *ListPatrolReportsResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) SetWorkspace(v string) *ListPatrolReportsResponseBodyDataItems {
	s.Workspace = &v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItems) Validate() error {
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPatrolReportsResponseBodyDataItemsScopeConfig struct {
	// The list of deployment IDs. This field is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This field is valid only when scopeType is set to TAGS. The key is the tag name, and the value is a list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListPatrolReportsResponseBodyDataItemsScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsResponseBodyDataItemsScopeConfig) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsResponseBodyDataItemsScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *ListPatrolReportsResponseBodyDataItemsScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *ListPatrolReportsResponseBodyDataItemsScopeConfig) SetDeploymentIds(v []*string) *ListPatrolReportsResponseBodyDataItemsScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItemsScopeConfig) SetTags(v map[string][]*string) *ListPatrolReportsResponseBodyDataItemsScopeConfig {
	s.Tags = v
	return s
}

func (s *ListPatrolReportsResponseBodyDataItemsScopeConfig) Validate() error {
	return dara.Validate(s)
}
