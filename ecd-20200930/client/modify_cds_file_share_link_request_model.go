// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdsFileShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ModifyCdsFileShareLinkRequest
	GetCdsId() *string
	SetDescription(v string) *ModifyCdsFileShareLinkRequest
	GetDescription() *string
	SetDisableDownload(v bool) *ModifyCdsFileShareLinkRequest
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *ModifyCdsFileShareLinkRequest
	GetDisablePreview() *bool
	SetDisableSave(v bool) *ModifyCdsFileShareLinkRequest
	GetDisableSave() *bool
	SetDownloadCount(v int64) *ModifyCdsFileShareLinkRequest
	GetDownloadCount() *int64
	SetDownloadLimit(v int64) *ModifyCdsFileShareLinkRequest
	GetDownloadLimit() *int64
	SetExpiration(v string) *ModifyCdsFileShareLinkRequest
	GetExpiration() *string
	SetPreviewCount(v int64) *ModifyCdsFileShareLinkRequest
	GetPreviewCount() *int64
	SetPreviewLimit(v int64) *ModifyCdsFileShareLinkRequest
	GetPreviewLimit() *int64
	SetReportCount(v int64) *ModifyCdsFileShareLinkRequest
	GetReportCount() *int64
	SetSaveCount(v int64) *ModifyCdsFileShareLinkRequest
	GetSaveCount() *int64
	SetSaveLimit(v int64) *ModifyCdsFileShareLinkRequest
	GetSaveLimit() *int64
	SetShareId(v string) *ModifyCdsFileShareLinkRequest
	GetShareId() *string
	SetShareName(v string) *ModifyCdsFileShareLinkRequest
	GetShareName() *string
	SetSharePwd(v string) *ModifyCdsFileShareLinkRequest
	GetSharePwd() *string
	SetStatus(v string) *ModifyCdsFileShareLinkRequest
	GetStatus() *string
	SetVideoPreviewCount(v int64) *ModifyCdsFileShareLinkRequest
	GetVideoPreviewCount() *int64
}

type ModifyCdsFileShareLinkRequest struct {
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The description of the file sharing task. The description must be 0 to 1,024 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to prohibit the download of the files that are being shared.
	//
	// Valid values:
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"DisableDownload,omitempty" xml:"DisableDownload,omitempty"`
	// Specifies whether to prohibit the preview of the files that are being shared.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"DisablePreview,omitempty" xml:"DisablePreview,omitempty"`
	// Specifies whether to prohibit the dump of the files that are being shared.
	//
	// Valid values:
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	DisableSave *bool `json:"DisableSave,omitempty" xml:"DisableSave,omitempty"`
	// The number of times that the shared files are downloaded. The value of this parameter must be equal to or greater than 0.
	//
	// example:
	//
	// 0
	DownloadCount *int64 `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	// The limit on the number of times that the shared files can be downloaded. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 0
	DownloadLimit *int64 `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	// The time when the file sharing link expires. The value of this parameter follows the RFC 3339 standard. Example: "2020-06-28T11:33:00.000+08:00". If this parameter is set to "", the file sharing link never expires.
	//
	// example:
	//
	// 2022-07-20T06:30:22.365Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The number of times that the shared files are previewed. The value of this parameter must be equal to or greater than 0.
	//
	// example:
	//
	// 0
	PreviewCount *int64 `json:"PreviewCount,omitempty" xml:"PreviewCount,omitempty"`
	// The limit on the number of times that the shared files can be previewed. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"PreviewLimit,omitempty" xml:"PreviewLimit,omitempty"`
	// The number of times that the shared files are reported. The value of this parameter must be equal to or greater than 0.
	//
	// example:
	//
	// 0
	ReportCount *int64 `json:"ReportCount,omitempty" xml:"ReportCount,omitempty"`
	// The number of times that the shared files are dumped. The value of this parameter must be equal to or greater than 0.
	//
	// example:
	//
	// 0
	SaveCount *int64 `json:"SaveCount,omitempty" xml:"SaveCount,omitempty"`
	// The limit on the number of times that the shared files can be dumped. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"SaveLimit,omitempty" xml:"SaveLimit,omitempty"`
	// The ID of the file sharing task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1Fs****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	// The name of the file sharing task. If you do not configure this parameter, the sharing task name is the first ID that is returned in the file_id_list value.
	//
	// >  The sharing task name must be 0 to 128 characters in length.
	ShareName *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	// The length of the access code. Valid values: 6 to 8. Unit: bytes. If you leave this parameter empty or set it to null, no access code is required. If you use a token to share files, you do not need to configure this parameter. The access code can contain only visible ASCII characters.
	//
	// example:
	//
	// 12345678
	SharePwd *string `json:"SharePwd,omitempty" xml:"SharePwd,omitempty"`
	// The sharing status.
	//
	// Valid values:
	//
	// 	- disabled: The sharing task is canceled.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- enabled: The sharing task is valid.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of times that the videos are previewed in the shared files. The value of this parameter must be equal to or greater than 0.
	//
	// example:
	//
	// 0
	VideoPreviewCount *int64 `json:"VideoPreviewCount,omitempty" xml:"VideoPreviewCount,omitempty"`
}

