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
  SetResourceId(v string) *EventRedeployInstanceRequest
  GetResourceId() *string 
}

type EventRedeployInstanceRequest struct {
  // The ID of the system event.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // e-d71ff150945b9c02eb6ebc0016328468
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  // The type of the O\\&M task. Valid values:
  // 
  // 	- immediate
  // 
  // 	- scheduled
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // immediate
  OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
  // The execution time of the reservation. The timestamp is measured in milliseconds. If the OpsType parameter is set to scheduled, this parameter is required.
  // 
  // example:
  // 
  // 1742452232000
  PlanTime *int64 `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
  // The ID of the resource.
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

func (s *EventRedeployInstanceRequest) SetResourceId(v string) *EventRedeployInstanceRequest {
  s.ResourceId = &v
  return s
}

func (s *EventRedeployInstanceRequest) Validate() error {
  return dara.Validate(s)
}

