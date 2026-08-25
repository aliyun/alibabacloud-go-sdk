// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableImageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableImageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableImageResponse
  GetStatusCode() *int32 
  SetBody(v *EnableImageResponseBody) *EnableImageResponse
  GetBody() *EnableImageResponseBody 
}

type EnableImageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableImageResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableImageResponse) GoString() string {
  return s.String()
}

func (s *EnableImageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableImageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableImageResponse) GetBody() *EnableImageResponseBody  {
  return s.Body
}

func (s *EnableImageResponse) SetHeaders(v map[string]*string) *EnableImageResponse {
  s.Headers = v
  return s
}

func (s *EnableImageResponse) SetStatusCode(v int32) *EnableImageResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableImageResponse) SetBody(v *EnableImageResponseBody) *EnableImageResponse {
  s.Body = v
  return s
}

func (s *EnableImageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

