// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *RemoveTagsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *RemoveTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveTagsRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *RemoveTagsRequest
	GetTags() *string
}

type RemoveTagsRequest struct {
	// The SLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1l5j******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A list of tags to be removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagKey":"Key1","TagValue":"Value1"},{"TagKey":"Key2","TagValue":"Value2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s RemoveTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *RemoveTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *RemoveTagsRequest) SetLoadBalancerId(v string) *RemoveTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *RemoveTagsRequest) SetOwnerAccount(v string) *RemoveTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTagsRequest) SetOwnerId(v int64) *RemoveTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetRegionId(v string) *RemoveTagsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerAccount(v string) *RemoveTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTagsRequest) SetResourceOwnerId(v int64) *RemoveTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTagsRequest) SetTags(v string) *RemoveTagsRequest {
	s.Tags = &v
	return s
}

func (s *RemoveTagsRequest) Validate() error {
	return dara.Validate(s)
}
