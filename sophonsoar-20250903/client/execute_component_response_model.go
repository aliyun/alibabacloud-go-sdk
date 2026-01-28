// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteComponentResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteComponentResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteComponentResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteComponentResponseBody) *ExecuteComponentResponse
  GetBody() *ExecuteComponentResponseBody 
}

type ExecuteComponentResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteComponentResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteComponentResponse) GoString() string {
  return s.String()
}

func (s *ExecuteComponentResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteComponentResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteComponentResponse) GetBody() *ExecuteComponentResponseBody  {
  return s.Body
}

func (s *ExecuteComponentResponse) SetHeaders(v map[string]*string) *ExecuteComponentResponse {
  s.Headers = v
  return s
}

func (s *ExecuteComponentResponse) SetStatusCode(v int32) *ExecuteComponentResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteComponentResponse) SetBody(v *ExecuteComponentResponseBody) *ExecuteComponentResponse {
  s.Body = v
  return s
}

func (s *ExecuteComponentResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

