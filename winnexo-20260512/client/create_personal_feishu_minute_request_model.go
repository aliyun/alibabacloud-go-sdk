// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuMinuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialId(v string) *CreatePersonalFeishuMinuteRequest
	GetCredentialId() *string
	SetDescription(v string) *CreatePersonalFeishuMinuteRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuMinuteRequest
	GetDirectoryId() *string
	SetMinuteToken(v string) *CreatePersonalFeishuMinuteRequest
	GetMinuteToken() *string
	SetName(v string) *CreatePersonalFeishuMinuteRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalFeishuMinuteRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalFeishuMinuteRequest
	GetTenantId() *string
}

type CreatePersonalFeishuMinuteRequest struct {
	// 凭证 ID（关联 rbj_credential 表，必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCredentialId
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// 资源描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 目标个人目录 ID；不传时自动绑定到当前数字员工默认根目录，传入时必须是当前用户在当前数字员工下的已有个人目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 飞书妙记 token（妙记唯一标识符，必填）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	MinuteToken *string `json:"minuteToken,omitempty" xml:"minuteToken,omitempty"`
	// 资源显示名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工名称（已废弃：不再作为个人资源隔离条件，仅保留用于来源追溯）
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalFeishuMinuteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuMinuteRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuMinuteRequest) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreatePersonalFeishuMinuteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuMinuteRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuMinuteRequest) GetMinuteToken() *string {
	return s.MinuteToken
}

func (s *CreatePersonalFeishuMinuteRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuMinuteRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuMinuteRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuMinuteRequest) SetCredentialId(v string) *CreatePersonalFeishuMinuteRequest {
	s.CredentialId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetDescription(v string) *CreatePersonalFeishuMinuteRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetDirectoryId(v string) *CreatePersonalFeishuMinuteRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetMinuteToken(v string) *CreatePersonalFeishuMinuteRequest {
	s.MinuteToken = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetName(v string) *CreatePersonalFeishuMinuteRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuMinuteRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) SetTenantId(v string) *CreatePersonalFeishuMinuteRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteRequest) Validate() error {
	return dara.Validate(s)
}
