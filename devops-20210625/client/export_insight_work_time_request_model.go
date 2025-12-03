// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightWorkTimeRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightWorkTimeRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkTimeRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightWorkTimeRequest
  GetStartTime() *string 
}

type ExportInsightWorkTimeRequest struct {
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

func (s ExportInsightWorkTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkTimeRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkTimeRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightWorkTimeRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkTimeRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkTimeRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightWorkTimeRequest) SetEndTime(v string) *ExportInsightWorkTimeRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightWorkTimeRequest) SetMaxResults(v int64) *ExportInsightWorkTimeRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkTimeRequest) SetNextToken(v string) *ExportInsightWorkTimeRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkTimeRequest) SetStartTime(v string) *ExportInsightWorkTimeRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightWorkTimeRequest) Validate() error {
  return dara.Validate(s)
}

