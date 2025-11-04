// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExitStandbyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExitStandbyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExitStandbyResponse
  GetStatusCode() *int32 
  SetBody(v *ExitStandbyResponseBody) *ExitStandbyResponse
  GetBody() *ExitStandbyResponseBody 
}

type ExitStandbyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExitStandbyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExitStandbyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExitStandbyResponse) GoString() string {
  return s.String()
}

func (s *ExitStandbyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExitStandbyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExitStandbyResponse) GetBody() *ExitStandbyResponseBody  {
  return s.Body
}

func (s *ExitStandbyResponse) SetHeaders(v map[string]*string) *ExitStandbyResponse {
  s.Headers = v
  return s
}

func (s *ExitStandbyResponse) SetStatusCode(v int32) *ExitStandbyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExitStandbyResponse) SetBody(v *ExitStandbyResponseBody) *ExitStandbyResponse {
  s.Body = v
  return s
}

func (s *ExitStandbyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

