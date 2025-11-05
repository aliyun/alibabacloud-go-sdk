// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedBlockStorageClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzone(v string) *CreateDedicatedBlockStorageClusterRequest
	GetAzone() *string
	SetCapacity(v int64) *CreateDedicatedBlockStorageClusterRequest
	GetCapacity() *int64
	SetDbscId(v string) *CreateDedicatedBlockStorageClusterRequest
	GetDbscId() *string
	SetDbscName(v string) *CreateDedicatedBlockStorageClusterRequest
	GetDbscName() *string
	SetPeriod(v int32) *CreateDedicatedBlockStorageClusterRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDedicatedBlockStorageClusterRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *CreateDedicatedBlockStorageClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDedicatedBlockStorageClusterRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateDedicatedBlockStorageClusterRequestTag) *CreateDedicatedBlockStorageClusterRequest
	GetTag() []*CreateDedicatedBlockStorageClusterRequestTag
	SetType(v string) *CreateDedicatedBlockStorageClusterRequest
	GetType() *string
}

type CreateDedicatedBlockStorageClusterRequest struct {
	// The ID of the zone in which to create the dedicated block storage cluster. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan-b
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// The capacity of the dedicated block storage cluster. Valid values: 61440 to 2334720. Unit: GiB. 2,334,720 GiB is equal to 2,280 TiB. The capacity increases in a minimum increment of 12,288 GiB.
	//
	// >  If the capacity of a dedicated block storage cluster is less than 576 TiB, the maximum throughput per TiB cannot exceed 52 MB/s. If the capacity of a dedicated block storage cluster is greater than 576 TiB, the maximum throughput per TiB cannot exceed 26 MB/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61440
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test1233
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The name of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// myDBSCCluster
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The subscription duration of the dedicated block storage cluster. Valid values: 6, 7, 8, 9, 10, 11, 12, 24, and 36.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration specified by `Period`. Set the value to Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the region in which to create the dedicated block storage cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the dedicated block storage cluster.
	//
	// example:
	//
	// rg-acfmvs*******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags to add to the dedicated block storage cluster. You can specify up to 20 tags.
	Tag []*CreateDedicatedBlockStorageClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the dedicated block storage cluster. Valid values:
	//
	// 	- Standard: basic dedicated block storage cluster. Enterprise SSDs (ESSDs) at performance level 0 (PL0 ESSDs) can be created in basic dedicated block storage clusters.
	//
	// 	- Premium: performance dedicated block storage cluster. ESSDs at performance level 1 (PL1 ESSDs) can be created in performance dedicated block storage clusters.
	//
	// Default value: Premium.
	//
	// For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Premium
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetAzone() *string {
	return s.Azone
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetDbscId() *string {
	return s.DbscId
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetDbscName() *string {
	return s.DbscName
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetTag() []*CreateDedicatedBlockStorageClusterRequestTag {
	return s.Tag
}

func (s *CreateDedicatedBlockStorageClusterRequest) GetType() *string {
	return s.Type
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetAzone(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Azone = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetCapacity(v int64) *CreateDedicatedBlockStorageClusterRequest {
	s.Capacity = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscName(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscName = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetPeriod(v int32) *CreateDedicatedBlockStorageClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetPeriodUnit(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetRegionId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetResourceGroupId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetTag(v []*CreateDedicatedBlockStorageClusterRequestTag) *CreateDedicatedBlockStorageClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetType(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Type = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDedicatedBlockStorageClusterRequestTag struct {
	// The key of tag N to add to the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) SetKey(v string) *CreateDedicatedBlockStorageClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) SetValue(v string) *CreateDedicatedBlockStorageClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequestTag) Validate() error {
	return dara.Validate(s)
}
