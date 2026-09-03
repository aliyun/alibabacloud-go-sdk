// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdsFileShareLinkModel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCount(v int64) *CdsFileShareLinkModel
	GetAccessCount() *int64
	SetCreateTime(v string) *CdsFileShareLinkModel
	GetCreateTime() *string
	SetCreator(v string) *CdsFileShareLinkModel
	GetCreator() *string
	SetDescription(v string) *CdsFileShareLinkModel
	GetDescription() *string
	SetDisableDownload(v bool) *CdsFileShareLinkModel
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *CdsFileShareLinkModel
	GetDisablePreview() *bool
	SetDisableSave(v bool) *CdsFileShareLinkModel
	GetDisableSave() *bool
	SetDownloadCount(v int64) *CdsFileShareLinkModel
	GetDownloadCount() *int64
	SetDownloadLimit(v int64) *CdsFileShareLinkModel
	GetDownloadLimit() *int64
	SetDriveId(v string) *CdsFileShareLinkModel
	GetDriveId() *string
	SetExpiration(v string) *CdsFileShareLinkModel
	GetExpiration() *string
	SetExpired(v bool) *CdsFileShareLinkModel
	GetExpired() *bool
	SetFileIds(v string) *CdsFileShareLinkModel
	GetFileIds() *string
	SetModifiyTime(v string) *CdsFileShareLinkModel
	GetModifiyTime() *string
	SetPreviewCount(v int64) *CdsFileShareLinkModel
	GetPreviewCount() *int64
	SetPreviewLimit(v int64) *CdsFileShareLinkModel
	GetPreviewLimit() *int64
	SetReportCount(v int64) *CdsFileShareLinkModel
	GetReportCount() *int64
	SetSaveCount(v int64) *CdsFileShareLinkModel
	GetSaveCount() *int64
	SetSaveLimit(v int64) *CdsFileShareLinkModel
	GetSaveLimit() *int64
	SetShareId(v string) *CdsFileShareLinkModel
	GetShareId() *string
	SetShareLink(v string) *CdsFileShareLinkModel
	GetShareLink() *string
	SetShareName(v string) *CdsFileShareLinkModel
	GetShareName() *string
	SetSharePwd(v string) *CdsFileShareLinkModel
	GetSharePwd() *string
	SetStatus(v string) *CdsFileShareLinkModel
	GetStatus() *string
	SetVideoPreviewCount(v int64) *CdsFileShareLinkModel
	GetVideoPreviewCount() *int64
}

