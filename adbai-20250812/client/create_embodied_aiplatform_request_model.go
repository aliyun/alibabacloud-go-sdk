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
	SetPlatformName(v string) *CreateEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRayConfig(v *CreateEmbodiedAIPlatformRequestRayConfig) *CreateEmbodiedAIPlatformRequest
	GetRayConfig() *CreateEmbodiedAIPlatformRequestRayConfig
	SetRegionId(v string) *CreateEmbodiedAIPlatformRequest
	GetRegionId() *string
	SetWebserverSpecName(v string) *CreateEmbodiedAIPlatformRequest
	GetWebserverSpecName() *string
}

type CreateEmbodiedAIPlatformRequest struct {
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
	// platform1
	PlatformName *string                                   `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	RayConfig    *CreateEmbodiedAIPlatformRequestRayConfig `json:"RayConfig,omitempty" xml:"RayConfig,omitempty" type:"Struct"`
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

func (s CreateEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *CreateEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *CreateEmbodiedAIPlatformRequest) GetRayConfig() *CreateEmbodiedAIPlatformRequestRayConfig {
	return s.RayConfig
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

func (s *CreateEmbodiedAIPlatformRequest) SetPlatformName(v string) *CreateEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *CreateEmbodiedAIPlatformRequest) SetRayConfig(v *CreateEmbodiedAIPlatformRequestRayConfig) *CreateEmbodiedAIPlatformRequest {
	s.RayConfig = v
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
	return nil
}

type CreateEmbodiedAIPlatformRequestRayConfig struct {
	// example:
	//
	// BASIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// xlarge
	HeadSpec     *string                                                 `json:"HeadSpec,omitempty" xml:"HeadSpec,omitempty"`
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
