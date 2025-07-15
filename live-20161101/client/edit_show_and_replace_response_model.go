// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditShowAndReplaceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditShowAndReplaceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditShowAndReplaceResponse
  GetStatusCode() *int32 
  SetBody(v *EditShowAndReplaceResponseBody) *EditShowAndReplaceResponse
  GetBody() *EditShowAndReplaceResponseBody 
}

type EditShowAndReplaceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditShowAndReplaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditShowAndReplaceResponse) String() string {
  return dara.Prettify(s)
}

func (s EditShowAndReplaceResponse) GoString() string {
  return s.String()
}

func (s *EditShowAndReplaceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditShowAndReplaceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditShowAndReplaceResponse) GetBody() *EditShowAndReplaceResponseBody  {
  return s.Body
}

func (s *EditShowAndReplaceResponse) SetHeaders(v map[string]*string) *EditShowAndReplaceResponse {
  s.Headers = v
  return s
}

func (s *EditShowAndReplaceResponse) SetStatusCode(v int32) *EditShowAndReplaceResponse {
  s.StatusCode = &v
  return s
}

func (s *EditShowAndReplaceResponse) SetBody(v *EditShowAndReplaceResponseBody) *EditShowAndReplaceResponse {
  s.Body = v
  return s
}

func (s *EditShowAndReplaceResponse) Validate() error {
  return dara.Validate(s)
}