func (s ModifyCdsFileShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdsFileShareLinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdsFileShareLinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ModifyCdsFileShareLinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCdsFileShareLinkRequest) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *ModifyCdsFileShareLinkRequest) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *ModifyCdsFileShareLinkRequest) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *ModifyCdsFileShareLinkRequest) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *ModifyCdsFileShareLinkRequest) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *ModifyCdsFileShareLinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *ModifyCdsFileShareLinkRequest) GetPreviewCount() *int64 {
	return s.PreviewCount
}

func (s *ModifyCdsFileShareLinkRequest) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *ModifyCdsFileShareLinkRequest) GetReportCount() *int64 {
	return s.ReportCount
}

func (s *ModifyCdsFileShareLinkRequest) GetSaveCount() *int64 {
	return s.SaveCount
}

func (s *ModifyCdsFileShareLinkRequest) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *ModifyCdsFileShareLinkRequest) GetShareId() *string {
	return s.ShareId
}

func (s *ModifyCdsFileShareLinkRequest) GetShareName() *string {
	return s.ShareName
}

func (s *ModifyCdsFileShareLinkRequest) GetSharePwd() *string {
	return s.SharePwd
}

func (s *ModifyCdsFileShareLinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyCdsFileShareLinkRequest) GetVideoPreviewCount() *int64 {
	return s.VideoPreviewCount
}

func (s *ModifyCdsFileShareLinkRequest) SetCdsId(v string) *ModifyCdsFileShareLinkRequest {
	s.CdsId = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDescription(v string) *ModifyCdsFileShareLinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDisableDownload(v bool) *ModifyCdsFileShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDisablePreview(v bool) *ModifyCdsFileShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDisableSave(v bool) *ModifyCdsFileShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDownloadCount(v int64) *ModifyCdsFileShareLinkRequest {
	s.DownloadCount = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetDownloadLimit(v int64) *ModifyCdsFileShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetExpiration(v string) *ModifyCdsFileShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetPreviewCount(v int64) *ModifyCdsFileShareLinkRequest {
	s.PreviewCount = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetPreviewLimit(v int64) *ModifyCdsFileShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetReportCount(v int64) *ModifyCdsFileShareLinkRequest {
	s.ReportCount = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetSaveCount(v int64) *ModifyCdsFileShareLinkRequest {
	s.SaveCount = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetSaveLimit(v int64) *ModifyCdsFileShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetShareId(v string) *ModifyCdsFileShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetShareName(v string) *ModifyCdsFileShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetSharePwd(v string) *ModifyCdsFileShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetStatus(v string) *ModifyCdsFileShareLinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) SetVideoPreviewCount(v int64) *ModifyCdsFileShareLinkRequest {
	s.VideoPreviewCount = &v
	return s
}

func (s *ModifyCdsFileShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
