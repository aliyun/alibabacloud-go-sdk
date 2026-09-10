// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunFailedResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckRunFailedResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckRunFailedResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckRunFailedResponseBody) *ExecDataCheckRunFailedResponse
  GetBody() *ExecDataCheckRunFailedResponseBody 
}

type ExecDataCheckRunFailedResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckRunFailedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckRunFailedResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunFailedResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunFailedResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckRunFailedResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckRunFailedResponse) GetBody() *ExecDataCheckRunFailedResponseBody  {
  return s.Body
}

func (s *ExecDataCheckRunFailedResponse) SetHeaders(v map[string]*string) *ExecDataCheckRunFailedResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckRunFailedResponse) SetStatusCode(v int32) *ExecDataCheckRunFailedResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckRunFailedResponse) SetBody(v *ExecDataCheckRunFailedResponseBody) *ExecDataCheckRunFailedResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckRunFailedResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

