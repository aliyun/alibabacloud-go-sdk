// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeAvailableClassesRequest
	GetCategory() *string
	SetCommodityCode(v string) *DescribeAvailableClassesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *DescribeAvailableClassesRequest
	GetDBInstanceId() *string
	SetDBInstanceStorageType(v string) *DescribeAvailableClassesRequest
	GetDBInstanceStorageType() *string
	SetEngine(v string) *DescribeAvailableClassesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableClassesRequest
	GetEngineVersion() *string
	SetInstanceChargeType(v string) *DescribeAvailableClassesRequest
	GetInstanceChargeType() *string
	SetOrderType(v string) *DescribeAvailableClassesRequest
	GetOrderType() *string
	SetRegionId(v string) *DescribeAvailableClassesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeAvailableClassesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeAvailableClassesRequest
	GetZoneId() *string
}

type DescribeAvailableClassesRequest struct {
	// The instance edition. Valid values:
	//
	// 	- Regular instances
	//
	//     	- **Basic**: Basic Edition
	//
	//     	- **HighAvailability**: high-availability series
	//
	//     	- **cluster**: Cluster Edition (applicable only to MySQL and PostgreSQL)
	//
	//     	- **AlwaysOn**: SQL Server Cluster Edition
	//
	//     	- **Finance**: RDS Enterprise Edition
	//
	// 	- Serverless instances
	//
	//     	- **serverless_basic**: Serverless Basic Edition (applicable only to MySQL and PostgreSQL)
	//
	//     	- **serverless_standard**: Serverless high availability series (applicable only to MySQL and PostgreSQL)
	//
	//     	- **serverless_ha**: SQL Server Serverless high availability series
	//
	//     > This parameter is required when you create a serverless instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the instance. Valid values:
	//
	// - **bards**: pay-as-you-go primary instance (China site)
	//
	// - **rds**: subscription primary instance (China site)
	//
	// - **rords**: pay-as-you-go read-only instance (China site)
	//
	// - **rds_rordspre_public_cn**: subscription read-only instance (China site)
	//
	// - **bards_intl**: pay-as-you-go primary instance (international site)
	//
	// - **rds_intl**: subscription primary instance (international site)
	//
	// - **rords_intl**: pay-as-you-go read-only instance (international site)
	//
	// - **rds_rordspre_public_intl**: subscription read-only instance (international site)
	//
	// - **rds_serverless_public_cn**: serverless (China site)
	//
	// - **rds_serverless_public_intl**: serverless (international site)
	//
	// > This parameter is required when you query a read-only instance.
	//
	// example:
	//
	// bards
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **general_essd**: premium performance disk
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd0**: PL0 ESSD cloud disk
	//
	// 	- **cloud_essd**: PL1 ESSD cloud disk
	//
	// 	- **cloud_essd2**: PL2 ESSD cloud disk
	//
	// 	- **cloud_essd3**: PL3 ESSD cloud disk
	//
	// > Serverless instances support only PL1 ESSD cloud disks. Set this parameter to **cloud_essd**.
	//
	// This parameter is required.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
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
	// The database engine version of the instance. Valid values:
	//
	// - Regular instances
	//
	//     - MySQL: **5.5, 5.6, 5.7, 8.0**
	//
	//     - SQL Server: **2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, 2019_ent**
	//
	//     - PostgreSQL: **10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0**
	//
	//     - MariaDB: **10.3**
	//
	// - Serverless instances
	//
	//     - MySQL: **5.7**, **8.0**
	//
	//     - SQL Server: **2016_std_sl**, **2017_std_sl**, **2019_std_sl**
	//
	//     - PostgreSQL: **14.0, 15.0, 16.0, 17.0**
	//
	//     > ApsaraDB RDS for MariaDB does not support serverless instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Prepaid**: subscription
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Serverless**: serverless
	//
	// > ApsaraDB RDS for MariaDB does not support serverless instances.
	//
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The order type. The only valid value is **BUY**.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID of the instance. You can call the DescribeDBInstanceAttribute operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the instance. You can call the DescribeDBInstanceAttribute operation to query the zone ID.
	//
	// >If DescribeDBInstanceAttribute returns a multi-zone value (such as `cn-hangzhou-MAZ9(g,h)`), specify a single zone. Example: `cn-hangzhou-g` or `cn-hangzhou-j`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableClassesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeAvailableClassesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAvailableClassesRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeAvailableClassesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableClassesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableClassesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableClassesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeAvailableClassesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableClassesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableClassesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableClassesRequest) SetCategory(v string) *DescribeAvailableClassesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetCommodityCode(v string) *DescribeAvailableClassesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetDBInstanceId(v string) *DescribeAvailableClassesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetDBInstanceStorageType(v string) *DescribeAvailableClassesRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetEngine(v string) *DescribeAvailableClassesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetEngineVersion(v string) *DescribeAvailableClassesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetInstanceChargeType(v string) *DescribeAvailableClassesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetOrderType(v string) *DescribeAvailableClassesRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetRegionId(v string) *DescribeAvailableClassesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetResourceOwnerId(v int64) *DescribeAvailableClassesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetZoneId(v string) *DescribeAvailableClassesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) Validate() error {
	return dara.Validate(s)
}
