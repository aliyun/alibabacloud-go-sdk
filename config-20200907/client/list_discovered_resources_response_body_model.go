// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiscoveredResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceProfiles(v *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListDiscoveredResourcesResponseBody
	GetDiscoveredResourceProfiles() *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles
	SetRequestId(v string) *ListDiscoveredResourcesResponseBody
	GetRequestId() *string
}

type ListDiscoveredResourcesResponseBody struct {
	// The information about the resources.
	DiscoveredResourceProfiles *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C7817373-78CB-4F9A-8AFA-E7A88E9D64A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDiscoveredResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBody) GetDiscoveredResourceProfiles() *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	return s.DiscoveredResourceProfiles
}

func (s *ListDiscoveredResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiscoveredResourcesResponseBody) SetDiscoveredResourceProfiles(v *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListDiscoveredResourcesResponseBody {
	s.DiscoveredResourceProfiles = v
	return s
}

func (s *ListDiscoveredResourcesResponseBody) SetRequestId(v string) *ListDiscoveredResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBody) Validate() error {
	if s.DiscoveredResourceProfiles != nil {
		if err := s.DiscoveredResourceProfiles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles struct {
	// The details of the resources.
	DiscoveredResourceProfileList []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	// The maximum number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 161259599160****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetDiscoveredResourceProfileList() []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	return s.DiscoveredResourceProfileList
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetDiscoveredResourceProfileList(v []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.DiscoveredResourceProfileList = v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetMaxResults(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.MaxResults = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetNextToken(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.NextToken = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) Validate() error {
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

type ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 161259599160****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created. Unit: milliseconds.
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
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// rg-acfmvoh45rhxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// The status of the resource. The value of this parameter varies based on the resource type and may be empty. Examples:
	//
	// 	- If the ResourceType parameter is set to ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	//
	// 	- If the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that is not in a specific state. In this case, this parameter is left empty.
	//
	// example:
	//
	// InUse
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
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
	// The version of the resource change.
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

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) String() string {
	return dara.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetRegion() *string {
	return s.Region
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceCreationTime() *int64 {
	return s.ResourceCreationTime
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceDeleted() *int32 {
	return s.ResourceDeleted
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetTags() *string {
	return s.Tags
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVersion() *int64 {
	return s.Version
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAccountId(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AccountId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAvailabilityZone(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AvailabilityZone = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetRegion(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Region = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceCreationTime(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceCreationTime = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceDeleted(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceGroupId(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceId(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceName(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceName = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceStatus(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceStatus = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceType(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceType = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetTags(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Tags = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetUpdateTime(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.UpdateTime = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVSwitchId(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.VSwitchId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVersion(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Version = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetVpcId(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.VpcId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) Validate() error {
	return dara.Validate(s)
}
