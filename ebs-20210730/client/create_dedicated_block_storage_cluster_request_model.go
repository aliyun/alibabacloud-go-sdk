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
	// The zone ID of the dedicated block storage cluster. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan-b
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// The capacity of the dedicated block storage cluster. Valid values: 61440 to 2334720 GiB (2280 TiB). Minimum increment: 12288 GiB.
	//
	// > When the capacity of the dedicated block storage cluster is less than 576 TiB, the maximum throughput per TiB does not exceed 52 MB/s. When the capacity of the dedicated block storage cluster is greater than 576 TiB, the maximum throughput per TiB does not exceed 26 MB/s.
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
	// The subscription duration of the instance. Valid values: 6, 7, 8, 9, 10, 11, 12, 24, and 36.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration specified by the `Period` parameter. Only Month is supported.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the dedicated block storage cluster. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated block storage cluster belongs.
	//
	// example:
	//
	// rg-acfmvs*******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags. A maximum of 20 tags can be specified.
	Tag []*CreateDedicatedBlockStorageClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The performance type of the dedicated block storage cluster. Valid values:
	//
	// - Standard: basic. You can create PL0 ESSDs in this type of dedicated block storage cluster.
	//
	// - Premium: performance. You can create PL1 ESSDs in this type of dedicated block storage cluster.
	//
	// Default value: Premium.
	//
	// For more information about standard SSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
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
	// The tag key of the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated block storage cluster.
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
