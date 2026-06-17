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
  // The link type of the cluster. The backend randomly selects a default value. Valid values:
  // 
  // - **lvs**: Linux Virtual Server.
  // 
  // - **proxy**: proxy server.
  // 
  // - **dns**: Domain Name System.
  // 
  // example:
  // 
  // lvs
  DBInstanceConnType *string `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
  // The node specifications. For more information, see the following documents:
  // 
  // - PolarDB for MySQL: [Compute node specifications](https://help.aliyun.com/document_detail/102542.html).
  // 
  // - PolarDB for PostgreSQL (Oracle Compatible): [Compute node specifications](https://help.aliyun.com/document_detail/207921.html).
  // 
  // - PolarDB for PostgreSQL: [Compute node specifications](https://help.aliyun.com/document_detail/209380.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // polar.mysql.x4.large
  DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
  // The database engine type. Valid values:
  // 
  // - **MySQL**
  // 
  // - **PostgreSQL**
  // 
  // - **Oracle**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MySQL
  DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
  // The version of the database engine.
  // 
  // - Valid values for MySQL:
  // 
  //   - **5.6**
  // 
  //   - **5.7**
  // 
  //   - **8.0**
  // 
  // - Valid values for PostgreSQL:
  // 
  //   - **11**
  // 
  //   - **14**
  // 
  // - Valid values for Oracle:
  // 
  //   - **11**
  // 
  //   - **14**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 8.0
  DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
  // Specifies whether to return the list of zones that support single-zone deployment. Valid values:
  // 
  // - **0*	- (default): The list is not returned.
  // 
  // - **1**: The list is returned.
  // 
  // example:
  // 
  // 1
  DispenseMode *string `json:"DispenseMode,omitempty" xml:"DispenseMode,omitempty"`
  // Specifies whether to create a MaxScale cluster. Valid values:
  // 
  // - **true*	- (default): A MaxScale cluster is created.
  // 
  // - **false**: A MaxScale cluster is not created.
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
  // > Call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
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
  // The subdomain. A subdomain is a level below a top-level domain. For example, if the parent domain is \\`cn-beijing\\`, a valid subdomain is \\`cn-beijing-i-aliyun\\`.
  // 
  // example:
  // 
  // cn-beijing-i-aliyun
  SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
  // The zone ID.
  // 
  // > Call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available zones.
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

