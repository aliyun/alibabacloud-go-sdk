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
	SetEffectiveTime(v string) *UpgradeDBInstanceRequest
	GetEffectiveTime() *string
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
	// The Serverless cache storage capacity. Unit: GB.
	//
	// > This parameter is required only for Serverless Pro instances.
	//
	// example:
	//
	// 800
	CacheStorageSize *string `json:"CacheStorageSize,omitempty" xml:"CacheStorageSize,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-rj***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate*	- (default): The change takes effect immediately.
	//
	// 	- **MaintainTime**: The change takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The specifications of segment nodes. For information about supported node specifications, see [Instance specifications](https://help.aliyun.com/document_detail/35406.html).
	//
	// > This parameter is supported only for elastic storage mode instances.
	//
	// example:
	//
	// 4C16G
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	MasterNodeNum *string `json:"MasterNodeNum,omitempty" xml:"MasterNodeNum,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// null
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the resource group ID, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The performance level (PL) of the disk. Valid values:
	//
	// - **pl0**: PL0.
	//
	// - **pl1**: PL1.
	//
	// - **pl2**: PL2.
	//
	// example:
	//
	// pl1
	SegDiskPerformanceLevel *string `json:"SegDiskPerformanceLevel,omitempty" xml:"SegDiskPerformanceLevel,omitempty"`
	// The number of segment nodes. The supported number of nodes varies based on the instance resource type and instance edition:
	//
	// - Elastic storage mode, high-availability edition: valid values are 4 to 512, in increments of 4.
	//
	// - Elastic storage mode, <props="china">basic edition (formerly high-performance edition)<props="intl">high-performance edition: valid values are 2 to 512, in increments of 2.
	//
	// - Serverless manual scheduling mode: valid values are 2 to 512, in increments of 2.
	//
	// example:
	//
	// 2
	SegNodeNum *string `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty"`
	// The cloud disk storage type after the change. Currently, only ESSD cloud disks are supported. Set the value to **cloud_essd**.
	//
	// example:
	//
	// cloud_essd
	SegStorageType *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty"`
	// - Serverless instances:
	//
	// The compute resource threshold. Valid values: 8 to 32, in increments of 8. Unit: ACU. Default value: 32.
	//
	// - Serverless Pro instances: The reserved compute resources. Valid values: 16 to 1024. Unit: ACU. Default value: 16. The increment rules are as follows:
	//
	//   - Range 16 to 32: increments of 4.
	//
	//   - Range 32 to 64: increments of 8.
	//
	//   - Range 64 to 128: increments of 16.
	//
	//   - Range 128 to 256: increments of 32.
	//
	//   - Range greater than 256: increments of 64.
	//
	// > This parameter is required only for Serverless automatic scheduling mode and Serverless Pro instances.
	//
	// example:
	//
	// 16
	ServerlessResource *string `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
	// The storage capacity of segment nodes. Unit: GB. Valid values: 50 to <props="china">8000<props="intl">6000, in increments of 50.
	//
	// > This parameter is supported only for elastic storage mode instances.
	//
	// example:
	//
	// 100
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The type of the specification change. Valid values:
	//
	// - **0*	- (default): Changes the number of segment nodes. SegNodeNum is required. Other parameters do not take effect.
	//
	// - **1**: Changes the segment node specifications and instance storage capacity. InstanceSpec is required. StorageSize is optional and must be greater than or equal to the current instance storage capacity.
	//
	// - **2**: Changes the number of master nodes. MasterNodeNum is required. Other parameters do not take effect.
	//
	// - **3**: Changes the cloud disk storage type and performance level (PL). SegDiskPerformanceLevel and SegStorageType are required. Other parameters do not take effect.
	//
	// > - Different instance resource types support different Upgrade/Downgrade operations for compute nodes. For more information, see [Precautions](https://help.aliyun.com/document_detail/50956.html).
	//
	// - After you select a specification change type, only the corresponding parameters take effect. Other parameters do not take effect. For example, if **UpgradeType*	- is set to 0 and you specify both the number of segment nodes and the number of master nodes, only the number of segment nodes is changed.
	//
	// - Changing the number of master nodes is supported only on the China site (aliyun.com).
	//
	// - You can change the cloud disk storage type only from standard SSD to ESSD cloud disk.
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

func (s *UpgradeDBInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
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

func (s *UpgradeDBInstanceRequest) SetEffectiveTime(v string) *UpgradeDBInstanceRequest {
	s.EffectiveTime = &v
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
