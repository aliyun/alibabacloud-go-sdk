// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableUserResponseBody
  GetRequestId() *string 
}

type EnableUserResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableUserResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableUserResponseBody) GoString() string {
  return s.String()
}

func (s *EnableUserResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableUserResponseBody) SetRequestId(v string) *EnableUserResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableUserResponseBody) Validate() error {
  return dara.Validate(s)
}

