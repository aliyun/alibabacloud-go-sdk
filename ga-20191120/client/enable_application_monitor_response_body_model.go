// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationMonitorResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationMonitorResponseBody
  GetRequestId() *string 
}

type EnableApplicationMonitorResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 04F0F334-1335-436C-A1D7-6C044FE73368
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationMonitorResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationMonitorResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationMonitorResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationMonitorResponseBody) SetRequestId(v string) *EnableApplicationMonitorResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationMonitorResponseBody) Validate() error {
  return dara.Validate(s)
}

