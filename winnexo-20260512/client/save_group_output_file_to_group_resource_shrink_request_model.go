// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToGroupResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetGroupId() *string
	SetItemIdsShrink(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToGroupResourceShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the target personal directory. If not specified, the user\"s default directory is automatically resolved.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// itemIds
	//
	// This parameter is required.
	//
	// example:
	//
	// ["item-1","item-2"]
	ItemIdsShrink *string `json:"itemIds,omitempty" xml:"itemIds,omitempty"`
	// The save mode. Valid values:
	//
	// - link: creates a link (1:1 idempotent, editing the output synchronizes the resource).
	//
	// - copy: creates a copy (unlimited times, snapshot).
	//
	// This parameter is required.
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
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveGroupOutputFileToGroupResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetAppToken(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetGroupId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetItemIdsShrink(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetMode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetOpsUserCode(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetOpsUserId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) SetTenantId(v string) *SaveGroupOutputFileToGroupResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
