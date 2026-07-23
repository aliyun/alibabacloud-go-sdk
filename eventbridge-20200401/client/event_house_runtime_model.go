// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventHouseRuntime interface {
  dara.Model
  String() string
  GoString() string
  SetCu(v int32) *EventHouseRuntime
  GetCu() *int32 
  SetErrorCode(v string) *EventHouseRuntime
  GetErrorCode() *string 
  SetErrorMessage(v string) *EventHouseRuntime
  GetErrorMessage() *string 
  SetName(v string) *EventHouseRuntime
  GetName() *string 
  SetProgress(v int32) *EventHouseRuntime
  GetProgress() *int32 
  SetStage(v string) *EventHouseRuntime
  GetStage() *string 
  SetStatus(v string) *EventHouseRuntime
  GetStatus() *string 
  SetTargetCu(v int32) *EventHouseRuntime
  GetTargetCu() *int32 
}

type EventHouseRuntime struct {
  // The number of CUs that last took effect for the EventHouse Runtime.
  // 
  // example:
  // 
  // 2
  Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
  // The stable error code returned when the creation or specification change operation fails.
  // 
  // example:
  // 
  // RUNTIME_OPERATION_TIMEOUT
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The desensitized error message returned when the creation or specification change operation fails.
  // 
  // example:
  // 
  // Runtime operation timed out
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The name of the EventHouse Runtime. Typically set to default in the initial phase.
  // 
  // example:
  // 
  // default
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The progress of the current creation or specification change operation. Valid values: 0 to 100.
  // 
  // example:
  // 
  // 80
  Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
  // The current stage of the creation or specification change operation.
  // 
  // example:
  // 
  // RUNTIME_HEALTH_CHECK
  Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
  // The current status of the EventHouse Runtime. RUNNING indicates that the Runtime is ready and can accept queries. Valid values: CREATING, RUNNING, UPDATING, RECOVERING, CLOSED, CREATE_FAILED, and UPDATE_FAILED.
  // 
  // example:
  // 
  // RUNNING
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // The target number of CUs during creation, specification change, or the corresponding failed state. This parameter is not returned when the Runtime is running stably.
  // 
  // example:
  // 
  // 2
  TargetCu *int32 `json:"TargetCu,omitempty" xml:"TargetCu,omitempty"`
}

func (s EventHouseRuntime) String() string {
  return dara.Prettify(s)
}

func (s EventHouseRuntime) GoString() string {
  return s.String()
}

func (s *EventHouseRuntime) GetCu() *int32  {
  return s.Cu
}

func (s *EventHouseRuntime) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EventHouseRuntime) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EventHouseRuntime) GetName() *string  {
  return s.Name
}

func (s *EventHouseRuntime) GetProgress() *int32  {
  return s.Progress
}

func (s *EventHouseRuntime) GetStage() *string  {
  return s.Stage
}

func (s *EventHouseRuntime) GetStatus() *string  {
  return s.Status
}

func (s *EventHouseRuntime) GetTargetCu() *int32  {
  return s.TargetCu
}

func (s *EventHouseRuntime) SetCu(v int32) *EventHouseRuntime {
  s.Cu = &v
  return s
}

func (s *EventHouseRuntime) SetErrorCode(v string) *EventHouseRuntime {
  s.ErrorCode = &v
  return s
}

func (s *EventHouseRuntime) SetErrorMessage(v string) *EventHouseRuntime {
  s.ErrorMessage = &v
  return s
}

func (s *EventHouseRuntime) SetName(v string) *EventHouseRuntime {
  s.Name = &v
  return s
}

func (s *EventHouseRuntime) SetProgress(v int32) *EventHouseRuntime {
  s.Progress = &v
  return s
}

func (s *EventHouseRuntime) SetStage(v string) *EventHouseRuntime {
  s.Stage = &v
  return s
}

func (s *EventHouseRuntime) SetStatus(v string) *EventHouseRuntime {
  s.Status = &v
  return s
}

func (s *EventHouseRuntime) SetTargetCu(v int32) *EventHouseRuntime {
  s.TargetCu = &v
  return s
}

func (s *EventHouseRuntime) Validate() error {
  return dara.Validate(s)
}

