// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAtmConfig(v *CreateDBResourceGroupRequestAtmConfig) *CreateDBResourceGroupRequest
	GetAtmConfig() *CreateDBResourceGroupRequestAtmConfig
	SetAutoStopInterval(v string) *CreateDBResourceGroupRequest
	GetAutoStopInterval() *string
	SetClassification(v string) *CreateDBResourceGroupRequest
	GetClassification() *string
	SetClusterMode(v string) *CreateDBResourceGroupRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *CreateDBResourceGroupRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *CreateDBResourceGroupRequest
	GetDBClusterId() *string
	SetEnableSpot(v bool) *CreateDBResourceGroupRequest
	GetEnableSpot() *bool
	SetEngine(v string) *CreateDBResourceGroupRequest
	GetEngine() *string
	SetEngineParams(v map[string]interface{}) *CreateDBResourceGroupRequest
	GetEngineParams() map[string]interface{}
	SetGpuElasticPlan(v *CreateDBResourceGroupRequestGpuElasticPlan) *CreateDBResourceGroupRequest
	GetGpuElasticPlan() *CreateDBResourceGroupRequestGpuElasticPlan
	SetGroupName(v string) *CreateDBResourceGroupRequest
	GetGroupName() *string
	SetGroupType(v string) *CreateDBResourceGroupRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *CreateDBResourceGroupRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *CreateDBResourceGroupRequest
	GetMaxComputeResource() *string
	SetMaxGpuQuantity(v int32) *CreateDBResourceGroupRequest
	GetMaxGpuQuantity() *int32
	SetMinClusterCount(v int32) *CreateDBResourceGroupRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *CreateDBResourceGroupRequest
	GetMinComputeResource() *string
	SetMinGpuQuantity(v int32) *CreateDBResourceGroupRequest
	GetMinGpuQuantity() *int32
	SetRayConfig(v *CreateDBResourceGroupRequestRayConfig) *CreateDBResourceGroupRequest
	GetRayConfig() *CreateDBResourceGroupRequestRayConfig
	SetRegionId(v string) *CreateDBResourceGroupRequest
	GetRegionId() *string
	SetRules(v []*CreateDBResourceGroupRequestRules) *CreateDBResourceGroupRequest
	GetRules() []*CreateDBResourceGroupRequestRules
	SetScalePolicy(v string) *CreateDBResourceGroupRequest
	GetScalePolicy() *string
	SetSpecName(v string) *CreateDBResourceGroupRequest
	GetSpecName() *string
	SetTargetResourceGroupName(v string) *CreateDBResourceGroupRequest
	GetTargetResourceGroupName() *string
}

