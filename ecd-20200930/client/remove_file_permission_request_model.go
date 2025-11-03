// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFilePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *RemoveFilePermissionRequest
	GetCdsId() *string
	SetEndUserId(v string) *RemoveFilePermissionRequest
	GetEndUserId() *string
	SetFileId(v string) *RemoveFilePermissionRequest
	GetFileId() *string
	SetGroupId(v string) *RemoveFilePermissionRequest
	GetGroupId() *string
	SetMemberList(v []*RemoveFilePermissionRequestMemberList) *RemoveFilePermissionRequest
	GetMemberList() []*RemoveFilePermissionRequestMemberList
	SetRegionId(v string) *RemoveFilePermissionRequest
	GetRegionId() *string
}

type RemoveFilePermissionRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-066224****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) operation to query the ID of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6333e553a133ce21e6f747cf948bb9ef95d7****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the team space.
	//
	// example:
	//
	// cg-1fbmvrc7ug5m7****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The users that you want to authorize to use the cloud disk.
	//
	// This parameter is required.
	MemberList []*RemoveFilePermissionRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveFilePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionRequest) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *RemoveFilePermissionRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RemoveFilePermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *RemoveFilePermissionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveFilePermissionRequest) GetMemberList() []*RemoveFilePermissionRequestMemberList {
	return s.MemberList
}

func (s *RemoveFilePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveFilePermissionRequest) SetCdsId(v string) *RemoveFilePermissionRequest {
	s.CdsId = &v
	return s
}

func (s *RemoveFilePermissionRequest) SetEndUserId(v string) *RemoveFilePermissionRequest {
	s.EndUserId = &v
	return s
}

func (s *RemoveFilePermissionRequest) SetFileId(v string) *RemoveFilePermissionRequest {
	s.FileId = &v
	return s
}

func (s *RemoveFilePermissionRequest) SetGroupId(v string) *RemoveFilePermissionRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveFilePermissionRequest) SetMemberList(v []*RemoveFilePermissionRequestMemberList) *RemoveFilePermissionRequest {
	s.MemberList = v
	return s
}

func (s *RemoveFilePermissionRequest) SetRegionId(v string) *RemoveFilePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveFilePermissionRequest) Validate() error {
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

type RemoveFilePermissionRequestMemberList struct {
	// The permission information.
	//
	// This parameter is required.
	CdsIdentity *RemoveFilePermissionRequestMemberListCdsIdentity `json:"CdsIdentity,omitempty" xml:"CdsIdentity,omitempty" type:"Struct"`
	// You can set permissions by specifying roles or by customizing operation permissions. This field is used to set permissions by specifying roles. This field is mutually exclusive with `ActionList`.
	//
	// Valid values:
	//
	// 	- SystemFileEditorWithoutShareLink: the role that has the permissions to edit files but cannot share files.
	//
	// 	- SystemFileUploaderAndDownloaderWithShareLink: the role that has the permissions to upload, download, and share files.
	//
	// 	- SystemFileDownloader: the role that has the permissions to download files.
	//
	// 	- SystemFileEditorWithoutDelete: the role that has the permissions to edit files but cannot delete files.
	//
	// 	- SystemFileOwner: the role that has the permissions to collaborate with others.
	//
	// 	- SystemFileDownloaderWithShareLink: the role that has the permissions to download and share files
	//
	// 	- SystemFileUploaderAndViewer: the role that has the permissions to preview or upload files.
	//
	// 	- SystemFileViewer: the role that has the permissions to preview files.
	//
	// 	- SystemFileEditor: the role that has the permissions to edit files
	//
	// 	- SystemFileUploaderWithShareLink: the role that has the permissions to upload or share files.
	//
	// 	- SystemFileUploader: the role that has the permission to upload files.
	//
	// 	- SystemFileUploaderAndDownloader: the role that has the permissions to upload or download files.
	//
	// 	- SystemFileMetaViewer: the role that has the permissions to view files
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemFileUploaderAndDownloader
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s RemoveFilePermissionRequestMemberList) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionRequestMemberList) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionRequestMemberList) GetCdsIdentity() *RemoveFilePermissionRequestMemberListCdsIdentity {
	return s.CdsIdentity
}

func (s *RemoveFilePermissionRequestMemberList) GetRoleId() *string {
	return s.RoleId
}

func (s *RemoveFilePermissionRequestMemberList) SetCdsIdentity(v *RemoveFilePermissionRequestMemberListCdsIdentity) *RemoveFilePermissionRequestMemberList {
	s.CdsIdentity = v
	return s
}

func (s *RemoveFilePermissionRequestMemberList) SetRoleId(v string) *RemoveFilePermissionRequestMemberList {
	s.RoleId = &v
	return s
}

func (s *RemoveFilePermissionRequestMemberList) Validate() error {
	if s.CdsIdentity != nil {
		if err := s.CdsIdentity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveFilePermissionRequestMemberListCdsIdentity struct {
	// The user ID or group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 249dsfseee643h33g3dv****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The object type.
	//
	// Valid values:
	//
	// 	- IT_Group: group.
	//
	// 	- IT_User: user.
	//
	// This parameter is required.
	//
	// example:
	//
	// IT_User
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RemoveFilePermissionRequestMemberListCdsIdentity) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionRequestMemberListCdsIdentity) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionRequestMemberListCdsIdentity) GetId() *string {
	return s.Id
}

func (s *RemoveFilePermissionRequestMemberListCdsIdentity) GetType() *string {
	return s.Type
}

func (s *RemoveFilePermissionRequestMemberListCdsIdentity) SetId(v string) *RemoveFilePermissionRequestMemberListCdsIdentity {
	s.Id = &v
	return s
}

func (s *RemoveFilePermissionRequestMemberListCdsIdentity) SetType(v string) *RemoveFilePermissionRequestMemberListCdsIdentity {
	s.Type = &v
	return s
}

func (s *RemoveFilePermissionRequestMemberListCdsIdentity) Validate() error {
	return dara.Validate(s)
}
