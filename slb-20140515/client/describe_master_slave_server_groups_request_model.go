// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeMasterSlaveServerGroupsRequest
	GetDescription() *string
	SetIncludeListener(v bool) *DescribeMasterSlaveServerGroupsRequest
	GetIncludeListener() *bool
	SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeMasterSlaveServerGroupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeMasterSlaveServerGroupsRequestTag) *DescribeMasterSlaveServerGroupsRequest
	GetTag() []*DescribeMasterSlaveServerGroupsRequestTag
}

type DescribeMasterSlaveServerGroupsRequest struct {
	// The description of the primary/secondary server group.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test-112
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to return information about the associated listeners. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IncludeListener *bool `json:"IncludeListener,omitempty" xml:"IncludeListener,omitempty"`
	// The CLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp14zi0n66zpg6o******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
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
	Tag []*DescribeMasterSlaveServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMasterSlaveServerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetIncludeListener() *bool {
	return s.IncludeListener
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMasterSlaveServerGroupsRequest) GetTag() []*DescribeMasterSlaveServerGroupsRequestTag {
	return s.Tag
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetDescription(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.Description = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetIncludeListener(v bool) *DescribeMasterSlaveServerGroupsRequest {
	s.IncludeListener = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetLoadBalancerId(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetRegionId(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) SetTag(v []*DescribeMasterSlaveServerGroupsRequestTag) *DescribeMasterSlaveServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequest) Validate() error {
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

type DescribeMasterSlaveServerGroupsRequestTag struct {
	// The key of tag N. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` and `acs:`.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMasterSlaveServerGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMasterSlaveServerGroupsRequestTag) SetKey(v string) *DescribeMasterSlaveServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequestTag) SetValue(v string) *DescribeMasterSlaveServerGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
