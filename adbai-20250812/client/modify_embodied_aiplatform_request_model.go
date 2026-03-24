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
	SetPlatformName(v string) *ModifyEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRayConfig(v *ModifyEmbodiedAIPlatformRequestRayConfig) *ModifyEmbodiedAIPlatformRequest
	GetRayConfig() *ModifyEmbodiedAIPlatformRequestRayConfig
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
	// This parameter is required.
	//
	// example:
	//
	// eap_platform
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// This parameter is required.
	RayConfig *ModifyEmbodiedAIPlatformRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
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

func (s *ModifyEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *ModifyEmbodiedAIPlatformRequest) GetRayConfig() *ModifyEmbodiedAIPlatformRequestRayConfig {
	return s.RayConfig
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

func (s *ModifyEmbodiedAIPlatformRequest) SetPlatformName(v string) *ModifyEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *ModifyEmbodiedAIPlatformRequest) SetRayConfig(v *ModifyEmbodiedAIPlatformRequestRayConfig) *ModifyEmbodiedAIPlatformRequest {
	s.RayConfig = v
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
