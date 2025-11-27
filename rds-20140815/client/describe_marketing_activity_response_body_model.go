// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketingActivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *DescribeMarketingActivityResponseBody
	GetAliUid() *int64
	SetBid(v string) *DescribeMarketingActivityResponseBody
	GetBid() *string
	SetItems(v []*DescribeMarketingActivityResponseBodyItems) *DescribeMarketingActivityResponseBody
	GetItems() []*DescribeMarketingActivityResponseBodyItems
	SetRegionId(v string) *DescribeMarketingActivityResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeMarketingActivityResponseBody
	GetRequestId() *string
}

type DescribeMarketingActivityResponseBody struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1979008652307170
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 	- China site: 26842
	//
	// 	- International site: 26888
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The activity parameters
	Items []*DescribeMarketingActivityResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7D8F09DB-1124-5D78-A520-FF88FAF4351B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMarketingActivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketingActivityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMarketingActivityResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeMarketingActivityResponseBody) GetBid() *string {
	return s.Bid
}

func (s *DescribeMarketingActivityResponseBody) GetItems() []*DescribeMarketingActivityResponseBodyItems {
	return s.Items
}

func (s *DescribeMarketingActivityResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMarketingActivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMarketingActivityResponseBody) SetAliUid(v int64) *DescribeMarketingActivityResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeMarketingActivityResponseBody) SetBid(v string) *DescribeMarketingActivityResponseBody {
	s.Bid = &v
	return s
}

