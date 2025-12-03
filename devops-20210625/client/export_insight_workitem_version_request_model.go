// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemVersionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightWorkitemVersionRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightWorkitemVersionRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemVersionRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightWorkitemVersionRequest
  GetStartTime() *string 
}

type ExportInsightWorkitemVersionRequest struct {
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

func (s ExportInsightWorkitemVersionRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemVersionRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemVersionRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightWorkitemVersionRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemVersionRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemVersionRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightWorkitemVersionRequest) SetEndTime(v string) *ExportInsightWorkitemVersionRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightWorkitemVersionRequest) SetMaxResults(v int64) *ExportInsightWorkitemVersionRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemVersionRequest) SetNextToken(v string) *ExportInsightWorkitemVersionRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemVersionRequest) SetStartTime(v string) *ExportInsightWorkitemVersionRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightWorkitemVersionRequest) Validate() error {
  return dara.Validate(s)
}

