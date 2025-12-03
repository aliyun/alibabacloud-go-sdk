// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) *DescribeMultiZoneAvailableResourceResponseBody
	GetAvailableZones() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones
	SetRequestId(v string) *DescribeMultiZoneAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeMultiZoneAvailableResourceResponseBody struct {
	AvailableZones *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// example:
	//
	// B2EEBBA9-C627-4415-81A0-B77BC54F1D52
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) GetAvailableZones() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetAvailableZones(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) *DescribeMultiZoneAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBody) Validate() error {
	if s.AvailableZones != nil {
		if err := s.AvailableZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) GetAvailableZone() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	return s.AvailableZone
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZones) Validate() error {
	if s.AvailableZone != nil {
		for _, item := range s.AvailableZone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	MasterResources *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources `json:"MasterResources,omitempty" xml:"MasterResources,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId         *string                                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportedEngines *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou-bef-aliyun
	ZoneCombination *string `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GetMasterResources() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	return s.MasterResources
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GetSupportedEngines() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	return s.SupportedEngines
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) GetZoneCombination() *string {
	return s.ZoneCombination
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetMasterResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.MasterResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetSupportedEngines(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.SupportedEngines = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZone) Validate() error {
	if s.MasterResources != nil {
		if err := s.MasterResources.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedEngines != nil {
		if err := s.SupportedEngines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources struct {
	MasterResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource `json:"MasterResource,omitempty" xml:"MasterResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) GetMasterResource() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	return s.MasterResource
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) SetMasterResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources {
	s.MasterResource = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResources) Validate() error {
	if s.MasterResource != nil {
		for _, item := range s.MasterResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource struct {
	// example:
	//
	// hbase.sn2.large
	InstanceType       *string                                                                                                                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) GetInstanceTypeDetail() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	return s.InstanceTypeDetail
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResource) Validate() error {
	if s.InstanceTypeDetail != nil {
		if err := s.InstanceTypeDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 16
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) GetMem() *int32 {
	return s.Mem
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneMasterResourcesMasterResourceInstanceTypeDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines struct {
	SupportedEngine []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) GetSupportedEngine() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	return s.SupportedEngine
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) SetSupportedEngine(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines {
	s.SupportedEngine = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEngines) Validate() error {
	if s.SupportedEngine != nil {
		for _, item := range s.SupportedEngine {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine struct {
	// example:
	//
	// hbaseue
	Engine                  *string                                                                                                                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
	SupportedEngineVersions *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GetEngine() *string {
	return s.Engine
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) GetSupportedEngineVersions() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	return s.SupportedEngineVersions
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) SetSupportedEngineVersions(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngine) Validate() error {
	if s.SupportedEngineVersions != nil {
		if err := s.SupportedEngineVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) GetSupportedEngineVersion() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	return s.SupportedEngineVersion
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersions) Validate() error {
	if s.SupportedEngineVersion != nil {
		for _, item := range s.SupportedEngineVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion struct {
	SupportedCategories *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Struct"`
	// example:
	//
	// 2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GetSupportedCategories() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	return s.SupportedCategories
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetSupportedCategories(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedCategories = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersion) Validate() error {
	if s.SupportedCategories != nil {
		if err := s.SupportedCategories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories struct {
	SupportedCategories []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories `json:"SupportedCategories,omitempty" xml:"SupportedCategories,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) GetSupportedCategories() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	return s.SupportedCategories
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) SetSupportedCategories(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories {
	s.SupportedCategories = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategories) Validate() error {
	if s.SupportedCategories != nil {
		for _, item := range s.SupportedCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories struct {
	// example:
	//
	// cluster
	Category              *string                                                                                                                                                                                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	SupportedStorageTypes *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GetCategory() *string {
	return s.Category
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) GetSupportedStorageTypes() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	return s.SupportedStorageTypes
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetCategory(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.Category = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) SetSupportedStorageTypes(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories {
	s.SupportedStorageTypes = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategories) Validate() error {
	if s.SupportedStorageTypes != nil {
		if err := s.SupportedStorageTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes struct {
	SupportedStorageType []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType `json:"SupportedStorageType,omitempty" xml:"SupportedStorageType,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) GetSupportedStorageType() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	return s.SupportedStorageType
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) SetSupportedStorageType(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes {
	s.SupportedStorageType = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypes) Validate() error {
	if s.SupportedStorageType != nil {
		for _, item := range s.SupportedStorageType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType struct {
	CoreResources *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources `json:"CoreResources,omitempty" xml:"CoreResources,omitempty" type:"Struct"`
	// example:
	//
	// cloud_efficiency
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GetCoreResources() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	return s.CoreResources
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetCoreResources(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.CoreResources = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) SetStorageType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType {
	s.StorageType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageType) Validate() error {
	if s.CoreResources != nil {
		if err := s.CoreResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources struct {
	CoreResource []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource `json:"CoreResource,omitempty" xml:"CoreResource,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) GetCoreResource() []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	return s.CoreResource
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) SetCoreResource(v []*DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources {
	s.CoreResource = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResources) Validate() error {
	if s.CoreResource != nil {
		for _, item := range s.CoreResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource struct {
	DBInstanceStorageRange *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	// example:
	//
	// hbase.sn2.2xlarge
	InstanceType       *string                                                                                                                                                                                                                                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeDetail *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail `json:"InstanceTypeDetail,omitempty" xml:"InstanceTypeDetail,omitempty" type:"Struct"`
	// example:
	//
	// 30
	MaxCoreCount *int32 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GetDBInstanceStorageRange() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	return s.DBInstanceStorageRange
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GetInstanceTypeDetail() *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	return s.InstanceTypeDetail
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) GetMaxCoreCount() *int32 {
	return s.MaxCoreCount
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetDBInstanceStorageRange(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceType(v string) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetInstanceTypeDetail(v *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.InstanceTypeDetail = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) SetMaxCoreCount(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource {
	s.MaxCoreCount = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResource) Validate() error {
	if s.DBInstanceStorageRange != nil {
		if err := s.DBInstanceStorageRange.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceTypeDetail != nil {
		if err := s.InstanceTypeDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange struct {
	// example:
	//
	// 64000
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// example:
	//
	// 400
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// example:
	//
	// 40
	StepSize *int32 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GetMinSize() *int32 {
	return s.MinSize
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) GetStepSize() *int32 {
	return s.StepSize
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMaxSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MaxSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetMinSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.MinSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) SetStepSize(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange {
	s.StepSize = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceDBInstanceStorageRange) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail struct {
	// example:
	//
	// 32
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 8
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) GetMem() *int32 {
	return s.Mem
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetCpu(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) SetMem(v int32) *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail {
	s.Mem = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponseBodyAvailableZonesAvailableZoneSupportedEnginesSupportedEngineSupportedEngineVersionsSupportedEngineVersionSupportedCategoriesSupportedCategoriesSupportedStorageTypesSupportedStorageTypeCoreResourcesCoreResourceInstanceTypeDetail) Validate() error {
	return dara.Validate(s)
}
