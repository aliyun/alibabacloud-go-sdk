// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDeploymentSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSets(v *DescribeRCDeploymentSetsResponseBodyDeploymentSets) *DescribeRCDeploymentSetsResponseBody
	GetDeploymentSets() *DescribeRCDeploymentSetsResponseBodyDeploymentSets
	SetPageNumber(v int32) *DescribeRCDeploymentSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCDeploymentSetsResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCDeploymentSetsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeRCDeploymentSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCDeploymentSetsResponseBody
	GetTotalCount() *int32
}

type DescribeRCDeploymentSetsResponseBody struct {
	// The details of the deployment set.
	DeploymentSets *DescribeRCDeploymentSetsResponseBodyDeploymentSets `json:"DeploymentSets,omitempty" xml:"DeploymentSets,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// 39265F46-EC77-4036-8AC4-F035F32F6BE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCDeploymentSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBody) GetDeploymentSets() *DescribeRCDeploymentSetsResponseBodyDeploymentSets {
	return s.DeploymentSets
}

func (s *DescribeRCDeploymentSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCDeploymentSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCDeploymentSetsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCDeploymentSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCDeploymentSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCDeploymentSetsResponseBody) SetDeploymentSets(v *DescribeRCDeploymentSetsResponseBodyDeploymentSets) *DescribeRCDeploymentSetsResponseBody {
	s.DeploymentSets = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) SetPageNumber(v int32) *DescribeRCDeploymentSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) SetPageSize(v int32) *DescribeRCDeploymentSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) SetRegionId(v string) *DescribeRCDeploymentSetsResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) SetRequestId(v string) *DescribeRCDeploymentSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) SetTotalCount(v int32) *DescribeRCDeploymentSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBody) Validate() error {
	if s.DeploymentSets != nil {
		if err := s.DeploymentSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSets struct {
	DeploymentSet []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet `json:"DeploymentSet,omitempty" xml:"DeploymentSet,omitempty" type:"Repeated"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSets) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSets) GetDeploymentSet() []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	return s.DeploymentSet
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSets) SetDeploymentSet(v []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) *DescribeRCDeploymentSetsResponseBodyDeploymentSets {
	s.DeploymentSet = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSets) Validate() error {
	if s.DeploymentSet != nil {
		for _, item := range s.DeploymentSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet struct {
	// The details of the capacities of the deployment set. This parameter is valid only when the deployment set contains existing RDS Custom instances. The value contains the details of the capacities of the deployment set in different zones.
	Capacities *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities `json:"Capacities,omitempty" xml:"Capacities,omitempty" type:"Struct"`
	// The time when the deployment set was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-06-19T07:15:44Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deployment set description.
	//
	// example:
	//
	// test
	DeploymentSetDescription *string `json:"DeploymentSetDescription,omitempty" xml:"DeploymentSetDescription,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-ob5n4rbgy****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The deployment set name.
	//
	// example:
	//
	// deployment_test
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// The deployment strategy. The return value of this parameter is the value of the `Strategy` request parameter.
	//
	// example:
	//
	// Availability
	DeploymentStrategy *string `json:"DeploymentStrategy,omitempty" xml:"DeploymentStrategy,omitempty"`
	// The deployment domain.
	//
	// example:
	//
	// default
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The deployment granularity.
	//
	// example:
	//
	// None
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of groups in the deployment set.
	//
	// >  This parameter is valid only when the Strategy request parameter is set to AvailabilityGroup.
	//
	// example:
	//
	// 3
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The number of RDS Custom instances in the deployment set.
	//
	// example:
	//
	// 1
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The ID of the RDS Custom instance in the deployment set.
	InstanceIds *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The deployment strategy.
	//
	// example:
	//
	// LooseDispersion
	Strategy *string                                                              `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	Tags     *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetCapacities() *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities {
	return s.Capacities
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetDescription() *string {
	return s.DeploymentSetDescription
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDeploymentStrategy() *string {
	return s.DeploymentStrategy
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetDomain() *string {
	return s.Domain
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetInstanceIds() *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds {
	return s.InstanceIds
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) GetTags() *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags {
	return s.Tags
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetCapacities(v *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Capacities = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetCreateTime(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.CreateTime = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetDescription(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetDescription = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetId(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentSetName(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentSetName = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDeploymentStrategy(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.DeploymentStrategy = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetDomain(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Domain = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetGranularity(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Granularity = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetGroupCount(v int32) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.GroupCount = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetInstanceAmount(v int32) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.InstanceAmount = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetInstanceIds(v *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.InstanceIds = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetStrategy(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Strategy = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) SetTags(v *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet {
	s.Tags = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSet) Validate() error {
	if s.Capacities != nil {
		if err := s.Capacities.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceIds != nil {
		if err := s.InstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities struct {
	Capacity []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity `json:"Capacity,omitempty" xml:"Capacity,omitempty" type:"Repeated"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) GetCapacity() []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	return s.Capacity
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) SetCapacity(v []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities {
	s.Capacity = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacities) Validate() error {
	if s.Capacity != nil {
		for _, item := range s.Capacity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity struct {
	// The number of RDS Custom instances that reside in the zone and can be added to the deployment set.
	//
	// example:
	//
	// 18
	AvailableAmount *int32 `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// The number of RDS Custom instances that reside in the zone in the deployment set.
	//
	// example:
	//
	// 2
	UsedAmount *int32 `json:"UsedAmount,omitempty" xml:"UsedAmount,omitempty"`
	// The zone ID. Only the IDs of the zones to which the existing RDS Custom instances in the deployment set belong are returned.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetAvailableAmount() *int32 {
	return s.AvailableAmount
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetUsedAmount() *int32 {
	return s.UsedAmount
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetAvailableAmount(v int32) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.AvailableAmount = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetUsedAmount(v int32) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.UsedAmount = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) SetZoneId(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetCapacitiesCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) SetInstanceId(v []*string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags struct {
	Tag []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) GetTag() []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag {
	return s.Tag
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) SetTag(v []*DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags {
	s.Tag = v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTags) Validate() error {
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

type DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) SetResourceId(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag {
	s.ResourceId = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) SetResourceType(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) SetTagKey(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) SetTagValue(v string) *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeRCDeploymentSetsResponseBodyDeploymentSetsDeploymentSetTagsTag) Validate() error {
	return dara.Validate(s)
}
