// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseBuildTasksResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBaseBuildTasksResponseBodyData) *ListKnowledgeBaseBuildTasksResponseBody
	GetData() *ListKnowledgeBaseBuildTasksResponseBodyData
	SetMaxResults(v int32) *ListKnowledgeBaseBuildTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgeBaseBuildTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBaseBuildTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBaseBuildTasksResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBaseBuildTasksResponseBody struct {
	// The business error code. You can use this code to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeBaseBuildTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetData() *ListKnowledgeBaseBuildTasksResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetCode(v string) *ListKnowledgeBaseBuildTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetData(v *ListKnowledgeBaseBuildTasksResponseBodyData) *ListKnowledgeBaseBuildTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetMaxResults(v int32) *ListKnowledgeBaseBuildTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetMessage(v string) *ListKnowledgeBaseBuildTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetRequestId(v string) *ListKnowledgeBaseBuildTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) SetSuccess(v bool) *ListKnowledgeBaseBuildTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseBuildTasksResponseBodyData struct {
	// The cursor for the next page. If there is no next page, the value is null.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The knowledge objects that you have permission to read on the current page.
	Tasks []*ListKnowledgeBaseBuildTasksResponseBodyDataTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyData) GetTasks() []*ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	return s.Tasks
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyData) SetNextToken(v string) *ListKnowledgeBaseBuildTasksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyData) SetTasks(v []*ListKnowledgeBaseBuildTasksResponseBodyDataTasks) *ListKnowledgeBaseBuildTasksResponseBodyData {
	s.Tasks = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyData) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseBuildTasksResponseBodyDataTasks struct {
	// The content language of the task. This field is optional and is explicitly selected when the task is created.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the task. This field is optional.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The typed input scope selected for the task. System-managed upload build tasks do not have fixed inputs, and this field is omitted when reading.
	Input *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The knowledge base to which the task belongs.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The model configured for tasks that require a model. You can specify a model ID or set this parameter to AUTO. If this parameter is omitted during creation, the default value AUTO is used. Do not configure this parameter for tasks that do not call a model.
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
	// The name of the task.
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
	RuntimeResource *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The stable task ID.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The manual or periodic trigger configuration.
	Trigger *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The task type. Valid values: PROJECT_SPACE_BUILD, SEMANTIC_GRAPH, DOCUMENT_GRAPH, DOCUMENT_SYNC, DOCUMENT_SYNC_AND_GRAPH, and SEMANTIC_MODEL. The SEMANTIC_MODEL value is used only for YAML import.
	//
	// example:
	//
	// PROJECT_SPACE_BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasks) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetInput() *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	return s.Input
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetLlmModel() *string {
	return s.LlmModel
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetRuntimeResource() *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetTrigger() *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger {
	return s.Trigger
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetContentLocale(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.ContentLocale = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetCreateTime(v int64) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetDescription(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.Description = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetInput(v *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.Input = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetLlmModel(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.LlmModel = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetModifyTime(v int64) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.ModifyTime = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetName(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.Name = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetProjectId(v int64) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.ProjectId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetRuntimeResource(v *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.RuntimeResource = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetTaskId(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.TaskId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetTrigger(v *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.Trigger = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) SetType(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasks {
	s.Type = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasks) Validate() error {
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

type ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values must be non-empty and deduplicated.
	Entities []*ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The IDs of the knowledge objects. This field is required, non-empty, and deduplicated when building from document objects in the knowledge base. Each item is a KnowledgeObjectId with Type=DOCUMENT in the knowledge base. Do not mix this field with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The IDs of the DataWorks workspaces selected for workspace-scoped builds. The values must be non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the knowledge base. The values are deduplicated. Whether this field is required and the allowed source types depend on the task input scenario. For more information, see the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GetEntities() []*ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities {
	return s.Entities
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) SetEntities(v []*ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	s.Entities = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) SetKnowledgeObjectIds(v []*string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) SetProjectIds(v []*int64) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	s.ProjectIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) SetSourceIds(v []*string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	s.SourceIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) SetYamlUploadId(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput {
	s.YamlUploadId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInput) Validate() error {
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

type ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server verifies this against Data Map identities. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) SetEntityId(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities {
	s.EntityId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) SetEntityType(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities {
	s.EntityType = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksInputEntities) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If this parameter is omitted during creation, the default value 4 is used. The actual value is returned when reading.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the DataWorks resource group used by the actual Builder nodes. The server verifies that the resource group is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) SetCu(v int32) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) SetResourceGroupId(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger struct {
	// The DataWorks scheduling cron expression. This field is provided only when the trigger mode is PERIODIC.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Specifies whether periodic automatic triggering is enabled. This field is provided only when the trigger mode is PERIODIC and does not indicate the instance status.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger mode. Valid values: MANUAL and PERIODIC. Manual tasks do not create instances automatically.
	//
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) GetMode() *string {
	return s.Mode
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) SetCron(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger {
	s.Cron = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) SetEnabled(v bool) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger {
	s.Enabled = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) SetMode(v string) *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger {
	s.Mode = &v
	return s
}

func (s *ListKnowledgeBaseBuildTasksResponseBodyDataTasksTrigger) Validate() error {
	return dara.Validate(s)
}
