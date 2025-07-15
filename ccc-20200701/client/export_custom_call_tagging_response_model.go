// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomCallTaggingResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportCustomCallTaggingResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportCustomCallTaggingResponse
  GetStatusCode() *int32 
  SetBody(v *ExportCustomCallTaggingResponseBody) *ExportCustomCallTaggingResponse
  GetBody() *ExportCustomCallTaggingResponseBody 
}

type ExportCustomCallTaggingResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCustomCallTaggingResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomCallTaggingResponse) GoString() string {
  return s.String()
}

func (s *ExportCustomCallTaggingResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportCustomCallTaggingResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportCustomCallTaggingResponse) GetBody() *ExportCustomCallTaggingResponseBody  {
  return s.Body
}

func (s *ExportCustomCallTaggingResponse) SetHeaders(v map[string]*string) *ExportCustomCallTaggingResponse {
  s.Headers = v
  return s
}

func (s *ExportCustomCallTaggingResponse) SetStatusCode(v int32) *ExportCustomCallTaggingResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportCustomCallTaggingResponse) SetBody(v *ExportCustomCallTaggingResponseBody) *ExportCustomCallTaggingResponse {
  s.Body = v
  return s
}

func (s *ExportCustomCallTaggingResponse) Validate() error {
  return dara.Validate(s)
}

