// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightSpaceResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSpaceResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightSpaceResponseBodyResult) *ExportInsightSpaceResponseBody
  GetResult() []*ExportInsightSpaceResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightSpaceResponseBody
  GetTotalCount() *int64 
}

type ExportInsightSpaceResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightSpaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightSpaceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSpaceResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSpaceResponseBody) GetResult() []*ExportInsightSpaceResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightSpaceResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightSpaceResponseBody) SetMaxResults(v int64) *ExportInsightSpaceResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSpaceResponseBody) SetNextToken(v string) *ExportInsightSpaceResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSpaceResponseBody) SetResult(v []*ExportInsightSpaceResponseBodyResult) *ExportInsightSpaceResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightSpaceResponseBody) SetTotalCount(v int64) *ExportInsightSpaceResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightSpaceResponseBody) Validate() error {
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

type ExportInsightSpaceResponseBodyResult struct {
  // example:
  // 
  // Project
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // example:
  // 
  // 1706510424000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1706511201000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 11034222
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 83a2861bbb43b270a04b42xxxx
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // 61db9af2148974246be6xxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 10
  Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // Project
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExportInsightSpaceResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceResponseBodyResult) GetCategory() *string  {
  return s.Category
}

func (s *ExportInsightSpaceResponseBodyResult) GetCustomCode() *string  {
  return s.CustomCode
}

func (s *ExportInsightSpaceResponseBodyResult) GetDescription() *string  {
  return s.Description
}

func (s *ExportInsightSpaceResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightSpaceResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightSpaceResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightSpaceResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightSpaceResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightSpaceResponseBodyResult) GetName() *string  {
  return s.Name
}

func (s *ExportInsightSpaceResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightSpaceResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightSpaceResponseBodyResult) GetStage() *string  {
  return s.Stage
}

func (s *ExportInsightSpaceResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *ExportInsightSpaceResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightSpaceResponseBodyResult) SetCategory(v string) *ExportInsightSpaceResponseBodyResult {
  s.Category = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetCustomCode(v string) *ExportInsightSpaceResponseBodyResult {
  s.CustomCode = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetDescription(v string) *ExportInsightSpaceResponseBodyResult {
  s.Description = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetGmtCreate(v int64) *ExportInsightSpaceResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetGmtModified(v int64) *ExportInsightSpaceResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetId(v int64) *ExportInsightSpaceResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetIdentifier(v string) *ExportInsightSpaceResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetIsDeleted(v string) *ExportInsightSpaceResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetName(v string) *ExportInsightSpaceResponseBodyResult {
  s.Name = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetOrganizationId(v string) *ExportInsightSpaceResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetSource(v string) *ExportInsightSpaceResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetStage(v string) *ExportInsightSpaceResponseBodyResult {
  s.Stage = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetStatus(v string) *ExportInsightSpaceResponseBodyResult {
  s.Status = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) SetType(v string) *ExportInsightSpaceResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightSpaceResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

