// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventRebootInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEventId(v string) *EventRebootInstanceRequest
  GetEventId() *string 
  SetOpsType(v string) *EventRebootInstanceRequest
  GetOpsType() *string 
  SetPlanTime(v int64) *EventRebootInstanceRequest
  GetPlanTime() *int64 
  SetResourceId(v string) *EventRebootInstanceRequest
  GetResourceId() *string 
}

type EventRebootInstanceRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // e-9d992570200d86b9126d78aa8c37b54b
  EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // immediate
  OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
  // example:
  // 
  // 1742452232000
  PlanTime *int64 `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // i-5****
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s EventRebootInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EventRebootInstanceRequest) GoString() string {
  return s.String()
}

func (s *EventRebootInstanceRequest) GetEventId() *string  {
  return s.EventId
}

func (s *EventRebootInstanceRequest) GetOpsType() *string  {
  return s.OpsType
}

func (s *EventRebootInstanceRequest) GetPlanTime() *int64  {
  return s.PlanTime
}

func (s *EventRebootInstanceRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EventRebootInstanceRequest) SetEventId(v string) *EventRebootInstanceRequest {
  s.EventId = &v
  return s
}

func (s *EventRebootInstanceRequest) SetOpsType(v string) *EventRebootInstanceRequest {
  s.OpsType = &v
  return s
}

func (s *EventRebootInstanceRequest) SetPlanTime(v int64) *EventRebootInstanceRequest {
  s.PlanTime = &v
  return s
}

func (s *EventRebootInstanceRequest) SetResourceId(v string) *EventRebootInstanceRequest {
  s.ResourceId = &v
  return s
}

func (s *EventRebootInstanceRequest) Validate() error {
  return dara.Validate(s)
}

