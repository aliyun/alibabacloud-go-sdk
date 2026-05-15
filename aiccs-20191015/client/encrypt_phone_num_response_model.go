// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptPhoneNumResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EncryptPhoneNumResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EncryptPhoneNumResponse
  GetStatusCode() *int32 
  SetBody(v *EncryptPhoneNumResponseBody) *EncryptPhoneNumResponse
  GetBody() *EncryptPhoneNumResponseBody 
}

type EncryptPhoneNumResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EncryptPhoneNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptPhoneNumResponse) String() string {
  return dara.Prettify(s)
}

func (s EncryptPhoneNumResponse) GoString() string {
  return s.String()
}

func (s *EncryptPhoneNumResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EncryptPhoneNumResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EncryptPhoneNumResponse) GetBody() *EncryptPhoneNumResponseBody  {
  return s.Body
}

func (s *EncryptPhoneNumResponse) SetHeaders(v map[string]*string) *EncryptPhoneNumResponse {
  s.Headers = v
  return s
}

func (s *EncryptPhoneNumResponse) SetStatusCode(v int32) *EncryptPhoneNumResponse {
  s.StatusCode = &v
  return s
}

func (s *EncryptPhoneNumResponse) SetBody(v *EncryptPhoneNumResponseBody) *EncryptPhoneNumResponse {
  s.Body = v
  return s
}

func (s *EncryptPhoneNumResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

