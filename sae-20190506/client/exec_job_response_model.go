// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecJobResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecJobResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecJobResponse
  GetStatusCode() *int32 
  SetBody(v *ExecJobResponseBody) *ExecJobResponse
  GetBody() *ExecJobResponseBody 
}

type ExecJobResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecJobResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecJobResponse) GoString() string {
  return s.String()
}

func (s *ExecJobResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecJobResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecJobResponse) GetBody() *ExecJobResponseBody  {
  return s.Body
}

func (s *ExecJobResponse) SetHeaders(v map[string]*string) *ExecJobResponse {
  s.Headers = v
  return s
}

func (s *ExecJobResponse) SetStatusCode(v int32) *ExecJobResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecJobResponse) SetBody(v *ExecJobResponseBody) *ExecJobResponse {
  s.Body = v
  return s
}

func (s *ExecJobResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

