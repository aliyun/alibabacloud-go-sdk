// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
	SetSupportedDBTypes(v *DescribeAvailableResourceResponseBodySupportedDBTypes) *DescribeAvailableResourceResponseBody
	GetSupportedDBTypes() *DescribeAvailableResourceResponseBodySupportedDBTypes
}

type DescribeAvailableResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 344EE51D-8850-4935-B68B-58A8F4DCE3BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported database types.
	SupportedDBTypes *DescribeAvailableResourceResponseBodySupportedDBTypes `json:"SupportedDBTypes,omitempty" xml:"SupportedDBTypes,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) GetSupportedDBTypes() *DescribeAvailableResourceResponseBodySupportedDBTypes {
	return s.SupportedDBTypes
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetSupportedDBTypes(v *DescribeAvailableResourceResponseBodySupportedDBTypes) *DescribeAvailableResourceResponseBody {
	s.SupportedDBTypes = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
	if s.SupportedDBTypes != nil {
		if err := s.SupportedDBTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypes struct {
	SupportedDBType []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType `json:"SupportedDBType,omitempty" xml:"SupportedDBType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypes) GetSupportedDBType() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType {
	return s.SupportedDBType
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypes) SetSupportedDBType(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) *DescribeAvailableResourceResponseBodySupportedDBTypes {
	s.SupportedDBType = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypes) Validate() error {
	if s.SupportedDBType != nil {
		for _, item := range s.SupportedDBType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType struct {
	// The available zones.
	AvailableZones *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// The architecture of the instance. Valid values:
	//
	// 	- **normal**: replica set instance
	//
	// 	- **sharding**: sharded cluster instance
	//
	// example:
	//
	// sharding
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) GetAvailableZones() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones {
	return s.AvailableZones
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) GetDbType() *string {
	return s.DbType
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) SetAvailableZones(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) SetDbType(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType {
	s.DbType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBType) Validate() error {
	if s.AvailableZones != nil {
		if err := s.AvailableZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) GetAvailableZone() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	return s.AvailableZone
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones {
	s.AvailableZone = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZones) Validate() error {
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

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The supported storage engine versions.
	SupportedEngineVersions *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions `json:"SupportedEngineVersions,omitempty" xml:"SupportedEngineVersions,omitempty" type:"Struct"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) GetSupportedEngineVersions() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions {
	return s.SupportedEngineVersions
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetSupportedEngineVersions(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.SupportedEngineVersions = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZone) Validate() error {
	if s.SupportedEngineVersions != nil {
		if err := s.SupportedEngineVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions struct {
	SupportedEngineVersion []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) GetSupportedEngineVersion() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion {
	return s.SupportedEngineVersion
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) SetSupportedEngineVersion(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions {
	s.SupportedEngineVersion = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersions) Validate() error {
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

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion struct {
	// The supported storage engines.
	SupportedEngines *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Struct"`
	// The database engine version of the instance.
	//
	// example:
	//
	// 4.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) GetSupportedEngines() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines {
	return s.SupportedEngines
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) SetSupportedEngines(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) SetVersion(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion {
	s.Version = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersion) Validate() error {
	if s.SupportedEngines != nil {
		if err := s.SupportedEngines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines struct {
	SupportedEngine []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine `json:"SupportedEngine,omitempty" xml:"SupportedEngine,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) GetSupportedEngine() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine {
	return s.SupportedEngine
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) SetSupportedEngine(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines {
	s.SupportedEngine = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEngines) Validate() error {
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

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine struct {
	// The storage engine of the instance.
	//
	// example:
	//
	// WiredTiger
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The supported instance types.
	SupportedNodeTypes *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes `json:"SupportedNodeTypes,omitempty" xml:"SupportedNodeTypes,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) GetSupportedNodeTypes() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes {
	return s.SupportedNodeTypes
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) SetEngine(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) SetSupportedNodeTypes(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine {
	s.SupportedNodeTypes = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngine) Validate() error {
	if s.SupportedNodeTypes != nil {
		if err := s.SupportedNodeTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes struct {
	SupportedNodeType []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType `json:"SupportedNodeType,omitempty" xml:"SupportedNodeType,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) GetSupportedNodeType() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	return s.SupportedNodeType
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) SetSupportedNodeType(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes {
	s.SupportedNodeType = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypes) Validate() error {
	if s.SupportedNodeType != nil {
		for _, item := range s.SupportedNodeType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType struct {
	// The details of the available resources.
	AvailableResources *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The network type of the instance.
	//
	// example:
	//
	// VPC
	NetworkTypes *string `json:"NetworkTypes,omitempty" xml:"NetworkTypes,omitempty"`
	// The number of nodes in the instance.
	//
	// example:
	//
	// 3
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) GetAvailableResources() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources {
	return s.AvailableResources
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) GetNetworkTypes() *string {
	return s.NetworkTypes
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetAvailableResources(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetNetworkTypes(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.NetworkTypes = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) SetNodeType(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType {
	s.NodeType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeType) Validate() error {
	if s.AvailableResources != nil {
		if err := s.AvailableResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) GetAvailableResource() []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	return s.AvailableResource
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources {
	s.AvailableResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResources) Validate() error {
	if s.AvailableResource != nil {
		for _, item := range s.AvailableResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource struct {
	// The storage capacity range of the instance.
	DBInstanceStorageRange *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange `json:"DBInstanceStorageRange,omitempty" xml:"DBInstanceStorageRange,omitempty" type:"Struct"`
	// The instance family.
	//
	// example:
	//
	// mdb.shard.2x.xlarge.d
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// 4 cores, 8 GB (Dedicated) (Current instance type: mdb.shard.2x.xlarge.d (4 cores, 8 GB (Dedicated cloud-disk), maximum connections: 3000, maximum IOPS: min{1800 + 50 × Storage capacity, 21000}))
	InstanceClassRemark *string `json:"InstanceClassRemark,omitempty" xml:"InstanceClassRemark,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) GetDBInstanceStorageRange() *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange {
	return s.DBInstanceStorageRange
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) GetInstanceClassRemark() *string {
	return s.InstanceClassRemark
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) SetDBInstanceStorageRange(v *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	s.DBInstanceStorageRange = v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) SetInstanceClassRemark(v string) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource {
	s.InstanceClassRemark = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResource) Validate() error {
	if s.DBInstanceStorageRange != nil {
		if err := s.DBInstanceStorageRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange struct {
	// The maximum storage capacity. Unit: GB.
	//
	// example:
	//
	// 16000
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum storage capacity. Unit: GB.
	//
	// example:
	//
	// 20
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The step size for adjusting the storage capacity. Unit: GB.
	//
	// example:
	//
	// 10
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) GetMax() *int32 {
	return s.Max
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) GetMin() *int32 {
	return s.Min
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) GetStep() *int32 {
	return s.Step
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) SetMax(v int32) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange {
	s.Max = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) SetMin(v int32) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange {
	s.Min = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) SetStep(v int32) *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodySupportedDBTypesSupportedDBTypeAvailableZonesAvailableZoneSupportedEngineVersionsSupportedEngineVersionSupportedEnginesSupportedEngineSupportedNodeTypesSupportedNodeTypeAvailableResourcesAvailableResourceDBInstanceStorageRange) Validate() error {
	return dara.Validate(s)
}
