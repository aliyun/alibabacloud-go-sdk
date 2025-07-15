// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EkycVerifyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EkycVerifyResponse
  GetStatusCode() *int32 
  SetBody(v *EkycVerifyResponseBody) *EkycVerifyResponse
  GetBody() *EkycVerifyResponseBody 
}

type EkycVerifyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EkycVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EkycVerifyResponse) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyResponse) GoString() string {
  return s.String()
}

func (s *EkycVerifyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EkycVerifyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EkycVerifyResponse) GetBody() *EkycVerifyResponseBody  {
  return s.Body
}

func (s *EkycVerifyResponse) SetHeaders(v map[string]*string) *EkycVerifyResponse {
  s.Headers = v
  return s
}

func (s *EkycVerifyResponse) SetStatusCode(v int32) *EkycVerifyResponse {
  s.StatusCode = &v
  return s
}

func (s *EkycVerifyResponse) SetBody(v *EkycVerifyResponseBody) *EkycVerifyResponse {
  s.Body = v
  return s
}

func (s *EkycVerifyResponse) Validate() error {
  return dara.Validate(s)
}

