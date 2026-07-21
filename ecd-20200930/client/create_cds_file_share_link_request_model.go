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
	// The enterprise cloud disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The share description. Length range: 0 to 1024 characters.
	//
	// example:
	//
	// SharedFile
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable downloading of files in the share. Valid values:
	//
	// - true: Downloading is disabled.
	//
	// - false: Downloading is enabled.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"DisableDownload,omitempty" xml:"DisableDownload,omitempty"`
	// Specifies whether to disable previewing of files in the share. Valid values:
	//
	// - true: Preview is disabled.
	//
	// - false: Preview is enabled.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"DisablePreview,omitempty" xml:"DisablePreview,omitempty"`
	// Specifies whether to disable saving of files in the share. Valid values:
	//
	// - true: Saving is disabled.
	//
	// - false: Saving is enabled.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"DisableSave,omitempty" xml:"DisableSave,omitempty"`
	// The maximum number of times the shared files can be downloaded. The value is an integer. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"DownloadLimit,omitempty" xml:"DownloadLimit,omitempty"`
	// The ID of the user who uses the cloud disk.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The expiration time in RFC 3339 format. If this parameter is left empty, the share is permanently valid.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The list of file IDs.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
	// The team space ID.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of times the shared files can be previewed. The value is an integer. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"PreviewLimit,omitempty" xml:"PreviewLimit,omitempty"`
	// The maximum number of times the shared files can be saved. The value is an integer. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"SaveLimit,omitempty" xml:"SaveLimit,omitempty"`
	// The share name. If this parameter is not set, the file name corresponding to the first ID in `file_id_list` is used by default. Length range: 0 to 128 characters.
	//
	// example:
	//
	// view.txt
	ShareName *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	// The extraction code. Length range: 0 to 64 bytes. If this parameter is not set or is set to empty, no extraction code is required, and you do not need to specify the extraction code parameter when obtaining the share token. Only printable ASCII characters are allowed.
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