type CreateDBResourceGroupRequest struct {
	AtmConfig *CreateDBResourceGroupRequestAtmConfig `json:"AtmConfig,omitempty" xml:"AtmConfig,omitempty" type:"Struct"`
	// The automatic stop interval. Unit: minutes (m).
	//
	// example:
	//
	// 5m
	AutoStopInterval *string `json:"AutoStopInterval,omitempty" xml:"AutoStopInterval,omitempty"`
	// The classification of the resource group. Valid values:
	//
	// - SQL
	//
	// - SparkSQL
	//
	// - MultiCluster
	//
	// - AI
	//
	// example:
	//
	// SQL
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// -
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// -
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the Dedicated Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. After the spot instance feature is enabled, the unit price of resources is reduced, but the resources may be released. Only Job resource groups support this feature. Valid values:
	//
	// - **True**: enables the spot instance feature.
	//
	// - **False**: disables the spot instance feature.
	//
	// example:
	//
	// True
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The database engine. Valid values:
	//
	// - **AnalyticDB*	- (default): the AnalyticDB for MySQL engine.
	//
	// - **SparkWarehouse**: the SparkWarehouse engine.
	//
	// example:
	//
	// SparkWarehouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine configuration.
	//
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The GPU time-sharing elastic plan.
	GpuElasticPlan *CreateDBResourceGroupRequestGpuElasticPlan `json:"GpuElasticPlan,omitempty" xml:"GpuElasticPlan,omitempty" type:"Struct"`
	// The name of the resource group.
	//
	// - The name can be up to 255 characters in length.
	//
	// - The name must start with a digit, an uppercase letter, or a lowercase letter.
	//
	// - The name can contain digits, uppercase letters, lowercase letters, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// - **Interactive**
	//
	// - **Job**
	//
	// > For more information about Data Lakehouse Edition resource groups, see [Resource group overview (Data Lakehouse Edition)](https://help.aliyun.com/document_detail/428610.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// -
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACUs.
	//
	// - If the resource group type is Interactive, the maximum reserved computing resources is the current unallocated resources of the cluster, in increments of 16 ACUs.
	//
	// - If the resource group type is Job, the maximum reserved computing resources is the current unallocated resources of the cluster, in increments of 8 ACUs.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// The maximum number of GPUs.
	//
	// example:
	//
	// 2
	MaxGpuQuantity *int32 `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// A reserved parameter (not applicable).
	//
	// example:
	//
	// -
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources. Unit: ACUs.
	//
	// - If the resource group type is Interactive, the minimum reserved computing resources is 16 ACUs.
	//
	// - If the resource group type is Job, the minimum reserved computing resources is 0 ACUs.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The minimum number of GPUs.
	//
	// example:
	//
	// 1
	MinGpuQuantity *int32 `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	// The Ray configuration.
	//
	// > This parameter is required when the resource group is an AI resource group and the corresponding engine is RayCluster.
	RayConfig *CreateDBResourceGroupRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the region IDs of AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job routing rules.
	Rules []*CreateDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The scaling policy of the resource group. Valid values:
	//
	// - AutoScaling: enables the AutoScaling automatic scaling policy.
	//
	// - Disable: disables automatic scaling.
	//
	// - MultiCluster: enables the MultiCluster automatic scaling policy.
	//
	// example:
	//
	// AutoScaling
	ScalePolicy *string `json:"ScalePolicy,omitempty" xml:"ScalePolicy,omitempty"`
	// The specification name.
	//
	// example:
	//
	// ADB.MLLarge.2
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// The name of the destination resource group.
	//
	// example:
	//
	// test
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s CreateDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) GetAtmConfig() *CreateDBResourceGroupRequestAtmConfig {
	return s.AtmConfig
}

func (s *CreateDBResourceGroupRequest) GetAutoStopInterval() *string {
	return s.AutoStopInterval
}

func (s *CreateDBResourceGroupRequest) GetClassification() *string {
	return s.Classification
}

func (s *CreateDBResourceGroupRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateDBResourceGroupRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *CreateDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBResourceGroupRequest) GetEnableSpot() *bool {
	return s.EnableSpot
}

func (s *CreateDBResourceGroupRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBResourceGroupRequest) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *CreateDBResourceGroupRequest) GetGpuElasticPlan() *CreateDBResourceGroupRequestGpuElasticPlan {
	return s.GpuElasticPlan
}

func (s *CreateDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDBResourceGroupRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *CreateDBResourceGroupRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *CreateDBResourceGroupRequest) GetMaxGpuQuantity() *int32 {
	return s.MaxGpuQuantity
}

func (s *CreateDBResourceGroupRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *CreateDBResourceGroupRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *CreateDBResourceGroupRequest) GetMinGpuQuantity() *int32 {
	return s.MinGpuQuantity
}

func (s *CreateDBResourceGroupRequest) GetRayConfig() *CreateDBResourceGroupRequestRayConfig {
	return s.RayConfig
}

func (s *CreateDBResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBResourceGroupRequest) GetRules() []*CreateDBResourceGroupRequestRules {
	return s.Rules
}

func (s *CreateDBResourceGroupRequest) GetScalePolicy() *string {
	return s.ScalePolicy
}

func (s *CreateDBResourceGroupRequest) GetSpecName() *string {
	return s.SpecName
}