type CdsFileShareLinkModel struct {
	// The number of times to access the shared file.
	//
	// example:
	//
	// 10000
	AccessCount *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-09-04T03:30:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user that creates the file sharing task.
	//
	// example:
	//
	// user01
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description.
	//
	// example:
	//
	// newDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The shared file is forbidden from being downloaded.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"DisableDownload,omitempty" xml:"DisableDownload,omitempty"`
	// The shared file is forbidden from being previewed.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"DisablePreview,omitempty" xml:"DisablePreview,omitempty"`
	// The shared file is forbidden from being dumped.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"DisableSave,omitempty" xml:"DisableSave,omitempty"`
	// The number of times that the shared file can be downloaded.
	//
	// example:
	//
	// 100
	DownloadCount *int64 `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	// The maximum number of times that the shared file can be downloaded.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	// The team space ID.
	//
	// example:
	//
	// dri-g0877jp3hu1ox****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The time when the file sharing link expires.
	//
	// >  The value must be in the FC3339 format. Example: 2020-06-28T11:33:00.000+08:00. If the parameter is left empty, the file sharing link is permanently valid.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// Specifies whether the file sharing link expires.
	//
	// example:
	//
	// False
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The file sharing task IDs.
	//
	// example:
	//
	// [\\"63886f1fe2014d9a5a3348768dcc27dfc57ee103\\"]
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2021-09-04T04:30:36Z
	ModifiyTime *string `json:"ModifiyTime,omitempty" xml:"ModifiyTime,omitempty"`
	// The number of times that the shared file is previewed.
	//
	// example:
	//
	// 0
	PreviewCount *int64 `json:"PreviewCount,omitempty" xml:"PreviewCount,omitempty"`
	// The maximum number of times that the shared file can be previewed.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"PreviewLimit,omitempty" xml:"PreviewLimit,omitempty"`
	// The number of times that the shared file is reported due to content violation.
	//
	// example:
	//
	// 0
	ReportCount *int64 `json:"ReportCount,omitempty" xml:"ReportCount,omitempty"`
	// The number of times that the shared files can be dumped.
	//
	// example:
	//
	// 100
	SaveCount *int64 `json:"SaveCount,omitempty" xml:"SaveCount,omitempty"`
	// The maximum number of times that the shared file can be saved.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"SaveLimit,omitempty" xml:"SaveLimit,omitempty"`
	// The file sharing task ID.
	//
	// example:
	//
	// 7JQX1Fs****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	// The file sharing link.
	//
	// example:
	//
	// https://stg109960.apps.aliyunpds.com/disk/s/7uLJanz****
	ShareLink *string `json:"ShareLink,omitempty" xml:"ShareLink,omitempty"`
	// The shared file name. By default, the name of the first shared file is used.
	//
	// example:
	//
	// view.txt
	ShareName *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	// The access code. It must contain up to 64 characters in length. 0 characters indicates that there is no access code.
	//
	// example:
	//
	// 12345678
	SharePwd *string `json:"SharePwd,omitempty" xml:"SharePwd,omitempty"`
	// The status of the file sharing link.
	//
	// Valid values:
	//
	// 	- forbidden_disabled
	//
	//     <!-- -->
	//
	//     : The file sharing link is canceled
	//
	//     <!-- -->
	//
	//     after it
	//
	//     <!-- -->
	//
	//     is not allowed.
	//
	// 	- forbidden
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The file sharing link is not allowed
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- disabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The file sharing link is canceled.
	//
	//     <!-- -->
	//
	// 	- enabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The file sharing link is valid.
	//
	//     <!-- -->
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of times that the audio and video file is played.
	//
	// example:
	//
	// 0
	VideoPreviewCount *int64 `json:"VideoPreviewCount,omitempty" xml:"VideoPreviewCount,omitempty"`
}

func (s CdsFileShareLinkModel) String() string {
	return dara.Prettify(s)
}

func (s CdsFileShareLinkModel) GoString() string {
	return s.String()
}

func (s *CdsFileShareLinkModel) GetAccessCount() *int64 {
	return s.AccessCount
}

func (s *CdsFileShareLinkModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CdsFileShareLinkModel) GetCreator() *string {
	return s.Creator
}

func (s *CdsFileShareLinkModel) GetDescription() *string {
	return s.Description
}

func (s *CdsFileShareLinkModel) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *CdsFileShareLinkModel) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *CdsFileShareLinkModel) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *CdsFileShareLinkModel) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *CdsFileShareLinkModel) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *CdsFileShareLinkModel) GetDriveId() *string {
	return s.DriveId
}

func (s *CdsFileShareLinkModel) GetExpiration() *string {
	return s.Expiration
}

func (s *CdsFileShareLinkModel) GetExpired() *bool {
	return s.Expired
}

func (s *CdsFileShareLinkModel) GetFileIds() *string {
	return s.FileIds
}

func (s *CdsFileShareLinkModel) GetModifiyTime() *string {
	return s.ModifiyTime
}

func (s *CdsFileShareLinkModel) GetPreviewCount() *int64 {
	return s.PreviewCount
}

func (s *CdsFileShareLinkModel) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *CdsFileShareLinkModel) GetReportCount() *int64 {
	return s.ReportCount
}

func (s *CdsFileShareLinkModel) GetSaveCount() *int64 {
	return s.SaveCount
}

func (s *CdsFileShareLinkModel) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *CdsFileShareLinkModel) GetShareId() *string {
	return s.ShareId
}

func (s *CdsFileShareLinkModel) GetShareLink() *string {
	return s.ShareLink
}

