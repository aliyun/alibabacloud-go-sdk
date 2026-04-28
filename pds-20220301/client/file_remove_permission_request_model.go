// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileRemovePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FileRemovePermissionRequest
	GetDriveId() *string
	SetFileId(v string) *FileRemovePermissionRequest
	GetFileId() *string
	SetMemberList(v []*FileRemovePermissionRequestMemberList) *FileRemovePermissionRequest
	GetMemberList() []*FileRemovePermissionRequestMemberList
}

type FileRemovePermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The identities with whom the file is shared.
	//
	// This parameter is required.
	MemberList []*FileRemovePermissionRequestMemberList `json:"member_list,omitempty" xml:"member_list,omitempty" type:"Repeated"`
}

func (s FileRemovePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s FileRemovePermissionRequest) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *FileRemovePermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *FileRemovePermissionRequest) GetMemberList() []*FileRemovePermissionRequestMemberList {
	return s.MemberList
}

func (s *FileRemovePermissionRequest) SetDriveId(v string) *FileRemovePermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileRemovePermissionRequest) SetFileId(v string) *FileRemovePermissionRequest {
	s.FileId = &v
	return s
}

func (s *FileRemovePermissionRequest) SetMemberList(v []*FileRemovePermissionRequestMemberList) *FileRemovePermissionRequest {
	s.MemberList = v
	return s
}

func (s *FileRemovePermissionRequest) Validate() error {
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FileRemovePermissionRequestMemberList struct {
	// The identity to whom the permissions are granted, which is a user or a group.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The role ID. You can grant permissions by assigning roles to identities, or you can customize the permissions. To grant permissions by assigning roles to identities, specify role_id. role_id and action_list are mutually exclusive. If both parameters are specified, role_id has a higher priority.
	//
	// Valid values:
	//
	// SystemFileOwner: collaborator.
	//
	// SystemFileDownloader: downloader.
	//
	// SystemFileEditor: editor.
	//
	// SystemFileEditorWithoutDelete: editor without permissions to delete the file.
	//
	// SystemFileEditorWithoutShareLink: editor without permissions to share the file.
	//
	// SystemFileMetaViewer: viewer of lists.
	//
	// SystemFileUploader: uploader. SystemFileUploaderAndDownloader: uploader and downloader.
	//
	// SystemFileDownloaderWithShareLink: downloader and sharer.
	//
	// SystemFileUploaderAndDownloaderWithShareLink: uploader, downloader, and sharer.
	//
	// SystemFileUploaderAndViewer: viewer and uploader.
	//
	// SystemFileUploaderWithShareLink: uploader and sharer.
	//
	// SystemFileViewer: viewer.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemFileDownloader
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s FileRemovePermissionRequestMemberList) String() string {
	return dara.Prettify(s)
}

func (s FileRemovePermissionRequestMemberList) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionRequestMemberList) GetIdentity() *Identity {
	return s.Identity
}

func (s *FileRemovePermissionRequestMemberList) GetRoleId() *string {
	return s.RoleId
}

func (s *FileRemovePermissionRequestMemberList) SetIdentity(v *Identity) *FileRemovePermissionRequestMemberList {
	s.Identity = v
	return s
}

func (s *FileRemovePermissionRequestMemberList) SetRoleId(v string) *FileRemovePermissionRequestMemberList {
	s.RoleId = &v
	return s
}

func (s *FileRemovePermissionRequestMemberList) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
