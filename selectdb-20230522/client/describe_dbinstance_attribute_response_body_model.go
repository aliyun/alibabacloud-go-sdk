// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanUpgradeVersions(v []*string) *DescribeDBInstanceAttributeResponseBody
	GetCanUpgradeVersions() []*string
	SetChargeType(v string) *DescribeDBInstanceAttributeResponseBody
	GetChargeType() *string
	SetConfigPatternType(v string) *DescribeDBInstanceAttributeResponseBody
	GetConfigPatternType() *string
	SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBody
	GetCreateTime() *string
	SetDBClusterList(v []*DescribeDBInstanceAttributeResponseBodyDBClusterList) *DescribeDBInstanceAttributeResponseBody
	GetDBClusterList() []*DescribeDBInstanceAttributeResponseBodyDBClusterList
	SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBody
	GetDBInstanceId() *string
	SetDeployScheme(v string) *DescribeDBInstanceAttributeResponseBody
	GetDeployScheme() *string
	SetDescription(v string) *DescribeDBInstanceAttributeResponseBody
	GetDescription() *string
	SetEngine(v string) *DescribeDBInstanceAttributeResponseBody
	GetEngine() *string
	SetEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBody
	GetEngineMinorVersion() *string
	SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBody
	GetEngineVersion() *string
	SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBody
	GetExpireTime() *string
	SetGmtModified(v string) *DescribeDBInstanceAttributeResponseBody
	GetGmtModified() *string
	SetLockMode(v int64) *DescribeDBInstanceAttributeResponseBody
	GetLockMode() *int64
	SetLockReason(v string) *DescribeDBInstanceAttributeResponseBody
	GetLockReason() *string
	SetMaintainEndtime(v string) *DescribeDBInstanceAttributeResponseBody
	GetMaintainEndtime() *string
	SetMaintainStarttime(v string) *DescribeDBInstanceAttributeResponseBody
	GetMaintainStarttime() *string
	SetMultiZone(v []*DescribeDBInstanceAttributeResponseBodyMultiZone) *DescribeDBInstanceAttributeResponseBody
	GetMultiZone() []*DescribeDBInstanceAttributeResponseBodyMultiZone
	SetObjectStoreSize(v int64) *DescribeDBInstanceAttributeResponseBody
	GetObjectStoreSize() *int64
	SetRegionId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRequestId() *string
	SetResourceCpu(v int64) *DescribeDBInstanceAttributeResponseBody
	GetResourceCpu() *int64
	SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBody
	GetResourceGroupId() *string
	SetSecGroupConnValid(v string) *DescribeDBInstanceAttributeResponseBody
	GetSecGroupConnValid() *string
	SetServerless(v bool) *DescribeDBInstanceAttributeResponseBody
	GetServerless() *bool
	SetStatus(v string) *DescribeDBInstanceAttributeResponseBody
	GetStatus() *string
	SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBody
	GetStorageSize() *int64
	SetSubDomain(v string) *DescribeDBInstanceAttributeResponseBody
	GetSubDomain() *string
	SetTags(v []*DescribeDBInstanceAttributeResponseBodyTags) *DescribeDBInstanceAttributeResponseBody
	GetTags() []*DescribeDBInstanceAttributeResponseBodyTags
	SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBody
	GetVSwitchId() *string
	SetVirtualClusterList(v []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList) *DescribeDBInstanceAttributeResponseBody
	GetVirtualClusterList() []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList
	SetVpcId(v string) *DescribeDBInstanceAttributeResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The information returned.
	CanUpgradeVersions []*string `json:"CanUpgradeVersions,omitempty" xml:"CanUpgradeVersions,omitempty" type:"Repeated"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	ChargeType        *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ConfigPatternType *string `json:"ConfigPatternType,omitempty" xml:"ConfigPatternType,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-08-14T03:00:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about each cluster returned.
	DBClusterList []*DescribeDBInstanceAttributeResponseBodyDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// multi_az
	DeployScheme *string `json:"DeployScheme,omitempty" xml:"DeployScheme,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor kernel version number of the instance.
	//
	// example:
	//
	// 3.0.1
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2023-09-17T00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance was last modified, such as when you restarted the instance or applied for a public endpoint for the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-08-17T09:58Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The lock mode of the instance. Set the value to **lock**, which specifies that the instance is locked when it automatically expires or has an overdue payment.
	//
	// example:
	//
	// lock
	LockMode *int64 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the instance is locked.
	//
	// example:
	//
	// nolock
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The end time of the instance maintenance window.
	//
	// example:
	//
	// 1970-01-01T05:00Z
	MaintainEndtime *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	// The start time of the instance maintenance window.
	//
	// example:
	//
	// 1970-01-01T02:00Z
	MaintainStarttime *string `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	// if can be null:
	// true
	MultiZone []*DescribeDBInstanceAttributeResponseBodyMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 0
	ObjectStoreSize *int64 `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The Region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of CPU cores of the instance.
	//
	// example:
	//
	// 8
	ResourceCpu *int64 `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzbck4asz3dsa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// false
	SecGroupConnValid *string `json:"SecGroupConnValid,omitempty" xml:"SecGroupConnValid,omitempty"`
	// example:
	//
	// false
	Serverless *bool `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **RESOURCE_CHANGING**: The resource configuration of the instance is being changed.
	//
	// 	- **ORDER_PREPARING**: The order is being confirmed.
	//
	// 	- **READONLY_RESOURCE_CHANGING**: The resource configuration of the instance is being changed and the instance is write-locked.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cache size.
	//
	// example:
	//
	// 400
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The subdomain zone ID.
	//
	// example:
	//
	// cn-beijing-h-aliyun
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The tags that are added to the instances. Each tag is a key-value pair that consists of two parts: TagKey and TagValue. Format: `{"key1":"value1"}`.
	Tags []*DescribeDBInstanceAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-bp18iztwqrs8qj2nc6nyu
	VSwitchId          *string                                                      `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VirtualClusterList []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList `json:"VirtualClusterList,omitempty" xml:"VirtualClusterList,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The Zone ID.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) GetCanUpgradeVersions() []*string {
	return s.CanUpgradeVersions
}

func (s *DescribeDBInstanceAttributeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstanceAttributeResponseBody) GetConfigPatternType() *string {
	return s.ConfigPatternType
}

func (s *DescribeDBInstanceAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDBClusterList() []*DescribeDBInstanceAttributeResponseBodyDBClusterList {
	return s.DBClusterList
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDeployScheme() *string {
	return s.DeployScheme
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceAttributeResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeResponseBody) GetEngineMinorVersion() *string {
	return s.EngineMinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstanceAttributeResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDBInstanceAttributeResponseBody) GetLockMode() *int64 {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBody) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceAttributeResponseBody) GetMaintainEndtime() *string {
	return s.MaintainEndtime
}

func (s *DescribeDBInstanceAttributeResponseBody) GetMaintainStarttime() *string {
	return s.MaintainStarttime
}

func (s *DescribeDBInstanceAttributeResponseBody) GetMultiZone() []*DescribeDBInstanceAttributeResponseBodyMultiZone {
	return s.MultiZone
}

func (s *DescribeDBInstanceAttributeResponseBody) GetObjectStoreSize() *int64 {
	return s.ObjectStoreSize
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetResourceCpu() *int64 {
	return s.ResourceCpu
}

func (s *DescribeDBInstanceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetSecGroupConnValid() *string {
	return s.SecGroupConnValid
}

func (s *DescribeDBInstanceAttributeResponseBody) GetServerless() *bool {
	return s.Serverless
}

func (s *DescribeDBInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBody) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribeDBInstanceAttributeResponseBody) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDBInstanceAttributeResponseBody) GetTags() []*DescribeDBInstanceAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeDBInstanceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetVirtualClusterList() []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	return s.VirtualClusterList
}

func (s *DescribeDBInstanceAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCanUpgradeVersions(v []*string) *DescribeDBInstanceAttributeResponseBody {
	s.CanUpgradeVersions = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetConfigPatternType(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ConfigPatternType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBClusterList(v []*DescribeDBInstanceAttributeResponseBodyDBClusterList) *DescribeDBInstanceAttributeResponseBody {
	s.DBClusterList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDeployScheme(v string) *DescribeDBInstanceAttributeResponseBody {
	s.DeployScheme = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDescription(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngine(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngineMinorVersion(v string) *DescribeDBInstanceAttributeResponseBody {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetGmtModified(v string) *DescribeDBInstanceAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetLockMode(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetLockReason(v string) *DescribeDBInstanceAttributeResponseBody {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetMaintainEndtime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.MaintainEndtime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetMaintainStarttime(v string) *DescribeDBInstanceAttributeResponseBody {
	s.MaintainStarttime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetMultiZone(v []*DescribeDBInstanceAttributeResponseBodyMultiZone) *DescribeDBInstanceAttributeResponseBody {
	s.MultiZone = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetObjectStoreSize(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetResourceCpu(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.ResourceCpu = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetSecGroupConnValid(v string) *DescribeDBInstanceAttributeResponseBody {
	s.SecGroupConnValid = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetServerless(v bool) *DescribeDBInstanceAttributeResponseBody {
	s.Serverless = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetStatus(v string) *DescribeDBInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetStorageSize(v int64) *DescribeDBInstanceAttributeResponseBody {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetSubDomain(v string) *DescribeDBInstanceAttributeResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetTags(v []*DescribeDBInstanceAttributeResponseBodyTags) *DescribeDBInstanceAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetVirtualClusterList(v []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList) *DescribeDBInstanceAttributeResponseBody {
	s.VirtualClusterList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetVpcId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) Validate() error {
	if s.DBClusterList != nil {
		for _, item := range s.DBClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MultiZone != nil {
		for _, item := range s.MultiZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VirtualClusterList != nil {
		for _, item := range s.VirtualClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBClusterList struct {
	// The cache size. Unit: GB.
	//
	// example:
	//
	// 200
	CacheStorageSizeGB *string `json:"CacheStorageSizeGB,omitempty" xml:"CacheStorageSizeGB,omitempty"`
	// The cache type.
	//
	// example:
	//
	// cloud_essd
	CacheStorageType *string `json:"CacheStorageType,omitempty" xml:"CacheStorageType,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// selectdb-cn-7213cjv****-be
	ClusterBinding   *string `json:"ClusterBinding,omitempty" xml:"ClusterBinding,omitempty"`
	ClusterNodeCount *int32  `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	ClusterNodeType  *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 8
	CpuCores *int64 `json:"CpuCores,omitempty" xml:"CpuCores,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2023-08-14T09:24:13Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The specifications of the cluster. Valid values:
	//
	// 	- **selectdb.xlarge**: 4 CPU cores and 32 GB of memory.
	//
	// 	- **selectdb.2xlarge**: 8 CPU cores and 64 GB of memory.
	//
	// 	- **selectdb.4xlarge**: 16 CPU cores and 128 GB of memory.
	//
	// 	- **selectdb.8xlarge**: 32 CPU cores and 256 GB of memory.
	//
	// 	- **selectdb.16xlarge**: 64 CPU cores and 512 GB of memory.
	//
	// 	- **selectdb.24xlarge**: 96 CPU cores and 768 GB of memory.
	//
	// 	- **selectdb.32xlarge**: 128 CPU cores and 1,024 GB of memory.
	//
	// example:
	//
	// selectdb.2xlarge
	DbClusterClass *string `json:"DbClusterClass,omitempty" xml:"DbClusterClass,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// selectdb-cn-h033cjs****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test01
	DbClusterName *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test instance
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 64
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The modified time.
	//
	// example:
	//
	// 2024-07-02T16:35:44+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The performance level.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 4
	ScaleMax *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 0.5
	ScaleMin *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// false
	ScalingRulesEnable *bool `json:"ScalingRulesEnable,omitempty" xml:"ScalingRulesEnable,omitempty"`
	// The time when the cluster started.
	//
	// example:
	//
	// 2023-08-14T09:24:13Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- **CREATING**: The cluster is being created.
	//
	// 	- **ACTIVATION**: The cluster is running.
	//
	// 	- **RESOURCE_CHANGING**: The resource configuration of the cluster is being changed.
	//
	// 	- **ORDER_PREPARING**: The order is being confirmed.
	//
	// 	- **READONLY_RESOURCE_CHANGING**: The resource configuration of the cluster is being changed and the cluster is write-locked.
	//
	// 	- **DELETING**: The cluster is being deleted.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 预留参数，暂不返回。
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// example:
	//
	// vsw-t4n8x7jcc8rknon85tqoa
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetCacheStorageSizeGB() *string {
	return s.CacheStorageSizeGB
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetCacheStorageType() *string {
	return s.CacheStorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetClusterBinding() *string {
	return s.ClusterBinding
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetClusterNodeCount() *int32 {
	return s.ClusterNodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetClusterNodeType() *string {
	return s.ClusterNodeType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetCpuCores() *int64 {
	return s.CpuCores
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetDbClusterClass() *string {
	return s.DbClusterClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetDbClusterName() *string {
	return s.DbClusterName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetScalingRulesEnable() *bool {
	return s.ScalingRulesEnable
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCacheStorageSizeGB(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CacheStorageSizeGB = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCacheStorageType(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CacheStorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetClusterBinding(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ClusterBinding = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetClusterNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ClusterNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetClusterNodeType(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ClusterNodeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCpuCores(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetCreatedTime(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterClass(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterId(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbClusterName(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbClusterName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetDbInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetMemory(v int64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.Memory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetModifiedTime(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetPerformanceLevel(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetScaleMax(v float64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetScaleMin(v float64) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetScalingRulesEnable(v bool) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ScalingRulesEnable = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetSubDomain(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.SubDomain = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBClusterList {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBClusterList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyMultiZone struct {
	// example:
	//
	// 4096
	AvailableIpCount *int64 `json:"AvailableIpCount,omitempty" xml:"AvailableIpCount,omitempty"`
	// example:
	//
	// 113.88.14.211/32
	Cidr       *string   `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyMultiZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyMultiZone) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) GetAvailableIpCount() *int64 {
	return s.AvailableIpCount
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) GetCidr() *string {
	return s.Cidr
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) SetAvailableIpCount(v int64) *DescribeDBInstanceAttributeResponseBodyMultiZone {
	s.AvailableIpCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) SetCidr(v string) *DescribeDBInstanceAttributeResponseBodyMultiZone {
	s.Cidr = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) SetVSwitchIds(v []*string) *DescribeDBInstanceAttributeResponseBodyMultiZone {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyMultiZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyMultiZone) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDBInstanceAttributeResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDBInstanceAttributeResponseBodyTags) SetTagKey(v string) *DescribeDBInstanceAttributeResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyTags) SetTagValue(v string) *DescribeDBInstanceAttributeResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyVirtualClusterList struct {
	ActiveClusterId    *string `json:"ActiveClusterId,omitempty" xml:"ActiveClusterId,omitempty"`
	ActiveClusterName  *string `json:"ActiveClusterName,omitempty" xml:"ActiveClusterName,omitempty"`
	CreatedTime        *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DbClusterId        *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	DbClusterName      *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	StandbyClusterId   *string `json:"StandbyClusterId,omitempty" xml:"StandbyClusterId,omitempty"`
	StandbyClusterName *string `json:"StandbyClusterName,omitempty" xml:"StandbyClusterName,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyVirtualClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetActiveClusterId() *string {
	return s.ActiveClusterId
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetActiveClusterName() *string {
	return s.ActiveClusterName
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetDbClusterName() *string {
	return s.DbClusterName
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetStandbyClusterId() *string {
	return s.StandbyClusterId
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetStandbyClusterName() *string {
	return s.StandbyClusterName
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetActiveClusterId(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.ActiveClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetActiveClusterName(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.ActiveClusterName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetCreatedTime(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetDbClusterId(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetDbClusterName(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.DbClusterName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetStandbyClusterId(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.StandbyClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetStandbyClusterName(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.StandbyClusterName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyVirtualClusterList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyVirtualClusterList) Validate() error {
	return dara.Validate(s)
}
