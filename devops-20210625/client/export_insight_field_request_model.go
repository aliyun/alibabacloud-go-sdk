// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightFieldRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightFieldRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightFieldRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightFieldRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightFieldRequest
  GetStartTime() *string 
}

type ExportInsightFieldRequest struct {
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

func (s ExportInsightFieldRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightFieldRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightFieldRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightFieldRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightFieldRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightFieldRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightFieldRequest) SetEndTime(v string) *ExportInsightFieldRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightFieldRequest) SetMaxResults(v int64) *ExportInsightFieldRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightFieldRequest) SetNextToken(v string) *ExportInsightFieldRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightFieldRequest) SetStartTime(v string) *ExportInsightFieldRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightFieldRequest) Validate() error {
  return dara.Validate(s)
}

