// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportScanResultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportScanResultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportScanResultResponse
  GetStatusCode() *int32 
  SetBody(v *ExportScanResultResponseBody) *ExportScanResultResponse
  GetBody() *ExportScanResultResponseBody 
}

type ExportScanResultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportScanResultResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportScanResultResponse) GoString() string {
  return s.String()
}

func (s *ExportScanResultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportScanResultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportScanResultResponse) GetBody() *ExportScanResultResponseBody  {
  return s.Body
}

func (s *ExportScanResultResponse) SetHeaders(v map[string]*string) *ExportScanResultResponse {
  s.Headers = v
  return s
}

func (s *ExportScanResultResponse) SetStatusCode(v int32) *ExportScanResultResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportScanResultResponse) SetBody(v *ExportScanResultResponseBody) *ExportScanResultResponse {
  s.Body = v
  return s
}

func (s *ExportScanResultResponse) Validate() error {
  return dara.Validate(s)
}

