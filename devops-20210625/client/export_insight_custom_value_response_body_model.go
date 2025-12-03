// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightCustomValueResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightCustomValueResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightCustomValueResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightCustomValueResponseBodyResult) *ExportInsightCustomValueResponseBody
  GetResult() []*ExportInsightCustomValueResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightCustomValueResponseBody
  GetTotalCount() *int64 
}

type ExportInsightCustomValueResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightCustomValueResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightCustomValueResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightCustomValueResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightCustomValueResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightCustomValueResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightCustomValueResponseBody) GetResult() []*ExportInsightCustomValueResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightCustomValueResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightCustomValueResponseBody) SetMaxResults(v int64) *ExportInsightCustomValueResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightCustomValueResponseBody) SetNextToken(v string) *ExportInsightCustomValueResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightCustomValueResponseBody) SetResult(v []*ExportInsightCustomValueResponseBodyResult) *ExportInsightCustomValueResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightCustomValueResponseBody) SetTotalCount(v int64) *ExportInsightCustomValueResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightCustomValueResponseBody) Validate() error {
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

type ExportInsightCustomValueResponseBodyResult struct {
  // example:
  // 
  // 66.6
  DoubleValue *float64 `json:"doubleValue,omitempty" xml:"doubleValue,omitempty"`
  // example:
  // 
  // 34dde3dfa5e3750151a7c4xxxx
  FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
  // example:
  // 
  // 1704950971000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1714669494000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 320737507
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 442d4a6a9980e841dc192a411080xxxx
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 66
  LongValue *int64 `json:"longValue,omitempty" xml:"longValue,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 1ee00fcb1a18c2dc83dafdxxxx
  TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
  // example:
  // 
  // Workitem
  TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
  // example:
  // 
  // string
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // example:
  // 
  // 66
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ExportInsightCustomValueResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightCustomValueResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightCustomValueResponseBodyResult) GetDoubleValue() *float64  {
  return s.DoubleValue
}

func (s *ExportInsightCustomValueResponseBodyResult) GetFieldId() *string  {
  return s.FieldId
}

func (s *ExportInsightCustomValueResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightCustomValueResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightCustomValueResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightCustomValueResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightCustomValueResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightCustomValueResponseBodyResult) GetLongValue() *int64  {
  return s.LongValue
}

func (s *ExportInsightCustomValueResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightCustomValueResponseBodyResult) GetTargetId() *string  {
  return s.TargetId
}

func (s *ExportInsightCustomValueResponseBodyResult) GetTargetType() *string  {
  return s.TargetType
}

func (s *ExportInsightCustomValueResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightCustomValueResponseBodyResult) GetValue() *string  {
  return s.Value
}

func (s *ExportInsightCustomValueResponseBodyResult) SetDoubleValue(v float64) *ExportInsightCustomValueResponseBodyResult {
  s.DoubleValue = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetFieldId(v string) *ExportInsightCustomValueResponseBodyResult {
  s.FieldId = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetGmtCreate(v int64) *ExportInsightCustomValueResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetGmtModified(v int64) *ExportInsightCustomValueResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetId(v int64) *ExportInsightCustomValueResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetIdentifier(v string) *ExportInsightCustomValueResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetIsDeleted(v string) *ExportInsightCustomValueResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetLongValue(v int64) *ExportInsightCustomValueResponseBodyResult {
  s.LongValue = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetOrganizationId(v string) *ExportInsightCustomValueResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetTargetId(v string) *ExportInsightCustomValueResponseBodyResult {
  s.TargetId = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetTargetType(v string) *ExportInsightCustomValueResponseBodyResult {
  s.TargetType = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetType(v string) *ExportInsightCustomValueResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) SetValue(v string) *ExportInsightCustomValueResponseBodyResult {
  s.Value = &v
  return s
}

func (s *ExportInsightCustomValueResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

