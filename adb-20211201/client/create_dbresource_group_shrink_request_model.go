// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStopInterval(v string) *CreateDBResourceGroupShrinkRequest
	GetAutoStopInterval() *string
	SetClusterMode(v string) *CreateDBResourceGroupShrinkRequest
	GetClusterMode() *string
	SetClusterSizeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetClusterSizeResource() *string
	SetDBClusterId(v string) *CreateDBResourceGroupShrinkRequest
	GetDBClusterId() *string
	SetEnableSpot(v bool) *CreateDBResourceGroupShrinkRequest
	GetEnableSpot() *bool
	SetEngine(v string) *CreateDBResourceGroupShrinkRequest
	GetEngine() *string
	SetEngineParamsShrink(v string) *CreateDBResourceGroupShrinkRequest
	GetEngineParamsShrink() *string
	SetGroupName(v string) *CreateDBResourceGroupShrinkRequest
	GetGroupName() *string
	SetGroupType(v string) *CreateDBResourceGroupShrinkRequest
	GetGroupType() *string
	SetMaxClusterCount(v int32) *CreateDBResourceGroupShrinkRequest
	GetMaxClusterCount() *int32
	SetMaxComputeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetMaxComputeResource() *string
	SetMaxGpuQuantity(v int32) *CreateDBResourceGroupShrinkRequest
	GetMaxGpuQuantity() *int32
	SetMinClusterCount(v int32) *CreateDBResourceGroupShrinkRequest
	GetMinClusterCount() *int32
	SetMinComputeResource(v string) *CreateDBResourceGroupShrinkRequest
	GetMinComputeResource() *string
	SetMinGpuQuantity(v int32) *CreateDBResourceGroupShrinkRequest
	GetMinGpuQuantity() *int32
	SetRayConfigShrink(v string) *CreateDBResourceGroupShrinkRequest
	GetRayConfigShrink() *string
	SetRegionId(v string) *CreateDBResourceGroupShrinkRequest
	GetRegionId() *string
	SetRulesShrink(v string) *CreateDBResourceGroupShrinkRequest
	GetRulesShrink() *string
	SetSpecName(v string) *CreateDBResourceGroupShrinkRequest
	GetSpecName() *string
	SetTargetResourceGroupName(v string) *CreateDBResourceGroupShrinkRequest
	GetTargetResourceGroupName() *string
}

type CreateDBResourceGroupShrinkRequest struct {
	// example:
	//
	// 5m
	AutoStopInterval *string `json:"AutoStopInterval,omitempty" xml:"AutoStopInterval,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. After you enable the spot instance feature, you are charged for resources at a lower unit price but the resources are probably released. You can enable the spot instance feature only for job resource groups. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// True
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// example:
	//
	// SparkWarehouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParamsShrink *string `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The name of the resource group.
	//
	// 	- The name can be up to 255 characters in length.
	//
	// 	- The name must start with a letter or a digit.
	//
	// 	- The name can contain letters, digits, hyphens (_), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// >  For more information about resource groups, see [Resource group overview](https://help.aliyun.com/document_detail/428610.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum reserved computing resources.
	//
	// 	- If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16ACU.
	//
	// 	- If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8ACU.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	MaxGpuQuantity *int32 `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum reserved computing resources.
	//
	// 	- When GroupType is set to Interactive, set this parameter to 16ACU.
	//
	// 	- When GroupType is set to Job, set this parameter to 0ACU.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// A reserved parameter.
	MinGpuQuantity  *int32  `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	RayConfigShrink *string `json:"RayConfig,omitempty" xml:"RayConfig,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job resubmission rules.
	RulesShrink *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	// A reserved parameter.
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// A reserved parameter.
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s CreateDBResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupShrinkRequest) GetAutoStopInterval() *string {
	return s.AutoStopInterval
}

func (s *CreateDBResourceGroupShrinkRequest) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *CreateDBResourceGroupShrinkRequest) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBResourceGroupShrinkRequest) GetEnableSpot() *bool {
	return s.EnableSpot
}

func (s *CreateDBResourceGroupShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateDBResourceGroupShrinkRequest) GetEngineParamsShrink() *string {
	return s.EngineParamsShrink
}

func (s *CreateDBResourceGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDBResourceGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDBResourceGroupShrinkRequest) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *CreateDBResourceGroupShrinkRequest) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetMaxGpuQuantity() *int32 {
	return s.MaxGpuQuantity
}

func (s *CreateDBResourceGroupShrinkRequest) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *CreateDBResourceGroupShrinkRequest) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *CreateDBResourceGroupShrinkRequest) GetMinGpuQuantity() *int32 {
	return s.MinGpuQuantity
}

func (s *CreateDBResourceGroupShrinkRequest) GetRayConfigShrink() *string {
	return s.RayConfigShrink
}

func (s *CreateDBResourceGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBResourceGroupShrinkRequest) GetRulesShrink() *string {
	return s.RulesShrink
}

func (s *CreateDBResourceGroupShrinkRequest) GetSpecName() *string {
	return s.SpecName
}

func (s *CreateDBResourceGroupShrinkRequest) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *CreateDBResourceGroupShrinkRequest) SetAutoStopInterval(v string) *CreateDBResourceGroupShrinkRequest {
	s.AutoStopInterval = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterMode(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetClusterSizeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.ClusterSizeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetDBClusterId(v string) *CreateDBResourceGroupShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEnableSpot(v bool) *CreateDBResourceGroupShrinkRequest {
	s.EnableSpot = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEngine(v string) *CreateDBResourceGroupShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetEngineParamsShrink(v string) *CreateDBResourceGroupShrinkRequest {
	s.EngineParamsShrink = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupName(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetGroupType(v string) *CreateDBResourceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MaxClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MaxComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMaxGpuQuantity(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MaxGpuQuantity = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinClusterCount(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MinClusterCount = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinComputeResource(v string) *CreateDBResourceGroupShrinkRequest {
	s.MinComputeResource = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetMinGpuQuantity(v int32) *CreateDBResourceGroupShrinkRequest {
	s.MinGpuQuantity = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetRayConfigShrink(v string) *CreateDBResourceGroupShrinkRequest {
	s.RayConfigShrink = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetRegionId(v string) *CreateDBResourceGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetRulesShrink(v string) *CreateDBResourceGroupShrinkRequest {
	s.RulesShrink = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetSpecName(v string) *CreateDBResourceGroupShrinkRequest {
	s.SpecName = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) SetTargetResourceGroupName(v string) *CreateDBResourceGroupShrinkRequest {
	s.TargetResourceGroupName = &v
	return s
}

func (s *CreateDBResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
