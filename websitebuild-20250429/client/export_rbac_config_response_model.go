// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRbacConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportRbacConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportRbacConfigResponse
  GetStatusCode() *int32 
  SetBody(v *ExportRbacConfigResponseBody) *ExportRbacConfigResponse
  GetBody() *ExportRbacConfigResponseBody 
}

type ExportRbacConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportRbacConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRbacConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportRbacConfigResponse) GoString() string {
  return s.String()
}

func (s *ExportRbacConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportRbacConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportRbacConfigResponse) GetBody() *ExportRbacConfigResponseBody  {
  return s.Body
}

func (s *ExportRbacConfigResponse) SetHeaders(v map[string]*string) *ExportRbacConfigResponse {
  s.Headers = v
  return s
}

func (s *ExportRbacConfigResponse) SetStatusCode(v int32) *ExportRbacConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportRbacConfigResponse) SetBody(v *ExportRbacConfigResponseBody) *ExportRbacConfigResponse {
  s.Body = v
  return s
}

func (s *ExportRbacConfigResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

