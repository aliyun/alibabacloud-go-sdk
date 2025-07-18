// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEnterpriseAcceleratePolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableEnterpriseAcceleratePolicyResponseBody
  GetRequestId() *string 
}

type EnableEnterpriseAcceleratePolicyResponseBody struct {
  // example:
  // 
  // 09D9F396-29C5-5F0F-9C12-83308062CA2F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableEnterpriseAcceleratePolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableEnterpriseAcceleratePolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableEnterpriseAcceleratePolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableEnterpriseAcceleratePolicyResponseBody) SetRequestId(v string) *EnableEnterpriseAcceleratePolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableEnterpriseAcceleratePolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

