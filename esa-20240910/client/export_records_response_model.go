// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportRecordsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportRecordsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportRecordsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportRecordsResponseBody) *ExportRecordsResponse
  GetBody() *ExportRecordsResponseBody 
}

type ExportRecordsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportRecordsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportRecordsResponse) GoString() string {
  return s.String()
}

func (s *ExportRecordsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportRecordsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportRecordsResponse) GetBody() *ExportRecordsResponseBody  {
  return s.Body
}

func (s *ExportRecordsResponse) SetHeaders(v map[string]*string) *ExportRecordsResponse {
  s.Headers = v
  return s
}

func (s *ExportRecordsResponse) SetStatusCode(v int32) *ExportRecordsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportRecordsResponse) SetBody(v *ExportRecordsResponseBody) *ExportRecordsResponse {
  s.Body = v
  return s
}

func (s *ExportRecordsResponse) Validate() error {
  return dara.Validate(s)
}

