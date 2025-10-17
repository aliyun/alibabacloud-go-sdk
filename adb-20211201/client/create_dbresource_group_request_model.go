// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStopInterval(v string) *CreateDBResourceGroupRequest
	GetAutoStopInterval() *string
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
	SetSpecName(v string) *CreateDBResourceGroupRequest
	GetSpecName() *string
	SetTargetResourceGroupName(v string) *CreateDBResourceGroupRequest
	GetTargetResourceGroupName() *string
}

type CreateDBResourceGroupRequest struct {
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
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
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
	MinGpuQuantity *int32                                 `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	RayConfig      *CreateDBResourceGroupRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job resubmission rules.
	Rules []*CreateDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// A reserved parameter.
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// A reserved parameter.
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s CreateDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupRequest) GetAutoStopInterval() *string {
	return s.AutoStopInterval
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

func (s *CreateDBResourceGroupRequest) GetSpecName() *string {
	return s.SpecName
}

func (s *CreateDBResourceGroupRequest) GetTargetResourceGroupName() *string {
	return s.TargetResourceGroupName
}

func (s *CreateDBResourceGroupRequest) SetAutoStopInterval(v string) *CreateDBResourceGroupRequest {
	s.AutoStopInterval = &v
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

func (s *CreateDBResourceGroupRequest) SetSpecName(v string) *CreateDBResourceGroupRequest {
	s.SpecName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) SetTargetResourceGroupName(v string) *CreateDBResourceGroupRequest {
	s.TargetResourceGroupName = &v
	return s
}

func (s *CreateDBResourceGroupRequest) Validate() error {
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

type CreateDBResourceGroupRequestRayConfig struct {
	// example:
	//
	// BASIC
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	EnableUserEni    *bool   `json:"EnableUserEni,omitempty" xml:"EnableUserEni,omitempty"`
	HeadAllocateUnit *string `json:"HeadAllocateUnit,omitempty" xml:"HeadAllocateUnit,omitempty"`
	HeadDiskCapacity *string `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	// example:
	//
	// xlarge
	HeadSpec     *string                                              `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	HeadSpecType *string                                              `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
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
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 2
	MaxWorkerQuantity *int32 `json:"MaxWorkerQuantity,omitempty" xml:"MaxWorkerQuantity,omitempty"`
	// example:
	//
	// 1
	MinWorkerQuantity *int32 `json:"MinWorkerQuantity,omitempty" xml:"MinWorkerQuantity,omitempty"`
	// example:
	//
	// 100G
	WorkerDiskCapacity *string `json:"WorkerDiskCapacity,omitempty" xml:"WorkerDiskCapacity,omitempty"`
	// example:
	//
	// xlarge
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
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
	// 	- The name can be up to 255 characters in length.
	//
	// 	- The name must start with a letter or digit.
	//
	// 	- The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 180000
	QueryTime *string `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
	// The name of the resource group to which you want to resubmit the query job.
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
