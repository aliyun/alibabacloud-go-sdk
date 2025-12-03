// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceRefRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightSpaceRefRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightSpaceRefRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSpaceRefRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightSpaceRefRequest
  GetStartTime() *string 
}

type ExportInsightSpaceRefRequest struct {
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

func (s ExportInsightSpaceRefRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceRefRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceRefRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightSpaceRefRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSpaceRefRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSpaceRefRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightSpaceRefRequest) SetEndTime(v string) *ExportInsightSpaceRefRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightSpaceRefRequest) SetMaxResults(v int64) *ExportInsightSpaceRefRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSpaceRefRequest) SetNextToken(v string) *ExportInsightSpaceRefRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSpaceRefRequest) SetStartTime(v string) *ExportInsightSpaceRefRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightSpaceRefRequest) Validate() error {
  return dara.Validate(s)
}

