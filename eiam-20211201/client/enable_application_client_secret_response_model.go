// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationClientSecretResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationClientSecretResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationClientSecretResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationClientSecretResponseBody) *EnableApplicationClientSecretResponse
  GetBody() *EnableApplicationClientSecretResponseBody 
}

type EnableApplicationClientSecretResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationClientSecretResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationClientSecretResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationClientSecretResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationClientSecretResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationClientSecretResponse) GetBody() *EnableApplicationClientSecretResponseBody  {
  return s.Body
}

func (s *EnableApplicationClientSecretResponse) SetHeaders(v map[string]*string) *EnableApplicationClientSecretResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationClientSecretResponse) SetStatusCode(v int32) *EnableApplicationClientSecretResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationClientSecretResponse) SetBody(v *EnableApplicationClientSecretResponseBody) *EnableApplicationClientSecretResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationClientSecretResponse) Validate() error {
  return dara.Validate(s)
}

