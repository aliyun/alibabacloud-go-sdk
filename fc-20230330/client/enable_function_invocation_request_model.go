// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFunctionInvocationRequest interface {
  dara.Model
  String() string
  GoString() string
}

type EnableFunctionInvocationRequest struct {
}

func (s EnableFunctionInvocationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableFunctionInvocationRequest) GoString() string {
  return s.String()
}

func (s *EnableFunctionInvocationRequest) Validate() error {
  return dara.Validate(s)
}