func (s *CdsFileShareLinkModel) GetShareName() *string {
	return s.ShareName
}

func (s *CdsFileShareLinkModel) GetSharePwd() *string {
	return s.SharePwd
}

func (s *CdsFileShareLinkModel) GetStatus() *string {
	return s.Status
}

func (s *CdsFileShareLinkModel) GetVideoPreviewCount() *int64 {
	return s.VideoPreviewCount
}

func (s *CdsFileShareLinkModel) SetAccessCount(v int64) *CdsFileShareLinkModel {
	s.AccessCount = &v
	return s
}

func (s *CdsFileShareLinkModel) SetCreateTime(v string) *CdsFileShareLinkModel {
	s.CreateTime = &v
	return s
}

func (s *CdsFileShareLinkModel) SetCreator(v string) *CdsFileShareLinkModel {
	s.Creator = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDescription(v string) *CdsFileShareLinkModel {
	s.Description = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDisableDownload(v bool) *CdsFileShareLinkModel {
	s.DisableDownload = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDisablePreview(v bool) *CdsFileShareLinkModel {
	s.DisablePreview = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDisableSave(v bool) *CdsFileShareLinkModel {
	s.DisableSave = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDownloadCount(v int64) *CdsFileShareLinkModel {
	s.DownloadCount = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDownloadLimit(v int64) *CdsFileShareLinkModel {
	s.DownloadLimit = &v
	return s
}

func (s *CdsFileShareLinkModel) SetDriveId(v string) *CdsFileShareLinkModel {
	s.DriveId = &v
	return s
}

func (s *CdsFileShareLinkModel) SetExpiration(v string) *CdsFileShareLinkModel {
	s.Expiration = &v
	return s
}

func (s *CdsFileShareLinkModel) SetExpired(v bool) *CdsFileShareLinkModel {
	s.Expired = &v
	return s
}

func (s *CdsFileShareLinkModel) SetFileIds(v string) *CdsFileShareLinkModel {
	s.FileIds = &v
	return s
}

func (s *CdsFileShareLinkModel) SetModifiyTime(v string) *CdsFileShareLinkModel {
	s.ModifiyTime = &v
	return s
}

func (s *CdsFileShareLinkModel) SetPreviewCount(v int64) *CdsFileShareLinkModel {
	s.PreviewCount = &v
	return s
}

func (s *CdsFileShareLinkModel) SetPreviewLimit(v int64) *CdsFileShareLinkModel {
	s.PreviewLimit = &v
	return s
}

func (s *CdsFileShareLinkModel) SetReportCount(v int64) *CdsFileShareLinkModel {
	s.ReportCount = &v
	return s
}

func (s *CdsFileShareLinkModel) SetSaveCount(v int64) *CdsFileShareLinkModel {
	s.SaveCount = &v
	return s
}

func (s *CdsFileShareLinkModel) SetSaveLimit(v int64) *CdsFileShareLinkModel {
	s.SaveLimit = &v
	return s
}

func (s *CdsFileShareLinkModel) SetShareId(v string) *CdsFileShareLinkModel {
	s.ShareId = &v
	return s
}

func (s *CdsFileShareLinkModel) SetShareLink(v string) *CdsFileShareLinkModel {
	s.ShareLink = &v
	return s
}

func (s *CdsFileShareLinkModel) SetShareName(v string) *CdsFileShareLinkModel {
	s.ShareName = &v
	return s
}

func (s *CdsFileShareLinkModel) SetSharePwd(v string) *CdsFileShareLinkModel {
	s.SharePwd = &v
	return s
}

func (s *CdsFileShareLinkModel) SetStatus(v string) *CdsFileShareLinkModel {
	s.Status = &v
	return s
}

func (s *CdsFileShareLinkModel) SetVideoPreviewCount(v int64) *CdsFileShareLinkModel {
	s.VideoPreviewCount = &v
	return s
}

func (s *CdsFileShareLinkModel) Validate() error {
	return dara.Validate(s)
}
