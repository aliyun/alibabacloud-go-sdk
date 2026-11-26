// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetMemoryDetailRequest
	GetAppToken() *string
	SetGroupId(v string) *GetMemoryDetailRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *GetMemoryDetailRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetMemoryDetailRequest
	GetOpsUserId() *string
	SetOwnerType(v string) *GetMemoryDetailRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryDetailRequest
	GetTenantId() *string
}

type GetMemoryDetailRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The workspace ID. This parameter is required when ownerType is set to group.
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
	// The owner type. A value of user indicates the personal memory of the authenticated user. A value of group indicates the memory of a specified workspace.
	//
	// example:
	//
	// group
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetMemoryDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetMemoryDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryDetailRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetMemoryDetailRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetMemoryDetailRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryDetailRequest) SetAppToken(v string) *GetMemoryDetailRequest {
	s.AppToken = &v
	return s
}

func (s *GetMemoryDetailRequest) SetGroupId(v string) *GetMemoryDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryDetailRequest) SetOpsUserCode(v string) *GetMemoryDetailRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetMemoryDetailRequest) SetOpsUserId(v string) *GetMemoryDetailRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetMemoryDetailRequest) SetOwnerType(v string) *GetMemoryDetailRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryDetailRequest) SetTenantId(v string) *GetMemoryDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryDetailRequest) Validate() error {
	return dara.Validate(s)
}
