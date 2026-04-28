// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIDPermission interface {
	dara.Model
	String() string
	GoString() string
	SetDisinheritSubGroup(v bool) *IDPermission
	GetDisinheritSubGroup() *bool
	SetExpireTime(v int64) *IDPermission
	GetExpireTime() *int64
	SetPermission(v *Permission) *IDPermission
	GetPermission() *Permission
	SetRoles(v []*string) *IDPermission
	GetRoles() []*string
}

type IDPermission struct {
	DisinheritSubGroup *bool  `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	ExpireTime         *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// if can be null:
	// false
	Permission *Permission `json:"permission,omitempty" xml:"permission,omitempty"`
	Roles      []*string   `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s IDPermission) String() string {
	return dara.Prettify(s)
}

func (s IDPermission) GoString() string {
	return s.String()
}

func (s *IDPermission) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *IDPermission) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *IDPermission) GetPermission() *Permission {
	return s.Permission
}

func (s *IDPermission) GetRoles() []*string {
	return s.Roles
}

func (s *IDPermission) SetDisinheritSubGroup(v bool) *IDPermission {
	s.DisinheritSubGroup = &v
	return s
}

func (s *IDPermission) SetExpireTime(v int64) *IDPermission {
	s.ExpireTime = &v
	return s
}

func (s *IDPermission) SetPermission(v *Permission) *IDPermission {
	s.Permission = v
	return s
}

func (s *IDPermission) SetRoles(v []*string) *IDPermission {
	s.Roles = v
	return s
}

func (s *IDPermission) Validate() error {
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	return nil
}
