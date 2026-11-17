// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetClientToken() *string
	SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetContentLocale() *string
	SetDescription(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetDescription() *string
	SetInput(v *CreateKnowledgeBaseBuildTaskRequestInput) *CreateKnowledgeBaseBuildTaskRequest
	GetInput() *CreateKnowledgeBaseBuildTaskRequestInput
	SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetKnowledgeBaseId() *string
	SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetLlmModel() *string
	SetName(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetName() *string
	SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskRequest
	GetProjectId() *int64
	SetRuntimeResource(v *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) *CreateKnowledgeBaseBuildTaskRequest
	GetRuntimeResource() *CreateKnowledgeBaseBuildTaskRequestRuntimeResource
	SetTrigger(v *CreateKnowledgeBaseBuildTaskRequestTrigger) *CreateKnowledgeBaseBuildTaskRequest
	GetTrigger() *CreateKnowledgeBaseBuildTaskRequestTrigger
	SetType(v string) *CreateKnowledgeBaseBuildTaskRequest
	GetType() *string
}

type CreateKnowledgeBaseBuildTaskRequest struct {
	// The idempotency token for the creation request.
	//
	// example:
	//
	// task-create-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the content.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The description of the task. The description can be up to 1,024 characters in length. If this parameter is omitted, no description is set.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies the input scope by selecting projects, entities, sources, or files based on the type.
	//
	// This parameter is required.
	Input *CreateKnowledgeBaseBuildTaskRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The specific model ID or AUTO.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// Project space knowledge building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace where the task runs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder compute units (CUs).
	//
	// This parameter is required.
	RuntimeResource *CreateKnowledgeBaseBuildTaskRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The manual or periodic trigger configuration.
	Trigger *CreateKnowledgeBaseBuildTaskRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The available build types.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetInput() *CreateKnowledgeBaseBuildTaskRequestInput {
	return s.Input
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetRuntimeResource() *CreateKnowledgeBaseBuildTaskRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetTrigger() *CreateKnowledgeBaseBuildTaskRequestTrigger {
	return s.Trigger
}

func (s *CreateKnowledgeBaseBuildTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetClientToken(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.ContentLocale = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetDescription(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetInput(v *CreateKnowledgeBaseBuildTaskRequestInput) *CreateKnowledgeBaseBuildTaskRequest {
	s.Input = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.LlmModel = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetName(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetRuntimeResource(v *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) *CreateKnowledgeBaseBuildTaskRequest {
	s.RuntimeResource = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetTrigger(v *CreateKnowledgeBaseBuildTaskRequestTrigger) *CreateKnowledgeBaseBuildTaskRequest {
	s.Trigger = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) SetType(v string) *CreateKnowledgeBaseBuildTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskRequestInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The list cannot be empty and must contain unique values.
	Entities []*CreateKnowledgeBaseBuildTaskRequestInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required when building from document objects in the current knowledge base. The list cannot be empty and must contain unique values. Each item is the KnowledgeObjectId of a DOCUMENT type object in the current knowledge base. This parameter cannot be mixed with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The IDs of the DataWorks workspaces selected for workspace-scoped builds. The list cannot be empty and must contain unique values.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the current knowledge base. The list must contain unique values. Whether this parameter is required and the allowed source types depend on the task input scenario. For more information, see the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the current knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) GetEntities() []*CreateKnowledgeBaseBuildTaskRequestInputEntities {
	return s.Entities
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) SetEntities(v []*CreateKnowledgeBaseBuildTaskRequestInputEntities) *CreateKnowledgeBaseBuildTaskRequestInput {
	s.Entities = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) SetKnowledgeObjectIds(v []*string) *CreateKnowledgeBaseBuildTaskRequestInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) SetProjectIds(v []*int64) *CreateKnowledgeBaseBuildTaskRequestInput {
	s.ProjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) SetSourceIds(v []*string) *CreateKnowledgeBaseBuildTaskRequestInput {
	s.SourceIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) SetYamlUploadId(v string) *CreateKnowledgeBaseBuildTaskRequestInput {
	s.YamlUploadId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInput) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskRequestInputEntities struct {
	// The metadata entity ID returned by Data Map. The client must pass this value as is and must not concatenate database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this value based on the Data Map identity. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskRequestInputEntities) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskRequestInputEntities) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskRequestInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeBaseBuildTaskRequestInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateKnowledgeBaseBuildTaskRequestInputEntities) SetEntityId(v string) *CreateKnowledgeBaseBuildTaskRequestInputEntities {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInputEntities) SetEntityType(v string) *CreateKnowledgeBaseBuildTaskRequestInputEntities {
	s.EntityType = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestInputEntities) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskRequestRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If this parameter is omitted during creation, the default value is 4. The actual valid value is returned in read operations.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the DataWorks resource group used by the actual Builder nodes. The server validates that the resource group is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) SetCu(v int32) *CreateKnowledgeBaseBuildTaskRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) SetResourceGroupId(v string) *CreateKnowledgeBaseBuildTaskRequestRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskRequestTrigger struct {
	// The DataWorks scheduling cron expression. This parameter is provided only when the mode is set to PERIODIC.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Specifies whether to enable or disable automatic periodic triggering. This parameter is provided only when the mode is set to PERIODIC and does not indicate the instance status.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger mode. Valid values: MANUAL and PERIODIC. Manual tasks do not create instances automatically.
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskRequestTrigger) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) GetCron() *string {
	return s.Cron
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) GetMode() *string {
	return s.Mode
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) SetCron(v string) *CreateKnowledgeBaseBuildTaskRequestTrigger {
	s.Cron = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) SetEnabled(v bool) *CreateKnowledgeBaseBuildTaskRequestTrigger {
	s.Enabled = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) SetMode(v string) *CreateKnowledgeBaseBuildTaskRequestTrigger {
	s.Mode = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskRequestTrigger) Validate() error {
	return dara.Validate(s)
}
