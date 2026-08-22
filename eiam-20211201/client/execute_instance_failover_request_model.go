// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteInstanceFailoverRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceFailoverStatus(v string) *ExecuteInstanceFailoverRequest
  GetInstanceFailoverStatus() *string 
  SetInstanceId(v string) *ExecuteInstanceFailoverRequest
  GetInstanceId() *string 
}

type ExecuteInstanceFailoverRequest struct {
  // The failover status. Valid values:
  // 
  // - inactive: The primary instance is active.
  // 
  // - active: The replica instance is active.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // active
  InstanceFailoverStatus *string `json:"InstanceFailoverStatus,omitempty" xml:"InstanceFailoverStatus,omitempty"`
  // The replica instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_xxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ExecuteInstanceFailoverRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteInstanceFailoverRequest) GoString() string {
  return s.String()
}

func (s *ExecuteInstanceFailoverRequest) GetInstanceFailoverStatus() *string  {
  return s.InstanceFailoverStatus
}

func (s *ExecuteInstanceFailoverRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExecuteInstanceFailoverRequest) SetInstanceFailoverStatus(v string) *ExecuteInstanceFailoverRequest {
  s.InstanceFailoverStatus = &v
  return s
}

func (s *ExecuteInstanceFailoverRequest) SetInstanceId(v string) *ExecuteInstanceFailoverRequest {
  s.InstanceId = &v
  return s
}

func (s *ExecuteInstanceFailoverRequest) Validate() error {
  return dara.Validate(s)
}

