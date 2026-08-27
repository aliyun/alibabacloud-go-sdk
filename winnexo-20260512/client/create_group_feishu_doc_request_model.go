// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupFeishuDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuDocRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreateGroupFeishuDocRequest
	GetDocUrl() *string
	SetGroupId(v string) *CreateGroupFeishuDocRequest
	GetGroupId() *string
	SetName(v string) *CreateGroupFeishuDocRequest
	GetName() *string
	SetNotes(v string) *CreateGroupFeishuDocRequest
	GetNotes() *string
	SetObjectBindings(v []*CreateGroupFeishuDocRequestObjectBindings) *CreateGroupFeishuDocRequest
	GetObjectBindings() []*CreateGroupFeishuDocRequestObjectBindings
	SetOperatingObjectName(v string) *CreateGroupFeishuDocRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuDocRequest
	GetSourceTags() *string
	SetSyncConfig(v *CreateGroupFeishuDocRequestSyncConfig) *CreateGroupFeishuDocRequest
	GetSyncConfig() *CreateGroupFeishuDocRequestSyncConfig
	SetTenantId(v string) *CreateGroupFeishuDocRequest
	GetTenantId() *string
}

type CreateGroupFeishuDocRequest struct {
	// The description of the AI assistant.
	//
	// example:
	//
	// Group collaboration document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
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
	// Project Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// Extract decisions and to-do items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The object bindings.
	ObjectBindings []*CreateGroupFeishuDocRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	SyncConfig *CreateGroupFeishuDocRequestSyncConfig `json:"syncConfig,omitempty" xml:"syncConfig,omitempty" type:"Struct"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateGroupFeishuDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuDocRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateGroupFeishuDocRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuDocRequest) GetName() *string {
	return s.Name
}

func (s *CreateGroupFeishuDocRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuDocRequest) GetObjectBindings() []*CreateGroupFeishuDocRequestObjectBindings {
	return s.ObjectBindings
}

func (s *CreateGroupFeishuDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuDocRequest) GetSyncConfig() *CreateGroupFeishuDocRequestSyncConfig {
	return s.SyncConfig
}

func (s *CreateGroupFeishuDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuDocRequest) SetDescription(v string) *CreateGroupFeishuDocRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetDirectoryId(v string) *CreateGroupFeishuDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetDocUrl(v string) *CreateGroupFeishuDocRequest {
	s.DocUrl = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetGroupId(v string) *CreateGroupFeishuDocRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetName(v string) *CreateGroupFeishuDocRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetNotes(v string) *CreateGroupFeishuDocRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetObjectBindings(v []*CreateGroupFeishuDocRequestObjectBindings) *CreateGroupFeishuDocRequest {
	s.ObjectBindings = v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetOperatingObjectName(v string) *CreateGroupFeishuDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetSourceTags(v string) *CreateGroupFeishuDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetSyncConfig(v *CreateGroupFeishuDocRequestSyncConfig) *CreateGroupFeishuDocRequest {
	s.SyncConfig = v
	return s
}

func (s *CreateGroupFeishuDocRequest) SetTenantId(v string) *CreateGroupFeishuDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuDocRequest) Validate() error {
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

type CreateGroupFeishuDocRequestObjectBindings struct {
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

func (s CreateGroupFeishuDocRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *CreateGroupFeishuDocRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateGroupFeishuDocRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateGroupFeishuDocRequestObjectBindings) SetGraphName(v string) *CreateGroupFeishuDocRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *CreateGroupFeishuDocRequestObjectBindings) SetObjectId(v string) *CreateGroupFeishuDocRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *CreateGroupFeishuDocRequestObjectBindings) SetObjectType(v string) *CreateGroupFeishuDocRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *CreateGroupFeishuDocRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}

type CreateGroupFeishuDocRequestSyncConfig struct {
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
	// The preset mode (can be ignored).
	//
	// example:
	//
	// custom
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreateGroupFeishuDocRequestSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuDocRequestSyncConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuDocRequestSyncConfig) GetCron() *string {
	return s.Cron
}

func (s *CreateGroupFeishuDocRequestSyncConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateGroupFeishuDocRequestSyncConfig) GetPreset() *string {
	return s.Preset
}

func (s *CreateGroupFeishuDocRequestSyncConfig) SetCron(v string) *CreateGroupFeishuDocRequestSyncConfig {
	s.Cron = &v
	return s
}

func (s *CreateGroupFeishuDocRequestSyncConfig) SetEnabled(v bool) *CreateGroupFeishuDocRequestSyncConfig {
	s.Enabled = &v
	return s
}

func (s *CreateGroupFeishuDocRequestSyncConfig) SetPreset(v string) *CreateGroupFeishuDocRequestSyncConfig {
	s.Preset = &v
	return s
}

func (s *CreateGroupFeishuDocRequestSyncConfig) Validate() error {
	return dara.Validate(s)
}
