// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkitemActivityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportWorkitemActivityRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportWorkitemActivityRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportWorkitemActivityRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportWorkitemActivityRequest
  GetStartTime() *string 
}

type ExportWorkitemActivityRequest struct {
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

func (s ExportWorkitemActivityRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkitemActivityRequest) GoString() string {
  return s.String()
}

func (s *ExportWorkitemActivityRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportWorkitemActivityRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportWorkitemActivityRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportWorkitemActivityRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportWorkitemActivityRequest) SetEndTime(v string) *ExportWorkitemActivityRequest {
  s.EndTime = &v
  return s
}

func (s *ExportWorkitemActivityRequest) SetMaxResults(v int64) *ExportWorkitemActivityRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportWorkitemActivityRequest) SetNextToken(v string) *ExportWorkitemActivityRequest {
  s.NextToken = &v
  return s
}

func (s *ExportWorkitemActivityRequest) SetStartTime(v string) *ExportWorkitemActivityRequest {
  s.StartTime = &v
  return s
}

func (s *ExportWorkitemActivityRequest) Validate() error {
  return dara.Validate(s)
}

