// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDataSourcesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportDataSourcesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportDataSourcesResponse
  GetStatusCode() *int32 
  SetBody(v *ExportDataSourcesResponseBody) *ExportDataSourcesResponse
  GetBody() *ExportDataSourcesResponseBody 
}

type ExportDataSourcesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportDataSourcesResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportDataSourcesResponse) GoString() string {
  return s.String()
}

func (s *ExportDataSourcesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportDataSourcesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportDataSourcesResponse) GetBody() *ExportDataSourcesResponseBody  {
  return s.Body
}

func (s *ExportDataSourcesResponse) SetHeaders(v map[string]*string) *ExportDataSourcesResponse {
  s.Headers = v
  return s
}

func (s *ExportDataSourcesResponse) SetStatusCode(v int32) *ExportDataSourcesResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportDataSourcesResponse) SetBody(v *ExportDataSourcesResponseBody) *ExportDataSourcesResponse {
  s.Body = v
  return s
}

func (s *ExportDataSourcesResponse) Validate() error {
  return dara.Validate(s)
}

