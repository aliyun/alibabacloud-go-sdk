// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *CreatePersonalDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreatePersonalDirectoryRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreatePersonalDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalDirectoryRequest
	GetOpsUserId() *string
	SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreatePersonalDirectoryRequest
	GetTenantId() *string
}

type CreatePersonalDirectoryRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The workspace description.
	//
	// example:
	//
	// hangzhou-release-version-3-eventbridge-numeric-queue-fix-20260529
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the digital employee.
	//
	// This parameter is required.
	//
	// example:
	//
	// sandbox-conversation-webpage-github-default-p32JG2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The folder ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 235454102432001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreatePersonalDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDirectoryRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreatePersonalDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreatePersonalDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDirectoryRequest) SetAppToken(v string) *CreatePersonalDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetDescription(v string) *CreatePersonalDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetName(v string) *CreatePersonalDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetOperatingObjectName(v string) *CreatePersonalDirectoryRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetOpsUserCode(v string) *CreatePersonalDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetOpsUserId(v string) *CreatePersonalDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetParentDirectoryId(v string) *CreatePersonalDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) SetTenantId(v string) *CreatePersonalDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
