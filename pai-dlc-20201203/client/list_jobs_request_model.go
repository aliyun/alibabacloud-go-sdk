// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListJobsRequest
	GetAccessibility() *string
	SetBusinessUserId(v string) *ListJobsRequest
	GetBusinessUserId() *string
	SetCaller(v string) *ListJobsRequest
	GetCaller() *string
	SetDisplayName(v string) *ListJobsRequest
	GetDisplayName() *string
	SetDisplayNameSearchMode(v string) *ListJobsRequest
	GetDisplayNameSearchMode() *string
	SetEnableAssignNode(v string) *ListJobsRequest
	GetEnableAssignNode() *string
	SetEndTime(v string) *ListJobsRequest
	GetEndTime() *string
	SetFromAllWorkspaces(v bool) *ListJobsRequest
	GetFromAllWorkspaces() *bool
	SetImageSearch(v string) *ListJobsRequest
	GetImageSearch() *string
	SetJobId(v string) *ListJobsRequest
	GetJobId() *string
	SetJobIds(v string) *ListJobsRequest
	GetJobIds() *string
	SetJobType(v string) *ListJobsRequest
	GetJobType() *string
	SetNumericRangeField(v string) *ListJobsRequest
	GetNumericRangeField() *string
	SetNumericRangeMax(v int64) *ListJobsRequest
	GetNumericRangeMax() *int64
	SetNumericRangeMin(v int64) *ListJobsRequest
	GetNumericRangeMin() *int64
	SetOrder(v string) *ListJobsRequest
	GetOrder() *string
	SetOversoldInfo(v string) *ListJobsRequest
	GetOversoldInfo() *string
	SetPageNumber(v int32) *ListJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListJobsRequest
	GetPaymentType() *string
	SetPipelineId(v string) *ListJobsRequest
	GetPipelineId() *string
	SetReasonSearch(v string) *ListJobsRequest
	GetReasonSearch() *string
	SetResourceId(v string) *ListJobsRequest
	GetResourceId() *string
	SetResourceQuotaName(v string) *ListJobsRequest
	GetResourceQuotaName() *string
	SetShowOwn(v bool) *ListJobsRequest
	GetShowOwn() *bool
	SetSortBy(v string) *ListJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListJobsRequest
	GetStatus() *string
	SetTags(v map[string]*string) *ListJobsRequest
	GetTags() map[string]*string
	SetTimeRangeField(v string) *ListJobsRequest
	GetTimeRangeField() *string
	SetUserCommandSearch(v string) *ListJobsRequest
	GetUserCommandSearch() *string
	SetUserIdForFilter(v string) *ListJobsRequest
	GetUserIdForFilter() *string
	SetUsername(v string) *ListJobsRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListJobsRequest
	GetWorkspaceId() *string
}

