// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseRequest
	GetDirectoryId() *string
	SetKbName(v string) *CreatePersonalAlidingKnowledgeBaseRequest
	GetKbName() *string
	SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseRequest
	GetKbUrl() *string
	SetObjectBindings(v []*CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) *CreatePersonalAlidingKnowledgeBaseRequest
	GetObjectBindings() []*CreatePersonalAlidingKnowledgeBaseRequestObjectBindings
	SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseRequest
	GetOperatingObjectName() *string
	SetSyncConfig(v *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) *CreatePersonalAlidingKnowledgeBaseRequest
	GetSyncConfig() *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig
	SetTenantId(v string) *CreatePersonalAlidingKnowledgeBaseRequest
	GetTenantId() *string
}

type CreatePersonalAlidingKnowledgeBaseRequest struct {
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
	ObjectBindings []*CreatePersonalAlidingKnowledgeBaseRequestObjectBindings `json:"objectBindings,omitempty" xml:"objectBindings,omitempty" type:"Repeated"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The synchronization settings.
	SyncConfig *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig `json:"syncConfig,omitempty" xml:"syncConfig,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// example:
	//
	// PiPklI1iSRTm6VFFqlY9VzbgiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetKbName() *string {
	return s.KbName
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetKbUrl() *string {
	return s.KbUrl
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetObjectBindings() []*CreatePersonalAlidingKnowledgeBaseRequestObjectBindings {
	return s.ObjectBindings
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetSyncConfig() *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig {
	return s.SyncConfig
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetDirectoryId(v string) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetKbName(v string) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.KbName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetKbUrl(v string) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.KbUrl = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetObjectBindings(v []*CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.ObjectBindings = v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetOperatingObjectName(v string) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetSyncConfig(v *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.SyncConfig = v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) SetTenantId(v string) *CreatePersonalAlidingKnowledgeBaseRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequest) Validate() error {
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

type CreatePersonalAlidingKnowledgeBaseRequestObjectBindings struct {
	// The ID of the recommended item, which can be a **feedId*	- or a micro-application ID.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The advanced field type.
	//
	// example:
	//
	// table
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) SetObjectId(v string) *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings {
	s.ObjectId = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) SetObjectType(v string) *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings {
	s.ObjectType = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestObjectBindings) Validate() error {
	return dara.Validate(s)
}

type CreatePersonalAlidingKnowledgeBaseRequestSyncConfig struct {
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// string_value
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Specifies whether to enable synchronization.
	//
	// example:
	//
	// False
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) GetCron() *string {
	return s.Cron
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) SetCron(v string) *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig {
	s.Cron = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) SetEnabled(v bool) *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig {
	s.Enabled = &v
	return s
}

func (s *CreatePersonalAlidingKnowledgeBaseRequestSyncConfig) Validate() error {
	return dara.Validate(s)
}
