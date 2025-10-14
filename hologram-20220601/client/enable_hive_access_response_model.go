// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHiveAccessResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableHiveAccessResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableHiveAccessResponse
  GetStatusCode() *int32 
  SetBody(v *EnableHiveAccessResponseBody) *EnableHiveAccessResponse
  GetBody() *EnableHiveAccessResponseBody 
}

type EnableHiveAccessResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableHiveAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHiveAccessResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableHiveAccessResponse) GoString() string {
  return s.String()
}

func (s *EnableHiveAccessResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableHiveAccessResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableHiveAccessResponse) GetBody() *EnableHiveAccessResponseBody  {
  return s.Body
}

func (s *EnableHiveAccessResponse) SetHeaders(v map[string]*string) *EnableHiveAccessResponse {
  s.Headers = v
  return s
}

func (s *EnableHiveAccessResponse) SetStatusCode(v int32) *EnableHiveAccessResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableHiveAccessResponse) SetBody(v *EnableHiveAccessResponseBody) *EnableHiveAccessResponse {
  s.Body = v
  return s
}

func (s *EnableHiveAccessResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

