// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody
	GetGroupsInfo() []*DescribeDBResourceGroupResponseBodyGroupsInfo
	SetRequestId(v string) *DescribeDBResourceGroupResponseBody
	GetRequestId() *string
}

type DescribeDBResourceGroupResponseBody struct {
	// The list of resource group information.
	GroupsInfo []*DescribeDBResourceGroupResponseBodyGroupsInfo `json:"GroupsInfo,omitempty" xml:"GroupsInfo,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A94B6C02-7BD4-5D67-9776-3AC8317E8DD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBody) GetGroupsInfo() []*DescribeDBResourceGroupResponseBodyGroupsInfo {
	return s.GroupsInfo
}

func (s *DescribeDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBResourceGroupResponseBody) SetGroupsInfo(v []*DescribeDBResourceGroupResponseBodyGroupsInfo) *DescribeDBResourceGroupResponseBody {
	s.GroupsInfo = v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetRequestId(v string) *DescribeDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) Validate() error {
	if s.GroupsInfo != nil {
		for _, item := range s.GroupsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyGroupsInfo struct {
	// The PromQL resource group configuration.
	AtmConfig *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig `json:"AtmConfig,omitempty" xml:"AtmConfig,omitempty" type:"Struct"`
	// The auto-stop interval. The value is in the format of a number followed by "m", which indicates minutes. Valid values range from 0m or 5m to 10080m. A value of 0m indicates that the resource group does not automatically stop.
	//
	// example:
	//
	// 0m
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
	// The reserved parameter. This parameter is not applicable.
	//
	// example:
	//
	// Reserved parameter
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The reserved parameter. This parameter is not applicable.
	//
	// example:
	//
	// Reserved parameter
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC in the <i>yyyy-MM-ddTHH:mm:ssZ</i> format.
	//
	// example:
	//
	// 2022-08-29T03:34:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The minimum elastic compute resources. Unit: ACU.
	//
	// example:
	//
	// 16ACU
	ElasticMinComputeResource *string `json:"ElasticMinComputeResource,omitempty" xml:"ElasticMinComputeResource,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. After you enable this feature, the unit price of resources decreases, but the instances may be released. Valid values:
	//
	// - **True**: enables the spot instance feature.
	//
	// - **False**: disables the spot instance feature.
	//
	// This parameter can be set to True only for Job resource groups.
	//
	// example:
	//
	// True
	EnableSpot *string `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// The engine type.
	//
	// example:
	//
	// AnalyticDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine parameters.
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The GPU time-based elastic plan.
	GpuElasticPlan *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan `json:"GpuElasticPlan,omitempty" xml:"GpuElasticPlan,omitempty" type:"Struct"`
	// The name of the resource group.
	//
	// example:
	//
	// test1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The resource group type. Valid values:
	//
	// - **Interactive**
	//
	// - **Job**
	//
	// > For more information about Data Lakehouse Edition resource groups, see [Resource group overview (Data Lakehouse Edition)](https://help.aliyun.com/document_detail/428610.html).
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The Resource Access Management (RAM) users attached to the resource group.
	//
	// example:
	//
	// testb,testc
	GroupUsers *string `json:"GroupUsers,omitempty" xml:"GroupUsers,omitempty"`
	// The reserved parameter. This parameter is not applicable.
	//
	// example:
	//
	// Reserved parameter
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum reserved compute resources. Unit: ACU.
	//
	// example:
	//
	// 512ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// The maximum number of GPUs.
	//
	// example:
	//
	// 2
	MaxGpuQuantity *int32 `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// The message of the job delivery rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reserved parameter. This parameter is not applicable.
	//
	// example:
	//
	// Reserved parameter
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum reserved compute resources. Unit: ACU.
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
	// The Ray configuration information.
	RayConfig *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The job delivery rules.
	Rules []*DescribeDBResourceGroupResponseBodyGroupsInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The reserved parameter. This parameter is not applicable.
	//
	// example:
	//
	// Reserved parameter
	RunningClusterCount *int32 `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	// The scale-out policy of the resource group. Valid values:
	//
	// - AutoScaling: enables the AutoScaling policy.
	//
	// - Disable: disables auto scaling.
	//
	// - MultiCluster: enables the MultiCluster policy.
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
	// The status of the resource group. Valid values:
	//
	// - **Pending**: The resource group is being created.
	//
	// - **Running**: The resource group is running.
	//
	// - **Scaling**: The resource group is being scaled.
	//
	// - **Deleting**: The resource group is being deleted.
	//
	// - **Deleted**: The resource group is deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the destination resource group.
	//
	// example:
	//
	// test
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
	// The time when the resource group was updated. The time is displayed in UTC in the <i>yyyy-MM-ddTHH:mm:ssZ</i> format.
	//
	// example:
	//
	// 2022-08-31T03:34:30Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetAtmConfig() *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	return s.AtmConfig
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetAutoStopInterval() *string {
	return s.AutoStopInterval
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetClassification() *string {
	return s.Classification
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetElasticMinComputeResource() *string {
	return s.ElasticMinComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetEnableSpot() *string {
	return s.EnableSpot
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetEngineParams() map[string]interface{} {
	return s.EngineParams
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGpuElasticPlan() *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan {
	return s.GpuElasticPlan
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetGroupUsers() *string {
	return s.GroupUsers
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMaxGpuQuantity() *int32 {
	return s.MaxGpuQuantity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetMinGpuQuantity() *int32 {
	return s.MinGpuQuantity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetRayConfig() *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	return s.RayConfig
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetRules() []*DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	return s.Rules
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetRunningClusterCount() *int32 {
	return s.RunningClusterCount
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetScalePolicy() *string {
	return s.ScalePolicy
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetSpecName() *string {
	return s.SpecName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetAtmConfig(v *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.AtmConfig = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetAutoStopInterval(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.AutoStopInterval = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClassification(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Classification = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterMode(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterMode = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetClusterSizeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ClusterSizeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetCreateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetElasticMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ElasticMinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEnableSpot(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.EnableSpot = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEngine(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Engine = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetEngineParams(v map[string]interface{}) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.EngineParams = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGpuElasticPlan(v *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GpuElasticPlan = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetGroupUsers(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.GroupUsers = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMaxGpuQuantity(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MaxGpuQuantity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMessage(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Message = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinComputeResource(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetMinGpuQuantity(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.MinGpuQuantity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRayConfig(v *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.RayConfig = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRules(v []*DescribeDBResourceGroupResponseBodyGroupsInfoRules) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Rules = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetRunningClusterCount(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.RunningClusterCount = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetScalePolicy(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.ScalePolicy = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetSpecName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.SpecName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetStatus(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.Status = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetTargetResourceGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.TargetResourceGroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetUpdateTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) Validate() error {
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

type DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig struct {
	// The number of authentication nodes.
	//
	// example:
	//
	// 2
	AuthNodeNum *string `json:"AuthNodeNum,omitempty" xml:"AuthNodeNum,omitempty"`
	// The authentication node specifications.
	//
	// example:
	//
	// 8ACU
	AuthNodeSpec *string `json:"AuthNodeSpec,omitempty" xml:"AuthNodeSpec,omitempty"`
	// The number of insert nodes.
	//
	// example:
	//
	// 1
	InsertNodeNum *string `json:"InsertNodeNum,omitempty" xml:"InsertNodeNum,omitempty"`
	// The insert node specifications.
	//
	// example:
	//
	// 8ACU
	InsertNodeSpec *string `json:"InsertNodeSpec,omitempty" xml:"InsertNodeSpec,omitempty"`
	// The cache size of the query node.
	//
	// example:
	//
	// 10
	SelectNodeCacheSize *string `json:"SelectNodeCacheSize,omitempty" xml:"SelectNodeCacheSize,omitempty"`
	// The number of query nodes.
	//
	// example:
	//
	// 1
	SelectNodeNum *string `json:"SelectNodeNum,omitempty" xml:"SelectNodeNum,omitempty"`
	// The query node specifications.
	//
	// example:
	//
	// 8ACU
	SelectNodeSpec *string `json:"SelectNodeSpec,omitempty" xml:"SelectNodeSpec,omitempty"`
	// The disk size of the storage node.
	//
	// example:
	//
	// 100
	StorageNodeDiskSize *string `json:"StorageNodeDiskSize,omitempty" xml:"StorageNodeDiskSize,omitempty"`
	// The disk type of the storage node.
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
	StorageNodeNum *string `json:"StorageNodeNum,omitempty" xml:"StorageNodeNum,omitempty"`
	// The storage node specifications.
	//
	// example:
	//
	// 8ACU
	StorageNodeSpec *string `json:"StorageNodeSpec,omitempty" xml:"StorageNodeSpec,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetAuthNodeNum() *string {
	return s.AuthNodeNum
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetAuthNodeSpec() *string {
	return s.AuthNodeSpec
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetInsertNodeNum() *string {
	return s.InsertNodeNum
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetInsertNodeSpec() *string {
	return s.InsertNodeSpec
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetSelectNodeCacheSize() *string {
	return s.SelectNodeCacheSize
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetSelectNodeNum() *string {
	return s.SelectNodeNum
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetSelectNodeSpec() *string {
	return s.SelectNodeSpec
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetStorageNodeDiskSize() *string {
	return s.StorageNodeDiskSize
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetStorageNodeDiskType() *string {
	return s.StorageNodeDiskType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetStorageNodeNum() *string {
	return s.StorageNodeNum
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) GetStorageNodeSpec() *string {
	return s.StorageNodeSpec
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetAuthNodeNum(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.AuthNodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetAuthNodeSpec(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.AuthNodeSpec = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetInsertNodeNum(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.InsertNodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetInsertNodeSpec(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.InsertNodeSpec = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetSelectNodeCacheSize(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.SelectNodeCacheSize = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetSelectNodeNum(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.SelectNodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetSelectNodeSpec(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.SelectNodeSpec = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetStorageNodeDiskSize(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.StorageNodeDiskSize = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetStorageNodeDiskType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.StorageNodeDiskType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetStorageNodeNum(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.StorageNodeNum = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) SetStorageNodeSpec(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig {
	s.StorageNodeSpec = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoAtmConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan struct {
	// Indicates whether the GPU time-based elastic plan is enabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of rules.
	Rules []*DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) GetRules() []*DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules {
	return s.Rules
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) SetEnabled(v bool) *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan {
	s.Enabled = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) SetRules(v []*DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan {
	s.Rules = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlan) Validate() error {
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

type DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules struct {
	// The end time in cron expression format. The interval must be no less than 1 hour.
	//
	// example:
	//
	// 0 0 3 	- 	- ?
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	// The start time in cron expression format. The interval must be no less than 1 hour.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	StartCronExpression *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) SetEndCronExpression(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules {
	s.EndCronExpression = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) SetStartCronExpression(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules {
	s.StartCronExpression = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoGpuElasticPlanRules) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig struct {
	// The Ray application configuration.
	//
	// if can be null:
	// true
	AppConfig *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig `json:"AppConfig,omitempty" xml:"AppConfig,omitempty" type:"Struct"`
	// The Ray cluster type. Valid values:
	//
	// - **BASIC**: the basic type, which is not highly available.
	//
	// - **HIGH_AVAILABILITY**: the high-availability type.
	//
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to enable the elastic network interface (ENI).
	//
	// if can be null:
	// false
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
	// 100Gi
	HeadDiskCapacity *string `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	// The node specifications of the head node.
	//
	// example:
	//
	// ADB.MLLarge.2
	HeadSpec *string `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	// The resource type of the head node.
	//
	// example:
	//
	// cpu
	HeadSpecType *string `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	// The address of the Ray cluster.
	//
	// example:
	//
	// http://ray-cluster-address.example.com
	RayClusterAddress *string `json:"RayClusterAddress,omitempty" xml:"RayClusterAddress,omitempty"`
	// The address of the Ray dashboard.
	//
	// example:
	//
	// http://ray-dashboard-address.example.com
	RayDashboardAddress *string `json:"RayDashboardAddress,omitempty" xml:"RayDashboardAddress,omitempty"`
	// The address of Ray Grafana.
	//
	// example:
	//
	// http://ray-grafana-address.example.com
	RayGrafanaAddress *string `json:"RayGrafanaAddress,omitempty" xml:"RayGrafanaAddress,omitempty"`
	// The public endpoint of Ray Serve.
	//
	// example:
	//
	// 1.2.3.4:8100
	RayServePublicAddress *string `json:"RayServePublicAddress,omitempty" xml:"RayServePublicAddress,omitempty"`
	// The list of mounted storage.
	StorageMounts []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts `json:"StorageMounts,omitempty" xml:"StorageMounts,omitempty" type:"Repeated"`
	// The custom Python dependencies for RayCluster. Specify the content in the requirements.txt format, with one pip dependency per line. A maximum of 200 lines are allowed, and each line cannot exceed 512 characters. Version constraints, extras, PEP 440 direct URLs, and # comments are supported, but shell metacharacters are not supported. This parameter applies only to AI resource groups of the RayCluster type.
	//
	// example:
	//
	// numpy==1.24.0
	UserDefinedRequirements *string `json:"UserDefinedRequirements,omitempty" xml:"UserDefinedRequirements,omitempty"`
	// The list of Ray worker groups.
	WorkerGroups []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetAppConfig() *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig {
	return s.AppConfig
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetEnableUserEni() *bool {
	return s.EnableUserEni
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetHeadAllocateUnit() *string {
	return s.HeadAllocateUnit
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetHeadDiskCapacity() *string {
	return s.HeadDiskCapacity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetHeadSpecType() *string {
	return s.HeadSpecType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetRayClusterAddress() *string {
	return s.RayClusterAddress
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetRayDashboardAddress() *string {
	return s.RayDashboardAddress
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetRayGrafanaAddress() *string {
	return s.RayGrafanaAddress
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetRayServePublicAddress() *string {
	return s.RayServePublicAddress
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetStorageMounts() []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	return s.StorageMounts
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetUserDefinedRequirements() *string {
	return s.UserDefinedRequirements
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetWorkerGroups() []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetAppConfig(v *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.AppConfig = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetCategory(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.Category = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetEnableUserEni(v bool) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.EnableUserEni = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetHeadAllocateUnit(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.HeadAllocateUnit = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetHeadDiskCapacity(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.HeadDiskCapacity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetHeadSpec(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetHeadSpecType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.HeadSpecType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetRayClusterAddress(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.RayClusterAddress = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetRayDashboardAddress(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.RayDashboardAddress = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetRayGrafanaAddress(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.RayGrafanaAddress = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetRayServePublicAddress(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.RayServePublicAddress = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetStorageMounts(v []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.StorageMounts = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetUserDefinedRequirements(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.UserDefinedRequirements = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetWorkerGroups(v []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) Validate() error {
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

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig struct {
	// The name of the Ray application.
	//
	// example:
	//
	// app01
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the Ray application.
	//
	// example:
	//
	// IsaacLab
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The image configuration.
	//
	// if can be null:
	// true
	ImageSelector *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector `json:"ImageSelector,omitempty" xml:"ImageSelector,omitempty" type:"Struct"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) GetAppType() *string {
	return s.AppType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) GetImageSelector() *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector {
	return s.ImageSelector
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) SetAppName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig {
	s.AppName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) SetAppType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig {
	s.AppType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) SetImageSelector(v *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig {
	s.ImageSelector = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig) Validate() error {
	if s.ImageSelector != nil {
		if err := s.ImageSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector struct {
	// The image.
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
	// The large language model (LLM).
	//
	// example:
	//
	// Deepseek-R1
	LlmModel *string `json:"LlmModel,omitempty" xml:"LlmModel,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) GetImage() *string {
	return s.Image
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) GetInferenceEngine() *string {
	return s.InferenceEngine
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) GetLlmModel() *string {
	return s.LlmModel
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) SetImage(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector {
	s.Image = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) SetInferenceEngine(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector {
	s.InferenceEngine = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) SetLlmModel(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector {
	s.LlmModel = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfigImageSelector) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts struct {
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
	StorageId *int64 `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
	// The name of the mounted storage.
	//
	// example:
	//
	// training-data
	StorageName *string `json:"StorageName,omitempty" xml:"StorageName,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) GetStorageId() *int64 {
	return s.StorageId
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) GetStorageName() *string {
	return s.StorageName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) SetMountPath(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) SetStorageId(v int64) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	s.StorageId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) SetStorageName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	s.StorageName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups struct {
	// The allocation unit.
	//
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// The name of the Ray worker group.
	//
	// example:
	//
	// g01
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The maximum number of workers.
	//
	// example:
	//
	// 1
	MaxWorkerQuantity *int32 `json:"MaxWorkerQuantity,omitempty" xml:"MaxWorkerQuantity,omitempty"`
	// The minimum number of workers.
	//
	// example:
	//
	// 1
	MinWorkerQuantity *int32 `json:"MinWorkerQuantity,omitempty" xml:"MinWorkerQuantity,omitempty"`
	// The disk size of each worker.
	//
	// example:
	//
	// 100G
	WorkerDiskCapacity *string `json:"WorkerDiskCapacity,omitempty" xml:"WorkerDiskCapacity,omitempty"`
	// The specification of the worker.
	//
	// example:
	//
	// large
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
	// The resource type of the Ray worker.
	//
	// example:
	//
	// CPU
	WorkerSpecType *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetAllocateUnit(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetWorkerSpecName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) SetWorkerSpecType(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRules struct {
	// The name of the resource group.
	//
	// example:
	//
	// user_default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The threshold for query execution time. Unit: milliseconds.
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

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyGroupsInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) GetQueryTime() *string {
	return s.QueryTime
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) GetTargetGroupName() *string {
	return s.TargetGroupName
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetQueryTime(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.QueryTime = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) SetTargetGroupName(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRules {
	s.TargetGroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRules) Validate() error {
	return dara.Validate(s)
}
