// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrollAccountResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnrollAccountResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnrollAccountResponse
  GetStatusCode() *int32 
  SetBody(v *EnrollAccountResponseBody) *EnrollAccountResponse
  GetBody() *EnrollAccountResponseBody 
}

type EnrollAccountResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnrollAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnrollAccountResponse) String() string {
  return dara.Prettify(s)
}

func (s EnrollAccountResponse) GoString() string {
  return s.String()
}

func (s *EnrollAccountResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnrollAccountResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnrollAccountResponse) GetBody() *EnrollAccountResponseBody  {
  return s.Body
}

func (s *EnrollAccountResponse) SetHeaders(v map[string]*string) *EnrollAccountResponse {
  s.Headers = v
  return s
}

func (s *EnrollAccountResponse) SetStatusCode(v int32) *EnrollAccountResponse {
  s.StatusCode = &v
  return s
}

func (s *EnrollAccountResponse) SetBody(v *EnrollAccountResponseBody) *EnrollAccountResponse {
  s.Body = v
  return s
}

func (s *EnrollAccountResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