func (s *CreateDBResourceGroupRequest) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *CreateDBResourceGroupRequest) SetAtmConfig(v *CreateDBResourceGroupRequestAtmConfig) *CreateDBResourceGroupRequest {
	s.AtmConfig = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetAutoStopInterval(v string) *CreateDBResourceGroupRequest {
	s.AutoStopInterval = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClassification(v string) *CreateDBResourceGroupRequest {
	s.Classification = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClusterMode(v string) *CreateDBResourceGroupRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetDBClusterId(v string) *CreateDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEnableSpot(v bool) *CreateDBResourceGroupRequest {
	s.EnableSpot = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEngine(v string) *CreateDBResourceGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetEngineParams(v map[string]interface{}) *CreateDBResourceGroupRequest {
	s.EngineParams = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGpuElasticPlan(v *CreateDBResourceGroupRequestGpuElasticPlan) *CreateDBResourceGroupRequest {
	s.GpuElasticPlan = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupName(v string) *CreateDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetGroupType(v string) *CreateDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMaxGpuQuantity(v int32) *CreateDBResourceGroupRequest {
	s.MaxGpuQuantity = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinComputeResource(v string) *CreateDBResourceGroupRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetMinGpuQuantity(v int32) *CreateDBResourceGroupRequest {
	s.MinGpuQuantity = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetRayConfig(v *CreateDBResourceGroupRequestRayConfig) *CreateDBResourceGroupRequest {
	s.RayConfig = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetRegionId(v string) *CreateDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetRules(v []*CreateDBResourceGroupRequestRules) *CreateDBResourceGroupRequest {
	s.Rules = v
	return s
}

func (s *CreateDBResourceGroupRequest) SetScalePolicy(v string) *CreateDBResourceGroupRequest {
	s.ScalePolicy = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetSpecName(v string) *CreateDBResourceGroupRequest {
	s.SpecName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetTargetResourceGroupName(v string) *CreateDBResourceGroupRequest {
	s.TargetResourceGroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) Validate() error {
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

type CreateDBResourceGroupRequestAtmConfig struct {
	// example:
	//
	// 2
	AuthNodeNum *int32 `json:"AuthNodeNum,omitempty" xml:"AuthNodeNum,omitempty"`
	// example:
	//
	// 8ACU
	AuthNodeSpec *string `json:"AuthNodeSpec,omitempty" xml:"AuthNodeSpec,omitempty"`
	// example:
	//
	// 2
	InsertNodeNum *int32 `json:"InsertNodeNum,omitempty" xml:"InsertNodeNum,omitempty"`
	// example:
	//
	// 8ACU
	InsertNodeSpec *string `json:"InsertNodeSpec,omitempty" xml:"InsertNodeSpec,omitempty"`
	// example:
	//
	// 10
	SelectNodeCacheSize *int32 `json:"SelectNodeCacheSize,omitempty" xml:"SelectNodeCacheSize,omitempty"`
	// example:
	//
	// 1
	SelectNodeNum *int32 `json:"SelectNodeNum,omitempty" xml:"SelectNodeNum,omitempty"`
	// example:
	//
	// 8ACU
	SelectNodeSpec *string `json:"SelectNodeSpec,omitempty" xml:"SelectNodeSpec,omitempty"`
	// example:
	//
	// 1
	StorageNodeDiskSize *int32 `json:"StorageNodeDiskSize,omitempty" xml:"StorageNodeDiskSize,omitempty"`
	// example:
	//
	// essd_pl1
	StorageNodeDiskType *string `json:"StorageNodeDiskType,omitempty" xml:"StorageNodeDiskType,omitempty"`
	// example:
	//
	// 2
	StorageNodeNum *int32 `json:"StorageNodeNum,omitempty" xml:"StorageNodeNum,omitempty"`
	// example:
	//
	// 8ACU
	StorageNodeSpec *string `json:"StorageNodeSpec,omitempty" xml:"StorageNodeSpec,omitempty"`
}

func (s CreateDBResourceGroupRequestAtmConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestAtmConfig) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetAuthNodeNum() *int32 {
	return s.AuthNodeNum
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetAuthNodeSpec() *string {
	return s.AuthNodeSpec
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetInsertNodeNum() *int32 {
	return s.InsertNodeNum
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetInsertNodeSpec() *string {
	return s.InsertNodeSpec
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetSelectNodeCacheSize() *int32 {
	return s.SelectNodeCacheSize
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetSelectNodeNum() *int32 {
	return s.SelectNodeNum
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetSelectNodeSpec() *string {
	return s.SelectNodeSpec
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetStorageNodeDiskSize() *int32 {
	return s.StorageNodeDiskSize
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetStorageNodeDiskType() *string {
	return s.StorageNodeDiskType
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetStorageNodeNum() *int32 {
	return s.StorageNodeNum
}

func (s *CreateDBResourceGroupRequestAtmConfig) GetStorageNodeSpec() *string {
	return s.StorageNodeSpec
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetAuthNodeNum(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.AuthNodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetAuthNodeSpec(v string) *CreateDBResourceGroupRequestAtmConfig {
	s.AuthNodeSpec = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetInsertNodeNum(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.InsertNodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetInsertNodeSpec(v string) *CreateDBResourceGroupRequestAtmConfig {
	s.InsertNodeSpec = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetSelectNodeCacheSize(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.SelectNodeCacheSize = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetSelectNodeNum(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.SelectNodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetSelectNodeSpec(v string) *CreateDBResourceGroupRequestAtmConfig {
	s.SelectNodeSpec = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetStorageNodeDiskSize(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.StorageNodeDiskSize = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetStorageNodeDiskType(v string) *CreateDBResourceGroupRequestAtmConfig {
	s.StorageNodeDiskType = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetStorageNodeNum(v int32) *CreateDBResourceGroupRequestAtmConfig {
	s.StorageNodeNum = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) SetStorageNodeSpec(v string) *CreateDBResourceGroupRequestAtmConfig {
	s.StorageNodeSpec = &v
	return s
}

func (s *CreateDBResourceGroupRequestAtmConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDBResourceGroupRequestGpuElasticPlan struct {
	// Specifies whether to enable the elastic plan immediately after creation. Valid values:
	//
	// - true: enables the elastic plan immediately.
	//
	// - false: does not enable the elastic plan.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of rules.
	Rules []*CreateDBResourceGroupRequestGpuElasticPlanRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateDBResourceGroupRequestGpuElasticPlan) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestGpuElasticPlan) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestGpuElasticPlan) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDBResourceGroupRequestGpuElasticPlan) GetRules() []*CreateDBResourceGroupRequestGpuElasticPlanRules {
	return s.Rules
}

func (s *CreateDBResourceGroupRequestGpuElasticPlan) SetEnabled(v bool) *CreateDBResourceGroupRequestGpuElasticPlan {
	s.Enabled = &v
	return s
}

func (s *CreateDBResourceGroupRequestGpuElasticPlan) SetRules(v []*CreateDBResourceGroupRequestGpuElasticPlanRules) *CreateDBResourceGroupRequestGpuElasticPlan {
	s.Rules = v
	return s
}

func (s *CreateDBResourceGroupRequestGpuElasticPlan) Validate() error {
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

type CreateDBResourceGroupRequestGpuElasticPlanRules struct {
	// The end time as a cron expression. The interval must be at least 1 hour.
	//
	// example:
	//
	// 0 0 3 	- 	- ?
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	// The start time as a cron expression. The interval must be at least 1 hour.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	StartCronExpression *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
}

func (s CreateDBResourceGroupRequestGpuElasticPlanRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestGpuElasticPlanRules) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestGpuElasticPlanRules) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *CreateDBResourceGroupRequestGpuElasticPlanRules) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *CreateDBResourceGroupRequestGpuElasticPlanRules) SetEndCronExpression(v string) *CreateDBResourceGroupRequestGpuElasticPlanRules {
	s.EndCronExpression = &v
	return s
}

func (s *CreateDBResourceGroupRequestGpuElasticPlanRules) SetStartCronExpression(v string) *CreateDBResourceGroupRequestGpuElasticPlanRules {
	s.StartCronExpression = &v
	return s
}

func (s *CreateDBResourceGroupRequestGpuElasticPlanRules) Validate() error {
	return dara.Validate(s)
}

type CreateDBResourceGroupRequestRayConfig struct {
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
	// Specifies whether to enable user ENI connectivity.
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
	HeadSpecType            *string `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	UserDefinedRequirements *string `json:"UserDefinedRequirements,omitempty" xml:"UserDefinedRequirements,omitempty"`
	// The list of Ray worker group configurations.
	WorkerGroups []*CreateDBResourceGroupRequestRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s CreateDBResourceGroupRequestRayConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestRayConfig) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestRayConfig) GetCategory() *string {
	return s.Category
}

func (s *CreateDBResourceGroupRequestRayConfig) GetEnableUserEni() *bool {
	return s.EnableUserEni
}

func (s *CreateDBResourceGroupRequestRayConfig) GetHeadAllocateUnit() *string {
	return s.HeadAllocateUnit
}

func (s *CreateDBResourceGroupRequestRayConfig) GetHeadDiskCapacity() *string {
	return s.HeadDiskCapacity
}

func (s *CreateDBResourceGroupRequestRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *CreateDBResourceGroupRequestRayConfig) GetHeadSpecType() *string {
	return s.HeadSpecType
}

func (s *CreateDBResourceGroupRequestRayConfig) GetUserDefinedRequirements() *string {
	return s.UserDefinedRequirements
}

func (s *CreateDBResourceGroupRequestRayConfig) GetWorkerGroups() []*CreateDBResourceGroupRequestRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *CreateDBResourceGroupRequestRayConfig) SetCategory(v string) *CreateDBResourceGroupRequestRayConfig {
	s.Category = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetEnableUserEni(v bool) *CreateDBResourceGroupRequestRayConfig {
	s.EnableUserEni = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetHeadAllocateUnit(v string) *CreateDBResourceGroupRequestRayConfig {
	s.HeadAllocateUnit = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetHeadDiskCapacity(v string) *CreateDBResourceGroupRequestRayConfig {
	s.HeadDiskCapacity = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetHeadSpec(v string) *CreateDBResourceGroupRequestRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetHeadSpecType(v string) *CreateDBResourceGroupRequestRayConfig {
	s.HeadSpecType = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetUserDefinedRequirements(v string) *CreateDBResourceGroupRequestRayConfig {
	s.UserDefinedRequirements = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) SetWorkerGroups(v []*CreateDBResourceGroupRequestRayConfigWorkerGroups) *CreateDBResourceGroupRequestRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfig) Validate() error {
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

type CreateDBResourceGroupRequestRayConfigWorkerGroups struct {
	// The allocation unit.
	//
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// The name of the worker group.
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

func (s CreateDBResourceGroupRequestRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetAllocateUnit(v string) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetGroupName(v string) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerSpecName(v string) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) SetWorkerSpecType(v string) *CreateDBResourceGroupRequestRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *CreateDBResourceGroupRequestRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type CreateDBResourceGroupRequestRules struct {
	// The name of the resource group.
	//
	// - The name can be up to 255 characters in length.
	//
	// - The name must start with a digit, an uppercase letter, or a lowercase letter.
	//
	// - The name can contain digits, uppercase letters, lowercase letters, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The query execution time threshold. Unit: milliseconds (ms).
	//
	// example:
	//
	// 180000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the destination resource group.
	//
	// example:
	//
	// job
	TargetGroupName *string `json:"TargetGroupName,omitempty" xml:"TargetGroupName,omitempty"`
}

func (s CreateDBResourceGroupRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequestRules) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequestRules) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupRequestRules) GetQueryTime() *string {
	return s.QueryTime
}

func (s *CreateDBResourceGroupRequestRules) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *CreateDBResourceGroupRequestRules) SetGroupName(v string) *CreateDBResourceGroupRequestRules {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequestRules) SetQueryTime(v string) *CreateDBResourceGroupRequestRules {
	s.QueryTime = &v
	return s
}

func (s *CreateDBResourceGroupRequestRules) SetTargetGroupName(v string) *CreateDBResourceGroupRequestRules {
	s.TargetGroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequestRules) Validate() error {
	return dara.Validate(s)
}
