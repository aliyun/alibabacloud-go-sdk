// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlarmResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAlarmResponseBody
  GetRequestId() *string 
}

type EnableAlarmResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 688B18B8-FB1E-42EB-A1ED-7F55B090****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAlarmResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAlarmResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAlarmResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAlarmResponseBody) SetRequestId(v string) *EnableAlarmResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAlarmResponseBody) Validate() error {
  return dara.Validate(s)
}

