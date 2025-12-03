// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSprintResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightSprintResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSprintResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightSprintResponseBodyResult) *ExportInsightSprintResponseBody
  GetResult() []*ExportInsightSprintResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightSprintResponseBody
  GetTotalCount() *int64 
}

type ExportInsightSprintResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightSprintResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightSprintResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSprintResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightSprintResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSprintResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSprintResponseBody) GetResult() []*ExportInsightSprintResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightSprintResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightSprintResponseBody) SetMaxResults(v int64) *ExportInsightSprintResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSprintResponseBody) SetNextToken(v string) *ExportInsightSprintResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSprintResponseBody) SetResult(v []*ExportInsightSprintResponseBodyResult) *ExportInsightSprintResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightSprintResponseBody) SetTotalCount(v int64) *ExportInsightSprintResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightSprintResponseBody) Validate() error {
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

type ExportInsightSprintResponseBodyResult struct {
  // example:
  // 
  // 1711936113000
  ActualEnd *int64 `json:"actualEnd,omitempty" xml:"actualEnd,omitempty"`
  // example:
  // 
  // 1711936113000
  ActualStart *int64 `json:"actualStart,omitempty" xml:"actualStart,omitempty"`
  // example:
  // 
  // 1710989643000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1711728000000
  GmtEnd *int64 `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
  // example:
  // 
  // 1711936113000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 1711936113000
  GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
  // example:
  // 
  // 18471761
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // e4895cadc86632f34dfaa7xxxx
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 385e7e5a4be6791f0a5185xxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 50
  Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 60.0
  WorkTimeCapacity *float64 `json:"workTimeCapacity,omitempty" xml:"workTimeCapacity,omitempty"`
}

func (s ExportInsightSprintResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSprintResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightSprintResponseBodyResult) GetActualEnd() *int64  {
  return s.ActualEnd
}

func (s *ExportInsightSprintResponseBodyResult) GetActualStart() *int64  {
  return s.ActualStart
}

func (s *ExportInsightSprintResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightSprintResponseBodyResult) GetGmtEnd() *int64  {
  return s.GmtEnd
}

func (s *ExportInsightSprintResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightSprintResponseBodyResult) GetGmtStart() *int64  {
  return s.GmtStart
}

func (s *ExportInsightSprintResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightSprintResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightSprintResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightSprintResponseBodyResult) GetName() *string  {
  return s.Name
}

func (s *ExportInsightSprintResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightSprintResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportInsightSprintResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightSprintResponseBodyResult) GetStatus() *int32  {
  return s.Status
}

func (s *ExportInsightSprintResponseBodyResult) GetWorkTimeCapacity() *float64  {
  return s.WorkTimeCapacity
}

func (s *ExportInsightSprintResponseBodyResult) SetActualEnd(v int64) *ExportInsightSprintResponseBodyResult {
  s.ActualEnd = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetActualStart(v int64) *ExportInsightSprintResponseBodyResult {
  s.ActualStart = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetGmtCreate(v int64) *ExportInsightSprintResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetGmtEnd(v int64) *ExportInsightSprintResponseBodyResult {
  s.GmtEnd = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetGmtModified(v int64) *ExportInsightSprintResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetGmtStart(v int64) *ExportInsightSprintResponseBodyResult {
  s.GmtStart = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetId(v int64) *ExportInsightSprintResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetIdentifier(v string) *ExportInsightSprintResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetIsDeleted(v string) *ExportInsightSprintResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetName(v string) *ExportInsightSprintResponseBodyResult {
  s.Name = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetOrganizationId(v string) *ExportInsightSprintResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetProjectId(v string) *ExportInsightSprintResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetSource(v string) *ExportInsightSprintResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetStatus(v int32) *ExportInsightSprintResponseBodyResult {
  s.Status = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) SetWorkTimeCapacity(v float64) *ExportInsightSprintResponseBodyResult {
  s.WorkTimeCapacity = &v
  return s
}

func (s *ExportInsightSprintResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

