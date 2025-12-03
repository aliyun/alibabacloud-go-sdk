// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DeleteLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DeleteLoadBalancerListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLoadBalancerListenerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLoadBalancerListenerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLoadBalancerListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLoadBalancerListenerRequest
	GetResourceOwnerId() *int64
}

type DeleteLoadBalancerListenerRequest struct {
	// The frontend port that is used by the CLB instance.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the CLB instance.
	//
	// >  This parameter is required if the same port is specified for listeners of different protocols.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp13jaf5qli5xmg******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// >  The **RegionId*	- parameter is required if the endpoint of the region is slb.aliyuncs.com.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DeleteLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DeleteLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteLoadBalancerListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLoadBalancerListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLoadBalancerListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLoadBalancerListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLoadBalancerListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerPort(v int32) *DeleteLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetListenerProtocol(v string) *DeleteLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetRegionId(v string) *DeleteLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *DeleteLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
