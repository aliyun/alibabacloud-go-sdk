// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFeishuDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetDocUrl() *string
	SetName(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetName() *string
	SetNotes(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetNotes() *string
	SetObjectBindingsShrink(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetObjectBindingsShrink() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetSourceTags() *string
	SetSyncConfigShrink(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseFeishuDocShrinkRequest struct {
	// The description of the alias.
	//
	// example:
	//
	// Enterprise policy document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_tenant_kb_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The mirror name.
	//
	// example:
	//
	// Enterprise Policy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). Used for auxiliary analysis.
	//
	// example:
	//
	// Extract applicable scope and key clauses
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// Enterprise Knowledge Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The resource tags (optional, a JSON string list such as ["tagA","tagB"]).
	//
	// example:
	//
	// ["policy"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The synchronization settings.
	SyncConfigShrink *string `json:"syncConfig,omitempty" xml:"syncConfig,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetDescription(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.DocUrl = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetName(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetNotes(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetObjectBindingsShrink(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetSourceTags(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetSyncConfigShrink(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.SyncConfigShrink = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) SetTenantId(v string) *CreateKnowledgeBaseFeishuDocShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
