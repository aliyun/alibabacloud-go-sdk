// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RenameSourceRequest
	GetAppToken() *string
	SetNewName(v string) *RenameSourceRequest
	GetNewName() *string
	SetOpsUserCode(v string) *RenameSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RenameSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *RenameSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *RenameSourceRequest
	GetTenantId() *string
}

type RenameSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new name of the data source.
	//
	// example:
	//
	// string_value
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
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
	// The data source ID, which is unique within the tenant.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. You can pass this parameter explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RenameSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameSourceRequest) GoString() string {
	return s.String()
}

func (s *RenameSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RenameSourceRequest) GetNewName() *string {
	return s.NewName
}

func (s *RenameSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RenameSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RenameSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RenameSourceRequest) SetAppToken(v string) *RenameSourceRequest {
	s.AppToken = &v
	return s
}

func (s *RenameSourceRequest) SetNewName(v string) *RenameSourceRequest {
	s.NewName = &v
	return s
}

func (s *RenameSourceRequest) SetOpsUserCode(v string) *RenameSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RenameSourceRequest) SetOpsUserId(v string) *RenameSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *RenameSourceRequest) SetSourceId(v string) *RenameSourceRequest {
	s.SourceId = &v
	return s
}

func (s *RenameSourceRequest) SetTenantId(v string) *RenameSourceRequest {
	s.TenantId = &v
	return s
}

func (s *RenameSourceRequest) Validate() error {
	return dara.Validate(s)
}
