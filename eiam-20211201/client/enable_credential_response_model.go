// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCredentialResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCredentialResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCredentialResponseBody) *EnableCredentialResponse
  GetBody() *EnableCredentialResponseBody 
}

type EnableCredentialResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCredentialResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialResponse) GoString() string {
  return s.String()
}

func (s *EnableCredentialResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCredentialResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCredentialResponse) GetBody() *EnableCredentialResponseBody  {
  return s.Body
}

func (s *EnableCredentialResponse) SetHeaders(v map[string]*string) *EnableCredentialResponse {
  s.Headers = v
  return s
}

func (s *EnableCredentialResponse) SetStatusCode(v int32) *EnableCredentialResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCredentialResponse) SetBody(v *EnableCredentialResponseBody) *EnableCredentialResponse {
  s.Body = v
  return s
}

func (s *EnableCredentialResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

