// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAIDBClusterDatasetServiceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableAIDBClusterDatasetServiceRequest
  GetDBClusterId() *string 
  SetOwnerAccount(v string) *EnableAIDBClusterDatasetServiceRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableAIDBClusterDatasetServiceRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableAIDBClusterDatasetServiceRequest
  GetRegionId() *string 
  SetRelativeDBClusterId(v string) *EnableAIDBClusterDatasetServiceRequest
  GetRelativeDBClusterId() *string 
  SetResourceOwnerAccount(v string) *EnableAIDBClusterDatasetServiceRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableAIDBClusterDatasetServiceRequest
  GetResourceOwnerId() *int64 
}

type EnableAIDBClusterDatasetServiceRequest struct {
  // The instance ID of the PolarFS cold storage instance that you attach to the dataset management service. Instructions to invoke: Note: This parameter is actually required. If it is not provided, the API returns InvalidDBClusterId.Malformed (400). Instructions to invoke: Valid ID format description: specifies the accepted prefix families (such as pcs-xxx), separator conventions (hyphen character vs. underscore character), and the corresponding resource types.
  // 
  // example:
  // 
  // pcs-2ze26***
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID.
  // 
  // example:
  // 
  // cn-beijing
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The instance ID of the database instance that you attach to the dataset management service.
  // 
  // example:
  // 
  // pc-2ze88***
  RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableAIDBClusterDatasetServiceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAIDBClusterDatasetServiceRequest) GoString() string {
  return s.String()
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetRelativeDBClusterId() *string  {
  return s.RelativeDBClusterId
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableAIDBClusterDatasetServiceRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetDBClusterId(v string) *EnableAIDBClusterDatasetServiceRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetOwnerAccount(v string) *EnableAIDBClusterDatasetServiceRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetOwnerId(v int64) *EnableAIDBClusterDatasetServiceRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetRegionId(v string) *EnableAIDBClusterDatasetServiceRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetRelativeDBClusterId(v string) *EnableAIDBClusterDatasetServiceRequest {
  s.RelativeDBClusterId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetResourceOwnerAccount(v string) *EnableAIDBClusterDatasetServiceRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) SetResourceOwnerId(v int64) *EnableAIDBClusterDatasetServiceRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceRequest) Validate() error {
  return dara.Validate(s)
}

