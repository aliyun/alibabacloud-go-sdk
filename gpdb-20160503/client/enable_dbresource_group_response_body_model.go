// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBResourceGroupResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDBResourceGroupResponseBody
  GetRequestId() *string 
}

type EnableDBResourceGroupResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 5850BF81-1A2B-5ACE-AF41-57**********
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBResourceGroupResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBResourceGroupResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBResourceGroupResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBResourceGroupResponseBody) SetRequestId(v string) *EnableDBResourceGroupResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBResourceGroupResponseBody) Validate() error {
  return dara.Validate(s)
}

