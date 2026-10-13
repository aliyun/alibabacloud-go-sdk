// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndPenetrationTestResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndPenetrationTestResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndPenetrationTestResponse
  GetStatusCode() *int32 
  SetBody(v *EndPenetrationTestResponseBody) *EndPenetrationTestResponse
  GetBody() *EndPenetrationTestResponseBody 
}

type EndPenetrationTestResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndPenetrationTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndPenetrationTestResponse) String() string {
  return dara.Prettify(s)
}

func (s EndPenetrationTestResponse) GoString() string {
  return s.String()
}

func (s *EndPenetrationTestResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndPenetrationTestResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndPenetrationTestResponse) GetBody() *EndPenetrationTestResponseBody  {
  return s.Body
}

func (s *EndPenetrationTestResponse) SetHeaders(v map[string]*string) *EndPenetrationTestResponse {
  s.Headers = v
  return s
}

func (s *EndPenetrationTestResponse) SetStatusCode(v int32) *EndPenetrationTestResponse {
  s.StatusCode = &v
  return s
}

func (s *EndPenetrationTestResponse) SetBody(v *EndPenetrationTestResponseBody) *EndPenetrationTestResponse {
  s.Body = v
  return s
}

func (s *EndPenetrationTestResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

