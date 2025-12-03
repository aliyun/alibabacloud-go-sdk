// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensionId(v string) *DescribeDomainExtensionsRequest
	GetDomainExtensionId() *string
	SetListenerPort(v int32) *DescribeDomainExtensionsRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeDomainExtensionsRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeDomainExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDomainExtensionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDomainExtensionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDomainExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDomainExtensionsRequest
	GetResourceOwnerId() *int64
}

type DescribeDomainExtensionsRequest struct {
	// The ID of the additional certificate.
	//
	// example:
	//
	// de-bp1rp7ta1****
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	// The frontend port of the HTTPS listener that is configured for the SLB instance. Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The SLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
}

func (s DescribeDomainExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionsRequest) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DescribeDomainExtensionsRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeDomainExtensionsRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeDomainExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDomainExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDomainExtensionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDomainExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDomainExtensionsRequest) SetDomainExtensionId(v string) *DescribeDomainExtensionsRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetListenerPort(v int32) *DescribeDomainExtensionsRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetLoadBalancerId(v string) *DescribeDomainExtensionsRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetOwnerAccount(v string) *DescribeDomainExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetOwnerId(v int64) *DescribeDomainExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetRegionId(v string) *DescribeDomainExtensionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetResourceOwnerAccount(v string) *DescribeDomainExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) SetResourceOwnerId(v int64) *DescribeDomainExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDomainExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
