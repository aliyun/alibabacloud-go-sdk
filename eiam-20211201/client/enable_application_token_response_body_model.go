// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationTokenResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationTokenResponseBody
  GetRequestId() *string 
}

type EnableApplicationTokenResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationTokenResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationTokenResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationTokenResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationTokenResponseBody) SetRequestId(v string) *EnableApplicationTokenResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationTokenResponseBody) Validate() error {
  return dara.Validate(s)
}

