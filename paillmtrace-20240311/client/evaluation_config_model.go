// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluationConfig interface {
  dara.Model
  String() string
  GoString() string
  SetAnswer(v *EvaluationConfigAnswer) *EvaluationConfig
  GetAnswer() *EvaluationConfigAnswer 
  SetContext(v *EvaluationConfigContext) *EvaluationConfig
  GetContext() *EvaluationConfigContext 
  SetQuery(v *EvaluationConfigQuery) *EvaluationConfig
  GetQuery() *EvaluationConfigQuery 
}

type EvaluationConfig struct {
  Answer *EvaluationConfigAnswer `json:"Answer,omitempty" xml:"Answer,omitempty" type:"Struct"`
  Context *EvaluationConfigContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
  Query *EvaluationConfigQuery `json:"Query,omitempty" xml:"Query,omitempty" type:"Struct"`
}

func (s EvaluationConfig) String() string {
  return dara.Prettify(s)
}

func (s EvaluationConfig) GoString() string {
  return s.String()
}

func (s *EvaluationConfig) GetAnswer() *EvaluationConfigAnswer  {
  return s.Answer
}

func (s *EvaluationConfig) GetContext() *EvaluationConfigContext  {
  return s.Context
}

func (s *EvaluationConfig) GetQuery() *EvaluationConfigQuery  {
  return s.Query
}

func (s *EvaluationConfig) SetAnswer(v *EvaluationConfigAnswer) *EvaluationConfig {
  s.Answer = v
  return s
}

func (s *EvaluationConfig) SetContext(v *EvaluationConfigContext) *EvaluationConfig {
  s.Context = v
  return s
}

func (s *EvaluationConfig) SetQuery(v *EvaluationConfigQuery) *EvaluationConfig {
  s.Query = v
  return s
}

func (s *EvaluationConfig) Validate() error {
  return dara.Validate(s)
}

type EvaluationConfigAnswer struct {
  JsonPathInSpan *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
  JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
  SpanName *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigAnswer) String() string {
  return dara.Prettify(s)
}

func (s EvaluationConfigAnswer) GoString() string {
  return s.String()
}

func (s *EvaluationConfigAnswer) GetJsonPathInSpan() *string  {
  return s.JsonPathInSpan
}

func (s *EvaluationConfigAnswer) GetJsonPathInSpanValue() *string  {
  return s.JsonPathInSpanValue
}

func (s *EvaluationConfigAnswer) GetSpanName() *string  {
  return s.SpanName
}

func (s *EvaluationConfigAnswer) SetJsonPathInSpan(v string) *EvaluationConfigAnswer {
  s.JsonPathInSpan = &v
  return s
}

func (s *EvaluationConfigAnswer) SetJsonPathInSpanValue(v string) *EvaluationConfigAnswer {
  s.JsonPathInSpanValue = &v
  return s
}

func (s *EvaluationConfigAnswer) SetSpanName(v string) *EvaluationConfigAnswer {
  s.SpanName = &v
  return s
}

func (s *EvaluationConfigAnswer) Validate() error {
  return dara.Validate(s)
}

type EvaluationConfigContext struct {
  JsonPathInSpan *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
  JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
  SpanName *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigContext) String() string {
  return dara.Prettify(s)
}

func (s EvaluationConfigContext) GoString() string {
  return s.String()
}

func (s *EvaluationConfigContext) GetJsonPathInSpan() *string  {
  return s.JsonPathInSpan
}

func (s *EvaluationConfigContext) GetJsonPathInSpanValue() *string  {
  return s.JsonPathInSpanValue
}

func (s *EvaluationConfigContext) GetSpanName() *string  {
  return s.SpanName
}

func (s *EvaluationConfigContext) SetJsonPathInSpan(v string) *EvaluationConfigContext {
  s.JsonPathInSpan = &v
  return s
}

func (s *EvaluationConfigContext) SetJsonPathInSpanValue(v string) *EvaluationConfigContext {
  s.JsonPathInSpanValue = &v
  return s
}

func (s *EvaluationConfigContext) SetSpanName(v string) *EvaluationConfigContext {
  s.SpanName = &v
  return s
}

func (s *EvaluationConfigContext) Validate() error {
  return dara.Validate(s)
}

type EvaluationConfigQuery struct {
  JsonPathInSpan *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
  JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
  SpanName *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigQuery) String() string {
  return dara.Prettify(s)
}

func (s EvaluationConfigQuery) GoString() string {
  return s.String()
}

func (s *EvaluationConfigQuery) GetJsonPathInSpan() *string  {
  return s.JsonPathInSpan
}

func (s *EvaluationConfigQuery) GetJsonPathInSpanValue() *string  {
  return s.JsonPathInSpanValue
}

func (s *EvaluationConfigQuery) GetSpanName() *string  {
  return s.SpanName
}

func (s *EvaluationConfigQuery) SetJsonPathInSpan(v string) *EvaluationConfigQuery {
  s.JsonPathInSpan = &v
  return s
}

func (s *EvaluationConfigQuery) SetJsonPathInSpanValue(v string) *EvaluationConfigQuery {
  s.JsonPathInSpanValue = &v
  return s
}

func (s *EvaluationConfigQuery) SetSpanName(v string) *EvaluationConfigQuery {
  s.SpanName = &v
  return s
}

func (s *EvaluationConfigQuery) Validate() error {
  return dara.Validate(s)
}

