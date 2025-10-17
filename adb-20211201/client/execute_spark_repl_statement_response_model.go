// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSparkReplStatementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteSparkReplStatementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteSparkReplStatementResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteSparkReplStatementResponseBody) *ExecuteSparkReplStatementResponse
  GetBody() *ExecuteSparkReplStatementResponseBody 
}

type ExecuteSparkReplStatementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteSparkReplStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSparkReplStatementResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSparkReplStatementResponse) GoString() string {
  return s.String()
}

func (s *ExecuteSparkReplStatementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteSparkReplStatementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteSparkReplStatementResponse) GetBody() *ExecuteSparkReplStatementResponseBody  {
  return s.Body
}

func (s *ExecuteSparkReplStatementResponse) SetHeaders(v map[string]*string) *ExecuteSparkReplStatementResponse {
  s.Headers = v
  return s
}

func (s *ExecuteSparkReplStatementResponse) SetStatusCode(v int32) *ExecuteSparkReplStatementResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteSparkReplStatementResponse) SetBody(v *ExecuteSparkReplStatementResponseBody) *ExecuteSparkReplStatementResponse {
  s.Body = v
  return s
}

func (s *ExecuteSparkReplStatementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

