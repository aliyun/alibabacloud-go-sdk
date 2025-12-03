// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightCustomValueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightCustomValueRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightCustomValueRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightCustomValueRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightCustomValueRequest
  GetStartTime() *string 
}

type ExportInsightCustomValueRequest struct {
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

func (s ExportInsightCustomValueRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightCustomValueRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightCustomValueRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightCustomValueRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightCustomValueRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightCustomValueRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightCustomValueRequest) SetEndTime(v string) *ExportInsightCustomValueRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightCustomValueRequest) SetMaxResults(v int64) *ExportInsightCustomValueRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightCustomValueRequest) SetNextToken(v string) *ExportInsightCustomValueRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightCustomValueRequest) SetStartTime(v string) *ExportInsightCustomValueRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightCustomValueRequest) Validate() error {
  return dara.Validate(s)
}

