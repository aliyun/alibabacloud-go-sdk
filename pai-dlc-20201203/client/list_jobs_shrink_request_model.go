// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListJobsShrinkRequest
	GetAccessibility() *string
	SetBusinessUserId(v string) *ListJobsShrinkRequest
	GetBusinessUserId() *string
	SetCaller(v string) *ListJobsShrinkRequest
	GetCaller() *string
	SetDisplayName(v string) *ListJobsShrinkRequest
	GetDisplayName() *string
	SetEndTime(v string) *ListJobsShrinkRequest
	GetEndTime() *string
	SetFromAllWorkspaces(v bool) *ListJobsShrinkRequest
	GetFromAllWorkspaces() *bool
	SetJobId(v string) *ListJobsShrinkRequest
	GetJobId() *string
	SetJobIds(v string) *ListJobsShrinkRequest
	GetJobIds() *string
	SetJobType(v string) *ListJobsShrinkRequest
	GetJobType() *string
	SetOrder(v string) *ListJobsShrinkRequest
	GetOrder() *string
	SetOversoldInfo(v string) *ListJobsShrinkRequest
	GetOversoldInfo() *string
	SetPageNumber(v int32) *ListJobsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsShrinkRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListJobsShrinkRequest
	GetPaymentType() *string
	SetPipelineId(v string) *ListJobsShrinkRequest
	GetPipelineId() *string
	SetResourceId(v string) *ListJobsShrinkRequest
	GetResourceId() *string
	SetResourceQuotaName(v string) *ListJobsShrinkRequest
	GetResourceQuotaName() *string
	SetShowOwn(v bool) *ListJobsShrinkRequest
	GetShowOwn() *bool
	SetSortBy(v string) *ListJobsShrinkRequest
	GetSortBy() *string
	SetStartTime(v string) *ListJobsShrinkRequest
	GetStartTime() *string
	SetStatus(v string) *ListJobsShrinkRequest
	GetStatus() *string
	SetTagsShrink(v string) *ListJobsShrinkRequest
	GetTagsShrink() *string
	SetUserIdForFilter(v string) *ListJobsShrinkRequest
	GetUserIdForFilter() *string
	SetUsername(v string) *ListJobsShrinkRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListJobsShrinkRequest
	GetWorkspaceId() *string
}

type ListJobsShrinkRequest struct {
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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
	FromAllWorkspaces *bool `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
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
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
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
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s ListJobsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListJobsShrinkRequest) GetBusinessUserId() *string {
	return s.BusinessUserId
}

func (s *ListJobsShrinkRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListJobsShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListJobsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobsShrinkRequest) GetFromAllWorkspaces() *bool {
	return s.FromAllWorkspaces
}

func (s *ListJobsShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsShrinkRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *ListJobsShrinkRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListJobsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListJobsShrinkRequest) GetOversoldInfo() *string {
	return s.OversoldInfo
}

func (s *ListJobsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListJobsShrinkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ListJobsShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListJobsShrinkRequest) GetResourceQuotaName() *string {
	return s.ResourceQuotaName
}

func (s *ListJobsShrinkRequest) GetShowOwn() *bool {
	return s.ShowOwn
}

func (s *ListJobsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListJobsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListJobsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListJobsShrinkRequest) GetUserIdForFilter() *string {
	return s.UserIdForFilter
}

func (s *ListJobsShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *ListJobsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobsShrinkRequest) SetAccessibility(v string) *ListJobsShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *ListJobsShrinkRequest) SetBusinessUserId(v string) *ListJobsShrinkRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetCaller(v string) *ListJobsShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsShrinkRequest) SetDisplayName(v string) *ListJobsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetEndTime(v string) *ListJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetFromAllWorkspaces(v bool) *ListJobsShrinkRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobId(v string) *ListJobsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobIds(v string) *ListJobsShrinkRequest {
	s.JobIds = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobType(v string) *ListJobsShrinkRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsShrinkRequest) SetOrder(v string) *ListJobsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListJobsShrinkRequest) SetOversoldInfo(v string) *ListJobsShrinkRequest {
	s.OversoldInfo = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v int32) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v int32) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPaymentType(v string) *ListJobsShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPipelineId(v string) *ListJobsShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceId(v string) *ListJobsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceQuotaName(v string) *ListJobsShrinkRequest {
	s.ResourceQuotaName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetShowOwn(v bool) *ListJobsShrinkRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortBy(v string) *ListJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStartTime(v string) *ListJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStatus(v string) *ListJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListJobsShrinkRequest) SetTagsShrink(v string) *ListJobsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetUserIdForFilter(v string) *ListJobsShrinkRequest {
	s.UserIdForFilter = &v
	return s
}

func (s *ListJobsShrinkRequest) SetUsername(v string) *ListJobsShrinkRequest {
	s.Username = &v
	return s
}

func (s *ListJobsShrinkRequest) SetWorkspaceId(v string) *ListJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
