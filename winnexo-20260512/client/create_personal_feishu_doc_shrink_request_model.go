// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalFeishuDocShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuDocShrinkRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreatePersonalFeishuDocShrinkRequest
	GetDocUrl() *string
	SetName(v string) *CreatePersonalFeishuDocShrinkRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalFeishuDocShrinkRequest
	GetNotes() *string
	SetObjectBindingsShrink(v string) *CreatePersonalFeishuDocShrinkRequest
	GetObjectBindingsShrink() *string
	SetOperatingObjectName(v string) *CreatePersonalFeishuDocShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalFeishuDocShrinkRequest
	GetSourceTags() *string
	SetSyncConfigShrink(v string) *CreatePersonalFeishuDocShrinkRequest
	GetSyncConfigShrink() *string
	SetTenantId(v string) *CreatePersonalFeishuDocShrinkRequest
	GetTenantId() *string
}

type CreatePersonalFeishuDocShrinkRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// Project design document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_personal_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The document URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.feishu.cn/docx/doxcnExample
	DocUrl *string `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	// The updated name of the filter view.
	//
	// example:
	//
	// ProjectPlan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// Extract decisions and to-dos
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindingsShrink *string `json:"objectBindings,omitempty" xml:"objectBindings,omitempty"`
	// The digital employee name (operating object name, optional).
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
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalFeishuDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetObjectBindingsShrink() *string {
	return s.ObjectBindingsShrink
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetSyncConfigShrink() *string {
	return s.SyncConfigShrink
}

func (s *CreatePersonalFeishuDocShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetDescription(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetDirectoryId(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetDocUrl(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.DocUrl = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetName(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetNotes(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetObjectBindingsShrink(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.ObjectBindingsShrink = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetSourceTags(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetSyncConfigShrink(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.SyncConfigShrink = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) SetTenantId(v string) *CreatePersonalFeishuDocShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
