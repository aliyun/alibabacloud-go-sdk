// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHighDefinationMonitorResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableHighDefinationMonitorResponseBody
  GetRequestId() *string 
  SetSuccess(v string) *EnableHighDefinationMonitorResponseBody
  GetSuccess() *string 
}

type EnableHighDefinationMonitorResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 2F398FF5-B349-5C01-8638-8E9A0BF1DBE6
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the call is successful. Valid values:
  // 
  // 	- **true**: yes
  // 
  // 	- **false**: no
  // 
  // example:
  // 
  // true
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHighDefinationMonitorResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHighDefinationMonitorResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHighDefinationMonitorResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHighDefinationMonitorResponseBody) GetSuccess() *string  {
  return s.Success
}

func (s *EnableHighDefinationMonitorResponseBody) SetRequestId(v string) *EnableHighDefinationMonitorResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHighDefinationMonitorResponseBody) SetSuccess(v string) *EnableHighDefinationMonitorResponseBody {
  s.Success = &v
  return s
}

func (s *EnableHighDefinationMonitorResponseBody) Validate() error {
  return dara.Validate(s)
}

