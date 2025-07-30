// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableConditionalAccessPolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableConditionalAccessPolicyResponseBody
  GetRequestId() *string 
}

type EnableConditionalAccessPolicyResponseBody struct {
  // 请求ID。
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableConditionalAccessPolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableConditionalAccessPolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableConditionalAccessPolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableConditionalAccessPolicyResponseBody) SetRequestId(v string) *EnableConditionalAccessPolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableConditionalAccessPolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

