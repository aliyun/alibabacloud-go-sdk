// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScheduledTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteScheduledTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteScheduledTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteScheduledTaskResponseBody) *ExecuteScheduledTaskResponse
  GetBody() *ExecuteScheduledTaskResponseBody 
}

type ExecuteScheduledTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteScheduledTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScheduledTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteScheduledTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteScheduledTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteScheduledTaskResponse) GetBody() *ExecuteScheduledTaskResponseBody  {
  return s.Body
}

func (s *ExecuteScheduledTaskResponse) SetHeaders(v map[string]*string) *ExecuteScheduledTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteScheduledTaskResponse) SetStatusCode(v int32) *ExecuteScheduledTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteScheduledTaskResponse) SetBody(v *ExecuteScheduledTaskResponseBody) *ExecuteScheduledTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteScheduledTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

