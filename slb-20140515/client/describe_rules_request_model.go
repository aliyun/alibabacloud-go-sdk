// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeRulesRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DescribeRulesRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *DescribeRulesRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRulesRequest
	GetResourceOwnerId() *int64
}

type DescribeRulesRequest struct {
	// The frontend listener port that is used by the Server Load Balancer (SLB) instance.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The frontend listener protocol that is used by the SLB instance.
	//
	// >  This parameter is required when listeners that use different protocols listen on the same port.
	//
	// example:
	//
	// https
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1ca0zt07t934****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeRulesRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeRulesRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRulesRequest) SetListenerPort(v int32) *DescribeRulesRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeRulesRequest) SetListenerProtocol(v string) *DescribeRulesRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeRulesRequest) SetLoadBalancerId(v string) *DescribeRulesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeRulesRequest) SetOwnerAccount(v string) *DescribeRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRulesRequest) SetOwnerId(v int64) *DescribeRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRulesRequest) SetRegionId(v string) *DescribeRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceOwnerAccount(v string) *DescribeRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceOwnerId(v int64) *DescribeRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRulesRequest) Validate() error {
	return dara.Validate(s)
}
