// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRedeployInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEventId(v string) *EventRedeployInstanceRequest
  GetEventId() *string 
  SetOpsType(v string) *EventRedeployInstanceRequest
  GetOpsType() *string 
  SetPlanTime(v int64) *EventRedeployInstanceRequest
  GetPlanTime() *int64 
  SetPlanUtcTime(v string) *EventRedeployInstanceRequest
  GetPlanUtcTime() *string 
  SetResourceId(v string) *EventRedeployInstanceRequest
  GetResourceId() *string 
}

type EventRedeployInstanceRequest struct {
  // The system event ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // e-d71ff150945b9c02eb6ebc0016328468
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  // The O\\&M task type. Valid values:
  // 
  // - immediate: Immediate execution.
  // 
  // - scheduled: Scheduled execution.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // immediate
  OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
  // The scheduled execution time, specified as a timestamp in milliseconds. This parameter is required if \\"OpsType\\" is set to \\"scheduled\\".
  // 
  // example:
  // 
  // 1742452232000
  PlanTime *int64 `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
  PlanUtcTime *string `json:"PlanUtcTime,omitempty" xml:"PlanUtcTime,omitempty"`
  // The resource ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // i-55qi8m11rr53c4i964md8a00l
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s EventRedeployInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EventRedeployInstanceRequest) GoString() string {
  return s.String()
}

func (s *EventRedeployInstanceRequest) GetEventId() *string  {
  return s.EventId
}

func (s *EventRedeployInstanceRequest) GetOpsType() *string  {
  return s.OpsType
}

func (s *EventRedeployInstanceRequest) GetPlanTime() *int64  {
  return s.PlanTime
}

func (s *EventRedeployInstanceRequest) GetPlanUtcTime() *string  {
  return s.PlanUtcTime
}

func (s *EventRedeployInstanceRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EventRedeployInstanceRequest) SetEventId(v string) *EventRedeployInstanceRequest {
  s.EventId = &v
  return s
}

func (s *EventRedeployInstanceRequest) SetOpsType(v string) *EventRedeployInstanceRequest {
  s.OpsType = &v
  return s
}

func (s *EventRedeployInstanceRequest) SetPlanTime(v int64) *EventRedeployInstanceRequest {
  s.PlanTime = &v
  return s
}

func (s *EventRedeployInstanceRequest) SetPlanUtcTime(v string) *EventRedeployInstanceRequest {
  s.PlanUtcTime = &v
  return s
}

func (s *EventRedeployInstanceRequest) SetResourceId(v string) *EventRedeployInstanceRequest {
  s.ResourceId = &v
  return s
}

func (s *EventRedeployInstanceRequest) Validate() error {
  return dara.Validate(s)
}

