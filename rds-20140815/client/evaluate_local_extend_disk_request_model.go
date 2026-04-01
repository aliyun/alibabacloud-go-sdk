// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateLocalExtendDiskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceName(v string) *EvaluateLocalExtendDiskRequest
  GetDBInstanceName() *string 
  SetOwnerId(v int64) *EvaluateLocalExtendDiskRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EvaluateLocalExtendDiskRequest
  GetRegionId() *string 
  SetResourceGroupId(v string) *EvaluateLocalExtendDiskRequest
  GetResourceGroupId() *string 
  SetResourceOwnerAccount(v string) *EvaluateLocalExtendDiskRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EvaluateLocalExtendDiskRequest
  GetResourceOwnerId() *int64 
  SetStorage(v int32) *EvaluateLocalExtendDiskRequest
  GetStorage() *int32 
}

type EvaluateLocalExtendDiskRequest struct {
  // The instance name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // rm-m5e999iqm65******
  DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The resource group ID.
  // 
  // example:
  // 
  // rg-ac****
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The new storage capacity. Unit: GB.
  // 
  // example:
  // 
  // 1000
  Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s EvaluateLocalExtendDiskRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateLocalExtendDiskRequest) GoString() string {
  return s.String()
}

func (s *EvaluateLocalExtendDiskRequest) GetDBInstanceName() *string  {
  return s.DBInstanceName
}

func (s *EvaluateLocalExtendDiskRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EvaluateLocalExtendDiskRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EvaluateLocalExtendDiskRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EvaluateLocalExtendDiskRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EvaluateLocalExtendDiskRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EvaluateLocalExtendDiskRequest) GetStorage() *int32  {
  return s.Storage
}

func (s *EvaluateLocalExtendDiskRequest) SetDBInstanceName(v string) *EvaluateLocalExtendDiskRequest {
  s.DBInstanceName = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetOwnerId(v int64) *EvaluateLocalExtendDiskRequest {
  s.OwnerId = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetRegionId(v string) *EvaluateLocalExtendDiskRequest {
  s.RegionId = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetResourceGroupId(v string) *EvaluateLocalExtendDiskRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetResourceOwnerAccount(v string) *EvaluateLocalExtendDiskRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetResourceOwnerId(v int64) *EvaluateLocalExtendDiskRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) SetStorage(v int32) *EvaluateLocalExtendDiskRequest {
  s.Storage = &v
  return s
}

func (s *EvaluateLocalExtendDiskRequest) Validate() error {
  return dara.Validate(s)
}

