// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody
	GetCode() *string
	SetData(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) *CreateKnowledgeBaseBuildTaskInstanceResponseBody
	GetData() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData
	SetMessage(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowledgeBaseBuildTaskInstanceResponseBody
	GetSuccess() *bool
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) GetData() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	return s.Data
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) SetCode(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) SetData(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) SetMessage(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) SetRequestId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) SetSuccess(v bool) *CreateKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyData struct {
	// The snapshot of the model, content language, and runtime resources for this execution. This snapshot does not change when the task is updated.
	ConfigSnapshot *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current business stage. This parameter is returned when reliable records are available. For common build stages and semantic model import stages, refer to the domain model. This parameter does not replace the overall status.
	//
	// example:
	//
	// INITIALIZE
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The error code and secure error message returned when the request fails.
	Error *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The time when the execution finished. This parameter is optional.
	//
	// example:
	//
	// 1790200060000
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The confirmed ModelId and VersionNo for model imports only. For their structure, see GetKnowledgeBaseBuildTaskInstance.
	GeneratedModels []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels `json:"GeneratedModels,omitempty" xml:"GeneratedModels,omitempty" type:"Repeated"`
	// The non-sensitive input snapshot after selection and parsing for this execution. This snapshot does not change when the task is updated.
	InputSnapshot *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot `json:"InputSnapshot,omitempty" xml:"InputSnapshot,omitempty" type:"Struct"`
	// The knowledge base ID, which is used for scope validation and queries.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The time when the execution started. This parameter is optional.
	//
	// example:
	//
	// 1790200001000
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The overall status of this execution.
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The source task ID.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of this execution.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The trigger type of this execution, such as manual, periodic, or supported events.
	//
	// example:
	//
	// MANUAL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetConfigSnapshot() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetError() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	return s.Error
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetGeneratedModels() []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	return s.GeneratedModels
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetInputSnapshot() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	return s.InputSnapshot
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetConfigSnapshot(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.ConfigSnapshot = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCreateTime(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCurrentStage(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CurrentStage = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetError(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Error = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetFinishedTime(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.FinishedTime = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetGeneratedModels(v []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.GeneratedModels = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetInputSnapshot(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.InputSnapshot = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStartedTime(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStatus(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskInstanceId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTriggerType(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyData) Validate() error {
	if s.ConfigSnapshot != nil {
		if err := s.ConfigSnapshot.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	if s.GeneratedModels != nil {
		for _, item := range s.GeneratedModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InputSnapshot != nil {
		if err := s.InputSnapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot struct {
	// The content language used for this execution.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The specific model ID actually used and frozen for this execution. AUTO is not returned. This parameter can be empty for tasks that do not call models, such as pure synchronization tasks.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The DataWorks workspace ID used for this execution. This ID is frozen from the task definition when the instance is created.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder CUs actually used for this execution.
	RuntimeResource *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetLlmModel() *string {
	return s.LlmModel
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetRuntimeResource() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetContentLocale(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ContentLocale = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetLlmModel(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.LlmModel = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetProjectId(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ProjectId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetRuntimeResource(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.RuntimeResource = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If this parameter is omitted during creation, the effective value is 4. The read operation returns the effective value.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The DataWorks resource group ID used by the actual Builder nodes. The server validates that this resource group is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetCu(v int32) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetResourceGroupId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The secure error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetCode(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetMessage(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels struct {
	// The model ID generated by the semantic layer. You can use this ID as a KnowledgeObjectId to read the corresponding semantic model.
	//
	// example:
	//
	// 01MODELULID
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The current version of the SEMANTIC_MODEL being read. Pass this value when updating content. The version can be edited in place.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetModelId() *string {
	return s.ModelId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetModelId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.ModelId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetVersionNo(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.VersionNo = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot struct {
	// The number of external document sources selected for this execution. This parameter is not returned when it is not applicable.
	//
	// example:
	//
	// 1
	DocumentSourceCount *int64 `json:"DocumentSourceCount,omitempty" xml:"DocumentSourceCount,omitempty"`
	// The actual Data Map entities resolved from the project, database, or schema scope for this execution. This parameter is not returned when it is not applicable.
	ResolvedEntities []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities `json:"ResolvedEntities,omitempty" xml:"ResolvedEntities,omitempty" type:"Repeated"`
	// The total number of resolved entities. An accurate count is returned even if the array is not returned due to its size.
	//
	// example:
	//
	// 42
	ResolvedEntityCount *int64 `json:"ResolvedEntityCount,omitempty" xml:"ResolvedEntityCount,omitempty"`
	// The snapshot of the selection conditions in the task definition when the instance is created. This snapshot does not change when the task is updated.
	SelectedInput *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput `json:"SelectedInput,omitempty" xml:"SelectedInput,omitempty" type:"Struct"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetDocumentSourceCount() *int64 {
	return s.DocumentSourceCount
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntities() []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	return s.ResolvedEntities
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntityCount() *int64 {
	return s.ResolvedEntityCount
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetSelectedInput() *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	return s.SelectedInput
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetDocumentSourceCount(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.DocumentSourceCount = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntities(v []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntities = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntityCount(v int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntityCount = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetSelectedInput(v *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.SelectedInput = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) Validate() error {
	if s.ResolvedEntities != nil {
		for _, item := range s.ResolvedEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SelectedInput != nil {
		if err := s.SelectedInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this parameter based on the Data Map identity. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityType(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityType = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values are non-empty and deduplicated.
	Entities []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The knowledge object IDs required when building from document objects in this knowledge base. The values must be non-empty and deduplicated. Each item is a KnowledgeObjectId with Type=DOCUMENT in this knowledge base. Do not mix these IDs with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The values are non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of registered sources in this knowledge base. The values are deduplicated. Whether this parameter is required and the allowed source types depend on the task input scenario. Refer to the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in this knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetEntities() []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	return s.Entities
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetEntities(v []*CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.Entities = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetKnowledgeObjectIds(v []*string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetProjectIds(v []*int64) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.ProjectIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetSourceIds(v []*string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.SourceIds = v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetYamlUploadId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.YamlUploadId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) Validate() error {
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

type CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this parameter based on the Data Map identity. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityId(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityType(v string) *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityType = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) Validate() error {
	return dara.Validate(s)
}
