// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckReRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckReRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckReRunResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckReRunResponseBody) *ExecDataCheckReRunResponse
  GetBody() *ExecDataCheckReRunResponseBody 
}

type ExecDataCheckReRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckReRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckReRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckReRunResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckReRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckReRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckReRunResponse) GetBody() *ExecDataCheckReRunResponseBody  {
  return s.Body
}

func (s *ExecDataCheckReRunResponse) SetHeaders(v map[string]*string) *ExecDataCheckReRunResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckReRunResponse) SetStatusCode(v int32) *ExecDataCheckReRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckReRunResponse) SetBody(v *ExecDataCheckReRunResponseBody) *ExecDataCheckReRunResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckReRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

