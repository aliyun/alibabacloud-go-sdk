// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicNetworkAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReleasePublicNetworkAddressRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest
	GetResourceOwnerId() *int64
}

type ReleasePublicNetworkAddressRequest struct {
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleasePublicNetworkAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicNetworkAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReleasePublicNetworkAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleasePublicNetworkAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleasePublicNetworkAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleasePublicNetworkAddressRequest) SetClusterId(v string) *ReleasePublicNetworkAddressRequest {
	s.ClusterId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerAccount(v string) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) SetResourceOwnerId(v int64) *ReleasePublicNetworkAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleasePublicNetworkAddressRequest) Validate() error {
	return dara.Validate(s)
}
