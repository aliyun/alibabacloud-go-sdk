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
  // The evaluator runtime configuration. For inline LLM evaluators, this must include configurations such as prompt. When referencing an existing evaluator, this parameter is typically not required and is only specified when runtime parameters such as version need to be set.
  // 
  // example:
  // 
  // {"version":"1.0.0"}
  Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
  // The reference name of a registered evaluator. When specified, the evaluator definition is loaded by this reference with higher priority. Both built-in evaluators and custom evaluators are supported.
  // 
  // example:
  // 
  // Builtin.agent_task_completion
  EvaluatorRef *string `json:"evaluatorRef,omitempty" xml:"evaluatorRef,omitempty"`
  // The evaluator-level data filter conditions. These take effect together with the task-level dataFilter.query.
  // 
  // example:
  // 
  // {"query":"serviceName=\\"checkout-service\\""}
  Filters map[string]interface{} `json:"filters,omitempty" xml:"filters,omitempty"`
  // The evaluator name. Required for inline evaluators when evaluatorRef is not specified. The evaluatorRef or name must be unique within the same task.
  // 
  // example:
  // 
  // agent_task_completion
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // The field name for the evaluation result. Required for inline evaluators. When referencing an existing evaluator, the metricName defined in the evaluator definition is used if this parameter is not specified.
  // 
  // example:
  // 
  // agent_task_completion
  ResultName *string `json:"resultName,omitempty" xml:"resultName,omitempty"`
  // The evaluation result type. Required for inline evaluators. Defaults to score when referencing an existing evaluator and this parameter is not specified.
  // 
  // example:
  // 
  // score
  ResultType *string `json:"resultType,omitempty" xml:"resultType,omitempty"`
  // The evaluator type. Defaults to LLM if not specified. Inline CODE evaluators are currently not supported. For CODE type evaluators, reference a previously created evaluator by using evaluatorRef.
  // 
  // example:
  // 
  // AGENT
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // The variable mapping that maps evaluator variables to evaluation data fields. Required for LLM/AGENT inline evaluators. When referencing an existing evaluator, the variable names must exist in the evaluator definition.
  // 
  // example:
  // 
  // {"input":"trace.input","output":"trace.output","agent_trajectory":"trace.agent_trajectory"}
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

