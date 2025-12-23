// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EncryptResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EncryptResponse
  GetStatusCode() *int32 
  SetBody(v *EncryptResponseBody) *EncryptResponse
  GetBody() *EncryptResponseBody 
}

type EncryptResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptResponse) String() string {
  return dara.Prettify(s)
}

func (s EncryptResponse) GoString() string {
  return s.String()
}

func (s *EncryptResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EncryptResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EncryptResponse) GetBody() *EncryptResponseBody  {
  return s.Body
}

func (s *EncryptResponse) SetHeaders(v map[string]*string) *EncryptResponse {
  s.Headers = v
  return s
}

func (s *EncryptResponse) SetStatusCode(v int32) *EncryptResponse {
  s.StatusCode = &v
  return s
}

func (s *EncryptResponse) SetBody(v *EncryptResponseBody) *EncryptResponse {
  s.Body = v
  return s
}

func (s *EncryptResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

