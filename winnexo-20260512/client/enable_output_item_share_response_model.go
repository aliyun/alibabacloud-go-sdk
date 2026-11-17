// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableOutputItemShareResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableOutputItemShareResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableOutputItemShareResponse
  GetStatusCode() *int32 
  SetBody(v *EnableOutputItemShareResponseBody) *EnableOutputItemShareResponse
  GetBody() *EnableOutputItemShareResponseBody 
}

type EnableOutputItemShareResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableOutputItemShareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableOutputItemShareResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableOutputItemShareResponse) GoString() string {
  return s.String()
}

func (s *EnableOutputItemShareResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableOutputItemShareResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableOutputItemShareResponse) GetBody() *EnableOutputItemShareResponseBody  {
  return s.Body
}

func (s *EnableOutputItemShareResponse) SetHeaders(v map[string]*string) *EnableOutputItemShareResponse {
  s.Headers = v
  return s
}

func (s *EnableOutputItemShareResponse) SetStatusCode(v int32) *EnableOutputItemShareResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableOutputItemShareResponse) SetBody(v *EnableOutputItemShareResponseBody) *EnableOutputItemShareResponse {
  s.Body = v
  return s
}

func (s *EnableOutputItemShareResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

