// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeInstanceMembersRequest
	GetCurrentPage() *string
	SetMemberDesc(v string) *DescribeInstanceMembersRequest
	GetMemberDesc() *string
	SetMemberDisplayName(v string) *DescribeInstanceMembersRequest
	GetMemberDisplayName() *string
	SetMemberUid(v string) *DescribeInstanceMembersRequest
	GetMemberUid() *string
	SetPageSize(v string) *DescribeInstanceMembersRequest
	GetPageSize() *string
}

type DescribeInstanceMembersRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The remarks of the member. The remarks must be 1 to 256 characters in length.
	//
	// example:
	//
	// renewal
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// cloudfirewall_2
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// The UID of the member.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries per page.
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeInstanceMembersRequest) GetMemberDesc() *string {
	return s.MemberDesc
}

func (s *DescribeInstanceMembersRequest) GetMemberDisplayName() *string {
	return s.MemberDisplayName
}

func (s *DescribeInstanceMembersRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeInstanceMembersRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstanceMembersRequest) SetCurrentPage(v string) *DescribeInstanceMembersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDesc(v string) *DescribeInstanceMembersRequest {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDisplayName(v string) *DescribeInstanceMembersRequest {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberUid(v string) *DescribeInstanceMembersRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetPageSize(v string) *DescribeInstanceMembersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMembersRequest) Validate() error {
	return dara.Validate(s)
}
