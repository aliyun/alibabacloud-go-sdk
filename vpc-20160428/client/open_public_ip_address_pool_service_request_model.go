// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPublicIpAddressPoolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenPublicIpAddressPoolServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *OpenPublicIpAddressPoolServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenPublicIpAddressPoolServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenPublicIpAddressPoolServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenPublicIpAddressPoolServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenPublicIpAddressPoolServiceRequest
	GetResourceOwnerId() *int64
}

type OpenPublicIpAddressPoolServiceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655442455
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenPublicIpAddressPoolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenPublicIpAddressPoolServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenPublicIpAddressPoolServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetClientToken(v string) *OpenPublicIpAddressPoolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetOwnerAccount(v string) *OpenPublicIpAddressPoolServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetOwnerId(v int64) *OpenPublicIpAddressPoolServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetRegionId(v string) *OpenPublicIpAddressPoolServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetResourceOwnerAccount(v string) *OpenPublicIpAddressPoolServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) SetResourceOwnerId(v int64) *OpenPublicIpAddressPoolServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenPublicIpAddressPoolServiceRequest) Validate() error {
	return dara.Validate(s)
}
