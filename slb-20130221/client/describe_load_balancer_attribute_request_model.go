// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerAttributeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeLoadBalancerAttributeRequest
	GetTags() *string
	SetAccessKeyId(v string) *DescribeLoadBalancerAttributeRequest
	GetAccessKeyId() *string
}

type DescribeLoadBalancerAttributeRequest struct {
	// This parameter is required.
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s DescribeLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerAttributeRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeLoadBalancerAttributeRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetTags(v string) *DescribeLoadBalancerAttributeRequest {
	s.Tags = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetAccessKeyId(v string) *DescribeLoadBalancerAttributeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
