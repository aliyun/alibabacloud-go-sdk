// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetSyncConfigShrink(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreatePersonalAlidingKnowledgeBaseShrinkRequest
	GetTenantId() *string
}

type CreatePersonalAlidingKnowledgeBaseShrinkRequest struct {
	// 目标个人目录 ID；不传时自动绑定到用户默认根目录，传入时必须是当前用户的已有个人目录（PERSONAL）
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 知识库显示名称；不传时由后台从远程拉取的根节点名称回填
	//
	// example:
	//
	// string_value
	KbName *string `json:"kbName,omitempty" xml:"kbName,omitempty"`
	// 阿里钉知识库的可公开访问 URL
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	KbUrl                *string `json:"kbUrl,omitempty" xml:"kbUrl,omitempty"`
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// Agent 命名空间标识，可选
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	SyncConfigShrink    *string `json:"syncConfig,omitempty" xml:"syncConfig,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseShrinkRequest) GoString() string {
	return s.String()
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

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreatePersonalAlidingKnowledgeBaseShrinkRequest) GetTenantId() *string {
	return s.TenantId
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
