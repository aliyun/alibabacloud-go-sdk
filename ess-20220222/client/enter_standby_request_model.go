// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterStandbyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAsync(v bool) *EnterStandbyRequest
  GetAsync() *bool 
  SetClientToken(v string) *EnterStandbyRequest
  GetClientToken() *string 
  SetInstanceIds(v []*string) *EnterStandbyRequest
  GetInstanceIds() []*string 
  SetOwnerId(v int64) *EnterStandbyRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnterStandbyRequest
  GetResourceOwnerAccount() *string 
  SetScalingGroupId(v string) *EnterStandbyRequest
  GetScalingGroupId() *string 
}

type EnterStandbyRequest struct {
  // Specifies whether to asynchronously put the ECS instance into the Standby state. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // Default value: false.
  // 
  // example:
  // 
  // false
  Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
  // The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-42665544****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The IDs of the ECS instances.
  // 
  // This parameter is required.
  InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  // The ID of the scaling group.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // asg-bp1fo0dbtsbmqa9h****
  ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s EnterStandbyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterStandbyRequest) GoString() string {
  return s.String()
}

func (s *EnterStandbyRequest) GetAsync() *bool  {
  return s.Async
}

func (s *EnterStandbyRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnterStandbyRequest) GetInstanceIds() []*string  {
  return s.InstanceIds
}

func (s *EnterStandbyRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnterStandbyRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnterStandbyRequest) GetScalingGroupId() *string  {
  return s.ScalingGroupId
}

func (s *EnterStandbyRequest) SetAsync(v bool) *EnterStandbyRequest {
  s.Async = &v
  return s
}

func (s *EnterStandbyRequest) SetClientToken(v string) *EnterStandbyRequest {
  s.ClientToken = &v
  return s
}

func (s *EnterStandbyRequest) SetInstanceIds(v []*string) *EnterStandbyRequest {
  s.InstanceIds = v
  return s
}

func (s *EnterStandbyRequest) SetOwnerId(v int64) *EnterStandbyRequest {
  s.OwnerId = &v
  return s
}

func (s *EnterStandbyRequest) SetResourceOwnerAccount(v string) *EnterStandbyRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnterStandbyRequest) SetScalingGroupId(v string) *EnterStandbyRequest {
  s.ScalingGroupId = &v
  return s
}

func (s *EnterStandbyRequest) Validate() error {
  return dara.Validate(s)
}

