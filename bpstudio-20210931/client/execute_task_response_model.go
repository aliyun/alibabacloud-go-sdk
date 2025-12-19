// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteTaskResponseBody) *ExecuteTaskResponse
  GetBody() *ExecuteTaskResponseBody 
}

type ExecuteTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteTaskResponse) GetBody() *ExecuteTaskResponseBody  {
  return s.Body
}

func (s *ExecuteTaskResponse) SetHeaders(v map[string]*string) *ExecuteTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteTaskResponse) SetStatusCode(v int32) *ExecuteTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteTaskResponse) SetBody(v *ExecuteTaskResponseBody) *ExecuteTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

