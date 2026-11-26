// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ClearAutoMemoryRequest
	GetAppToken() *string
	SetGroupId(v string) *ClearAutoMemoryRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *ClearAutoMemoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ClearAutoMemoryRequest
	GetOpsUserId() *string
	SetOwnerType(v string) *ClearAutoMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ClearAutoMemoryRequest
	GetTenantId() *string
	SetType(v string) *ClearAutoMemoryRequest
	GetType() *string
}

type ClearAutoMemoryRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID. This parameter is required when ownerType is set to group.
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// The owner type. Valid values: user: the personal memory of the authenticated user; group: the memory of the specified collaboration workspace.
	//
	// example:
	//
	// group
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The memory type. Valid values: user, feedback, reference, project, or all. The value all does not clear manual memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ClearAutoMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryRequest) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ClearAutoMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ClearAutoMemoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ClearAutoMemoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ClearAutoMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ClearAutoMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearAutoMemoryRequest) GetType() *string {
	return s.Type
}

func (s *ClearAutoMemoryRequest) SetAppToken(v string) *ClearAutoMemoryRequest {
	s.AppToken = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetGroupId(v string) *ClearAutoMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetOpsUserCode(v string) *ClearAutoMemoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetOpsUserId(v string) *ClearAutoMemoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetOwnerType(v string) *ClearAutoMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetTenantId(v string) *ClearAutoMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetType(v string) *ClearAutoMemoryRequest {
	s.Type = &v
	return s
}

func (s *ClearAutoMemoryRequest) Validate() error {
	return dara.Validate(s)
}
