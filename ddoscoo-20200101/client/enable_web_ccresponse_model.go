// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableWebCCResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableWebCCResponse
  GetStatusCode() *int32 
  SetBody(v *EnableWebCCResponseBody) *EnableWebCCResponse
  GetBody() *EnableWebCCResponseBody 
}

type EnableWebCCResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableWebCCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableWebCCResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCResponse) GoString() string {
  return s.String()
}

func (s *EnableWebCCResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableWebCCResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableWebCCResponse) GetBody() *EnableWebCCResponseBody  {
  return s.Body
}

func (s *EnableWebCCResponse) SetHeaders(v map[string]*string) *EnableWebCCResponse {
  s.Headers = v
  return s
}

func (s *EnableWebCCResponse) SetStatusCode(v int32) *EnableWebCCResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableWebCCResponse) SetBody(v *EnableWebCCResponseBody) *EnableWebCCResponse {
  s.Body = v
  return s
}

func (s *EnableWebCCResponse) Validate() error {
  return dara.Validate(s)
}

