// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPasswordResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EncryptPasswordResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EncryptPasswordResponse
  GetStatusCode() *int32 
  SetBody(v *EncryptPasswordResponseBody) *EncryptPasswordResponse
  GetBody() *EncryptPasswordResponseBody 
}

type EncryptPasswordResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EncryptPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptPasswordResponse) String() string {
  return dara.Prettify(s)
}

func (s EncryptPasswordResponse) GoString() string {
  return s.String()
}

func (s *EncryptPasswordResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EncryptPasswordResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EncryptPasswordResponse) GetBody() *EncryptPasswordResponseBody  {
  return s.Body
}

func (s *EncryptPasswordResponse) SetHeaders(v map[string]*string) *EncryptPasswordResponse {
  s.Headers = v
  return s
}

func (s *EncryptPasswordResponse) SetStatusCode(v int32) *EncryptPasswordResponse {
  s.StatusCode = &v
  return s
}

func (s *EncryptPasswordResponse) SetBody(v *EncryptPasswordResponseBody) *EncryptPasswordResponse {
  s.Body = v
  return s
}

func (s *EncryptPasswordResponse) Validate() error {
  return dara.Validate(s)
}

