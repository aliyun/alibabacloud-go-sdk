// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckToggleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckToggleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckToggleResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckToggleResponseBody) *ExecDataCheckToggleResponse
  GetBody() *ExecDataCheckToggleResponseBody 
}

type ExecDataCheckToggleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckToggleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckToggleResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckToggleResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckToggleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckToggleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckToggleResponse) GetBody() *ExecDataCheckToggleResponseBody  {
  return s.Body
}

func (s *ExecDataCheckToggleResponse) SetHeaders(v map[string]*string) *ExecDataCheckToggleResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckToggleResponse) SetStatusCode(v int32) *ExecDataCheckToggleResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckToggleResponse) SetBody(v *ExecDataCheckToggleResponseBody) *ExecDataCheckToggleResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckToggleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

