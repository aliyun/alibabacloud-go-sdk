// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterStandbyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnterStandbyResponseBody
  GetRequestId() *string 
  SetScalingActivityId(v string) *EnterStandbyResponseBody
  GetScalingActivityId() *string 
}

type EnterStandbyResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The ID of the scaling activity.
  // 
  // example:
  // 
  // asa-2zeb04oym05qaceq****
  ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s EnterStandbyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterStandbyResponseBody) GoString() string {
  return s.String()
}

func (s *EnterStandbyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterStandbyResponseBody) GetScalingActivityId() *string  {
  return s.ScalingActivityId
}

func (s *EnterStandbyResponseBody) SetRequestId(v string) *EnterStandbyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterStandbyResponseBody) SetScalingActivityId(v string) *EnterStandbyResponseBody {
  s.ScalingActivityId = &v
  return s
}

func (s *EnterStandbyResponseBody) Validate() error {
  return dara.Validate(s)
}

