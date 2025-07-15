// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilePermissionMember interface {
	dara.Model
	String() string
	GoString() string
	SetCdsIdentity(v *FilePermissionMemberCdsIdentity) *FilePermissionMember
	GetCdsIdentity() *FilePermissionMemberCdsIdentity
	SetDisinheritSubGroup(v bool) *FilePermissionMember
	GetDisinheritSubGroup() *bool
	SetExpireTime(v int64) *FilePermissionMember
	GetExpireTime() *int64
	SetRoleId(v string) *FilePermissionMember
	GetRoleId() *string
}

type FilePermissionMember struct {
	// This parameter is required.
	CdsIdentity        *FilePermissionMemberCdsIdentity `json:"CdsIdentity,omitempty" xml:"CdsIdentity,omitempty" type:"Struct"`
	DisinheritSubGroup *bool                            `json:"DisinheritSubGroup,omitempty" xml:"DisinheritSubGroup,omitempty"`
	ExpireTime         *int64                           `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s FilePermissionMember) String() string {
	return dara.Prettify(s)
}

func (s FilePermissionMember) GoString() string {
	return s.String()
}

func (s *FilePermissionMember) GetCdsIdentity() *FilePermissionMemberCdsIdentity {
	return s.CdsIdentity
}

func (s *FilePermissionMember) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *FilePermissionMember) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *FilePermissionMember) GetRoleId() *string {
	return s.RoleId
}

func (s *FilePermissionMember) SetCdsIdentity(v *FilePermissionMemberCdsIdentity) *FilePermissionMember {
	s.CdsIdentity = v
	return s
}

func (s *FilePermissionMember) SetDisinheritSubGroup(v bool) *FilePermissionMember {
	s.DisinheritSubGroup = &v
	return s
}

func (s *FilePermissionMember) SetExpireTime(v int64) *FilePermissionMember {
	s.ExpireTime = &v
	return s
}

func (s *FilePermissionMember) SetRoleId(v string) *FilePermissionMember {
	s.RoleId = &v
	return s
}

func (s *FilePermissionMember) Validate() error {
	return dara.Validate(s)
}

type FilePermissionMemberCdsIdentity struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FilePermissionMemberCdsIdentity) String() string {
	return dara.Prettify(s)
}

func (s FilePermissionMemberCdsIdentity) GoString() string {
	return s.String()
}

func (s *FilePermissionMemberCdsIdentity) GetId() *string {
	return s.Id
}

func (s *FilePermissionMemberCdsIdentity) GetType() *string {
	return s.Type
}

func (s *FilePermissionMemberCdsIdentity) SetId(v string) *FilePermissionMemberCdsIdentity {
	s.Id = &v
	return s
}

func (s *FilePermissionMemberCdsIdentity) SetType(v string) *FilePermissionMemberCdsIdentity {
	s.Type = &v
	return s
}

func (s *FilePermissionMemberCdsIdentity) Validate() error {
	return dara.Validate(s)
}
