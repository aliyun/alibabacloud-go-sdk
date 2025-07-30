// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationSsoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationSsoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationSsoResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationSsoResponseBody) *EnableApplicationSsoResponse
  GetBody() *EnableApplicationSsoResponseBody 
}

type EnableApplicationSsoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationSsoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationSsoResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationSsoResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationSsoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationSsoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationSsoResponse) GetBody() *EnableApplicationSsoResponseBody  {
  return s.Body
}

func (s *EnableApplicationSsoResponse) SetHeaders(v map[string]*string) *EnableApplicationSsoResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationSsoResponse) SetStatusCode(v int32) *EnableApplicationSsoResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationSsoResponse) SetBody(v *EnableApplicationSsoResponseBody) *EnableApplicationSsoResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationSsoResponse) Validate() error {
  return dara.Validate(s)
}

