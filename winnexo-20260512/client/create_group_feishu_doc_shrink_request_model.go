// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupFeishuDocShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuDocShrinkRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateGroupFeishuDocShrinkRequest
	GetDocUrl() *string
	SetGroupId(v string) *CreateGroupFeishuDocShrinkRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupFeishuDocShrinkRequest
	GetName() *string
	SetNotes(v string) *CreateGroupFeishuDocShrinkRequest
	GetNotes() *string
	SetObjectBindingsShrink(v string) *CreateGroupFeishuDocShrinkRequest
	GetObjectBindingsShrink() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuDocShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuDocShrinkRequest
	GetSourceTags() *string
	SetSyncConfigShrink(v string) *CreateGroupFeishuDocShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreateGroupFeishuDocShrinkRequest
	GetTenantId() *string
}

type CreateGroupFeishuDocShrinkRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// Group collaboration document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dir_group_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The project group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_delivery
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The image name.
	//
	// example:
	//
	// ProjectPlan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// Extract decisions and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The name of the operating object.
	//
	// example:
	//
	// R&D Assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The resource tags (optional, a JSON string list, such as ["tagA","tagB"]).
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The synchronization settings.
	SyncConfigShrink *string `json:"syncConfig,omitempty" xml:"syncConfig,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupFeishuDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuDocShrinkRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateGroupFeishuDocShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuDocShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupFeishuDocShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuDocShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *CreateGroupFeishuDocShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuDocShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuDocShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreateGroupFeishuDocShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDescription(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDirectoryId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetDocUrl(v string) *CreateGroupFeishuDocShrinkRequest {
	s.DocUrl = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetGroupId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetName(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetNotes(v string) *CreateGroupFeishuDocShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetObjectBindingsShrink(v string) *CreateGroupFeishuDocShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetOperatingObjectName(v string) *CreateGroupFeishuDocShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetSourceTags(v string) *CreateGroupFeishuDocShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetSyncConfigShrink(v string) *CreateGroupFeishuDocShrinkRequest {
	s.SyncConfigShrink = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) SetTenantId(v string) *CreateGroupFeishuDocShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
