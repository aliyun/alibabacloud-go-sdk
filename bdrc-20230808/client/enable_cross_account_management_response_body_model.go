// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCrossAccountManagementResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCrossAccountManagementResponseBody
  GetRequestId() *string 
}

type EnableCrossAccountManagementResponseBody struct {
  // example:
  // 
  // 14DFF801-A4E3-5136-AAB8-7D246012CD7A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCrossAccountManagementResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCrossAccountManagementResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCrossAccountManagementResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCrossAccountManagementResponseBody) SetRequestId(v string) *EnableCrossAccountManagementResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCrossAccountManagementResponseBody) Validate() error {
  return dara.Validate(s)
}

