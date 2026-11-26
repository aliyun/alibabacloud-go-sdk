// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToGroupResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetAppToken() *string
	SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetDirectoryId() *string
	SetGroupId(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetGroupId() *string
	SetItemIds(v []*string) *SaveGroupOutputFileToGroupResourceRequest
	GetItemIds() []*string
	SetMode(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetMode() *string
	SetOpsUserCode(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetOpsUserId() *string
	SetTenantId(v string) *SaveGroupOutputFileToGroupResourceRequest
	GetTenantId() *string
}

type SaveGroupOutputFileToGroupResourceRequest struct {
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
	ItemIds []*string `json:"itemIds,omitempty" xml:"itemIds,omitempty" type:"Repeated"`
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

func (s SaveGroupOutputFileToGroupResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetMode() *string {
	return s.Mode
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGroupOutputFileToGroupResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetAppToken(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetDirectoryId(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.DirectoryId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetGroupId(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.GroupId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetItemIds(v []*string) *SaveGroupOutputFileToGroupResourceRequest {
	s.ItemIds = v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetMode(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.Mode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetOpsUserCode(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetOpsUserId(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) SetTenantId(v string) *SaveGroupOutputFileToGroupResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceRequest) Validate() error {
	return dara.Validate(s)
}
