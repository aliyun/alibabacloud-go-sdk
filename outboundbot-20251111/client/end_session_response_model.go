// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndSessionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndSessionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndSessionResponse
  GetStatusCode() *int32 
  SetBody(v *EndSessionResponseBody) *EndSessionResponse
  GetBody() *EndSessionResponseBody 
}

type EndSessionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndSessionResponse) String() string {
  return dara.Prettify(s)
}

func (s EndSessionResponse) GoString() string {
  return s.String()
}

func (s *EndSessionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndSessionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndSessionResponse) GetBody() *EndSessionResponseBody  {
  return s.Body
}

func (s *EndSessionResponse) SetHeaders(v map[string]*string) *EndSessionResponse {
  s.Headers = v
  return s
}

func (s *EndSessionResponse) SetStatusCode(v int32) *EndSessionResponse {
  s.StatusCode = &v
  return s
}

func (s *EndSessionResponse) SetBody(v *EndSessionResponseBody) *EndSessionResponse {
  s.Body = v
  return s
}

func (s *EndSessionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

