// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInternalAuthenticationSourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableInternalAuthenticationSourceResponseBody
  GetRequestId() *string 
}

type EnableInternalAuthenticationSourceResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInternalAuthenticationSourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInternalAuthenticationSourceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInternalAuthenticationSourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInternalAuthenticationSourceResponseBody) SetRequestId(v string) *EnableInternalAuthenticationSourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInternalAuthenticationSourceResponseBody) Validate() error {
  return dara.Validate(s)
}

