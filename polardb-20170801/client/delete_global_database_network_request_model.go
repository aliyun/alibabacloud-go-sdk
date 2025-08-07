// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGDNId(v string) *DeleteGlobalDatabaseNetworkRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeleteGlobalDatabaseNetworkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type DeleteGlobalDatabaseNetworkRequest struct {
	// The ID of the GDN.
	//
	// This parameter is required.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId        *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetGDNId(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetResourceGroupId(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
