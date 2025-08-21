// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebAccessLogConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWebAccessLogConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWebAccessLogConfigResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWebAccessLogConfigResponseBody) *EnableWebAccessLogConfigResponse
  GetBody() *EnableWebAccessLogConfigResponseBody 
}

type EnableWebAccessLogConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWebAccessLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWebAccessLogConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWebAccessLogConfigResponse) GoString() string {
  return s.String()
}

func (s *EnableWebAccessLogConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWebAccessLogConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWebAccessLogConfigResponse) GetBody() *EnableWebAccessLogConfigResponseBody  {
  return s.Body
}

func (s *EnableWebAccessLogConfigResponse) SetHeaders(v map[string]*string) *EnableWebAccessLogConfigResponse {
  s.Headers = v
  return s
}

func (s *EnableWebAccessLogConfigResponse) SetStatusCode(v int32) *EnableWebAccessLogConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWebAccessLogConfigResponse) SetBody(v *EnableWebAccessLogConfigResponseBody) *EnableWebAccessLogConfigResponse {
  s.Body = v
  return s
}

func (s *EnableWebAccessLogConfigResponse) Validate() error {
  return dara.Validate(s)
}

