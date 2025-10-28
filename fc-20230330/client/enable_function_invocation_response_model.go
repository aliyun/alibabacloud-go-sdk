// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFunctionInvocationResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableFunctionInvocationResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableFunctionInvocationResponse
  GetStatusCode() *int32 
  SetBody(v *EnableFunctionInvocationResponseBody) *EnableFunctionInvocationResponse
  GetBody() *EnableFunctionInvocationResponseBody 
}

type EnableFunctionInvocationResponse struct {
  Headers map[string]*string `json:"headers" xml:"headers"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableFunctionInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFunctionInvocationResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableFunctionInvocationResponse) GoString() string {
  return s.String()
}

func (s *EnableFunctionInvocationResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableFunctionInvocationResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableFunctionInvocationResponse) GetBody() *EnableFunctionInvocationResponseBody  {
  return s.Body
}

func (s *EnableFunctionInvocationResponse) SetHeaders(v map[string]*string) *EnableFunctionInvocationResponse {
  s.Headers = v
  return s
}

func (s *EnableFunctionInvocationResponse) SetStatusCode(v int32) *EnableFunctionInvocationResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableFunctionInvocationResponse) SetBody(v *EnableFunctionInvocationResponseBody) *EnableFunctionInvocationResponse {
  s.Body = v
  return s
}

func (s *EnableFunctionInvocationResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

