// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnalysisTagDetailByTaskIdResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportAnalysisTagDetailByTaskIdResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportAnalysisTagDetailByTaskIdResponse
  GetStatusCode() *int32 
  SetBody(v *ExportAnalysisTagDetailByTaskIdResponseBody) *ExportAnalysisTagDetailByTaskIdResponse
  GetBody() *ExportAnalysisTagDetailByTaskIdResponseBody 
}

type ExportAnalysisTagDetailByTaskIdResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportAnalysisTagDetailByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAnalysisTagDetailByTaskIdResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportAnalysisTagDetailByTaskIdResponse) GoString() string {
  return s.String()
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) GetBody() *ExportAnalysisTagDetailByTaskIdResponseBody  {
  return s.Body
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) SetHeaders(v map[string]*string) *ExportAnalysisTagDetailByTaskIdResponse {
  s.Headers = v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) SetStatusCode(v int32) *ExportAnalysisTagDetailByTaskIdResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) SetBody(v *ExportAnalysisTagDetailByTaskIdResponseBody) *ExportAnalysisTagDetailByTaskIdResponse {
  s.Body = v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

