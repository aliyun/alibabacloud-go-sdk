// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroup interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Group
	GetCreatedAt() *int64
	SetCreator(v string) *Group
	GetCreator() *string
	SetDescription(v string) *Group
	GetDescription() *string
	SetDomainId(v string) *Group
	GetDomainId() *string
	SetGroupId(v string) *Group
	GetGroupId() *string
	SetGroupName(v string) *Group
	GetGroupName() *string
	SetUpdatedAt(v int64) *Group
	GetUpdatedAt() *int64
}

type Group struct {
	// The time when the group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The ID of the user who created the group.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description of the group.
	//
	// example:
	//
	// created by system
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the domain.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// b45c0c0c373c41ec9ebb5c85a025a08f
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// test group
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// The time when the group was modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Group) String() string {
	return dara.Prettify(s)
}

func (s Group) GoString() string {
	return s.String()
}

func (s *Group) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Group) GetCreator() *string {
	return s.Creator
}

func (s *Group) GetDescription() *string {
	return s.Description
}

func (s *Group) GetDomainId() *string {
	return s.DomainId
}

func (s *Group) GetGroupId() *string {
	return s.GroupId
}

func (s *Group) GetGroupName() *string {
	return s.GroupName
}

func (s *Group) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Group) SetCreatedAt(v int64) *Group {
	s.CreatedAt = &v
	return s
}

func (s *Group) SetCreator(v string) *Group {
	s.Creator = &v
	return s
}

func (s *Group) SetDescription(v string) *Group {
	s.Description = &v
	return s
}

func (s *Group) SetDomainId(v string) *Group {
	s.DomainId = &v
	return s
}

func (s *Group) SetGroupId(v string) *Group {
	s.GroupId = &v
	return s
}

func (s *Group) SetGroupName(v string) *Group {
	s.GroupName = &v
	return s
}

func (s *Group) SetUpdatedAt(v int64) *Group {
	s.UpdatedAt = &v
	return s
}

func (s *Group) Validate() error {
	return dara.Validate(s)
}
