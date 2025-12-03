// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeListenerAccessControlAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeListenerAccessControlAttributeRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *DescribeListenerAccessControlAttributeRequest
	GetListenerProtocol() *string
	SetLoadBalancerId(v string) *DescribeListenerAccessControlAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeListenerAccessControlAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeListenerAccessControlAttributeRequest struct {
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
	// > This parameter is required if the same port is specified for listeners of different protocols.
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
	// lb-8vb86hxixo8lvsja8****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the Classic Load Balancer (CLB) instance is created.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeListenerAccessControlAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeListenerAccessControlAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeListenerAccessControlAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeListenerAccessControlAttributeRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeListenerAccessControlAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeListenerAccessControlAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeListenerAccessControlAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeListenerAccessControlAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeListenerAccessControlAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeListenerAccessControlAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeListenerAccessControlAttributeRequest) SetListenerPort(v int32) *DescribeListenerAccessControlAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetListenerProtocol(v string) *DescribeListenerAccessControlAttributeRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetLoadBalancerId(v string) *DescribeListenerAccessControlAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetRegionId(v string) *DescribeListenerAccessControlAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerAccount(v string) *DescribeListenerAccessControlAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerId(v int64) *DescribeListenerAccessControlAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeListenerAccessControlAttributeRequest) Validate() error {
	return dara.Validate(s)
}
