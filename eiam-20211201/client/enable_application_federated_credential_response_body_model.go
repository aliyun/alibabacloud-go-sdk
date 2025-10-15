// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationFederatedCredentialResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationFederatedCredentialResponseBody
  GetRequestId() *string 
}

type EnableApplicationFederatedCredentialResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationFederatedCredentialResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationFederatedCredentialResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationFederatedCredentialResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationFederatedCredentialResponseBody) SetRequestId(v string) *EnableApplicationFederatedCredentialResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationFederatedCredentialResponseBody) Validate() error {
  return dara.Validate(s)
}

