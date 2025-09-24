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
}

type ListMembersRequest struct {
	// The member name. Fuzzy match is supported.
	//
	// example:
	//
	// zhangsan
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The page number of the workspace list. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The roles that are used to filter members. Multiple roles are separated by commas (,). Valid values:
	//
	// 	- PAI.AlgoDeveloper: algorithm developer
	//
	// 	- PAI.AlgoOperator: algorithm O\\&M engineer
	//
	// 	- PAI.LabelManager: labeling administrator
	//
	// 	- PAI.MaxComputeDeveloper: MaxCompute developer
	//
	// 	- PAI.WorkspaceAdmin: administrator
	//
	// 	- PAI.WorkspaceGuest: guest
	//
	// 	- PAI.WorkspaceOwner: owner
	//
	// example:
	//
	// PAI.AlgoDeveloper
	Roles *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
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

func (s *ListMembersRequest) Validate() error {
	return dara.Validate(s)
}
