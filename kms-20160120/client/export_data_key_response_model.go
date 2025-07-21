// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportDataKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportDataKeyResponse
  GetStatusCode() *int32 
  SetBody(v *ExportDataKeyResponseBody) *ExportDataKeyResponse
  GetBody() *ExportDataKeyResponseBody 
}

type ExportDataKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportDataKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDataKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportDataKeyResponse) GoString() string {
  return s.String()
}

func (s *ExportDataKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportDataKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportDataKeyResponse) GetBody() *ExportDataKeyResponseBody  {
  return s.Body
}

func (s *ExportDataKeyResponse) SetHeaders(v map[string]*string) *ExportDataKeyResponse {
  s.Headers = v
  return s
}

func (s *ExportDataKeyResponse) SetStatusCode(v int32) *ExportDataKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportDataKeyResponse) SetBody(v *ExportDataKeyResponseBody) *ExportDataKeyResponse {
  s.Body = v
  return s
}

func (s *ExportDataKeyResponse) Validate() error {
  return dara.Validate(s)
}

