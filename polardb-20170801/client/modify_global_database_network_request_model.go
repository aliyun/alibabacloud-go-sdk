// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalDatabaseNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableGlobalDomainName(v bool) *ModifyGlobalDatabaseNetworkRequest
	GetEnableGlobalDomainName() *bool
	SetGDNDescription(v string) *ModifyGlobalDatabaseNetworkRequest
	GetGDNDescription() *string
	SetGDNId(v string) *ModifyGlobalDatabaseNetworkRequest
	GetGDNId() *string
	SetOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyGlobalDatabaseNetworkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyGlobalDatabaseNetworkRequest
	GetSecurityToken() *string
}

type ModifyGlobalDatabaseNetworkRequest struct {
	// Create a global domain
	//
	// example:
	//
	// false
	EnableGlobalDomainName *bool `json:"EnableGlobalDomainName,omitempty" xml:"EnableGlobalDomainName,omitempty"`
	// The description of the GDN. The description must meet the following requirements:
	//
	// 	- The description cannot start with http:// or https://.
	//
	// 	- The description must start with a letter.
	//
	// 	- The description can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The description must be 2 to 126 characters in length.
	//
	// example:
	//
	// GDN-fortest
	GDNDescription *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	// The GDN ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId        *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGlobalDatabaseNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetEnableGlobalDomainName() *bool {
	return s.EnableGlobalDomainName
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetGDNDescription() *string {
	return s.GDNDescription
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetGDNId() *string {
	return s.GDNId
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalDatabaseNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetEnableGlobalDomainName(v bool) *ModifyGlobalDatabaseNetworkRequest {
	s.EnableGlobalDomainName = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetGDNDescription(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.GDNDescription = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetGDNId(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetResourceGroupId(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) Validate() error {
	return dara.Validate(s)
}
