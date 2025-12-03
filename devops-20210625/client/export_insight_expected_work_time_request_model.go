// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightExpectedWorkTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightExpectedWorkTimeRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightExpectedWorkTimeRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightExpectedWorkTimeRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightExpectedWorkTimeRequest
  GetStartTime() *string 
}

type ExportInsightExpectedWorkTimeRequest struct {
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

func (s ExportInsightExpectedWorkTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightExpectedWorkTimeRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightExpectedWorkTimeRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightExpectedWorkTimeRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightExpectedWorkTimeRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightExpectedWorkTimeRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightExpectedWorkTimeRequest) SetEndTime(v string) *ExportInsightExpectedWorkTimeRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeRequest) SetMaxResults(v int64) *ExportInsightExpectedWorkTimeRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeRequest) SetNextToken(v string) *ExportInsightExpectedWorkTimeRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeRequest) SetStartTime(v string) *ExportInsightExpectedWorkTimeRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeRequest) Validate() error {
  return dara.Validate(s)
}

