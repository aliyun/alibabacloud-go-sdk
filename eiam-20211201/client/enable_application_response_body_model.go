// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationResponseBody
  GetRequestId() *string 
}

type EnableApplicationResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationResponseBody) SetRequestId(v string) *EnableApplicationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationResponseBody) Validate() error {
  return dara.Validate(s)
}

