// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAuthnResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableIdentityProviderAuthnResponseBody
  GetRequestId() *string 
}

type EnableIdentityProviderAuthnResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableIdentityProviderAuthnResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAuthnResponseBody) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAuthnResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableIdentityProviderAuthnResponseBody) SetRequestId(v string) *EnableIdentityProviderAuthnResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableIdentityProviderAuthnResponseBody) Validate() error {
  return dara.Validate(s)
}

