// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndTime(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest
  GetEndTime() *string 
  SetMaxResults(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest
  GetMaxResults() *int64 
  SetNextToken(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest
  GetNextToken() *string 
  SetStartTime(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest
  GetStartTime() *string 
}

type ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest struct {
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

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) GetEndTime() *string  {
  return s.EndTime
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) GetMaxResults() *int64  {
  return s.MaxResults
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) GetStartTime() *string  {
  return s.StartTime
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) SetEndTime(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest {
  s.EndTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) SetMaxResults(v int64) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest {
  s.MaxResults = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) SetNextToken(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest {
  s.NextToken = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) SetStartTime(v string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest {
  s.StartTime = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest) Validate() error {
  return dara.Validate(s)
}

