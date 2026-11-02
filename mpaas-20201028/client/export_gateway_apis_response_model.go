// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGatewayApisResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportGatewayApisResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportGatewayApisResponse
  GetStatusCode() *int32 
  SetBody(v *ExportGatewayApisResponseBody) *ExportGatewayApisResponse
  GetBody() *ExportGatewayApisResponseBody 
}

type ExportGatewayApisResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportGatewayApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportGatewayApisResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportGatewayApisResponse) GoString() string {
  return s.String()
}

func (s *ExportGatewayApisResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportGatewayApisResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportGatewayApisResponse) GetBody() *ExportGatewayApisResponseBody  {
  return s.Body
}

func (s *ExportGatewayApisResponse) SetHeaders(v map[string]*string) *ExportGatewayApisResponse {
  s.Headers = v
  return s
}

func (s *ExportGatewayApisResponse) SetStatusCode(v int32) *ExportGatewayApisResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportGatewayApisResponse) SetBody(v *ExportGatewayApisResponseBody) *ExportGatewayApisResponse {
  s.Body = v
  return s
}

func (s *ExportGatewayApisResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

