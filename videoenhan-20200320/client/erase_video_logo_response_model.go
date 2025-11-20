// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEraseVideoLogoResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EraseVideoLogoResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EraseVideoLogoResponse
  GetStatusCode() *int32 
  SetBody(v *EraseVideoLogoResponseBody) *EraseVideoLogoResponse
  GetBody() *EraseVideoLogoResponseBody 
}

type EraseVideoLogoResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EraseVideoLogoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EraseVideoLogoResponse) String() string {
  return dara.Prettify(s)
}

func (s EraseVideoLogoResponse) GoString() string {
  return s.String()
}

func (s *EraseVideoLogoResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EraseVideoLogoResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EraseVideoLogoResponse) GetBody() *EraseVideoLogoResponseBody  {
  return s.Body
}

func (s *EraseVideoLogoResponse) SetHeaders(v map[string]*string) *EraseVideoLogoResponse {
  s.Headers = v
  return s
}

func (s *EraseVideoLogoResponse) SetStatusCode(v int32) *EraseVideoLogoResponse {
  s.StatusCode = &v
  return s
}

func (s *EraseVideoLogoResponse) SetBody(v *EraseVideoLogoResponseBody) *EraseVideoLogoResponse {
  s.Body = v
  return s
}

func (s *EraseVideoLogoResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

