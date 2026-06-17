// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAvailableResourcesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAvailableResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeAvailableResourcesResponseBodyResources) *DescribeAvailableResourcesResponseBody
	GetResources() []*DescribeAvailableResourcesResponseBodyResources
}

type DescribeAvailableResourcesResponseBody struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 61DC563C-F8E4-593A-8D27-CE**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The available resources.
	Resources []*DescribeAvailableResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourcesResponseBody) GetResources() []*DescribeAvailableResourcesResponseBodyResources {
	return s.Resources
}

func (s *DescribeAvailableResourcesResponseBody) SetRegionId(v string) *DescribeAvailableResourcesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetRequestId(v string) *DescribeAvailableResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) SetResources(v []*DescribeAvailableResourcesResponseBodyResources) *DescribeAvailableResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeAvailableResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourcesResponseBodyResources struct {
	// Details of the supported engines and instance types.
	SupportedEngines []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResources) GetSupportedEngines() []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	return s.SupportedEngines
}

func (s *DescribeAvailableResourcesResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetSupportedEngines(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) *DescribeAvailableResourcesResponseBodyResources {
	s.SupportedEngines = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResources) SetZoneId(v string) *DescribeAvailableResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResources) Validate() error {
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

type DescribeAvailableResourcesResponseBodyResourcesSupportedEngines struct {
	// The instance resource type. Valid values:
	//
	// - **ecs**: Elastic storage mode.
	//
	// - **serverless**: Serverless mode.
	//
	// example:
	//
	// ecs
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The supported engine version.
	//
	// example:
	//
	// 6.0
	SupportedEngineVersion *string `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty"`
	// The supported instance types.
	SupportedInstanceClasses []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses `json:"SupportedInstanceClasses,omitempty" xml:"SupportedInstanceClasses,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GetMode() *string {
	return s.Mode
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GetSupportedEngineVersion() *string {
	return s.SupportedEngineVersion
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) GetSupportedInstanceClasses() []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	return s.SupportedInstanceClasses
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetMode(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedEngineVersion(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedEngineVersion = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) SetSupportedInstanceClasses(v []*DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines {
	s.SupportedInstanceClasses = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEngines) Validate() error {
	if s.SupportedInstanceClasses != nil {
		for _, item := range s.SupportedInstanceClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses struct {
	// The instance series. Valid values:
	//
	// - **HighAvailability**: High-availability edition.
	//
	// - **Basic**: Basic edition
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the segment node instance type.
	//
	// example:
	//
	// Single-segment node with 2 cores and 16 GB of memory.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the segment node instance type.
	//
	// example:
	//
	// 2C16G
	DisplayClass *string `json:"DisplayClass,omitempty" xml:"DisplayClass,omitempty"`
	// The instance type of the segment node.
	//
	// example:
	//
	// 2C16G
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// Details of the segment nodes.
	NodeCount *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	// The Segment storage capacity.
	StorageSize *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Struct"`
	// The storage type. Valid values:
	//
	// - **cloud_essd**: ESSD cloud disk.
	//
	// - **cloud_efficiency**: Standard cloud disk.
	//
	// - **oss**:OSS.
	//
	// example:
	//
	// cloud_essd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetDescription() *string {
	return s.Description
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetDisplayClass() *string {
	return s.DisplayClass
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetNodeCount() *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	return s.NodeCount
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetStorageSize() *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	return s.StorageSize
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetCategory(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDescription(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.Description = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetDisplayClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.DisplayClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetInstanceClass(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetNodeCount(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageSize(v *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageSize = v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) SetStorageType(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClasses) Validate() error {
	if s.NodeCount != nil {
		if err := s.NodeCount.Validate(); err != nil {
			return err
		}
	}
	if s.StorageSize != nil {
		if err := s.StorageSize.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount struct {
	// The maximum number of Segment nodes.
	//
	// example:
	//
	// 256
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of Segment nodes.
	//
	// example:
	//
	// 4
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size for adding nodes.
	//
	// For example, if this parameter is set to 4, the number of nodes added must be a multiple of 4 when adding Segment nodes.
	//
	// example:
	//
	// 4
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GetMaxCount() *string {
	return s.MaxCount
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GetMinCount() *string {
	return s.MinCount
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GetStep() *string {
	return s.Step
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesNodeCount) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize struct {
	// The maximum Segment storage capacity.
	//
	// example:
	//
	// 1000
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum Segment storage capacity.
	//
	// example:
	//
	// 50
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The Segment storage expansion step size.
	//
	// example:
	//
	// 50
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GetMaxCount() *string {
	return s.MaxCount
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GetMinCount() *string {
	return s.MinCount
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GetStep() *string {
	return s.Step
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMaxCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMinCount(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetStep(v string) *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourcesResponseBodyResourcesSupportedEnginesSupportedInstanceClassesStorageSize) Validate() error {
	return dara.Validate(s)
}
