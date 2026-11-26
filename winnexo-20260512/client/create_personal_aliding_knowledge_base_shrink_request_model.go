// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetAppToken() *string
	SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetDirectoryId() *string
	SetKbName(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetKbName() *string
	SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetKbUrl() *string
	SetObjectBindingsShrink(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetObjectBindingsShrink() *string
	SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetSourceTags() *string
	SetSyncConfigShrink(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetTenantId() *string
}

type CreatePersonalAlidingKnowledgeBaseShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The display name of the knowledge base. If not provided, the name is populated from the root node name pulled from the remote source.
	//
	// example:
	//
	// string_value
	KbName *string `json:"kbName,omitempty" xml:"kbName,omitempty"`
	// The publicly accessible URL of the AliDing knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	KbUrl *string `json:"kbUrl,omitempty" xml:"kbUrl,omitempty"`
	// The object bindings.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
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
	// The list of resource tag JSON strings applied to all child sources created during knowledge base synchronization.
	//
	// example:
	//
	// ["Key","KnowledgeBase"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The synchronization settings.
	SyncConfigShrink *string `json:"syncConfig,omitempty" xml:"syncConfig,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// PiPklI1iSRTm6VFFqlY9VzbgiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetKbName() *string {
	return s.KbName
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetKbUrl() *string {
	return s.KbUrl
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetAppToken(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetKbName(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.KbName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.KbUrl = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetObjectBindingsShrink(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetOpsUserCode(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetOpsUserId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetSourceTags(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetSyncConfigShrink(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.SyncConfigShrink = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) SetTenantId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
