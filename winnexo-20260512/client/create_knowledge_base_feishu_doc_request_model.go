// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFeishuDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetDocUrl() *string
	SetName(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetName() *string
	SetNotes(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetNotes() *string
	SetObjectBindings(v []*CreateKnowledgeBaseFeishuDocRequestObjectBindings) *CreateKnowledgeBaseFeishuDocRequest
	GetObjectBindings() []*CreateKnowledgeBaseFeishuDocRequestObjectBindings
	SetOperatingObjectName(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetSourceTags() *string
	SetSyncConfig(v *CreateKnowledgeBaseFeishuDocRequestSyncConfig) *CreateKnowledgeBaseFeishuDocRequest
	GetSyncConfig() *CreateKnowledgeBaseFeishuDocRequestSyncConfig
	SetTenantId(v string) *CreateKnowledgeBaseFeishuDocRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseFeishuDocRequest struct {
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
	ObjectBindings []*CreateKnowledgeBaseFeishuDocRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	SyncConfig *CreateKnowledgeBaseFeishuDocRequestSyncConfig `json:"syncConfig,omitempty" xml:"syncConfig,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetObjectBindings() []*CreateKnowledgeBaseFeishuDocRequestObjectBindings {
	return s.ObjectBindings
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetSyncConfig() *CreateKnowledgeBaseFeishuDocRequestSyncConfig {
	return s.SyncConfig
}

func (s *CreateKnowledgeBaseFeishuDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetDescription(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetDirectoryId(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetDocUrl(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.DocUrl = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetName(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetNotes(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.Notes = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetObjectBindings(v []*CreateKnowledgeBaseFeishuDocRequestObjectBindings) *CreateKnowledgeBaseFeishuDocRequest {
	s.ObjectBindings = v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetSourceTags(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetSyncConfig(v *CreateKnowledgeBaseFeishuDocRequestSyncConfig) *CreateKnowledgeBaseFeishuDocRequest {
	s.SyncConfig = v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) SetTenantId(v string) *CreateKnowledgeBaseFeishuDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequest) Validate() error {
	if s.ObjectBindings != nil {
		for _, item := range s.ObjectBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SyncConfig != nil {
		if err := s.SyncConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseFeishuDocRequestObjectBindings struct {
	// The name of the semantic graph to which the object belongs.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The ID of the recommended item, which can be a **feedId*	- or a mini-app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The advanced field type.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) SetGraphName(v string) *CreateKnowledgeBaseFeishuDocRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) SetObjectId(v string) *CreateKnowledgeBaseFeishuDocRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) SetObjectType(v string) *CreateKnowledgeBaseFeishuDocRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseFeishuDocRequestSyncConfig struct {
	// The cron expression for the timed scheduling task.
	//
	// example:
	//
	// 0 	- 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether to enable or disable synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The synchronization preset: hourly or daily_2am.
	//
	// example:
	//
	// custom
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreateKnowledgeBaseFeishuDocRequestSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFeishuDocRequestSyncConfig) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) GetCron() *string {
	return s.Cron
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) GetPreset() *string {
	return s.Preset
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) SetCron(v string) *CreateKnowledgeBaseFeishuDocRequestSyncConfig {
	s.Cron = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) SetEnabled(v bool) *CreateKnowledgeBaseFeishuDocRequestSyncConfig {
	s.Enabled = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) SetPreset(v string) *CreateKnowledgeBaseFeishuDocRequestSyncConfig {
	s.Preset = &v
	return s
}

func (s *CreateKnowledgeBaseFeishuDocRequestSyncConfig) Validate() error {
	return dara.Validate(s)
}
