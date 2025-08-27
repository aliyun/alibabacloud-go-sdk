// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityAddResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntityAddResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntityAddResponse
  GetStatusCode() *int32 
  SetBody(v *EntityAddResponseBody) *EntityAddResponse
  GetBody() *EntityAddResponseBody 
}

type EntityAddResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntityAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntityAddResponse) String() string {
  return dara.Prettify(s)
}

func (s EntityAddResponse) GoString() string {
  return s.String()
}

func (s *EntityAddResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntityAddResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntityAddResponse) GetBody() *EntityAddResponseBody  {
  return s.Body
}

func (s *EntityAddResponse) SetHeaders(v map[string]*string) *EntityAddResponse {
  s.Headers = v
  return s
}

func (s *EntityAddResponse) SetStatusCode(v int32) *EntityAddResponse {
  s.StatusCode = &v
  return s
}

func (s *EntityAddResponse) SetBody(v *EntityAddResponseBody) *EntityAddResponse {
  s.Body = v
  return s
}

func (s *EntityAddResponse) Validate() error {
  return dara.Validate(s)
}

