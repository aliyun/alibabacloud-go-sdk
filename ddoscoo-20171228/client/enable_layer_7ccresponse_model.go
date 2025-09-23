// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableLayer7CCResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableLayer7CCResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableLayer7CCResponse
  GetStatusCode() *int32 
  SetBody(v *EnableLayer7CCResponseBody) *EnableLayer7CCResponse
  GetBody() *EnableLayer7CCResponseBody 
}

type EnableLayer7CCResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableLayer7CCResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableLayer7CCResponse) GoString() string {
  return s.String()
}

func (s *EnableLayer7CCResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableLayer7CCResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableLayer7CCResponse) GetBody() *EnableLayer7CCResponseBody  {
  return s.Body
}

func (s *EnableLayer7CCResponse) SetHeaders(v map[string]*string) *EnableLayer7CCResponse {
  s.Headers = v
  return s
}

func (s *EnableLayer7CCResponse) SetStatusCode(v int32) *EnableLayer7CCResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableLayer7CCResponse) SetBody(v *EnableLayer7CCResponseBody) *EnableLayer7CCResponse {
  s.Body = v
  return s
}

func (s *EnableLayer7CCResponse) Validate() error {
  return dara.Validate(s)
}

