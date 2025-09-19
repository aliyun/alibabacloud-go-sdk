// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecStrategyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecStrategyResponseBody
  GetRequestId() *string 
}

type ExecStrategyResponseBody struct {
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // 7F84EBCA-86F8-5AA0-BF74-A0276ECB****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecStrategyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecStrategyResponseBody) GoString() string {
  return s.String()
}

func (s *ExecStrategyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecStrategyResponseBody) SetRequestId(v string) *ExecStrategyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecStrategyResponseBody) Validate() error {
  return dara.Validate(s)
}

