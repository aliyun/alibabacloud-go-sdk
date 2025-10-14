// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteLogQueryResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteLogQueryResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteLogQueryResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteLogQueryResponseBody) *ExecuteLogQueryResponse
  GetBody() *ExecuteLogQueryResponseBody 
}

type ExecuteLogQueryResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteLogQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteLogQueryResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteLogQueryResponse) GoString() string {
  return s.String()
}

func (s *ExecuteLogQueryResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteLogQueryResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteLogQueryResponse) GetBody() *ExecuteLogQueryResponseBody  {
  return s.Body
}

func (s *ExecuteLogQueryResponse) SetHeaders(v map[string]*string) *ExecuteLogQueryResponse {
  s.Headers = v
  return s
}

func (s *ExecuteLogQueryResponse) SetStatusCode(v int32) *ExecuteLogQueryResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteLogQueryResponse) SetBody(v *ExecuteLogQueryResponseBody) *ExecuteLogQueryResponse {
  s.Body = v
  return s
}

func (s *ExecuteLogQueryResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

