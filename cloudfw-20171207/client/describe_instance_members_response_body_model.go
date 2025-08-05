// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*DescribeInstanceMembersResponseBodyMembers) *DescribeInstanceMembersResponseBody
	GetMembers() []*DescribeInstanceMembersResponseBodyMembers
	SetPageInfo(v *DescribeInstanceMembersResponseBodyPageInfo) *DescribeInstanceMembersResponseBody
	GetPageInfo() *DescribeInstanceMembersResponseBodyPageInfo
	SetRequestId(v string) *DescribeInstanceMembersResponseBody
	GetRequestId() *string
}

type DescribeInstanceMembersResponseBody struct {
	// The information about the member.
	Members []*DescribeInstanceMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeInstanceMembersResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A531AE1A-FBA2-48B6-BAB8-84D02BD409EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBody) GetMembers() []*DescribeInstanceMembersResponseBodyMembers {
	return s.Members
}

func (s *DescribeInstanceMembersResponseBody) GetPageInfo() *DescribeInstanceMembersResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeInstanceMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMembersResponseBody) SetMembers(v []*DescribeInstanceMembersResponseBodyMembers) *DescribeInstanceMembersResponseBody {
	s.Members = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetPageInfo(v *DescribeInstanceMembersResponseBodyPageInfo) *DescribeInstanceMembersResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetRequestId(v string) *DescribeInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMembersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceMembersResponseBodyMembers struct {
	// The time when the member was added to Cloud Firewall. The value is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615189819
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The remarks of the member.
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
	// The status of the member. Valid values:
	//
	// 	- **normal**
	//
	// 	- **deleting**
	//
	// example:
	//
	// normal
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// The UID of the member.
	//
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the member was last modified. The value is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1615189819
	ModifyTime *int32 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetMemberDesc() *string {
	return s.MemberDesc
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetMemberDisplayName() *string {
	return s.MemberDisplayName
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetMemberStatus() *string {
	return s.MemberStatus
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeInstanceMembersResponseBodyMembers) GetModifyTime() *int32 {
	return s.ModifyTime
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetCreateTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDesc(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDisplayName(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberStatus(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberStatus = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberUid(v int64) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetModifyTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.ModifyTime = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceMembersResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of the members.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetPageSize(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
