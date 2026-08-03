// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanUpgradeVersionCommunityMap(v map[string]*string) *DescribeDBInstanceAttributeResponseBody
	GetCanUpgradeVersionCommunityMap() map[string]*string
	SetCanUpgradeVersions(v []*string) *DescribeDBInstanceAttributeResponseBody
	GetCanUpgradeVersions() []*string
	SetChargeType(v string) *DescribeDBInstanceAttributeResponseBody
	GetChargeType() *string
	SetCommunityVersion(v string) *DescribeDBInstanceAttributeResponseBody
	GetCommunityVersion() *string
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
	SetFEClusterList(v []*DescribeDBInstanceAttributeResponseBodyFEClusterList) *DescribeDBInstanceAttributeResponseBody
	GetFEClusterList() []*DescribeDBInstanceAttributeResponseBodyFEClusterList
	SetGmtModified(v string) *DescribeDBInstanceAttributeResponseBody
	GetGmtModified() *string
	SetLangfuseInstanceIds(v []*string) *DescribeDBInstanceAttributeResponseBody
	GetLangfuseInstanceIds() []*string
	SetLockMode(v int64) *DescribeDBInstanceAttributeResponseBody
	GetLockMode() *int64
	SetLockReason(v string) *DescribeDBInstanceAttributeResponseBody
	GetLockReason() *string
	SetMCPServerServiceStatus(v string) *DescribeDBInstanceAttributeResponseBody
	GetMCPServerServiceStatus() *string
	SetMaintainEndtime(v string) *DescribeDBInstanceAttributeResponseBody
	GetMaintainEndtime() *string
	SetMaintainStarttime(v string) *DescribeDBInstanceAttributeResponseBody
	GetMaintainStarttime() *string
	SetMultiZone(v []*DescribeDBInstanceAttributeResponseBodyMultiZone) *DescribeDBInstanceAttributeResponseBody
	GetMultiZone() []*DescribeDBInstanceAttributeResponseBodyMultiZone
	SetOTelBearerToken(v string) *DescribeDBInstanceAttributeResponseBody
	GetOTelBearerToken() *string
	SetOTelGrafanaServiceStatus(v string) *DescribeDBInstanceAttributeResponseBody
	GetOTelGrafanaServiceStatus() *string
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
	CanUpgradeVersionCommunityMap map[string]*string `json:"CanUpgradeVersionCommunityMap,omitempty" xml:"CanUpgradeVersionCommunityMap,omitempty"`
	// The list of versions to which the instance can be upgraded.
	CanUpgradeVersions []*string `json:"CanUpgradeVersions,omitempty" xml:"CanUpgradeVersions,omitempty" type:"Repeated"`
	// The billing type of the instance. Valid values:
	//
	// example:
	//
	// Prepaid
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CommunityVersion *string `json:"CommunityVersion,omitempty" xml:"CommunityVersion,omitempty"`
	// The configuration template applied to the instance.
	//
	// example:
	//
	// log
	ConfigPatternType *string `json:"ConfigPatternType,omitempty" xml:"ConfigPatternType,omitempty"`
	// The creation time of the instance.
	//
	// example:
	//
	// 2023-08-14T03:00:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of clusters that belong to the instance.
	DBClusterList []*DescribeDBInstanceAttributeResponseBodyDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The deployment mode of the instance.
	//
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
	// The database engine type.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The minor engine version of the instance.
	//
	// example:
	//
	// 3.0.1
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The database engine version.
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
	ExpireTime    *string                                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FEClusterList []*DescribeDBInstanceAttributeResponseBodyFEClusterList `json:"FEClusterList,omitempty" xml:"FEClusterList,omitempty" type:"Repeated"`
	// The time when the instance was last modified (for example, restarted or had public network access enabled). The time is in the yyyy-MM-ddTHH:mmZ format (UTC).
	//
	// example:
	//
	// 2023-08-17T09:58Z
	GmtModified         *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LangfuseInstanceIds []*string `json:"LangfuseInstanceIds,omitempty" xml:"LangfuseInstanceIds,omitempty" type:"Repeated"`
	// The lock mode of the instance. The value is **lock**, which indicates that the instance is automatically expired or has an overdue payment.
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
	// example:
	//
	// not_enabled
	MCPServerServiceStatus *string `json:"MCPServerServiceStatus,omitempty" xml:"MCPServerServiceStatus,omitempty"`
	// The end time of the maintenance window of the instance.
	//
	// example:
	//
	// Reserved parameter.
	MaintainEndtime *string `json:"MaintainEndtime,omitempty" xml:"MaintainEndtime,omitempty"`
	// The start time of the maintenance window of the instance.
	//
	// example:
	//
	// Reserved parameter.
	MaintainStarttime *string `json:"MaintainStarttime,omitempty" xml:"MaintainStarttime,omitempty"`
	// The multi-zone configuration.
	//
	// if can be null:
	// true
	MultiZone                []*DescribeDBInstanceAttributeResponseBodyMultiZone `json:"MultiZone,omitempty" xml:"MultiZone,omitempty" type:"Repeated"`
	OTelBearerToken          *string                                             `json:"OTelBearerToken,omitempty" xml:"OTelBearerToken,omitempty"`
	OTelGrafanaServiceStatus *string                                             `json:"OTelGrafanaServiceStatus,omitempty" xml:"OTelGrafanaServiceStatus,omitempty"`
	// The storage space.
	//
	// example:
	//
	// 0
	ObjectStoreSize *int64 `json:"ObjectStoreSize,omitempty" xml:"ObjectStoreSize,omitempty"`
	// The region ID.
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
	// The number of CPU resources.
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
	// Indicates whether the direct port connection feature is enabled for the VPC in which the instance resides.
	//
	// example:
	//
	// false
	SecGroupConnValid *string `json:"SecGroupConnValid,omitempty" xml:"SecGroupConnValid,omitempty"`
	// Indicates whether the serverless feature is enabled for the instance.
	//
	// example:
	//
	// false
	Serverless *bool `json:"Serverless,omitempty" xml:"Serverless,omitempty"`
	// The status of the instance. Valid values:
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 400
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The zone.
	//
	// example:
	//
	// Reserved parameter.
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The list of instance labels.
	Tags []*DescribeDBInstanceAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp18iztwqrs8qj2nc6nyu
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of virtual clusters.
	VirtualClusterList []*DescribeDBInstanceAttributeResponseBodyVirtualClusterList `json:"VirtualClusterList,omitempty" xml:"VirtualClusterList,omitempty" type:"Repeated"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
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

func (s *DescribeDBInstanceAttributeResponseBody) GetCanUpgradeVersionCommunityMap() map[string]*string {
	return s.CanUpgradeVersionCommunityMap
}

func (s *DescribeDBInstanceAttributeResponseBody) GetCanUpgradeVersions() []*string {
	return s.CanUpgradeVersions
}

func (s *DescribeDBInstanceAttributeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstanceAttributeResponseBody) GetCommunityVersion() *string {
	return s.CommunityVersion
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

func (s *DescribeDBInstanceAttributeResponseBody) GetFEClusterList() []*DescribeDBInstanceAttributeResponseBodyFEClusterList {
	return s.FEClusterList
}

func (s *DescribeDBInstanceAttributeResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDBInstanceAttributeResponseBody) GetLangfuseInstanceIds() []*string {
	return s.LangfuseInstanceIds
}

func (s *DescribeDBInstanceAttributeResponseBody) GetLockMode() *int64 {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBody) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceAttributeResponseBody) GetMCPServerServiceStatus() *string {
	return s.MCPServerServiceStatus
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

func (s *DescribeDBInstanceAttributeResponseBody) GetOTelBearerToken() *string {
	return s.OTelBearerToken
}

func (s *DescribeDBInstanceAttributeResponseBody) GetOTelGrafanaServiceStatus() *string {
	return s.OTelGrafanaServiceStatus
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

func (s *DescribeDBInstanceAttributeResponseBody) SetCanUpgradeVersionCommunityMap(v map[string]*string) *DescribeDBInstanceAttributeResponseBody {
	s.CanUpgradeVersionCommunityMap = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCanUpgradeVersions(v []*string) *DescribeDBInstanceAttributeResponseBody {
	s.CanUpgradeVersions = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetChargeType(v string) *DescribeDBInstanceAttributeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetCommunityVersion(v string) *DescribeDBInstanceAttributeResponseBody {
	s.CommunityVersion = &v
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

func (s *DescribeDBInstanceAttributeResponseBody) SetFEClusterList(v []*DescribeDBInstanceAttributeResponseBodyFEClusterList) *DescribeDBInstanceAttributeResponseBody {
	s.FEClusterList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetGmtModified(v string) *DescribeDBInstanceAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetLangfuseInstanceIds(v []*string) *DescribeDBInstanceAttributeResponseBody {
	s.LangfuseInstanceIds = v
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

func (s *DescribeDBInstanceAttributeResponseBody) SetMCPServerServiceStatus(v string) *DescribeDBInstanceAttributeResponseBody {
	s.MCPServerServiceStatus = &v
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

func (s *DescribeDBInstanceAttributeResponseBody) SetOTelBearerToken(v string) *DescribeDBInstanceAttributeResponseBody {
	s.OTelBearerToken = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetOTelGrafanaServiceStatus(v string) *DescribeDBInstanceAttributeResponseBody {
	s.OTelGrafanaServiceStatus = &v
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
	// The cache storage size. Unit: GB.
	//
	// example:
	//
	// 200
	CacheStorageSizeGB *string `json:"CacheStorageSizeGB,omitempty" xml:"CacheStorageSizeGB,omitempty"`
	// The cache storage type.
	//
	// example:
	//
	// cloud_essd
	CacheStorageType *string `json:"CacheStorageType,omitempty" xml:"CacheStorageType,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The bound target cluster.
	//
	// example:
	//
	// selectdb-cn-7213cjv****-be
	ClusterBinding *string `json:"ClusterBinding,omitempty" xml:"ClusterBinding,omitempty"`
	// The number of cluster nodes. This parameter takes effect only in serverless mode.
	//
	// example:
	//
	// 1
	ClusterNodeCount *int32 `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	// The cluster node type. This parameter takes effect only in serverless mode.
	//
	// example:
	//
	// base
	ClusterNodeType *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
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
	// The cluster specifications. Valid values:
	//
	// example:
	//
	// selectdb.2xlarge
	DbClusterClass *string `json:"DbClusterClass,omitempty" xml:"DbClusterClass,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-h033cjs****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test01
	DbClusterName *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Instance test
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The memory size.
	//
	// example:
	//
	// 64
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-07-02T16:35:44+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The performance level (PL).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The maximum value of the automatic scaling range for the instance RCU (RDS Capacity Unit).
	//
	// example:
	//
	// 4
	ScaleMax *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum value of the automatic scaling range for the instance RCU (RDS Capacity Unit).
	//
	// example:
	//
	// 0.5
	ScaleMin *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// Indicates whether the time-based elastic policy is enabled.
	//
	// example:
	//
	// false
	ScalingRulesEnable *bool `json:"ScalingRulesEnable,omitempty" xml:"ScalingRulesEnable,omitempty"`
	// The start time of the cluster.
	//
	// example:
	//
	// 2023-08-14T09:24:13Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the cluster. Valid values:
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// Reserved parameter. Not returned
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-t4n8x7jcc8rknon85tqoa
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
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

