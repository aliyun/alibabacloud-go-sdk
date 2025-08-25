// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnhanceImageColorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnhanceImageColorResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnhanceImageColorResponse
  GetStatusCode() *int32 
  SetBody(v *EnhanceImageColorResponseBody) *EnhanceImageColorResponse
  GetBody() *EnhanceImageColorResponseBody 
}

type EnhanceImageColorResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnhanceImageColorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnhanceImageColorResponse) String() string {
  return dara.Prettify(s)
}

func (s EnhanceImageColorResponse) GoString() string {
  return s.String()
}

func (s *EnhanceImageColorResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnhanceImageColorResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnhanceImageColorResponse) GetBody() *EnhanceImageColorResponseBody  {
  return s.Body
}

func (s *EnhanceImageColorResponse) SetHeaders(v map[string]*string) *EnhanceImageColorResponse {
  s.Headers = v
  return s
}

func (s *EnhanceImageColorResponse) SetStatusCode(v int32) *EnhanceImageColorResponse {
  s.StatusCode = &v
  return s
}

func (s *EnhanceImageColorResponse) SetBody(v *EnhanceImageColorResponseBody) *EnhanceImageColorResponse {
  s.Body = v
  return s
}

func (s *EnhanceImageColorResponse) Validate() error {
  return dara.Validate(s)
}

