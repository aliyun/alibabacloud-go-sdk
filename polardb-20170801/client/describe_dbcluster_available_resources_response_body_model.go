// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAvailableResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) *DescribeDBClusterAvailableResourcesResponseBody
	GetAvailableZones() []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones
	SetRequestId(v string) *DescribeDBClusterAvailableResourcesResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAvailableResourcesResponseBody struct {
	// The available zones of the cluster.
	AvailableZones []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2B19F698-8FFC-4918-B9E2-58D878******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) GetAvailableZones() []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) SetAvailableZones(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) *DescribeDBClusterAvailableResourcesResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) SetRequestId(v string) *DescribeDBClusterAvailableResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZones struct {
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The database engines that the available resources support.
	SupportedEngines []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) GetSupportedEngines() []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines {
	return s.SupportedEngines
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetRegionId(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetSupportedEngines(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.SupportedEngines = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetZoneId(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines struct {
	// The available resources.
	AvailableResources []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	// The version of the database engine.
	//
	// example:
	//
	// mysql57
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) GetAvailableResources() []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources {
	return s.AvailableResources
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) SetAvailableResources(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines {
	s.AvailableResources = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) SetEngine(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources struct {
	// The edition of the cluster. Valid values:
	//
	// 	- **Normal**: Cluster Edition.
	//
	// 	- **Basic**: Single Node Edition.
	//
	// 	- **ArchiveNormal**: X-Engine.
	//
	// 	- **NormalMultimaster**: Multi-master Cluster (Database/Table) Edition.
	//
	// 	- **SENormal**: Standard Edition.
	//
	// >- Only PolarDB for MySQL supports Single Node Edition.
	//
	// >- Only PolarDB for MySQL 8.0 supports X-Engine Edition and Multi-master Cluster (Database/Table) Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The specifications of the node.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) SetCategory(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) Validate() error {
	return dara.Validate(s)
}
