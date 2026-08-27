// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalFeishuDocRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalFeishuDocRequest
	GetDirectoryId() *string
	SetDocUrl(v string) *CreatePersonalFeishuDocRequest
	GetDocUrl() *string
	SetName(v string) *CreatePersonalFeishuDocRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalFeishuDocRequest
	GetNotes() *string
	SetObjectBindings(v []*CreatePersonalFeishuDocRequestObjectBindings) *CreatePersonalFeishuDocRequest
	GetObjectBindings() []*CreatePersonalFeishuDocRequestObjectBindings
	SetOperatingObjectName(v string) *CreatePersonalFeishuDocRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreatePersonalFeishuDocRequest
	GetSourceTags() *string
	SetSyncConfig(v *CreatePersonalFeishuDocRequestSyncConfig) *CreatePersonalFeishuDocRequest
	GetSyncConfig() *CreatePersonalFeishuDocRequestSyncConfig
	SetTenantId(v string) *CreatePersonalFeishuDocRequest
	GetTenantId() *string
}

type CreatePersonalFeishuDocRequest struct {
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
	ObjectBindings []*CreatePersonalFeishuDocRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
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
	SyncConfig *CreatePersonalFeishuDocRequestSyncConfig `json:"syncConfig,omitempty" xml:"syncConfig,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalFeishuDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalFeishuDocRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuDocRequest) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreatePersonalFeishuDocRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuDocRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalFeishuDocRequest) GetObjectBindings() []*CreatePersonalFeishuDocRequestObjectBindings {
	return s.ObjectBindings
}

func (s *CreatePersonalFeishuDocRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalFeishuDocRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalFeishuDocRequest) GetSyncConfig() *CreatePersonalFeishuDocRequestSyncConfig {
	return s.SyncConfig
}

func (s *CreatePersonalFeishuDocRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalFeishuDocRequest) SetDescription(v string) *CreatePersonalFeishuDocRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetDirectoryId(v string) *CreatePersonalFeishuDocRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetDocUrl(v string) *CreatePersonalFeishuDocRequest {
	s.DocUrl = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetName(v string) *CreatePersonalFeishuDocRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetNotes(v string) *CreatePersonalFeishuDocRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetObjectBindings(v []*CreatePersonalFeishuDocRequestObjectBindings) *CreatePersonalFeishuDocRequest {
	s.ObjectBindings = v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetOperatingObjectName(v string) *CreatePersonalFeishuDocRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetSourceTags(v string) *CreatePersonalFeishuDocRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetSyncConfig(v *CreatePersonalFeishuDocRequestSyncConfig) *CreatePersonalFeishuDocRequest {
	s.SyncConfig = v
	return s
}

func (s *CreatePersonalFeishuDocRequest) SetTenantId(v string) *CreatePersonalFeishuDocRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalFeishuDocRequest) Validate() error {
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

type CreatePersonalFeishuDocRequestObjectBindings struct {
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

func (s CreatePersonalFeishuDocRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) GetGraphName() *string {
	return s.GraphName
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) SetGraphName(v string) *CreatePersonalFeishuDocRequestObjectBindings {
	s.GraphName = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) SetObjectId(v string) *CreatePersonalFeishuDocRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) SetObjectType(v string) *CreatePersonalFeishuDocRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}

type CreatePersonalFeishuDocRequestSyncConfig struct {
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

func (s CreatePersonalFeishuDocRequestSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuDocRequestSyncConfig) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) GetCron() *string {
	return s.Cron
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) GetPreset() *string {
	return s.Preset
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) SetCron(v string) *CreatePersonalFeishuDocRequestSyncConfig {
	s.Cron = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) SetEnabled(v bool) *CreatePersonalFeishuDocRequestSyncConfig {
	s.Enabled = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) SetPreset(v string) *CreatePersonalFeishuDocRequestSyncConfig {
	s.Preset = &v
	return s
}

func (s *CreatePersonalFeishuDocRequestSyncConfig) Validate() error {
	return dara.Validate(s)
}
