// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMembership interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Membership
	GetCreatedAt() *int64
	SetCreator(v string) *Membership
	GetCreator() *string
	SetDescription(v string) *Membership
	GetDescription() *string
	SetDomainId(v string) *Membership
	GetDomainId() *string
	SetGroupId(v string) *Membership
	GetGroupId() *string
	SetMemberRole(v string) *Membership
	GetMemberRole() *string
	SetMemberType(v string) *Membership
	GetMemberType() *string
	SetSubGroupId(v string) *Membership
	GetSubGroupId() *string
	SetUpdatedAt(v int64) *Membership
	GetUpdatedAt() *int64
	SetUserId(v string) *Membership
	GetUserId() *string
}

type Membership struct {
	CreatedAt   *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	GroupId     *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	MemberRole  *string `json:"member_role,omitempty" xml:"member_role,omitempty"`
	MemberType  *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
	SubGroupId  *string `json:"sub_group_id,omitempty" xml:"sub_group_id,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s Membership) String() string {
	return dara.Prettify(s)
}

func (s Membership) GoString() string {
	return s.String()
}

func (s *Membership) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Membership) GetCreator() *string {
	return s.Creator
}

func (s *Membership) GetDescription() *string {
	return s.Description
}

func (s *Membership) GetDomainId() *string {
	return s.DomainId
}

func (s *Membership) GetGroupId() *string {
	return s.GroupId
}

func (s *Membership) GetMemberRole() *string {
	return s.MemberRole
}

func (s *Membership) GetMemberType() *string {
	return s.MemberType
}

func (s *Membership) GetSubGroupId() *string {
	return s.SubGroupId
}

func (s *Membership) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Membership) GetUserId() *string {
	return s.UserId
}

func (s *Membership) SetCreatedAt(v int64) *Membership {
	s.CreatedAt = &v
	return s
}

func (s *Membership) SetCreator(v string) *Membership {
	s.Creator = &v
	return s
}

func (s *Membership) SetDescription(v string) *Membership {
	s.Description = &v
	return s
}

func (s *Membership) SetDomainId(v string) *Membership {
	s.DomainId = &v
	return s
}

func (s *Membership) SetGroupId(v string) *Membership {
	s.GroupId = &v
	return s
}

func (s *Membership) SetMemberRole(v string) *Membership {
	s.MemberRole = &v
	return s
}

func (s *Membership) SetMemberType(v string) *Membership {
	s.MemberType = &v
	return s
}

func (s *Membership) SetSubGroupId(v string) *Membership {
	s.SubGroupId = &v
	return s
}

func (s *Membership) SetUpdatedAt(v int64) *Membership {
	s.UpdatedAt = &v
	return s
}

func (s *Membership) SetUserId(v string) *Membership {
	s.UserId = &v
	return s
}

func (s *Membership) Validate() error {
	return dara.Validate(s)
}
