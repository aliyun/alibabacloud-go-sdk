// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetAppToken() *string
	SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetClearExisting() *bool
	SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetEnabled() *bool
	SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetOpsUserId() *string
	SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type UpdateMemoryAutoExtractionSwitchRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to synchronously clear the automatically collected memory. This parameter can be used when enabling or disabling the feature.
	//
	// example:
	//
	// true
	ClearExisting *bool `json:"clearExisting,omitempty" xml:"clearExisting,omitempty"`
	// Specifies whether to enable automatic memory collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The ID of the collaboration workspace. This parameter is required when ownerType is set to group.
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
	// - user: the personal memory of the authenticated user
	//
	// - group: the memory of the specified collaboration workspace
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

func (s UpdateMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetClearExisting() *bool {
	return s.ClearExisting
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetAppToken(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.ClearExisting = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetOpsUserCode(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetOpsUserId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
