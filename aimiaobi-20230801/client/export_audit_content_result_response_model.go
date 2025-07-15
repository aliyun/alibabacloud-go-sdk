// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAuditContentResultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportAuditContentResultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportAuditContentResultResponse
  GetStatusCode() *int32 
  SetBody(v *ExportAuditContentResultResponseBody) *ExportAuditContentResultResponse
  GetBody() *ExportAuditContentResultResponseBody 
}

type ExportAuditContentResultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportAuditContentResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAuditContentResultResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportAuditContentResultResponse) GoString() string {
  return s.String()
}

func (s *ExportAuditContentResultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportAuditContentResultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportAuditContentResultResponse) GetBody() *ExportAuditContentResultResponseBody  {
  return s.Body
}

func (s *ExportAuditContentResultResponse) SetHeaders(v map[string]*string) *ExportAuditContentResultResponse {
  s.Headers = v
  return s
}

func (s *ExportAuditContentResultResponse) SetStatusCode(v int32) *ExportAuditContentResultResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportAuditContentResultResponse) SetBody(v *ExportAuditContentResultResponseBody) *ExportAuditContentResultResponse {
  s.Body = v
  return s
}

func (s *ExportAuditContentResultResponse) Validate() error {
  return dara.Validate(s)
}

