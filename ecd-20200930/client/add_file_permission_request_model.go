// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *AddFilePermissionRequest
	GetCdsId() *string
	SetEndUserId(v string) *AddFilePermissionRequest
	GetEndUserId() *string
	SetFileId(v string) *AddFilePermissionRequest
	GetFileId() *string
	SetGroupId(v string) *AddFilePermissionRequest
	GetGroupId() *string
	SetMemberList(v []*AddFilePermissionRequestMemberList) *AddFilePermissionRequest
	GetMemberList() []*AddFilePermissionRequestMemberList
	SetRegionId(v string) *AddFilePermissionRequest
	GetRegionId() *string
}

type AddFilePermissionRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user who uses the network disk.
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
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The users that you want to authorize to use the cloud disk.
	//
	// This parameter is required.
	MemberList []*AddFilePermissionRequestMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddFilePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionRequest) GoString() string {
	return s.String()
}

func (s *AddFilePermissionRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *AddFilePermissionRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *AddFilePermissionRequest) GetFileId() *string {
	return s.FileId
}

func (s *AddFilePermissionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddFilePermissionRequest) GetMemberList() []*AddFilePermissionRequestMemberList {
	return s.MemberList
}

func (s *AddFilePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddFilePermissionRequest) SetCdsId(v string) *AddFilePermissionRequest {
	s.CdsId = &v
	return s
}

func (s *AddFilePermissionRequest) SetEndUserId(v string) *AddFilePermissionRequest {
	s.EndUserId = &v
	return s
}

func (s *AddFilePermissionRequest) SetFileId(v string) *AddFilePermissionRequest {
	s.FileId = &v
	return s
}

func (s *AddFilePermissionRequest) SetGroupId(v string) *AddFilePermissionRequest {
	s.GroupId = &v
	return s
}

func (s *AddFilePermissionRequest) SetMemberList(v []*AddFilePermissionRequestMemberList) *AddFilePermissionRequest {
	s.MemberList = v
	return s
}

func (s *AddFilePermissionRequest) SetRegionId(v string) *AddFilePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *AddFilePermissionRequest) Validate() error {
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

type AddFilePermissionRequestMemberList struct {
	// The user of the cloud disk.
	//
	// This parameter is required.
	CdsIdentity *AddFilePermissionRequestMemberListCdsIdentity `json:"CdsIdentity,omitempty" xml:"CdsIdentity,omitempty" type:"Struct"`
	// Specifies whether the users of the child group can inherit the folder permissions.
	//
	// example:
	//
	// false
	DisinheritSubGroup *bool `json:"DisinheritSubGroup,omitempty" xml:"DisinheritSubGroup,omitempty"`
	// The time when the authorization expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. The value never expires. You can specify a value that is predefined by the system for this parameter. Example: 4775500800000.
	//
	// example:
	//
	// 4775500800000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// You can set permissions by specifying roles or by customizing operation permissions. This field is used to set permissions by specifying roles. This field is mutually exclusive with `ActionList`.
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
	// SystemFileUploaderAndDownloader
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s AddFilePermissionRequestMemberList) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionRequestMemberList) GoString() string {
	return s.String()
}

func (s *AddFilePermissionRequestMemberList) GetCdsIdentity() *AddFilePermissionRequestMemberListCdsIdentity {
	return s.CdsIdentity
}

func (s *AddFilePermissionRequestMemberList) GetDisinheritSubGroup() *bool {
	return s.DisinheritSubGroup
}

func (s *AddFilePermissionRequestMemberList) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *AddFilePermissionRequestMemberList) GetRoleId() *string {
	return s.RoleId
}

func (s *AddFilePermissionRequestMemberList) SetCdsIdentity(v *AddFilePermissionRequestMemberListCdsIdentity) *AddFilePermissionRequestMemberList {
	s.CdsIdentity = v
	return s
}

func (s *AddFilePermissionRequestMemberList) SetDisinheritSubGroup(v bool) *AddFilePermissionRequestMemberList {
	s.DisinheritSubGroup = &v
	return s
}

func (s *AddFilePermissionRequestMemberList) SetExpireTime(v int64) *AddFilePermissionRequestMemberList {
	s.ExpireTime = &v
	return s
}

func (s *AddFilePermissionRequestMemberList) SetRoleId(v string) *AddFilePermissionRequestMemberList {
	s.RoleId = &v
	return s
}

func (s *AddFilePermissionRequestMemberList) Validate() error {
	if s.CdsIdentity != nil {
		if err := s.CdsIdentity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFilePermissionRequestMemberListCdsIdentity struct {
	// The ID of the convenience user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user01
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user type.
	//
	// Set the value to TENANT_ADMIN.
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

func (s AddFilePermissionRequestMemberListCdsIdentity) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionRequestMemberListCdsIdentity) GoString() string {
	return s.String()
}

func (s *AddFilePermissionRequestMemberListCdsIdentity) GetId() *string {
	return s.Id
}

func (s *AddFilePermissionRequestMemberListCdsIdentity) GetType() *string {
	return s.Type
}

func (s *AddFilePermissionRequestMemberListCdsIdentity) SetId(v string) *AddFilePermissionRequestMemberListCdsIdentity {
	s.Id = &v
	return s
}

func (s *AddFilePermissionRequestMemberListCdsIdentity) SetType(v string) *AddFilePermissionRequestMemberListCdsIdentity {
	s.Type = &v
	return s
}

func (s *AddFilePermissionRequestMemberListCdsIdentity) Validate() error {
	return dara.Validate(s)
}
