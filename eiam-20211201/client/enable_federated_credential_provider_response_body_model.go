// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFederatedCredentialProviderResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableFederatedCredentialProviderResponseBody
  GetRequestId() *string 
}

type EnableFederatedCredentialProviderResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableFederatedCredentialProviderResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableFederatedCredentialProviderResponseBody) GoString() string {
  return s.String()
}

func (s *EnableFederatedCredentialProviderResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableFederatedCredentialProviderResponseBody) SetRequestId(v string) *EnableFederatedCredentialProviderResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableFederatedCredentialProviderResponseBody) Validate() error {
  return dara.Validate(s)
}

