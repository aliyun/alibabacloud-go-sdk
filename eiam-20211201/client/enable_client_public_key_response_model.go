// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableClientPublicKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableClientPublicKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableClientPublicKeyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableClientPublicKeyResponseBody) *EnableClientPublicKeyResponse
  GetBody() *EnableClientPublicKeyResponseBody 
}

type EnableClientPublicKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableClientPublicKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableClientPublicKeyResponse) GoString() string {
  return s.String()
}

func (s *EnableClientPublicKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableClientPublicKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableClientPublicKeyResponse) GetBody() *EnableClientPublicKeyResponseBody  {
  return s.Body
}

func (s *EnableClientPublicKeyResponse) SetHeaders(v map[string]*string) *EnableClientPublicKeyResponse {
  s.Headers = v
  return s
}

func (s *EnableClientPublicKeyResponse) SetStatusCode(v int32) *EnableClientPublicKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableClientPublicKeyResponse) SetBody(v *EnableClientPublicKeyResponseBody) *EnableClientPublicKeyResponse {
  s.Body = v
  return s
}

func (s *EnableClientPublicKeyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

