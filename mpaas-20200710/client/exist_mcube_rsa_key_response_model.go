// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistMcubeRsaKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExistMcubeRsaKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExistMcubeRsaKeyResponse
  GetStatusCode() *int32 
  SetBody(v *ExistMcubeRsaKeyResponseBody) *ExistMcubeRsaKeyResponse
  GetBody() *ExistMcubeRsaKeyResponseBody 
}

type ExistMcubeRsaKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExistMcubeRsaKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExistMcubeRsaKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExistMcubeRsaKeyResponse) GoString() string {
  return s.String()
}

func (s *ExistMcubeRsaKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExistMcubeRsaKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExistMcubeRsaKeyResponse) GetBody() *ExistMcubeRsaKeyResponseBody  {
  return s.Body
}

func (s *ExistMcubeRsaKeyResponse) SetHeaders(v map[string]*string) *ExistMcubeRsaKeyResponse {
  s.Headers = v
  return s
}

func (s *ExistMcubeRsaKeyResponse) SetStatusCode(v int32) *ExistMcubeRsaKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExistMcubeRsaKeyResponse) SetBody(v *ExistMcubeRsaKeyResponseBody) *ExistMcubeRsaKeyResponse {
  s.Body = v
  return s
}

func (s *ExistMcubeRsaKeyResponse) Validate() error {
  return dara.Validate(s)
}

