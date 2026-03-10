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
  SetEvaluators(v []*Evaluator) *ExperimentRecord
  GetEvaluators() []*Evaluator 
  SetExecutedAt(v int64) *ExperimentRecord
  GetExecutedAt() *int64 
  SetExperimentConfig(v *ExperimentConfig) *ExperimentRecord
  GetExperimentConfig() *ExperimentConfig 
  SetExperimentName(v string) *ExperimentRecord
  GetExperimentName() *string 
  SetFailedTasks(v int32) *ExperimentRecord
  GetFailedTasks() *int32 
  SetInput(v map[string]interface{}) *ExperimentRecord
  GetInput() map[string]interface{} 
  SetModelName(v string) *ExperimentRecord
  GetModelName() *string 
  SetPlanId(v string) *ExperimentRecord
  GetPlanId() *string 
  SetPlanName(v string) *ExperimentRecord
  GetPlanName() *string 
  SetProgress(v float32) *ExperimentRecord
  GetProgress() *float32 
  SetRecordId(v string) *ExperimentRecord
  GetRecordId() *string 
  SetSelectedItemIds(v []*string) *ExperimentRecord
  GetSelectedItemIds() []*string 
  SetStatus(v string) *ExperimentRecord
  GetStatus() *string 
  SetTotalTasks(v int32) *ExperimentRecord
  GetTotalTasks() *int32 
}

type ExperimentRecord struct {
  CompletedAt *int64 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
  CompletedTasks *int32 `json:"completedTasks,omitempty" xml:"completedTasks,omitempty"`
  DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
  DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
  DatasetProject *string `json:"datasetProject,omitempty" xml:"datasetProject,omitempty"`
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
  ExecutedAt *int64 `json:"executedAt,omitempty" xml:"executedAt,omitempty"`
  ExperimentConfig *ExperimentConfig `json:"experimentConfig,omitempty" xml:"experimentConfig,omitempty"`
  ExperimentName *string `json:"experimentName,omitempty" xml:"experimentName,omitempty"`
  FailedTasks *int32 `json:"failedTasks,omitempty" xml:"failedTasks,omitempty"`
  Input map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
  ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
  PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
  PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
  Progress *float32 `json:"progress,omitempty" xml:"progress,omitempty"`
  RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
  SelectedItemIds []*string `json:"selectedItemIds,omitempty" xml:"selectedItemIds,omitempty" type:"Repeated"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *ExperimentRecord) GetEvaluators() []*Evaluator  {
  return s.Evaluators
}

func (s *ExperimentRecord) GetExecutedAt() *int64  {
  return s.ExecutedAt
}

func (s *ExperimentRecord) GetExperimentConfig() *ExperimentConfig  {
  return s.ExperimentConfig
}

func (s *ExperimentRecord) GetExperimentName() *string  {
  return s.ExperimentName
}

func (s *ExperimentRecord) GetFailedTasks() *int32  {
  return s.FailedTasks
}

func (s *ExperimentRecord) GetInput() map[string]interface{}  {
  return s.Input
}

func (s *ExperimentRecord) GetModelName() *string  {
  return s.ModelName
}

func (s *ExperimentRecord) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExperimentRecord) GetPlanName() *string  {
  return s.PlanName
}

func (s *ExperimentRecord) GetProgress() *float32  {
  return s.Progress
}

func (s *ExperimentRecord) GetRecordId() *string  {
  return s.RecordId
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

func (s *ExperimentRecord) SetEvaluators(v []*Evaluator) *ExperimentRecord {
  s.Evaluators = v
  return s
}

func (s *ExperimentRecord) SetExecutedAt(v int64) *ExperimentRecord {
  s.ExecutedAt = &v
  return s
}

func (s *ExperimentRecord) SetExperimentConfig(v *ExperimentConfig) *ExperimentRecord {
  s.ExperimentConfig = v
  return s
}

func (s *ExperimentRecord) SetExperimentName(v string) *ExperimentRecord {
  s.ExperimentName = &v
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

func (s *ExperimentRecord) SetModelName(v string) *ExperimentRecord {
  s.ModelName = &v
  return s
}

func (s *ExperimentRecord) SetPlanId(v string) *ExperimentRecord {
  s.PlanId = &v
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

func (s *ExperimentRecord) SetRecordId(v string) *ExperimentRecord {
  s.RecordId = &v
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
    if err := s.ExperimentConfig.Validate(); err != nil {
      return err
    }
  }
  return nil
}

