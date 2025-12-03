// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeHealthStatusRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DescribeHealthStatusRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *DescribeHealthStatusRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeHealthStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHealthStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHealthStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHealthStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHealthStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeHealthStatusRequest struct {
	// The frontend port that is used by the SLB instance. Valid values: **1 to 65535**.
	//
	// >  If you do not specify this parameter, the health status of all ports is returned.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend protocol that is used by the SLB instance.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the Classic Load Balancer (CLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1qjwo61pqz3ah****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeHealthStatusRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeHealthStatusRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeHealthStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHealthStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHealthStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHealthStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHealthStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHealthStatusRequest) SetListenerPort(v int32) *DescribeHealthStatusRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetListenerProtocol(v string) *DescribeHealthStatusRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetLoadBalancerId(v string) *DescribeHealthStatusRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetOwnerAccount(v string) *DescribeHealthStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetOwnerId(v int64) *DescribeHealthStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetRegionId(v string) *DescribeHealthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetResourceOwnerAccount(v string) *DescribeHealthStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetResourceOwnerId(v int64) *DescribeHealthStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
