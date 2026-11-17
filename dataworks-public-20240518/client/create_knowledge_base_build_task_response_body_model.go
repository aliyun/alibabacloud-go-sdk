// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseBuildTaskResponseBody
	GetCode() *string
	SetData(v *CreateKnowledgeBaseBuildTaskResponseBodyData) *CreateKnowledgeBaseBuildTaskResponseBody
	GetData() *CreateKnowledgeBaseBuildTaskResponseBodyData
	SetMessage(v string) *CreateKnowledgeBaseBuildTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKnowledgeBaseBuildTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowledgeBaseBuildTaskResponseBody
	GetSuccess() *bool
}

type CreateKnowledgeBaseBuildTaskResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *CreateKnowledgeBaseBuildTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The secure error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) GetData() *CreateKnowledgeBaseBuildTaskResponseBodyData {
	return s.Data
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) SetCode(v string) *CreateKnowledgeBaseBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) SetData(v *CreateKnowledgeBaseBuildTaskResponseBodyData) *CreateKnowledgeBaseBuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) SetMessage(v string) *CreateKnowledgeBaseBuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) SetRequestId(v string) *CreateKnowledgeBaseBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) SetSuccess(v bool) *CreateKnowledgeBaseBuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskResponseBodyData struct {
	// The language of the content for this task. This parameter is optional and is explicitly selected when the task is created.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1790200000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the task. This parameter is optional.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input scope specified based on the task type.
	Input *CreateKnowledgeBaseBuildTaskResponseBodyDataInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the task belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The specific model ID or AUTO for tasks that require a model. If this parameter is omitted during creation, it is equivalent to AUTO. This parameter is not configured for tasks that do not call a model.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The time when the definition was last modified.
	//
	// example:
	//
	// 1790200300000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// Project space knowledge building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace where the task runs. This is different from the build asset scope specified by Input.ProjectIds.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder compute specifications for running the task.
	RuntimeResource *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The stable ID of the task.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The manual or periodic trigger configuration.
	Trigger *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The task type. Valid values: PROJECT_SPACE_BUILD, SEMANTIC_GRAPH, DOCUMENT_GRAPH, DOCUMENT_SYNC, DOCUMENT_SYNC_AND_GRAPH, and SEMANTIC_MODEL. SEMANTIC_MODEL is used only for YAML imports.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetInput() *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	return s.Input
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetRuntimeResource() *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetTrigger() *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	return s.Trigger
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.ContentLocale = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetCreateTime(v int64) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetDescription(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetInput(v *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.Input = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.LlmModel = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetModifyTime(v int64) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetName(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetRuntimeResource(v *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.RuntimeResource = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetTaskId(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetTrigger(v *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.Trigger = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) SetType(v string) *CreateKnowledgeBaseBuildTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyData) Validate() error {
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

type CreateKnowledgeBaseBuildTaskResponseBodyDataInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The list cannot be empty and must contain unique values.
	Entities []*CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
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

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataInput) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GetEntities() []*CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	return s.Entities
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) SetEntities(v []*CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.Entities = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) SetKnowledgeObjectIds(v []*string) *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) SetProjectIds(v []*int64) *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.ProjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) SetSourceIds(v []*string) *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.SourceIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) SetYamlUploadId(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.YamlUploadId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInput) Validate() error {
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

type CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities struct {
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

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityId(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityType(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityType = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource struct {
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

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetCu(v int32) *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetResourceGroupId(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger struct {
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

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetCron() *string {
	return s.Cron
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetMode() *string {
	return s.Mode
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetCron(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Cron = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetEnabled(v bool) *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Enabled = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetMode(v string) *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Mode = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskResponseBodyDataTrigger) Validate() error {
	return dara.Validate(s)
}
