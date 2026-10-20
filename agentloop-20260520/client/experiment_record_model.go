// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentRecord interface {
  dara.Model
  String() string
  GoString() string
  SetCompletedAt(v int64) *ExperimentRecord
  GetCompletedAt() *int64 
  SetCompletedTasks(v int32) *ExperimentRecord
  GetCompletedTasks() *int32 
  SetDataSourceType(v string) *ExperimentRecord
  GetDataSourceType() *string 
  SetDatasetId(v string) *ExperimentRecord
  GetDatasetId() *string 
  SetDatasetProject(v string) *ExperimentRecord
  GetDatasetProject() *string 
  SetErrorMessage(v string) *ExperimentRecord
  GetErrorMessage() *string 
  SetEvaluationTaskId(v string) *ExperimentRecord
  GetEvaluationTaskId() *string 
  SetEvaluators(v []*Evaluator) *ExperimentRecord
  GetEvaluators() []*Evaluator 
  SetExecutedAt(v int64) *ExperimentRecord
  GetExecutedAt() *int64 
  SetExperimentConfig(v []*ExperimentConfig) *ExperimentRecord
  GetExperimentConfig() []*ExperimentConfig 
  SetExperimentPlanId(v string) *ExperimentRecord
  GetExperimentPlanId() *string 
  SetFailedTasks(v int32) *ExperimentRecord
  GetFailedTasks() *int32 
  SetInput(v map[string]interface{}) *ExperimentRecord
  GetInput() map[string]interface{} 
  SetModelNames(v []*string) *ExperimentRecord
  GetModelNames() []*string 
  SetPlanName(v string) *ExperimentRecord
  GetPlanName() *string 
  SetProgress(v float32) *ExperimentRecord
  GetProgress() *float32 
  SetQuerySql(v string) *ExperimentRecord
  GetQuerySql() *string 
  SetRecordId(v string) *ExperimentRecord
  GetRecordId() *string 
  SetRecordName(v string) *ExperimentRecord
  GetRecordName() *string 
  SetSelectedItemIds(v []*string) *ExperimentRecord
  GetSelectedItemIds() []*string 
  SetStatus(v string) *ExperimentRecord
  GetStatus() *string 
  SetTotalTasks(v int32) *ExperimentRecord
  GetTotalTasks() *int32 
}

