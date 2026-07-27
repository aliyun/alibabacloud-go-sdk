// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportKgSchemaResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportKgSchemaResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportKgSchemaResponse
  GetStatusCode() *int32 
  SetBody(v *ExportKgSchemaResponseBody) *ExportKgSchemaResponse
  GetBody() *ExportKgSchemaResponseBody 
}

type ExportKgSchemaResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportKgSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportKgSchemaResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportKgSchemaResponse) GoString() string {
  return s.String()
}

func (s *ExportKgSchemaResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportKgSchemaResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportKgSchemaResponse) GetBody() *ExportKgSchemaResponseBody  {
  return s.Body
}

func (s *ExportKgSchemaResponse) SetHeaders(v map[string]*string) *ExportKgSchemaResponse {
  s.Headers = v
  return s
}

func (s *ExportKgSchemaResponse) SetStatusCode(v int32) *ExportKgSchemaResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportKgSchemaResponse) SetBody(v *ExportKgSchemaResponseBody) *ExportKgSchemaResponse {
  s.Body = v
  return s
}

func (s *ExportKgSchemaResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

