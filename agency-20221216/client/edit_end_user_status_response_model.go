// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditEndUserStatusResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditEndUserStatusResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditEndUserStatusResponse
  GetStatusCode() *int32 
  SetBody(v *EditEndUserStatusResponseBody) *EditEndUserStatusResponse
  GetBody() *EditEndUserStatusResponseBody 
}

type EditEndUserStatusResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditEndUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditEndUserStatusResponse) String() string {
  return dara.Prettify(s)
}

func (s EditEndUserStatusResponse) GoString() string {
  return s.String()
}

func (s *EditEndUserStatusResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditEndUserStatusResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditEndUserStatusResponse) GetBody() *EditEndUserStatusResponseBody  {
  return s.Body
}

func (s *EditEndUserStatusResponse) SetHeaders(v map[string]*string) *EditEndUserStatusResponse {
  s.Headers = v
  return s
}

func (s *EditEndUserStatusResponse) SetStatusCode(v int32) *EditEndUserStatusResponse {
  s.StatusCode = &v
  return s
}

func (s *EditEndUserStatusResponse) SetBody(v *EditEndUserStatusResponseBody) *EditEndUserStatusResponse {
  s.Body = v
  return s
}

func (s *EditEndUserStatusResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