type ExperimentRecord struct {
  // The completion timestamp in milliseconds.
  // 
  // example:
  // 
  // 1767229200000
  CompletedAt *int64 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
  // The number of completed tasks.
  // 
  // example:
  // 
  // 5
  CompletedTasks *int32 `json:"completedTasks,omitempty" xml:"completedTasks,omitempty"`
  // The data source type. Valid values: simple, dataset-full, and dataset-partial.
  // 
  // example:
  // 
  // dataset-full
  DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
  // The dataset ID.
  // 
  // example:
  // 
  // dataset-123456
  DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
  // The Simple Log Service (SLS) project where the dataset is located.
  // 
  // example:
  // 
  // agentspace-project
  DatasetProject *string `json:"datasetProject,omitempty" xml:"datasetProject,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // Model invocation timeout
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // The ID of the associated evaluation task.
  // 
  // example:
  // 
  // eval-task-6bec93bfa03740dd86ce2bf1496e65fb
  EvaluationTaskId *string `json:"evaluationTaskId,omitempty" xml:"evaluationTaskId,omitempty"`
  // The list of evaluators.
  // 
  // example:
  // 
  // [{"name":"Builtin.agent_correctness"}]
  Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
  // The execution timestamp in milliseconds.
  // 
  // example:
  // 
  // 1767225600000
  ExecutedAt *int64 `json:"executedAt,omitempty" xml:"executedAt,omitempty"`
  // The list of experiment configurations.
  // 
  // example:
  // 
  // [{"label":"A","name":"baseline","modelName":"qwen-max"}]
  ExperimentConfig []*ExperimentConfig `json:"experimentConfig,omitempty" xml:"experimentConfig,omitempty" type:"Repeated"`
  // The ID of the associated experiment plan.
  // 
  // example:
  // 
  // exp-plan-0242d983f5d340fd8479cf2c19eb279e
  ExperimentPlanId *string `json:"experimentPlanId,omitempty" xml:"experimentPlanId,omitempty"`
  // The number of failed tasks.
  // 
  // example:
  // 
  // 1
  FailedTasks *int32 `json:"failedTasks,omitempty" xml:"failedTasks,omitempty"`
  // The single-line input in SIMPLE mode.
  // 
  // example:
  // 
  // {"test":"test","test2":1}
  Input map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
  // The list of model names used.
  // 
  // example:
  // 
  // ["qwen-max"]
  ModelNames []*string `json:"modelNames,omitempty" xml:"modelNames,omitempty" type:"Repeated"`
  // The name of the associated experiment plan.
  // 
  // example:
  // 
  // customer-service-eval
  PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
  // The execution progress, ranging from 0 to 1.
  // 
  // example:
  // 
  // 0.5
  Progress *float32 `json:"progress,omitempty" xml:"progress,omitempty"`
  // The custom query SQL clause in dataset-partial mode.
  // 
  // example:
  // 
  // level > 2
  QuerySql *string `json:"querySql,omitempty" xml:"querySql,omitempty"`
  // The experiment record ID.
  // 
  // example:
  // 
  // exp-run-f6d419b0ed3d43a7b585948a55efc07b
  RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
  // The experiment record name.
  // 
  // example:
  // 
  // baseline-run
  RecordName *string `json:"recordName,omitempty" xml:"recordName,omitempty"`
  // The list of selected data item IDs in dataset-partial mode.
  // 
  // example:
  // 
  // ["item-123456"]
  SelectedItemIds []*string `json:"selectedItemIds,omitempty" xml:"selectedItemIds,omitempty" type:"Repeated"`
  // The status. Valid values: pending, running, completed, and failed.
  // 
  // example:
  // 
  // running
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The total number of tasks.
  // 
  // example:
  // 
  // 10
  TotalTasks *int32 `json:"totalTasks,omitempty" xml:"totalTasks,omitempty"`
}

func (s ExperimentRecord) String() string {
  return dara.Prettify(s)
}

func (s ExperimentRecord) GoString() string {
  return s.String()
}

func (s *ExperimentRecord) GetCompletedAt() *int64  {
  return s.CompletedAt
}

func (s *ExperimentRecord) GetCompletedTasks() *int32  {
  return s.CompletedTasks
}

func (s *ExperimentRecord) GetDataSourceType() *string  {
  return s.DataSourceType
}

func (s *ExperimentRecord) GetDatasetId() *string  {
  return s.DatasetId
}

func (s *ExperimentRecord) GetDatasetProject() *string  {
  return s.DatasetProject
}

func (s *ExperimentRecord) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExperimentRecord) GetEvaluationTaskId() *string  {
  return s.EvaluationTaskId
}

func (s *ExperimentRecord) GetEvaluators() []*Evaluator  {
  return s.Evaluators
}

func (s *ExperimentRecord) GetExecutedAt() *int64  {
  return s.ExecutedAt
}

func (s *ExperimentRecord) GetExperimentConfig() []*ExperimentConfig  {
  return s.ExperimentConfig
}

func (s *ExperimentRecord) GetExperimentPlanId() *string  {
  return s.ExperimentPlanId
}

func (s *ExperimentRecord) GetFailedTasks() *int32  {
  return s.FailedTasks
}

func (s *ExperimentRecord) GetInput() map[string]interface{}  {
  return s.Input
}

func (s *ExperimentRecord) GetModelNames() []*string  {
  return s.ModelNames
}

func (s *ExperimentRecord) GetPlanName() *string  {
  return s.PlanName
}

func (s *ExperimentRecord) GetProgress() *float32  {
  return s.Progress
}

func (s *ExperimentRecord) GetQuerySql() *string  {
  return s.QuerySql
}

func (s *ExperimentRecord) GetRecordId() *string  {
  return s.RecordId
}

func (s *ExperimentRecord) GetRecordName() *string  {
  return s.RecordName
}

func (s *ExperimentRecord) GetSelectedItemIds() []*string  {
  return s.SelectedItemIds
}

func (s *ExperimentRecord) GetStatus() *string  {
  return s.Status
}

func (s *ExperimentRecord) GetTotalTasks() *int32  {
  return s.TotalTasks
}

func (s *ExperimentRecord) SetCompletedAt(v int64) *ExperimentRecord {
  s.CompletedAt = &v
  return s
}

func (s *ExperimentRecord) SetCompletedTasks(v int32) *ExperimentRecord {
  s.CompletedTasks = &v
  return s
}

func (s *ExperimentRecord) SetDataSourceType(v string) *ExperimentRecord {
  s.DataSourceType = &v
  return s
}

func (s *ExperimentRecord) SetDatasetId(v string) *ExperimentRecord {
  s.DatasetId = &v
  return s
}

func (s *ExperimentRecord) SetDatasetProject(v string) *ExperimentRecord {
  s.DatasetProject = &v
  return s
}

func (s *ExperimentRecord) SetErrorMessage(v string) *ExperimentRecord {
  s.ErrorMessage = &v
  return s
}

func (s *ExperimentRecord) SetEvaluationTaskId(v string) *ExperimentRecord {
  s.EvaluationTaskId = &v
  return s
}

func (s *ExperimentRecord) SetEvaluators(v []*Evaluator) *ExperimentRecord {
  s.Evaluators = v
  return s
}

func (s *ExperimentRecord) SetExecutedAt(v int64) *ExperimentRecord {
  s.ExecutedAt = &v
  return s
}

func (s *ExperimentRecord) SetExperimentConfig(v []*ExperimentConfig) *ExperimentRecord {
  s.ExperimentConfig = v
  return s
}

func (s *ExperimentRecord) SetExperimentPlanId(v string) *ExperimentRecord {
  s.ExperimentPlanId = &v
  return s
}

func (s *ExperimentRecord) SetFailedTasks(v int32) *ExperimentRecord {
  s.FailedTasks = &v
  return s
}

func (s *ExperimentRecord) SetInput(v map[string]interface{}) *ExperimentRecord {
  s.Input = v
  return s
}

func (s *ExperimentRecord) SetModelNames(v []*string) *ExperimentRecord {
  s.ModelNames = v
  return s
}

func (s *ExperimentRecord) SetPlanName(v string) *ExperimentRecord {
  s.PlanName = &v
  return s
}

func (s *ExperimentRecord) SetProgress(v float32) *ExperimentRecord {
  s.Progress = &v
  return s
}

func (s *ExperimentRecord) SetQuerySql(v string) *ExperimentRecord {
  s.QuerySql = &v
  return s
}

func (s *ExperimentRecord) SetRecordId(v string) *ExperimentRecord {
  s.RecordId = &v
  return s
}

func (s *ExperimentRecord) SetRecordName(v string) *ExperimentRecord {
  s.RecordName = &v
  return s
}

func (s *ExperimentRecord) SetSelectedItemIds(v []*string) *ExperimentRecord {
  s.SelectedItemIds = v
  return s
}

func (s *ExperimentRecord) SetStatus(v string) *ExperimentRecord {
  s.Status = &v
  return s
}

func (s *ExperimentRecord) SetTotalTasks(v int32) *ExperimentRecord {
  s.TotalTasks = &v
  return s
}

func (s *ExperimentRecord) Validate() error {
  if s.Evaluators != nil {
    for _, item := range s.Evaluators {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.ExperimentConfig != nil {
    for _, item := range s.ExperimentConfig {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

