// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePhysicalConnectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnablePhysicalConnectionResponseBody
  GetRequestId() *string 
}

type EnablePhysicalConnectionResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 20365164-5b0d-460a-83c2-2189972b3349
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnablePhysicalConnectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnablePhysicalConnectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnablePhysicalConnectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnablePhysicalConnectionResponseBody) SetRequestId(v string) *EnablePhysicalConnectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnablePhysicalConnectionResponseBody) Validate() error {
  return dara.Validate(s)
}

