// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *AddTagsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *AddTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTagsRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *AddTagsRequest
	GetTags() *string
}

type AddTagsRequest struct {
	// The ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1kuzyb******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SLB instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags that need to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagKey":"Key1","TagValue":"Value1"},{"TagKey":"Key2","TagValue":"Value2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *AddTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *AddTagsRequest) SetLoadBalancerId(v string) *AddTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *AddTagsRequest) SetOwnerAccount(v string) *AddTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetOwnerId(v int64) *AddTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsRequest) SetRegionId(v string) *AddTagsRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerAccount(v string) *AddTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerId(v int64) *AddTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsRequest) SetTags(v string) *AddTagsRequest {
	s.Tags = &v
	return s
}

func (s *AddTagsRequest) Validate() error {
	return dara.Validate(s)
}
