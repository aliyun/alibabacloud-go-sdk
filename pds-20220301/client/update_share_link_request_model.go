// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateShareLinkRequest
	GetDescription() *string
	SetDisableDownload(v bool) *UpdateShareLinkRequest
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *UpdateShareLinkRequest
	GetDisablePreview() *bool
	SetDisableSave(v bool) *UpdateShareLinkRequest
	GetDisableSave() *bool
	SetDownloadCount(v int64) *UpdateShareLinkRequest
	GetDownloadCount() *int64
	SetDownloadLimit(v int64) *UpdateShareLinkRequest
	GetDownloadLimit() *int64
	SetExpiration(v string) *UpdateShareLinkRequest
	GetExpiration() *string
	SetOfficeEditable(v bool) *UpdateShareLinkRequest
	GetOfficeEditable() *bool
	SetPreviewCount(v int64) *UpdateShareLinkRequest
	GetPreviewCount() *int64
	SetPreviewLimit(v int64) *UpdateShareLinkRequest
	GetPreviewLimit() *int64
	SetReportCount(v int64) *UpdateShareLinkRequest
	GetReportCount() *int64
	SetSaveCount(v int64) *UpdateShareLinkRequest
	GetSaveCount() *int64
	SetSaveLimit(v int64) *UpdateShareLinkRequest
	GetSaveLimit() *int64
	SetShareId(v string) *UpdateShareLinkRequest
	GetShareId() *string
	SetShareName(v string) *UpdateShareLinkRequest
	GetShareName() *string
	SetSharePwd(v string) *UpdateShareLinkRequest
	GetSharePwd() *string
	SetStatus(v string) *UpdateShareLinkRequest
	GetStatus() *string
	SetVideoPreviewCount(v int64) *UpdateShareLinkRequest
	GetVideoPreviewCount() *int64
}

type UpdateShareLinkRequest struct {
	// The description of the share link. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to prohibit the downloads of the shared files.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Specifies whether to prohibit the previews of the shared files.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Specifies whether to prohibit the saves of the shared files.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The number of times that the shared files are downloaded. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 30
	DownloadCount *int64 `json:"download_count,omitempty" xml:"download_count,omitempty"`
	// The maximum number of times that the shared files can be downloaded. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The time when the share link expires. The time follows the RFC 3339 standard. Example: 2020-06-28T11:33:00.000+08:00. If you leave this parameter empty, the share link never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration     *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	OfficeEditable *bool   `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	// The number of times that the shared files are previewed. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 3
	PreviewCount *int64 `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	// The maximum number of times that the shared files can be previewed. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	// The number of times that the shared files are reported. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 1
	ReportCount *int64 `json:"report_count,omitempty" xml:"report_count,omitempty"`
	// The number of times that the shared files are saved. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 5
	SaveCount *int64 `json:"save_count,omitempty" xml:"save_count,omitempty"`
	// The maximum number of times that the shared files can be saved. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The name of the share link. By default, the name of the first file is used. The name can be up to 128 characters in length.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The access code. The access code can be up to 64 characters in length. A value of 0 specifies an empty string.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	// The state of the share link. Valid values:
	//
	// 	- disabled: The share link is canceled.
	//
	// 	- enabled: The share link is effective.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of times that the videos are previewed in the shared files. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 100
	VideoPreviewCount *int64 `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s UpdateShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateShareLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateShareLinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateShareLinkRequest) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *UpdateShareLinkRequest) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *UpdateShareLinkRequest) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *UpdateShareLinkRequest) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *UpdateShareLinkRequest) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *UpdateShareLinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *UpdateShareLinkRequest) GetOfficeEditable() *bool {
	return s.OfficeEditable
}

func (s *UpdateShareLinkRequest) GetPreviewCount() *int64 {
	return s.PreviewCount
}

func (s *UpdateShareLinkRequest) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *UpdateShareLinkRequest) GetReportCount() *int64 {
	return s.ReportCount
}

func (s *UpdateShareLinkRequest) GetSaveCount() *int64 {
	return s.SaveCount
}

func (s *UpdateShareLinkRequest) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *UpdateShareLinkRequest) GetShareId() *string {
	return s.ShareId
}

func (s *UpdateShareLinkRequest) GetShareName() *string {
	return s.ShareName
}

func (s *UpdateShareLinkRequest) GetSharePwd() *string {
	return s.SharePwd
}

func (s *UpdateShareLinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateShareLinkRequest) GetVideoPreviewCount() *int64 {
	return s.VideoPreviewCount
}

func (s *UpdateShareLinkRequest) SetDescription(v string) *UpdateShareLinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisableDownload(v bool) *UpdateShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisablePreview(v bool) *UpdateShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisableSave(v bool) *UpdateShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDownloadCount(v int64) *UpdateShareLinkRequest {
	s.DownloadCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDownloadLimit(v int64) *UpdateShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetExpiration(v string) *UpdateShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateShareLinkRequest) SetOfficeEditable(v bool) *UpdateShareLinkRequest {
	s.OfficeEditable = &v
	return s
}

func (s *UpdateShareLinkRequest) SetPreviewCount(v int64) *UpdateShareLinkRequest {
	s.PreviewCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetPreviewLimit(v int64) *UpdateShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetReportCount(v int64) *UpdateShareLinkRequest {
	s.ReportCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSaveCount(v int64) *UpdateShareLinkRequest {
	s.SaveCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSaveLimit(v int64) *UpdateShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetShareId(v string) *UpdateShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *UpdateShareLinkRequest) SetShareName(v string) *UpdateShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSharePwd(v string) *UpdateShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *UpdateShareLinkRequest) SetStatus(v string) *UpdateShareLinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateShareLinkRequest) SetVideoPreviewCount(v int64) *UpdateShareLinkRequest {
	s.VideoPreviewCount = &v
	return s
}

func (s *UpdateShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
