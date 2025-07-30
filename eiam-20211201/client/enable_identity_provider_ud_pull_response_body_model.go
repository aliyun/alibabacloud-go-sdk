// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderUdPullResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableIdentityProviderUdPullResponseBody
  GetRequestId() *string 
}

type EnableIdentityProviderUdPullResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableIdentityProviderUdPullResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderUdPullResponseBody) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderUdPullResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableIdentityProviderUdPullResponseBody) SetRequestId(v string) *EnableIdentityProviderUdPullResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableIdentityProviderUdPullResponseBody) Validate() error {
  return dara.Validate(s)
}

