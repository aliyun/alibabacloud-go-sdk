// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMaterialFileResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportMaterialFileResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportMaterialFileResponse
  GetStatusCode() *int32 
  SetBody(v *ExportMaterialFileResponseBody) *ExportMaterialFileResponse
  GetBody() *ExportMaterialFileResponseBody 
}

type ExportMaterialFileResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportMaterialFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportMaterialFileResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportMaterialFileResponse) GoString() string {
  return s.String()
}

func (s *ExportMaterialFileResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportMaterialFileResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportMaterialFileResponse) GetBody() *ExportMaterialFileResponseBody  {
  return s.Body
}

func (s *ExportMaterialFileResponse) SetHeaders(v map[string]*string) *ExportMaterialFileResponse {
  s.Headers = v
  return s
}

func (s *ExportMaterialFileResponse) SetStatusCode(v int32) *ExportMaterialFileResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportMaterialFileResponse) SetBody(v *ExportMaterialFileResponseBody) *ExportMaterialFileResponse {
  s.Body = v
  return s
}

func (s *ExportMaterialFileResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

