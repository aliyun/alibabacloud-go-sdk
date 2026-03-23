// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeAvailableZonesResponseBodyAvailableZones) *DescribeAvailableZonesResponseBody
	GetAvailableZones() []*DescribeAvailableZonesResponseBodyAvailableZones
	SetRequestId(v string) *DescribeAvailableZonesResponseBody
	GetRequestId() *string
}

type DescribeAvailableZonesResponseBody struct {
	AvailableZones []*DescribeAvailableZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBody) GetAvailableZones() []*DescribeAvailableZonesResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableZonesResponseBody) SetAvailableZones(v []*DescribeAvailableZonesResponseBodyAvailableZones) *DescribeAvailableZonesResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableZonesResponseBody) SetRequestId(v string) *DescribeAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableZonesResponseBody) Validate() error {
	if s.AvailableZones != nil {
		for _, item := range s.AvailableZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyAvailableZones struct {
	RegionId         *string                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportedEngines []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	ZoneId           *string                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableZonesResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) GetSupportedEngines() []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines {
	return s.SupportedEngines
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) SetRegionId(v string) *DescribeAvailableZonesResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) SetSupportedEngines(v []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) *DescribeAvailableZonesResponseBodyAvailableZones {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) SetZoneId(v string) *DescribeAvailableZonesResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZones) Validate() error {
	if s.SupportedEngines != nil {
		for _, item := range s.SupportedEngines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines struct {
	Engine                  *string                                                                                    `json:"Engine,omitempty" xml:"Engine,omitempty"`
	SupportedEngineVersions []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Repeated"`
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) GetSupportedEngineVersions() []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions {
	return s.SupportedEngineVersions
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) SetEngine(v string) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) SetSupportedEngineVersions(v []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEngines) Validate() error {
	if s.SupportedEngineVersions != nil {
		for _, item := range s.SupportedEngineVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions struct {
	SupportedCategorys []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys `json:"SupportedCategorys,omitempty" xml:"SupportedCategorys,omitempty" type:"Repeated"`
	Version            *string                                                                                                      `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) GetSupportedCategorys() []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys {
	return s.SupportedCategorys
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) GetVersion() *string {
	return s.Version
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) SetSupportedCategorys(v []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions {
	s.SupportedCategorys = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) SetVersion(v string) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions {
	s.Version = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersions) Validate() error {
	if s.SupportedCategorys != nil {
		for _, item := range s.SupportedCategorys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys struct {
	Category              *string                                                                                                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	SupportedStorageTypes []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes `json:"SupportedStorageTypes,omitempty" xml:"SupportedStorageTypes,omitempty" type:"Repeated"`
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) GetSupportedStorageTypes() []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes {
	return s.SupportedStorageTypes
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) SetCategory(v string) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys {
	s.Category = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) SetSupportedStorageTypes(v []*DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys {
	s.SupportedStorageTypes = v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorys) Validate() error {
	if s.SupportedStorageTypes != nil {
		for _, item := range s.SupportedStorageTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes struct {
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) SetStorageType(v string) *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableZonesResponseBodyAvailableZonesSupportedEnginesSupportedEngineVersionsSupportedCategorysSupportedStorageTypes) Validate() error {
	return dara.Validate(s)
}
