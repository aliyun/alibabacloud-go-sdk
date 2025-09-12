// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptInvokeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EncryptInvokeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EncryptInvokeResponse
  GetStatusCode() *int32 
  SetBody(v *EncryptInvokeResponseBody) *EncryptInvokeResponse
  GetBody() *EncryptInvokeResponseBody 
}

type EncryptInvokeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EncryptInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptInvokeResponse) String() string {
  return dara.Prettify(s)
}

func (s EncryptInvokeResponse) GoString() string {
  return s.String()
}

func (s *EncryptInvokeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EncryptInvokeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EncryptInvokeResponse) GetBody() *EncryptInvokeResponseBody  {
  return s.Body
}

func (s *EncryptInvokeResponse) SetHeaders(v map[string]*string) *EncryptInvokeResponse {
  s.Headers = v
  return s
}

func (s *EncryptInvokeResponse) SetStatusCode(v int32) *EncryptInvokeResponse {
  s.StatusCode = &v
  return s
}

func (s *EncryptInvokeResponse) SetBody(v *EncryptInvokeResponseBody) *EncryptInvokeResponse {
  s.Body = v
  return s
}

func (s *EncryptInvokeResponse) Validate() error {
  return dara.Validate(s)
}

