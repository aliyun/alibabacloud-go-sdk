// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *GetMemberResponseBody
	GetAccountType() *string
	SetDisplayName(v string) *GetMemberResponseBody
	GetDisplayName() *string
	SetGmtCreateTime(v string) *GetMemberResponseBody
	GetGmtCreateTime() *string
	SetMemberId(v string) *GetMemberResponseBody
	GetMemberId() *string
	SetMemberName(v string) *GetMemberResponseBody
	GetMemberName() *string
	SetRequestId(v string) *GetMemberResponseBody
	GetRequestId() *string
	SetRoles(v []*string) *GetMemberResponseBody
	GetRoles() []*string
	SetUserId(v string) *GetMemberResponseBody
	GetUserId() *string
}

type GetMemberResponseBody struct {
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The display name of the member.
	//
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the workspace is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The member ID.
	//
	// example:
	//
	// 145883-21513926******88039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The username.
	//
	// example:
	//
	// user1
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of roles.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberResponseBody) GetAccountType() *string {
	return s.AccountType
}

func (s *GetMemberResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetMemberResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetMemberResponseBody) GetMemberId() *string {
	return s.MemberId
}

func (s *GetMemberResponseBody) GetMemberName() *string {
	return s.MemberName
}

func (s *GetMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemberResponseBody) GetRoles() []*string {
	return s.Roles
}

func (s *GetMemberResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetMemberResponseBody) SetAccountType(v string) *GetMemberResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetMemberResponseBody) SetDisplayName(v string) *GetMemberResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetMemberResponseBody) SetGmtCreateTime(v string) *GetMemberResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMemberResponseBody) SetMemberId(v string) *GetMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *GetMemberResponseBody) SetMemberName(v string) *GetMemberResponseBody {
	s.MemberName = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberResponseBody) SetRoles(v []*string) *GetMemberResponseBody {
	s.Roles = v
	return s
}

func (s *GetMemberResponseBody) SetUserId(v string) *GetMemberResponseBody {
	s.UserId = &v
	return s
}

func (s *GetMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
