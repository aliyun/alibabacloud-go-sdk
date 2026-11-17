// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentLocale(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetContentLocale() *string
	SetDescription(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetDescription() *string
	SetInput(v *UpdateKnowledgeBaseBuildTaskRequestInput) *UpdateKnowledgeBaseBuildTaskRequest
	GetInput() *UpdateKnowledgeBaseBuildTaskRequestInput
	SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetKnowledgeBaseId() *string
	SetLlmModel(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetLlmModel() *string
	SetName(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetName() *string
	SetProjectId(v int64) *UpdateKnowledgeBaseBuildTaskRequest
	GetProjectId() *int64
	SetRuntimeResource(v *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) *UpdateKnowledgeBaseBuildTaskRequest
	GetRuntimeResource() *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource
	SetTaskId(v string) *UpdateKnowledgeBaseBuildTaskRequest
	GetTaskId() *string
	SetTrigger(v *UpdateKnowledgeBaseBuildTaskRequestTrigger) *UpdateKnowledgeBaseBuildTaskRequest
	GetTrigger() *UpdateKnowledgeBaseBuildTaskRequestTrigger
}

type UpdateKnowledgeBaseBuildTaskRequest struct {
	// The content language.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The new description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input configuration that completely replaces the original input.
	Input *UpdateKnowledgeBaseBuildTaskRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The associated knowledge base.
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
	// AUTO
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The new task name.
	//
	// example:
	//
	// Sales Project Knowledge Building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workspace used for future runs.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Updates the runtime resources based on the submitted fields. If ResourceGroupId or Cu is not provided, the original value is retained.
	RuntimeResource *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger configuration that completely replaces the original trigger.
	Trigger *UpdateKnowledgeBaseBuildTaskRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeBaseBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetInput() *UpdateKnowledgeBaseBuildTaskRequestInput {
	return s.Input
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetLlmModel() *string {
	return s.LlmModel
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetRuntimeResource() *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) GetTrigger() *UpdateKnowledgeBaseBuildTaskRequestTrigger {
	return s.Trigger
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetContentLocale(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.ContentLocale = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetDescription(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetInput(v *UpdateKnowledgeBaseBuildTaskRequestInput) *UpdateKnowledgeBaseBuildTaskRequest {
	s.Input = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetLlmModel(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.LlmModel = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetName(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetProjectId(v int64) *UpdateKnowledgeBaseBuildTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetRuntimeResource(v *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) *UpdateKnowledgeBaseBuildTaskRequest {
	s.RuntimeResource = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetTaskId(v string) *UpdateKnowledgeBaseBuildTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) SetTrigger(v *UpdateKnowledgeBaseBuildTaskRequestTrigger) *UpdateKnowledgeBaseBuildTaskRequest {
	s.Trigger = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequest) Validate() error {
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

type UpdateKnowledgeBaseBuildTaskRequestInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The list cannot be empty and must contain unique values.
	Entities []*UpdateKnowledgeBaseBuildTaskRequestInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required when building from document objects in this knowledge base. The list cannot be empty and must contain unique values. Each item is a KnowledgeObjectId with Type=DOCUMENT in this knowledge base. Do not mix this with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The list cannot be empty and must contain unique values.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of registered sources in this knowledge base. The list must contain unique values. Whether this parameter is required and the allowed source types depend on the task input scenario. See the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in this knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskRequestInput) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) GetEntities() []*UpdateKnowledgeBaseBuildTaskRequestInputEntities {
	return s.Entities
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) SetEntities(v []*UpdateKnowledgeBaseBuildTaskRequestInputEntities) *UpdateKnowledgeBaseBuildTaskRequestInput {
	s.Entities = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) SetKnowledgeObjectIds(v []*string) *UpdateKnowledgeBaseBuildTaskRequestInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) SetProjectIds(v []*int64) *UpdateKnowledgeBaseBuildTaskRequestInput {
	s.ProjectIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) SetSourceIds(v []*string) *UpdateKnowledgeBaseBuildTaskRequestInput {
	s.SourceIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) SetYamlUploadId(v string) *UpdateKnowledgeBaseBuildTaskRequestInput {
	s.YamlUploadId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInput) Validate() error {
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

type UpdateKnowledgeBaseBuildTaskRequestInputEntities struct {
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

func (s UpdateKnowledgeBaseBuildTaskRequestInputEntities) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskRequestInputEntities) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInputEntities) SetEntityId(v string) *UpdateKnowledgeBaseBuildTaskRequestInputEntities {
	s.EntityId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInputEntities) SetEntityType(v string) *UpdateKnowledgeBaseBuildTaskRequestInputEntities {
	s.EntityType = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestInputEntities) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseBuildTaskRequestRuntimeResource struct {
	// The compute specifications of the Builder. Valid values: 4 to 16. If this parameter is omitted during creation, the valid value is 4. The read operation returns the valid value.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The DataWorks resource group ID used by the actual Builder node. The server validates that it is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) SetCu(v int32) *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) SetResourceGroupId(v string) *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseBuildTaskRequestTrigger struct {
	// The DataWorks scheduling cron expression. This parameter is provided only when Mode is set to PERIODIC.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Specifies whether periodic automatic triggering is enabled. This parameter is provided only when Mode is set to PERIODIC. This parameter does not indicate the instance status.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger mode. Valid values: MANUAL and PERIODIC. Manual tasks do not create instances automatically.
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) GetCron() *string {
	return s.Cron
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) GetMode() *string {
	return s.Mode
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) SetCron(v string) *UpdateKnowledgeBaseBuildTaskRequestTrigger {
	s.Cron = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) SetEnabled(v bool) *UpdateKnowledgeBaseBuildTaskRequestTrigger {
	s.Enabled = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) SetMode(v string) *UpdateKnowledgeBaseBuildTaskRequestTrigger {
	s.Mode = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskRequestTrigger) Validate() error {
	return dara.Validate(s)
}
