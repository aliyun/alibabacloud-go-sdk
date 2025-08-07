// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateRegionResourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceConnType(v string) *EvaluateRegionResourceRequest
  GetDBInstanceConnType() *string 
  SetDBNodeClass(v string) *EvaluateRegionResourceRequest
  GetDBNodeClass() *string 
  SetDBType(v string) *EvaluateRegionResourceRequest
  GetDBType() *string 
  SetDBVersion(v string) *EvaluateRegionResourceRequest
  GetDBVersion() *string 
  SetDispenseMode(v string) *EvaluateRegionResourceRequest
  GetDispenseMode() *string 
  SetNeedMaxScaleLink(v string) *EvaluateRegionResourceRequest
  GetNeedMaxScaleLink() *string 
  SetOwnerAccount(v string) *EvaluateRegionResourceRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EvaluateRegionResourceRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EvaluateRegionResourceRequest
  GetRegionId() *string 
  SetResourceGroupId(v string) *EvaluateRegionResourceRequest
  GetResourceGroupId() *string 
  SetResourceOwnerAccount(v string) *EvaluateRegionResourceRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EvaluateRegionResourceRequest
  GetResourceOwnerId() *int64 
  SetSubDomain(v string) *EvaluateRegionResourceRequest
  GetSubDomain() *string 
  SetZoneId(v string) *EvaluateRegionResourceRequest
  GetZoneId() *string 
}

type EvaluateRegionResourceRequest struct {
  // The cluster link type. The backend randomly selects the default value. Valid values:
  // 
  // 	- **lvs**: Linux virtual server.
  // 
  // 	- **proxy**: proxy server.
  // 
  // 	- **dns**: domain name system.
  // 
  // example:
  // 
  // lvs
  DBInstanceConnType *string `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
  // The specifications of the node. For information about node specifications, see the following topics:
  // 
  // 	- PolarDB for MySQL: [Specifications of compute nodes](https://help.aliyun.com/document_detail/102542.html)
  // 
  // 	- PolarDB for Oracle: [Specifications of compute nodes](https://help.aliyun.com/document_detail/207921.html)
  // 
  // 	- PolarDB for PostgreSQL: [Specifications of compute nodes](https://help.aliyun.com/document_detail/209380.html)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // polar.mysql.x4.large
  DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
  // The type of the database engine. Valid values:
  // 
  // 	- **MySQL**
  // 
  // 	- **PostgreSQL**
  // 
  // 	- **Oracle**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MySQL
  DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
  // The version of the database engine
  // 
  // 	- Valid values for the MySQL database engine:
  // 
  //     	- **5.6**
  // 
  //     	- **5.7**
  // 
  //     	- **8.0**
  // 
  // 	- Valid values for the PostgreSQL database engine:
  // 
  //     	- **11**
  // 
  //     	- **14**
  // 
  // 	- Valid value for the Oracle database engine: **11**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 8.0
  DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
  // Specifies whether to return the zones in which the single-zone deployment method is supported. Default value: 0. Valid values:
  // 
  // 	- **0**: no value returned
  // 
  // 	- **1**: returns the zones.
  // 
  // example:
  // 
  // 1
  DispenseMode *string `json:"DispenseMode,omitempty" xml:"DispenseMode,omitempty"`
  // Specifies whether to create Maxscale. Valid values:
  // 
  // 	- **true*	- (default)
  // 
  // 	- **false**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  NeedMaxScaleLink *string `json:"NeedMaxScaleLink,omitempty" xml:"NeedMaxScaleLink,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID.
  // 
  // > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the resource group.
  // 
  // example:
  // 
  // rg-************
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The subdomain. It is the child domain of the top-level domain name or parent domain. For example, if the parent domain name is cn-beijing, its child domain can be cn-beijing-i-aliyun.
  // 
  // example:
  // 
  // cn-beijing-i-aliyun
  SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
  // The zone ID.
  // 
  // > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available zones.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-g
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EvaluateRegionResourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateRegionResourceRequest) GoString() string {
  return s.String()
}

func (s *EvaluateRegionResourceRequest) GetDBInstanceConnType() *string  {
  return s.DBInstanceConnType
}

func (s *EvaluateRegionResourceRequest) GetDBNodeClass() *string  {
  return s.DBNodeClass
}

func (s *EvaluateRegionResourceRequest) GetDBType() *string  {
  return s.DBType
}

func (s *EvaluateRegionResourceRequest) GetDBVersion() *string  {
  return s.DBVersion
}

func (s *EvaluateRegionResourceRequest) GetDispenseMode() *string  {
  return s.DispenseMode
}

func (s *EvaluateRegionResourceRequest) GetNeedMaxScaleLink() *string  {
  return s.NeedMaxScaleLink
}

func (s *EvaluateRegionResourceRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EvaluateRegionResourceRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EvaluateRegionResourceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EvaluateRegionResourceRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EvaluateRegionResourceRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EvaluateRegionResourceRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EvaluateRegionResourceRequest) GetSubDomain() *string  {
  return s.SubDomain
}

func (s *EvaluateRegionResourceRequest) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EvaluateRegionResourceRequest) SetDBInstanceConnType(v string) *EvaluateRegionResourceRequest {
  s.DBInstanceConnType = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetDBNodeClass(v string) *EvaluateRegionResourceRequest {
  s.DBNodeClass = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetDBType(v string) *EvaluateRegionResourceRequest {
  s.DBType = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetDBVersion(v string) *EvaluateRegionResourceRequest {
  s.DBVersion = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetDispenseMode(v string) *EvaluateRegionResourceRequest {
  s.DispenseMode = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetNeedMaxScaleLink(v string) *EvaluateRegionResourceRequest {
  s.NeedMaxScaleLink = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetOwnerAccount(v string) *EvaluateRegionResourceRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetOwnerId(v int64) *EvaluateRegionResourceRequest {
  s.OwnerId = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetRegionId(v string) *EvaluateRegionResourceRequest {
  s.RegionId = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetResourceGroupId(v string) *EvaluateRegionResourceRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetResourceOwnerAccount(v string) *EvaluateRegionResourceRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetResourceOwnerId(v int64) *EvaluateRegionResourceRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetSubDomain(v string) *EvaluateRegionResourceRequest {
  s.SubDomain = &v
  return s
}

func (s *EvaluateRegionResourceRequest) SetZoneId(v string) *EvaluateRegionResourceRequest {
  s.ZoneId = &v
  return s
}

func (s *EvaluateRegionResourceRequest) Validate() error {
  return dara.Validate(s)
}

