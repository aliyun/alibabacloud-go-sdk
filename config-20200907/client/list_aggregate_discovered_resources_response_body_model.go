// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateDiscoveredResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceProfiles(v *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListAggregateDiscoveredResourcesResponseBody
	GetDiscoveredResourceProfiles() *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles
	SetRequestId(v string) *ListAggregateDiscoveredResourcesResponseBody
	GetRequestId() *string
}

type ListAggregateDiscoveredResourcesResponseBody struct {
	// The information about the resources.
	DiscoveredResourceProfiles *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBody) GetDiscoveredResourceProfiles() *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	return s.DiscoveredResourceProfiles
}

func (s *ListAggregateDiscoveredResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateDiscoveredResourcesResponseBody) SetDiscoveredResourceProfiles(v *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListAggregateDiscoveredResourcesResponseBody {
	s.DiscoveredResourceProfiles = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBody) SetRequestId(v string) *ListAggregateDiscoveredResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBody) Validate() error {
	if s.DiscoveredResourceProfiles != nil {
		if err := s.DiscoveredResourceProfiles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles struct {
	// The details of the resources.
	DiscoveredResourceProfileList []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	// The maximum number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that was used to initiate the next request.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetDiscoveredResourceProfileList() []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	return s.DiscoveredResourceProfileList
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetDiscoveredResourceProfileList(v []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.DiscoveredResourceProfileList = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetMaxResults(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetNextToken(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.NextToken = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) Validate() error {
	if s.DiscoveredResourceProfileList != nil {
		for _, item := range s.DiscoveredResourceProfileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs. We recommend that you use the ResourceOwnerId parameter.
	//
	// example:
	//
	// 161259599160****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone in which the resource resides.
	//
	// example:
	//
	// cn-huhehaote-a
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-huhehaote
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1618675206000
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// The status of the resource. Valid values:
	//
	// 	- 0: The resource is deleted.
	//
	// 	- 1: The resource is retained.
	//
	// example:
	//
	// 1
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eni-hp31cqoba96jagtz****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// Cloud Firewall
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 161259599160****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the resource. The value of this parameter varies with the resource type and may be empty. Examples:
	//
	// 	- If the value of the ResourceType parameter is ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that is in a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// 	- If the value of the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that is not in a specific state. In this case, this parameter is empty.
	//
	// example:
	//
	// InUse
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	//
	// example:
	//
	// {\\"key1\\":[\\"value2\\"]}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the resource was last updated. The value must be a timestamp in milliseconds.
	//
	// example:
	//
	// 1722441600000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-t4n7pokxxxxxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The build version of the resource.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-t4nhheyvay74fp7n0hxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetRegion() *string {
	return s.Region
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetTags() *string {
	return s.Tags
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVersion() *int64 {
	return s.Version
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAccountId(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AccountId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAvailabilityZone(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AvailabilityZone = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetRegion(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Region = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceCreationTime(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceCreationTime = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceDeleted(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceDeleted = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceId(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceName(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceOwnerId(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceStatus(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceStatus = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceType(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceType = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetTags(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Tags = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetUpdateTime(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.UpdateTime = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVSwitchId(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.VSwitchId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVersion(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Version = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVpcId(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.VpcId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) Validate() error {
	return dara.Validate(s)
}
