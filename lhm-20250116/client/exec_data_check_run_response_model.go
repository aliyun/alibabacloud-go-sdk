// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckRunResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckRunResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckRunResponseBody) *ExecDataCheckRunResponse
  GetBody() *ExecDataCheckRunResponseBody 
}

type ExecDataCheckRunResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckRunResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckRunResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckRunResponse) GetBody() *ExecDataCheckRunResponseBody  {
  return s.Body
}

func (s *ExecDataCheckRunResponse) SetHeaders(v map[string]*string) *ExecDataCheckRunResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckRunResponse) SetStatusCode(v int32) *ExecDataCheckRunResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckRunResponse) SetBody(v *ExecDataCheckRunResponseBody) *ExecDataCheckRunResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckRunResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

