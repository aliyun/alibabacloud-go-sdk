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
	AccessCount       *int64  `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator           *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableDownload   *bool   `json:"DisableDownload,omitempty" xml:"DisableDownload,omitempty"`
	DisablePreview    *bool   `json:"DisablePreview,omitempty" xml:"DisablePreview,omitempty"`
	DisableSave       *bool   `json:"DisableSave,omitempty" xml:"DisableSave,omitempty"`
	DownloadCount     *int64  `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	DownloadLimit     *int64  `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	DriveId           *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	Expiration        *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	Expired           *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	FileIds           *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	ModifiyTime       *string `json:"ModifiyTime,omitempty" xml:"ModifiyTime,omitempty"`
	PreviewCount      *int64  `json:"PreviewCount,omitempty" xml:"PreviewCount,omitempty"`
	PreviewLimit      *int64  `json:"PreviewLimit,omitempty" xml:"PreviewLimit,omitempty"`
	ReportCount       *int64  `json:"ReportCount,omitempty" xml:"ReportCount,omitempty"`
	SaveCount         *int64  `json:"SaveCount,omitempty" xml:"SaveCount,omitempty"`
	SaveLimit         *int64  `json:"SaveLimit,omitempty" xml:"SaveLimit,omitempty"`
	ShareId           *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	ShareLink         *string `json:"ShareLink,omitempty" xml:"ShareLink,omitempty"`
	ShareName         *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	SharePwd          *string `json:"SharePwd,omitempty" xml:"SharePwd,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoPreviewCount *int64  `json:"VideoPreviewCount,omitempty" xml:"VideoPreviewCount,omitempty"`
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
