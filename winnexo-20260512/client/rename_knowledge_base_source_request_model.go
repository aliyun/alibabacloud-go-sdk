// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameKnowledgeBaseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RenameKnowledgeBaseSourceRequest
	GetAppToken() *string
	SetNewName(v string) *RenameKnowledgeBaseSourceRequest
	GetNewName() *string
	SetOpsUserCode(v string) *RenameKnowledgeBaseSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RenameKnowledgeBaseSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *RenameKnowledgeBaseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *RenameKnowledgeBaseSourceRequest
	GetTenantId() *string
}

type RenameKnowledgeBaseSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The new name of the data source.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RenameKnowledgeBaseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameKnowledgeBaseSourceRequest) GoString() string {
	return s.String()
}

func (s *RenameKnowledgeBaseSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RenameKnowledgeBaseSourceRequest) GetNewName() *string {
	return s.NewName
}

func (s *RenameKnowledgeBaseSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RenameKnowledgeBaseSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RenameKnowledgeBaseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameKnowledgeBaseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RenameKnowledgeBaseSourceRequest) SetAppToken(v string) *RenameKnowledgeBaseSourceRequest {
	s.AppToken = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetNewName(v string) *RenameKnowledgeBaseSourceRequest {
	s.NewName = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetOpsUserCode(v string) *RenameKnowledgeBaseSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetOpsUserId(v string) *RenameKnowledgeBaseSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetSourceId(v string) *RenameKnowledgeBaseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetTenantId(v string) *RenameKnowledgeBaseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) Validate() error {
	return dara.Validate(s)
}