type ListJobsRequest struct {
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The ID of the user associated with the job.
	//
	// example:
	//
	// 16****
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// The caller.
	//
	// example:
	//
	// local
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The job name. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. For example, if you enter test, test-job1, job-test, job-test2, or job-test can be matched, and job-t1 cannot be matched. The default value null indicates any job name.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName           *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameSearchMode *string `json:"DisplayNameSearchMode,omitempty" xml:"DisplayNameSearchMode,omitempty"`
	EnableAssignNode      *string `json:"EnableAssignNode,omitempty" xml:"EnableAssignNode,omitempty"`
	// The end time of the query. Use the job creation time to filter data. The default value is the current time.
	//
	// example:
	//
	// 2020-11-09T14:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query a list of jobs across workspaces. This parameter must be used together with `ShowOwn=true`. You can use this parameter to query a list of jobs recently submitted by the current user.
	//
	// example:
	//
	// false
	FromAllWorkspaces *bool   `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	ImageSearch       *string `json:"ImageSearch,omitempty" xml:"ImageSearch,omitempty"`
	// The job ID. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. The default value null indicates any job ID.
	//
	// example:
	//
	// dlc********
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The job type. The default value null indicates any type. Valid values:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// 	- XGBoostJob
	//
	// 	- OneFlowJob
	//
	// 	- ElasticBatchJob
	//
	// example:
	//
	// TFJob
	JobType           *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	NumericRangeField *string `json:"NumericRangeField,omitempty" xml:"NumericRangeField,omitempty"`
	NumericRangeMax   *int64  `json:"NumericRangeMax,omitempty" xml:"NumericRangeMax,omitempty"`
	NumericRangeMin   *int64  `json:"NumericRangeMin,omitempty" xml:"NumericRangeMin,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default)
	//
	// 	- asc
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Idle resource information. Valid values:
	//
	// 	- ForbiddenQuotaOverSold
	//
	// 	- ForceQuotaOverSold
	//
	// 	- AcceptQuotaOverSold-true (true indicates that the job uses idle resources.)
	//
	// 	- AcceptQuotaOverSold-false (false indicates that the job uses guaranteed resources.)
	//
	// example:
	//
	// ForbiddenQuotaOverSold
	OversoldInfo *string `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	// The number of the page to return for the current query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- PrePaid: Resource quota
	//
	// 	- Spot: Preemptible resources
	//
	// 	- PostPaid: Public resources
	//
	// example:
	//
	// PostPaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The specific pipeline ID used to filter jobs.
	//
	// example:
	//
	// flow-*******
	PipelineId   *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ReasonSearch *string `json:"ReasonSearch,omitempty" xml:"ReasonSearch,omitempty"`
	// The resource group ID. For information about how to obtain the ID of a dedicated resource group, see [Manage resource quota](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// r*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource quota name used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by resource quota name.
	//
	// example:
	//
	// quota***
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// Specifies whether to query only the jobs submitted by the current user.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The sorting field. Valid values:
	//
	// 	- DisplayName
	//
	// 	- JobType
	//
	// 	- Status
	//
	// 	- GmtCreateTime
	//
	// 	- GmtFinishTime
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the query. Use the job creation time to filter data. The default value is the current time minus seven days. In other words, if you do not configure the StartTime and EndTime parameters, the system queries the job list in the last seven days.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job status. Valid values:
	//
	// 	- Creating
	//
	// 	- Queuing
	//
	// 	- Bidding (only available for spot jobs that use Lingjun resources)
	//
	// 	- EnvPreparing
	//
	// 	- SanityChecking
	//
	// 	- Running
	//
	// 	- Restarting
	//
	// 	- Stopping
	//
	// 	- SucceededReserving
	//
	// 	- FailedReserving
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags              map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TimeRangeField    *string            `json:"TimeRangeField,omitempty" xml:"TimeRangeField,omitempty"`
	UserCommandSearch *string            `json:"UserCommandSearch,omitempty" xml:"UserCommandSearch,omitempty"`
	// The user ID used to filter jobs.
	//
	// example:
	//
	// 20**************
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// The username used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by username.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListJobsRequest) GetBusinessUserId() *string {
	return s.BusinessUserId
}

func (s *ListJobsRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListJobsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListJobsRequest) GetDisplayNameSearchMode() *string {
	return s.DisplayNameSearchMode
}

func (s *ListJobsRequest) GetEnableAssignNode() *string {
	return s.EnableAssignNode
}

func (s *ListJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobsRequest) GetFromAllWorkspaces() *bool {
	return s.FromAllWorkspaces
}

func (s *ListJobsRequest) GetImageSearch() *string {
	return s.ImageSearch
}

func (s *ListJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *ListJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListJobsRequest) GetNumericRangeField() *string {
	return s.NumericRangeField
}

func (s *ListJobsRequest) GetNumericRangeMax() *int64 {
	return s.NumericRangeMax
}

func (s *ListJobsRequest) GetNumericRangeMin() *int64 {
	return s.NumericRangeMin
}

func (s *ListJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobsRequest) GetOversoldInfo() *string {
	return s.OversoldInfo
}

func (s *ListJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListJobsRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListJobsRequest) GetReasonSearch() *string {
	return s.ReasonSearch
}

func (s *ListJobsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListJobsRequest) GetResourceQuotaName() *string {
	return s.ResourceQuotaName
}

func (s *ListJobsRequest) GetShowOwn() *bool {
	return s.ShowOwn
}

func (s *ListJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *ListJobsRequest) GetTimeRangeField() *string {
	return s.TimeRangeField
}

func (s *ListJobsRequest) GetUserCommandSearch() *string {
	return s.UserCommandSearch
}

func (s *ListJobsRequest) GetUserIdForFilter() *string {
	return s.UserIdForFilter
}

func (s *ListJobsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListJobsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobsRequest) SetAccessibility(v string) *ListJobsRequest {
	s.Accessibility = &v
	return s
}

func (s *ListJobsRequest) SetBusinessUserId(v string) *ListJobsRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsRequest) SetCaller(v string) *ListJobsRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsRequest) SetDisplayName(v string) *ListJobsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsRequest) SetDisplayNameSearchMode(v string) *ListJobsRequest {
	s.DisplayNameSearchMode = &v
	return s
}

func (s *ListJobsRequest) SetEnableAssignNode(v string) *ListJobsRequest {
	s.EnableAssignNode = &v
	return s
}

func (s *ListJobsRequest) SetEndTime(v string) *ListJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsRequest) SetFromAllWorkspaces(v bool) *ListJobsRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsRequest) SetImageSearch(v string) *ListJobsRequest {
	s.ImageSearch = &v
	return s
}

func (s *ListJobsRequest) SetJobId(v string) *ListJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsRequest) SetJobIds(v string) *ListJobsRequest {
	s.JobIds = &v
	return s
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsRequest) SetNumericRangeField(v string) *ListJobsRequest {
	s.NumericRangeField = &v
	return s
}

func (s *ListJobsRequest) SetNumericRangeMax(v int64) *ListJobsRequest {
	s.NumericRangeMax = &v
	return s
}

func (s *ListJobsRequest) SetNumericRangeMin(v int64) *ListJobsRequest {
	s.NumericRangeMin = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
	return s
}

func (s *ListJobsRequest) SetOversoldInfo(v string) *ListJobsRequest {
	s.OversoldInfo = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetPaymentType(v string) *ListJobsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListJobsRequest) SetPipelineId(v string) *ListJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsRequest) SetReasonSearch(v string) *ListJobsRequest {
	s.ReasonSearch = &v
	return s
}

func (s *ListJobsRequest) SetResourceId(v string) *ListJobsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsRequest) SetResourceQuotaName(v string) *ListJobsRequest {
	s.ResourceQuotaName = &v
	return s
}

func (s *ListJobsRequest) SetShowOwn(v bool) *ListJobsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v string) *ListJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsRequest) SetStartTime(v string) *ListJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetTags(v map[string]*string) *ListJobsRequest {
	s.Tags = v
	return s
}

func (s *ListJobsRequest) SetTimeRangeField(v string) *ListJobsRequest {
	s.TimeRangeField = &v
	return s
}

func (s *ListJobsRequest) SetUserCommandSearch(v string) *ListJobsRequest {
	s.UserCommandSearch = &v
	return s
}

func (s *ListJobsRequest) SetUserIdForFilter(v string) *ListJobsRequest {
	s.UserIdForFilter = &v
	return s
}

func (s *ListJobsRequest) SetUsername(v string) *ListJobsRequest {
	s.Username = &v
	return s
}

func (s *ListJobsRequest) SetWorkspaceId(v string) *ListJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
