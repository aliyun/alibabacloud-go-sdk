// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalAuthenticationSourceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInternalAuthenticationSourceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInternalAuthenticationSourceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInternalAuthenticationSourceResponseBody) *EnableInternalAuthenticationSourceResponse
  GetBody() *EnableInternalAuthenticationSourceResponseBody 
}

type EnableInternalAuthenticationSourceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInternalAuthenticationSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInternalAuthenticationSourceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalAuthenticationSourceResponse) GoString() string {
  return s.String()
}

func (s *EnableInternalAuthenticationSourceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInternalAuthenticationSourceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInternalAuthenticationSourceResponse) GetBody() *EnableInternalAuthenticationSourceResponseBody  {
  return s.Body
}

func (s *EnableInternalAuthenticationSourceResponse) SetHeaders(v map[string]*string) *EnableInternalAuthenticationSourceResponse {
  s.Headers = v
  return s
}

func (s *EnableInternalAuthenticationSourceResponse) SetStatusCode(v int32) *EnableInternalAuthenticationSourceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInternalAuthenticationSourceResponse) SetBody(v *EnableInternalAuthenticationSourceResponseBody) *EnableInternalAuthenticationSourceResponse {
  s.Body = v
  return s
}

func (s *EnableInternalAuthenticationSourceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

