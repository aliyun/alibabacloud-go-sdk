// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCipStatsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportCipStatsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportCipStatsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportCipStatsResponseBody) *ExportCipStatsResponse
  GetBody() *ExportCipStatsResponseBody 
}

type ExportCipStatsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportCipStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCipStatsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportCipStatsResponse) GoString() string {
  return s.String()
}

func (s *ExportCipStatsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportCipStatsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportCipStatsResponse) GetBody() *ExportCipStatsResponseBody  {
  return s.Body
}

func (s *ExportCipStatsResponse) SetHeaders(v map[string]*string) *ExportCipStatsResponse {
  s.Headers = v
  return s
}

func (s *ExportCipStatsResponse) SetStatusCode(v int32) *ExportCipStatsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportCipStatsResponse) SetBody(v *ExportCipStatsResponseBody) *ExportCipStatsResponse {
  s.Body = v
  return s
}

func (s *ExportCipStatsResponse) Validate() error {
  return dara.Validate(s)
}

