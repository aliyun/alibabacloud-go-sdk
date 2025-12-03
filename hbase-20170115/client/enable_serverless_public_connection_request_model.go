// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServerlessPublicConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableServerlessPublicConnectionRequest
  GetClientToken() *string 
  SetInstanceId(v string) *EnableServerlessPublicConnectionRequest
  GetInstanceId() *string 
  SetOwnerId(v int64) *EnableServerlessPublicConnectionRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableServerlessPublicConnectionRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableServerlessPublicConnectionRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableServerlessPublicConnectionRequest
  GetResourceOwnerId() *int64 
  SetZoneId(v string) *EnableServerlessPublicConnectionRequest
  GetZoneId() *string 
}

type EnableServerlessPublicConnectionRequest struct {
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // This parameter is required.
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // This parameter is required.
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // This parameter is required.
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EnableServerlessPublicConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableServerlessPublicConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnableServerlessPublicConnectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableServerlessPublicConnectionRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableServerlessPublicConnectionRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableServerlessPublicConnectionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableServerlessPublicConnectionRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableServerlessPublicConnectionRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableServerlessPublicConnectionRequest) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableServerlessPublicConnectionRequest) SetClientToken(v string) *EnableServerlessPublicConnectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetInstanceId(v string) *EnableServerlessPublicConnectionRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetOwnerId(v int64) *EnableServerlessPublicConnectionRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetRegionId(v string) *EnableServerlessPublicConnectionRequest {
  s.RegionId = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetResourceOwnerAccount(v string) *EnableServerlessPublicConnectionRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetResourceOwnerId(v int64) *EnableServerlessPublicConnectionRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) SetZoneId(v string) *EnableServerlessPublicConnectionRequest {
  s.ZoneId = &v
  return s
}

func (s *EnableServerlessPublicConnectionRequest) Validate() error {
  return dara.Validate(s)
}

