// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody
	GetCode() *string
	SetData(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) *GetKnowledgeBaseBuildTaskInstanceResponseBody
	GetData() *GetKnowledgeBaseBuildTaskInstanceResponseBodyData
	SetMessage(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgeBaseBuildTaskInstanceResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBaseBuildTaskInstanceResponseBody struct {
	// The business error code. You can use this code to determine the error type when a request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *GetKnowledgeBaseBuildTaskInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetKnowledgeBaseBuildTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) GetData() *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) SetCode(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) SetData(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) SetMessage(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) SetRequestId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) SetSuccess(v bool) *GetKnowledgeBaseBuildTaskInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyData struct {
	// The snapshot of the model, content language, and runtime resources for the current run. This snapshot does not change when the task is updated.
	ConfigSnapshot *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty" type:"Struct"`
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
	Error *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The time when the execution finished. This field is optional.
	//
	// example:
	//
	// 1790200060000
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The confirmed materialization results of the semantic model import instance. For the structures of ModelId and VersionNo, refer to the YAML import results section.
	GeneratedModels []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels `json:"GeneratedModels,omitempty" xml:"GeneratedModels,omitempty" type:"Repeated"`
	// The non-sensitive input snapshot after selection and parsing for the current run. This snapshot does not change when the task is updated.
	InputSnapshot *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot `json:"InputSnapshot,omitempty" xml:"InputSnapshot,omitempty" type:"Struct"`
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
	// The overall status. Valid values: PENDING, RUNNING, SUCCEEDED, FAILED, and STOPPED. The STOPPED status includes instances that are manually stopped or terminated because they are replaced by subsequent builds. The system no longer waits for these instances to succeed.
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
	// The execution ID.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The trigger type of the current run, such as manual, periodic, or a supported event.
	//
	// example:
	//
	// MANUAL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetConfigSnapshot() *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetError() *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	return s.Error
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetGeneratedModels() []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	return s.GeneratedModels
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetInputSnapshot() *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	return s.InputSnapshot
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetConfigSnapshot(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.ConfigSnapshot = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCreateTime(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetCurrentStage(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.CurrentStage = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetError(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Error = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetFinishedTime(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.FinishedTime = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetGeneratedModels(v []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.GeneratedModels = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetInputSnapshot(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.InputSnapshot = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStartedTime(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetStatus(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TaskInstanceId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) SetTriggerType(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyData) Validate() error {
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

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot struct {
	// The content language used for the current run.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The specific model ID actually used and frozen for the current run. The value AUTO is not returned. This field can be empty for tasks that do not call models, such as pure synchronization tasks.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The DataWorks workspace ID used for the current run. This ID is frozen from the task definition when the instance is created.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder compute units (CUs) actually used for the current run.
	RuntimeResource *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetLlmModel() *string {
	return s.LlmModel
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) GetRuntimeResource() *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	return s.RuntimeResource
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetContentLocale(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ContentLocale = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetLlmModel(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.LlmModel = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetProjectId(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.ProjectId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) SetRuntimeResource(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot {
	s.RuntimeResource = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshot) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If this field is omitted during creation, the valid value is 4. The read operation returns the valid value.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The DataWorks resource group ID used by the actual Builder nodes. The server verifies that this resource group is available in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetCu(v int32) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) SetResourceGroupId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataConfigSnapshotRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError struct {
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

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetCode(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) SetMessage(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels struct {
	// The stable model ID in the semantic layer, which can be used as a KnowledgeObjectId.
	//
	// example:
	//
	// 01MODELULID
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The materialization version for the current run. The same version can be edited later and does not represent an immutable file snapshot.
	//
	// example:
	//
	// 1
	VersionNo *int64 `json:"VersionNo,omitempty" xml:"VersionNo,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetModelId() *string {
	return s.ModelId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetModelId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.ModelId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) SetVersionNo(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels {
	s.VersionNo = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataGeneratedModels) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot struct {
	// The number of external document sources selected for the current run. This field is not returned when it is not applicable.
	//
	// example:
	//
	// 1
	DocumentSourceCount *int64 `json:"DocumentSourceCount,omitempty" xml:"DocumentSourceCount,omitempty"`
	// The relative file paths within the knowledge base processed for the current page upload build. These paths are from the frozen upload list of the instance. OSS addresses or verification information is not returned, and the current task or file list is not re-read.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	// The actual Data Map entities resolved from the project, database, or schema scope for the current run. This field is not returned when it is not applicable.
	ResolvedEntities []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities `json:"ResolvedEntities,omitempty" xml:"ResolvedEntities,omitempty" type:"Repeated"`
	// The total number of resolved entities. An accurate count is returned even if the array is not returned due to its size.
	//
	// example:
	//
	// 42
	ResolvedEntityCount *int64 `json:"ResolvedEntityCount,omitempty" xml:"ResolvedEntityCount,omitempty"`
	// The snapshot of the selection conditions in the task definition when the instance is created. This snapshot does not change when the task is updated. This field is omitted for page upload builds that do not have fixed object selections.
	SelectedInput *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput `json:"SelectedInput,omitempty" xml:"SelectedInput,omitempty" type:"Struct"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetDocumentSourceCount() *int64 {
	return s.DocumentSourceCount
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntities() []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	return s.ResolvedEntities
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetResolvedEntityCount() *int64 {
	return s.ResolvedEntityCount
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) GetSelectedInput() *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	return s.SelectedInput
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetDocumentSourceCount(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.DocumentSourceCount = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetFilePaths(v []*string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.FilePaths = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntities(v []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntities = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetResolvedEntityCount(v int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.ResolvedEntityCount = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) SetSelectedInput(v *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot {
	s.SelectedInput = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshot) Validate() error {
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

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
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

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) SetEntityType(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities {
	s.EntityType = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotResolvedEntities) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput struct {
	// The Data Map entities such as databases, schemas, and tables selected for data asset-scoped builds. The values are non-empty and deduplicated.
	Entities []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The knowledge object IDs required when building from document objects in the current knowledge base. The values must be non-empty and deduplicated. Each item is a KnowledgeObjectId with Type=DOCUMENT in the current knowledge base. Do not mix these IDs with the SourceIds used for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The values are non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the current knowledge base. The values are deduplicated. Whether this field is required and the allowed source types depend on the task input scenario. Refer to the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the current knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetEntities() []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	return s.Entities
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetEntities(v []*GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.Entities = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetKnowledgeObjectIds(v []*string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetProjectIds(v []*int64) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.ProjectIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetSourceIds(v []*string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.SourceIds = v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) SetYamlUploadId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput {
	s.YamlUploadId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInput) Validate() error {
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

type GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities struct {
	// The metadata entity ID returned by Data Map. The client passes this ID as is and does not concatenate database, schema, or table identifiers.
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

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityId(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) SetEntityType(v string) *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities {
	s.EntityType = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceResponseBodyDataInputSnapshotSelectedInputEntities) Validate() error {
	return dara.Validate(s)
}
