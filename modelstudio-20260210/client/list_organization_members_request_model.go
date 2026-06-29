// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasSeat(v bool) *ListOrganizationMembersRequest
	GetHasSeat() *bool
	SetName(v string) *ListOrganizationMembersRequest
	GetName() *string
	SetPageNum(v int32) *ListOrganizationMembersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListOrganizationMembersRequest
	GetPageSize() *int32
	SetStatus(v string) *ListOrganizationMembersRequest
	GetStatus() *string
}

type ListOrganizationMembersRequest struct {
	// Specifies whether to filter by seat assignment.
	//
	// example:
	//
	// true
	HasSeat *bool `json:"HasSeat,omitempty" xml:"HasSeat,omitempty"`
	// Fuzzy filter by member name. Matches accountName or email and is case-insensitive.
	//
	// example:
	//
	// 成员名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by member status, such as ACTIVE or FROZEN. Set to null to disable filtering.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListOrganizationMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersRequest) GetHasSeat() *bool {
	return s.HasSeat
}

func (s *ListOrganizationMembersRequest) GetName() *string {
	return s.Name
}

func (s *ListOrganizationMembersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListOrganizationMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrganizationMembersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListOrganizationMembersRequest) SetHasSeat(v bool) *ListOrganizationMembersRequest {
	s.HasSeat = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetName(v string) *ListOrganizationMembersRequest {
	s.Name = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetPageNum(v int32) *ListOrganizationMembersRequest {
	s.PageNum = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetPageSize(v int32) *ListOrganizationMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetStatus(v string) *ListOrganizationMembersRequest {
	s.Status = &v
	return s
}

func (s *ListOrganizationMembersRequest) Validate() error {
	return dara.Validate(s)
}
