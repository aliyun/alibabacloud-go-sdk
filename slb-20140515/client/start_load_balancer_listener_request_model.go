// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoadBalancerListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *StartLoadBalancerListenerRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *StartLoadBalancerListenerRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *StartLoadBalancerListenerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *StartLoadBalancerListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartLoadBalancerListenerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartLoadBalancerListenerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartLoadBalancerListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartLoadBalancerListenerRequest
	GetResourceOwnerId() *int64
}

type StartLoadBalancerListenerRequest struct {
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
	// lb-bp13jaf5qli5*********
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the CLB instance is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartLoadBalancerListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLoadBalancerListenerRequest) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *StartLoadBalancerListenerRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *StartLoadBalancerListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *StartLoadBalancerListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartLoadBalancerListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartLoadBalancerListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartLoadBalancerListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartLoadBalancerListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartLoadBalancerListenerRequest) SetListenerPort(v int32) *StartLoadBalancerListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetListenerProtocol(v string) *StartLoadBalancerListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetLoadBalancerId(v string) *StartLoadBalancerListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetOwnerAccount(v string) *StartLoadBalancerListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetOwnerId(v int64) *StartLoadBalancerListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetRegionId(v string) *StartLoadBalancerListenerRequest {
	s.RegionId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetResourceOwnerAccount(v string) *StartLoadBalancerListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) SetResourceOwnerId(v int64) *StartLoadBalancerListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartLoadBalancerListenerRequest) Validate() error {
	return dara.Validate(s)
}
