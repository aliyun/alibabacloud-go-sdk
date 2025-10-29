// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventMigrateInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDataPolicy(v string) *EventMigrateInstanceRequest
  GetDataPolicy() *string 
  SetEventId(v string) *EventMigrateInstanceRequest
  GetEventId() *string 
  SetOpsType(v string) *EventMigrateInstanceRequest
  GetOpsType() *string 
  SetPassword(v string) *EventMigrateInstanceRequest
  GetPassword() *string 
  SetPlanTime(v int64) *EventMigrateInstanceRequest
  GetPlanTime() *int64 
  SetResourceId(v string) *EventMigrateInstanceRequest
  GetResourceId() *string 
}

type EventMigrateInstanceRequest struct {
  // The data migration policy. Valid values:
  // 
  // 	- abandon: does not migrate data. This is the default value.
  // 
  // 	- force_transfer: forcibly migrates data.
  // 
  // 	- try_transfer: Migrate data as much as possible.
  // 
  // example:
  // 
  // abandon
  DataPolicy *string `json:"DataPolicy,omitempty" xml:"DataPolicy,omitempty"`
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
  // The password of the instance. This parameter is optional. If you do not specify this parameter, a random password is used.
  // 
  // The password must be 8 to 30 characters in length. The password must contain uppercase letters, lowercase letters, digits, and special characters.
  // 
  // Note that you cannot enter a password for scheduled execution.
  // 
  // example:
  // 
  // AAaa1234
  Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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

func (s EventMigrateInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s EventMigrateInstanceRequest) GoString() string {
  return s.String()
}

func (s *EventMigrateInstanceRequest) GetDataPolicy() *string  {
  return s.DataPolicy
}

func (s *EventMigrateInstanceRequest) GetEventId() *string  {
  return s.EventId
}

func (s *EventMigrateInstanceRequest) GetOpsType() *string  {
  return s.OpsType
}

func (s *EventMigrateInstanceRequest) GetPassword() *string  {
  return s.Password
}

func (s *EventMigrateInstanceRequest) GetPlanTime() *int64  {
  return s.PlanTime
}

func (s *EventMigrateInstanceRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EventMigrateInstanceRequest) SetDataPolicy(v string) *EventMigrateInstanceRequest {
  s.DataPolicy = &v
  return s
}

func (s *EventMigrateInstanceRequest) SetEventId(v string) *EventMigrateInstanceRequest {
  s.EventId = &v
  return s
}

func (s *EventMigrateInstanceRequest) SetOpsType(v string) *EventMigrateInstanceRequest {
  s.OpsType = &v
  return s
}

func (s *EventMigrateInstanceRequest) SetPassword(v string) *EventMigrateInstanceRequest {
  s.Password = &v
  return s
}

func (s *EventMigrateInstanceRequest) SetPlanTime(v int64) *EventMigrateInstanceRequest {
  s.PlanTime = &v
  return s
}

func (s *EventMigrateInstanceRequest) SetResourceId(v string) *EventMigrateInstanceRequest {
  s.ResourceId = &v
  return s
}

func (s *EventMigrateInstanceRequest) Validate() error {
  return dara.Validate(s)
}

