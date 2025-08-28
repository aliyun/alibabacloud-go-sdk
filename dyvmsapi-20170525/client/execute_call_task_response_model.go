// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCallTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteCallTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteCallTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteCallTaskResponseBody) *ExecuteCallTaskResponse
  GetBody() *ExecuteCallTaskResponseBody 
}

type ExecuteCallTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteCallTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCallTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteCallTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteCallTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteCallTaskResponse) GetBody() *ExecuteCallTaskResponseBody  {
  return s.Body
}

func (s *ExecuteCallTaskResponse) SetHeaders(v map[string]*string) *ExecuteCallTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteCallTaskResponse) SetStatusCode(v int32) *ExecuteCallTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteCallTaskResponse) SetBody(v *ExecuteCallTaskResponseBody) *ExecuteCallTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteCallTaskResponse) Validate() error {
  return dara.Validate(s)
}

