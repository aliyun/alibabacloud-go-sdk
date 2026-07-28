// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody
	GetItems() []*DescribeDBInstancesResponseBodyItems
	SetPageNumber(v int64) *DescribeDBInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDBInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDBInstancesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeDBInstancesResponseBody
	GetTotalRecordCount() *int64
}

type DescribeDBInstancesResponseBody struct {
	// The list of instance details.
	Items []*DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries to return per page. Valid values:
	//
	// - **30*	- (default value)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC854513-E85E-54F3-9842-B9CCD3308CDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) GetItems() []*DescribeDBInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDBInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancesResponseBody) SetItems(v []*DescribeDBInstancesResponseBodyItems) *DescribeDBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int64) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int64) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeDBInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) Validate() error {
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

type DescribeDBInstancesResponseBodyItems struct {
	// The instance edition. The default value is basic.
	//
	// example:
	//
	// basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go
	//
	// - **Prepaid**: subscription
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The total number of clusters.
	//
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The deployment mode of the instance:
	//
	// - multi_az: zone-redundant storage.
	//
	// - single_az: locally redundant storage.
	//
	// example:
	//
	// single_az
	DeployScheme *string `json:"DeployScheme,omitempty" xml:"DeployScheme,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// New instance test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The database type.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor engine version of the instance.
	//
	// example:
	//
	// 4.0.4
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The database version.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time of the cluster.
	//
	// > This parameter is returned only for **Prepaid*	- (subscription) clusters. For **Postpaid*	- (pay-as-you-go) clusters, this parameter is empty.
	//
	// example:
	//
	// 2024-03-29T03:47:05Z
	ExpireTime    *string                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FEClusterList []*DescribeDBInstancesResponseBodyItemsFEClusterList `json:"FEClusterList,omitempty" xml:"FEClusterList,omitempty" type:"Repeated"`
	// The time when the task was created (GMT).
	//
	// example:
	//
	// 2023-08-12T04:14Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the task was last modified (GMT).
	//
	// example:
	//
	// 2023-08-12T19:05Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance usage type.
	//
	// example:
	//
	// Instance
	InstanceUsedType *string `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	// Indicates whether the instance is deleted. Valid values:
	//
	// - **true**: The instance is deleted.
	//
	// - **false**: The instance is not deleted.
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The lock mode of the instance.
	//
	// example:
	//
	// 0
	LockMode *int64 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the instance is locked.
	//
	// example:
	//
	// nolock
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The timestamp that indicates the end of the maintenance window.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	MaintainEndTimeStr *string `json:"MaintainEndTimeStr,omitempty" xml:"MaintainEndTimeStr,omitempty"`
	// The end time of the maintenance window for the instance.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	MaintainEndtime *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	// The timestamp that indicates the start of the maintenance window.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	MaintainStartTimeStr *string `json:"MaintainStartTimeStr,omitempty" xml:"MaintainStartTimeStr,omitempty"`
	// The start time of the maintenance window for the instance.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	MaintainStarttime *string `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	// The multi-zone configuration.
	MultiZone []*DescribeDBInstancesResponseBodyItemsMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// The instance storage size. Unit: GB.
	//
	// example:
	//
	// 200
	ObjectStoreSize *int64 `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The creation time.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	ParentInstance *string `json:"ParentInstance,omitempty" xml:"ParentInstance,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The allocated CPU for the resource.
	//
	// example:
	//
	// 8
	ResourceCpu *int64 `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 64
	ResourceMemory *int64 `json:"ResourceMemory,omitempty" xml:"ResourceMemory,omitempty"`
	// The maximum number of RDS Capacity Units (RCUs) for the instance.
	//
	// example:
	//
	// 0
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of RDS Capacity Units (RCUs) for the instance.
	//
	// example:
	//
	// 0
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// This field is redundant.
	//
	// example:
	//
	// 空
	ScaleReplica *int64 `json:"ScaleReplica,omitempty" xml:"ScaleReplica,omitempty"`
	// Indicates whether the instance is a serverless instance.
	//
	// example:
	//
	// false
	Serverless *bool `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// The state of the instance. Valid values:
	//
	// - **CREATING**: The instance is being created.
	//
	// - **ACTIVATION**: The instance is running.
	//
	// - **RESOURCE_CHANGING**: The instance is being upgraded or downgraded.
	//
	// - **ORDER_PREPARING**: The order is being confirmed.
	//
	// - **READONLY_RESOURCE_CHANGING**: The instance configuration is being changed, and the instance is write-locked.
	//
	// - **DELETING**: The instance is being deleted.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 100
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage class of the instance.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The list of tags of the instance.
	Tags []*DescribeDBInstancesResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Prometheus monitoring cluster.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	TenantClusterId *string `json:"TenantClusterId,omitempty" xml:"TenantClusterId,omitempty"`
	// The token for connecting to Prometheus monitoring.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	TenantToken *string `json:"TenantToken,omitempty" xml:"TenantToken,omitempty"`
	// The user account label for Prometheus monitoring.
	//
	// example:
	//
	// Reserved parameter. Not returned.
	TenantUserId *string `json:"TenantUserId,omitempty" xml:"TenantUserId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The connection address.
	//
	// example:
	//
	// Not applicable.
	ConnectionString *string `json:"connectionString,omitempty" xml:"connectionString,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstancesResponseBodyItems) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesResponseBodyItems) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *DescribeDBInstancesResponseBodyItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesResponseBodyItems) GetDeployScheme() *string {
	return s.DeployScheme
}

func (s *DescribeDBInstancesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstancesResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyItems) GetEngineMinorVersion() *string {
	return s.EngineMinorVersion
}

func (s *DescribeDBInstancesResponseBodyItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesResponseBodyItems) GetFEClusterList() []*DescribeDBInstancesResponseBodyItemsFEClusterList {
	return s.FEClusterList
}

func (s *DescribeDBInstancesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeDBInstancesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDBInstancesResponseBodyItems) GetInstanceUsedType() *string {
	return s.InstanceUsedType
}

func (s *DescribeDBInstancesResponseBodyItems) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *DescribeDBInstancesResponseBodyItems) GetLockMode() *int64 {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyItems) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesResponseBodyItems) GetMaintainEndTimeStr() *string {
	return s.MaintainEndTimeStr
}

