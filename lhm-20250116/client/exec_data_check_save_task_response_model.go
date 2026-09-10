// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckSaveTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckSaveTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckSaveTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckSaveTaskResponseBody) *ExecDataCheckSaveTaskResponse
  GetBody() *ExecDataCheckSaveTaskResponseBody 
}

type ExecDataCheckSaveTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckSaveTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckSaveTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckSaveTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckSaveTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckSaveTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckSaveTaskResponse) GetBody() *ExecDataCheckSaveTaskResponseBody  {
  return s.Body
}

func (s *ExecDataCheckSaveTaskResponse) SetHeaders(v map[string]*string) *ExecDataCheckSaveTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckSaveTaskResponse) SetStatusCode(v int32) *ExecDataCheckSaveTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckSaveTaskResponse) SetBody(v *ExecDataCheckSaveTaskResponseBody) *ExecDataCheckSaveTaskResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckSaveTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

