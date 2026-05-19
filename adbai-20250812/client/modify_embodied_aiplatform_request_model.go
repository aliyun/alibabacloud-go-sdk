// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetDeviceCount(v string) *ModifyEmbodiedAIPlatformRequest
	GetDeviceCount() *string
	SetPlatformName(v string) *ModifyEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRayConfig(v *ModifyEmbodiedAIPlatformRequestRayConfig) *ModifyEmbodiedAIPlatformRequest
	GetRayConfig() *ModifyEmbodiedAIPlatformRequestRayConfig
	SetRayTrainConfig(v *ModifyEmbodiedAIPlatformRequestRayTrainConfig) *ModifyEmbodiedAIPlatformRequest
	GetRayTrainConfig() *ModifyEmbodiedAIPlatformRequestRayTrainConfig
	SetRegionId(v string) *ModifyEmbodiedAIPlatformRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *ModifyEmbodiedAIPlatformRequest
	GetWebserverSpecName() *string
}

type ModifyEmbodiedAIPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DeviceCount *string `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eap_platform
	PlatformName   *string                                        `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	RayConfig      *ModifyEmbodiedAIPlatformRequestRayConfig      `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	RayTrainConfig *ModifyEmbodiedAIPlatformRequestRayTrainConfig `json:"RayTrainConfig,omitempty" xml:"RayTrainConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// large
	WebserverSpecName *string `json:"WebserverSpecName,omitempty" xml:"WebserverSpecName,omitempty"`
}

func (s ModifyEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEmbodiedAIPlatformRequest) GetDeviceCount() *string {
	return s.DeviceCount
}

func (s *ModifyEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ModifyEmbodiedAIPlatformRequest) GetRayConfig() *ModifyEmbodiedAIPlatformRequestRayConfig {
	return s.RayConfig
}

func (s *ModifyEmbodiedAIPlatformRequest) GetRayTrainConfig() *ModifyEmbodiedAIPlatformRequestRayTrainConfig {
	return s.RayTrainConfig
}

func (s *ModifyEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEmbodiedAIPlatformRequest) GetWebserverSpecName() *string {
	return s.WebserverSpecName
}

func (s *ModifyEmbodiedAIPlatformRequest) SetDBClusterId(v string) *ModifyEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetDeviceCount(v string) *ModifyEmbodiedAIPlatformRequest {
	s.DeviceCount = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetPlatformName(v string) *ModifyEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetRayConfig(v *ModifyEmbodiedAIPlatformRequestRayConfig) *ModifyEmbodiedAIPlatformRequest {
	s.RayConfig = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetRayTrainConfig(v *ModifyEmbodiedAIPlatformRequestRayTrainConfig) *ModifyEmbodiedAIPlatformRequest {
	s.RayTrainConfig = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetRegionId(v string) *ModifyEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetWebserverSpecName(v string) *ModifyEmbodiedAIPlatformRequest {
	s.WebserverSpecName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) Validate() error {
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

type ModifyEmbodiedAIPlatformRequestRayConfig struct {
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// xlarge
	HeadSpec     *string                                                 `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
	WorkerGroups []*ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups `json:"WorkerGroups,omitempty" xml:"WorkerGroups,omitempty" type:"Repeated"`
}

func (s ModifyEmbodiedAIPlatformRequestRayConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayConfig) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) GetCategory() *string {
	return s.Category
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) GetHeadSpec() *string {
	return s.HeadSpec
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) GetWorkerGroups() []*ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	return s.WorkerGroups
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) SetCategory(v string) *ModifyEmbodiedAIPlatformRequestRayConfig {
	s.Category = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) SetHeadSpec(v string) *ModifyEmbodiedAIPlatformRequestRayConfig {
	s.HeadSpec = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) SetWorkerGroups(v []*ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) *ModifyEmbodiedAIPlatformRequestRayConfig {
	s.WorkerGroups = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfig) Validate() error {
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

type ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups struct {
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
	// CPU
	WorkerSpecType *string `json:"WorkerSpecType,omitempty" xml:"WorkerSpecType,omitempty"`
}

func (s ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetMaxWorkerQuantity() *int32 {
	return s.MaxWorkerQuantity
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetMinWorkerQuantity() *int32 {
	return s.MinWorkerQuantity
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerDiskCapacity() *string {
	return s.WorkerDiskCapacity
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerSpecName() *string {
	return s.WorkerSpecName
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) GetWorkerSpecType() *string {
	return s.WorkerSpecType
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetAllocateUnit(v string) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.AllocateUnit = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetGroupName(v string) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.GroupName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetMaxWorkerQuantity(v int32) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.MaxWorkerQuantity = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetMinWorkerQuantity(v int32) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.MinWorkerQuantity = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerDiskCapacity(v string) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerDiskCapacity = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerSpecName(v string) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerSpecName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) SetWorkerSpecType(v string) *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups {
	s.WorkerSpecType = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayConfigWorkerGroups) Validate() error {
	return dara.Validate(s)
}

type ModifyEmbodiedAIPlatformRequestRayTrainConfig struct {
	// example:
	//
	// 10
	CpuAcu         *int64                                                       `json:"CpuAcu,omitempty" xml:"CpuAcu,omitempty"`
	GpuSpecs       []*ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs     `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	TerminalConfig *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig `json:"TerminalConfig,omitempty" xml:"TerminalConfig,omitempty" type:"Struct"`
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfig) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) GetCpuAcu() *int64 {
	return s.CpuAcu
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) GetGpuSpecs() []*ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	return s.GpuSpecs
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) GetTerminalConfig() *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig {
	return s.TerminalConfig
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) SetCpuAcu(v int64) *ModifyEmbodiedAIPlatformRequestRayTrainConfig {
	s.CpuAcu = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) SetGpuSpecs(v []*ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) *ModifyEmbodiedAIPlatformRequestRayTrainConfig {
	s.GpuSpecs = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) SetTerminalConfig(v *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) *ModifyEmbodiedAIPlatformRequestRayTrainConfig {
	s.TerminalConfig = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfig) Validate() error {
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

type ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs struct {
	// example:
	//
	// "1"
	AllocateUnit *string `json:"AllocateUnit,omitempty" xml:"AllocateUnit,omitempty"`
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ADB.MLTensor.2
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetAllocateUnit() *string {
	return s.AllocateUnit
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetCount() *int64 {
	return s.Count
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) GetSpecName() *string {
	return s.SpecName
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetAllocateUnit(v string) *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.AllocateUnit = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetCount(v int64) *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.Count = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) SetSpecName(v string) *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs {
	s.SpecName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigGpuSpecs) Validate() error {
	return dara.Validate(s)
}

type ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig struct {
	AcrConfig *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig `json:"AcrConfig,omitempty" xml:"AcrConfig,omitempty" type:"Struct"`
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) GetAcrConfig() *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	return s.AcrConfig
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) SetAcrConfig(v *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig {
	s.AcrConfig = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfig) Validate() error {
	if s.AcrConfig != nil {
		if err := s.AcrConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig struct {
	// example:
	//
	// i-8vb32jdd306b50rodza7
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// example-vpc.example-region.cr.aliyuncs.com
	Registry *string `json:"Registry,omitempty" xml:"Registry,omitempty"`
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GoString() string {
	return s.String()
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) GetRegistry() *string {
	return s.Registry
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetInstanceId(v string) *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.InstanceId = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetNamespaces(v []*string) *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.Namespaces = v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) SetRegistry(v string) *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig {
	s.Registry = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequestRayTrainConfigTerminalConfigAcrConfig) Validate() error {
	return dara.Validate(s)
}
