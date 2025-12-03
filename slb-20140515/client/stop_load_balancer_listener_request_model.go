// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *StopLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *StopLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *StopLoadBalancerListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *StopLoadBalancerListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopLoadBalancerListenerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopLoadBalancerListenerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopLoadBalancerListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopLoadBalancerListenerRequest
	GetResourceOwnerId() *int64
}

type StopLoadBalancerListenerRequest struct {
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
	// >  This parameter is required if the same port is used by listeners of different protocols.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The CLB instance ID.
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
	// To query the region ID, refer to the list of [regions and zones](https://help.aliyun.com/document_detail/40654.html) or call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *StopLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *StopLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *StopLoadBalancerListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopLoadBalancerListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopLoadBalancerListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopLoadBalancerListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopLoadBalancerListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopLoadBalancerListenerRequest) SetListenerPort(v int32) *StopLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetListenerProtocol(v string) *StopLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StopLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetOwnerAccount(v string) *StopLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetOwnerId(v int64) *StopLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetRegionId(v string) *StopLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *StopLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *StopLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
