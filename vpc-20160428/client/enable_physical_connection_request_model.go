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
  // Specifies whether to bypass the service provider (SP) subscription lifecycle. Valid values:
  // 
  // - **true**: Bypasses the SP subscription lifecycle. This means the instance is no longer managed by the Alibaba Cloud subscription system.
  // 
  // - **false*	- (default): Does not bypass the SP subscription lifecycle.
  // 
  // > To use this feature, contact your account manager.
  // 
  // example:
  // 
  // false
  ByPassSp *bool `json:"ByPassSp,omitempty" xml:"ByPassSp,omitempty"`
  // A client-generated token that is used to ensure the idempotence of the request.
  // 
  // The token must be unique across requests. The token can contain a maximum of 64 ASCII characters.
  // 
  // example:
  // 
  // 02fb3da4-130e-11e9-8e44-0016e04115b
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the physical connection.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-2zeoaxkq3x****
  PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
  // The ID of the region where the physical connection is located.
  // 
  // You can call `DescribeRegions` to get the latest region IDs.
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

