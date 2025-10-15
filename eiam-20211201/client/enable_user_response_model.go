// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableUserResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableUserResponse
  GetStatusCode() *int32 
  SetBody(v *EnableUserResponseBody) *EnableUserResponse
  GetBody() *EnableUserResponseBody 
}

type EnableUserResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableUserResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableUserResponse) GoString() string {
  return s.String()
}

func (s *EnableUserResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableUserResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableUserResponse) GetBody() *EnableUserResponseBody  {
  return s.Body
}

func (s *EnableUserResponse) SetHeaders(v map[string]*string) *EnableUserResponse {
  s.Headers = v
  return s
}

func (s *EnableUserResponse) SetStatusCode(v int32) *EnableUserResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableUserResponse) SetBody(v *EnableUserResponseBody) *EnableUserResponse {
  s.Body = v
  return s
}

func (s *EnableUserResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

