// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScriptResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteScriptResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteScriptResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteScriptResponseBody) *ExecuteScriptResponse
  GetBody() *ExecuteScriptResponseBody 
}

type ExecuteScriptResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteScriptResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScriptResponse) GoString() string {
  return s.String()
}

func (s *ExecuteScriptResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteScriptResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteScriptResponse) GetBody() *ExecuteScriptResponseBody  {
  return s.Body
}

func (s *ExecuteScriptResponse) SetHeaders(v map[string]*string) *ExecuteScriptResponse {
  s.Headers = v
  return s
}

func (s *ExecuteScriptResponse) SetStatusCode(v int32) *ExecuteScriptResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteScriptResponse) SetBody(v *ExecuteScriptResponseBody) *ExecuteScriptResponse {
  s.Body = v
  return s
}

func (s *ExecuteScriptResponse) Validate() error {
  return dara.Validate(s)
}

