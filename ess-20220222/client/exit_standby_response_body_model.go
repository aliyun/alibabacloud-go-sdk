// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExitStandbyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExitStandbyResponseBody
  GetRequestId() *string 
  SetScalingActivityId(v string) *ExitStandbyResponseBody
  GetScalingActivityId() *string 
}

type ExitStandbyResponseBody struct {
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

func (s ExitStandbyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExitStandbyResponseBody) GoString() string {
  return s.String()
}

func (s *ExitStandbyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExitStandbyResponseBody) GetScalingActivityId() *string  {
  return s.ScalingActivityId
}

func (s *ExitStandbyResponseBody) SetRequestId(v string) *ExitStandbyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExitStandbyResponseBody) SetScalingActivityId(v string) *ExitStandbyResponseBody {
  s.ScalingActivityId = &v
  return s
}

func (s *ExitStandbyResponseBody) Validate() error {
  return dara.Validate(s)
}

