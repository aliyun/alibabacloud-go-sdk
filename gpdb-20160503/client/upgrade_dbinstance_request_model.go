// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheStorageSize(v string) *UpgradeDBInstanceRequest
	GetCacheStorageSize() *string
	SetDBInstanceClass(v string) *UpgradeDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceGroupCount(v string) *UpgradeDBInstanceRequest
	GetDBInstanceGroupCount() *string
	SetDBInstanceId(v string) *UpgradeDBInstanceRequest
	GetDBInstanceId() *string
	SetInstanceSpec(v string) *UpgradeDBInstanceRequest
	GetInstanceSpec() *string
	SetMasterNodeNum(v string) *UpgradeDBInstanceRequest
	GetMasterNodeNum() *string
	SetOwnerId(v int64) *UpgradeDBInstanceRequest
	GetOwnerId() *int64
	SetPayType(v string) *UpgradeDBInstanceRequest
	GetPayType() *string
	SetRegionId(v string) *UpgradeDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpgradeDBInstanceRequest
	GetResourceGroupId() *string
	SetSegDiskPerformanceLevel(v string) *UpgradeDBInstanceRequest
	GetSegDiskPerformanceLevel() *string
	SetSegNodeNum(v string) *UpgradeDBInstanceRequest
	GetSegNodeNum() *string
	SetSegStorageType(v string) *UpgradeDBInstanceRequest
	GetSegStorageType() *string
	SetServerlessResource(v string) *UpgradeDBInstanceRequest
	GetServerlessResource() *string
	SetStorageSize(v string) *UpgradeDBInstanceRequest
	GetStorageSize() *string
	SetUpgradeType(v int64) *UpgradeDBInstanceRequest
	GetUpgradeType() *int64
}

type UpgradeDBInstanceRequest struct {
	CacheStorageSize *string `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// null
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// null
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-rj***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The specifications of each compute node. For information about the supported specifications, see [Instance specifications](https://help.aliyun.com/document_detail/35406.html).
	//
	// > This parameter is available only for instances in elastic storage mode.
	//
	// example:
	//
	// 4C16G
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// 2
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// null
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The performance level of Enterprise SSDs (ESSDs). Valid values:
	//
	// 	- **pl0**
	//
	// 	- **pl1**
	//
	// 	- **pl2**
	//
	// example:
	//
	// pl1
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of compute nodes. The number of compute nodes varies based on the instance resource type and edition.
	//
	// 	- Valid values for High-availability Edition instances in elastic storage mode: 4 to 512, in 4 increments.
	//
	// 	- Valid values for High-performance Edition instances in elastic storage mode: 2 to 512, in 2 increments.
	//
	// 	- Valid values for instances in manual Serverless mode: 2 to 512, in 2 increments.
	//
	// example:
	//
	// 2
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The disk storage type of the instance after the change. The disk storage type can be changed only to ESSD. Set the value to **cloud_essd**.
	//
	// example:
	//
	// cloud_essd
	SegStorageType     *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	ServerlessResource *string `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The storage capacity of each compute node. Unit: GB. Valid values: 50 to 6000, in 50 increments.
	//
	// >  This parameter is available only for instances in elastic storage mode.
	//
	// example:
	//
	// 100
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The type of the instance configuration change. Valid values:
	//
	// 	- **0*	- (default): changes the number of compute nodes.
	//
	// 	- **1**: changes the specifications and storage capacity of each compute node.
	//
	// 	- **2**: changes the number of coordinator nodes.
	//
	// 	- **3**: changes the disk storage type and ESSD performance level of the instance.
	//
	// >
	//
	// 	- The supported changes to compute node configurations vary based on the instance resource type. For more information, see the "Usage notes" section of the [Change compute node configurations](https://help.aliyun.com/document_detail/50956.html) topic.
	//
	// 	- After you specify a change type, only the corresponding parameters take effect. For example, if you set **UpgradeType*	- to 0, the parameter that is used to change the number of compute nodes takes effect, but the parameter that is used to change the number of coordinator nodes does not.
	//
	// 	- The number of coordinator nodes can be changed only on the China site (aliyun.com).
	//
	// 	- The disk storage type can be changed only from ultra disks to ESSDs.
	//
	// example:
	//
	// 0
	UpgradeType *int64 `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceRequest) GetCacheStorageSize() *string {
	return s.CacheStorageSize
}

func (s *UpgradeDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *UpgradeDBInstanceRequest) GetDBInstanceGroupCount() *string {
	return s.DBInstanceGroupCount
}

func (s *UpgradeDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceRequest) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *UpgradeDBInstanceRequest) GetMasterNodeNum() *string {
	return s.MasterNodeNum
}

func (s *UpgradeDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *UpgradeDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpgradeDBInstanceRequest) GetSegDiskPerformanceLevel() *string {
	return s.SegDiskPerformanceLevel
}

func (s *UpgradeDBInstanceRequest) GetSegNodeNum() *string {
	return s.SegNodeNum
}

func (s *UpgradeDBInstanceRequest) GetSegStorageType() *string {
	return s.SegStorageType
}

func (s *UpgradeDBInstanceRequest) GetServerlessResource() *string {
	return s.ServerlessResource
}

func (s *UpgradeDBInstanceRequest) GetStorageSize() *string {
	return s.StorageSize
}

func (s *UpgradeDBInstanceRequest) GetUpgradeType() *int64 {
	return s.UpgradeType
}

func (s *UpgradeDBInstanceRequest) SetCacheStorageSize(v string) *UpgradeDBInstanceRequest {
	s.CacheStorageSize = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceGroupCount(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceId(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetInstanceSpec(v string) *UpgradeDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetMasterNodeNum(v string) *UpgradeDBInstanceRequest {
	s.MasterNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetOwnerId(v int64) *UpgradeDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetPayType(v string) *UpgradeDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetRegionId(v string) *UpgradeDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetResourceGroupId(v string) *UpgradeDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegDiskPerformanceLevel(v string) *UpgradeDBInstanceRequest {
	s.SegDiskPerformanceLevel = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegNodeNum(v string) *UpgradeDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetSegStorageType(v string) *UpgradeDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetServerlessResource(v string) *UpgradeDBInstanceRequest {
	s.ServerlessResource = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetStorageSize(v string) *UpgradeDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetUpgradeType(v int64) *UpgradeDBInstanceRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
