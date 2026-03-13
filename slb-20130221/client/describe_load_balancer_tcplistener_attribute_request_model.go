// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetTags() *string
	SetAccessKeyId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetAccessKeyId() *string
}

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	// This parameter is required.
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// This parameter is required.
	LoadBalancerId       *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	AccessKeyId          *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetTags(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.Tags = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetAccessKeyId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
