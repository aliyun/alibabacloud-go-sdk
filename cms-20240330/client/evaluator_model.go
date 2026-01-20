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
  SetDataScope(v string) *Evaluator
  GetDataScope() *string 
  SetFilters(v map[string]*string) *Evaluator
  GetFilters() map[string]*string 
  SetName(v string) *Evaluator
  GetName() *string 
  SetResultName(v string) *Evaluator
  GetResultName() *string 
  SetResultType(v string) *Evaluator
  GetResultType() *string 
  SetVariableMapping(v map[string]*string) *Evaluator
  GetVariableMapping() map[string]*string 
}

type Evaluator struct {
  Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
  DataScope *string `json:"dataScope,omitempty" xml:"dataScope,omitempty"`
  Filters map[string]*string `json:"filters,omitempty" xml:"filters,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  ResultName *string `json:"resultName,omitempty" xml:"resultName,omitempty"`
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
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

func (s *Evaluator) GetDataScope() *string  {
  return s.DataScope
}

func (s *Evaluator) GetFilters() map[string]*string  {
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

func (s *Evaluator) GetVariableMapping() map[string]*string  {
  return s.VariableMapping
}

func (s *Evaluator) SetConfig(v map[string]interface{}) *Evaluator {
  s.Config = v
  return s
}

func (s *Evaluator) SetDataScope(v string) *Evaluator {
  s.DataScope = &v
  return s
}

func (s *Evaluator) SetFilters(v map[string]*string) *Evaluator {
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

func (s *Evaluator) SetVariableMapping(v map[string]*string) *Evaluator {
  s.VariableMapping = v
  return s
}

func (s *Evaluator) Validate() error {
  return dara.Validate(s)
}

