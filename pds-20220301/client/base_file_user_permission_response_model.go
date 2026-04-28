// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseFileUserPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetCanAccess(v bool) *BaseFileUserPermissionResponse
	GetCanAccess() *bool
	SetCreatedAt(v int64) *BaseFileUserPermissionResponse
	GetCreatedAt() *int64
	SetCreator(v string) *BaseFileUserPermissionResponse
	GetCreator() *string
	SetDisinheritSubGroup(v bool) *BaseFileUserPermissionResponse
	GetDisinheritSubGroup() *bool
	SetDomainId(v string) *BaseFileUserPermissionResponse
	GetDomainId() *string
	SetDriveId(v string) *BaseFileUserPermissionResponse
	GetDriveId() *string
	SetExpireTime(v int64) *BaseFileUserPermissionResponse
	GetExpireTime() *int64
	SetFileFullPath(v string) *BaseFileUserPermissionResponse
	GetFileFullPath() *string
	SetFileId(v string) *BaseFileUserPermissionResponse
	GetFileId() *string
	SetIdentity(v *Identity) *BaseFileUserPermissionResponse
	GetIdentity() *Identity
	SetRoleId(v string) *BaseFileUserPermissionResponse
	GetRoleId() *string
}

type BaseFileUserPermissionResponse struct {
	CanAccess          *bool   `json:"can_access,omitempty" xml:"can_access,omitempty"`
	CreatedAt          *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator            *string `json:"creator,omitempty" xml:"creator,omitempty"`
	DisinheritSubGroup *bool   `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	// example:
	//
	// bj23
	DomainId     *string   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId      *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ExpireTime   *int64    `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	FileFullPath *string   `json:"file_full_path,omitempty" xml:"file_full_path,omitempty"`
	FileId       *string   `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Identity     *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	RoleId       *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s BaseFileUserPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseFileUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *BaseFileUserPermissionResponse) GetCanAccess() *bool {
	return s.CanAccess
}

func (s *BaseFileUserPermissionResponse) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *BaseFileUserPermissionResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseFileUserPermissionResponse) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *BaseFileUserPermissionResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseFileUserPermissionResponse) GetDriveId() *string {
	return s.DriveId
}

func (s *BaseFileUserPermissionResponse) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *BaseFileUserPermissionResponse) GetFileFullPath() *string {
	return s.FileFullPath
}

func (s *BaseFileUserPermissionResponse) GetFileId() *string {
	return s.FileId
}

func (s *BaseFileUserPermissionResponse) GetIdentity() *Identity {
	return s.Identity
}

func (s *BaseFileUserPermissionResponse) GetRoleId() *string {
	return s.RoleId
}

func (s *BaseFileUserPermissionResponse) SetCanAccess(v bool) *BaseFileUserPermissionResponse {
	s.CanAccess = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetCreatedAt(v int64) *BaseFileUserPermissionResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetCreator(v string) *BaseFileUserPermissionResponse {
	s.Creator = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDisinheritSubGroup(v bool) *BaseFileUserPermissionResponse {
	s.DisinheritSubGroup = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDomainId(v string) *BaseFileUserPermissionResponse {
	s.DomainId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDriveId(v string) *BaseFileUserPermissionResponse {
	s.DriveId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetExpireTime(v int64) *BaseFileUserPermissionResponse {
	s.ExpireTime = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetFileFullPath(v string) *BaseFileUserPermissionResponse {
	s.FileFullPath = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetFileId(v string) *BaseFileUserPermissionResponse {
	s.FileId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetIdentity(v *Identity) *BaseFileUserPermissionResponse {
	s.Identity = v
	return s
}

func (s *BaseFileUserPermissionResponse) SetRoleId(v string) *BaseFileUserPermissionResponse {
	s.RoleId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
