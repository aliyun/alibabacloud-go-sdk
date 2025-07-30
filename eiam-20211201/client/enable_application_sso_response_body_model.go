// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationSsoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationSsoResponseBody
  GetRequestId() *string 
}

type EnableApplicationSsoResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationSsoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationSsoResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationSsoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationSsoResponseBody) SetRequestId(v string) *EnableApplicationSsoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationSsoResponseBody) Validate() error {
  return dara.Validate(s)
}

