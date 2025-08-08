// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportMappCenterAppConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportMappCenterAppConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportMappCenterAppConfigResponse
  GetStatusCode() *int32 
  SetBody(v *ExportMappCenterAppConfigResponseBody) *ExportMappCenterAppConfigResponse
  GetBody() *ExportMappCenterAppConfigResponseBody 
}

type ExportMappCenterAppConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportMappCenterAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportMappCenterAppConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportMappCenterAppConfigResponse) GoString() string {
  return s.String()
}

func (s *ExportMappCenterAppConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportMappCenterAppConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportMappCenterAppConfigResponse) GetBody() *ExportMappCenterAppConfigResponseBody  {
  return s.Body
}

func (s *ExportMappCenterAppConfigResponse) SetHeaders(v map[string]*string) *ExportMappCenterAppConfigResponse {
  s.Headers = v
  return s
}

func (s *ExportMappCenterAppConfigResponse) SetStatusCode(v int32) *ExportMappCenterAppConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportMappCenterAppConfigResponse) SetBody(v *ExportMappCenterAppConfigResponseBody) *ExportMappCenterAppConfigResponse {
  s.Body = v
  return s
}

func (s *ExportMappCenterAppConfigResponse) Validate() error {
  return dara.Validate(s)
}

