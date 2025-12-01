// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDatamaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDatamaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDatamaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDatamaskResponseBody) *ExecDatamaskResponse
  GetBody() *ExecDatamaskResponseBody 
}

type ExecDatamaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDatamaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDatamaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDatamaskResponse) GoString() string {
  return s.String()
}

func (s *ExecDatamaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDatamaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDatamaskResponse) GetBody() *ExecDatamaskResponseBody  {
  return s.Body
}

func (s *ExecDatamaskResponse) SetHeaders(v map[string]*string) *ExecDatamaskResponse {
  s.Headers = v
  return s
}

func (s *ExecDatamaskResponse) SetStatusCode(v int32) *ExecDatamaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDatamaskResponse) SetBody(v *ExecDatamaskResponseBody) *ExecDatamaskResponse {
  s.Body = v
  return s
}

func (s *ExecDatamaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

