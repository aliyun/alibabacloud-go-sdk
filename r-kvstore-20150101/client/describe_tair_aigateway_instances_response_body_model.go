// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairAIGatewayInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeTairAIGatewayInstancesResponseBodyInstances) *DescribeTairAIGatewayInstancesResponseBody
	GetInstances() *DescribeTairAIGatewayInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeTairAIGatewayInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairAIGatewayInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTairAIGatewayInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTairAIGatewayInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeTairAIGatewayInstancesResponseBody struct {
	Instances *DescribeTairAIGatewayInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instance records for batch instance creation. This parameter is not returned when a single instance is created.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTairAIGatewayInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponseBody) GetInstances() *DescribeTairAIGatewayInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeTairAIGatewayInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairAIGatewayInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairAIGatewayInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairAIGatewayInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTairAIGatewayInstancesResponseBody) SetInstances(v *DescribeTairAIGatewayInstancesResponseBodyInstances) *DescribeTairAIGatewayInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBody) SetPageNumber(v int32) *DescribeTairAIGatewayInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBody) SetPageSize(v int32) *DescribeTairAIGatewayInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBody) SetRequestId(v string) *DescribeTairAIGatewayInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBody) SetTotalCount(v int32) *DescribeTairAIGatewayInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairAIGatewayInstancesResponseBodyInstances struct {
	TairSemanticCacheInstances []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances `json:"TairSemanticCacheInstances,omitempty" xml:"TairSemanticCacheInstances,omitempty" type:"Repeated"`
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstances) GetTairSemanticCacheInstances() []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	return s.TairSemanticCacheInstances
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstances) SetTairSemanticCacheInstances(v []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) *DescribeTairAIGatewayInstancesResponseBodyInstances {
	s.TairSemanticCacheInstances = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstances) Validate() error {
	if s.TairSemanticCacheInstances != nil {
		for _, item := range s.TairSemanticCacheInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances struct {
	ChargeType      *string                                                                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime      *string                                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime         *string                                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceClass   *string                                                                            `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId      *string                                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                                                            `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus  *string                                                                            `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType    *string                                                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NetworkType     *string                                                                            `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RegionId        *string                                                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                                                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId       *string                                                                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string                                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string                                                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetTags() *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags {
	return s.Tags
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetChargeType(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetCreateTime(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetEndTime(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.EndTime = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetInstanceClass(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetInstanceId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetInstanceName(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetInstanceStatus(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetInstanceType(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetNetworkType(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetRegionId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetResourceGroupId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetTags(v *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.Tags = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetVSwitchId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetVpcId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) SetZoneId(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstances) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags struct {
	Tag []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) GetTag() []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag {
	return s.Tag
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) SetTag(v []*DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags {
	s.Tag = v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTags) Validate() error {
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

type DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) SetKey(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) SetValue(v string) *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTairAIGatewayInstancesResponseBodyInstancesTairSemanticCacheInstancesTagsTag) Validate() error {
	return dara.Validate(s)
}
