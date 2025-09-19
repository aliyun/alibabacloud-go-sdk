// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecStrategyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecStrategyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecStrategyResponse
  GetStatusCode() *int32 
  SetBody(v *ExecStrategyResponseBody) *ExecStrategyResponse
  GetBody() *ExecStrategyResponseBody 
}

type ExecStrategyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecStrategyResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecStrategyResponse) GoString() string {
  return s.String()
}

func (s *ExecStrategyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecStrategyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecStrategyResponse) GetBody() *ExecStrategyResponseBody  {
  return s.Body
}

func (s *ExecStrategyResponse) SetHeaders(v map[string]*string) *ExecStrategyResponse {
  s.Headers = v
  return s
}

func (s *ExecStrategyResponse) SetStatusCode(v int32) *ExecStrategyResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecStrategyResponse) SetBody(v *ExecStrategyResponseBody) *ExecStrategyResponse {
  s.Body = v
  return s
}

func (s *ExecStrategyResponse) Validate() error {
  return dara.Validate(s)
}

