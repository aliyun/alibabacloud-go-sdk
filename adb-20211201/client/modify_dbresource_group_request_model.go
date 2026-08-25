// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAtmConfig(v *ModifyDBResourceGroupRequestAtmConfig) *ModifyDBResourceGroupRequest
	GetAtmConfig() *ModifyDBResourceGroupRequestAtmConfig
	SetAutoStopInterval(v string) *ModifyDBResourceGroupRequest
	GetAutoStopInterval() *string
	SetClusterMode(v string) *ModifyDBResourceGroupRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *ModifyDBResourceGroupRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *ModifyDBResourceGroupRequest
	GetDBClusterId() *string
	SetEnableSpot(v bool) *ModifyDBResourceGroupRequest
	GetEnableSpot() *bool
	SetEngineParams(v map[string]interface{}) *ModifyDBResourceGroupRequest
	GetEngineParams() map[string]interface{}
	SetGpuElasticPlan(v *ModifyDBResourceGroupRequestGpuElasticPlan) *ModifyDBResourceGroupRequest
	GetGpuElasticPlan() *ModifyDBResourceGroupRequestGpuElasticPlan
	SetGroupName(v string) *ModifyDBResourceGroupRequest
	GetGroupName() *string
	SetGroupType(v string) *ModifyDBResourceGroupRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *ModifyDBResourceGroupRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *ModifyDBResourceGroupRequest
	GetMaxComputeResource() *string
	SetMaxGpuQuantity(v int32) *ModifyDBResourceGroupRequest
	GetMaxGpuQuantity() *int32
	SetMinClusterCount(v int32) *ModifyDBResourceGroupRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *ModifyDBResourceGroupRequest
	GetMinComputeResource() *string
	SetMinGpuQuantity(v int32) *ModifyDBResourceGroupRequest
	GetMinGpuQuantity() *int32
	SetRayConfig(v *ModifyDBResourceGroupRequestRayConfig) *ModifyDBResourceGroupRequest
	GetRayConfig() *ModifyDBResourceGroupRequestRayConfig
	SetRegionId(v string) *ModifyDBResourceGroupRequest
	GetRegionId() *string
	SetRules(v []*ModifyDBResourceGroupRequestRules) *ModifyDBResourceGroupRequest
	GetRules() []*ModifyDBResourceGroupRequestRules
	SetSpecName(v string) *ModifyDBResourceGroupRequest
	GetSpecName() *string
	SetStatus(v string) *ModifyDBResourceGroupRequest
	GetStatus() *string
	SetTargetResourceGroupName(v string) *ModifyDBResourceGroupRequest
	GetTargetResourceGroupName() *string
}

type ModifyDBResourceGroupRequest struct {
	// The PromQL resource group configuration.
	AtmConfig *ModifyDBResourceGroupRequestAtmConfig `json:"AtmConfig,omitempty" xml:"AtmConfig,omitempty" type:"Struct"`
	// The automatic stop interval.
	//
	// example:
	//
	// 5m
	AutoStopInterval *string `json:"AutoStopInterval,omitempty" xml:"AutoStopInterval,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// 无
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// 无
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// <props="china">The cluster ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The cluster ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. After the spot instance feature is enabled, the unit price of resources is reduced, but the resources may be released. Only Job resource groups support this feature. Valid values:
	//
	// - **True**: Enables the spot instance feature.
	//
	// - **False**: Disables the spot instance feature.
	//
	// example:
	//
	// True
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The engine configuration.
	//
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The GPU time-sharing elastic plan.
	GpuElasticPlan *ModifyDBResourceGroupRequestGpuElasticPlan `json:"GpuElasticPlan,omitempty" xml:"GpuElasticPlan,omitempty" type:"Struct"`
	// The resource group name.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group names of a specified cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The resource group type. Valid values:
	//
	// - **Interactive**
	//
	// - **Job**
	//
	// > For more information about Data Lakehouse Edition resource groups, see [Resource group overview](https://help.aliyun.com/document_detail/428610.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Interactive
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// 无
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum reserved computing resources.
	//
	// - If the resource group type is Interactive, the maximum reserved computing resources is the unallocated resources of the cluster, in increments of 16 ACUs.
	//
	// - If the resource group type is Job, the maximum reserved computing resources is the unallocated resources of the cluster, in increments of 8 ACUs.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	MaxGpuQuantity *int32 `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// 无
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum reserved computing resources.
	//
	// - If the resource group type is Interactive, the minimum reserved computing resources is 16 ACUs.
	//
	// - If the resource group type is Job, the minimum reserved computing resources is 0 ACUs.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	MinGpuQuantity *int32 `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	// The Ray configuration. This parameter is required when the resource group is an AI resource group and the corresponding engine is RayCluster.
	RayConfig *ModifyDBResourceGroupRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID of a specified cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job routing rules.
	Rules []*ModifyDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// The resource group status. **starting*	- indicates that the resource group is being started. **stopping*	- indicates that the resource group is being stopped.
	//
	// example:
	//
	// starting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequest) GetAtmConfig() *ModifyDBResourceGroupRequestAtmConfig {
	return s.AtmConfig
}

