// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationResponseBody) *EnableApplicationResponse
  GetBody() *EnableApplicationResponseBody 
}

type EnableApplicationResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationResponse) GetBody() *EnableApplicationResponseBody  {
  return s.Body
}

func (s *EnableApplicationResponse) SetHeaders(v map[string]*string) *EnableApplicationResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationResponse) SetStatusCode(v int32) *EnableApplicationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationResponse) SetBody(v *EnableApplicationResponseBody) *EnableApplicationResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationResponse) Validate() error {
  return dara.Validate(s)
}

