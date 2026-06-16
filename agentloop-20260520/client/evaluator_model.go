// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluator interface {
  dara.Model
  String() string
  GoString() string
  SetConfig(v map[string]interface{}) *Evaluator
  GetConfig() map[string]interface{} 
  SetEvaluatorRef(v string) *Evaluator
  GetEvaluatorRef() *string 
  SetFilters(v map[string]interface{}) *Evaluator
  GetFilters() map[string]interface{} 
  SetName(v string) *Evaluator
  GetName() *string 
  SetResultName(v string) *Evaluator
  GetResultName() *string 
  SetResultType(v string) *Evaluator
  GetResultType() *string 
  SetType(v string) *Evaluator
  GetType() *string 
  SetVariableMapping(v map[string]*string) *Evaluator
  GetVariableMapping() map[string]*string 
}

type Evaluator struct {
  Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
  EvaluatorRef *string `json:"evaluatorRef,omitempty" xml:"evaluatorRef,omitempty"`
  Filters map[string]interface{} `json:"filters,omitempty" xml:"filters,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  ResultName *string `json:"resultName,omitempty" xml:"resultName,omitempty"`
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  VariableMapping map[string]*string `json:"variableMapping,omitempty" xml:"variableMapping,omitempty"`
}

func (s Evaluator) String() string {
  return dara.Prettify(s)
}

func (s Evaluator) GoString() string {
  return s.String()
}

func (s *Evaluator) GetConfig() map[string]interface{}  {
  return s.Config
}

func (s *Evaluator) GetEvaluatorRef() *string  {
  return s.EvaluatorRef
}

func (s *Evaluator) GetFilters() map[string]interface{}  {
  return s.Filters
}

func (s *Evaluator) GetName() *string  {
  return s.Name
}

func (s *Evaluator) GetResultName() *string  {
  return s.ResultName
}

func (s *Evaluator) GetResultType() *string  {
  return s.ResultType
}

func (s *Evaluator) GetType() *string  {
  return s.Type
}

func (s *Evaluator) GetVariableMapping() map[string]*string  {
  return s.VariableMapping
}

func (s *Evaluator) SetConfig(v map[string]interface{}) *Evaluator {
  s.Config = v
  return s
}

func (s *Evaluator) SetEvaluatorRef(v string) *Evaluator {
  s.EvaluatorRef = &v
  return s
}

func (s *Evaluator) SetFilters(v map[string]interface{}) *Evaluator {
  s.Filters = v
  return s
}

func (s *Evaluator) SetName(v string) *Evaluator {
  s.Name = &v
  return s
}

func (s *Evaluator) SetResultName(v string) *Evaluator {
  s.ResultName = &v
  return s
}

func (s *Evaluator) SetResultType(v string) *Evaluator {
  s.ResultType = &v
  return s
}

func (s *Evaluator) SetType(v string) *Evaluator {
  s.Type = &v
  return s
}

func (s *Evaluator) SetVariableMapping(v map[string]*string) *Evaluator {
  s.VariableMapping = v
  return s
}

func (s *Evaluator) Validate() error {
  return dara.Validate(s)
}

