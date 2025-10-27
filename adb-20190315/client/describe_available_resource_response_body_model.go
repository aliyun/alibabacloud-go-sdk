// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZoneList(v []*DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody
	GetAvailableZoneList() []*DescribeAvailableResourceResponseBodyAvailableZoneList
	SetRegionId(v string) *DescribeAvailableResourceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeAvailableResourceResponseBody struct {
	// The supported zones.
	AvailableZoneList []*DescribeAvailableResourceResponseBodyAvailableZoneList `json:"AvailableZoneList,omitempty" xml:"AvailableZoneList,omitempty" type:"Repeated"`
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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetAvailableZoneList() []*DescribeAvailableResourceResponseBodyAvailableZoneList {
	return s.AvailableZoneList
}

func (s *DescribeAvailableResourceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZoneList(v []*DescribeAvailableResourceResponseBodyAvailableZoneList) *DescribeAvailableResourceResponseBody {
	s.AvailableZoneList = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRegionId(v string) *DescribeAvailableResourceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
	if s.AvailableZoneList != nil {
		for _, item := range s.AvailableZoneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneList struct {
	// A reserved parameter.
	SupportedComputeResource []*string `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
	// The supported modes.
	SupportedMode []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode `json:"SupportedMode,omitempty" xml:"SupportedMode,omitempty" type:"Repeated"`
	// A reserved parameter.
	SupportedStorageResource []*string `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) GetSupportedComputeResource() []*string {
	return s.SupportedComputeResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) GetSupportedMode() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	return s.SupportedMode
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) GetSupportedStorageResource() []*string {
	return s.SupportedStorageResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedComputeResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedMode(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedMode = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.SupportedStorageResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) SetZoneName(v string) *DescribeAvailableResourceResponseBodyAvailableZoneList {
	s.ZoneName = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneList) Validate() error {
	if s.SupportedMode != nil {
		for _, item := range s.SupportedMode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode struct {
	// The supported mode. Valid values:
	//
	// 	- **flexible**: elastic mode.
	//
	// 	- **reserver**: reserved mode.
	//
	// example:
	//
	// flexible
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The supported editions.
	SupportedSerialList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList `json:"SupportedSerialList,omitempty" xml:"SupportedSerialList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GetMode() *string {
	return s.Mode
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) GetSupportedSerialList() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	return s.SupportedSerialList
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetMode(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) SetSupportedSerialList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode {
	s.SupportedSerialList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedMode) Validate() error {
	if s.SupportedSerialList != nil {
		for _, item := range s.SupportedSerialList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList struct {
	// The supported edition. Valid values:
	//
	// 	- **basic**: Basic Edition.
	//
	// 	- **cluster**: Cluster Edition.
	//
	// 	- **mixed_storage**: elastic mode for Cluster Edition.
	//
	// example:
	//
	// mixed_storage
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The supported resources in elastic mode.
	SupportedFlexibleResource []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource `json:"SupportedFlexibleResource,omitempty" xml:"SupportedFlexibleResource,omitempty" type:"Repeated"`
	// The supported resources in reserved mode.
	SupportedInstanceClassList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList `json:"SupportedInstanceClassList,omitempty" xml:"SupportedInstanceClassList,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GetSerial() *string {
	return s.Serial
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GetSupportedFlexibleResource() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	return s.SupportedFlexibleResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) GetSupportedInstanceClassList() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	return s.SupportedInstanceClassList
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSerial(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.Serial = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedFlexibleResource(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedFlexibleResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) SetSupportedInstanceClassList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList {
	s.SupportedInstanceClassList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialList) Validate() error {
	if s.SupportedFlexibleResource != nil {
		for _, item := range s.SupportedFlexibleResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportedInstanceClassList != nil {
		for _, item := range s.SupportedInstanceClassList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource struct {
	// The disk storage type. Valid values:
	//
	// 	- **hdd**
	//
	// 	- **ssd**
	//
	// example:
	//
	// hdd
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The supported computing resources.
	SupportedComputeResource []*string `json:"SupportedComputeResource,omitempty" xml:"SupportedComputeResource,omitempty" type:"Repeated"`
	// The supported elastic I/O resources.
	SupportedElasticIOResource *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource `json:"SupportedElasticIOResource,omitempty" xml:"SupportedElasticIOResource,omitempty" type:"Struct"`
	// The supported storage resources.
	SupportedStorageResource []*string `json:"SupportedStorageResource,omitempty" xml:"SupportedStorageResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GetSupportedComputeResource() []*string {
	return s.SupportedComputeResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GetSupportedElasticIOResource() *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	return s.SupportedElasticIOResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) GetSupportedStorageResource() []*string {
	return s.SupportedStorageResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetStorageType(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedComputeResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedComputeResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedElasticIOResource(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedElasticIOResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) SetSupportedStorageResource(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource {
	s.SupportedStorageResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResource) Validate() error {
	if s.SupportedElasticIOResource != nil {
		if err := s.SupportedElasticIOResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource struct {
	// The maximum amount of elastic I/O resources.
	//
	// example:
	//
	// 200
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum amount of elastic I/O resources.
	//
	// example:
	//
	// 0
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size.
	//
	// example:
	//
	// 1
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GetMaxCount() *string {
	return s.MaxCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GetMinCount() *string {
	return s.MinCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) GetStep() *string {
	return s.Step
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedFlexibleResourceSupportedElasticIOResource) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList struct {
	// The supported instance type.
	//
	// example:
	//
	// C32
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// A reserved parameter.
	SupportedExecutorList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList `json:"SupportedExecutorList,omitempty" xml:"SupportedExecutorList,omitempty" type:"Repeated"`
	// The supported compute nodes.
	SupportedNodeCountList []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList `json:"SupportedNodeCountList,omitempty" xml:"SupportedNodeCountList,omitempty" type:"Repeated"`
	// The description of the instance type.
	//
	// example:
	//
	// C32
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GetSupportedExecutorList() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList {
	return s.SupportedExecutorList
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GetSupportedNodeCountList() []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	return s.SupportedNodeCountList
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) GetTips() *string {
	return s.Tips
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetInstanceClass(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedExecutorList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedExecutorList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetSupportedNodeCountList(v []*DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.SupportedNodeCountList = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) SetTips(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList {
	s.Tips = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassList) Validate() error {
	if s.SupportedExecutorList != nil {
		for _, item := range s.SupportedExecutorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportedNodeCountList != nil {
		for _, item := range s.SupportedNodeCountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList struct {
	// The information about the supported compute nodes.
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) GetNodeCount() *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	return s.NodeCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorList) Validate() error {
	if s.NodeCount != nil {
		if err := s.NodeCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount struct {
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// N/A
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GetMaxCount() *string {
	return s.MaxCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GetMinCount() *string {
	return s.MinCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) GetStep() *string {
	return s.Step
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedExecutorListNodeCount) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList struct {
	// The number of the supported compute nodes.
	NodeCount *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" type:"Struct"`
	// The support storage capacity. Unit: GB.
	StorageSize []*string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) GetNodeCount() *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	return s.NodeCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) GetStorageSize() []*string {
	return s.StorageSize
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetNodeCount(v *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) SetStorageSize(v []*string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList {
	s.StorageSize = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountList) Validate() error {
	if s.NodeCount != nil {
		if err := s.NodeCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount struct {
	// The maximum number of compute nodes.
	//
	// example:
	//
	// 200
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The minimum number of compute nodes.
	//
	// example:
	//
	// 1
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty"`
	// The step size.
	//
	// example:
	//
	// 1
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GetMaxCount() *string {
	return s.MaxCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GetMinCount() *string {
	return s.MinCount
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) GetStep() *string {
	return s.Step
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMaxCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetMinCount(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) SetStep(v string) *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount {
	s.Step = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZoneListSupportedModeSupportedSerialListSupportedInstanceClassListSupportedNodeCountListNodeCount) Validate() error {
	return dara.Validate(s)
}
