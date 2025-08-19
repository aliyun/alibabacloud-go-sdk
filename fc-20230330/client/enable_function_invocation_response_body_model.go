// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFunctionInvocationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetSuccess(v bool) *EnableFunctionInvocationResponseBody
  GetSuccess() *bool 
}

type EnableFunctionInvocationResponseBody struct {
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EnableFunctionInvocationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableFunctionInvocationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableFunctionInvocationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableFunctionInvocationResponseBody) SetSuccess(v bool) *EnableFunctionInvocationResponseBody {
  s.Success = &v
  return s
}

func (s *EnableFunctionInvocationResponseBody) Validate() error {
  return dara.Validate(s)
}

