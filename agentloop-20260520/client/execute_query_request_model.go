// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAnnotationFilter(v *ExecuteQueryRequestAnnotationFilter) *ExecuteQueryRequest
  GetAnnotationFilter() *ExecuteQueryRequestAnnotationFilter 
  SetFrom(v int32) *ExecuteQueryRequest
  GetFrom() *int32 
  SetLength(v int32) *ExecuteQueryRequest
  GetLength() *int32 
  SetMaxOutputLength(v int32) *ExecuteQueryRequest
  GetMaxOutputLength() *int32 
  SetOffset(v int32) *ExecuteQueryRequest
  GetOffset() *int32 
  SetQuery(v string) *ExecuteQueryRequest
  GetQuery() *string 
  SetTo(v int32) *ExecuteQueryRequest
  GetTo() *int32 
  SetType(v string) *ExecuteQueryRequest
  GetType() *string 
  SetVersion(v string) *ExecuteQueryRequest
  GetVersion() *string 
}

type ExecuteQueryRequest struct {
  // The annotation filter.
  AnnotationFilter *ExecuteQueryRequestAnnotationFilter `json:"annotationFilter,omitempty" xml:"annotationFilter,omitempty" type:"Struct"`
  // The start time of the query.
  // 
  // example:
  // 
  // 1760925728
  From *int32 `json:"from,omitempty" xml:"from,omitempty"`
  // The page size.
  // 
  // example:
  // 
  // 100
  Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
  // The maximum output length.
  // 
  // example:
  // 
  // 100
  MaxOutputLength *int32 `json:"maxOutputLength,omitempty" xml:"maxOutputLength,omitempty"`
  // The pagination offset.
  // 
  // example:
  // 
  // 0
  Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
  // The query entered by the user.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // select count(*) from product_faq_dataset
  Query *string `json:"query,omitempty" xml:"query,omitempty"`
  // The end time of the query.
  // 
  // example:
  // 
  // 1760925788
  To *int32 `json:"to,omitempty" xml:"to,omitempty"`
  // The statement type. Currently, only SQL is supported.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SQL
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // The dataset version.
  // 
  // example:
  // 
  // 1.0.0
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ExecuteQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryRequest) GoString() string {
  return s.String()
}

func (s *ExecuteQueryRequest) GetAnnotationFilter() *ExecuteQueryRequestAnnotationFilter  {
  return s.AnnotationFilter
}

func (s *ExecuteQueryRequest) GetFrom() *int32  {
  return s.From
}

func (s *ExecuteQueryRequest) GetLength() *int32  {
  return s.Length
}

func (s *ExecuteQueryRequest) GetMaxOutputLength() *int32  {
  return s.MaxOutputLength
}

func (s *ExecuteQueryRequest) GetOffset() *int32  {
  return s.Offset
}

func (s *ExecuteQueryRequest) GetQuery() *string  {
  return s.Query
}

func (s *ExecuteQueryRequest) GetTo() *int32  {
  return s.To
}

func (s *ExecuteQueryRequest) GetType() *string  {
  return s.Type
}

func (s *ExecuteQueryRequest) GetVersion() *string  {
  return s.Version
}

func (s *ExecuteQueryRequest) SetAnnotationFilter(v *ExecuteQueryRequestAnnotationFilter) *ExecuteQueryRequest {
  s.AnnotationFilter = v
  return s
}

func (s *ExecuteQueryRequest) SetFrom(v int32) *ExecuteQueryRequest {
  s.From = &v
  return s
}

func (s *ExecuteQueryRequest) SetLength(v int32) *ExecuteQueryRequest {
  s.Length = &v
  return s
}

func (s *ExecuteQueryRequest) SetMaxOutputLength(v int32) *ExecuteQueryRequest {
  s.MaxOutputLength = &v
  return s
}

func (s *ExecuteQueryRequest) SetOffset(v int32) *ExecuteQueryRequest {
  s.Offset = &v
  return s
}

func (s *ExecuteQueryRequest) SetQuery(v string) *ExecuteQueryRequest {
  s.Query = &v
  return s
}

func (s *ExecuteQueryRequest) SetTo(v int32) *ExecuteQueryRequest {
  s.To = &v
  return s
}

func (s *ExecuteQueryRequest) SetType(v string) *ExecuteQueryRequest {
  s.Type = &v
  return s
}

func (s *ExecuteQueryRequest) SetVersion(v string) *ExecuteQueryRequest {
  s.Version = &v
  return s
}

func (s *ExecuteQueryRequest) Validate() error {
  if s.AnnotationFilter != nil {
    if err := s.AnnotationFilter.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteQueryRequestAnnotationFilter struct {
  // The annotation filter conditions.
  Conditions []*ExecuteQueryRequestAnnotationFilterConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s ExecuteQueryRequestAnnotationFilter) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryRequestAnnotationFilter) GoString() string {
  return s.String()
}

func (s *ExecuteQueryRequestAnnotationFilter) GetConditions() []*ExecuteQueryRequestAnnotationFilterConditions  {
  return s.Conditions
}

func (s *ExecuteQueryRequestAnnotationFilter) SetConditions(v []*ExecuteQueryRequestAnnotationFilterConditions) *ExecuteQueryRequestAnnotationFilter {
  s.Conditions = v
  return s
}

func (s *ExecuteQueryRequestAnnotationFilter) Validate() error {
  if s.Conditions != nil {
    for _, item := range s.Conditions {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteQueryRequestAnnotationFilterConditions struct {
  // The annotation key.
  // 
  // example:
  // 
  // answer_quality
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // The operator.
  // 
  // example:
  // 
  // eq
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // The annotation value.
  // 
  // example:
  // 
  // GOOD
  Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ExecuteQueryRequestAnnotationFilterConditions) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryRequestAnnotationFilterConditions) GoString() string {
  return s.String()
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) GetKey() *string  {
  return s.Key
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) GetOperator() *string  {
  return s.Operator
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) GetValue() interface{}  {
  return s.Value
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) SetKey(v string) *ExecuteQueryRequestAnnotationFilterConditions {
  s.Key = &v
  return s
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) SetOperator(v string) *ExecuteQueryRequestAnnotationFilterConditions {
  s.Operator = &v
  return s
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) SetValue(v interface{}) *ExecuteQueryRequestAnnotationFilterConditions {
  s.Value = v
  return s
}

func (s *ExecuteQueryRequestAnnotationFilterConditions) Validate() error {
  return dara.Validate(s)
}

