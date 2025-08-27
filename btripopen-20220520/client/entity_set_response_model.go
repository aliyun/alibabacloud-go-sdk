// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntitySetResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntitySetResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntitySetResponse
  GetStatusCode() *int32 
  SetBody(v *EntitySetResponseBody) *EntitySetResponse
  GetBody() *EntitySetResponseBody 
}

type EntitySetResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntitySetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntitySetResponse) String() string {
  return dara.Prettify(s)
}

func (s EntitySetResponse) GoString() string {
  return s.String()
}

func (s *EntitySetResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntitySetResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntitySetResponse) GetBody() *EntitySetResponseBody  {
  return s.Body
}

func (s *EntitySetResponse) SetHeaders(v map[string]*string) *EntitySetResponse {
  s.Headers = v
  return s
}

func (s *EntitySetResponse) SetStatusCode(v int32) *EntitySetResponse {
  s.StatusCode = &v
  return s
}

func (s *EntitySetResponse) SetBody(v *EntitySetResponseBody) *EntitySetResponse {
  s.Body = v
  return s
}

func (s *EntitySetResponse) Validate() error {
  return dara.Validate(s)
}

