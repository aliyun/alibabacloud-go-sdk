// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportZookeeperDataResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportZookeeperDataResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportZookeeperDataResponse
  GetStatusCode() *int32 
  SetBody(v *ExportZookeeperDataResponseBody) *ExportZookeeperDataResponse
  GetBody() *ExportZookeeperDataResponseBody 
}

type ExportZookeeperDataResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportZookeeperDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportZookeeperDataResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportZookeeperDataResponse) GoString() string {
  return s.String()
}

func (s *ExportZookeeperDataResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportZookeeperDataResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportZookeeperDataResponse) GetBody() *ExportZookeeperDataResponseBody  {
  return s.Body
}

func (s *ExportZookeeperDataResponse) SetHeaders(v map[string]*string) *ExportZookeeperDataResponse {
  s.Headers = v
  return s
}

func (s *ExportZookeeperDataResponse) SetStatusCode(v int32) *ExportZookeeperDataResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportZookeeperDataResponse) SetBody(v *ExportZookeeperDataResponseBody) *ExportZookeeperDataResponse {
  s.Body = v
  return s
}

func (s *ExportZookeeperDataResponse) Validate() error {
  return dara.Validate(s)
}

