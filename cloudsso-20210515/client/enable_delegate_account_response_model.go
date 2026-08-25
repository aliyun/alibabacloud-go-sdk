// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDelegateAccountResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDelegateAccountResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDelegateAccountResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDelegateAccountResponseBody) *EnableDelegateAccountResponse
  GetBody() *EnableDelegateAccountResponseBody 
}

type EnableDelegateAccountResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDelegateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDelegateAccountResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDelegateAccountResponse) GoString() string {
  return s.String()
}

func (s *EnableDelegateAccountResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDelegateAccountResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDelegateAccountResponse) GetBody() *EnableDelegateAccountResponseBody  {
  return s.Body
}

func (s *EnableDelegateAccountResponse) SetHeaders(v map[string]*string) *EnableDelegateAccountResponse {
  s.Headers = v
  return s
}

func (s *EnableDelegateAccountResponse) SetStatusCode(v int32) *EnableDelegateAccountResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDelegateAccountResponse) SetBody(v *EnableDelegateAccountResponseBody) *EnableDelegateAccountResponse {
  s.Body = v
  return s
}

func (s *EnableDelegateAccountResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

