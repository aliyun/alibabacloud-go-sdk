// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportJobsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportJobsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportJobsResponse
  GetStatusCode() *int32 
  SetBody(v []byte) *ExportJobsResponse
  GetBody() []byte 
}

type ExportJobsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body []byte `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportJobsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportJobsResponse) GoString() string {
  return s.String()
}

func (s *ExportJobsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportJobsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportJobsResponse) GetBody() []byte  {
  return s.Body
}

func (s *ExportJobsResponse) SetHeaders(v map[string]*string) *ExportJobsResponse {
  s.Headers = v
  return s
}

func (s *ExportJobsResponse) SetStatusCode(v int32) *ExportJobsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportJobsResponse) SetBody(v []byte) *ExportJobsResponse {
  s.Body = v
  return s
}

func (s *ExportJobsResponse) Validate() error {
  return dara.Validate(s)
}

