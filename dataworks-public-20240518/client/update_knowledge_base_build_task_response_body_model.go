// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateKnowledgeBaseBuildTaskResponseBody
	GetCode() *string
	SetData(v *UpdateKnowledgeBaseBuildTaskResponseBodyData) *UpdateKnowledgeBaseBuildTaskResponseBody
	GetData() *UpdateKnowledgeBaseBuildTaskResponseBodyData
	SetMessage(v string) *UpdateKnowledgeBaseBuildTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateKnowledgeBaseBuildTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKnowledgeBaseBuildTaskResponseBody
	GetSuccess() *bool
}

type UpdateKnowledgeBaseBuildTaskResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *UpdateKnowledgeBaseBuildTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The secure error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) GetData() *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	return s.Data
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) SetCode(v string) *UpdateKnowledgeBaseBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) SetData(v *UpdateKnowledgeBaseBuildTaskResponseBodyData) *UpdateKnowledgeBaseBuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) SetMessage(v string) *UpdateKnowledgeBaseBuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) SetRequestId(v string) *UpdateKnowledgeBaseBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) SetSuccess(v bool) *UpdateKnowledgeBaseBuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeBaseBuildTaskResponseBodyData struct {
	// The content language of this task. This parameter is optional and is explicitly selected when the task is created.
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
	// The input scope specified based on the task type.
	Input *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The associated knowledge base.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The specific model ID or AUTO configured for tasks that require a model. If this parameter is omitted during creation, it is equivalent to AUTO. This parameter is not configured for tasks that do not call a model.
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
	// Sales Project Knowledge Building
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID where the task runs. This is different from the build asset scope specified by Input.ProjectIds.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The task runtime resource group and Builder compute specifications.
	RuntimeResource *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The stable task ID.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The manual or periodic trigger configuration.
	Trigger *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The task type. Valid values: PROJECT_SPACE_BUILD, SEMANTIC_GRAPH, DOCUMENT_GRAPH, DOCUMENT_SYNC, DOCUMENT_SYNC_AND_GRAPH, and SEMANTIC_MODEL. SEMANTIC_MODEL is used only for YAML import.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetInput() *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	return s.Input
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetLlmModel() *string {
	return s.LlmModel
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetRuntimeResource() *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetTrigger() *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	return s.Trigger
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetContentLocale(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.ContentLocale = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetCreateTime(v int64) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetDescription(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetInput(v *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.Input = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetKnowledgeBaseId(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetLlmModel(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.LlmModel = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetModifyTime(v int64) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetName(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetProjectId(v int64) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetRuntimeResource(v *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.RuntimeResource = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetTaskId(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetTrigger(v *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.Trigger = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) SetType(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyData) Validate() error {
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

type UpdateKnowledgeBaseBuildTaskResponseBodyDataInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The list cannot be empty and must contain unique values.
	Entities []*UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
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

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GetEntities() []*UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	return s.Entities
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) SetEntities(v []*UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.Entities = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) SetKnowledgeObjectIds(v []*string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) SetProjectIds(v []*int64) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.ProjectIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) SetSourceIds(v []*string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.SourceIds = v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) SetYamlUploadId(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput {
	s.YamlUploadId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInput) Validate() error {
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

type UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities struct {
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

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityId(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) SetEntityType(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities {
	s.EntityType = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataInputEntities) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource struct {
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

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetCu(v int32) *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) SetResourceGroupId(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger struct {
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

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetCron() *string {
	return s.Cron
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) GetMode() *string {
	return s.Mode
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetCron(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Cron = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetEnabled(v bool) *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Enabled = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) SetMode(v string) *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger {
	s.Mode = &v
	return s
}

func (s *UpdateKnowledgeBaseBuildTaskResponseBodyDataTrigger) Validate() error {
	return dara.Validate(s)
}
