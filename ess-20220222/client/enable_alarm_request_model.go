// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlarmRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlarmTaskId(v string) *EnableAlarmRequest
  GetAlarmTaskId() *string 
  SetOwnerId(v int64) *EnableAlarmRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableAlarmRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableAlarmRequest
  GetResourceOwnerAccount() *string 
}

type EnableAlarmRequest struct {
  // The ID of the event-triggered task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
  AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the region.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-qingdao
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s EnableAlarmRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAlarmRequest) GoString() string {
  return s.String()
}

func (s *EnableAlarmRequest) GetAlarmTaskId() *string  {
  return s.AlarmTaskId
}

func (s *EnableAlarmRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableAlarmRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAlarmRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableAlarmRequest) SetAlarmTaskId(v string) *EnableAlarmRequest {
  s.AlarmTaskId = &v
  return s
}

func (s *EnableAlarmRequest) SetOwnerId(v int64) *EnableAlarmRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableAlarmRequest) SetRegionId(v string) *EnableAlarmRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAlarmRequest) SetResourceOwnerAccount(v string) *EnableAlarmRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableAlarmRequest) Validate() error {
  return dara.Validate(s)
}

