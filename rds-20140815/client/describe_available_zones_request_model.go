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
	// The RDS edition of the instance. Valid values:
	//
	// 	- Regular instance
	//
	//     	- **Basic**: RDS Basic Edition.
	//
	//     	- **HighAvailability**: RDS High-availability Edition.
	//
	//     	- **cluster**: RDS Cluster Edition for ApsaraDB RDS for MySQL.
	//
	//     	- **AlwaysOn**: RDS Cluster Edition for ApsaraDB RDS for SQL Server.
	//
	//     	- **Finance**: RDS Enterprise Edition.
	//
	// 	- Serverless instance
	//
	//     	- **serverless_basic**: RDS Basic Edition. This edition is available only for instances that run MySQL and PostgreSQL.
	//
	//     	- **serverless_standard**: RDS High-availability Edition for ApsaraDB RDS for MySQL.
	//
	//     	- **serverless_ha**: RDS High-availability Edition for ApsaraDB RDS for SQL Server.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the instance. This operation can return the resources that you can purchase based on the specified commodity code. Valid values:
	//
	// 	- **bards**: The instance is a pay-as-you-go primary instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rds**: The instance is a subscription primary instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rords**: The instance is a pay-as-you-go read-only instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available at the China site (aliyun.com).
	//
	// 	- **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available at the International site (alibabacloud.com).
	//
	// 	- **rds_intl**: The instance is a subscription primary instance. This value is available at the International site (alibabacloud.com).
	//
	// 	- **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available at the International site (alibabacloud.com).
	//
	// 	- **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available at the International site (alibabacloud.com).
	//
	// 	- **rds_serverless_public_cn**: The instance is a serverless instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rds_serverless_public_intl**: The instance is a serverless instance. This value is available at the International site (alibabacloud.com).
	//
	// example:
	//
	// bards
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the primary instance. If you want to query the read-only instances that you can purchase for a primary instance, you can specify this parameter.
	//
	// If you set **CommodityCode*	- to one of the following values, you must specify this parameter:
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
	// rm-uf6wjk5xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Specifies whether to return the zones in which the single-zone deployment method is supported. Valid values:
	//
	// 	- **1*	- (default): returns the zones.
	//
	// 	- **0**: does not return the zones.
	//
	// >  The single-zone deployment method allows you to deploy an instance that runs RDS Enterprise Edition in a single zone.
	//
	// example:
	//
	// 0
	DispenseMode *string `json:"DispenseMode,omitempty" xml:"DispenseMode,omitempty"`
	// The database engine of the instance. Valid values:
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
	// 	- Regular instance
	//
	//     	- Valid values if you set Engine to MySQL: **5.5**, **5.6**, **5.7**, and **8.0**
	//
	//     	- Valid values if you set Engine to SQLServer: **2008r2**, **08r2_ent_ha**, **2012**, **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_std_ha**, **2017_ent**, **2019_std_ha**, and **2019_ent**
	//
	//     	- Valid values if you set Engine to PostgreSQL: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, and **15.0**
	//
	//     	- Valid value when you set Engine to MariaDB: **10.3**
	//
	// 	- Serverless instance
	//
	//     	- Valid values if you set Engine to MySQL: **5.7*	- and **8.0**
	//
	//     	- Valid values if you set Engine to SQLServer: **2016_std_sl**, **2017_std_sl**, and **2019_std_sl**
	//
	//     	- Valid value if you set Engine to PostgreSQL: **14.0**
	//
	//     **
	//
	//     **Note**ApsaraDB RDS for MariaDB does not support serverless instances.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID. If the instance spans more than one zone, the value of this parameter contains an `MAZ` part, such as `cn-hangzhou-MAZ6(b,f)` and `cn-hangzhou-MAZ5(b,e,f)`. You can call the DescribeRegions operation to query the most recent zone list.
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
