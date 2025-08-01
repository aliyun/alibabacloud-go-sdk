// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTraceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EvaluateTraceRequest
  GetAppName() *string 
  SetEvaluationConfig(v *EvaluationConfig) *EvaluateTraceRequest
  GetEvaluationConfig() *EvaluationConfig 
  SetEvaluationId(v string) *EvaluateTraceRequest
  GetEvaluationId() *string 
  SetMaxTime(v string) *EvaluateTraceRequest
  GetMaxTime() *string 
  SetMinTime(v string) *EvaluateTraceRequest
  GetMinTime() *string 
  SetModelConfig(v *ModelConfig) *EvaluateTraceRequest
  GetModelConfig() *ModelConfig 
}

type EvaluateTraceRequest struct {
  // The name of the application to which the trace belongs.
  // 
  // example:
  // 
  // my-llm-app
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // If the value retrieved at the JSON path is itself a JSON string, further JSON path definitions within this JSON are necessary to get the actual value.
  // 
  // This parameter is required.
  EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
  // The ID of the evaluation task. If not specified, the system randomly generates and returns an ID. You can use this ID to quickly search for evaluation results.
  // 
  // example:
  // 
  // 44aea0ee00000000be5be24b2abb8f98
  EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
  // The end time of the search time range, in UTC format.
  // 
  // example:
  // 
  // 2025-04-05 13:24:25
  // 
  // 2025-04-05
  MaxTime *string `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
  // The start time of the search time range, in UTC format.
  // 
  // example:
  // 
  // 2025-04-05 13:24:25
  // 
  // 2025-04-05
  MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
  // The configuration structure to access the model used internally by the evaluation trace.
  ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
}

func (s EvaluateTraceRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTraceRequest) GoString() string {
  return s.String()
}

func (s *EvaluateTraceRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EvaluateTraceRequest) GetEvaluationConfig() *EvaluationConfig  {
  return s.EvaluationConfig
}

func (s *EvaluateTraceRequest) GetEvaluationId() *string  {
  return s.EvaluationId
}

func (s *EvaluateTraceRequest) GetMaxTime() *string  {
  return s.MaxTime
}

func (s *EvaluateTraceRequest) GetMinTime() *string  {
  return s.MinTime
}

func (s *EvaluateTraceRequest) GetModelConfig() *ModelConfig  {
  return s.ModelConfig
}

func (s *EvaluateTraceRequest) SetAppName(v string) *EvaluateTraceRequest {
  s.AppName = &v
  return s
}

func (s *EvaluateTraceRequest) SetEvaluationConfig(v *EvaluationConfig) *EvaluateTraceRequest {
  s.EvaluationConfig = v
  return s
}

func (s *EvaluateTraceRequest) SetEvaluationId(v string) *EvaluateTraceRequest {
  s.EvaluationId = &v
  return s
}

func (s *EvaluateTraceRequest) SetMaxTime(v string) *EvaluateTraceRequest {
  s.MaxTime = &v
  return s
}

func (s *EvaluateTraceRequest) SetMinTime(v string) *EvaluateTraceRequest {
  s.MinTime = &v
  return s
}

func (s *EvaluateTraceRequest) SetModelConfig(v *ModelConfig) *EvaluateTraceRequest {
  s.ModelConfig = v
  return s
}

func (s *EvaluateTraceRequest) Validate() error {
  return dara.Validate(s)
}

