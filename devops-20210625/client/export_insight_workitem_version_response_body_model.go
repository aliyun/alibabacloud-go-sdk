// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemVersionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportInsightWorkitemVersionResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemVersionResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportInsightWorkitemVersionResponseBodyResult) *ExportInsightWorkitemVersionResponseBody
  GetResult() []*ExportInsightWorkitemVersionResponseBodyResult 
  SetTotalCount(v int64) *ExportInsightWorkitemVersionResponseBody
  GetTotalCount() *int64 
}

type ExportInsightWorkitemVersionResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportInsightWorkitemVersionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportInsightWorkitemVersionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemVersionResponseBody) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemVersionResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemVersionResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemVersionResponseBody) GetResult() []*ExportInsightWorkitemVersionResponseBodyResult  {
  return s.Result
}

func (s *ExportInsightWorkitemVersionResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportInsightWorkitemVersionResponseBody) SetMaxResults(v int64) *ExportInsightWorkitemVersionResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBody) SetNextToken(v string) *ExportInsightWorkitemVersionResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBody) SetResult(v []*ExportInsightWorkitemVersionResponseBodyResult) *ExportInsightWorkitemVersionResponseBody {
  s.Result = v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBody) SetTotalCount(v int64) *ExportInsightWorkitemVersionResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBody) Validate() error {
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

type ExportInsightWorkitemVersionResponseBodyResult struct {
  // example:
  // 
  // 1704251228000
  GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
  // example:
  // 
  // 1704251228000
  GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
  // example:
  // 
  // 1704902400000
  GmtPublish *int64 `json:"gmtPublish,omitempty" xml:"gmtPublish,omitempty"`
  // example:
  // 
  // 1704297600000
  GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
  // example:
  // 
  // 648131
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 7ba6e8261b973c976df76b7de1
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // N
  IsDeleted *string `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
  // example:
  // 
  // 0
  LockStatus *int32 `json:"lockStatus,omitempty" xml:"lockStatus,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // 100
  Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
  // example:
  // 
  // 7eee44ec7f699d4e6980faxxxx
  TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
  // example:
  // 
  // Project
  TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ExportInsightWorkitemVersionResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemVersionResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetGmtCreate() *int64  {
  return s.GmtCreate
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetGmtModified() *int64  {
  return s.GmtModified
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetGmtPublish() *int64  {
  return s.GmtPublish
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetGmtStart() *int64  {
  return s.GmtStart
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetIsDeleted() *string  {
  return s.IsDeleted
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetLockStatus() *int32  {
  return s.LockStatus
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetName() *string  {
  return s.Name
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetStatus() *int32  {
  return s.Status
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetTargetId() *string  {
  return s.TargetId
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) GetTargetType() *string  {
  return s.TargetType
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetGmtCreate(v int64) *ExportInsightWorkitemVersionResponseBodyResult {
  s.GmtCreate = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetGmtModified(v int64) *ExportInsightWorkitemVersionResponseBodyResult {
  s.GmtModified = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetGmtPublish(v int64) *ExportInsightWorkitemVersionResponseBodyResult {
  s.GmtPublish = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetGmtStart(v int64) *ExportInsightWorkitemVersionResponseBodyResult {
  s.GmtStart = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetId(v int64) *ExportInsightWorkitemVersionResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetIdentifier(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetIsDeleted(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.IsDeleted = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetLockStatus(v int32) *ExportInsightWorkitemVersionResponseBodyResult {
  s.LockStatus = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetName(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.Name = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetOrganizationId(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetSource(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetStatus(v int32) *ExportInsightWorkitemVersionResponseBodyResult {
  s.Status = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetTargetId(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.TargetId = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) SetTargetType(v string) *ExportInsightWorkitemVersionResponseBodyResult {
  s.TargetType = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

