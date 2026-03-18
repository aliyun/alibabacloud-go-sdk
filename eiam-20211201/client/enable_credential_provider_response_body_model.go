// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialProviderResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCredentialProviderResponseBody
  GetRequestId() *string 
}

type EnableCredentialProviderResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCredentialProviderResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialProviderResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCredentialProviderResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCredentialProviderResponseBody) SetRequestId(v string) *EnableCredentialProviderResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCredentialProviderResponseBody) Validate() error {
  return dara.Validate(s)
}

