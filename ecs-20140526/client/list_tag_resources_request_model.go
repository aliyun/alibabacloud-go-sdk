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
	// The pagination token to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the resource is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of an ECS resource. The value of N ranges from 1 to 50.
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
	// - disk: disk
	//
	// - snapshot: snapshot
	//
	// - image: image
	//
	// - securitygroup: security group
	//
	// - volume: volume
	//
	// - eni: elastic network interface
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
	// - snapshotpolicy: snapshot policy
	//
	// - elasticityassurance: Elasticity Assurance
	//
	// - capacityreservation: capacity reservation
	//
	// - command: Cloud Assistant command
	//
	// - invocation: The result of a command execution or file delivery in Cloud Assistant
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
	// A list of tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// A list of tag filters.
	//
	// > This parameter is in invitation-only preview and is not yet available.
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
	// The tag key to use for an exact match. The tag key must be 1 to 128 characters in length. The value of N ranges from 1 to 20.
	//
	// Usage notes for the `Tag.N` parameter:
	//
	// - Method 1: To find ECS resources that have specific tags.
	//
	//   - If you specify only `Tag.N.Key`, the operation returns all resources that have the specified tag key.
	//
	//   - If you specify only `Tag.N.Value`, the operation returns an `InvalidParameter.TagValue` error.
	//
	//   - If you specify multiple tag key-value pairs, the operation returns only the ECS resources that match all specified pairs.
	//
	// - Method 2: To query resources in a non-default resource group.
	//
	//   - If you set `Key` to `acs:rm:rgId`, you must set `Value` to the ID of a non-default resource group. If you specify the ID of the default resource group, the operation returns an error.
	//
	//   - If you set `Key` to `acs:rm:rgId`, you cannot specify other tag key-value pairs. If you use multiple `Tag.N` parameters to query for resources by both resource group and tag, the operation returns an error.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to use for an exact match. The tag value must be 1 to 128 characters in length. The value of N ranges from 1 to 20.
	//
	// > When `Key` is `acs:rm:rgId`, you must set this parameter to the ID of a non-default resource group.
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
	// The tag key to use for a fuzzy match. The tag key must be 1 to 128 characters in length. The value of N ranges from 1 to 5.
	//
	// Use the `TagFilter.N` parameter to perform a fuzzy match on tags to find matching ECS resources. Each filter consists of one key and one or more values. A fuzzy match may have a 2-second latency and is supported only for queries that return 5,000 or fewer resources after filtering.
	//
	// - To perform a fuzzy match by tag key (`TagFilter.N.TagKey`), you must leave the tag values (`TagFilter.N.TagValues.N`) empty. For example, to search for ECS resources that have the tag key `environment`, you can set `TagFilter.1.TagKey` to `env*` (prefix match), `*env*` (substring match), or `env` (exact match), but you must leave `TagFilter.1.TagValues` empty.
	//
	// - To perform a fuzzy match by tag value (`TagFilter.N.TagValues.N`), you must set the tag key (`TagFilter.N.TagKey`) to an exact value. For example, to search for ECS resources with the tag key `env` and the tag value `product`, you must set `TagFilter.1.TagKey` to `env`. You can then set `TagFilter.1.TagValues.1` to `proc*` (prefix match), `*proc*` (substring match), or `proc` (exact match). For the same `TagKey`, you can use only one search pattern. If you specify multiple patterns, the system uses only the first pattern.
	//
	// - Tag keys are combined by using a logical AND. The operation returns only the ECS resources that match all specified tag keys.
	//
	// - Tag values for the same tag key are combined by using a logical OR. The operation returns the ECS resources that match any of the specified tag values for that tag key.
	//
	// > You cannot specify both the `TagFilter.N` and `Tag.N` parameters in the same request.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag values to use for a fuzzy match. The tag value must be 1 to 128 characters in length. The value of N ranges from 1 to 5. For more information, see the description of the `TagFilter.N.TagKey` parameter.
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
