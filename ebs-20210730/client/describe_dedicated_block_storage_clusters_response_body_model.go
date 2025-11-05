// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedBlockStorageClusters(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) *DescribeDedicatedBlockStorageClustersResponseBody
	GetDedicatedBlockStorageClusters() []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters
	SetNextToken(v string) *DescribeDedicatedBlockStorageClustersResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDedicatedBlockStorageClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDedicatedBlockStorageClustersResponseBody
	GetTotalCount() *int64
}

type DescribeDedicatedBlockStorageClustersResponseBody struct {
	// Details about the dedicated block storage clusters.
	DedicatedBlockStorageClusters []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters `json:"DedicatedBlockStorageClusters,omitempty" xml:"DedicatedBlockStorageClusters,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetDedicatedBlockStorageClusters() []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	return s.DedicatedBlockStorageClusters
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetDedicatedBlockStorageClusters(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.DedicatedBlockStorageClusters = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetPageNumber(v int32) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetPageSize(v int32) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetTotalCount(v int64) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) Validate() error {
	if s.DedicatedBlockStorageClusters != nil {
		for _, item := range s.DedicatedBlockStorageClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters struct {
	// The user ID.
	//
	// example:
	//
	// 12345601234560***
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The category of disks that can be created in the dedicated block storage cluster.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the dedicated block storage cluster was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1657113211
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Details about the storage capacity of the dedicated block storage cluster.
	DedicatedBlockStorageClusterCapacity *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity `json:"DedicatedBlockStorageClusterCapacity,omitempty" xml:"DedicatedBlockStorageClusterCapacity,omitempty" type:"Struct"`
	// The ID of the dedicated block storage cluster.
	//
	// example:
	//
	// dbsc-f8z4d3k4nsgg9okb****
	DedicatedBlockStorageClusterId *string `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty"`
	// The name of the dedicated block storage cluster.
	//
	// example:
	//
	// myDBSCCluster
	DedicatedBlockStorageClusterName *string `json:"DedicatedBlockStorageClusterName,omitempty" xml:"DedicatedBlockStorageClusterName,omitempty"`
	// The description of the dedicated block storage cluster.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether Thin Provision is enabled.
	//
	// example:
	//
	// true
	EnableThinProvision *bool `json:"EnableThinProvision,omitempty" xml:"EnableThinProvision,omitempty"`
	// The time when the dedicated block storage cluster expires. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1673020800
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The performance level of disks. Valid values:
	//
	// 	- PL0
	//
	// 	- PL1
	//
	// 	- PL2
	//
	// 	- PL3
	//
	// >  This parameter is valid only when the SupportedCategory value is cloud_essd.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The region ID of the dedicated block storage cluster.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated block storage cluster belongs. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the ID of the resource group.
	//
	// example:
	//
	// rg-aekzsoux****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The capacity oversold ratio.
	//
	// example:
	//
	// 1.2
	SizeOverSoldRatio *float64 `json:"SizeOverSoldRatio,omitempty" xml:"SizeOverSoldRatio,omitempty"`
	// The state of the dedicated block storage cluster. Valid values:
	//
	// 	- Preparing
	//
	// 	- Running
	//
	// 	- Expired
	//
	// 	- Offline
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// StorageDomain
	//
	// example:
	//
	// StorageDomain
	StorageDomain *string `json:"StorageDomain,omitempty" xml:"StorageDomain,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// cloud_essd
	SupportedCategory *string `json:"SupportedCategory,omitempty" xml:"SupportedCategory,omitempty"`
	// The tags of the dedicated block storage cluster.
	Tags []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the dedicated block storage cluster. Valid values:
	//
	// 	- Standard: basic dedicated block storage cluster. ESSDs at performance level 0 (PL0 ESSDs) can be created in basic dedicated block storage clusters.
	//
	// 	- Premium: performance dedicated block storage cluster. ESSDs at performance level 1 (PL1 ESSDs) can be created in performance dedicated block storage clusters.
	//
	// example:
	//
	// Standard
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of the dedicated block storage cluster.
	//
	// example:
	//
	// cn-heyuan-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetCategory() *string {
	return s.Category
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetDedicatedBlockStorageClusterCapacity() *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	return s.DedicatedBlockStorageClusterCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetDedicatedBlockStorageClusterId() *string {
	return s.DedicatedBlockStorageClusterId
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetDedicatedBlockStorageClusterName() *string {
	return s.DedicatedBlockStorageClusterName
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetDescription() *string {
	return s.Description
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetEnableThinProvision() *bool {
	return s.EnableThinProvision
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetSizeOverSoldRatio() *float64 {
	return s.SizeOverSoldRatio
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetStorageDomain() *string {
	return s.StorageDomain
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetSupportedCategory() *string {
	return s.SupportedCategory
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetTags() []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags {
	return s.Tags
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetType() *string {
	return s.Type
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetAliUid(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.AliUid = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCreateTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterCapacity(v *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterCapacity = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterName(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDescription(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetEnableThinProvision(v bool) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.EnableThinProvision = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetExpiredTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetResourceGroupId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetSizeOverSoldRatio(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.SizeOverSoldRatio = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetStatus(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetStorageDomain(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.StorageDomain = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetSupportedCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.SupportedCategory = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetTags(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetType(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetZoneId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) Validate() error {
	if s.DedicatedBlockStorageClusterCapacity != nil {
		if err := s.DedicatedBlockStorageClusterCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity struct {
	// The available capacity of the dedicated block storage cluster. Unit: GiB.
	//
	// example:
	//
	// 61440
	AvailableCapacity *int64 `json:"AvailableCapacity,omitempty" xml:"AvailableCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster that was delivered in disk creation orders. Unit: GB.
	//
	// example:
	//
	// 61440
	AvailableDeviceCapacity *int64 `json:"AvailableDeviceCapacity,omitempty" xml:"AvailableDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	//
	// example:
	//
	// 40000.3
	AvailableSpaceCapacity *float64 `json:"AvailableSpaceCapacity,omitempty" xml:"AvailableSpaceCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that was delivered in orders. Unit: GB.
	//
	// example:
	//
	// 61440
	ClusterAvailableCapacity *int64 `json:"ClusterAvailableCapacity,omitempty" xml:"ClusterAvailableCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that is to be delivered in orders. Unit: GB.
	//
	// example:
	//
	// 0
	ClusterDeliveryCapacity *int64 `json:"ClusterDeliveryCapacity,omitempty" xml:"ClusterDeliveryCapacity,omitempty"`
	// The capacity to be delivered for the dedicated block storage cluster. Unit: GiB.
	//
	// example:
	//
	// 0
	DeliveryCapacity *int64 `json:"DeliveryCapacity,omitempty" xml:"DeliveryCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster. Unit: GiB.
	//
	// example:
	//
	// 61440
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The total capacity of the dedicated block storage cluster that is to be delivered in disk creation orders. Unit: GB.
	//
	// example:
	//
	// 61440
	TotalDeviceCapacity *int64 `json:"TotalDeviceCapacity,omitempty" xml:"TotalDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	//
	// example:
	//
	// 73728
	TotalSpaceCapacity *int64 `json:"TotalSpaceCapacity,omitempty" xml:"TotalSpaceCapacity,omitempty"`
	// The used capacity of the dedicated block storage cluster. Unit: GiB.
	//
	// example:
	//
	// 1440
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
	// The capacity of the dedicated block storage cluster that was used to create disks. Unit: GB.
	//
	// example:
	//
	// 32000
	UsedDeviceCapacity *int64 `json:"UsedDeviceCapacity,omitempty" xml:"UsedDeviceCapacity,omitempty"`
	// This parameter is displayed only if Thin Provision is enabled.
	//
	// example:
	//
	// 33727.7
	UsedSpaceCapacity *float64 `json:"UsedSpaceCapacity,omitempty" xml:"UsedSpaceCapacity,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetAvailableCapacity() *int64 {
	return s.AvailableCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetAvailableDeviceCapacity() *int64 {
	return s.AvailableDeviceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetAvailableSpaceCapacity() *float64 {
	return s.AvailableSpaceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetClusterAvailableCapacity() *int64 {
	return s.ClusterAvailableCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetClusterDeliveryCapacity() *int64 {
	return s.ClusterDeliveryCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetDeliveryCapacity() *int64 {
	return s.DeliveryCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetTotalCapacity() *int64 {
	return s.TotalCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetTotalDeviceCapacity() *int64 {
	return s.TotalDeviceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetTotalSpaceCapacity() *int64 {
	return s.TotalSpaceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetUsedCapacity() *int64 {
	return s.UsedCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetUsedDeviceCapacity() *int64 {
	return s.UsedDeviceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GetUsedSpaceCapacity() *float64 {
	return s.UsedSpaceCapacity
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableSpaceCapacity(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableSpaceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetClusterAvailableCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.ClusterAvailableCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetClusterDeliveryCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.ClusterDeliveryCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetDeliveryCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.DeliveryCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalSpaceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalSpaceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedDeviceCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedDeviceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedSpaceCapacity(v float64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedSpaceCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags struct {
	// The tag key of the dedicated block storage cluster.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the dedicated block storage cluster.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) SetTagKey(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) SetTagValue(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersTags) Validate() error {
	return dara.Validate(s)
}
