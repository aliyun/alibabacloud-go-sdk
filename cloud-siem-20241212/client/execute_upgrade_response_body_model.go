// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteUpgradeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteUpgradeResponseBody
  GetRequestId() *string 
}

type ExecuteUpgradeResponseBody struct {
  // example:
  // 
  // 6276D891-*****-55B2-87B9-74D413F7****。
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteUpgradeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteUpgradeResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteUpgradeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteUpgradeResponseBody) SetRequestId(v string) *ExecuteUpgradeResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteUpgradeResponseBody) Validate() error {
  return dara.Validate(s)
}