func (s *DescribeMarketingActivityResponseBody) SetItems(v []*DescribeMarketingActivityResponseBodyItems) *DescribeMarketingActivityResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMarketingActivityResponseBody) SetRegionId(v string) *DescribeMarketingActivityResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeMarketingActivityResponseBody) SetRequestId(v string) *DescribeMarketingActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMarketingActivityResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMarketingActivityResponseBodyItems struct {
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **AlwaysOn**: RDS Cluster Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// Basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The payment type. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The instance type. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html) and [Read-only ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/145759.html).
	//
	// example:
	//
	// rds.mysql.s3.large
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The instance family. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/57184.html).
	//
	// example:
	//
	// x
	ClassGroup *string `json:"ClassGroup,omitempty" xml:"ClassGroup,omitempty"`
	// The number of CPU cores that are supported by the instance type. Unit: cores.
	//
	// example:
	//
	// 2
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disk capacity per node. Unit: GB.
	//
	// example:
	//
	// 900
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- MySQL
	//
	// 	- SQLServer
	//
	// 	- PostgreSQL
	//
	// 	- PPAS
	//
	// 	- MariaDB
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 60
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum I/O throughput. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	MaxIombps *int32 `json:"MaxIombps,omitempty" xml:"MaxIombps,omitempty"`
	// The maximum IOPS.
	//
	// example:
	//
	// 30
	MaxIops *int32 `json:"MaxIops,omitempty" xml:"MaxIops,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 1024
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: performance level 1 (PL1) enhanced SSD (ESSD)
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The RDS edition after the upgrade.
	//
	// example:
	//
	// HighAvailability
	UpgradeCategory *string `json:"UpgradeCategory,omitempty" xml:"UpgradeCategory,omitempty"`
	// The instance type after the upgrade.
	//
	// example:
	//
	// rds.mysql.s3.large
	UpgradeClassCode *string `json:"UpgradeClassCode,omitempty" xml:"UpgradeClassCode,omitempty"`
	// The instance family after the upgrade.
	//
	// example:
	//
	// d
	UpgradeClassGroup *string `json:"UpgradeClassGroup,omitempty" xml:"UpgradeClassGroup,omitempty"`
	// The number of CPU cores after the upgrade.
	//
	// example:
	//
	// 8
	UpgradeCpu *string `json:"UpgradeCpu,omitempty" xml:"UpgradeCpu,omitempty"`
	// The description of the upgrade.
	//
	// example:
	//
	// test
	UpgradeDescContent *string `json:"UpgradeDescContent,omitempty" xml:"UpgradeDescContent,omitempty"`
	// The disk capacity after the upgrade.
	//
	// example:
	//
	// 1024
	UpgradeDiskSize *int32 `json:"UpgradeDiskSize,omitempty" xml:"UpgradeDiskSize,omitempty"`
	// The maximum number of concurrent connections after the upgrade.
	//
	// example:
	//
	// 70
	UpgradeMaxConnections *int32 `json:"UpgradeMaxConnections,omitempty" xml:"UpgradeMaxConnections,omitempty"`
	// The maximum I/O throughput after the upgrade. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	UpgradeMaxIombps *int32 `json:"UpgradeMaxIombps,omitempty" xml:"UpgradeMaxIombps,omitempty"`
	// The maximum IOPS after the upgrade.
	//
	// example:
	//
	// 70
	UpgradeMaxIops *int32 `json:"UpgradeMaxIops,omitempty" xml:"UpgradeMaxIops,omitempty"`
	// The memory size after the upgrade.
	//
	// example:
	//
	// 1024
	UpgradeMemory *int64 `json:"UpgradeMemory,omitempty" xml:"UpgradeMemory,omitempty"`
	// The reference price of the upgrade.
	//
	// example:
	//
	// 23333.1
	UpgradeReferencePrice *string `json:"UpgradeReferencePrice,omitempty" xml:"UpgradeReferencePrice,omitempty"`
	// The storage type after the upgrade.
	//
	// example:
	//
	// cloud_essd
	UpgradeStorageType *string `json:"UpgradeStorageType,omitempty" xml:"UpgradeStorageType,omitempty"`
}

func (s DescribeMarketingActivityResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketingActivityResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMarketingActivityResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeMarketingActivityResponseBodyItems) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeMarketingActivityResponseBodyItems) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeMarketingActivityResponseBodyItems) GetClassGroup() *string {
	return s.ClassGroup
}

func (s *DescribeMarketingActivityResponseBodyItems) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeMarketingActivityResponseBodyItems) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeMarketingActivityResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeMarketingActivityResponseBodyItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeMarketingActivityResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMarketingActivityResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeMarketingActivityResponseBodyItems) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeMarketingActivityResponseBodyItems) GetMaxIombps() *int32 {
	return s.MaxIombps
}

func (s *DescribeMarketingActivityResponseBodyItems) GetMaxIops() *int32 {
	return s.MaxIops
}

func (s *DescribeMarketingActivityResponseBodyItems) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeMarketingActivityResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeCategory() *string {
	return s.UpgradeCategory
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeClassCode() *string {
	return s.UpgradeClassCode
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeClassGroup() *string {
	return s.UpgradeClassGroup
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeCpu() *string {
	return s.UpgradeCpu
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeDescContent() *string {
	return s.UpgradeDescContent
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeDiskSize() *int32 {
	return s.UpgradeDiskSize
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeMaxConnections() *int32 {
	return s.UpgradeMaxConnections
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeMaxIombps() *int32 {
	return s.UpgradeMaxIombps
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeMaxIops() *int32 {
	return s.UpgradeMaxIops
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeMemory() *int64 {
	return s.UpgradeMemory
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeReferencePrice() *string {
	return s.UpgradeReferencePrice
}

func (s *DescribeMarketingActivityResponseBodyItems) GetUpgradeStorageType() *string {
	return s.UpgradeStorageType
}

func (s *DescribeMarketingActivityResponseBodyItems) SetCategory(v string) *DescribeMarketingActivityResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetChargeType(v string) *DescribeMarketingActivityResponseBodyItems {
	s.ChargeType = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetClassCode(v string) *DescribeMarketingActivityResponseBodyItems {
	s.ClassCode = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetClassGroup(v string) *DescribeMarketingActivityResponseBodyItems {
	s.ClassGroup = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetCpu(v string) *DescribeMarketingActivityResponseBodyItems {
	s.Cpu = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetDiskSize(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.DiskSize = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetEngine(v string) *DescribeMarketingActivityResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetEngineVersion(v string) *DescribeMarketingActivityResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetInstanceId(v string) *DescribeMarketingActivityResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetInstanceName(v string) *DescribeMarketingActivityResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetMaxConnections(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.MaxConnections = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetMaxIombps(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.MaxIombps = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetMaxIops(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.MaxIops = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetMemory(v int64) *DescribeMarketingActivityResponseBodyItems {
	s.Memory = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetStorageType(v string) *DescribeMarketingActivityResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeCategory(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeCategory = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeClassCode(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeClassCode = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeClassGroup(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeClassGroup = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeCpu(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeCpu = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeDescContent(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeDescContent = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeDiskSize(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeDiskSize = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeMaxConnections(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeMaxConnections = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeMaxIombps(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeMaxIombps = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeMaxIops(v int32) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeMaxIops = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeMemory(v int64) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeMemory = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeReferencePrice(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeReferencePrice = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) SetUpgradeStorageType(v string) *DescribeMarketingActivityResponseBodyItems {
	s.UpgradeStorageType = &v
	return s
}

func (s *DescribeMarketingActivityResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
