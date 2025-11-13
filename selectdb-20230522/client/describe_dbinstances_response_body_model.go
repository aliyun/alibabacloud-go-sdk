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
	// The details about each instance returned.
	Items []*DescribeDBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
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
	// The total number of entries returned.
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
	// The edition of the instance. Default value: basic.
	//
	// example:
	//
	// basic
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
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
	DeployScheme *string `json:"DeployScheme,omitempty" xml:"DeployScheme,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// selectdb
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expires.
	//
	// >  A specific value is returned only for subscription clusters whose billing method is **Prepaid**. For pay-as-you-go clusters whose billing method is **Postpaid**, no value is returned.
	//
	// example:
	//
	// 2024-03-29T03:47:05Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-12T04:14Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the task was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-12T19:05Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// Instance
	InstanceUsedType *string `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	// Indicates whether the instance is deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// The end timestamp of the maintenance window.
	MaintainEndTimeStr *string `json:"MaintainEndTimeStr,omitempty" xml:"MaintainEndTimeStr,omitempty"`
	// The end time of the instance maintenance window.
	MaintainEndtime *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	// The start timestamp of the maintenance window.
	MaintainStartTimeStr *string `json:"MaintainStartTimeStr,omitempty" xml:"MaintainStartTimeStr,omitempty"`
	// The start time of the instance maintenance window.
	MaintainStarttime *string                                          `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	MultiZone         []*DescribeDBInstancesResponseBodyItemsMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	// The storage capacity of the instance. Unit: GB.
	//
	// example:
	//
	// 200
	ObjectStoreSize *int64 `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The time when the instance was created.
	ParentInstance *string `json:"ParentInstance,omitempty" xml:"ParentInstance,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of CPU cores of the instance.
	//
	// example:
	//
	// 8
	ResourceCpu *int64 `json:"ResourceCpu,omitempty" xml:"ResourceCpu,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The memory capacity of the instance.
	//
	// example:
	//
	// 64
	ResourceMemory *int64 `json:"ResourceMemory,omitempty" xml:"ResourceMemory,omitempty"`
	// The maximum number of RCUs.
	//
	// example:
	//
	// 0
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of RDS capacity units (RCUs).
	//
	// example:
	//
	// 0
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// This parameter is not returned.
	ScaleReplica *int64 `json:"ScaleReplica,omitempty" xml:"ScaleReplica,omitempty"`
	Serverless   *bool  `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
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
	// 100
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The storage type of the instance.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The details about each tag returned.
	Tags []*DescribeDBInstancesResponseBodyItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the cluster that is monitored by Managed Service for Prometheus.
	TenantClusterId *string `json:"TenantClusterId,omitempty" xml:"TenantClusterId,omitempty"`
	// The token that is used to access Managed Service for Prometheus.
	TenantToken *string `json:"TenantToken,omitempty" xml:"TenantToken,omitempty"`
	// The ID of the account that uses Managed Service for Prometheus.
	TenantUserId *string `json:"TenantUserId,omitempty" xml:"TenantUserId,omitempty"`
	// The virtual private cloud (VPC) ID.
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
	// The connection string of the instance.
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

type DescribeDBInstancesResponseBodyItemsMultiZone struct {
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	ZoneId     *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
