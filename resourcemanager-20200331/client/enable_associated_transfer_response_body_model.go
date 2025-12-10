// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAssociatedTransferResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAssociatedTransferResponseBody
  GetRequestId() *string 
}

type EnableAssociatedTransferResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 2D69A58F-345C-4FDE-88E4-BF5189484114
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAssociatedTransferResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAssociatedTransferResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAssociatedTransferResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAssociatedTransferResponseBody) SetRequestId(v string) *EnableAssociatedTransferResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAssociatedTransferResponseBody) Validate() error {
  return dara.Validate(s)
}

