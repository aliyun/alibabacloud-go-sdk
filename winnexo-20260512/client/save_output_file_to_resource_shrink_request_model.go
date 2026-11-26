// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveOutputFileToResourceShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveOutputFileToResourceShrinkRequest
	GetDirectoryId() *string
	SetItemIdsShrink(v string) *SaveOutputFileToResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveOutputFileToResourceShrinkRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveOutputFileToResourceShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveOutputFileToResourceShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveOutputFileToResourceShrinkRequest
	GetTenantId() *string
}

type SaveOutputFileToResourceShrinkRequest struct {
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
	ItemIdsShrink *string `json:"itemIds,omitempty" xml:"itemIds,omitempty"`
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

func (s SaveOutputFileToResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveOutputFileToResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveOutputFileToResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveOutputFileToResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveOutputFileToResourceShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveOutputFileToResourceShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveOutputFileToResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveOutputFileToResourceShrinkRequest) SetAppToken(v string) *SaveOutputFileToResourceShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetDirectoryId(v string) *SaveOutputFileToResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetItemIdsShrink(v string) *SaveOutputFileToResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetMode(v string) *SaveOutputFileToResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetOpsUserCode(v string) *SaveOutputFileToResourceShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetOpsUserId(v string) *SaveOutputFileToResourceShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) SetTenantId(v string) *SaveOutputFileToResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveOutputFileToResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
