// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomSourceAnalysisTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportCustomSourceAnalysisTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportCustomSourceAnalysisTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExportCustomSourceAnalysisTaskResponseBody) *ExportCustomSourceAnalysisTaskResponse
  GetBody() *ExportCustomSourceAnalysisTaskResponseBody 
}

type ExportCustomSourceAnalysisTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportCustomSourceAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCustomSourceAnalysisTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomSourceAnalysisTaskResponse) GoString() string {
  return s.String()
}

func (s *ExportCustomSourceAnalysisTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportCustomSourceAnalysisTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportCustomSourceAnalysisTaskResponse) GetBody() *ExportCustomSourceAnalysisTaskResponseBody  {
  return s.Body
}

func (s *ExportCustomSourceAnalysisTaskResponse) SetHeaders(v map[string]*string) *ExportCustomSourceAnalysisTaskResponse {
  s.Headers = v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponse) SetStatusCode(v int32) *ExportCustomSourceAnalysisTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponse) SetBody(v *ExportCustomSourceAnalysisTaskResponseBody) *ExportCustomSourceAnalysisTaskResponse {
  s.Body = v
  return s
}

func (s *ExportCustomSourceAnalysisTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

