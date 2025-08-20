// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdviceServiceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableAdviceServiceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableAdviceServiceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableAdviceServiceResponseBody) *EnableAdviceServiceResponse
  GetBody() *EnableAdviceServiceResponseBody 
}

type EnableAdviceServiceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableAdviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAdviceServiceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableAdviceServiceResponse) GoString() string {
  return s.String()
}

func (s *EnableAdviceServiceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableAdviceServiceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableAdviceServiceResponse) GetBody() *EnableAdviceServiceResponseBody  {
  return s.Body
}

func (s *EnableAdviceServiceResponse) SetHeaders(v map[string]*string) *EnableAdviceServiceResponse {
  s.Headers = v
  return s
}

func (s *EnableAdviceServiceResponse) SetStatusCode(v int32) *EnableAdviceServiceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableAdviceServiceResponse) SetBody(v *EnableAdviceServiceResponseBody) *EnableAdviceServiceResponse {
  s.Body = v
  return s
}

func (s *EnableAdviceServiceResponse) Validate() error {
  return dara.Validate(s)
}

