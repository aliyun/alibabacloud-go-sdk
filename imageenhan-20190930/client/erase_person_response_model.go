// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErasePersonResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ErasePersonResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ErasePersonResponse
  GetStatusCode() *int32 
  SetBody(v *ErasePersonResponseBody) *ErasePersonResponse
  GetBody() *ErasePersonResponseBody 
}

type ErasePersonResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ErasePersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ErasePersonResponse) String() string {
  return dara.Prettify(s)
}

func (s ErasePersonResponse) GoString() string {
  return s.String()
}

func (s *ErasePersonResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ErasePersonResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ErasePersonResponse) GetBody() *ErasePersonResponseBody  {
  return s.Body
}

func (s *ErasePersonResponse) SetHeaders(v map[string]*string) *ErasePersonResponse {
  s.Headers = v
  return s
}

func (s *ErasePersonResponse) SetStatusCode(v int32) *ErasePersonResponse {
  s.StatusCode = &v
  return s
}

func (s *ErasePersonResponse) SetBody(v *ErasePersonResponseBody) *ErasePersonResponse {
  s.Body = v
  return s
}

func (s *ErasePersonResponse) Validate() error {
  return dara.Validate(s)
}

