// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAtmConfigShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetAtmConfigShrink() *string
	SetAutoStopInterval(v string) *ModifyDBResourceGroupShrinkRequest
	GetAutoStopInterval() *string
	SetClusterMode(v string) *ModifyDBResourceGroupShrinkRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *ModifyDBResourceGroupShrinkRequest
	GetDBClusterId() *string
	SetEnableSpot(v bool) *ModifyDBResourceGroupShrinkRequest
	GetEnableSpot() *bool
	SetEngineParamsShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetEngineParamsShrink() *string
	SetGpuElasticPlanShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetGpuElasticPlanShrink() *string
	SetGroupName(v string) *ModifyDBResourceGroupShrinkRequest
	GetGroupName() *string
	SetGroupType(v string) *ModifyDBResourceGroupShrinkRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetMaxComputeResource() *string
	SetMaxGpuQuantity(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMaxGpuQuantity() *int32
	SetMinClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *ModifyDBResourceGroupShrinkRequest
	GetMinComputeResource() *string
	SetMinGpuQuantity(v int32) *ModifyDBResourceGroupShrinkRequest
	GetMinGpuQuantity() *int32
	SetRayConfigShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetRayConfigShrink() *string
	SetRegionId(v string) *ModifyDBResourceGroupShrinkRequest
	GetRegionId() *string
	SetRulesShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetRulesShrink() *string
	SetSpecName(v string) *ModifyDBResourceGroupShrinkRequest
	GetSpecName() *string
	SetStatus(v string) *ModifyDBResourceGroupShrinkRequest
	GetStatus() *string
	SetTargetResourceGroupName(v string) *ModifyDBResourceGroupShrinkRequest
	GetTargetResourceGroupName() *string
}

type ModifyDBResourceGroupShrinkRequest struct {
	AtmConfigShrink *string `json:"AtmConfig,omitempty" xml:"AtmConfig,omitempty"`
	// The idle duration after which the resource group is automatically stopped.
	//
	// example:
	//
	// 5m
	AutoStopInterval *string `json:"AutoStopInterval,omitempty" xml:"AutoStopInterval,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// 无
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// 无
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// <props="china">The ID of the Data Lakehouse Edition, Enterprise Edition, or Basic Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. This feature provides resources at a lower unit price, but they can be reclaimed at any time. Only `Job` resource groups support this feature. Valid values:
	//
	// - **True**: enables the spot instance feature.
	//
	// - **False**: disables the spot instance feature.
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
	EngineParamsShrink *string `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The time-based scaling plan for GPUs.
	GpuElasticPlanShrink *string `json:"GpuElasticPlan,omitempty" xml:"GpuElasticPlan,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name for a specific cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// - **Interactive**
	//
	// - **Job**
	//
	// > For more information about resource groups in Data Lakehouse Edition clusters, see [Resource groups](https://help.aliyun.com/document_detail/428610.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Interactive
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// 无
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources. The value cannot exceed the unallocated computing resources of the cluster.
	//
	// - If the resource group type is `Interactive`, the value is specified in increments of 16 ACU.
	//
	// - If the resource group type is `Job`, the value is specified in increments of 8 ACU.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	MaxGpuQuantity *int32 `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// 无
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources.
	//
	// - If the resource group type is `Interactive`, the minimum amount of reserved computing resources is 16 ACU.
	//
	// - If the resource group type is `Job`, the minimum amount of reserved computing resources is 0 ACU.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	MinGpuQuantity *int32 `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	// The Ray configuration. This parameter is required if the resource group is an AI group and uses a Ray cluster as its engine.
	RayConfigShrink *string `json:"RayConfig,omitempty" xml:"RayConfig,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job submission rules.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// The desired state of the resource group. Specify **starting*	- to start the resource group or **stopping*	- to stop it.
	//
	// example:
	//
	// starting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// Reserved parameter. Not applicable.
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s ModifyDBResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupShrinkRequest) GetAtmConfigShrink() *string {
	return s.AtmConfigShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetAutoStopInterval() *string {
	return s.AutoStopInterval
}

func (s *ModifyDBResourceGroupShrinkRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *ModifyDBResourceGroupShrinkRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetEnableSpot() *bool {
	return s.EnableSpot
}

func (s *ModifyDBResourceGroupShrinkRequest) GetEngineParamsShrink() *string {
	return s.EngineParamsShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetGpuElasticPlanShrink() *string {
	return s.GpuElasticPlanShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDBResourceGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMaxGpuQuantity() *int32 {
	return s.MaxGpuQuantity
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *ModifyDBResourceGroupShrinkRequest) GetMinGpuQuantity() *int32 {
	return s.MinGpuQuantity
}

func (s *ModifyDBResourceGroupShrinkRequest) GetRayConfigShrink() *string {
	return s.RayConfigShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) GetSpecName() *string {
	return s.SpecName
}

func (s *ModifyDBResourceGroupShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyDBResourceGroupShrinkRequest) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *ModifyDBResourceGroupShrinkRequest) SetAtmConfigShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.AtmConfigShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetAutoStopInterval(v string) *ModifyDBResourceGroupShrinkRequest {
	s.AutoStopInterval = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterMode(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetDBClusterId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetEnableSpot(v bool) *ModifyDBResourceGroupShrinkRequest {
	s.EnableSpot = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetEngineParamsShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.EngineParamsShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGpuElasticPlanShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GpuElasticPlanShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupName(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetGroupType(v string) *ModifyDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMaxGpuQuantity(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MaxGpuQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *ModifyDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetMinGpuQuantity(v int32) *ModifyDBResourceGroupShrinkRequest {
	s.MinGpuQuantity = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetRayConfigShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.RayConfigShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetRegionId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetRulesShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetSpecName(v string) *ModifyDBResourceGroupShrinkRequest {
	s.SpecName = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetStatus(v string) *ModifyDBResourceGroupShrinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetTargetResourceGroupName(v string) *ModifyDBResourceGroupShrinkRequest {
	s.TargetResourceGroupName = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
