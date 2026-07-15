// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayHistoryServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListRayHistoryServersRequest
	GetDisplayName() *string
	SetEndTime(v string) *ListRayHistoryServersRequest
	GetEndTime() *string
	SetIdPrefix(v string) *ListRayHistoryServersRequest
	GetIdPrefix() *string
	SetModifiedAfter(v string) *ListRayHistoryServersRequest
	GetModifiedAfter() *string
	SetOrder(v string) *ListRayHistoryServersRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListRayHistoryServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRayHistoryServersRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListRayHistoryServersRequest
	GetPaymentType() *string
	SetResourceId(v string) *ListRayHistoryServersRequest
	GetResourceId() *string
	SetShowOwn(v bool) *ListRayHistoryServersRequest
	GetShowOwn() *bool
	SetSortBy(v string) *ListRayHistoryServersRequest
	GetSortBy() *string
	SetStartTime(v string) *ListRayHistoryServersRequest
	GetStartTime() *string
	SetStatus(v string) *ListRayHistoryServersRequest
	GetStatus() *string
	SetStoragePath(v string) *ListRayHistoryServersRequest
	GetStoragePath() *string
	SetUserIdForFilter(v string) *ListRayHistoryServersRequest
	GetUserIdForFilter() *string
	SetUsername(v string) *ListRayHistoryServersRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListRayHistoryServersRequest
	GetWorkspaceId() *string
}

type ListRayHistoryServersRequest struct {
	// The display name of the job.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the query range. The job creation time is used for filtering.
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID prefix.
	//
	// example:
	//
	// 按ID前缀过滤
	IdPrefix *string `json:"IdPrefix,omitempty" xml:"IdPrefix,omitempty"`
	// Filters results by the time after which they were modified.
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	ModifiedAfter *string `json:"ModifiedAfter,omitempty" xml:"ModifiedAfter,omitempty"`
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
	// The page number of the page to return in a paged query. Paging starts from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of RayHistoryServer entries to return on each page in a paged query. Paging is used to return results in batches.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method. Valid values:
	//
	// - PrePaid
	//
	// - PostPaid.
	//
	// example:
	//
	// Postpaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The resource group ID. For information about how to query the ID of a dedicated resource group, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// quotaxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Specifies whether to return only the RayHistoryServer entries created by the current user.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The field by which to sort the returned results. Valid values:
	//
	// - DisplayName
	//
	// - GmtCreateTime
	//
	// - UserId
	//
	// - ResourceId
	//
	// - Status
	//
	// - GmtModifyTime.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The RayHistoryServer status. Valid values:
	//
	// - Creating: being created.
	//
	// - Queuing: waiting in queue.
	//
	// - Running: running.
	//
	// - Stopped: stopped.
	//
	// - Failed: failed.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage path of Ray logs.
	//
	// example:
	//
	// oss://bucket-test-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/tmp
	StoragePath *string `json:"StoragePath,omitempty" xml:"StoragePath,omitempty"`
	// Filters results by user ID.
	//
	// example:
	//
	// 123456789
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// Filters results by username.
	//
	// example:
	//
	// myusername
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html)..
	//
	// example:
	//
	// 268
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRayHistoryServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRayHistoryServersRequest) GoString() string {
	return s.String()
}

func (s *ListRayHistoryServersRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListRayHistoryServersRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRayHistoryServersRequest) GetIdPrefix() *string {
	return s.IdPrefix
}

func (s *ListRayHistoryServersRequest) GetModifiedAfter() *string {
	return s.ModifiedAfter
}

func (s *ListRayHistoryServersRequest) GetOrder() *string {
	return s.Order
}

func (s *ListRayHistoryServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRayHistoryServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRayHistoryServersRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListRayHistoryServersRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListRayHistoryServersRequest) GetShowOwn() *bool {
	return s.ShowOwn
}

func (s *ListRayHistoryServersRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRayHistoryServersRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRayHistoryServersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRayHistoryServersRequest) GetStoragePath() *string {
	return s.StoragePath
}

func (s *ListRayHistoryServersRequest) GetUserIdForFilter() *string {
	return s.UserIdForFilter
}

func (s *ListRayHistoryServersRequest) GetUsername() *string {
	return s.Username
}

func (s *ListRayHistoryServersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListRayHistoryServersRequest) SetDisplayName(v string) *ListRayHistoryServersRequest {
	s.DisplayName = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetEndTime(v string) *ListRayHistoryServersRequest {
	s.EndTime = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetIdPrefix(v string) *ListRayHistoryServersRequest {
	s.IdPrefix = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetModifiedAfter(v string) *ListRayHistoryServersRequest {
	s.ModifiedAfter = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetOrder(v string) *ListRayHistoryServersRequest {
	s.Order = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetPageNumber(v int32) *ListRayHistoryServersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetPageSize(v int32) *ListRayHistoryServersRequest {
	s.PageSize = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetPaymentType(v string) *ListRayHistoryServersRequest {
	s.PaymentType = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetResourceId(v string) *ListRayHistoryServersRequest {
	s.ResourceId = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetShowOwn(v bool) *ListRayHistoryServersRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetSortBy(v string) *ListRayHistoryServersRequest {
	s.SortBy = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetStartTime(v string) *ListRayHistoryServersRequest {
	s.StartTime = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetStatus(v string) *ListRayHistoryServersRequest {
	s.Status = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetStoragePath(v string) *ListRayHistoryServersRequest {
	s.StoragePath = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetUserIdForFilter(v string) *ListRayHistoryServersRequest {
	s.UserIdForFilter = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetUsername(v string) *ListRayHistoryServersRequest {
	s.Username = &v
	return s
}

func (s *ListRayHistoryServersRequest) SetWorkspaceId(v string) *ListRayHistoryServersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListRayHistoryServersRequest) Validate() error {
	return dara.Validate(s)
}
