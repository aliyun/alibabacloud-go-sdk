// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightWorkitemStatusRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightWorkitemStatusRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemStatusRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightWorkitemStatusRequest
  GetStartTime() *string 
}

type ExportInsightWorkitemStatusRequest struct {
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

func (s ExportInsightWorkitemStatusRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightWorkitemStatusRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemStatusRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemStatusRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightWorkitemStatusRequest) SetEndTime(v string) *ExportInsightWorkitemStatusRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusRequest) SetMaxResults(v int64) *ExportInsightWorkitemStatusRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemStatusRequest) SetNextToken(v string) *ExportInsightWorkitemStatusRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemStatusRequest) SetStartTime(v string) *ExportInsightWorkitemStatusRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusRequest) Validate() error {
  return dara.Validate(s)
}

