// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatable(v bool) *CreateShareLinkRequest
	GetCreatable() *bool
	SetCreatableFileIdList(v []*string) *CreateShareLinkRequest
	GetCreatableFileIdList() []*string
	SetDescription(v string) *CreateShareLinkRequest
	GetDescription() *string
	SetDisableDownload(v bool) *CreateShareLinkRequest
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *CreateShareLinkRequest
	GetDisablePreview() *bool
	SetDisableSave(v bool) *CreateShareLinkRequest
	GetDisableSave() *bool
	SetDownloadLimit(v int64) *CreateShareLinkRequest
	GetDownloadLimit() *int64
	SetDriveId(v string) *CreateShareLinkRequest
	GetDriveId() *string
	SetExpiration(v string) *CreateShareLinkRequest
	GetExpiration() *string
	SetFileIdList(v []*string) *CreateShareLinkRequest
	GetFileIdList() []*string
	SetOfficeEditable(v bool) *CreateShareLinkRequest
	GetOfficeEditable() *bool
	SetPreviewLimit(v int64) *CreateShareLinkRequest
	GetPreviewLimit() *int64
	SetRequireLogin(v bool) *CreateShareLinkRequest
	GetRequireLogin() *bool
	SetSaveLimit(v int64) *CreateShareLinkRequest
	GetSaveLimit() *int64
	SetShareAllFiles(v bool) *CreateShareLinkRequest
	GetShareAllFiles() *bool
	SetShareName(v string) *CreateShareLinkRequest
	GetShareName() *string
	SetSharePwd(v string) *CreateShareLinkRequest
	GetSharePwd() *string
	SetUserId(v string) *CreateShareLinkRequest
	GetUserId() *string
}

type CreateShareLinkRequest struct {
	Creatable           *bool     `json:"creatable,omitempty" xml:"creatable,omitempty"`
	CreatableFileIdList []*string `json:"creatable_file_id_list,omitempty" xml:"creatable_file_id_list,omitempty" type:"Repeated"`
	// The description of the share. The description must be 0 to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable the download feature.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Specifies whether to disable the preview feature.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Specifies whether to disable the dump feature.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The limit on the number of times that the shared files can be downloaded. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The time when the share URL expires. The value of this parameter follows the RFC 3339 standard. Example: "2020-06-28T11:33:00.000+08:00". If expiration is set to "", the share URL never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The IDs of the files to share in the parent path. The number of files in the parent path ranges from 1 to 100. If share_all_files is set to true, this parameter does not take effect. Otherwise, you must specify this parameter.``
	//
	// example:
	//
	// ["520b217f13adf4fc24f2191991b1664ce045b393"]
	FileIdList     []*string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty" type:"Repeated"`
	OfficeEditable *bool     `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	// The limit on the number of times that the shared files can be previewed. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	RequireLogin *bool  `json:"require_login,omitempty" xml:"require_login,omitempty"`
	// The limit on the number of times that the shared files can be dumped. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// Specifies whether to share all files in the drive.
	//
	// example:
	//
	// true
	ShareAllFiles *bool `json:"share_all_files,omitempty" xml:"share_all_files,omitempty"`
	// The name of the share. If you leave this parameter empty, the file name that corresponds to the first ID in the file ID list is used. The name must be 0 to 128 characters in length.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The access code. An access code must be 0 to 64 bytes in length. If you do not specify this parameter or leave this parameter empty, the files can be accessed without an access code. In this case, you do not need to specify the share_pwd parameter when you call an operation to query the share URL. The access code can contain only visible ASCII characters.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CreateShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateShareLinkRequest) GetCreatable() *bool {
	return s.Creatable
}

func (s *CreateShareLinkRequest) GetCreatableFileIdList() []*string {
	return s.CreatableFileIdList
}

func (s *CreateShareLinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateShareLinkRequest) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *CreateShareLinkRequest) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *CreateShareLinkRequest) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *CreateShareLinkRequest) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *CreateShareLinkRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CreateShareLinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *CreateShareLinkRequest) GetFileIdList() []*string {
	return s.FileIdList
}

func (s *CreateShareLinkRequest) GetOfficeEditable() *bool {
	return s.OfficeEditable
}

func (s *CreateShareLinkRequest) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *CreateShareLinkRequest) GetRequireLogin() *bool {
	return s.RequireLogin
}

func (s *CreateShareLinkRequest) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *CreateShareLinkRequest) GetShareAllFiles() *bool {
	return s.ShareAllFiles
}

func (s *CreateShareLinkRequest) GetShareName() *string {
	return s.ShareName
}

func (s *CreateShareLinkRequest) GetSharePwd() *string {
	return s.SharePwd
}

func (s *CreateShareLinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateShareLinkRequest) SetCreatable(v bool) *CreateShareLinkRequest {
	s.Creatable = &v
	return s
}

func (s *CreateShareLinkRequest) SetCreatableFileIdList(v []*string) *CreateShareLinkRequest {
	s.CreatableFileIdList = v
	return s
}

func (s *CreateShareLinkRequest) SetDescription(v string) *CreateShareLinkRequest {
	s.Description = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisableDownload(v bool) *CreateShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisablePreview(v bool) *CreateShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisableSave(v bool) *CreateShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *CreateShareLinkRequest) SetDownloadLimit(v int64) *CreateShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetDriveId(v string) *CreateShareLinkRequest {
	s.DriveId = &v
	return s
}

func (s *CreateShareLinkRequest) SetExpiration(v string) *CreateShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *CreateShareLinkRequest) SetFileIdList(v []*string) *CreateShareLinkRequest {
	s.FileIdList = v
	return s
}

func (s *CreateShareLinkRequest) SetOfficeEditable(v bool) *CreateShareLinkRequest {
	s.OfficeEditable = &v
	return s
}

func (s *CreateShareLinkRequest) SetPreviewLimit(v int64) *CreateShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetRequireLogin(v bool) *CreateShareLinkRequest {
	s.RequireLogin = &v
	return s
}

func (s *CreateShareLinkRequest) SetSaveLimit(v int64) *CreateShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetShareAllFiles(v bool) *CreateShareLinkRequest {
	s.ShareAllFiles = &v
	return s
}

func (s *CreateShareLinkRequest) SetShareName(v string) *CreateShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *CreateShareLinkRequest) SetSharePwd(v string) *CreateShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *CreateShareLinkRequest) SetUserId(v string) *CreateShareLinkRequest {
	s.UserId = &v
	return s
}

func (s *CreateShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
