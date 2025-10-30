// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomPrivacyPolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCustomPrivacyPolicyResponseBody
  GetRequestId() *string 
}

type EnableCustomPrivacyPolicyResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCustomPrivacyPolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomPrivacyPolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCustomPrivacyPolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCustomPrivacyPolicyResponseBody) SetRequestId(v string) *EnableCustomPrivacyPolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCustomPrivacyPolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

