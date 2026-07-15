// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTensorboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListTensorboardsRequest
	GetAccessibility() *string
	SetDisplayName(v string) *ListTensorboardsRequest
	GetDisplayName() *string
	SetEndTime(v string) *ListTensorboardsRequest
	GetEndTime() *string
	SetJobId(v string) *ListTensorboardsRequest
	GetJobId() *string
	SetOrder(v string) *ListTensorboardsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTensorboardsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTensorboardsRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListTensorboardsRequest
	GetPaymentType() *string
	SetQuotaId(v string) *ListTensorboardsRequest
	GetQuotaId() *string
	SetShowOwn(v bool) *ListTensorboardsRequest
	GetShowOwn() *bool
	SetSortBy(v string) *ListTensorboardsRequest
	GetSortBy() *string
	SetSourceId(v string) *ListTensorboardsRequest
	GetSourceId() *string
	SetSourceType(v string) *ListTensorboardsRequest
	GetSourceType() *string
	SetStartTime(v string) *ListTensorboardsRequest
	GetStartTime() *string
	SetStatus(v string) *ListTensorboardsRequest
	GetStatus() *string
	SetTensorboardId(v string) *ListTensorboardsRequest
	GetTensorboardId() *string
	SetUserId(v string) *ListTensorboardsRequest
	GetUserId() *string
	SetUsername(v string) *ListTensorboardsRequest
	GetUsername() *string
	SetVerbose(v bool) *ListTensorboardsRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListTensorboardsRequest
	GetWorkspaceId() *string
}

type ListTensorboardsRequest struct {
	// The visibility of the Tensorboard instance. Valid values:
	//
	// - PUBLIC: visible to all members in the workspace.
	//
	// - PRIVATE: visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The display name of the Tensorboard instance.
	//
	// example:
	//
	// TestTensorboard
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the query range. Tensorboard instances are filtered by their creation time in UTC. If this parameter is left empty, the default value is the current time.
	//
	// example:
	//
	// 2020-11-09T14:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The DLC job ID used to filter Tensorboard instances. Call [ListJobs](https://help.aliyun.com/document_detail/459676.html) to obtain the job ID.
	//
	// example:
	//
	// dlc-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The sort order. Valid values:
	//
	// - desc: descending order.
	//
	// - asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number for paging. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of Tensorboard instances to return on each page for paging.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the Tensorboard instance. Valid values:
	//
	// - Free: a Tensorboard instance that uses free resources.
	//
	// - Postpaid: a Tensorboard instance that uses pay-as-you-go resources.
	//
	// example:
	//
	// Postpaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The resource quota ID.
	//
	// > - Currently, only whitelisted users can use resource quota resources to create Tensorboard instances. To use this feature, contact us.
	//
	// > - This parameter takes effect only when the Tensorboard instance uses resource quota resources.
	//
	// example:
	//
	// quota12***
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// Specifies whether to return only Tensorboard instances created by the current user.
	//
	// example:
	//
	// false
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The field by which the results are sorted. Valid values:
	//
	// - DisplayName: the job name.
	//
	// - GmtCreateTime: the job creation time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The data source ID. Call [ListJobs](https://help.aliyun.com/document_detail/459676.html) to obtain the job ID.
	//
	// example:
	//
	// dlc-xxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The data source type. Currently, only DLC training jobs are supported, which means the value is job.
	//
	// example:
	//
	// job
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the query range. Tensorboard instances are filtered by their creation time in UTC. If this parameter is left empty, the default value is 7 days before the current time.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The Tensorboard status. Valid values:
	//
	// - Creating: being created.
	//
	// - Running: running.
	//
	// - Stopped: stopped.
	//
	// - Succeeded: succeeded.
	//
	// - Failed: failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Tensorboard ID used to filter the Tensorboard list.
	//
	// example:
	//
	// tensorboard-xxx
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 161****3000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// she****mo
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Specifies whether to display details.
	//
	// - true: Display details.
	//
	// - false: Do not display details.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The workspace ID. Tensorboard instances are filtered by workspace ID. <props="china">Call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID..
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTensorboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTensorboardsRequest) GoString() string {
	return s.String()
}

func (s *ListTensorboardsRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListTensorboardsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListTensorboardsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTensorboardsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListTensorboardsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTensorboardsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTensorboardsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTensorboardsRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListTensorboardsRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListTensorboardsRequest) GetShowOwn() *bool {
	return s.ShowOwn
}

func (s *ListTensorboardsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTensorboardsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListTensorboardsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTensorboardsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTensorboardsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTensorboardsRequest) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *ListTensorboardsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListTensorboardsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListTensorboardsRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListTensorboardsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTensorboardsRequest) SetAccessibility(v string) *ListTensorboardsRequest {
	s.Accessibility = &v
	return s
}

func (s *ListTensorboardsRequest) SetDisplayName(v string) *ListTensorboardsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListTensorboardsRequest) SetEndTime(v string) *ListTensorboardsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetJobId(v string) *ListTensorboardsRequest {
	s.JobId = &v
	return s
}

func (s *ListTensorboardsRequest) SetOrder(v string) *ListTensorboardsRequest {
	s.Order = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageNumber(v int32) *ListTensorboardsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageSize(v int32) *ListTensorboardsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTensorboardsRequest) SetPaymentType(v string) *ListTensorboardsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListTensorboardsRequest) SetQuotaId(v string) *ListTensorboardsRequest {
	s.QuotaId = &v
	return s
}

func (s *ListTensorboardsRequest) SetShowOwn(v bool) *ListTensorboardsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListTensorboardsRequest) SetSortBy(v string) *ListTensorboardsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTensorboardsRequest) SetSourceId(v string) *ListTensorboardsRequest {
	s.SourceId = &v
	return s
}

func (s *ListTensorboardsRequest) SetSourceType(v string) *ListTensorboardsRequest {
	s.SourceType = &v
	return s
}

func (s *ListTensorboardsRequest) SetStartTime(v string) *ListTensorboardsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetStatus(v string) *ListTensorboardsRequest {
	s.Status = &v
	return s
}

func (s *ListTensorboardsRequest) SetTensorboardId(v string) *ListTensorboardsRequest {
	s.TensorboardId = &v
	return s
}

func (s *ListTensorboardsRequest) SetUserId(v string) *ListTensorboardsRequest {
	s.UserId = &v
	return s
}

func (s *ListTensorboardsRequest) SetUsername(v string) *ListTensorboardsRequest {
	s.Username = &v
	return s
}

func (s *ListTensorboardsRequest) SetVerbose(v bool) *ListTensorboardsRequest {
	s.Verbose = &v
	return s
}

func (s *ListTensorboardsRequest) SetWorkspaceId(v string) *ListTensorboardsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTensorboardsRequest) Validate() error {
	return dara.Validate(s)
}
