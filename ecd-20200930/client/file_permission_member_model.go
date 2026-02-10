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
	// The object that you want to grant permissions. The object can be a user or a group.
	//
	// This parameter is required.
	CdsIdentity *FilePermissionMemberCdsIdentity `json:"CdsIdentity,omitempty" xml:"CdsIdentity,omitempty" type:"Struct"`
	// Indicates whether to disable the permission from users in the subgroup.
	//
	// example:
	//
	// true
	DisinheritSubGroup *bool `json:"DisinheritSubGroup,omitempty" xml:"DisinheritSubGroup,omitempty"`
	// The expiration time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC
	//
	// example:
	//
	// 1633598866642
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The role.
	//
	// Valid values:
	//
	// 	- SystemFileEditorWithoutShareLink: The role that can edit but cannot share files.
	//
	// 	- SystemFileUploaderAndDownloaderWithShareLink: The role that can upload, download, and share files.
	//
	// 	- SystemFileDownloader: The role that can download files.
	//
	// 	- SystemFileEditorWithoutDelete: The role that can edit but cannot edit files.
	//
	// 	- SystemFileOwner: The role that can collaborate with others on files.
	//
	// 	- SystemFileDownloaderWithShareLink: The role that can download and share files.
	//
	// 	- SystemFileUploaderAndViewer: The role that can preview and upload files.
	//
	// 	- SystemFileViewer: The role that can preview files.
	//
	// 	- SystemFileEditor: The role that can edit files.
	//
	// 	- SystemFileUploaderWithShareLink: The role that can upload and share files.
	//
	// 	- SystemFileUploader: The role that can upload files.
	//
	// 	- SystemFileUploaderAndDownloader: The role that can upload and download files.
	//
	// 	- SystemFileMetaViewer: The role that can view file list.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemFileEditor
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
	if s.CdsIdentity != nil {
		if err := s.CdsIdentity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FilePermissionMemberCdsIdentity struct {
	// The user ID or a team ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16435bdf934248b788b7b3771ee9****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The object type.
	//
	// Valid values:
	//
	// 	- IT_Group: team
	//
	// 	- IT_User: user
	//
	// This parameter is required.
	//
	// example:
	//
	// IT_User
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
