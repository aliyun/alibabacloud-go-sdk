// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpireLoginTokenResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExpireLoginTokenResponseBody
  GetRequestId() *string 
}

type ExpireLoginTokenResponseBody struct {
  // example:
  // 
  // 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExpireLoginTokenResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExpireLoginTokenResponseBody) GoString() string {
  return s.String()
}

func (s *ExpireLoginTokenResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExpireLoginTokenResponseBody) SetRequestId(v string) *ExpireLoginTokenResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExpireLoginTokenResponseBody) Validate() error {
  return dara.Validate(s)
}

