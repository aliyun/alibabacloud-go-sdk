// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckStopResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckStopResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckStopResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckStopResponseBody) *ExecDataCheckStopResponse
  GetBody() *ExecDataCheckStopResponseBody 
}

type ExecDataCheckStopResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckStopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckStopResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckStopResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckStopResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckStopResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckStopResponse) GetBody() *ExecDataCheckStopResponseBody  {
  return s.Body
}

func (s *ExecDataCheckStopResponse) SetHeaders(v map[string]*string) *ExecDataCheckStopResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckStopResponse) SetStatusCode(v int32) *ExecDataCheckStopResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckStopResponse) SetBody(v *ExecDataCheckStopResponseBody) *ExecDataCheckStopResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckStopResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

