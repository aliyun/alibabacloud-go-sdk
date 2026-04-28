// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareLinkByAnonymousResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetAccessCount() *int64
	SetAvatar(v string) *GetShareLinkByAnonymousResponseBody
	GetAvatar() *string
	SetCreatorId(v string) *GetShareLinkByAnonymousResponseBody
	GetCreatorId() *string
	SetCreatorName(v string) *GetShareLinkByAnonymousResponseBody
	GetCreatorName() *string
	SetCreatorPhone(v string) *GetShareLinkByAnonymousResponseBody
	GetCreatorPhone() *string
	SetDisableDownload(v bool) *GetShareLinkByAnonymousResponseBody
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *GetShareLinkByAnonymousResponseBody
	GetDisablePreview() *bool
	SetDisableSave(v bool) *GetShareLinkByAnonymousResponseBody
	GetDisableSave() *bool
	SetDownloadCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetDownloadCount() *int64
	SetDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody
	GetDownloadLimit() *int64
	SetExpiration(v string) *GetShareLinkByAnonymousResponseBody
	GetExpiration() *string
	SetHasPwd(v bool) *GetShareLinkByAnonymousResponseBody
	GetHasPwd() *bool
	SetPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetPreviewCount() *int64
	SetPreviewLimit(v int64) *GetShareLinkByAnonymousResponseBody
	GetPreviewLimit() *int64
	SetReportCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetReportCount() *int64
	SetSaveCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetSaveCount() *int64
	SetSaveDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody
	GetSaveDownloadLimit() *int64
	SetSaveLimit(v int64) *GetShareLinkByAnonymousResponseBody
	GetSaveLimit() *int64
	SetShareName(v string) *GetShareLinkByAnonymousResponseBody
	GetShareName() *string
	SetUpdatedAt(v string) *GetShareLinkByAnonymousResponseBody
	GetUpdatedAt() *string
	SetVideoPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody
	GetVideoPreviewCount() *int64
}

type GetShareLinkByAnonymousResponseBody struct {
	// The number of times that the shared files are visited.
	//
	// example:
	//
	// 30
	AccessCount *int64 `json:"access_count,omitempty" xml:"access_count,omitempty"`
	// The profile picture of the user who created the share link.
	//
	// example:
	//
	// https://aliyunpds.com/a.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The ID of the user who created the share link.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// The name of the user who created the share link. The value is masked.
	//
	// example:
	//
	// AB***CD
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// The mobile number of the user who created the share link. The value is masked.
	//
	// example:
	//
	// 136****00
	CreatorPhone *string `json:"creator_phone,omitempty" xml:"creator_phone,omitempty"`
	// Indicates whether the downloads of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Indicates whether the previews of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Indicates whether the saves of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The number of times that the shared files are downloaded.
	//
	// example:
	//
	// 50
	DownloadCount *int64 `json:"download_count,omitempty" xml:"download_count,omitempty"`
	// The maximum number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The time when the share link expires.
	//
	// example:
	//
	// 2020-08-20T06:51:27.292Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// Indicates whether a password is specified for the share link.
	//
	// example:
	//
	// true
	HasPwd *bool `json:"has_pwd,omitempty" xml:"has_pwd,omitempty"`
	// The number of times that the shared files are previewed.
	//
	// example:
	//
	// 80
	PreviewCount *int64 `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	// The maximum number of times that the shared files can be previewed.
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
	SaveCount *int64 `json:"save_count,omitempty" xml:"save_count,omitempty"`
	// The maximum number of times that the shared files can be saved and downloaded.
	//
	// example:
	//
	// 200
	SaveDownloadLimit *int64 `json:"save_download_limit,omitempty" xml:"save_download_limit,omitempty"`
	// The maximum number of times that the shared files can be saved.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// The name of the share link.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The time when the share link was last modified.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The number of times that the videos are previewed in the shared files.
	//
	// example:
	//
	// 5
	VideoPreviewCount *int64 `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s GetShareLinkByAnonymousResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetShareLinkByAnonymousResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousResponseBody) GetAccessCount() *int64 {
	return s.AccessCount
}

func (s *GetShareLinkByAnonymousResponseBody) GetAvatar() *string {
	return s.Avatar
}

func (s *GetShareLinkByAnonymousResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetShareLinkByAnonymousResponseBody) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetShareLinkByAnonymousResponseBody) GetCreatorPhone() *string {
	return s.CreatorPhone
}

func (s *GetShareLinkByAnonymousResponseBody) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *GetShareLinkByAnonymousResponseBody) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *GetShareLinkByAnonymousResponseBody) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *GetShareLinkByAnonymousResponseBody) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetShareLinkByAnonymousResponseBody) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *GetShareLinkByAnonymousResponseBody) GetExpiration() *string {
	return s.Expiration
}

func (s *GetShareLinkByAnonymousResponseBody) GetHasPwd() *bool {
	return s.HasPwd
}

func (s *GetShareLinkByAnonymousResponseBody) GetPreviewCount() *int64 {
	return s.PreviewCount
}

func (s *GetShareLinkByAnonymousResponseBody) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *GetShareLinkByAnonymousResponseBody) GetReportCount() *int64 {
	return s.ReportCount
}

func (s *GetShareLinkByAnonymousResponseBody) GetSaveCount() *int64 {
	return s.SaveCount
}

func (s *GetShareLinkByAnonymousResponseBody) GetSaveDownloadLimit() *int64 {
	return s.SaveDownloadLimit
}

func (s *GetShareLinkByAnonymousResponseBody) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *GetShareLinkByAnonymousResponseBody) GetShareName() *string {
	return s.ShareName
}

func (s *GetShareLinkByAnonymousResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetShareLinkByAnonymousResponseBody) GetVideoPreviewCount() *int64 {
	return s.VideoPreviewCount
}

func (s *GetShareLinkByAnonymousResponseBody) SetAccessCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.AccessCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetAvatar(v string) *GetShareLinkByAnonymousResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorId(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorName(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorPhone(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorPhone = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisableDownload(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisableDownload = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisablePreview(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisablePreview = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisableSave(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisableSave = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDownloadCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.DownloadCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.DownloadLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetExpiration(v string) *GetShareLinkByAnonymousResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetHasPwd(v bool) *GetShareLinkByAnonymousResponseBody {
	s.HasPwd = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.PreviewCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetPreviewLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.PreviewLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetReportCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.ReportCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveDownloadLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetShareName(v string) *GetShareLinkByAnonymousResponseBody {
	s.ShareName = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetUpdatedAt(v string) *GetShareLinkByAnonymousResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetVideoPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.VideoPreviewCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) Validate() error {
	return dara.Validate(s)
}
