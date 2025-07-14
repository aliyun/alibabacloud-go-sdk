// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody
	GetAvailableZones() *DescribeAvailableResourceResponseBodyAvailableZones
	SetRequestId(v string) *DescribeAvailableResourceResponseBody
	GetRequestId() *string
}

type DescribeAvailableResourceResponseBody struct {
	// Details about the zones in which resources are available.
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0041D94C-FB92-4C49-B115-259DA1C*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) GetAvailableZones() *DescribeAvailableResourceResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *DescribeAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) GetAvailableZone() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	return s.AvailableZone
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	// The resources that are available in the zone.
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of resources in the zone. Valid values:
	//
	// 	- Available
	//
	// 	- SoldOut
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource category based on the stock level in the zone. Valid values:
	//
	// 	- WithStock: Resources are in sufficient stock.
	//
	// 	- ClosedWithStock: Resources are in insufficient stock. We recommend that you use other resources that are in sufficient stock.
	//
	// 	- WithoutStock: Resources are out of stock and will be replenished. We recommend that you use other resources that are in sufficient stock.
	//
	// 	- ClosedWithoutStock: Resources are out of stock and will not be replenished. We recommend that you use other resources that are in sufficient stock.
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetAvailableResources() *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	return s.AvailableResources
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetStatus() *string {
	return s.Status
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetStatus(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.Status = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GetAvailableResource() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	return s.AvailableResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	s.AvailableResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource struct {
	// The information about the resources.
	SupportedResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	// The resource type. Valid values:
	//
	// 	- Zone: zone
	//
	// 	- IoOptimized: I/O optimized resource
	//
	// 	- InstanceType: instance type
	//
	// 	- SystemDisk: system disk
	//
	// 	- DataDisk: data disk
	//
	// 	- Network: network type
	//
	// 	- ddh: dedicated host
	//
	// example:
	//
	// InstanceType
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetSupportedResources() *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	return s.SupportedResources
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GetType() *string {
	return s.Type
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetSupportedResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.SupportedResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.Type = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources struct {
	SupportedResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource `json:"SupportedResource,omitempty" xml:"SupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GetSupportedResource() []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	return s.SupportedResource
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) SetSupportedResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	s.SupportedResource = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) Validate() error {
	return dara.Validate(s)
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource struct {
	// The maximum disk capacity.
	//
	// This parameter takes effect only if DestinationResource is set to SystemDisk or DataDisk.
	//
	// example:
	//
	// 2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum disk capacity.
	//
	// This parameter takes effect only if DestinationResource is set to SystemDisk or DataDisk.
	//
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The status of the resource. Valid values:
	//
	// 	- Available
	//
	// 	- SoldOut
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resource category based on the stock level. Valid values:
	//
	// 	- WithStock: Resources are in sufficient stock.
	//
	// 	- ClosedWithStock: Resources are in insufficient stock. We recommend that you use other resources that are in sufficient stock.
	//
	// 	- WithoutStock: Resources are out of stock and will be replenished. We recommend that you use other resources that are in sufficient stock.
	//
	// 	- ClosedWithoutStock: Resources are out of stock and will not be replenished. We recommend that you use other resources that are in sufficient stock.
	//
	// example:
	//
	// WithStock
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	// The unit of the disk capacity.
	//
	// This parameter takes effect only if DestinationResource is set to SystemDisk or DataDisk.
	//
	// example:
	//
	// null
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The resource.
	//
	// example:
	//
	// ecs.d1ne.xlarge
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMax() *int32 {
	return s.Max
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetMin() *int32 {
	return s.Min
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetStatusCategory() *string {
	return s.StatusCategory
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetUnit() *string {
	return s.Unit
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GetValue() *string {
	return s.Value
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMax(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Max = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetMin(v int32) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Min = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatus(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Status = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetUnit(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Unit = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetValue(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Value = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) Validate() error {
	return dara.Validate(s)
}
