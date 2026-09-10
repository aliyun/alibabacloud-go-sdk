// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskPopResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditTaskPopResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditTaskPopResponse
  GetStatusCode() *int32 
  SetBody(v *EditTaskPopResponseBody) *EditTaskPopResponse
  GetBody() *EditTaskPopResponseBody 
}

type EditTaskPopResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditTaskPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditTaskPopResponse) String() string {
  return dara.Prettify(s)
}

func (s EditTaskPopResponse) GoString() string {
  return s.String()
}

func (s *EditTaskPopResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditTaskPopResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditTaskPopResponse) GetBody() *EditTaskPopResponseBody  {
  return s.Body
}

func (s *EditTaskPopResponse) SetHeaders(v map[string]*string) *EditTaskPopResponse {
  s.Headers = v
  return s
}

func (s *EditTaskPopResponse) SetStatusCode(v int32) *EditTaskPopResponse {
  s.StatusCode = &v
  return s
}

func (s *EditTaskPopResponse) SetBody(v *EditTaskPopResponseBody) *EditTaskPopResponse {
  s.Body = v
  return s
}

func (s *EditTaskPopResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

