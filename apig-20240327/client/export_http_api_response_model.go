// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportHttpApiResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportHttpApiResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportHttpApiResponse
  GetStatusCode() *int32 
  SetBody(v *ExportHttpApiResponseBody) *ExportHttpApiResponse
  GetBody() *ExportHttpApiResponseBody 
}

type ExportHttpApiResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportHttpApiResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportHttpApiResponse) GoString() string {
  return s.String()
}

func (s *ExportHttpApiResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportHttpApiResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportHttpApiResponse) GetBody() *ExportHttpApiResponseBody  {
  return s.Body
}

func (s *ExportHttpApiResponse) SetHeaders(v map[string]*string) *ExportHttpApiResponse {
  s.Headers = v
  return s
}

func (s *ExportHttpApiResponse) SetStatusCode(v int32) *ExportHttpApiResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportHttpApiResponse) SetBody(v *ExportHttpApiResponseBody) *ExportHttpApiResponse {
  s.Body = v
  return s
}

func (s *ExportHttpApiResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

