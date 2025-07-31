// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnswerSampleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportAnswerSampleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportAnswerSampleResponse
  GetStatusCode() *int32 
  SetBody(v *ExportAnswerSampleResponseBody) *ExportAnswerSampleResponse
  GetBody() *ExportAnswerSampleResponseBody 
}

type ExportAnswerSampleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportAnswerSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAnswerSampleResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportAnswerSampleResponse) GoString() string {
  return s.String()
}

func (s *ExportAnswerSampleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportAnswerSampleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportAnswerSampleResponse) GetBody() *ExportAnswerSampleResponseBody  {
  return s.Body
}

func (s *ExportAnswerSampleResponse) SetHeaders(v map[string]*string) *ExportAnswerSampleResponse {
  s.Headers = v
  return s
}

func (s *ExportAnswerSampleResponse) SetStatusCode(v int32) *ExportAnswerSampleResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportAnswerSampleResponse) SetBody(v *ExportAnswerSampleResponseBody) *ExportAnswerSampleResponse {
  s.Body = v
  return s
}

func (s *ExportAnswerSampleResponse) Validate() error {
  return dara.Validate(s)
}

