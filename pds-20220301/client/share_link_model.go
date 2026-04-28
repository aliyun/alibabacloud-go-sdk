// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareLink interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCount(v int64) *ShareLink
	GetAccessCount() *int64
	SetCreatedAt(v string) *ShareLink
	GetCreatedAt() *string
	SetCreator(v string) *ShareLink
	GetCreator() *string
	SetDescription(v string) *ShareLink
	GetDescription() *string
	SetDisableDownload(v bool) *ShareLink
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *ShareLink
	GetDisablePreview() *bool
	SetDisableSave(v bool) *ShareLink
	GetDisableSave() *bool
	SetDownloadCount(v int64) *ShareLink
	GetDownloadCount() *int64
	SetDownloadLimit(v int64) *ShareLink
	GetDownloadLimit() *int64
	SetDriveId(v string) *ShareLink
	GetDriveId() *string
	SetExpiration(v string) *ShareLink
	GetExpiration() *string
	SetExpired(v bool) *ShareLink
	GetExpired() *bool
	SetFileIdList(v []*string) *ShareLink
	GetFileIdList() []*string
	SetOfficeEditable(v bool) *ShareLink
	GetOfficeEditable() *bool
	SetPreviewCount(v int64) *ShareLink
	GetPreviewCount() *int64
	SetPreviewLimit(v int64) *ShareLink
	GetPreviewLimit() *int64
	SetReportCount(v int64) *ShareLink
	GetReportCount() *int64
	SetSaveCount(v int64) *ShareLink
	GetSaveCount() *int64
	SetSaveDownloadLimit(v int64) *ShareLink
	GetSaveDownloadLimit() *int64
	SetSaveLimit(v int64) *ShareLink
	GetSaveLimit() *int64
	SetShareAllFiles(v bool) *ShareLink
	GetShareAllFiles() *bool
	SetShareId(v string) *ShareLink
	GetShareId() *string
	SetShareName(v string) *ShareLink
	GetShareName() *string
	SetSharePwd(v string) *ShareLink
	GetSharePwd() *string
	SetStatus(v string) *ShareLink
	GetStatus() *string
	SetUpdatedAt(v string) *ShareLink
	GetUpdatedAt() *string
	SetVideoPreviewCount(v int64) *ShareLink
	GetVideoPreviewCount() *int64
}

