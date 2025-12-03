// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceRefResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightSpaceRefResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSpaceRefResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightSpaceRefResponseBodyResult) *ExportInsightSpaceRefResponseBody
  GetResult() []*ExportInsightSpaceRefResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightSpaceRefResponseBody
  GetTotalCount() *int64 
}

type ExportInsightSpaceRefResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightSpaceRefResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightSpaceRefResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceRefResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceRefResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSpaceRefResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSpaceRefResponseBody) GetResult() []*ExportInsightSpaceRefResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightSpaceRefResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightSpaceRefResponseBody) SetMaxResults(v int64) *ExportInsightSpaceRefResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBody) SetNextToken(v string) *ExportInsightSpaceRefResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBody) SetResult(v []*ExportInsightSpaceRefResponseBodyResult) *ExportInsightSpaceRefResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightSpaceRefResponseBody) SetTotalCount(v int64) *ExportInsightSpaceRefResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBody) Validate() error {
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

type ExportInsightSpaceRefResponseBodyResult struct {
  // example:
  // 
  // 65659358c319d2a0f912xxxx
  CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
  // example:
  // 
  // 7bc2be989727d0d4c9801fxxxx
  FromId *string `json:"fromId,omitempty" xml:"fromId,omitempty"`
  // example:
  // 
  // 1704267849000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1704267849000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 356525
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 49565
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 65659358c319d2a0f912xxxx
  ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 732026500a48d7a74f8b43xxxx
  ToId *string `json:"toId,omitempty" xml:"toId,omitempty"`
  // example:
  // 
  // ASSOCIATED
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExportInsightSpaceRefResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceRefResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetCreatorId() *string  {
  return s.CreatorId
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetFromId() *string  {
  return s.FromId
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetModifierId() *string  {
  return s.ModifierId
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetToId() *string  {
  return s.ToId
}

func (s *ExportInsightSpaceRefResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetCreatorId(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.CreatorId = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetFromId(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.FromId = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetGmtCreate(v int64) *ExportInsightSpaceRefResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetGmtModified(v int64) *ExportInsightSpaceRefResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetId(v int64) *ExportInsightSpaceRefResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetIdentifier(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetIsDeleted(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetModifierId(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.ModifierId = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetOrganizationId(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetToId(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.ToId = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) SetType(v string) *ExportInsightSpaceRefResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportInsightSpaceRefResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

