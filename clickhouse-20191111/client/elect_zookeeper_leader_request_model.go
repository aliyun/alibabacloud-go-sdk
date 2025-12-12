// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElectZookeeperLeaderRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *ElectZookeeperLeaderRequest
  GetDBClusterId() *string 
  SetDisableWrite(v bool) *ElectZookeeperLeaderRequest
  GetDisableWrite() *bool 
  SetElectTime(v string) *ElectZookeeperLeaderRequest
  GetElectTime() *string 
  SetOwnerAccount(v string) *ElectZookeeperLeaderRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *ElectZookeeperLeaderRequest
  GetOwnerId() *int64 
  SetPageNumber(v int32) *ElectZookeeperLeaderRequest
  GetPageNumber() *int32 
  SetPageSize(v int32) *ElectZookeeperLeaderRequest
  GetPageSize() *int32 
  SetRegionId(v string) *ElectZookeeperLeaderRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *ElectZookeeperLeaderRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *ElectZookeeperLeaderRequest
  GetResourceOwnerId() *int64 
}

type ElectZookeeperLeaderRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // cc-bp108z124a8o7****
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // example:
  // 
  // true
  DisableWrite *bool `json:"DisableWrite,omitempty" xml:"DisableWrite,omitempty"`
  // example:
  // 
  // 2025-12-12T07:28:00Z
  ElectTime *string `json:"ElectTime,omitempty" xml:"ElectTime,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // example:
  // 
  // 1
  PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
  // example:
  // 
  // 30
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ElectZookeeperLeaderRequest) String() string {
  return dara.Prettify(s)
}

func (s ElectZookeeperLeaderRequest) GoString() string {
  return s.String()
}

func (s *ElectZookeeperLeaderRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *ElectZookeeperLeaderRequest) GetDisableWrite() *bool  {
  return s.DisableWrite
}

func (s *ElectZookeeperLeaderRequest) GetElectTime() *string  {
  return s.ElectTime
}

func (s *ElectZookeeperLeaderRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *ElectZookeeperLeaderRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ElectZookeeperLeaderRequest) GetPageNumber() *int32  {
  return s.PageNumber
}

func (s *ElectZookeeperLeaderRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *ElectZookeeperLeaderRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ElectZookeeperLeaderRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ElectZookeeperLeaderRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *ElectZookeeperLeaderRequest) SetDBClusterId(v string) *ElectZookeeperLeaderRequest {
  s.DBClusterId = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetDisableWrite(v bool) *ElectZookeeperLeaderRequest {
  s.DisableWrite = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetElectTime(v string) *ElectZookeeperLeaderRequest {
  s.ElectTime = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetOwnerAccount(v string) *ElectZookeeperLeaderRequest {
  s.OwnerAccount = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetOwnerId(v int64) *ElectZookeeperLeaderRequest {
  s.OwnerId = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetPageNumber(v int32) *ElectZookeeperLeaderRequest {
  s.PageNumber = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetPageSize(v int32) *ElectZookeeperLeaderRequest {
  s.PageSize = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetRegionId(v string) *ElectZookeeperLeaderRequest {
  s.RegionId = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetResourceOwnerAccount(v string) *ElectZookeeperLeaderRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) SetResourceOwnerId(v int64) *ElectZookeeperLeaderRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *ElectZookeeperLeaderRequest) Validate() error {
  return dara.Validate(s)
}

