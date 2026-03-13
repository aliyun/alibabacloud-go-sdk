// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateLoadBalancerRequest
	GetAddress() *string
	SetClientToken(v string) *CreateLoadBalancerRequest
	GetClientToken() *string
	SetIsPublicAddress(v string) *CreateLoadBalancerRequest
	GetIsPublicAddress() *string
	SetLoadBalancerMode(v string) *CreateLoadBalancerRequest
	GetLoadBalancerMode() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerRequest
	GetLoadBalancerName() *string
	SetOwnerAccount(v string) *CreateLoadBalancerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLoadBalancerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateLoadBalancerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerRequest
	GetResourceOwnerId() *int64
}

type CreateLoadBalancerRequest struct {
	Address              *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	IsPublicAddress      *string `json:"IsPublicAddress,omitempty" xml:"IsPublicAddress,omitempty"`
	LoadBalancerMode     *string `json:"LoadBalancerMode,omitempty" xml:"LoadBalancerMode,omitempty"`
	LoadBalancerName     *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLoadBalancerRequest) GetIsPublicAddress() *string {
	return s.IsPublicAddress
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerMode() *string {
	return s.LoadBalancerMode
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerRequest) SetAddress(v string) *CreateLoadBalancerRequest {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetIsPublicAddress(v string) *CreateLoadBalancerRequest {
	s.IsPublicAddress = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerMode(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerMode = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetOwnerId(v int64) *CreateLoadBalancerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
