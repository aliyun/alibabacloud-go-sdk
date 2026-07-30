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
	SetDescription(v string) *ListJobsRequest
	GetDescription() *string
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
	SetResourceIds(v string) *ListJobsRequest
	GetResourceIds() *string
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
	SetTemplateId(v string) *ListJobsRequest
	GetTemplateId() *string
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
	// The visibility of the job. Valid values:
	//
	// - PUBLIC: Visible to all members in the workspace.
	//
	// - PRIVATE (default): Visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The user ID associated with the job.
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
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job name. Supports fuzzy match and is case-insensitive. Wildcards are not supported.
	//
	// For example, entering test matches test-job1, job-test, job-test2, or job-Test, but does not match job-t1.
	//
	// Default value: empty, which indicates all job names.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The search mode for DisplayName. Default value: wildcard match.
	//
	// example:
	//
	// wildcard
	DisplayNameSearchMode *string `json:"DisplayNameSearchMode,omitempty" xml:"DisplayNameSearchMode,omitempty"`
	// Filters jobs based on whether running on specified nodes is enabled.
	//
	// example:
	//
	// true
	EnableAssignNode *string `json:"EnableAssignNode,omitempty" xml:"EnableAssignNode,omitempty"`
	// The end time of the query range. The job creation time is used for filtering. Default value: the current time.
	//
	// example:
	//
	// 2025-04-16T07:26:41Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to retrieve jobs across all workspaces. This parameter must be used together with `ShowOwn=true` to query jobs recently submitted by the current user.
	//
	// example:
	//
	// false
	FromAllWorkspaces *bool `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	// Retrieves nodes by performing a full-text index on the images field. Supports Chinese and English tokenization.
	//
	// example:
	//
	// pytorch
	ImageSearch *string `json:"ImageSearch,omitempty" xml:"ImageSearch,omitempty"`
	// The job ID. Fuzzy match is not supported. Case-insensitive. Wildcards are not supported.
	//
	// Default value: empty, which indicates all job IDs.
	//
	// example:
	//
	// dlc********
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// A list of job IDs separated by commas. If both JobIds and JobId are specified, JobId takes precedence.
	//
	// example:
	//
	// dlc123abc
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The job type. Default value: empty, which indicates all types. Valid values:
	//
	// - TFJob
	//
	// - PyTorchJob
	//
	// - XGBoostJob
	//
	// - OneFlowJob
	//
	// - ElasticBatchJob
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The field name for numeric range filtering. Must be used together with NumericRangeMin/NumericRangeMax.
	//
	// example:
	//
	// RequestGPU
	NumericRangeField *string `json:"NumericRangeField,omitempty" xml:"NumericRangeField,omitempty"`
	// The maximum value (inclusive) for numeric range filtering. Must be used together with NumericRangeField.
	//
	// example:
	//
	// 8
	NumericRangeMax *int64 `json:"NumericRangeMax,omitempty" xml:"NumericRangeMax,omitempty"`
	// The minimum value (inclusive) for numeric range filtering. Must be used together with NumericRangeField.
	//
	// example:
	//
	// 4
	NumericRangeMin *int64 `json:"NumericRangeMin,omitempty" xml:"NumericRangeMin,omitempty"`
	// The sort order. Valid values:
	//
	// - desc: Descending order. This is the default value.
	//
	// - asc: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The idle resource information. Valid values:
	//
	// - ForbiddenQuotaOverSold
	//
	// - ForceQuotaOverSold
	//
	// - AcceptQuotaOverSold-true (true indicates the job actually used idle resources)
	//
	// - AcceptQuotaOverSold-false (false indicates the job actually used guaranteed resources)
	//
	// example:
	//
	// ForbiddenQuotaOverSold
	OversoldInfo *string `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	// The page number to return in a paged query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs to return per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type. Valid values:
	//
	// - PrePaid: resource quota.
	//
	// - Spot: preemptible resources.
	//
	// - PostPaid: public resources.
	//
	// example:
	//
	// PostPaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Filters jobs created by the specified workflow ID.
	//
	// example:
	//
	// flow-*******
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// Retrieves nodes by performing a full-text index on the node failed reason field. Supports Chinese and English tokenization.
	//
	// example:
	//
	// OOM
	ReasonSearch *string `json:"ReasonSearch,omitempty" xml:"ReasonSearch,omitempty"`
	// The resource group ID. For information about how to query the dedicated resource group ID, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// r*****
	ResourceId  *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// Filters the job list by the resource quota name. Supports fuzzy match. Wildcards are not supported. Default value: empty, which indicates no filtering by resource quota.
	//
	// example:
	//
	// quota***
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// Specifies whether to return only jobs submitted by the current user.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The field by which to sort results:
	//
	// - DisplayName
	//
	// - JobType
	//
	// - Status
	//
	// - GmtCreateTime
	//
	// - GmtFinishTime
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the query range. The job creation time is used for filtering. Default value: the current time minus 7 days. If neither StartTime nor EndTime is specified, jobs from the last 7 days are returned by default.
	//
	// example:
	//
	// 2025-04-16T07:25:34Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job status. Valid values:
	//
	// - Creating
	//
	// - Queuing
	//
	// - Bidding (currently only for Lingjun Spot jobs)
	//
	// - EnvPreparing
	//
	// - SanityChecking
	//
	// - Running
	//
	// - Restarting
	//
	// - Stopping
	//
	// - SucceededReserving
	//
	// - FailedReserving
	//
	// - Succeeded
	//
	// - Failed
	//
	// - Stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags used for filtering.
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The template ID. Filters jobs created from the specified template.
	//
	// example:
	//
	// tmlabc123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The time field used for StartTime/EndTime filtering. Default value: creation time.
	//
	// example:
	//
	// GmtFinishTime
	TimeRangeField *string `json:"TimeRangeField,omitempty" xml:"TimeRangeField,omitempty"`
	// Retrieves nodes by performing a full-text index on the user_command field. Supports Chinese and English tokenization.
	//
	// example:
	//
	// python train.py
	UserCommandSearch *string `json:"UserCommandSearch,omitempty" xml:"UserCommandSearch,omitempty"`
	// Filters the job list by the user ID of the job submitter.
	//
	// example:
	//
	// 20**************
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// Filters the job list by the username of the job submitter. Supports fuzzy match. Wildcards are not supported. Default value: empty, which indicates no filtering by username.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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

func (s *ListJobsRequest) GetDescription() *string {
	return s.Description
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

func (s *ListJobsRequest) GetResourceIds() *string {
	return s.ResourceIds
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

func (s *ListJobsRequest) GetTemplateId() *string {
	return s.TemplateId
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

func (s *ListJobsRequest) SetDescription(v string) *ListJobsRequest {
	s.Description = &v
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

func (s *ListJobsRequest) SetResourceIds(v string) *ListJobsRequest {
	s.ResourceIds = &v
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

func (s *ListJobsRequest) SetTemplateId(v string) *ListJobsRequest {
	s.TemplateId = &v
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