func (s *ModifyDBResourceGroupRequest) GetAutoStopInterval() *string {
	return s.AutoStopInterval
}

func (s *ModifyDBResourceGroupRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *ModifyDBResourceGroupRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *ModifyDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBResourceGroupRequest) GetEnableSpot() *bool {
	return s.EnableSpot
}

func (s *ModifyDBResourceGroupRequest) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *ModifyDBResourceGroupRequest) GetGpuElasticPlan() *ModifyDBResourceGroupRequestGpuElasticPlan {
	return s.GpuElasticPlan
}

func (s *ModifyDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ModifyDBResourceGroupRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *ModifyDBResourceGroupRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *ModifyDBResourceGroupRequest) GetMaxGpuQuantity() *int32 {
	return s.MaxGpuQuantity
}

func (s *ModifyDBResourceGroupRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *ModifyDBResourceGroupRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *ModifyDBResourceGroupRequest) GetMinGpuQuantity() *int32 {
	return s.MinGpuQuantity
}

func (s *ModifyDBResourceGroupRequest) GetRayConfig() *ModifyDBResourceGroupRequestRayConfig {
	return s.RayConfig
}

func (s *ModifyDBResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBResourceGroupRequest) GetRules() []*ModifyDBResourceGroupRequestRules {
	return s.Rules
}

func (s *ModifyDBResourceGroupRequest) GetSpecName() *string {
	return s.SpecName
}

func (s *ModifyDBResourceGroupRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyDBResourceGroupRequest) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *ModifyDBResourceGroupRequest) SetAtmConfig(v *ModifyDBResourceGroupRequestAtmConfig) *ModifyDBResourceGroupRequest {
	s.AtmConfig = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetAutoStopInterval(v string) *ModifyDBResourceGroupRequest {
	s.AutoStopInterval = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetClusterMode(v string) *ModifyDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetDBClusterId(v string) *ModifyDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetEnableSpot(v bool) *ModifyDBResourceGroupRequest {
	s.EnableSpot = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetEngineParams(v map[string]interface{}) *ModifyDBResourceGroupRequest {
	s.EngineParams = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGpuElasticPlan(v *ModifyDBResourceGroupRequestGpuElasticPlan) *ModifyDBResourceGroupRequest {
	s.GpuElasticPlan = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupName(v string) *ModifyDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetGroupType(v string) *ModifyDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMaxGpuQuantity(v int32) *ModifyDBResourceGroupRequest {
	s.MaxGpuQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetMinGpuQuantity(v int32) *ModifyDBResourceGroupRequest {
	s.MinGpuQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetRayConfig(v *ModifyDBResourceGroupRequestRayConfig) *ModifyDBResourceGroupRequest {
	s.RayConfig = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetRegionId(v string) *ModifyDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetRules(v []*ModifyDBResourceGroupRequestRules) *ModifyDBResourceGroupRequest {
	s.Rules = v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetSpecName(v string) *ModifyDBResourceGroupRequest {
	s.SpecName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetStatus(v string) *ModifyDBResourceGroupRequest {
	s.Status = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) SetTargetResourceGroupName(v string) *ModifyDBResourceGroupRequest {
	s.TargetResourceGroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequest) Validate() error {
	if s.AtmConfig != nil {
		if err := s.AtmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.GpuElasticPlan != nil {
		if err := s.GpuElasticPlan.Validate(); err != nil {
			return err
		}
	}
	if s.RayConfig != nil {
		if err := s.RayConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBResourceGroupRequestAtmConfig struct {
	// The number of authentication nodes.
	//
	// example:
	//
	// 2
	AuthNodeNum *int32 `json:"AuthNodeNum,omitempty" xml:"AuthNodeNum,omitempty"`
	// The authentication node specifications in ACU ([0-9+]ACU).
	//
	// example:
	//
	// 8ACU
	AuthNodeSpec *string `json:"AuthNodeSpec,omitempty" xml:"AuthNodeSpec,omitempty"`
	// The number of insert nodes.
	//
	// example:
	//
	// 2
	InsertNodeNum *int32 `json:"InsertNodeNum,omitempty" xml:"InsertNodeNum,omitempty"`
	// The insert node specifications in ACU ([0-9+]ACU).
	//
	// example:
	//
	// 8ACU
	InsertNodeSpec *string `json:"InsertNodeSpec,omitempty" xml:"InsertNodeSpec,omitempty"`
	// The query node cache size in GB.
	//
	// example:
	//
	// 10
	SelectNodeCacheSize *int32 `json:"SelectNodeCacheSize,omitempty" xml:"SelectNodeCacheSize,omitempty"`
	// The number of query nodes.
	//
	// example:
	//
	// 1
	SelectNodeNum *int32 `json:"SelectNodeNum,omitempty" xml:"SelectNodeNum,omitempty"`
	// The query node specifications ([0-9+]ACU).
	//
	// example:
	//
	// 8ACU
	SelectNodeSpec *string `json:"SelectNodeSpec,omitempty" xml:"SelectNodeSpec,omitempty"`
	// The disk size of storage nodes.
	//
	// example:
	//
	// 1
	StorageNodeDiskSize *int32 `json:"StorageNodeDiskSize,omitempty" xml:"StorageNodeDiskSize,omitempty"`
	// The disk type of storage nodes (essd_pl1, essd_pl2).
	//
	// example:
	//
	// essd_pl1
	StorageNodeDiskType *string `json:"StorageNodeDiskType,omitempty" xml:"StorageNodeDiskType,omitempty"`
	// The number of storage nodes.
	//
	// example:
	//
	// 2
	StorageNodeNum *int32 `json:"StorageNodeNum,omitempty" xml:"StorageNodeNum,omitempty"`
	// The storage node specifications in ACU ([0-9+]ACU).
	//
	// example:
	//
	// 8ACU
	StorageNodeSpec *string `json:"StorageNodeSpec,omitempty" xml:"StorageNodeSpec,omitempty"`
}

func (s ModifyDBResourceGroupRequestAtmConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestAtmConfig) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetAuthNodeNum() *int32 {
	return s.AuthNodeNum
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetAuthNodeSpec() *string {
	return s.AuthNodeSpec
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetInsertNodeNum() *int32 {
	return s.InsertNodeNum
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetInsertNodeSpec() *string {
	return s.InsertNodeSpec
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetSelectNodeCacheSize() *int32 {
	return s.SelectNodeCacheSize
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetSelectNodeNum() *int32 {
	return s.SelectNodeNum
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetSelectNodeSpec() *string {
	return s.SelectNodeSpec
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetStorageNodeDiskSize() *int32 {
	return s.StorageNodeDiskSize
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetStorageNodeDiskType() *string {
	return s.StorageNodeDiskType
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetStorageNodeNum() *int32 {
	return s.StorageNodeNum
}

func (s *ModifyDBResourceGroupRequestAtmConfig) GetStorageNodeSpec() *string {
	return s.StorageNodeSpec
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetAuthNodeNum(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.AuthNodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetAuthNodeSpec(v string) *ModifyDBResourceGroupRequestAtmConfig {
	s.AuthNodeSpec = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetInsertNodeNum(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.InsertNodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetInsertNodeSpec(v string) *ModifyDBResourceGroupRequestAtmConfig {
	s.InsertNodeSpec = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetSelectNodeCacheSize(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.SelectNodeCacheSize = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetSelectNodeNum(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.SelectNodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetSelectNodeSpec(v string) *ModifyDBResourceGroupRequestAtmConfig {
	s.SelectNodeSpec = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetStorageNodeDiskSize(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.StorageNodeDiskSize = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetStorageNodeDiskType(v string) *ModifyDBResourceGroupRequestAtmConfig {
	s.StorageNodeDiskType = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetStorageNodeNum(v int32) *ModifyDBResourceGroupRequestAtmConfig {
	s.StorageNodeNum = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) SetStorageNodeSpec(v string) *ModifyDBResourceGroupRequestAtmConfig {
	s.StorageNodeSpec = &v
	return s
}

func (s *ModifyDBResourceGroupRequestAtmConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestGpuElasticPlan struct {
	// Specifies whether to enable the elastic plan immediately after creation. Valid values:
	//
	// - **true**: Enables the elastic plan immediately.
	//
	// - **false**: Does not enable the elastic plan.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of rules.
	Rules []*ModifyDBResourceGroupRequestGpuElasticPlanRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ModifyDBResourceGroupRequestGpuElasticPlan) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestGpuElasticPlan) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlan) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlan) GetRules() []*ModifyDBResourceGroupRequestGpuElasticPlanRules {
	return s.Rules
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlan) SetEnabled(v bool) *ModifyDBResourceGroupRequestGpuElasticPlan {
	s.Enabled = &v
	return s
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlan) SetRules(v []*ModifyDBResourceGroupRequestGpuElasticPlanRules) *ModifyDBResourceGroupRequestGpuElasticPlan {
	s.Rules = v
	return s
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlan) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBResourceGroupRequestGpuElasticPlanRules struct {
	// The end time, specified as a cron expression. The interval must be at least 1 hour.
	//
	// example:
	//
	// 0 0 3 	- 	- ?
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	// The start time, specified as a cron expression. The interval must be at least 1 hour.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	StartCronExpression *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
}

func (s ModifyDBResourceGroupRequestGpuElasticPlanRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestGpuElasticPlanRules) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlanRules) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlanRules) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlanRules) SetEndCronExpression(v string) *ModifyDBResourceGroupRequestGpuElasticPlanRules {
	s.EndCronExpression = &v
	return s
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlanRules) SetStartCronExpression(v string) *ModifyDBResourceGroupRequestGpuElasticPlanRules {
	s.StartCronExpression = &v
	return s
}

func (s *ModifyDBResourceGroupRequestGpuElasticPlanRules) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestRayConfig struct {
	// The Ray application configuration.
	AppConfig *ModifyDBResourceGroupRequestRayConfigAppConfig `json:"AppConfig,omitempty" xml:"AppConfig,omitempty" type:"Struct"`
	// The Ray cluster type. Valid values:
	//
	// - BASIC: basic type, non-high-availability
	//
	// - HIGH_AVAILABILITY: high-availability type
	//
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to enable ENI.
	EnableUserEni *bool `json:"EnableUserEni,omitempty" xml:"EnableUserEni,omitempty"`
	// The allocation unit of the head node.
	//
	// example:
	//
	// 1
	HeadAllocateUnit *string `json:"HeadAllocateUnit,omitempty" xml:"HeadAllocateUnit,omitempty"`
	// The disk size of the head node.
	//
	// example:
	//
	// 100G
	HeadDiskCapacity *string `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	// The node specifications of the head node.
	//
	// example:
	//
	// xlarge
	HeadSpec *string `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	// The resource type of the head node.
	//
	// example:
	//
	// CPU
	HeadSpecType *string `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	// The storage mount list.
	StorageMounts           []*ModifyDBResourceGroupRequestRayConfigStorageMounts `json:"StorageMounts,omitempty" xml:"StorageMounts,omitempty" type:"Repeated"`
	UserDefinedRequirements *string                                               `json:"UserDefinedRequirements,omitempty" xml:"UserDefinedRequirements,omitempty"`
	// The list of Ray worker group configurations.
	WorkerGroups []*ModifyDBResourceGroupRequestRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s ModifyDBResourceGroupRequestRayConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRayConfig) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetAppConfig() *ModifyDBResourceGroupRequestRayConfigAppConfig {
	return s.AppConfig
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetEnableUserEni() *bool {
	return s.EnableUserEni
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetHeadAllocateUnit() *string {
	return s.HeadAllocateUnit
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetHeadDiskCapacity() *string {
	return s.HeadDiskCapacity
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetHeadSpecType() *string {
	return s.HeadSpecType
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetStorageMounts() []*ModifyDBResourceGroupRequestRayConfigStorageMounts {
	return s.StorageMounts
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetUserDefinedRequirements() *string {
	return s.UserDefinedRequirements
}

func (s *ModifyDBResourceGroupRequestRayConfig) GetWorkerGroups() []*ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetAppConfig(v *ModifyDBResourceGroupRequestRayConfigAppConfig) *ModifyDBResourceGroupRequestRayConfig {
	s.AppConfig = v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetCategory(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.Category = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetEnableUserEni(v bool) *ModifyDBResourceGroupRequestRayConfig {
	s.EnableUserEni = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetHeadAllocateUnit(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.HeadAllocateUnit = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetHeadDiskCapacity(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.HeadDiskCapacity = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetHeadSpec(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetHeadSpecType(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.HeadSpecType = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetStorageMounts(v []*ModifyDBResourceGroupRequestRayConfigStorageMounts) *ModifyDBResourceGroupRequestRayConfig {
	s.StorageMounts = v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetUserDefinedRequirements(v string) *ModifyDBResourceGroupRequestRayConfig {
	s.UserDefinedRequirements = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) SetWorkerGroups(v []*ModifyDBResourceGroupRequestRayConfigWorkerGroups) *ModifyDBResourceGroupRequestRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfig) Validate() error {
	if s.AppConfig != nil {
		if err := s.AppConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StorageMounts != nil {
		for _, item := range s.StorageMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkerGroups != nil {
		for _, item := range s.WorkerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBResourceGroupRequestRayConfigAppConfig struct {
	// The application name.
	//
	// example:
	//
	// app01
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// IsaacLab
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The image configuration.
	ImageSelector *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector `json:"ImageSelector,omitempty" xml:"ImageSelector,omitempty" type:"Struct"`
}

func (s ModifyDBResourceGroupRequestRayConfigAppConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRayConfigAppConfig) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) GetAppName() *string {
	return s.AppName
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) GetAppType() *string {
	return s.AppType
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) GetImageSelector() *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector {
	return s.ImageSelector
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) SetAppName(v string) *ModifyDBResourceGroupRequestRayConfigAppConfig {
	s.AppName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) SetAppType(v string) *ModifyDBResourceGroupRequestRayConfigAppConfig {
	s.AppType = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) SetImageSelector(v *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) *ModifyDBResourceGroupRequestRayConfigAppConfig {
	s.ImageSelector = v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfig) Validate() error {
	if s.ImageSelector != nil {
		if err := s.ImageSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector struct {
	// The image name.
	//
	// example:
	//
	// lab2.10.0-ray2.43.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The inference engine.
	//
	// example:
	//
	// vLLM
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
	// The LLM model.
	//
	// example:
	//
	// Deepseek-R1
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
}

func (s ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) GetImage() *string {
	return s.Image
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) GetLlmModel() *string {
	return s.LlmModel
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) SetImage(v string) *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector {
	s.Image = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) SetInferenceEngine(v string) *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector {
	s.InferenceEngine = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) SetLlmModel(v string) *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector {
	s.LlmModel = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigAppConfigImageSelector) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestRayConfigStorageMounts struct {
	// The mount path.
	//
	// example:
	//
	// /mnt/data01
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The storage ID.
	//
	// example:
	//
	// 1
	StorageId   *int64  `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
	StorageName *string `json:"StorageName,omitempty" xml:"StorageName,omitempty"`
}

func (s ModifyDBResourceGroupRequestRayConfigStorageMounts) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRayConfigStorageMounts) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) GetStorageId() *int64 {
	return s.StorageId
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) GetStorageName() *string {
	return s.StorageName
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) SetMountPath(v string) *ModifyDBResourceGroupRequestRayConfigStorageMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) SetStorageId(v int64) *ModifyDBResourceGroupRequestRayConfigStorageMounts {
	s.StorageId = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) SetStorageName(v string) *ModifyDBResourceGroupRequestRayConfigStorageMounts {
	s.StorageName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestRayConfigWorkerGroups struct {
	// The allocation unit.
	//
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// The worker group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The maximum number of workers.
	//
	// example:
	//
	// 2
	MaxWorkerQuantity *int32 `json:"MaxWorkerQuantity,omitempty" xml:"MaxWorkerQuantity,omitempty"`
	// The minimum number of workers.
	//
	// example:
	//
	// 1
	MinWorkerQuantity *int32 `json:"MinWorkerQuantity,omitempty" xml:"MinWorkerQuantity,omitempty"`
	// The disk size of the worker node.
	//
	// example:
	//
	// 100G
	WorkerDiskCapacity *string `json:"WorkerDiskCapacity,omitempty" xml:"WorkerDiskCapacity,omitempty"`
	// The node specifications of the worker node.
	//
	// example:
	//
	// xlarge
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
	// The resource type of the worker node.
	//
	// example:
	//
	// GPU
	WorkerSpecType *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
}

func (s ModifyDBResourceGroupRequestRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetAllocateUnit(v string) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetGroupName(v string) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerSpecName(v string) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerSpecType(v string) *ModifyDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestRules struct {
	// The resource group name.
	//
	// example:
	//
	// user_default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution time threshold. Unit: milliseconds (ms).
	//
	// example:
	//
	// 180000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The target resource group name.
	//
	// example:
	//
	// job
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s ModifyDBResourceGroupRequestRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupRequestRules) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupRequestRules) GetQueryTime() *string {
	return s.QueryTime
}

func (s *ModifyDBResourceGroupRequestRules) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *ModifyDBResourceGroupRequestRules) SetGroupName(v string) *ModifyDBResourceGroupRequestRules {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRules) SetQueryTime(v string) *ModifyDBResourceGroupRequestRules {
	s.QueryTime = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRules) SetTargetGroupName(v string) *ModifyDBResourceGroupRequestRules {
	s.TargetGroupName = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRules) Validate() error {
	return dara.Validate(s)
}
