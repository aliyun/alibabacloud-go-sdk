// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueBackupResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *EnableHBaseueBackupResponseBody
  GetClusterId() *string 
  SetOrderId(v string) *EnableHBaseueBackupResponseBody
  GetOrderId() *string 
  SetRequestId(v string) *EnableHBaseueBackupResponseBody
  GetRequestId() *string 
}

type EnableHBaseueBackupResponseBody struct {
  // example:
  // 
  // bds-m5e54q06ceyhxxxx
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // example:
  // 
  // 1449xxx
  OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // example:
  // 
  // 15272D5D-46E8-4400-9CC8-A7E7B589F575
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableHBaseueBackupResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueBackupResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHBaseueBackupResponseBody) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnableHBaseueBackupResponseBody) GetOrderId() *string  {
  return s.OrderId
}

func (s *EnableHBaseueBackupResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHBaseueBackupResponseBody) SetClusterId(v string) *EnableHBaseueBackupResponseBody {
  s.ClusterId = &v
  return s
}

func (s *EnableHBaseueBackupResponseBody) SetOrderId(v string) *EnableHBaseueBackupResponseBody {
  s.OrderId = &v
  return s
}

func (s *EnableHBaseueBackupResponseBody) SetRequestId(v string) *EnableHBaseueBackupResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHBaseueBackupResponseBody) Validate() error {
  return dara.Validate(s)
}

