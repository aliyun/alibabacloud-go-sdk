// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemberName(v string) *ListMembersRequest
	GetMemberName() *string
	SetPageNumber(v int64) *ListMembersRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *ListMembersRequest
	GetPageSize() *int32
	SetRoles(v string) *ListMembersRequest
	GetRoles() *string
	SetUserId(v string) *ListMembersRequest
	GetUserId() *string
}

type ListMembersRequest struct {
	// The member name. Fuzzy match is supported.
	//
	// example:
	//
	// zhangsan
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The page number of the workspace list. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The paging Settings for the number of entries per page in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters members by role. Separate multiple roles with commas (,).
	//
	// Valid values:
	//
	// - PAI.AlgoDeveloper: Algorithm developer.
	//
	// - PAI.AlgoOperator: Algorithm O&M.
	//
	// - PAI.LabelManager: Labeling administrator.
	//
	// - PAI.MaxComputeDeveloper: MaxCompute developer.
	//
	// - PAI.WorkspaceAdmin: Administrator.
	//
	// - PAI.WorkspaceGuest: Guest.
	//
	// - PAI.WorkspaceOwner: Owner.
	//
	// example:
	//
	// PAI.AlgoDeveloper
	Roles *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
	// The ID of the user to query.
	//
	// example:
	//
	// 2788******129
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *ListMembersRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMembersRequest) GetRoles() *string {
	return s.Roles
}

func (s *ListMembersRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListMembersRequest) SetMemberName(v string) *ListMembersRequest {
	s.MemberName = &v
	return s
}

func (s *ListMembersRequest) SetPageNumber(v int64) *ListMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMembersRequest) SetPageSize(v int32) *ListMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMembersRequest) SetRoles(v string) *ListMembersRequest {
	s.Roles = &v
	return s
}

func (s *ListMembersRequest) SetUserId(v string) *ListMembersRequest {
	s.UserId = &v
	return s
}

func (s *ListMembersRequest) Validate() error {
	return dara.Validate(s)
}
