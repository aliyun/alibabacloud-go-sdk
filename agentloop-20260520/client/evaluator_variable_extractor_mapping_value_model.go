// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluatorVariableExtractorMappingValue interface {
  dara.Model
  String() string
  GoString() string
  SetOriginField(v string) *EvaluatorVariableExtractorMappingValue
  GetOriginField() *string 
  SetType(v string) *EvaluatorVariableExtractorMappingValue
  GetType() *string 
  SetExpression(v string) *EvaluatorVariableExtractorMappingValue
  GetExpression() *string 
}

type EvaluatorVariableExtractorMappingValue struct {
  // The evaluation data field from which content is extracted. The extraction expression is applied to the content of this field. Required when saving with the evaluation task. For the trial run API, this parameter can be omitted and the backend derives it from the expression. Multiple variables can share the same source field.
  // 
  // example:
  // 
  // trace.output
  OriginField *string `json:"originField,omitempty" xml:"originField,omitempty"`
  // The extraction method. jsonpath extracts values from the JSON content of the field by using JSONPath. regex performs regular expression matching on the full text of the field. When capturing groups are present, the first capturing group is returned. When no capturing group is present, the entire match is returned.
  // 
  // example:
  // 
  // jsonpath
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // The extraction expression. Its meaning is determined by type. When type is jsonpath, specify a JSONPath expression. You can use either a relative path relative to originField (such as $.order.expected) or an absolute path from the root (such as $trace.output.order.expected). When type is regex, specify a regular expression. Note that backslashes must be escaped in JSON. The expression syntax is validated upon saving. For regular expressions, RE2 compatibility is additionally validated. Patterns such as lookahead assertions, lookbehind assertions, backreferences, named groups, atomic groups, and possessive quantifiers are rejected.
  // 
  // example:
  // 
  // $.order.expected
  Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
}

func (s EvaluatorVariableExtractorMappingValue) String() string {
  return dara.Prettify(s)
}

func (s EvaluatorVariableExtractorMappingValue) GoString() string {
  return s.String()
}

func (s *EvaluatorVariableExtractorMappingValue) GetOriginField() *string  {
  return s.OriginField
}

func (s *EvaluatorVariableExtractorMappingValue) GetType() *string  {
  return s.Type
}

func (s *EvaluatorVariableExtractorMappingValue) GetExpression() *string  {
  return s.Expression
}

func (s *EvaluatorVariableExtractorMappingValue) SetOriginField(v string) *EvaluatorVariableExtractorMappingValue {
  s.OriginField = &v
  return s
}

func (s *EvaluatorVariableExtractorMappingValue) SetType(v string) *EvaluatorVariableExtractorMappingValue {
  s.Type = &v
  return s
}

func (s *EvaluatorVariableExtractorMappingValue) SetExpression(v string) *EvaluatorVariableExtractorMappingValue {
  s.Expression = &v
  return s
}

func (s *EvaluatorVariableExtractorMappingValue) Validate() error {
  return dara.Validate(s)
}

