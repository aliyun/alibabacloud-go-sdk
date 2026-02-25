// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaActiveUserUsagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListQuotaActiveUserUsagesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListQuotaActiveUserUsagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQuotaActiveUserUsagesRequest
	GetPageSize() *int32
	SetSelfOnly(v bool) *ListQuotaActiveUserUsagesRequest
	GetSelfOnly() *bool
	SetSortBy(v string) *ListQuotaActiveUserUsagesRequest
	GetSortBy() *string
	SetUserId(v string) *ListQuotaActiveUserUsagesRequest
	GetUserId() *string
	SetUsername(v string) *ListQuotaActiveUserUsagesRequest
	GetUsername() *string
	SetWorkspaceId(v string) *ListQuotaActiveUserUsagesRequest
	GetWorkspaceId() *string
}

type ListQuotaActiveUserUsagesRequest struct {
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 999
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	SelfOnly *bool `json:"SelfOnly,omitempty" xml:"SelfOnly,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 200xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListQuotaActiveUserUsagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaActiveUserUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaActiveUserUsagesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListQuotaActiveUserUsagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQuotaActiveUserUsagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQuotaActiveUserUsagesRequest) GetSelfOnly() *bool {
	return s.SelfOnly
}

func (s *ListQuotaActiveUserUsagesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListQuotaActiveUserUsagesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListQuotaActiveUserUsagesRequest) GetUsername() *string {
	return s.Username
}

func (s *ListQuotaActiveUserUsagesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListQuotaActiveUserUsagesRequest) SetOrder(v string) *ListQuotaActiveUserUsagesRequest {
	s.Order = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetPageNumber(v int32) *ListQuotaActiveUserUsagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetPageSize(v int32) *ListQuotaActiveUserUsagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetSelfOnly(v bool) *ListQuotaActiveUserUsagesRequest {
	s.SelfOnly = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetSortBy(v string) *ListQuotaActiveUserUsagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetUserId(v string) *ListQuotaActiveUserUsagesRequest {
	s.UserId = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetUsername(v string) *ListQuotaActiveUserUsagesRequest {
	s.Username = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) SetWorkspaceId(v string) *ListQuotaActiveUserUsagesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListQuotaActiveUserUsagesRequest) Validate() error {
	return dara.Validate(s)
}
