// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkitemActivityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMaxResults(v int64) *ExportWorkitemActivityResponseBody
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportWorkitemActivityResponseBody
  GetNextToken() *string 
  SetResult(v []*ExportWorkitemActivityResponseBodyResult) *ExportWorkitemActivityResponseBody
  GetResult() []*ExportWorkitemActivityResponseBodyResult 
  SetTotalCount(v int64) *ExportWorkitemActivityResponseBody
  GetTotalCount() *int64 
}

type ExportWorkitemActivityResponseBody struct {
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 2
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  Result []*ExportWorkitemActivityResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
  // example:
  // 
  // 100
  TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ExportWorkitemActivityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkitemActivityResponseBody) GoString() string {
  return s.String()
}

func (s *ExportWorkitemActivityResponseBody) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportWorkitemActivityResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportWorkitemActivityResponseBody) GetResult() []*ExportWorkitemActivityResponseBodyResult  {
  return s.Result
}

func (s *ExportWorkitemActivityResponseBody) GetTotalCount() *int64  {
  return s.TotalCount
}

func (s *ExportWorkitemActivityResponseBody) SetMaxResults(v int64) *ExportWorkitemActivityResponseBody {
  s.MaxResults = &v
  return s
}

func (s *ExportWorkitemActivityResponseBody) SetNextToken(v string) *ExportWorkitemActivityResponseBody {
  s.NextToken = &v
  return s
}

func (s *ExportWorkitemActivityResponseBody) SetResult(v []*ExportWorkitemActivityResponseBodyResult) *ExportWorkitemActivityResponseBody {
  s.Result = v
  return s
}

func (s *ExportWorkitemActivityResponseBody) SetTotalCount(v int64) *ExportWorkitemActivityResponseBody {
  s.TotalCount = &v
  return s
}

func (s *ExportWorkitemActivityResponseBody) Validate() error {
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

type ExportWorkitemActivityResponseBodyResult struct {
  // example:
  // 
  // 1714961337000
  GmtEvent *int64 `json:"gmtEvent,omitempty" xml:"gmtEvent,omitempty"`
  // example:
  // 
  // 4406380356
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 254662353
  Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
  // example:
  // 
  // 1
  NewValue *string `json:"newValue,omitempty" xml:"newValue,omitempty"`
  // example:
  // 
  // 2
  OldValue *string `json:"oldValue,omitempty" xml:"oldValue,omitempty"`
  // example:
  // 
  // 61db9af2148974246bexxxx
  OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
  // example:
  // 
  // 2a62349afcbef7f23d8f31xxxx
  ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
  // example:
  // 
  // projex
  Source *string `json:"source,omitempty" xml:"source,omitempty"`
  // example:
  // 
  // workitem.update.priority
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // example:
  // 
  // ec69eae498acce08ff7260xxxx
  WorkitemId *string `json:"workitemId,omitempty" xml:"workitemId,omitempty"`
}

func (s ExportWorkitemActivityResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkitemActivityResponseBodyResult) GoString() string {
  return s.String()
}

func (s *ExportWorkitemActivityResponseBodyResult) GetGmtEvent() *int64  {
  return s.GmtEvent
}

func (s *ExportWorkitemActivityResponseBodyResult) GetId() *int64  {
  return s.Id
}

func (s *ExportWorkitemActivityResponseBodyResult) GetIdentifier() *string  {
  return s.Identifier
}

func (s *ExportWorkitemActivityResponseBodyResult) GetNewValue() *string  {
  return s.NewValue
}

func (s *ExportWorkitemActivityResponseBodyResult) GetOldValue() *string  {
  return s.OldValue
}

func (s *ExportWorkitemActivityResponseBodyResult) GetOrganizationId() *string  {
  return s.OrganizationId
}

func (s *ExportWorkitemActivityResponseBodyResult) GetProjectId() *string  {
  return s.ProjectId
}

func (s *ExportWorkitemActivityResponseBodyResult) GetSource() *string  {
  return s.Source
}

func (s *ExportWorkitemActivityResponseBodyResult) GetType() *string  {
  return s.Type
}

func (s *ExportWorkitemActivityResponseBodyResult) GetWorkitemId() *string  {
  return s.WorkitemId
}

func (s *ExportWorkitemActivityResponseBodyResult) SetGmtEvent(v int64) *ExportWorkitemActivityResponseBodyResult {
  s.GmtEvent = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetId(v int64) *ExportWorkitemActivityResponseBodyResult {
  s.Id = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetIdentifier(v string) *ExportWorkitemActivityResponseBodyResult {
  s.Identifier = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetNewValue(v string) *ExportWorkitemActivityResponseBodyResult {
  s.NewValue = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetOldValue(v string) *ExportWorkitemActivityResponseBodyResult {
  s.OldValue = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetOrganizationId(v string) *ExportWorkitemActivityResponseBodyResult {
  s.OrganizationId = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetProjectId(v string) *ExportWorkitemActivityResponseBodyResult {
  s.ProjectId = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetSource(v string) *ExportWorkitemActivityResponseBodyResult {
  s.Source = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetType(v string) *ExportWorkitemActivityResponseBodyResult {
  s.Type = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) SetWorkitemId(v string) *ExportWorkitemActivityResponseBodyResult {
  s.WorkitemId = &v
  return s
}

func (s *ExportWorkitemActivityResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

