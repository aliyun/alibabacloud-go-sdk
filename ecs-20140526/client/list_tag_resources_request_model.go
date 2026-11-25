// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *ListTagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest
	GetTag() []*ListTagResourcesRequestTag
	SetTagFilter(v []*ListTagResourcesRequestTagFilter) *ListTagResourcesRequest
	GetTagFilter() []*ListTagResourcesRequestTagFilter
}

type ListTagResourcesRequest struct {
	// The token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ECS resource. Valid values of N: 1 to 50.
	//
	// example:
	//
	// i-bp1j6qtvdm8w0z1o****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - instance: ECS instance
	//
	// - disk: cloud disk
	//
	// - snapshot: snapshot
	//
	// - image: image
	//
	// - securitygroup: security group
	//
	// - volume: storage volume
	//
	// - eni: Elastic Network Interface (ENI)
	//
	// - ddh: dedicated host
	//
	// - ddhcluster: dedicated host cluster
	//
	// - keypair: SSH key pair
	//
	// - launchtemplate: launch template
	//
	// - reservedinstance: reserved instance
	//
	// - snapshotpolicy: automatic snapshot policy
	//
	// - elasticityassurance: elasticity assurance
	//
	// - capacityreservation: capacity reservation
	//
	// - command: Cloud Assistant command
	//
	// - invocation: Cloud Assistant command execution or file delivery result
	//
	// - activation: Cloud Assistant managed instance activation code
	//
	// - managedinstance: Cloud Assistant managed instance
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The tag filtering rules.
	//
	//
	// > This parameter is in invitational preview and is not publicly available.
	TagFilter []*ListTagResourcesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTag() []*ListTagResourcesRequestTag {
	return s.Tag
}

func (s *ListTagResourcesRequest) GetTagFilter() []*ListTagResourcesRequestTagFilter {
	return s.TagFilter
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetTagFilter(v []*ListTagResourcesRequestTagFilter) *ListTagResourcesRequest {
	s.TagFilter = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagFilter != nil {
		for _, item := range s.TagFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesRequestTag struct {
	// The tag key used for exact match when querying ECS resources. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// Usage notes for the Tag.N parameter:
	//
	// - Method 1: Used to exactly query ECS resources bound to specific tags. It consists of a key-value pair.
	//
	//     - If only Tag.N.Key is specified, all resources associated with the tag key are returned.
	//
	//     - If only Tag.N.Value is specified, the InvalidParameter.TagValue error is returned.
	//
	//     - If multiple tag key-value pairs are specified, only ECS resources that meet all the tag key-value pairs are returned.
	//
	// - Method 2: Used to query resources in non-default resource groups. The Key is fixed as acs:rm:rgId, and the corresponding Value is the resource group ID.
	//
	//     - If Key is set to acs:rm:rgId, Value can only be a non-default resource group ID. If the specified resource group ID is the default resource group, an error message is returned.
	//
	//     - If Key is set to acs:rm:rgId, you cannot specify other tag key-value pairs. If you query resources corresponding to both resource groups and tags by using multiple Tag.N parameters at the same time, an error message is returned.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value used for exact match when querying ECS resources. The tag value must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// > When Key is set to acs:rm:rgId, this parameter can only be set to a resource group ID, and the resource group ID cannot be the default resource group.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesRequestTagFilter struct {
	// The tag key used for fuzzy match when querying ECS resources. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 5.
	//
	// TagFilter.N is used to fuzzy query ECS resources bound to specific tags. It consists of a key and one or more values. Fuzzy queries may have a delay of 2 seconds and only support scenarios where the number of resources after fuzzy filtering is less than or equal to 5,000.
	//
	// - When you fuzzy query ECS resources by tag key (TagFilter.N.TagKey), the tag value (TagFilter.N.TagValues.N) must be empty. For example, to fuzzy search for ECS resources with the tag key environment, you can set TagFilter.1.TagKey to env	- (prefix match), *env	- (infix match), or env (exact match), and TagFilter.1.TagValues must be empty.
	//
	// - When you fuzzy query ECS resources by tag value (TagFilter.N.TagValues.N), the tag key (TagFilter.N.TagKey) must be set to an exact value. For example, to fuzzy search for ECS resources with the tag key env and the tag value product, you must set TagFilter.1.TagKey exactly to env, and you can set TagFilter.1.TagValues.1 to proc	- (prefix match), *proc	- (infix match), or proc (exact match). You can use only one search method under the same TagKey. If multiple different search methods are set, the first method prevails.
	//
	// - Tag keys have an intersection relationship. Only ECS resources that meet all the specified tag keys are returned.
	//
	// - Tag values under the same tag key have a union relationship. ECS resources that meet any of the specified tag values for the tag key are returned.
	//
	// > You cannot use TagFilter.N and Tag.N at the same time. Otherwise, an error message is returned.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value used for fuzzy match when querying ECS resources. The tag value must be 1 to 128 characters in length. Valid values of N: 1 to 5. For the specific metric description, see the description of the TagFilter.N.TagKey parameter.
	//
	// example:
	//
	// TestTagFilter
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTagFilter) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesRequestTagFilter) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListTagResourcesRequestTagFilter) SetTagKey(v string) *ListTagResourcesRequestTagFilter {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesRequestTagFilter) SetTagValues(v []*string) *ListTagResourcesRequestTagFilter {
	s.TagValues = v
	return s
}

func (s *ListTagResourcesRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
