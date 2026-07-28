// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetDeviceCount(v int32) *CreateEmbodiedAIPlatformRequest
	GetDeviceCount() *int32
	SetPlatformName(v string) *CreateEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRayConfig(v *CreateEmbodiedAIPlatformRequestRayConfig) *CreateEmbodiedAIPlatformRequest
	GetRayConfig() *CreateEmbodiedAIPlatformRequestRayConfig
	SetRayTrainConfig(v *CreateEmbodiedAIPlatformRequestRayTrainConfig) *CreateEmbodiedAIPlatformRequest
	GetRayTrainConfig() *CreateEmbodiedAIPlatformRequestRayTrainConfig
	SetRegionId(v string) *CreateEmbodiedAIPlatformRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformRequest
	GetWebserverSpecName() *string
}

type CreateEmbodiedAIPlatformRequest struct {
	// The instance cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ontology count.
	//
	// example:
	//
	// 3
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// The name of the embodied intelligence multimodal data platform.
	//
	// > The name can contain lowercase letters, digits, and underscores. It must start with a letter and end with a letter or digit. The name can be up to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// The Ray specification information of the platform.
	RayConfig *CreateEmbodiedAIPlatformRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// The development and training resource configuration.
	RayTrainConfig *CreateEmbodiedAIPlatformRequestRayTrainConfig `json:"RayTrainConfig,omitempty" xml:"RayTrainConfig,omitempty" type:"Struct"`
	// The region ID.
	//
	// > You can call the DescribeRegions operation to query the region ID of a specified Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The webserver specification of the platform.
	//
	// example:
	//
	// large
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
}

func (s CreateEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateEmbodiedAIPlatformRequest) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *CreateEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *CreateEmbodiedAIPlatformRequest) GetRayConfig() *CreateEmbodiedAIPlatformRequestRayConfig {
	return s.RayConfig
}

func (s *CreateEmbodiedAIPlatformRequest) GetRayTrainConfig() *CreateEmbodiedAIPlatformRequestRayTrainConfig {
	return s.RayTrainConfig
}

func (s *CreateEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEmbodiedAIPlatformRequest) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *CreateEmbodiedAIPlatformRequest) SetDBClusterId(v string) *CreateEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetDeviceCount(v int32) *CreateEmbodiedAIPlatformRequest {
	s.DeviceCount = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetPlatformName(v string) *CreateEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetRayConfig(v *CreateEmbodiedAIPlatformRequestRayConfig) *CreateEmbodiedAIPlatformRequest {
	s.RayConfig = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetRayTrainConfig(v *CreateEmbodiedAIPlatformRequestRayTrainConfig) *CreateEmbodiedAIPlatformRequest {
	s.RayTrainConfig = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetRegionId(v string) *CreateEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformRequest {
	s.WebserverSpecName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) Validate() error {
	if s.RayConfig != nil {
		if err := s.RayConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RayTrainConfig != nil {
		if err := s.RayTrainConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEmbodiedAIPlatformRequestRayConfig struct {
	// The Ray cluster type. Valid values:
	//
	// - BASIC: basic type, non-high-availability.
	//
	// - HIGH_AVAILABILITY: high-availability type.
	//
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The node specifications of the head node.
	//
	// example:
	//
	// xlarge
	HeadSpec *string `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	// The list of Ray worker group configurations.
	WorkerGroups []*CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s CreateEmbodiedAIPlatformRequestRayConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayConfig) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) GetCategory() *string {
	return s.Category
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) GetWorkerGroups() []*CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) SetCategory(v string) *CreateEmbodiedAIPlatformRequestRayConfig {
	s.Category = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) SetHeadSpec(v string) *CreateEmbodiedAIPlatformRequestRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) SetWorkerGroups(v []*CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) *CreateEmbodiedAIPlatformRequestRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfig) Validate() error {
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

type CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups struct {
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
	// CPU
	WorkerSpecType *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
}

func (s CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetAllocateUnit(v string) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetGroupName(v string) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerSpecName(v string) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerSpecType(v string) *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type CreateEmbodiedAIPlatformRequestRayTrainConfig struct {
	// The number of CPU ACUs to purchase.
	CpuAcu *int64 `json:"CpuAcu,omitempty" xml:"CpuAcu,omitempty"`
	// The list of GPU model specifications to purchase.
	GpuSpecs []*CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The development machine configuration.
	TerminalConfig *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig `json:"TerminalConfig,omitempty" xml:"TerminalConfig,omitempty" type:"Struct"`
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfig) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) GetCpuAcu() *int64 {
	return s.CpuAcu
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) GetGpuSpecs() []*CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	return s.GpuSpecs
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) GetTerminalConfig() *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig {
	return s.TerminalConfig
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) SetCpuAcu(v int64) *CreateEmbodiedAIPlatformRequestRayTrainConfig {
	s.CpuAcu = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) SetGpuSpecs(v []*CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) *CreateEmbodiedAIPlatformRequestRayTrainConfig {
	s.GpuSpecs = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) SetTerminalConfig(v *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) *CreateEmbodiedAIPlatformRequestRayTrainConfig {
	s.TerminalConfig = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfig) Validate() error {
	if s.GpuSpecs != nil {
		for _, item := range s.GpuSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TerminalConfig != nil {
		if err := s.TerminalConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs struct {
	// The allocation unit.
	//
	// example:
	//
	// "1"
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// The number of GPU cards.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The GPU model.
	//
	// example:
	//
	// ADB.MLGrand.4
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetCount() *int64 {
	return s.Count
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetSpecName() *string {
	return s.SpecName
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetAllocateUnit(v string) *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.AllocateUnit = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetCount(v int64) *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.Count = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetSpecName(v string) *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.SpecName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) Validate() error {
	return dara.Validate(s)
}

type CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig struct {
	// The image repository configuration.
	AcrConfig *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig `json:"AcrConfig,omitempty" xml:"AcrConfig,omitempty" type:"Struct"`
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) GetAcrConfig() *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	return s.AcrConfig
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) SetAcrConfig(v *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig {
	s.AcrConfig = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) Validate() error {
	if s.AcrConfig != nil {
		if err := s.AcrConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig struct {
	// The instance ID of the image repository.
	//
	// example:
	//
	// cri-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The address of the image repository instance.
	//
	// example:
	//
	// example-vpc.example-region.cr.aliyuncs.com
	Registry *string `json:"Registry,omitempty" xml:"Registry,omitempty"`
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetRegistry() *string {
	return s.Registry
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetInstanceId(v string) *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.InstanceId = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetNamespaces(v []*string) *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.Namespaces = v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetRegistry(v string) *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.Registry = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) Validate() error {
	return dara.Validate(s)
}
