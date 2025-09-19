// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportVulResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportVulResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportVulResponse
  GetStatusCode() *int32 
  SetBody(v *ExportVulResponseBody) *ExportVulResponse
  GetBody() *ExportVulResponseBody 
}

type ExportVulResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportVulResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportVulResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportVulResponse) GoString() string {
  return s.String()
}

func (s *ExportVulResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportVulResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportVulResponse) GetBody() *ExportVulResponseBody  {
  return s.Body
}

func (s *ExportVulResponse) SetHeaders(v map[string]*string) *ExportVulResponse {
  s.Headers = v
  return s
}

func (s *ExportVulResponse) SetStatusCode(v int32) *ExportVulResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportVulResponse) SetBody(v *ExportVulResponseBody) *ExportVulResponse {
  s.Body = v
  return s
}

func (s *ExportVulResponse) Validate() error {
  return dara.Validate(s)
}

