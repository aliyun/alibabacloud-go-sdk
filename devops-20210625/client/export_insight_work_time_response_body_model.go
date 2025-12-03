// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkTimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightWorkTimeResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkTimeResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightWorkTimeResponseBodyResult) *ExportInsightWorkTimeResponseBody
  GetResult() []*ExportInsightWorkTimeResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightWorkTimeResponseBody
  GetTotalCount() *int64 
}

type ExportInsightWorkTimeResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightWorkTimeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightWorkTimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkTimeResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkTimeResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkTimeResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkTimeResponseBody) GetResult() []*ExportInsightWorkTimeResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightWorkTimeResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightWorkTimeResponseBody) SetMaxResults(v int64) *ExportInsightWorkTimeResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBody) SetNextToken(v string) *ExportInsightWorkTimeResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBody) SetResult(v []*ExportInsightWorkTimeResponseBodyResult) *ExportInsightWorkTimeResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightWorkTimeResponseBody) SetTotalCount(v int64) *ExportInsightWorkTimeResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBody) Validate() error {
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

type ExportInsightWorkTimeResponseBodyResult struct {
  // example:
  // 
  // 120
  ActualValue *float64 `json:"actualValue,omitempty" xml:"actualValue,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // example:
  // 
  // 1714978610000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1715011199999
  GmtEnd *int64 `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
  // example:
  // 
  // 1714978610000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 1714924800000
  GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
  // example:
  // 
  // 49506082
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // da70ce5824231ca3c04ef808e0
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 09670872890eb1a0bb998exxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  // example:
  // 
  // 65659358c319d2a0f912xxxx
  RecorderId *string `json:"recorderId,omitempty" xml:"recorderId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // example:
  // 
  // 000000000cd82d3df50d5e5a5c094094fd7b4461
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
  // example:
  // 
  // 120
  Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
  // example:
  // 
  // 17bc1cf9a037a15fc9ce76xxxx
  WorkitemId *string `json:"workitemId,omitempty" xml:"workitemId,omitempty"`
}

func (s ExportInsightWorkTimeResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkTimeResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetActualValue() *float64  {
  return s.ActualValue
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetDescription() *string  {
  return s.Description
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetGmtEnd() *int64  {
  return s.GmtEnd
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetGmtStart() *int64  {
  return s.GmtStart
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetRecorderId() *string  {
  return s.RecorderId
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetUuid() *string  {
  return s.Uuid
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetValue() *int64  {
  return s.Value
}

func (s *ExportInsightWorkTimeResponseBodyResult) GetWorkitemId() *string  {
  return s.WorkitemId
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetActualValue(v float64) *ExportInsightWorkTimeResponseBodyResult {
  s.ActualValue = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetDescription(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.Description = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetGmtCreate(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetGmtEnd(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.GmtEnd = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetGmtModified(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetGmtStart(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.GmtStart = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetId(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetIdentifier(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetIsDeleted(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetOrganizationId(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetProjectId(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetRecorderId(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.RecorderId = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetSource(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetType(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetUuid(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.Uuid = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetValue(v int64) *ExportInsightWorkTimeResponseBodyResult {
  s.Value = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) SetWorkitemId(v string) *ExportInsightWorkTimeResponseBodyResult {
  s.WorkitemId = &v
  return s
}

func (s *ExportInsightWorkTimeResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

