// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityDeleteResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntityDeleteResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntityDeleteResponse
  GetStatusCode() *int32 
  SetBody(v *EntityDeleteResponseBody) *EntityDeleteResponse
  GetBody() *EntityDeleteResponseBody 
}

type EntityDeleteResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntityDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntityDeleteResponse) String() string {
  return dara.Prettify(s)
}

func (s EntityDeleteResponse) GoString() string {
  return s.String()
}

func (s *EntityDeleteResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntityDeleteResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntityDeleteResponse) GetBody() *EntityDeleteResponseBody  {
  return s.Body
}

func (s *EntityDeleteResponse) SetHeaders(v map[string]*string) *EntityDeleteResponse {
  s.Headers = v
  return s
}

func (s *EntityDeleteResponse) SetStatusCode(v int32) *EntityDeleteResponse {
  s.StatusCode = &v
  return s
}

func (s *EntityDeleteResponse) SetBody(v *EntityDeleteResponseBody) *EntityDeleteResponse {
  s.Body = v
  return s
}

func (s *EntityDeleteResponse) Validate() error {
  return dara.Validate(s)
}

