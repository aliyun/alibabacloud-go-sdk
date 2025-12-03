// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightSpaceRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightSpaceRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightSpaceRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightSpaceRequest
  GetStartTime() *string 
}

type ExportInsightSpaceRequest struct {
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

func (s ExportInsightSpaceRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightSpaceRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightSpaceRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightSpaceRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightSpaceRequest) SetEndTime(v string) *ExportInsightSpaceRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightSpaceRequest) SetMaxResults(v int64) *ExportInsightSpaceRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightSpaceRequest) SetNextToken(v string) *ExportInsightSpaceRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightSpaceRequest) SetStartTime(v string) *ExportInsightSpaceRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightSpaceRequest) Validate() error {
  return dara.Validate(s)
}

