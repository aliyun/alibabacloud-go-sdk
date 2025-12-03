// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistinctKey(v bool) *DescribeTagsRequest
	GetDistinctKey() *bool
	SetLoadBalancerId(v string) *DescribeTagsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagsRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeTagsRequest
	GetTags() *string
}

type DescribeTagsRequest struct {
	// Specifies whether the tags contain distinct keys.
	//
	// Valid values: true and false.
	//
	// example:
	//
	// false
	DistinctKey *bool `json:"DistinctKey,omitempty" xml:"DistinctKey,omitempty"`
	// The SLB instance ID.
	//
	// example:
	//
	// lb-bp1kuzybm******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 50. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that you want to query.
	//
	// example:
	//
	// [{"TagKey":"Key1","TagValue":"Value1"},{"TagKey":"Key2","TagValue":"Value2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetDistinctKey() *bool {
	return s.DistinctKey
}

func (s *DescribeTagsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeTagsRequest) SetDistinctKey(v bool) *DescribeTagsRequest {
	s.DistinctKey = &v
	return s
}

func (s *DescribeTagsRequest) SetLoadBalancerId(v string) *DescribeTagsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetTags(v string) *DescribeTagsRequest {
	s.Tags = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
