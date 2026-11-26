// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveOutputFileToResourceRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveOutputFileToResourceRequest
	GetDirectoryId() *string
	SetItemIds(v []*string) *SaveOutputFileToResourceRequest
	GetItemIds() []*string
	SetMode(v string) *SaveOutputFileToResourceRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveOutputFileToResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveOutputFileToResourceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveOutputFileToResourceRequest
	GetTenantId() *string
}

type SaveOutputFileToResourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the target personal folder. If not specified, the user\"s default folder is automatically resolved.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// itemIds
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
	// The save mode. Valid values:
	//
	// - link: Links the resource to the output in a 1:1 idempotent manner. Edits to the output are synchronized to the resource.
	//
	// - copy: Creates a snapshot copy. No limit on the number of copies.
	//
	// example:
	//
	// link
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// string_value
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveOutputFileToResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveOutputFileToResourceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveOutputFileToResourceRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *SaveOutputFileToResourceRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveOutputFileToResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveOutputFileToResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveOutputFileToResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveOutputFileToResourceRequest) SetAppToken(v string) *SaveOutputFileToResourceRequest {
	s.AppToken = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetDirectoryId(v string) *SaveOutputFileToResourceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetItemIds(v []*string) *SaveOutputFileToResourceRequest {
	s.ItemIds = v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetMode(v string) *SaveOutputFileToResourceRequest {
	s.Mode = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetOpsUserCode(v string) *SaveOutputFileToResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetOpsUserId(v string) *SaveOutputFileToResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) SetTenantId(v string) *SaveOutputFileToResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveOutputFileToResourceRequest) Validate() error {
	return dara.Validate(s)
}
