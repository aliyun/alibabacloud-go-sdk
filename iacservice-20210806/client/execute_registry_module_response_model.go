// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRegistryModuleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteRegistryModuleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteRegistryModuleResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteRegistryModuleResponseBody) *ExecuteRegistryModuleResponse
  GetBody() *ExecuteRegistryModuleResponseBody 
}

type ExecuteRegistryModuleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteRegistryModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteRegistryModuleResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRegistryModuleResponse) GoString() string {
  return s.String()
}

func (s *ExecuteRegistryModuleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteRegistryModuleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteRegistryModuleResponse) GetBody() *ExecuteRegistryModuleResponseBody  {
  return s.Body
}

func (s *ExecuteRegistryModuleResponse) SetHeaders(v map[string]*string) *ExecuteRegistryModuleResponse {
  s.Headers = v
  return s
}

func (s *ExecuteRegistryModuleResponse) SetStatusCode(v int32) *ExecuteRegistryModuleResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteRegistryModuleResponse) SetBody(v *ExecuteRegistryModuleResponseBody) *ExecuteRegistryModuleResponse {
  s.Body = v
  return s
}

func (s *ExecuteRegistryModuleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

