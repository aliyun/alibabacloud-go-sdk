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
  SetPlanId(v string) *ExperimentPlanData
  GetPlanId() *string 
  SetPlanName(v string) *ExperimentPlanData
  GetPlanName() *string 
  SetStatus(v string) *ExperimentPlanData
  GetStatus() *string 
  SetUpdatedAt(v int64) *ExperimentPlanData
  GetUpdatedAt() *int64 
}

type ExperimentPlanData struct {
  CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
  DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  ExperimentCount *int32 `json:"experimentCount,omitempty" xml:"experimentCount,omitempty"`
  PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
  PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *ExperimentPlanData) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExperimentPlanData) GetPlanName() *string  {
  return s.PlanName
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

func (s *ExperimentPlanData) SetPlanId(v string) *ExperimentPlanData {
  s.PlanId = &v
  return s
}

func (s *ExperimentPlanData) SetPlanName(v string) *ExperimentPlanData {
  s.PlanName = &v
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

