// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the spot instance feature for the resource group. After you enable the spot instance feature, you are charged for resources at a lower unit price but the resources are probably released. You can enable the spot instance feature only for job resource groups. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// true
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// example:
	//
	// {\\"spark.adb.version\\":\\"3.5\\"}
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the name of a resource group in a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// > For information about resource groups of Data Lakehouse Edition, see [Resource groups](https://help.aliyun.com/document_detail/428610.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Interactive
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources.
	//
	// 	- If GroupType is set to Interactive, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 16ACU.
	//
	// 	- If GroupType is set to Job, the maximum amount of reserved computing resources refers to the amount of resources that are not allocated in the cluster. Set this parameter to a value in increments of 8ACU.
	//
	// example:
	//
	// 48ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	MaxGpuQuantity     *int32  `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources.
	//
	// 	- If the GroupType parameter is set to Interactive, set the value to 16ACU.
	//
	// 	- If GroupType is set to Job, set the value to 0ACU.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string                                `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	MinGpuQuantity     *int32                                 `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	RayConfig          *ModifyDBResourceGroupRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job resubmission rules.
	Rules    []*ModifyDBResourceGroupRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	SpecName *string                              `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// example:
	//
	// starting
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
}

func (s ModifyDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupRequest) GoString() string {
	return s.String()
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

type ModifyDBResourceGroupRequestRayConfig struct {
	AppConfig        *ModifyDBResourceGroupRequestRayConfigAppConfig       `json:"AppConfig,omitempty" xml:"AppConfig,omitempty" type:"Struct"`
	Category         *string                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	EnableUserEni    *bool                                                 `json:"EnableUserEni,omitempty" xml:"EnableUserEni,omitempty"`
	HeadAllocateUnit *string                                               `json:"HeadAllocateUnit,omitempty" xml:"HeadAllocateUnit,omitempty"`
	HeadDiskCapacity *string                                               `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	HeadSpec         *string                                               `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	HeadSpecType     *string                                               `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	StorageMounts    []*ModifyDBResourceGroupRequestRayConfigStorageMounts `json:"StorageMounts,omitempty" xml:"StorageMounts,omitempty" type:"Repeated"`
	WorkerGroups     []*ModifyDBResourceGroupRequestRayConfigWorkerGroups  `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
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
	// example:
	//
	// app01
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// IsaacLab
	AppType       *string                                                      `json:"AppType,omitempty" xml:"AppType,omitempty"`
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
	// example:
	//
	// lab2.10.0-ray2.43.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// vLLM
	InferenceEngine *string `json:"InferenceEngine,omitempty" xml:"InferenceEngine,omitempty"`
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
	// example:
	//
	// /mnt/data01
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// 1
	StorageId *int64 `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
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

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) SetMountPath(v string) *ModifyDBResourceGroupRequestRayConfigStorageMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) SetStorageId(v int64) *ModifyDBResourceGroupRequestRayConfigStorageMounts {
	s.StorageId = &v
	return s
}

func (s *ModifyDBResourceGroupRequestRayConfigStorageMounts) Validate() error {
	return dara.Validate(s)
}

type ModifyDBResourceGroupRequestRayConfigWorkerGroups struct {
	AllocateUnit       *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	GroupName          *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MaxWorkerQuantity  *int32  `json:"MaxWorkerQuantity,omitempty" xml:"MaxWorkerQuantity,omitempty"`
	MinWorkerQuantity  *int32  `json:"MinWorkerQuantity,omitempty" xml:"MinWorkerQuantity,omitempty"`
	WorkerDiskCapacity *string `json:"WorkerDiskCapacity,omitempty" xml:"WorkerDiskCapacity,omitempty"`
	WorkerSpecName     *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
	WorkerSpecType     *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
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
	// The name of the resource group.
	//
	// example:
	//
	// user_default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The execution duration of the query. Unit: milliseconds.
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
