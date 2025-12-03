// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightTagRefRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightTagRefRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightTagRefRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightTagRefRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightTagRefRequest
  GetStartTime() *string 
}

type ExportInsightTagRefRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 2024-06-01 00:00:00
  EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
  // example:
  // 
  // 10
  MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
  // example:
  // 
  // 1
  NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2024-05-01 00:00:00
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ExportInsightTagRefRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightTagRefRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightTagRefRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightTagRefRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightTagRefRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightTagRefRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightTagRefRequest) SetEndTime(v string) *ExportInsightTagRefRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightTagRefRequest) SetMaxResults(v int64) *ExportInsightTagRefRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightTagRefRequest) SetNextToken(v string) *ExportInsightTagRefRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightTagRefRequest) SetStartTime(v string) *ExportInsightTagRefRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightTagRefRequest) Validate() error {
  return dara.Validate(s)
}