func (s *DescribeDBInstancesResponseBodyItems) GetMaintainEndtime() *string {
	return s.MaintainEndtime
}

func (s *DescribeDBInstancesResponseBodyItems) GetMaintainStartTimeStr() *string {
	return s.MaintainStartTimeStr
}

func (s *DescribeDBInstancesResponseBodyItems) GetMaintainStarttime() *string {
	return s.MaintainStarttime
}

func (s *DescribeDBInstancesResponseBodyItems) GetMultiZone() []*DescribeDBInstancesResponseBodyItemsMultiZone {
	return s.MultiZone
}

func (s *DescribeDBInstancesResponseBodyItems) GetObjectStoreSize() *int64 {
	return s.ObjectStoreSize
}

func (s *DescribeDBInstancesResponseBodyItems) GetParentInstance() *string {
	return s.ParentInstance
}

func (s *DescribeDBInstancesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyItems) GetResourceCpu() *int64 {
	return s.ResourceCpu
}

func (s *DescribeDBInstancesResponseBodyItems) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyItems) GetResourceMemory() *int64 {
	return s.ResourceMemory
}

func (s *DescribeDBInstancesResponseBodyItems) GetScaleMax() *int64 {
	return s.ScaleMax
}

func (s *DescribeDBInstancesResponseBodyItems) GetScaleMin() *int64 {
	return s.ScaleMin
}

func (s *DescribeDBInstancesResponseBodyItems) GetScaleReplica() *int64 {
	return s.ScaleReplica
}

func (s *DescribeDBInstancesResponseBodyItems) GetServerless() *bool {
	return s.Serverless
}

func (s *DescribeDBInstancesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstancesResponseBodyItems) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribeDBInstancesResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstancesResponseBodyItems) GetTags() []*DescribeDBInstancesResponseBodyItemsTags {
	return s.Tags
}

func (s *DescribeDBInstancesResponseBodyItems) GetTenantClusterId() *string {
	return s.TenantClusterId
}

func (s *DescribeDBInstancesResponseBodyItems) GetTenantToken() *string {
	return s.TenantToken
}

func (s *DescribeDBInstancesResponseBodyItems) GetTenantUserId() *string {
	return s.TenantUserId
}

