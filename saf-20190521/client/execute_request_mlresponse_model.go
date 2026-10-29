// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRequestMLResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteRequestMLResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteRequestMLResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteRequestMLResponseBody) *ExecuteRequestMLResponse
  GetBody() *ExecuteRequestMLResponseBody 
}

type ExecuteRequestMLResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteRequestMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteRequestMLResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRequestMLResponse) GoString() string {
  return s.String()
}

func (s *ExecuteRequestMLResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteRequestMLResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteRequestMLResponse) GetBody() *ExecuteRequestMLResponseBody  {
  return s.Body
}

func (s *ExecuteRequestMLResponse) SetHeaders(v map[string]*string) *ExecuteRequestMLResponse {
  s.Headers = v
  return s
}

func (s *ExecuteRequestMLResponse) SetStatusCode(v int32) *ExecuteRequestMLResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteRequestMLResponse) SetBody(v *ExecuteRequestMLResponseBody) *ExecuteRequestMLResponse {
  s.Body = v
  return s
}

func (s *ExecuteRequestMLResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

