// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteQueryResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteQueryResponseBody) *ExecuteQueryResponse
  GetBody() *ExecuteQueryResponseBody 
}

type ExecuteQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteQueryResponse) GoString() string {
  return s.String()
}

func (s *ExecuteQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteQueryResponse) GetBody() *ExecuteQueryResponseBody  {
  return s.Body
}

func (s *ExecuteQueryResponse) SetHeaders(v map[string]*string) *ExecuteQueryResponse {
  s.Headers = v
  return s
}

func (s *ExecuteQueryResponse) SetStatusCode(v int32) *ExecuteQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteQueryResponse) SetBody(v *ExecuteQueryResponseBody) *ExecuteQueryResponse {
  s.Body = v
  return s
}

func (s *ExecuteQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

