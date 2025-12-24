// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAdvancedAbilityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableIdentityProviderAdvancedAbilityResponseBody
  GetRequestId() *string 
}

type EnableIdentityProviderAdvancedAbilityResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableIdentityProviderAdvancedAbilityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAdvancedAbilityResponseBody) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAdvancedAbilityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableIdentityProviderAdvancedAbilityResponseBody) SetRequestId(v string) *EnableIdentityProviderAdvancedAbilityResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityResponseBody) Validate() error {
  return dara.Validate(s)
}

