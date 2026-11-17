// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseBuildTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody
	GetCode() *string
	SetData(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) *RetryKnowledgeBaseBuildTaskInstanceResponseBody
	GetData() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData
	SetMessage(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryKnowledgeBaseBuildTaskInstanceResponseBody
	GetSuccess() *bool
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBody struct {
	// The business error code. You can use this code to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned upon a successful request.
	Data *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The secure error message returned when the request fails.
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

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) GetData() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	return s.Data
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) SetCode(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) SetData(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) SetMessage(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) SetRequestId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) SetSuccess(v bool) *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyData struct {
	// The snapshot of the model, content language, and runtime resources for the current execution. This snapshot does not change when the task is updated.
	ConfigSnapshot *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current business stage. This field is returned when reliable records are available. For common build stages and semantic model import stages, refer to the domain model. This field does not replace the overall status.
	//
	// example:
	//
	// INITIALIZE
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The error code and secure error message returned when the request fails.
	Error *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The time when the execution finished. This field is optional.
	//
	// example:
	//
	// 1790200060000
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The confirmed ModelId and VersionNo for model imports only. For their structure, see GetKnowledgeBaseBuildTaskInstance.
	GeneratedModels []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels `json:"GeneratedModels,omitempty" xml:"GeneratedModels,omitempty" type:"Repeated"`
	// The non-sensitive input snapshot after selection and parsing for the current execution. This snapshot does not change when the task is updated.
	InputSnapshot *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot `json:"InputSnapshot,omitempty" xml:"InputSnapshot,omitempty" type:"Struct"`
	// The ID of the associated knowledge base, which facilitates scope validation and queries.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The time when the execution started. This field is optional.
	//
	// example:
	//
	// 1790200001000
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The overall status of the current execution.
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the source task.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the current execution.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The trigger type of the current execution, which can be manual, periodic, or a supported event.
	//
	// example:
	//
	// MANUAL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetConfigSnapshot() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetError() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	return s.Error
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetGeneratedModels() []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	return s.GeneratedModels
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetInputSnapshot() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	return s.InputSnapshot
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTriggerType() *string {
	return s.TriggerType
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetConfigSnapshot(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.ConfigSnapshot = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCreateTime(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCurrentStage(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CurrentStage = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetError(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Error = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetFinishedTime(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.FinishedTime = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetGeneratedModels(v []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.GeneratedModels = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetInputSnapshot(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.InputSnapshot = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetKnowledgeBaseId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStartedTime(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStatus(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskInstanceId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTriggerType(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyData) Validate() error {
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

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot struct {
	// The content language used for the current execution.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The specific model ID actually used and frozen for the current execution. The value AUTO is not returned. This field can be empty for tasks that do not call models, such as pure synchronization tasks.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The DataWorks workspace ID used for the current execution. This ID is frozen from the task definition when the instance is created.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder CUs actually used for the current execution.
	RuntimeResource *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetLlmModel() *string {
	return s.LlmModel
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetRuntimeResource() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	return s.RuntimeResource
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetContentLocale(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ContentLocale = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetLlmModel(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.LlmModel = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetProjectId(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ProjectId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetRuntimeResource(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.RuntimeResource = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource struct {
	// The Builder compute capacity. Valid values: 4 to 16. If this parameter is omitted during creation, the default value is 4. The read operation returns the effective value.
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

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetCu(v int32) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.Cu = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetResourceGroupId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError struct {
	// The business error code. You can use this code to determine the error type when a request fails.
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

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetCode(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetMessage(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels struct {
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

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetModelId() *string {
	return s.ModelId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetModelId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.ModelId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetVersionNo(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.VersionNo = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) Validate() error {
	return dara.Validate(s)
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot struct {
	// The number of external document sources selected for the current execution. This field is not returned when it is not applicable.
	//
	// example:
	//
	// 1
	DocumentSourceCount *int64 `json:"DocumentSourceCount,omitempty" xml:"DocumentSourceCount,omitempty"`
	// The actual Data Map entities resolved from the project, database, or schema scope for the current execution. This field is not returned when it is not applicable.
	ResolvedEntities []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities `json:"ResolvedEntities,omitempty" xml:"ResolvedEntities,omitempty" type:"Repeated"`
	// The total number of resolved entities. An accurate count is returned even when the array is not returned due to its size.
	//
	// example:
	//
	// 42
	ResolvedEntityCount *int64 `json:"ResolvedEntityCount,omitempty" xml:"ResolvedEntityCount,omitempty"`
	// The snapshot of the selection conditions in the task definition when the instance is created. This snapshot does not change when the task is updated.
	SelectedInput *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput `json:"SelectedInput,omitempty" xml:"SelectedInput,omitempty" type:"Struct"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetDocumentSourceCount() *int64 {
	return s.DocumentSourceCount
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntities() []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	return s.ResolvedEntities
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntityCount() *int64 {
	return s.ResolvedEntityCount
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetSelectedInput() *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	return s.SelectedInput
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetDocumentSourceCount(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.DocumentSourceCount = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntities(v []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntities = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntityCount(v int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntityCount = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetSelectedInput(v *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.SelectedInput = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) Validate() error {
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

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as-is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this field based on the Data Map identity. The presence of an enumeration value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityType(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityType = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) Validate() error {
	return dara.Validate(s)
}

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values are non-empty and deduplicated.
	Entities []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required, non-empty, and deduplicated when building from document objects in the knowledge base. Each item is a KnowledgeObjectId with Type=DOCUMENT in the knowledge base. Do not mix this field with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The values are non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the knowledge base. The values are deduplicated. Whether this field is required and the allowed source types are determined by the task input scenario. Refer to the following table for details.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetEntities() []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	return s.Entities
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetEntities(v []*RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.Entities = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetKnowledgeObjectIds(v []*string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetProjectIds(v []*int64) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.ProjectIds = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetSourceIds(v []*string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.SourceIds = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetYamlUploadId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.YamlUploadId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) Validate() error {
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

type RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as-is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this field based on the Data Map identity. The presence of an enumeration value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityId(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityType(v string) *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityType = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) Validate() error {
	return dara.Validate(s)
}
