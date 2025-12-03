// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeployKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDeployKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDeployKeyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDeployKeyResponseBody) *EnableDeployKeyResponse
  GetBody() *EnableDeployKeyResponseBody 
}

type EnableDeployKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDeployKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDeployKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDeployKeyResponse) GoString() string {
  return s.String()
}

func (s *EnableDeployKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDeployKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDeployKeyResponse) GetBody() *EnableDeployKeyResponseBody  {
  return s.Body
}

func (s *EnableDeployKeyResponse) SetHeaders(v map[string]*string) *EnableDeployKeyResponse {
  s.Headers = v
  return s
}

func (s *EnableDeployKeyResponse) SetStatusCode(v int32) *EnableDeployKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDeployKeyResponse) SetBody(v *EnableDeployKeyResponseBody) *EnableDeployKeyResponse {
  s.Body = v
  return s
}

func (s *EnableDeployKeyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

