// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPluginConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditPluginConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditPluginConfigResponse
  GetStatusCode() *int32 
  SetBody(v *EditPluginConfigResponseBody) *EditPluginConfigResponse
  GetBody() *EditPluginConfigResponseBody 
}

type EditPluginConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditPluginConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditPluginConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s EditPluginConfigResponse) GoString() string {
  return s.String()
}

func (s *EditPluginConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditPluginConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditPluginConfigResponse) GetBody() *EditPluginConfigResponseBody  {
  return s.Body
}

func (s *EditPluginConfigResponse) SetHeaders(v map[string]*string) *EditPluginConfigResponse {
  s.Headers = v
  return s
}

func (s *EditPluginConfigResponse) SetStatusCode(v int32) *EditPluginConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *EditPluginConfigResponse) SetBody(v *EditPluginConfigResponseBody) *EditPluginConfigResponse {
  s.Body = v
  return s
}

func (s *EditPluginConfigResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

