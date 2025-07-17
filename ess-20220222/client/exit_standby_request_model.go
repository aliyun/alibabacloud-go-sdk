// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExitStandbyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAsync(v bool) *ExitStandbyRequest
  GetAsync() *bool 
  SetClientToken(v string) *ExitStandbyRequest
  GetClientToken() *string 
  SetInstanceIds(v []*string) *ExitStandbyRequest
  GetInstanceIds() []*string 
  SetOwnerId(v int64) *ExitStandbyRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *ExitStandbyRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *ExitStandbyRequest
  GetResourceOwnerAccount() *string 
  SetScalingGroupId(v string) *ExitStandbyRequest
  GetScalingGroupId() *string 
}

type ExitStandbyRequest struct {
  // Specifies whether to remove the instance from the Standby state in an asynchronous manner. Valid values:
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
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-42665544****
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The IDs of the ECS instances. The value can be a JSON array that consists of up to 20 instance IDs. Separate multiple instance IDs with commas (,).
  // 
  // This parameter is required.
  InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the scaling group.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s ExitStandbyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExitStandbyRequest) GoString() string {
  return s.String()
}

func (s *ExitStandbyRequest) GetAsync() *bool  {
  return s.Async
}

func (s *ExitStandbyRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExitStandbyRequest) GetInstanceIds() []*string  {
  return s.InstanceIds
}

func (s *ExitStandbyRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExitStandbyRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExitStandbyRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ExitStandbyRequest) GetScalingGroupId() *string  {
  return s.ScalingGroupId
}

func (s *ExitStandbyRequest) SetAsync(v bool) *ExitStandbyRequest {
  s.Async = &v
  return s
}

func (s *ExitStandbyRequest) SetClientToken(v string) *ExitStandbyRequest {
  s.ClientToken = &v
  return s
}

func (s *ExitStandbyRequest) SetInstanceIds(v []*string) *ExitStandbyRequest {
  s.InstanceIds = v
  return s
}

func (s *ExitStandbyRequest) SetOwnerId(v int64) *ExitStandbyRequest {
  s.OwnerId = &v
  return s
}

func (s *ExitStandbyRequest) SetRegionId(v string) *ExitStandbyRequest {
  s.RegionId = &v
  return s
}

func (s *ExitStandbyRequest) SetResourceOwnerAccount(v string) *ExitStandbyRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ExitStandbyRequest) SetScalingGroupId(v string) *ExitStandbyRequest {
  s.ScalingGroupId = &v
  return s
}

func (s *ExitStandbyRequest) Validate() error {
  return dara.Validate(s)
}

