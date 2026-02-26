// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCredentialResponseBody
  GetRequestId() *string 
}

type EnableCredentialResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCredentialResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCredentialResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCredentialResponseBody) SetRequestId(v string) *EnableCredentialResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCredentialResponseBody) Validate() error {
  return dara.Validate(s)
}

