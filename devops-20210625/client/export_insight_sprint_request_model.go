// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSprintRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightSprintRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightSprintRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSprintRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightSprintRequest
  GetStartTime() *string 
}

type ExportInsightSprintRequest struct {
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

func (s ExportInsightSprintRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSprintRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightSprintRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightSprintRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSprintRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSprintRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightSprintRequest) SetEndTime(v string) *ExportInsightSprintRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightSprintRequest) SetMaxResults(v int64) *ExportInsightSprintRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSprintRequest) SetNextToken(v string) *ExportInsightSprintRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSprintRequest) SetStartTime(v string) *ExportInsightSprintRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightSprintRequest) Validate() error {
  return dara.Validate(s)
}

