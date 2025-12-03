// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AllocatePublicNetworkAddressRequest
	GetClusterId() *string
	SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type AllocatePublicNetworkAddressRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocatePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AllocatePublicNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocatePublicNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocatePublicNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocatePublicNetworkAddressRequest) SetClusterId(v string) *AllocatePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *AllocatePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocatePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
