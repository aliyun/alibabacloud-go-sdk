// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilePermissionMember interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*string) *FilePermissionMember
	GetActionList() []*string
	SetDisinheritSubGroup(v bool) *FilePermissionMember
	GetDisinheritSubGroup() *bool
	SetExpireTime(v int64) *FilePermissionMember
	GetExpireTime() *int64
	SetIdentity(v *Identity) *FilePermissionMember
	GetIdentity() *Identity
	SetRoleId(v string) *FilePermissionMember
	GetRoleId() *string
}

type FilePermissionMember struct {
	// The list of permissions to grant. You can grant permissions by assigning roles to identities, or you can customize the permissions. To grant permissions by assigning roles to identities, specify role_id. role_id and action_list are mutually exclusive. If both parameters are specified, the value of role_id prevails. When you specify action_list, the system automatically generates a temporary role_id. You can use this role_id to revoke the permissions.
	ActionList []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// Specifies whether the users of subgroups can inherit the permissions. For example, a user named user1 belongs to the group1 group, and a user named user2 belongs to the group2 group. group2 is the subgroup of group1. If you set disinherit_sub_group to true, only user1 is granted the permissions. user2 is not granted the permissions.
	//
	// example:
	//
	// false
	DisinheritSubGroup *bool `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	// The time when the permissions expire. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. A value of 4775500800000 indicates that the permissions never expire.
	//
	// example:
	//
	// 1633598085642
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The identity to which the permissions are granted, which is a user or a group.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The role ID. You can grant permissions by assigning roles to identities, or you can customize the permissions. To grant permissions by assigning roles to identities, specify role_id. role_id and action_list are mutually exclusive. If both parameters are specified, the value of role_id prevails.
	//
	// Valid values:
	//
	// SystemFileOwner: collaborator
	//
	// SystemFileDownloader: downloader
	//
	// SystemFileEditor: editor
	//
	// SystemFileEditorWithoutDelete: editor without permissions to delete the file
	//
	// SystemFileEditorWithoutShareLink: editor without permissions to share the file
	//
	// SystemFileMetaViewer: viewer of lists
	//
	// SystemFileUploader: uploader. SystemFileUploaderAndDownloader: uploader and downloader
	//
	// SystemFileDownloaderWithShareLink: downloader and sharer
	//
	// SystemFileUploaderAndDownloaderWithShareLink: uploader, downloader, and sharer
	//
	// SystemFileUploaderAndViewer: viewer and uploader
	//
	// SystemFileUploaderWithShareLink: uploader and sharer
	//
	// SystemFileViewer: viewer
	//
	// example:
	//
	// SystemFileDownloader
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s FilePermissionMember) String() string {
	return dara.Prettify(s)
}

func (s FilePermissionMember) GoString() string {
	return s.String()
}

func (s *FilePermissionMember) GetActionList() []*string {
	return s.ActionList
}

func (s *FilePermissionMember) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *FilePermissionMember) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *FilePermissionMember) GetIdentity() *Identity {
	return s.Identity
}

func (s *FilePermissionMember) GetRoleId() *string {
	return s.RoleId
}

func (s *FilePermissionMember) SetActionList(v []*string) *FilePermissionMember {
	s.ActionList = v
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

func (s *FilePermissionMember) SetIdentity(v *Identity) *FilePermissionMember {
	s.Identity = v
	return s
}

func (s *FilePermissionMember) SetRoleId(v string) *FilePermissionMember {
	s.RoleId = &v
	return s
}

func (s *FilePermissionMember) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
