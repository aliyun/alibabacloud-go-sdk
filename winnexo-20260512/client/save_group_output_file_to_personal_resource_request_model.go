// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetGroupId() *string
	SetItemIds(v []*string) *SaveGroupOutputFileToPersonalResourceRequest
	GetItemIds() []*string
	SetMode(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToPersonalResourceRequest struct {
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
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
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

func (s SaveGroupOutputFileToPersonalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetAppToken(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetDirectoryId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetGroupId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetItemIds(v []*string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.ItemIds = v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetMode(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetOpsUserCode(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetOpsUserId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) SetTenantId(v string) *SaveGroupOutputFileToPersonalResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceRequest) Validate() error {
	return dara.Validate(s)
}