type ShareLink struct {
	// The number of times that the shared files are visited.
	//
	// example:
	//
	// 4
	AccessCount *int64 `json:"access_count,omitempty" xml:"access_count,omitempty"`
	// The time when the share was created.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The user who created the share.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The description of the share.
	//
	// example:
	//
	// videos
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
	// Specifies whether to disable the save feature.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The number of times that the shared files are downloaded.
	//
	// example:
	//
	// 5
	DownloadCount *int64 `json:"download_count,omitempty" xml:"download_count,omitempty"`
	// The limit on the number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The time when the share URL expires. The value of this parameter follows the RFC 3339 standard. Example: "2020-06-28T11:33:00.000+08:00". If you set the value to "", the share URL never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// Specifies whether the share is expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"expired,omitempty" xml:"expired,omitempty"`
	// The IDs of the files to share in the parent path.
	//
	// example:
	//
	// ["520b217f13adf4fc24f2191991b1664ce045b393"]
	FileIdList     []*string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty" type:"Repeated"`
	OfficeEditable *bool     `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	// The number of times that the shared files are previewed.
	//
	// example:
	//
	// 10
	PreviewCount *int64 `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	// The limit on the number of times that the shared files can be previewed.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	// The number of times that the shared files are reported.
	//
	// example:
	//
	// 0
	ReportCount *int64 `json:"report_count,omitempty" xml:"report_count,omitempty"`
	// The number of times that the shared files are saved.
	//
	// example:
	//
	// 2
	SaveCount         *int64 `json:"save_count,omitempty" xml:"save_count,omitempty"`
	SaveDownloadLimit *int64 `json:"save_download_limit,omitempty" xml:"save_download_limit,omitempty"`
	// The limit on the number of times that the shared files can be saved.
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
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The name of the share. By default, the file name that corresponds to the first ID in the file ID list is used.
	//
	// example:
	//
	// video-1.MP4
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The access code. An access code can be up to 64 characters in length. If you do not specify a value, files can be accessed without an access code.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	// The status of the share. Valid values:
	//
	// 	- disabled: The share is canceled.
	//
	// 	- enabled: The share is effective.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the share was last modified.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The number of times that the shared audio and video files are played.
	//
	// example:
	//
	// 1
	VideoPreviewCount *int64 `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s ShareLink) String() string {
	return dara.Prettify(s)
}

func (s ShareLink) GoString() string {
	return s.String()
}

func (s *ShareLink) GetAccessCount() *int64 {
	return s.AccessCount
}

func (s *ShareLink) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ShareLink) GetCreator() *string {
	return s.Creator
}

func (s *ShareLink) GetDescription() *string {
	return s.Description
}

func (s *ShareLink) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *ShareLink) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *ShareLink) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *ShareLink) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *ShareLink) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *ShareLink) GetDriveId() *string {
	return s.DriveId
}

func (s *ShareLink) GetExpiration() *string {
	return s.Expiration
}

func (s *ShareLink) GetExpired() *bool {
	return s.Expired
}

func (s *ShareLink) GetFileIdList() []*string {
	return s.FileIdList
}

func (s *ShareLink) GetOfficeEditable() *bool {
	return s.OfficeEditable
}

func (s *ShareLink) GetPreviewCount() *int64 {
	return s.PreviewCount
}

func (s *ShareLink) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *ShareLink) GetReportCount() *int64 {
	return s.ReportCount
}

func (s *ShareLink) GetSaveCount() *int64 {
	return s.SaveCount
}

func (s *ShareLink) GetSaveDownloadLimit() *int64 {
	return s.SaveDownloadLimit
}

func (s *ShareLink) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *ShareLink) GetShareAllFiles() *bool {
	return s.ShareAllFiles
}

func (s *ShareLink) GetShareId() *string {
	return s.ShareId
}

func (s *ShareLink) GetShareName() *string {
	return s.ShareName
}

func (s *ShareLink) GetSharePwd() *string {
	return s.SharePwd
}

func (s *ShareLink) GetStatus() *string {
	return s.Status
}

func (s *ShareLink) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ShareLink) GetVideoPreviewCount() *int64 {
	return s.VideoPreviewCount
}

func (s *ShareLink) SetAccessCount(v int64) *ShareLink {
	s.AccessCount = &v
	return s
}

func (s *ShareLink) SetCreatedAt(v string) *ShareLink {
	s.CreatedAt = &v
	return s
}

func (s *ShareLink) SetCreator(v string) *ShareLink {
	s.Creator = &v
	return s
}

func (s *ShareLink) SetDescription(v string) *ShareLink {
	s.Description = &v
	return s
}

func (s *ShareLink) SetDisableDownload(v bool) *ShareLink {
	s.DisableDownload = &v
	return s
}

func (s *ShareLink) SetDisablePreview(v bool) *ShareLink {
	s.DisablePreview = &v
	return s
}

func (s *ShareLink) SetDisableSave(v bool) *ShareLink {
	s.DisableSave = &v
	return s
}

func (s *ShareLink) SetDownloadCount(v int64) *ShareLink {
	s.DownloadCount = &v
	return s
}

func (s *ShareLink) SetDownloadLimit(v int64) *ShareLink {
	s.DownloadLimit = &v
	return s
}

func (s *ShareLink) SetDriveId(v string) *ShareLink {
	s.DriveId = &v
	return s
}

func (s *ShareLink) SetExpiration(v string) *ShareLink {
	s.Expiration = &v
	return s
}

func (s *ShareLink) SetExpired(v bool) *ShareLink {
	s.Expired = &v
	return s
}

func (s *ShareLink) SetFileIdList(v []*string) *ShareLink {
	s.FileIdList = v
	return s
}

func (s *ShareLink) SetOfficeEditable(v bool) *ShareLink {
	s.OfficeEditable = &v
	return s
}

func (s *ShareLink) SetPreviewCount(v int64) *ShareLink {
	s.PreviewCount = &v
	return s
}

func (s *ShareLink) SetPreviewLimit(v int64) *ShareLink {
	s.PreviewLimit = &v
	return s
}

func (s *ShareLink) SetReportCount(v int64) *ShareLink {
	s.ReportCount = &v
	return s
}

func (s *ShareLink) SetSaveCount(v int64) *ShareLink {
	s.SaveCount = &v
	return s
}

func (s *ShareLink) SetSaveDownloadLimit(v int64) *ShareLink {
	s.SaveDownloadLimit = &v
	return s
}

func (s *ShareLink) SetSaveLimit(v int64) *ShareLink {
	s.SaveLimit = &v
	return s
}

func (s *ShareLink) SetShareAllFiles(v bool) *ShareLink {
	s.ShareAllFiles = &v
	return s
}

func (s *ShareLink) SetShareId(v string) *ShareLink {
	s.ShareId = &v
	return s
}

func (s *ShareLink) SetShareName(v string) *ShareLink {
	s.ShareName = &v
	return s
}

func (s *ShareLink) SetSharePwd(v string) *ShareLink {
	s.SharePwd = &v
	return s
}

func (s *ShareLink) SetStatus(v string) *ShareLink {
	s.Status = &v
	return s
}

func (s *ShareLink) SetUpdatedAt(v string) *ShareLink {
	s.UpdatedAt = &v
	return s
}

func (s *ShareLink) SetVideoPreviewCount(v int64) *ShareLink {
	s.VideoPreviewCount = &v
	return s
}

func (s *ShareLink) Validate() error {
	return dara.Validate(s)
}
