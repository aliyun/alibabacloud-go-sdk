// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeVServerGroupsRequest
	GetDescription() *string
	SetIncludeListener(v bool) *DescribeVServerGroupsRequest
	GetIncludeListener() *bool
	SetIncludeRule(v bool) *DescribeVServerGroupsRequest
	GetIncludeRule() *bool
	SetLoadBalancerId(v string) *DescribeVServerGroupsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeVServerGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVServerGroupsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeVServerGroupsRequestTag) *DescribeVServerGroupsRequest
	GetTag() []*DescribeVServerGroupsRequestTag
}

type DescribeVServerGroupsRequest struct {
	// The name of the vServer group.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Group3
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to return information about the associated listeners. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	IncludeListener *bool `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	// Specifies whether to return the forwarding rules associated with the vServer groups. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	IncludeRule *bool `json:"IncludeRule,omitempty" xml:"IncludeRule,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea*******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeVServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeVServerGroupsRequest) GetIncludeListener() *bool {
	return s.IncludeListener
}

func (s *DescribeVServerGroupsRequest) GetIncludeRule() *bool {
	return s.IncludeRule
}

func (s *DescribeVServerGroupsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeVServerGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVServerGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVServerGroupsRequest) GetTag() []*DescribeVServerGroupsRequestTag {
	return s.Tag
}

func (s *DescribeVServerGroupsRequest) SetDescription(v string) *DescribeVServerGroupsRequest {
	s.Description = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetIncludeListener(v bool) *DescribeVServerGroupsRequest {
	s.IncludeListener = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetIncludeRule(v bool) *DescribeVServerGroupsRequest {
	s.IncludeRule = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetLoadBalancerId(v string) *DescribeVServerGroupsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetOwnerAccount(v string) *DescribeVServerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetOwnerId(v int64) *DescribeVServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetRegionId(v string) *DescribeVServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeVServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetResourceOwnerId(v int64) *DescribeVServerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVServerGroupsRequest) SetTag(v []*DescribeVServerGroupsRequestTag) *DescribeVServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVServerGroupsRequest) Validate() error {
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

type DescribeVServerGroupsRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs`:.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVServerGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVServerGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVServerGroupsRequestTag) SetKey(v string) *DescribeVServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVServerGroupsRequestTag) SetValue(v string) *DescribeVServerGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVServerGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
