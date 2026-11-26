// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetMemoryAutoExtractionSwitchRequest
	GetAppToken() *string
	SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *GetMemoryAutoExtractionSwitchRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetOpsUserId() *string
	SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type GetMemoryAutoExtractionSwitchRequest struct {
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
	// The owner type. A value of user indicates the personal memory of the authenticated user. A value of group indicates the memory of the specified collaboration workspace.
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
}

func (s GetMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetAppToken(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.AppToken = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetOpsUserCode(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetOpsUserId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
