// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CreateCdsFileShareLinkRequest
	GetCdsId() *string
	SetDescription(v string) *CreateCdsFileShareLinkRequest
	GetDescription() *string
	SetDisableDownload(v bool) *CreateCdsFileShareLinkRequest
	GetDisableDownload() *bool
	SetDisablePreview(v bool) *CreateCdsFileShareLinkRequest
	GetDisablePreview() *bool
	SetDisableSave(v bool) *CreateCdsFileShareLinkRequest
	GetDisableSave() *bool
	SetDownloadLimit(v int64) *CreateCdsFileShareLinkRequest
	GetDownloadLimit() *int64
	SetEndUserId(v string) *CreateCdsFileShareLinkRequest
	GetEndUserId() *string
	SetExpiration(v string) *CreateCdsFileShareLinkRequest
	GetExpiration() *string
	SetFileIds(v []*string) *CreateCdsFileShareLinkRequest
	GetFileIds() []*string
	SetGroupId(v string) *CreateCdsFileShareLinkRequest
	GetGroupId() *string
	SetPreviewLimit(v int64) *CreateCdsFileShareLinkRequest
	GetPreviewLimit() *int64
	SetSaveLimit(v int64) *CreateCdsFileShareLinkRequest
	GetSaveLimit() *int64
	SetShareName(v string) *CreateCdsFileShareLinkRequest
	GetShareName() *string
	SetSharePwd(v string) *CreateCdsFileShareLinkRequest
	GetSharePwd() *string
}

type CreateCdsFileShareLinkRequest struct {
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The description of the file sharing task. The description must be 0 to 1,024 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to prohibit the download of the files that are being shared.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     prohibits file download
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     allows file download
	//
	//     <!-- -->
	//
	//     .
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
	//     :
	//
	//     <!-- -->
	//
	//     prohibits file preview
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     allows file preview
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"DisablePreview,omitempty" xml:"DisablePreview,omitempty"`
	// Specifies whether to prohibit the dump of the files that are being shared.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     prohibits file dump
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     allows file dump
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// false
	DisableSave *bool `json:"DisableSave,omitempty" xml:"DisableSave,omitempty"`
	// The limit on the number of times that the shared files can be downloaded. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the file sharing link expires. The value of this parameter follows the RFC 3339 standard. Example: "2020-06-28T11:33:00.000+08:00". If this parameter is set to "", the file sharing link never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The file IDs.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
	GroupId *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The limit on the number of times that the shared files can be previewed. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be previewed.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"PreviewLimit,omitempty" xml:"PreviewLimit,omitempty"`
	// The limit on the number of times that the shared files can be dumped. The value of this parameter must be equal to or greater than 0. The value 0 specifies that no limit is imposed on the number of times that the shared files can be dumped.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"SaveLimit,omitempty" xml:"SaveLimit,omitempty"`
	// The name of the file sharing task. If you leave this parameter empty, the file name that corresponds to the first ID in the file ID list is used. The name must be 0 to 128 characters in length.
	//
	// example:
	//
	// view.txt
	ShareName *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	// The length of the access code. Valid values: 6 to 8. Unit: bytes. If you leave this parameter empty or set it to null, no access code is required. If you use a token to share files, you do not need to configure this parameter. The access code can contain only visible ASCII characters.
	//
	// example:
	//
	// 12345678
	SharePwd *string `json:"SharePwd,omitempty" xml:"SharePwd,omitempty"`
}

func (s CreateCdsFileShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCdsFileShareLinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCdsFileShareLinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCdsFileShareLinkRequest) GetDisableDownload() *bool {
	return s.DisableDownload
}

func (s *CreateCdsFileShareLinkRequest) GetDisablePreview() *bool {
	return s.DisablePreview
}

func (s *CreateCdsFileShareLinkRequest) GetDisableSave() *bool {
	return s.DisableSave
}

func (s *CreateCdsFileShareLinkRequest) GetDownloadLimit() *int64 {
	return s.DownloadLimit
}

func (s *CreateCdsFileShareLinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateCdsFileShareLinkRequest) GetExpiration() *string {
	return s.Expiration
}

func (s *CreateCdsFileShareLinkRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *CreateCdsFileShareLinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateCdsFileShareLinkRequest) GetPreviewLimit() *int64 {
	return s.PreviewLimit
}

func (s *CreateCdsFileShareLinkRequest) GetSaveLimit() *int64 {
	return s.SaveLimit
}

func (s *CreateCdsFileShareLinkRequest) GetShareName() *string {
	return s.ShareName
}

func (s *CreateCdsFileShareLinkRequest) GetSharePwd() *string {
	return s.SharePwd
}

func (s *CreateCdsFileShareLinkRequest) SetCdsId(v string) *CreateCdsFileShareLinkRequest {
	s.CdsId = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetDescription(v string) *CreateCdsFileShareLinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetDisableDownload(v bool) *CreateCdsFileShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetDisablePreview(v bool) *CreateCdsFileShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetDisableSave(v bool) *CreateCdsFileShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetDownloadLimit(v int64) *CreateCdsFileShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetEndUserId(v string) *CreateCdsFileShareLinkRequest {
	s.EndUserId = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetExpiration(v string) *CreateCdsFileShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetFileIds(v []*string) *CreateCdsFileShareLinkRequest {
	s.FileIds = v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetGroupId(v string) *CreateCdsFileShareLinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetPreviewLimit(v int64) *CreateCdsFileShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetSaveLimit(v int64) *CreateCdsFileShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetShareName(v string) *CreateCdsFileShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) SetSharePwd(v string) *CreateCdsFileShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *CreateCdsFileShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
