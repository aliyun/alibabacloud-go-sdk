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
  // Specifies whether to skip the SP (sales and billing system) billing order flow and directly enable the Express Connect circuit. Valid values:
  // 
  // - false (default): Standard flow. A billing order is automatically created when the circuit is enabled. The circuit enters the Enabled state after the order is created.
  // 
  // - true: Skips the billing order and directly enables the circuit. This capability is available only to whitelist accounts.
  // 
  // example:
  // 
  // false
  ByPassSp *bool `json:"ByPassSp,omitempty" xml:"ByPassSp,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // The client generates the value of this parameter. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
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
  // pc-2zeoaxkq3x****
  PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
  // The region where the Express Connect circuit resides.
  // 
  // You can call the DescribeRegions operation to query region IDs.
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