type DescribeDBInstanceAttributeResponseBodyFEClusterList struct {
	DbClusterId          *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	NodeCount            *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	SingleNodeCpuCores   *int64  `json:"SingleNodeCpuCores,omitempty" xml:"SingleNodeCpuCores,omitempty"`
	SingleNodeMemoryInGB *int64  `json:"SingleNodeMemoryInGB,omitempty" xml:"SingleNodeMemoryInGB,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyFEClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyFEClusterList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) GetDbClusterId() *string {
	return s.DbClusterId
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) GetSingleNodeCpuCores() *int64 {
	return s.SingleNodeCpuCores
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) GetSingleNodeMemoryInGB() *int64 {
	return s.SingleNodeMemoryInGB
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) SetDbClusterId(v string) *DescribeDBInstanceAttributeResponseBodyFEClusterList {
	s.DbClusterId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) SetNodeCount(v int64) *DescribeDBInstanceAttributeResponseBodyFEClusterList {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) SetSingleNodeCpuCores(v int64) *DescribeDBInstanceAttributeResponseBodyFEClusterList {
	s.SingleNodeCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) SetSingleNodeMemoryInGB(v int64) *DescribeDBInstanceAttributeResponseBodyFEClusterList {
	s.SingleNodeMemoryInGB = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyFEClusterList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyFEClusterList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyMultiZone struct {
	// The number of available IP addresses in the zone.
	//
	// example:
	//
	// 4096
	AvailableIpCount *int64 `json:"AvailableIpCount,omitempty" xml:"AvailableIpCount,omitempty"`
	// The Classless Inter-Domain Routing block of the prefix list entry.
	//
	// example:
	//
	// 113.88.14.211/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone ID.
	//
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
	// The key of the tag.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
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
	// The ID of the primary cluster.
	//
	// example:
	//
	// selectdb-xx78***-be
	ActiveClusterId *string `json:"ActiveClusterId,omitempty" xml:"ActiveClusterId,omitempty"`
	// The name of the primary cluster.
	//
	// example:
	//
	// test1
	ActiveClusterName *string `json:"ActiveClusterName,omitempty" xml:"ActiveClusterName,omitempty"`
	// The creation time of the instance.
	//
	// example:
	//
	// 2025-05-31T21:01:09Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// selectdb-vcg-33cjs****-be
	DbClusterId *string `json:"DbClusterId,omitempty" xml:"DbClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// vgcdemo
	DbClusterName *string `json:"DbClusterName,omitempty" xml:"DbClusterName,omitempty"`
	// The ID of the standby cluster.
	//
	// example:
	//
	// selectdb-x6u7***-be
	StandbyClusterId *string `json:"StandbyClusterId,omitempty" xml:"StandbyClusterId,omitempty"`
	// The name of the standby cluster.
	//
	// example:
	//
	// test2
	StandbyClusterName *string `json:"StandbyClusterName,omitempty" xml:"StandbyClusterName,omitempty"`
	// The status of the instance. Valid values:
	//
	// example:
	//
	// UPDATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
