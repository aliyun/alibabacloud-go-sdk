// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKeyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableKeyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableKeyResponse
  GetStatusCode() *int32 
  SetBody(v *EnableKeyResponseBody) *EnableKeyResponse
  GetBody() *EnableKeyResponseBody 
}

type EnableKeyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableKeyResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableKeyResponse) GoString() string {
  return s.String()
}

func (s *EnableKeyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableKeyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableKeyResponse) GetBody() *EnableKeyResponseBody  {
  return s.Body
}

func (s *EnableKeyResponse) SetHeaders(v map[string]*string) *EnableKeyResponse {
  s.Headers = v
  return s
}

func (s *EnableKeyResponse) SetStatusCode(v int32) *EnableKeyResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableKeyResponse) SetBody(v *EnableKeyResponseBody) *EnableKeyResponse {
  s.Body = v
  return s
}

func (s *EnableKeyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

