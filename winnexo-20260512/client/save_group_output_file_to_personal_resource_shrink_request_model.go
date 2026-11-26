// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetGroupId() *string
	SetItemIdsShrink(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetItemIdsShrink() *string
	SetMode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToPersonalResourceShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The enterprise knowledge base directory ID.
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
	// - link: Links the output. This mode is idempotent on a 1:1 basis. Editing the output synchronizes the resource.
	//
	// - copy: Copies the output. This mode allows unlimited copies and creates a snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// copy
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
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SaveGroupOutputFileToPersonalResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetItemIdsShrink() *string {
	return s.ItemIdsShrink
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetAppToken(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetItemIdsShrink(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetMode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetOpsUserCode(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetOpsUserId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
