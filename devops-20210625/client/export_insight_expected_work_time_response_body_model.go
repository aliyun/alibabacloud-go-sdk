// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightExpectedWorkTimeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightExpectedWorkTimeResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightExpectedWorkTimeResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightExpectedWorkTimeResponseBodyResult) *ExportInsightExpectedWorkTimeResponseBody
  GetResult() []*ExportInsightExpectedWorkTimeResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightExpectedWorkTimeResponseBody
  GetTotalCount() *int64 
}

type ExportInsightExpectedWorkTimeResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightExpectedWorkTimeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightExpectedWorkTimeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightExpectedWorkTimeResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightExpectedWorkTimeResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightExpectedWorkTimeResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightExpectedWorkTimeResponseBody) GetResult() []*ExportInsightExpectedWorkTimeResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightExpectedWorkTimeResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightExpectedWorkTimeResponseBody) SetMaxResults(v int64) *ExportInsightExpectedWorkTimeResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBody) SetNextToken(v string) *ExportInsightExpectedWorkTimeResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBody) SetResult(v []*ExportInsightExpectedWorkTimeResponseBodyResult) *ExportInsightExpectedWorkTimeResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBody) SetTotalCount(v int64) *ExportInsightExpectedWorkTimeResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBody) Validate() error {
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

type ExportInsightExpectedWorkTimeResponseBodyResult struct {
  // example:
  // 
  // 1714976497000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1714976520000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 26281535
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // bd4ddc7b0ea0ef2ab52699xxxx
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 61db9af2148974246be6xxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 6c4687b0179e1d458fedf1xxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  // example:
  // 
  // 63466a385dc8531eebd7xxxx
  RecorderId *string `json:"recorderId,omitempty" xml:"recorderId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // example:
  // 
  // 180
  Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
  // example:
  // 
  // de7c6fd3bd4b53f4d9e279xxxx
  WorkitemId *string `json:"workitemId,omitempty" xml:"workitemId,omitempty"`
}

func (s ExportInsightExpectedWorkTimeResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightExpectedWorkTimeResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetRecorderId() *string  {
  return s.RecorderId
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetValue() *float64  {
  return s.Value
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) GetWorkitemId() *string  {
  return s.WorkitemId
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetGmtCreate(v int64) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetGmtModified(v int64) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetId(v int64) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetIdentifier(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetIsDeleted(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetOrganizationId(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetProjectId(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetRecorderId(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.RecorderId = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetSource(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetType(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetValue(v float64) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.Value = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) SetWorkitemId(v string) *ExportInsightExpectedWorkTimeResponseBodyResult {
  s.WorkitemId = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

