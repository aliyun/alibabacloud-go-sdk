// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdHocTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteAdHocTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteAdHocTaskResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteAdHocTaskResponseBody) *ExecuteAdHocTaskResponse
  GetBody() *ExecuteAdHocTaskResponseBody 
}

type ExecuteAdHocTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteAdHocTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAdHocTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdHocTaskResponse) GoString() string {
  return s.String()
}

func (s *ExecuteAdHocTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteAdHocTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteAdHocTaskResponse) GetBody() *ExecuteAdHocTaskResponseBody  {
  return s.Body
}

func (s *ExecuteAdHocTaskResponse) SetHeaders(v map[string]*string) *ExecuteAdHocTaskResponse {
  s.Headers = v
  return s
}

func (s *ExecuteAdHocTaskResponse) SetStatusCode(v int32) *ExecuteAdHocTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteAdHocTaskResponse) SetBody(v *ExecuteAdHocTaskResponseBody) *ExecuteAdHocTaskResponse {
  s.Body = v
  return s
}

func (s *ExecuteAdHocTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

