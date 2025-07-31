// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportTextScanResultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportTextScanResultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportTextScanResultResponse
  GetStatusCode() *int32 
  SetBody(v *ExportTextScanResultResponseBody) *ExportTextScanResultResponse
  GetBody() *ExportTextScanResultResponseBody 
}

type ExportTextScanResultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportTextScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportTextScanResultResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportTextScanResultResponse) GoString() string {
  return s.String()
}

func (s *ExportTextScanResultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportTextScanResultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportTextScanResultResponse) GetBody() *ExportTextScanResultResponseBody  {
  return s.Body
}

func (s *ExportTextScanResultResponse) SetHeaders(v map[string]*string) *ExportTextScanResultResponse {
  s.Headers = v
  return s
}

func (s *ExportTextScanResultResponse) SetStatusCode(v int32) *ExportTextScanResultResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportTextScanResultResponse) SetBody(v *ExportTextScanResultResponseBody) *ExportTextScanResultResponse {
  s.Body = v
  return s
}

func (s *ExportTextScanResultResponse) Validate() error {
  return dara.Validate(s)
}

