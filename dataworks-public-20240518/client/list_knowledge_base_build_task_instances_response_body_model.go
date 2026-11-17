// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseBuildTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetData() *ListKnowledgeBaseBuildTaskInstancesResponseBodyData
	SetMaxResults(v int32) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeBaseBuildTaskInstancesResponseBody
	GetSuccess() *bool
}

type ListKnowledgeBaseBuildTaskInstancesResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeBaseBuildTaskInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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

func (s ListKnowledgeBaseBuildTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetData() *ListKnowledgeBaseBuildTaskInstancesResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetCode(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetData(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetMaxResults(v int32) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetMessage(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetRequestId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) SetSuccess(v bool) *ListKnowledgeBaseBuildTaskInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyData struct {
	// The knowledge objects that you have permissions to read on the current page.
	Instances []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The pagination token for the next page. This field is null if there is no next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) GetInstances() []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) SetInstances(v []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) *ListKnowledgeBaseBuildTaskInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) SetNextToken(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances struct {
	// The snapshot of the model, content language, and runtime resources for the execution. This snapshot does not change when the task is updated.
	ConfigSnapshot *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current business stage, returned when reliable records are available. For standard build stages and semantic model import stages, refer to the domain model. This field does not replace the overall status.
	//
	// example:
	//
	// INITIALIZE
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The error code and secure error message returned when the execution fails.
	Error *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The time when the execution finished. This field is optional.
	//
	// example:
	//
	// 1790200060000
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The confirmed ModelId and VersionNo for model imports only. For their structure, see GetKnowledgeBaseBuildTaskInstance.
	GeneratedModels []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels `json:"GeneratedModels,omitempty" xml:"GeneratedModels,omitempty" type:"Repeated"`
	// The non-sensitive input snapshot after selection and parsing for the execution. This snapshot does not change when the task is updated.
	InputSnapshot *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot `json:"InputSnapshot,omitempty" xml:"InputSnapshot,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the instance belongs. Use it for scope validation and queries.
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
	// The overall status of the execution.
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
	// The task instance ID.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
	// The trigger type of the execution, which can be manual, periodic, or a supported event.
	//
	// example:
	//
	// MANUAL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetConfigSnapshot() *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot {
	return s.ConfigSnapshot
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetError() *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError {
	return s.Error
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetGeneratedModels() []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels {
	return s.GeneratedModels
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetInputSnapshot() *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	return s.InputSnapshot
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetTaskId() *string {
	return s.TaskId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetConfigSnapshot(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.ConfigSnapshot = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetCreateTime(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetCurrentStage(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.CurrentStage = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetError(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.Error = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetFinishedTime(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.FinishedTime = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetGeneratedModels(v []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.GeneratedModels = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetInputSnapshot(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.InputSnapshot = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetKnowledgeBaseId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetStartedTime(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.StartedTime = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetStatus(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetTaskId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.TaskId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetTaskInstanceId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.TaskInstanceId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) SetTriggerType(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances {
	s.TriggerType = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstances) Validate() error {
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

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot struct {
	// The content language used for the execution.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// The specific model ID actually used and frozen for the execution. AUTO is not returned. This field can be empty for tasks that do not call models, such as pure synchronization tasks.
	//
	// example:
	//
	// model-001
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
	// The DataWorks workspace ID used for the execution, frozen from the task definition when the instance was created.
	//
	// example:
	//
	// 67890
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group and Builder CUs actually used for the execution.
	RuntimeResource *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) GetLlmModel() *string {
	return s.LlmModel
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) GetRuntimeResource() *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource {
	return s.RuntimeResource
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) SetContentLocale(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot {
	s.ContentLocale = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) SetLlmModel(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot {
	s.LlmModel = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) SetProjectId(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot {
	s.ProjectId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) SetRuntimeResource(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot {
	s.RuntimeResource = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshot) Validate() error {
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource struct {
	// The Builder compute specifications. Valid values: 4 to 16. If omitted during creation, the default value is 4. The read operation returns the effective value.
	//
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The DataWorks resource group ID used by the actual Builder nodes. The server verifies its availability in the workspace corresponding to the top-level ProjectId of the task.
	//
	// example:
	//
	// rg-001
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) GetCu() *int32 {
	return s.Cu
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) SetCu(v int32) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) SetResourceGroupId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesConfigSnapshotRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError struct {
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

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) SetCode(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) SetMessage(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesError) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels struct {
	// The model ID generated by the semantic layer. You can use it as a KnowledgeObjectId to read the corresponding semantic model.
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

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) GetModelId() *string {
	return s.ModelId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) SetModelId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels {
	s.ModelId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) SetVersionNo(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels {
	s.VersionNo = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesGeneratedModels) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot struct {
	// The number of external document sources selected for the execution. This field is not returned when not applicable.
	//
	// example:
	//
	// 1
	DocumentSourceCount *int64 `json:"DocumentSourceCount,omitempty" xml:"DocumentSourceCount,omitempty"`
	// The relative file paths in the knowledge base processed for page upload builds in the execution. These paths are from the frozen upload list of the instance. OSS addresses or verification information are not returned, and the current task or file list is not re-read.
	FilePaths []*string `json:"FilePaths,omitempty" xml:"FilePaths,omitempty" type:"Repeated"`
	// The actual Data Map entities resolved from the project, database, or schema scope for the execution. This field is not returned when not applicable.
	ResolvedEntities []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities `json:"ResolvedEntities,omitempty" xml:"ResolvedEntities,omitempty" type:"Repeated"`
	// The total number of resolved entities. An accurate count is returned even if the array is omitted due to its size.
	//
	// example:
	//
	// 42
	ResolvedEntityCount *int64 `json:"ResolvedEntityCount,omitempty" xml:"ResolvedEntityCount,omitempty"`
	// The snapshot of the selection conditions in the task definition when the instance was created. This snapshot does not change when the task is updated. This field is omitted if no fixed object is selected for page upload builds.
	SelectedInput *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput `json:"SelectedInput,omitempty" xml:"SelectedInput,omitempty" type:"Struct"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GetDocumentSourceCount() *int64 {
	return s.DocumentSourceCount
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GetResolvedEntities() []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities {
	return s.ResolvedEntities
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GetResolvedEntityCount() *int64 {
	return s.ResolvedEntityCount
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) GetSelectedInput() *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	return s.SelectedInput
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) SetDocumentSourceCount(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	s.DocumentSourceCount = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) SetFilePaths(v []*string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	s.FilePaths = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) SetResolvedEntities(v []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	s.ResolvedEntities = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) SetResolvedEntityCount(v int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	s.ResolvedEntityCount = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) SetSelectedInput(v *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot {
	s.SelectedInput = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshot) Validate() error {
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

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities struct {
	// The metadata entity ID returned by Data Map. The client must pass this ID as is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this based on Data Map identities. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) SetEntityId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities {
	s.EntityId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) SetEntityType(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities {
	s.EntityType = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotResolvedEntities) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput struct {
	// The Data Map entities, such as databases, schemas, and tables, selected for data asset-scoped builds. The list is non-empty and deduplicated.
	Entities []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// Required, non-empty, and deduplicated when building from document objects in the knowledge base. Each item is a KnowledgeObjectId with Type=DOCUMENT in the knowledge base. This field cannot be mixed with SourceIds used for external connector synchronization scopes.
	KnowledgeObjectIds []*string `json:"KnowledgeObjectIds,omitempty" xml:"KnowledgeObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace IDs selected for workspace-scoped builds. The list is non-empty and deduplicated.
	ProjectIds []*int64 `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The IDs of the registered sources in the knowledge base. The list is deduplicated. Whether this field is required and the allowed source types depend on the task input scenario. Refer to the following table.
	SourceIds []*string `json:"SourceIds,omitempty" xml:"SourceIds,omitempty" type:"Repeated"`
	// The YAML upload session ID in the knowledge base used for SEMANTIC_MODEL import tasks. This ID is bound when the task is created and cannot be mixed with other input scopes.
	//
	// example:
	//
	// upload-yaml-001
	YamlUploadId *string `json:"YamlUploadId,omitempty" xml:"YamlUploadId,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GetEntities() []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities {
	return s.Entities
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GetKnowledgeObjectIds() []*string {
	return s.KnowledgeObjectIds
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) GetYamlUploadId() *string {
	return s.YamlUploadId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) SetEntities(v []*ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	s.Entities = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) SetKnowledgeObjectIds(v []*string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	s.KnowledgeObjectIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) SetProjectIds(v []*int64) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	s.ProjectIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) SetSourceIds(v []*string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	s.SourceIds = v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) SetYamlUploadId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput {
	s.YamlUploadId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInput) Validate() error {
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

type ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities struct {
	// The metadata entity ID returned by Data Map. The client must pass this ID as is without concatenating database, schema, or table identifiers.
	//
	// example:
	//
	// maxcompute-table:::project_name:schema_name:table_name
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity level, such as DATABASE, SCHEMA, or TABLE. The server validates this based on Data Map identities. The presence of an enum value does not guarantee that all engines support it.
	//
	// example:
	//
	// TABLE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) SetEntityId(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities {
	s.EntityId = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) SetEntityType(v string) *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities {
	s.EntityType = &v
	return s
}

func (s *ListKnowledgeBaseBuildTaskInstancesResponseBodyDataInstancesInputSnapshotSelectedInputEntities) Validate() error {
	return dara.Validate(s)
}
