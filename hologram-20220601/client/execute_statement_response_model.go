// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStatementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteStatementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteStatementResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteStatementResponseBody) *ExecuteStatementResponse
  GetBody() *ExecuteStatementResponseBody 
}

type ExecuteStatementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteStatementResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStatementResponse) GoString() string {
  return s.String()
}

func (s *ExecuteStatementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteStatementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteStatementResponse) GetBody() *ExecuteStatementResponseBody  {
  return s.Body
}

func (s *ExecuteStatementResponse) SetHeaders(v map[string]*string) *ExecuteStatementResponse {
  s.Headers = v
  return s
}

func (s *ExecuteStatementResponse) SetStatusCode(v int32) *ExecuteStatementResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteStatementResponse) SetBody(v *ExecuteStatementResponseBody) *ExecuteStatementResponse {
  s.Body = v
  return s
}

func (s *ExecuteStatementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

