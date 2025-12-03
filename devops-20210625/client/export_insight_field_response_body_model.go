// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightFieldResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightFieldResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightFieldResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightFieldResponseBodyResult) *ExportInsightFieldResponseBody
  GetResult() []*ExportInsightFieldResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightFieldResponseBody
  GetTotalCount() *int64 
}

type ExportInsightFieldResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightFieldResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightFieldResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightFieldResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightFieldResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightFieldResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightFieldResponseBody) GetResult() []*ExportInsightFieldResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightFieldResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightFieldResponseBody) SetMaxResults(v int64) *ExportInsightFieldResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightFieldResponseBody) SetNextToken(v string) *ExportInsightFieldResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightFieldResponseBody) SetResult(v []*ExportInsightFieldResponseBodyResult) *ExportInsightFieldResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightFieldResponseBody) SetTotalCount(v int64) *ExportInsightFieldResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightFieldResponseBody) Validate() error {
  if s.Result != nil {
    for _, item := range s.Result {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExportInsightFieldResponseBodyResult struct {
  // example:
  // 
  // 9798551
  FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
  FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
  // example:
  // 
  // 1713752162000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1714977502000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 666666
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // field-444153
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // N
  IsSystem *string `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
  OptionValue *string `json:"optionValue,omitempty" xml:"optionValue,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 150
  Position *int32 `json:"position,omitempty" xml:"position,omitempty"`
  // example:
  // 
  // global
  Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
  // example:
  // 
  // organization
  TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
  // example:
  // 
  // string
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExportInsightFieldResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightFieldResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightFieldResponseBodyResult) GetFieldId() *string  {
  return s.FieldId
}

func (s *ExportInsightFieldResponseBodyResult) GetFieldName() *string  {
  return s.FieldName
}

func (s *ExportInsightFieldResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightFieldResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightFieldResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightFieldResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightFieldResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightFieldResponseBodyResult) GetIsSystem() *string  {
  return s.IsSystem
}

func (s *ExportInsightFieldResponseBodyResult) GetOptionValue() *string  {
  return s.OptionValue
}

func (s *ExportInsightFieldResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightFieldResponseBodyResult) GetPosition() *int32  {
  return s.Position
}

func (s *ExportInsightFieldResponseBodyResult) GetScope() *string  {
  return s.Scope
}

func (s *ExportInsightFieldResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightFieldResponseBodyResult) GetTargetId() *string  {
  return s.TargetId
}

func (s *ExportInsightFieldResponseBodyResult) GetTargetType() *string  {
  return s.TargetType
}

func (s *ExportInsightFieldResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightFieldResponseBodyResult) SetFieldId(v string) *ExportInsightFieldResponseBodyResult {
  s.FieldId = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetFieldName(v string) *ExportInsightFieldResponseBodyResult {
  s.FieldName = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetGmtCreate(v int64) *ExportInsightFieldResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetGmtModified(v int64) *ExportInsightFieldResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetId(v int64) *ExportInsightFieldResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetIdentifier(v string) *ExportInsightFieldResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetIsDeleted(v string) *ExportInsightFieldResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetIsSystem(v string) *ExportInsightFieldResponseBodyResult {
  s.IsSystem = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetOptionValue(v string) *ExportInsightFieldResponseBodyResult {
  s.OptionValue = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetOrganizationId(v string) *ExportInsightFieldResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetPosition(v int32) *ExportInsightFieldResponseBodyResult {
  s.Position = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetScope(v string) *ExportInsightFieldResponseBodyResult {
  s.Scope = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetSource(v string) *ExportInsightFieldResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetTargetId(v string) *ExportInsightFieldResponseBodyResult {
  s.TargetId = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetTargetType(v string) *ExportInsightFieldResponseBodyResult {
  s.TargetType = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) SetType(v string) *ExportInsightFieldResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightFieldResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

