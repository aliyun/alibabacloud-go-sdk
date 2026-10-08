// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeAvailableZonesRequest
	GetCategory() *string
	SetCommodityCode(v string) *DescribeAvailableZonesRequest
	GetCommodityCode() *string
	SetDBInstanceName(v string) *DescribeAvailableZonesRequest
	GetDBInstanceName() *string
	SetDispenseMode(v string) *DescribeAvailableZonesRequest
	GetDispenseMode() *string
	SetEngine(v string) *DescribeAvailableZonesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableZonesRequest
	GetEngineVersion() *string
	SetRegionId(v string) *DescribeAvailableZonesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeAvailableZonesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeAvailableZonesRequest
	GetZoneId() *string
}

type DescribeAvailableZonesRequest struct {
	// The instance edition. Valid values:
	//
	// 	- Regular instances
	//
	//     	- **Basic**: Basic Edition
	//
	//     	- **HighAvailability**: High-availability Edition
	//
	//     	- **cluster**: MySQL Cluster Edition
	//
	//     	- **AlwaysOn**: SQL Server Cluster Edition
	//
	//     	- **Finance**: RDS Enterprise Edition
	//
	// 	- Serverless instances
	//
	//     	- **serverless_basic**: Serverless Basic Edition (applicable only to MySQL and PostgreSQL)
	//
	//     	- **serverless_standard**: MySQL Serverless High-availability Edition
	//
	//     	- **serverless_ha**: SQL Server Serverless High-availability Edition
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the instance. The operation queries available resources for sale based on the specified commodity code. Valid values:
	//
	// 	- **bards**: pay-as-you-go primary instance (China site)
	//
	// 	- **rds**: subscription primary instance (China site)
	//
	// 	- **rords**: pay-as-you-go read-only instance (China site)
	//
	// 	- **rds_rordspre_public_cn**: subscription read-only instance (China site)
	//
	// 	- **bards_intl**: pay-as-you-go primary instance (international site)
	//
	// 	- **rds_intl**: subscription primary instance (international site)
	//
	// 	- **rords_intl**: pay-as-you-go read-only instance (international site)
	//
	// 	- **rds_rordspre_public_intl**: subscription read-only instance (international site)
	//
	// 	- **rds_serverless_public_cn**: serverless (China site)
	//
	// 	- **rds_serverless_public_intl**: serverless (international site)
	//
	// example:
	//
	// bards
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance ID of the primary instance. This parameter is used to query available read-only instance resources for the specified primary instance.
	//
	// This parameter is required when **CommodityCode*	- is set to one of the following values:
	//
	// 	- **rords_intl**
	//
	// 	- **rds_rordspre_public_intl**
	//
	// 	- **rords**
	//
	// 	- **rds_rordspre_public_cn**
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Specifies whether to return the list of zones that support single-zone deployment. Valid values:
	//
	// 	- **1*	- (default): Returns the list.
	//
	// 	- **0**: Does not return the list.
	//
	// > The single-zone deployment feature allows you to deploy RDS Enterprise Edition instances in a single zone.
	//
	// example:
	//
	// 0
	DispenseMode *string `json:"DispenseMode,omitempty" xml:"DispenseMode,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. Valid values:
	//
	// - Regular instances
	//
	//     - MySQL: **5.5**, **5.6**, **5.7**, **8.0**
	//
	//     - SQL Server: **2008r2**, **08r2_ent_ha**, **2012**, **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_std_ha**, **2017_ent**, **2019_std_ha**, **2019_ent**
	//
	//     - PostgreSQL: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**
	//
	//     - MariaDB: **10.3**
	//
	// - Serverless instances
	//
	//     - MySQL: **5.7**, **8.0**
	//
	//     - SQL Server: **2016_std_sl**, **2017_std_sl**, **2019_std_sl**
	//
	//     - PostgreSQL: **14.0**
	//
	//     > MariaDB does not support serverless instances.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The region ID. You can call DescribeRegions to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID. The format of multi-zone IDs differs from that of single-zone IDs and contains `MAZ`, such as `cn-hangzhou-MAZ6(b,f)` and `cn-hangzhou-MAZ5(b,e,f)`. You can call DescribeRegions to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableZonesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeAvailableZonesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAvailableZonesRequest) GetDispenseMode() *string {
	return s.DispenseMode
}

func (s *DescribeAvailableZonesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableZonesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableZonesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableZonesRequest) SetCategory(v string) *DescribeAvailableZonesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetCommodityCode(v string) *DescribeAvailableZonesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetDBInstanceName(v string) *DescribeAvailableZonesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetDispenseMode(v string) *DescribeAvailableZonesRequest {
	s.DispenseMode = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetEngine(v string) *DescribeAvailableZonesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetEngineVersion(v string) *DescribeAvailableZonesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetRegionId(v string) *DescribeAvailableZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetResourceOwnerId(v int64) *DescribeAvailableZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetZoneId(v string) *DescribeAvailableZonesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
