// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseBuildTaskResponseBody
	GetCode() *string
	SetData(v *GetKnowledgeBaseBuildTaskResponseBodyData) *GetKnowledgeBaseBuildTaskResponseBody
	GetData() *GetKnowledgeBaseBuildTaskResponseBodyData
	SetMessage(v string) *GetKnowledgeBaseBuildTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseBuildTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeBaseBuildTaskResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBaseBuildTaskResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *GetKnowledgeBaseBuildTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The security error message returned when the request fails.
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

func (s GetKnowledgeBaseBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) GetData() *GetKnowledgeBaseBuildTaskResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) SetCode(v string) *GetKnowledgeBaseBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) SetData(v *GetKnowledgeBaseBuildTaskResponseBodyData) *GetKnowledgeBaseBuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) SetMessage(v string) *GetKnowledgeBaseBuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) SetRequestId(v string) *GetKnowledgeBaseBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) SetSuccess(v bool) *GetKnowledgeBaseBuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBaseBuildTaskResponseBodyData struct {
	// The content language of the task. This field is optional and is specified explicitly when the task is created.
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
	// The task description. This parameter is optional.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The typed input scope selected for the task. System-managed upload build tasks do not have fixed inputs, and this field is omitted when reading.
	Input *GetKnowledgeBaseBuildTaskResponseBodyDataInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The ID of the associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The large language model (LLM) configured for tasks that require a model. You can specify a model ID or set this parameter to AUTO. If this parameter is omitted during creation, the default value is AUTO. Do not configure this parameter for tasks that do not call a model.
	//
	// example:
	//
	// AUTO
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The time when the definition was last modified.
	//
	// example:
	//
	// 1790200300000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// Project workspace knowledge build
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace where the task runs. This is different from the build asset scope specified by Input.ProjectIds.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder compute specifications for running the task.
	RuntimeResource *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The stable task ID.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The manual or periodic trigger configuration.
	Trigger *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The task type. Valid values: PROJECT_SPACE_BUILD, SEMANTIC_GRAPH, DOCUMENT_GRAPH, DOCUMENT_SYNC, DOCUMENT_SYNC_AND_GRAPH, and SEMANTIC_MODEL. SEMANTIC_MODEL is used only for YAML import.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetInput() *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	return s.Input
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetLlmModel() *string {
	return s.LlmModel
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetRuntimeResource() *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	return s.RuntimeResource
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetTrigger() *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	return s.Trigger
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetContentLocale(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.ContentLocale = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetCreateTime(v int64) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetDescription(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetInput(v *GetKnowledgeBaseBuildTaskResponseBodyDataInput) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.Input = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetLlmModel(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.LlmModel = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetModifyTime(v int64) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetName(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetProjectId(v int64) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetRuntimeResource(v *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.RuntimeResource = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetTaskId(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetTrigger(v *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.Trigger = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) SetType(v string) *GetKnowledgeBaseBuildTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyData) Validate() error {
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

type GetKnowledgeBaseBuildTaskResponseBodyDataInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values must be non-empty and deduplicated.
	Entities []*GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required, non-empty, and deduplicated when building from document objects in the knowledge base. Each item is a KnowledgeObjectId with Type=DOCUMENT in the knowledge base. Do not mix this with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The IDs of the DataWorks workspaces selected for workspace-scoped builds. The values must be non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the knowledge base. The values are deduplicated. Whether this parameter is required and the allowed source types depend on the task input scenario. For more information, see the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataInput) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataInput) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) GetEntities() []*GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	return s.Entities
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) SetEntities(v []*GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.Entities = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) SetKnowledgeObjectIds(v []*string) *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) SetProjectIds(v []*int64) *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.ProjectIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) SetSourceIds(v []*string) *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.SourceIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) SetYamlUploadId(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.YamlUploadId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInput) Validate() error {
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

type GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this based on the Data Map identity. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityId(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityType(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityType = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataInputEntities) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If this parameter is omitted during creation, the default value is 4. The read operation returns the effective value.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the DataWorks resource group used by the actual Builder nodes. The server validates that it is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetCu(v int32) *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetResourceGroupId(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskResponseBodyDataTrigger struct {
	// The DataWorks scheduling cron expression. This parameter is provided only when the mode is PERIODIC.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Specifies whether periodic automatic triggering is enabled. This parameter is provided only when the mode is PERIODIC and does not indicate the instance status.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger mode. Valid values: MANUAL and PERIODIC. Manual tasks do not create instances automatically.
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetCron() *string {
	return s.Cron
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetMode() *string {
	return s.Mode
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetCron(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Cron = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetEnabled(v bool) *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Enabled = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetMode(v string) *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Mode = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResponseBodyDataTrigger) Validate() error {
	return dara.Validate(s)
}
