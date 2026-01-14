// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditTaskResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditTaskResponse
  GetStatusCode() *int32 
  SetBody(v *EditTaskResponseBody) *EditTaskResponse
  GetBody() *EditTaskResponseBody 
}

type EditTaskResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditTaskResponse) String() string {
  return dara.Prettify(s)
}

func (s EditTaskResponse) GoString() string {
  return s.String()
}

func (s *EditTaskResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditTaskResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditTaskResponse) GetBody() *EditTaskResponseBody  {
  return s.Body
}

func (s *EditTaskResponse) SetHeaders(v map[string]*string) *EditTaskResponse {
  s.Headers = v
  return s
}

func (s *EditTaskResponse) SetStatusCode(v int32) *EditTaskResponse {
  s.StatusCode = &v
  return s
}

func (s *EditTaskResponse) SetBody(v *EditTaskResponseBody) *EditTaskResponse {
  s.Body = v
  return s
}

func (s *EditTaskResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

