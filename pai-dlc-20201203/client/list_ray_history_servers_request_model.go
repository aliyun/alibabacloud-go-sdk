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
	SetUserIdForFilter(v string) *ListRayHistoryServersRequest
	GetUserIdForFilter() *string
	SetUsername(v string) *ListRayHistoryServersRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListRayHistoryServersRequest
	GetWorkspaceId() *string
}

type ListRayHistoryServersRequest struct {
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IdPrefix *string `json:"IdPrefix,omitempty" xml:"IdPrefix,omitempty"`
	// example:
	//
	// 2020-11-09T16:00:00Z
	ModifiedAfter *string `json:"ModifiedAfter,omitempty" xml:"ModifiedAfter,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Postpaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// quotaxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123456789
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// example:
	//
	// myusername
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
