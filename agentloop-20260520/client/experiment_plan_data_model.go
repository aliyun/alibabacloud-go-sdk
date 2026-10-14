// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentPlanData interface {
  dara.Model
  String() string
  GoString() string
  SetCreatedAt(v int64) *ExperimentPlanData
  GetCreatedAt() *int64 
  SetDatasetId(v string) *ExperimentPlanData
  GetDatasetId() *string 
  SetDescription(v string) *ExperimentPlanData
  GetDescription() *string 
  SetExperimentCount(v int32) *ExperimentPlanData
  GetExperimentCount() *int32 
  SetExperimentType(v string) *ExperimentPlanData
  GetExperimentType() *string 
  SetPipelineName(v string) *ExperimentPlanData
  GetPipelineName() *string 
  SetPlanId(v string) *ExperimentPlanData
  GetPlanId() *string 
  SetPlanName(v string) *ExperimentPlanData
  GetPlanName() *string 
  SetQuerySql(v string) *ExperimentPlanData
  GetQuerySql() *string 
  SetStatus(v string) *ExperimentPlanData
  GetStatus() *string 
  SetUpdatedAt(v int64) *ExperimentPlanData
  GetUpdatedAt() *int64 
}

type ExperimentPlanData struct {
  // The creation time in milliseconds.
  // 
  // example:
  // 
  // 1767225600000
  CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  // The ID of the associated dataset.
  // 
  // example:
  // 
  // dataset-123456
  DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
  // The description.
  // 
  // example:
  // 
  // Offline evaluation plan for customer service Q&A
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // The number of experiment executions under this plan.
  // 
  // example:
  // 
  // 3
  ExperimentCount *int32 `json:"experimentCount,omitempty" xml:"experimentCount,omitempty"`
  // The experiment type. Valid values: offline (offline experiment) and online (online experiment).
  // 
  // example:
  // 
  // offline
  ExperimentType *string `json:"experimentType,omitempty" xml:"experimentType,omitempty"`
  // The name of the associated data processing pipeline. This field is empty if no pipeline is associated. After association, when an experiment under this plan completes and writes results to a Logstore, the system filters by the trace ID of the experiment trace and calls PreviewPipeline to write the pipeline-processed results.
  // 
  // example:
  // 
  // trace-process-pipeline
  PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
  // The experiment plan ID.
  PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
  // The experiment plan name.
  // 
  // example:
  // 
  // customer-service-eval
  PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
  // The custom query SQL clause in dataset-partial mode.
  // 
  // example:
  // 
  // level > 2
  QuerySql *string `json:"querySql,omitempty" xml:"querySql,omitempty"`
  // The status. Valid values: pending, running, and stopped.
  // 
  // example:
  // 
  // pending
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The update time in milliseconds.
  // 
  // example:
  // 
  // 1767229200000
  UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ExperimentPlanData) String() string {
  return dara.Prettify(s)
}

func (s ExperimentPlanData) GoString() string {
  return s.String()
}

func (s *ExperimentPlanData) GetCreatedAt() *int64  {
  return s.CreatedAt
}

func (s *ExperimentPlanData) GetDatasetId() *string  {
  return s.DatasetId
}

func (s *ExperimentPlanData) GetDescription() *string  {
  return s.Description
}

func (s *ExperimentPlanData) GetExperimentCount() *int32  {
  return s.ExperimentCount
}

func (s *ExperimentPlanData) GetExperimentType() *string  {
  return s.ExperimentType
}

func (s *ExperimentPlanData) GetPipelineName() *string  {
  return s.PipelineName
}

func (s *ExperimentPlanData) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExperimentPlanData) GetPlanName() *string  {
  return s.PlanName
}

func (s *ExperimentPlanData) GetQuerySql() *string  {
  return s.QuerySql
}

func (s *ExperimentPlanData) GetStatus() *string  {
  return s.Status
}

func (s *ExperimentPlanData) GetUpdatedAt() *int64  {
  return s.UpdatedAt
}

func (s *ExperimentPlanData) SetCreatedAt(v int64) *ExperimentPlanData {
  s.CreatedAt = &v
  return s
}

func (s *ExperimentPlanData) SetDatasetId(v string) *ExperimentPlanData {
  s.DatasetId = &v
  return s
}

func (s *ExperimentPlanData) SetDescription(v string) *ExperimentPlanData {
  s.Description = &v
  return s
}

func (s *ExperimentPlanData) SetExperimentCount(v int32) *ExperimentPlanData {
  s.ExperimentCount = &v
  return s
}

func (s *ExperimentPlanData) SetExperimentType(v string) *ExperimentPlanData {
  s.ExperimentType = &v
  return s
}

func (s *ExperimentPlanData) SetPipelineName(v string) *ExperimentPlanData {
  s.PipelineName = &v
  return s
}

func (s *ExperimentPlanData) SetPlanId(v string) *ExperimentPlanData {
  s.PlanId = &v
  return s
}

func (s *ExperimentPlanData) SetPlanName(v string) *ExperimentPlanData {
  s.PlanName = &v
  return s
}

func (s *ExperimentPlanData) SetQuerySql(v string) *ExperimentPlanData {
  s.QuerySql = &v
  return s
}

func (s *ExperimentPlanData) SetStatus(v string) *ExperimentPlanData {
  s.Status = &v
  return s
}

func (s *ExperimentPlanData) SetUpdatedAt(v int64) *ExperimentPlanData {
  s.UpdatedAt = &v
  return s
}

func (s *ExperimentPlanData) Validate() error {
  return dara.Validate(s)
}

