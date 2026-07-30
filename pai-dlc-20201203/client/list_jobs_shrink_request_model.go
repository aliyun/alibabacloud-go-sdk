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
	SetDescription(v string) *ListJobsShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *ListJobsShrinkRequest
	GetDisplayName() *string
	SetDisplayNameSearchMode(v string) *ListJobsShrinkRequest
	GetDisplayNameSearchMode() *string
	SetEnableAssignNode(v string) *ListJobsShrinkRequest
	GetEnableAssignNode() *string
	SetEndTime(v string) *ListJobsShrinkRequest
	GetEndTime() *string
	SetFromAllWorkspaces(v bool) *ListJobsShrinkRequest
	GetFromAllWorkspaces() *bool
	SetImageSearch(v string) *ListJobsShrinkRequest
	GetImageSearch() *string
	SetJobId(v string) *ListJobsShrinkRequest
	GetJobId() *string
	SetJobIds(v string) *ListJobsShrinkRequest
	GetJobIds() *string
	SetJobType(v string) *ListJobsShrinkRequest
	GetJobType() *string
	SetNumericRangeField(v string) *ListJobsShrinkRequest
	GetNumericRangeField() *string
	SetNumericRangeMax(v int64) *ListJobsShrinkRequest
	GetNumericRangeMax() *int64
	SetNumericRangeMin(v int64) *ListJobsShrinkRequest
	GetNumericRangeMin() *int64
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
	SetReasonSearch(v string) *ListJobsShrinkRequest
	GetReasonSearch() *string
	SetResourceId(v string) *ListJobsShrinkRequest
	GetResourceId() *string
	SetResourceIds(v string) *ListJobsShrinkRequest
	GetResourceIds() *string
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
	SetTemplateId(v string) *ListJobsShrinkRequest
	GetTemplateId() *string
	SetTimeRangeField(v string) *ListJobsShrinkRequest
	GetTimeRangeField() *string
	SetUserCommandSearch(v string) *ListJobsShrinkRequest
	GetUserCommandSearch() *string
	SetUserIdForFilter(v string) *ListJobsShrinkRequest
	GetUserIdForFilter() *string
	SetUsername(v string) *ListJobsShrinkRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListJobsShrinkRequest
	GetWorkspaceId() *string
}

type ListJobsShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *ListJobsShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ListJobsShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListJobsShrinkRequest) GetDisplayNameSearchMode() *string {
	return s.DisplayNameSearchMode
}

func (s *ListJobsShrinkRequest) GetEnableAssignNode() *string {
	return s.EnableAssignNode
}

func (s *ListJobsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobsShrinkRequest) GetFromAllWorkspaces() *bool {
	return s.FromAllWorkspaces
}

func (s *ListJobsShrinkRequest) GetImageSearch() *string {
	return s.ImageSearch
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

func (s *ListJobsShrinkRequest) GetNumericRangeField() *string {
	return s.NumericRangeField
}

func (s *ListJobsShrinkRequest) GetNumericRangeMax() *int64 {
	return s.NumericRangeMax
}

func (s *ListJobsShrinkRequest) GetNumericRangeMin() *int64 {
	return s.NumericRangeMin
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

func (s *ListJobsShrinkRequest) GetReasonSearch() *string {
	return s.ReasonSearch
}

func (s *ListJobsShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListJobsShrinkRequest) GetResourceIds() *string {
	return s.ResourceIds
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

func (s *ListJobsShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListJobsShrinkRequest) GetTimeRangeField() *string {
	return s.TimeRangeField
}

func (s *ListJobsShrinkRequest) GetUserCommandSearch() *string {
	return s.UserCommandSearch
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

func (s *ListJobsShrinkRequest) SetDescription(v string) *ListJobsShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListJobsShrinkRequest) SetDisplayName(v string) *ListJobsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetDisplayNameSearchMode(v string) *ListJobsShrinkRequest {
	s.DisplayNameSearchMode = &v
	return s
}

func (s *ListJobsShrinkRequest) SetEnableAssignNode(v string) *ListJobsShrinkRequest {
	s.EnableAssignNode = &v
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

func (s *ListJobsShrinkRequest) SetImageSearch(v string) *ListJobsShrinkRequest {
	s.ImageSearch = &v
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

func (s *ListJobsShrinkRequest) SetNumericRangeField(v string) *ListJobsShrinkRequest {
	s.NumericRangeField = &v
	return s
}

func (s *ListJobsShrinkRequest) SetNumericRangeMax(v int64) *ListJobsShrinkRequest {
	s.NumericRangeMax = &v
	return s
}

func (s *ListJobsShrinkRequest) SetNumericRangeMin(v int64) *ListJobsShrinkRequest {
	s.NumericRangeMin = &v
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

func (s *ListJobsShrinkRequest) SetReasonSearch(v string) *ListJobsShrinkRequest {
	s.ReasonSearch = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceId(v string) *ListJobsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceIds(v string) *ListJobsShrinkRequest {
	s.ResourceIds = &v
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

func (s *ListJobsShrinkRequest) SetTemplateId(v string) *ListJobsShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetTimeRangeField(v string) *ListJobsShrinkRequest {
	s.TimeRangeField = &v
	return s
}

func (s *ListJobsShrinkRequest) SetUserCommandSearch(v string) *ListJobsShrinkRequest {
	s.UserCommandSearch = &v
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
