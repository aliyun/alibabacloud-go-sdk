// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableExecuteStatementResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableExecuteStatementResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableExecuteStatementResponse
  GetStatusCode() *int32 
  SetBody(v *EnableExecuteStatementResponseBody) *EnableExecuteStatementResponse
  GetBody() *EnableExecuteStatementResponseBody 
}

type EnableExecuteStatementResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableExecuteStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableExecuteStatementResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableExecuteStatementResponse) GoString() string {
  return s.String()
}

func (s *EnableExecuteStatementResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableExecuteStatementResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableExecuteStatementResponse) GetBody() *EnableExecuteStatementResponseBody  {
  return s.Body
}

func (s *EnableExecuteStatementResponse) SetHeaders(v map[string]*string) *EnableExecuteStatementResponse {
  s.Headers = v
  return s
}

func (s *EnableExecuteStatementResponse) SetStatusCode(v int32) *EnableExecuteStatementResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableExecuteStatementResponse) SetBody(v *EnableExecuteStatementResponseBody) *EnableExecuteStatementResponse {
  s.Body = v
  return s
}

func (s *EnableExecuteStatementResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

