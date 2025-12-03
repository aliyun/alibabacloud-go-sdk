// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse
  GetBody() *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody 
}

type ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) GetBody() *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody  {
  return s.Body
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) SetHeaders(v map[string]*string) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) SetStatusCode(v int32) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) SetBody(v *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponseBody) *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse {
  s.Body = v
  return s
}

func (s *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

