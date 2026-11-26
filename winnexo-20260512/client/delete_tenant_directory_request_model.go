// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *DeleteTenantDirectoryRequest
	GetAppToken() *string
	SetDeleteMode(v string) *DeleteTenantDirectoryRequest
	GetDeleteMode() *string
	SetDirectoryId(v string) *DeleteTenantDirectoryRequest
	GetDirectoryId() *string
	SetOpsUserCode(v string) *DeleteTenantDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *DeleteTenantDirectoryRequest
	GetOpsUserId() *string
	SetTenantId(v string) *DeleteTenantDirectoryRequest
	GetTenantId() *string
}

type DeleteTenantDirectoryRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The deletion mode: reject / recursive / move_to_root.
	//
	// example:
	//
	// reject
	DeleteMode *string `json:"deleteMode,omitempty" xml:"deleteMode,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// The tenant ID that takes effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *DeleteTenantDirectoryRequest) GetDeleteMode() *string {
	return s.DeleteMode
}

func (s *DeleteTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteTenantDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *DeleteTenantDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *DeleteTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteTenantDirectoryRequest) SetAppToken(v string) *DeleteTenantDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetDeleteMode(v string) *DeleteTenantDirectoryRequest {
	s.DeleteMode = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetDirectoryId(v string) *DeleteTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetOpsUserCode(v string) *DeleteTenantDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetOpsUserId(v string) *DeleteTenantDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) SetTenantId(v string) *DeleteTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
