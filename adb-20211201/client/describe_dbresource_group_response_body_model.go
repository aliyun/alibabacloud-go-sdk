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
	// The queried resource groups.
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
	// The time when the resource group was created. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-29T03:34:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The minimum amount of elastic computing resources.
	//
	// example:
	//
	// 16ACU
	ElasticMinComputeResource *string `json:"ElasticMinComputeResource,omitempty" xml:"ElasticMinComputeResource,omitempty"`
	// Indicates whether the preemptible instance feature is enabled for the resource group. After the preemptible instance feature is enabled, you are charged for resources at a lower unit price but the resources are probably released. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// The True value is returned only for job resource groups.
	//
	// example:
	//
	// True
	EnableSpot   *string                `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	Engine       *string                `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineParams map[string]interface{} `json:"EngineParams,omitempty" xml:"EngineParams,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// >  For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/428610.html).
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The Resource Access Management (RAM) user that is associated with the resource group.
	//
	// example:
	//
	// testb,testc
	GroupUsers *string `json:"GroupUsers,omitempty" xml:"GroupUsers,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources.
	//
	// example:
	//
	// 512ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	MaxGpuQuantity     *int32  `json:"MaxGpuQuantity,omitempty" xml:"MaxGpuQuantity,omitempty"`
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources.
	//
	// example:
	//
	// 0ACU
	MinComputeResource *string                                                 `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	MinGpuQuantity     *int32                                                  `json:"MinGpuQuantity,omitempty" xml:"MinGpuQuantity,omitempty"`
	RayConfig          *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The job resubmission rules.
	Rules []*DescribeDBResourceGroupResponseBodyGroupsInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	RunningClusterCount *int32  `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	SpecName            *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- **creating**: The resource group is being created.
	//
	// 	- **ok**: The resource group is created.
	//
	// 	- **pendingdelete**: The resource group is pending to be deleted.
	//
	// example:
	//
	// ok
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetResourceGroupName *string `json:"TargetResourceGroupName,omitempty" xml:"TargetResourceGroupName,omitempty"`
	// The time when the resource group was updated. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
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

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) GetAutoStopInterval() *string {
	return s.AutoStopInterval
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

func (s *DescribeDBResourceGroupResponseBodyGroupsInfo) SetAutoStopInterval(v string) *DescribeDBResourceGroupResponseBodyGroupsInfo {
	s.AutoStopInterval = &v
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

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig struct {
	// if can be null:
	// true
	AppConfig *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigAppConfig `json:"AppConfig,omitempty" xml:"AppConfig,omitempty" type:"Struct"`
	Category  *string                                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	// if can be null:
	// false
	EnableUserEni    *bool   `json:"EnableUserEni,omitempty" xml:"EnableUserEni,omitempty"`
	HeadAllocateUnit *string `json:"HeadAllocateUnit,omitempty" xml:"HeadAllocateUnit,omitempty"`
	HeadDiskCapacity *string `json:"HeadDiskCapacity,omitempty" xml:"HeadDiskCapacity,omitempty"`
	HeadSpec         *string `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	HeadSpecType     *string `json:"HeadSpecType,omitempty" xml:"HeadSpecType,omitempty"`
	// example:
	//
	// http://ray-cluster-address.example.com
	RayClusterAddress *string `json:"RayClusterAddress,omitempty" xml:"RayClusterAddress,omitempty"`
	// example:
	//
	// http://ray-dashboard-address.example.com
	RayDashboardAddress *string `json:"RayDashboardAddress,omitempty" xml:"RayDashboardAddress,omitempty"`
	// example:
	//
	// http://ray-grafana-address.example.com
	RayGrafanaAddress *string                                                                `json:"RayGrafanaAddress,omitempty" xml:"RayGrafanaAddress,omitempty"`
	StorageMounts     []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts `json:"StorageMounts,omitempty" xml:"StorageMounts,omitempty" type:"Repeated"`
	WorkerGroups      []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups  `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
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

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) GetStorageMounts() []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	return s.StorageMounts
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

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig) SetStorageMounts(v []*DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfig {
	s.StorageMounts = v
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
	// example:
	//
	// app01
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// IsaacLab
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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
	// example:
	//
	// /mnt/data01
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// 1
	StorageId *int64 `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
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

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) SetMountPath(v string) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) SetStorageId(v int64) *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts {
	s.StorageId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigStorageMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyGroupsInfoRayConfigWorkerGroups struct {
	// example:
	//
	// 1
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// example:
	//
	// g01
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1
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
	// large
	WorkerSpecName *string `json:"WorkerSpecName,omitempty" xml:"WorkerSpecName,omitempty"`
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
