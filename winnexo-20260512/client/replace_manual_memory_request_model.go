// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ReplaceManualMemoryRequest
	GetAppToken() *string
	SetContent(v string) *ReplaceManualMemoryRequest
	GetContent() *string
	SetGroupId(v string) *ReplaceManualMemoryRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *ReplaceManualMemoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ReplaceManualMemoryRequest
	GetOpsUserId() *string
	SetOwnerType(v string) *ReplaceManualMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ReplaceManualMemoryRequest
	GetTenantId() *string
}

type ReplaceManualMemoryRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The complete content of profile.md. An empty string indicates that the memory is cleared. This field must be provided, but an empty value is allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prefers concise answers
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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
	// The owner type. Valid values:
	//
	// - user: the personal memory of the authenticated user.
	//
	// - group: the memory of the specified workspace.
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

func (s ReplaceManualMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryRequest) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ReplaceManualMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *ReplaceManualMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReplaceManualMemoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ReplaceManualMemoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ReplaceManualMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ReplaceManualMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceManualMemoryRequest) SetAppToken(v string) *ReplaceManualMemoryRequest {
	s.AppToken = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetContent(v string) *ReplaceManualMemoryRequest {
	s.Content = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetGroupId(v string) *ReplaceManualMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetOpsUserCode(v string) *ReplaceManualMemoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetOpsUserId(v string) *ReplaceManualMemoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetOwnerType(v string) *ReplaceManualMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetTenantId(v string) *ReplaceManualMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) Validate() error {
	return dara.Validate(s)
}
