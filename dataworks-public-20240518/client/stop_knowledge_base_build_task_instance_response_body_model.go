// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopKnowledgeBaseBuildTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody
	GetCode() *string
	SetData(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) *StopKnowledgeBaseBuildTaskInstanceResponseBody
	GetData() *StopKnowledgeBaseBuildTaskInstanceResponseBodyData
	SetMessage(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopKnowledgeBaseBuildTaskInstanceResponseBody
	GetSuccess() *bool
}

type StopKnowledgeBaseBuildTaskInstanceResponseBody struct {
	// The business error code, used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *StopKnowledgeBaseBuildTaskInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s StopKnowledgeBaseBuildTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) GetData() *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	return s.Data
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) SetCode(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) SetData(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) SetMessage(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) SetRequestId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) SetSuccess(v bool) *StopKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyData struct {
	// The snapshot of the model, content language, and runtime resources for the current run. This snapshot does not change when the task is updated.
	ConfigSnapshot *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current business stage. Returned when reliable records are available. For common build stages and semantic model import stages, refer to the domain model. This does not replace the overall status.
	//
	// example:
	//
	// INITIALIZE
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The error code and secure error message returned when the request fails.
	Error *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The time when execution ended. This field is optional.
	//
	// example:
	//
	// 1790200060000
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The confirmed ModelId and VersionNo for model imports only. For their structure, see GetKnowledgeBaseBuildTaskInstance.
	GeneratedModels []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels `json:"GeneratedModels,omitempty" xml:"GeneratedModels,omitempty" type:"Repeated"`
	// The non-sensitive input snapshot after selection and resolution for the current run. This snapshot does not change when the task is updated.
	InputSnapshot *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot `json:"InputSnapshot,omitempty" xml:"InputSnapshot,omitempty" type:"Struct"`
	// The ID of the associated knowledge base, used for scope validation and queries.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The time when execution started. This field is optional.
	//
	// example:
	//
	// 1790200001000
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The overall status of the current execution.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The source task ID.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The trigger type of the current run, which can be manual, periodic, or a supported event.
	//
	// example:
	//
	// MANUAL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetConfigSnapshot() *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetError() *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	return s.Error
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetGeneratedModels() []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	return s.GeneratedModels
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetInputSnapshot() *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	return s.InputSnapshot
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTriggerType() *string {
	return s.TriggerType
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetConfigSnapshot(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.ConfigSnapshot = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCreateTime(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCurrentStage(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CurrentStage = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetError(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Error = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetFinishedTime(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.FinishedTime = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetGeneratedModels(v []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.GeneratedModels = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetInputSnapshot(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.InputSnapshot = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetKnowledgeBaseId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStartedTime(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStatus(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskInstanceId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTriggerType(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyData) Validate() error {
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

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot struct {
	// The content language used for the current run.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The specific model ID actually used and frozen for the current run. AUTO is not returned. This value can be empty for tasks that do not call models, such as pure synchronization tasks.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The DataWorks workspace ID used for the current run. This value is frozen from the task definition when the instance is created.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder CUs actually used for the current run.
	RuntimeResource *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetLlmModel() *string {
	return s.LlmModel
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetRuntimeResource() *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	return s.RuntimeResource
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetContentLocale(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ContentLocale = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetLlmModel(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.LlmModel = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetProjectId(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ProjectId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetRuntimeResource(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.RuntimeResource = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If omitted during creation, the effective value is 4. The effective value is returned when queried.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The DataWorks resource group ID used by the actual Builder nodes. The server validates that it is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetCu(v int32) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.Cu = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetResourceGroupId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError struct {
	// The business error code, used to determine the error type when the request fails.
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

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetCode(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetMessage(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels struct {
	// The model ID generated by the semantic layer, which can be used as a KnowledgeObjectId to read the corresponding semantic model.
	//
	// example:
	//
	// 01MODELULID
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The currently read version of the SEMANTIC_MODEL. Pass this value when updating content. The version can be edited in place.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetModelId() *string {
	return s.ModelId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetModelId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.ModelId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetVersionNo(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.VersionNo = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) Validate() error {
	return dara.Validate(s)
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot struct {
	// The number of external document sources selected for the current run. Not returned when not applicable.
	//
	// example:
	//
	// 1
	DocumentSourceCount *int64 `json:"DocumentSourceCount,omitempty" xml:"DocumentSourceCount,omitempty"`
	// The actual Data Map entities resolved from the project, database, or schema scope for the current run. Not returned when not applicable.
	ResolvedEntities []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities `json:"ResolvedEntities,omitempty" xml:"ResolvedEntities,omitempty" type:"Repeated"`
	// The total number of resolved entities. An accurate count is returned even if the array is not returned due to its size.
	//
	// example:
	//
	// 42
	ResolvedEntityCount *int64 `json:"ResolvedEntityCount,omitempty" xml:"ResolvedEntityCount,omitempty"`
	// The snapshot of the selection conditions in the task definition when the instance is created. This snapshot does not change when the task is updated.
	SelectedInput *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput `json:"SelectedInput,omitempty" xml:"SelectedInput,omitempty" type:"Struct"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetDocumentSourceCount() *int64 {
	return s.DocumentSourceCount
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntities() []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	return s.ResolvedEntities
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntityCount() *int64 {
	return s.ResolvedEntityCount
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetSelectedInput() *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	return s.SelectedInput
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetDocumentSourceCount(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.DocumentSourceCount = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntities(v []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntities = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntityCount(v int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntityCount = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetSelectedInput(v *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.SelectedInput = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) Validate() error {
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

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this value as-is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this value based on the Data Map identity. The presence of an enumeration value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityType(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityType = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) Validate() error {
	return dara.Validate(s)
}

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values are non-empty and deduplicated.
	Entities []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required, non-empty, and deduplicated when building from document objects in the current knowledge base. Each item is a KnowledgeObjectId with Type=DOCUMENT in the current knowledge base. This cannot be mixed with SourceIds for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The values are non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the current knowledge base. The values are deduplicated. Whether this parameter is required and the allowed source types depend on the task input scenario. For more information, see the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the current knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetEntities() []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	return s.Entities
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetEntities(v []*StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.Entities = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetKnowledgeObjectIds(v []*string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetProjectIds(v []*int64) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.ProjectIds = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetSourceIds(v []*string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.SourceIds = v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetYamlUploadId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.YamlUploadId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) Validate() error {
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

type StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this value as-is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this value based on the Data Map identity. The presence of an enumeration value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityId(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityType(v string) *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityType = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) Validate() error {
	return dara.Validate(s)
}
