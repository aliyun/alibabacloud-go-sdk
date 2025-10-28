// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody
	GetMembers() []*ListMembersResponseBodyMembers
	SetRequestId(v string) *ListMembersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListMembersResponseBody
	GetTotalCount() *int64
}

type ListMembersResponseBody struct {
	// The members.
	Members []*ListMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of members that meet the filter conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) GetMembers() []*ListMembersResponseBodyMembers {
	return s.Members
}

func (s *ListMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMembersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMembersResponseBody) SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetTotalCount(v int64) *ListMembersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMembersResponseBodyMembers struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the user is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The member ID.
	//
	// example:
	//
	// 14588*****51688039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The username.
	//
	// example:
	//
	// user1
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The list of roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 215139******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembers) GetAccountName() *string {
	return s.AccountName
}

func (s *ListMembersResponseBodyMembers) GetAccountType() *string {
	return s.AccountType
}

func (s *ListMembersResponseBodyMembers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMembersResponseBodyMembers) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListMembersResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *ListMembersResponseBodyMembers) GetMemberName() *string {
	return s.MemberName
}

func (s *ListMembersResponseBodyMembers) GetRoles() []*string {
	return s.Roles
}

func (s *ListMembersResponseBodyMembers) GetUserId() *string {
	return s.UserId
}

func (s *ListMembersResponseBodyMembers) SetAccountName(v string) *ListMembersResponseBodyMembers {
	s.AccountName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetAccountType(v string) *ListMembersResponseBodyMembers {
	s.AccountType = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetDisplayName(v string) *ListMembersResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetGmtCreateTime(v string) *ListMembersResponseBodyMembers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberId(v string) *ListMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberName(v string) *ListMembersResponseBodyMembers {
	s.MemberName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetRoles(v []*string) *ListMembersResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserId(v string) *ListMembersResponseBodyMembers {
	s.UserId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
