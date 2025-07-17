// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableScalingGroupResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableScalingGroupResponseBody
  GetRequestId() *string 
}

type EnableScalingGroupResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableScalingGroupResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableScalingGroupResponseBody) GoString() string {
  return s.String()
}

func (s *EnableScalingGroupResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableScalingGroupResponseBody) SetRequestId(v string) *EnableScalingGroupResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableScalingGroupResponseBody) Validate() error {
  return dara.Validate(s)
}

