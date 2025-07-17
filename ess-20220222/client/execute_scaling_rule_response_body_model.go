// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScalingRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteScalingRuleResponseBody
  GetRequestId() *string 
  SetScalingActivityId(v string) *ExecuteScalingRuleResponseBody
  GetScalingActivityId() *string 
}

type ExecuteScalingRuleResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The ID of the scaling activity.
  // 
  // example:
  // 
  // asa-bp13o672yeautiil****
  ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ExecuteScalingRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScalingRuleResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteScalingRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteScalingRuleResponseBody) GetScalingActivityId() *string  {
  return s.ScalingActivityId
}

func (s *ExecuteScalingRuleResponseBody) SetRequestId(v string) *ExecuteScalingRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteScalingRuleResponseBody) SetScalingActivityId(v string) *ExecuteScalingRuleResponseBody {
  s.ScalingActivityId = &v
  return s
}

func (s *ExecuteScalingRuleResponseBody) Validate() error {
  return dara.Validate(s)
}