func (s *DescribeDBInstancesResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesResponseBodyItems) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDBInstancesResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstancesResponseBodyItems) SetCategory(v string) *DescribeDBInstancesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetChargeType(v string) *DescribeDBInstancesResponseBodyItems {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetClusterCount(v int32) *DescribeDBInstancesResponseBodyItems {
	s.ClusterCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDBInstanceId(v string) *DescribeDBInstancesResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDeployScheme(v string) *DescribeDBInstancesResponseBodyItems {
	s.DeployScheme = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetDescription(v string) *DescribeDBInstancesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetEngine(v string) *DescribeDBInstancesResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetEngineMinorVersion(v string) *DescribeDBInstancesResponseBodyItems {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetExpireTime(v string) *DescribeDBInstancesResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetFEClusterList(v []*DescribeDBInstancesResponseBodyItemsFEClusterList) *DescribeDBInstancesResponseBodyItems {
	s.FEClusterList = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetGmtCreated(v string) *DescribeDBInstancesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetGmtModified(v string) *DescribeDBInstancesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetInstanceUsedType(v string) *DescribeDBInstancesResponseBodyItems {
	s.InstanceUsedType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetIsDeleted(v bool) *DescribeDBInstancesResponseBodyItems {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetLockMode(v int64) *DescribeDBInstancesResponseBodyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetLockReason(v string) *DescribeDBInstancesResponseBodyItems {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainEndTimeStr(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainEndTimeStr = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainEndtime(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainEndtime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainStartTimeStr(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainStartTimeStr = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMaintainStarttime(v string) *DescribeDBInstancesResponseBodyItems {
	s.MaintainStarttime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetMultiZone(v []*DescribeDBInstancesResponseBodyItemsMultiZone) *DescribeDBInstancesResponseBodyItems {
	s.MultiZone = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetObjectStoreSize(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ObjectStoreSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetParentInstance(v string) *DescribeDBInstancesResponseBodyItems {
	s.ParentInstance = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetRegionId(v string) *DescribeDBInstancesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceCpu(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ResourceCpu = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyItems {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetResourceMemory(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ResourceMemory = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleMax(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleMax = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleMin(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleMin = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetScaleReplica(v int64) *DescribeDBInstancesResponseBodyItems {
	s.ScaleReplica = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetServerless(v bool) *DescribeDBInstancesResponseBodyItems {
	s.Serverless = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStatus(v string) *DescribeDBInstancesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStorageSize(v int64) *DescribeDBInstancesResponseBodyItems {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetStorageType(v string) *DescribeDBInstancesResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTags(v []*DescribeDBInstancesResponseBodyItemsTags) *DescribeDBInstancesResponseBodyItems {
	s.Tags = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantClusterId(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantClusterId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantToken(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantToken = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetTenantUserId(v string) *DescribeDBInstancesResponseBodyItems {
	s.TenantUserId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetVpcId(v string) *DescribeDBInstancesResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetVswitchId(v string) *DescribeDBInstancesResponseBodyItems {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetZoneId(v string) *DescribeDBInstancesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) SetConnectionString(v string) *DescribeDBInstancesResponseBodyItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItems) Validate() error {
	if s.FEClusterList != nil {
		for _, item := range s.FEClusterList {
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
	return nil
}

type DescribeDBInstancesResponseBodyItemsFEClusterList struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	NodeCount            *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	SingleNodeCpuCores   *int64  `json:"SingleNodeCpuCores,omitempty" xml:"SingleNodeCpuCores,omitempty"`
	SingleNodeMemoryInGB *int64  `json:"SingleNodeMemoryInGB,omitempty" xml:"SingleNodeMemoryInGB,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsFEClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsFEClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) GetSingleNodeCpuCores() *int64 {
	return s.SingleNodeCpuCores
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) GetSingleNodeMemoryInGB() *int64 {
	return s.SingleNodeMemoryInGB
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) SetDbClusterId(v string) *DescribeDBInstancesResponseBodyItemsFEClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) SetNodeCount(v int64) *DescribeDBInstancesResponseBodyItemsFEClusterList {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) SetSingleNodeCpuCores(v int64) *DescribeDBInstancesResponseBodyItemsFEClusterList {
	s.SingleNodeCpuCores = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) SetSingleNodeMemoryInGB(v int64) *DescribeDBInstancesResponseBodyItemsFEClusterList {
	s.SingleNodeMemoryInGB = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) SetStatus(v string) *DescribeDBInstancesResponseBodyItemsFEClusterList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsFEClusterList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyItemsMultiZone struct {
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyItemsMultiZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsMultiZone) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsMultiZone) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeDBInstancesResponseBodyItemsMultiZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyItemsMultiZone) SetVSwitchIds(v []*string) *DescribeDBInstancesResponseBodyItemsMultiZone {
	s.VSwitchIds = v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsMultiZone) SetZoneId(v string) *DescribeDBInstancesResponseBodyItemsMultiZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsMultiZone) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyItemsTags struct {
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

func (s DescribeDBInstancesResponseBodyItemsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyItemsTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyItemsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDBInstancesResponseBodyItemsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDBInstancesResponseBodyItemsTags) SetTagKey(v string) *DescribeDBInstancesResponseBodyItemsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsTags) SetTagValue(v string) *DescribeDBInstancesResponseBodyItemsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyItemsTags) Validate() error {
	return dara.Validate(s)
}
