// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightTagRefResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightTagRefResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightTagRefResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightTagRefResponseBodyResult) *ExportInsightTagRefResponseBody
  GetResult() []*ExportInsightTagRefResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightTagRefResponseBody
  GetTotalCount() *int64 
}

type ExportInsightTagRefResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightTagRefResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightTagRefResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightTagRefResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightTagRefResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightTagRefResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightTagRefResponseBody) GetResult() []*ExportInsightTagRefResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightTagRefResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightTagRefResponseBody) SetMaxResults(v int64) *ExportInsightTagRefResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightTagRefResponseBody) SetNextToken(v string) *ExportInsightTagRefResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightTagRefResponseBody) SetResult(v []*ExportInsightTagRefResponseBodyResult) *ExportInsightTagRefResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightTagRefResponseBody) SetTotalCount(v int64) *ExportInsightTagRefResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightTagRefResponseBody) Validate() error {
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

type ExportInsightTagRefResponseBodyResult struct {
  // example:
  // 
  // 1696660187000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1696660187000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 41317426
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 8545272
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 1
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 19e0bc5348ccbe6c0d00fbxxxx
  TagId *string `json:"tagId,omitempty" xml:"tagId,omitempty"`
  // example:
  // 
  // bde89961b5a4acc8cf54eaxxxx
  TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
  // example:
  // 
  // Workitem
  TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ExportInsightTagRefResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightTagRefResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightTagRefResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightTagRefResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightTagRefResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightTagRefResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightTagRefResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightTagRefResponseBodyResult) GetName() *string  {
  return s.Name
}

func (s *ExportInsightTagRefResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightTagRefResponseBodyResult) GetTagId() *string  {
  return s.TagId
}

func (s *ExportInsightTagRefResponseBodyResult) GetTargetId() *string  {
  return s.TargetId
}

func (s *ExportInsightTagRefResponseBodyResult) GetTargetType() *string  {
  return s.TargetType
}

func (s *ExportInsightTagRefResponseBodyResult) SetGmtCreate(v int64) *ExportInsightTagRefResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetGmtModified(v int64) *ExportInsightTagRefResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetId(v int64) *ExportInsightTagRefResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetIdentifier(v string) *ExportInsightTagRefResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetIsDeleted(v string) *ExportInsightTagRefResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetName(v string) *ExportInsightTagRefResponseBodyResult {
  s.Name = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetOrganizationId(v string) *ExportInsightTagRefResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetTagId(v string) *ExportInsightTagRefResponseBodyResult {
  s.TagId = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetTargetId(v string) *ExportInsightTagRefResponseBodyResult {
  s.TargetId = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) SetTargetType(v string) *ExportInsightTagRefResponseBodyResult {
  s.TargetType = &v
  return s
}

func (s *ExportInsightTagRefResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

