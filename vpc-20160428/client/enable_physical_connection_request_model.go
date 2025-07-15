// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePhysicalConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetByPassSp(v bool) *EnablePhysicalConnectionRequest
  GetByPassSp() *bool 
  SetClientToken(v string) *EnablePhysicalConnectionRequest
  GetClientToken() *string 
  SetOwnerAccount(v string) *EnablePhysicalConnectionRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnablePhysicalConnectionRequest
  GetOwnerId() *int64 
  SetPhysicalConnectionId(v string) *EnablePhysicalConnectionRequest
  GetPhysicalConnectionId() *string 
  SetRegionId(v string) *EnablePhysicalConnectionRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnablePhysicalConnectionRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnablePhysicalConnectionRequest
  GetResourceOwnerId() *int64 
}

type EnablePhysicalConnectionRequest struct {
  // Specifies whether to skip the order lifecycle. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false*	- (default)
  // 
  // >  To use this feature, you must contact your account manager.
  // 
  // example:
  // 
  // false
  ByPassSp *bool `json:"ByPassSp,omitempty" xml:"ByPassSp,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
  // 
  // example:
  // 
  // 02fb3da4-130e-11e9-8e44-0016e04115b
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the Express Connect circuit.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-119mfjz****
  PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
  // The region where the Express Connect circuit is deployed.
  // 
  // You can call the DescribeRegions operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnablePhysicalConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePhysicalConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnablePhysicalConnectionRequest) GetByPassSp() *bool  {
  return s.ByPassSp
}

func (s *EnablePhysicalConnectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnablePhysicalConnectionRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnablePhysicalConnectionRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnablePhysicalConnectionRequest) GetPhysicalConnectionId() *string  {
  return s.PhysicalConnectionId
}

func (s *EnablePhysicalConnectionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnablePhysicalConnectionRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnablePhysicalConnectionRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnablePhysicalConnectionRequest) SetByPassSp(v bool) *EnablePhysicalConnectionRequest {
  s.ByPassSp = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetClientToken(v string) *EnablePhysicalConnectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetOwnerAccount(v string) *EnablePhysicalConnectionRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetOwnerId(v int64) *EnablePhysicalConnectionRequest {
  s.OwnerId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetPhysicalConnectionId(v string) *EnablePhysicalConnectionRequest {
  s.PhysicalConnectionId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetRegionId(v string) *EnablePhysicalConnectionRequest {
  s.RegionId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *EnablePhysicalConnectionRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetResourceOwnerId(v int64) *EnablePhysicalConnectionRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) Validate() error {
  return dara.Validate(s)
}

